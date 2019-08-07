/*
 * MinIO Cloud Storage, (C) 2018-2019 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package openid

import (
	"crypto"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/minio/minio/cmd/config"
	"github.com/minio/minio/pkg/env"
	xnet "github.com/minio/minio/pkg/net"
)

// JWKSArgs - RSA authentication target arguments
type JWKSArgs struct {
	URL         *xnet.URL `json:"url"`
	publicKeys  map[string]crypto.PublicKey
	transport   *http.Transport
	closeRespFn func(io.ReadCloser)
}

// Config - OpenID Config
type Config struct {
	JWKS              JWKSArgs `json:"jwks"`
	ConfigURL         string   `json:"configURL"`
	PolicyClaimPrefix string   `json:"policyClaimPrefix"`
}

// PopulatePublicKey - populates a new publickey from the JWKS URL.
func (r *JWKSArgs) PopulatePublicKey() error {
	if r.URL == nil {
		return nil
	}
	client := &http.Client{}
	if r.transport != nil {
		client.Transport = r.transport
	}
	resp, err := client.Get(r.URL.String())
	if err != nil {
		return err
	}
	defer r.closeRespFn(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	var jwk JWKS
	if err = json.NewDecoder(resp.Body).Decode(&jwk); err != nil {
		return err
	}

	for _, key := range jwk.Keys {
		r.publicKeys[key.Kid], err = key.DecodePublicKey()
		if err != nil {
			return err
		}
	}

	return nil
}

// UnmarshalJSON - decodes JSON data.
func (r *JWKSArgs) UnmarshalJSON(data []byte) error {
	// subtype to avoid recursive call to UnmarshalJSON()
	type subJWKSArgs JWKSArgs
	var sr subJWKSArgs

	if err := json.Unmarshal(data, &sr); err != nil {
		return err
	}

	ar := JWKSArgs(sr)
	if ar.URL == nil || ar.URL.String() == "" {
		*r = ar
		return nil
	}

	*r = ar
	return nil
}

// JWT - rs client grants provider details.
type JWT struct {
	args JWKSArgs
}

func expToInt64(expI interface{}) (expAt int64, err error) {
	switch exp := expI.(type) {
	case float64:
		expAt = int64(exp)
	case int64:
		expAt = exp
	case json.Number:
		expAt, err = exp.Int64()
		if err != nil {
			return 0, err
		}
	default:
		return 0, ErrInvalidDuration
	}
	return expAt, nil
}

// GetDefaultExpiration - returns the expiration seconds expected.
func GetDefaultExpiration(dsecs string) (time.Duration, error) {
	defaultExpiryDuration := time.Duration(60) * time.Minute // Defaults to 1hr.
	if dsecs != "" {
		expirySecs, err := strconv.ParseInt(dsecs, 10, 64)
		if err != nil {
			return 0, ErrInvalidDuration
		}
		// The duration, in seconds, of the role session.
		// The value can range from 900 seconds (15 minutes)
		// to 12 hours.
		if expirySecs < 900 || expirySecs > 43200 {
			return 0, ErrInvalidDuration
		}

		defaultExpiryDuration = time.Duration(expirySecs) * time.Second
	}
	return defaultExpiryDuration, nil
}

// Validate - validates the access token.
func (p *JWT) Validate(token, dsecs string) (map[string]interface{}, error) {
	jp := new(jwtgo.Parser)
	jp.ValidMethods = []string{"RS256", "RS384", "RS512", "ES256", "ES384", "ES512"}

	keyFuncCallback := func(jwtToken *jwtgo.Token) (interface{}, error) {
		kid, ok := jwtToken.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("Invalid kid value %v", jwtToken.Header["kid"])
		}
		return p.args.publicKeys[kid], nil
	}

	var claims jwtgo.MapClaims
	jwtToken, err := jp.ParseWithClaims(token, &claims, keyFuncCallback)
	if err != nil {
		if err = p.args.PopulatePublicKey(); err != nil {
			return nil, err
		}
		jwtToken, err = jwtgo.ParseWithClaims(token, &claims, keyFuncCallback)
		if err != nil {
			return nil, err
		}
	}

	if !jwtToken.Valid {
		return nil, ErrTokenExpired
	}

	expAt, err := expToInt64(claims["exp"])
	if err != nil {
		return nil, err
	}

	defaultExpiryDuration, err := GetDefaultExpiration(dsecs)
	if err != nil {
		return nil, err
	}

	if time.Unix(expAt, 0).UTC().Sub(time.Now().UTC()) < defaultExpiryDuration {
		defaultExpiryDuration = time.Unix(expAt, 0).UTC().Sub(time.Now().UTC())
	}

	expiry := time.Now().UTC().Add(defaultExpiryDuration).Unix()
	if expAt < expiry {
		claims["exp"] = strconv.FormatInt(expAt, 64)
	}

	return claims, nil

}

// ID returns the provider name and authentication type.
func (p *JWT) ID() ID {
	return "jwt"
}

// OpenID keys and envs.
const (
	JwksURL           = "jwks_url"
	ConfigURL         = "config_url"
	PolicyClaimPrefix = "policy_claim_prefix"

	EnvIdentityOpenID                  = "MINIO_IDENTITY_OPENID_JWKS_URL"
	EnvIdentityOpenIDConfigURL         = "MINIO_IDENTITY_OPENID_CONFIG_URL"
	EnvIdentityOpenIDPolicyClaimPrefix = "MINIO_IDENTITY_OPENID_POLICY_CLAIM_PREFIX"
)

// LookupConfig lookup jwks from config, override with any ENVs.
func LookupConfig(kv config.KVS, transport *http.Transport, closeRespFn func(io.ReadCloser)) (c Config, err error) {
	c = Config{}

	jwksURL := env.Get(EnvIamJwksURL, "")
	if jwksURL == "" {
		jwksURL = env.Get(EnvIdentityOpenID, kv.Get(JwksURL))
	}

	if jwksURL != "" {
		u, err := xnet.ParseURL(jwksURL)
		if err != nil {
			return c, err
		}

		c.JWKS = JWKSArgs{
			URL:         u,
			publicKeys:  make(map[string]crypto.PublicKey),
			transport:   transport,
			closeRespFn: closeRespFn,
		}

		if err = c.JWKS.PopulatePublicKey(); err != nil {
			return c, err
		}
	}

	c.ConfigURL = env.Get(EnvIdentityOpenIDConfigURL, kv.Get(ConfigURL))
	c.PolicyClaimPrefix = env.Get(EnvIdentityOpenIDPolicyClaimPrefix, kv.Get(PolicyClaimPrefix))
	return c, nil
}

// NewJWT - initialize new jwt authenticator.
func NewJWT(args JWKSArgs) *JWT {
	return &JWT{
		args: args,
	}
}
