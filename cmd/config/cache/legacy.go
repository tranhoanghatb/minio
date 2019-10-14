/*
 * MinIO Cloud Storage, (C) 2019 MinIO, Inc.
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

package cache

import (
	"fmt"
	"strings"

	"github.com/minio/minio/cmd/config"
)

// SetCacheConfig - One time migration code needed, for migrating from older config to new for Cache.
func SetCacheConfig(s config.Config, cfg Config) {
	s[config.CacheSubSys][config.Default] = config.KVS{
		Drives:  strings.Join(cfg.Drives, ","),
		Exclude: strings.Join(cfg.Exclude, ","),
		Expiry:  fmt.Sprintf("%d", cfg.Expiry),
		Quota:   fmt.Sprintf("%d", cfg.MaxUse),
		config.State: func() string {
			if len(cfg.Drives) > 0 {
				return config.StateOn
			}
			return config.StateOff
		}(),
	}
}
