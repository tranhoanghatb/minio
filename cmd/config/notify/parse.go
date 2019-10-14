package notify

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio/cmd/config"
	"github.com/minio/minio/cmd/logger"
	"github.com/minio/minio/pkg/env"
	"github.com/minio/minio/pkg/event"
	"github.com/minio/minio/pkg/event/target"
	xnet "github.com/minio/minio/pkg/net"
)

// TestNotificationTargets is similar to GetNotificationTargets()
// avoids explicit registration.
func TestNotificationTargets(cfg config.Config, doneCh <-chan struct{}, rootCAs *x509.CertPool) error {
	_, err := RegisterNotificationTargets(cfg, doneCh, rootCAs, true)
	return err
}

// GetNotificationTargets registers and initializes all notification
// targets, returns error if any.
func GetNotificationTargets(cfg config.Config, doneCh <-chan struct{}, rootCAs *x509.CertPool) (*event.TargetList, error) {
	return RegisterNotificationTargets(cfg, doneCh, rootCAs, false)
}

// RegisterNotificationTargets - returns TargetList which contains enabled targets in serverConfig.
// A new notification target is added like below
// * Add a new target in pkg/event/target package.
// * Add newly added target configuration to serverConfig.Notify.<TARGET_NAME>.
// * Handle the configuration in this function to create/add into TargetList.
func RegisterNotificationTargets(cfg config.Config, doneCh <-chan struct{}, rootCAs *x509.CertPool, test bool) (*event.TargetList, error) {
	targetList := event.NewTargetList()
	for id, args := range GetNotifyAMQP(cfg) {
		if !args.Enable {
			continue
		}
		newTarget, err := target.NewAMQPTarget(id, args, doneCh, logger.LogOnceIf)
		if err != nil {
			return nil, err
		}
		if test {
			newTarget.Close()
			continue
		}
		if err = targetList.Add(newTarget); err != nil {
			return nil, err
		}
	}

	for id, args := range GetNotifyES(cfg) {
		if !args.Enable {
			continue
		}
		newTarget, err := target.NewElasticsearchTarget(id, args, doneCh, logger.LogOnceIf)
		if err != nil {
			return nil, err

		}
		if test {
			newTarget.Close()
			continue
		}
		if err = targetList.Add(newTarget); err != nil {
			return nil, err
		}
	}

	for id, args := range GetNotifyKafka(cfg) {
		if !args.Enable {
			continue
		}
		args.TLS.RootCAs = rootCAs
		newTarget, err := target.NewKafkaTarget(id, args, doneCh, logger.LogOnceIf)
		if err != nil {
			return nil, err
		}
		if test {
			newTarget.Close()
			continue
		}
		if err = targetList.Add(newTarget); err != nil {
			return nil, err
		}
	}

	for id, args := range GetNotifyMQTT(cfg, rootCAs) {
		if !args.Enable {
			continue
		}
		args.RootCAs = rootCAs
		newTarget, err := target.NewMQTTTarget(id, args, doneCh, logger.LogOnceIf)
		if err != nil {
			return nil, err
		}
		if test {
			newTarget.Close()
			continue
		}
		if err = targetList.Add(newTarget); err != nil {
			return nil, err
		}
	}

	for id, args := range GetNotifyMySQL(cfg) {
		if !args.Enable {
			continue
		}
		newTarget, err := target.NewMySQLTarget(id, args, doneCh, logger.LogOnceIf)
		if err != nil {
			return nil, err
		}
		if test {
			newTarget.Close()
			continue
		}
		if err = targetList.Add(newTarget); err != nil {
			return nil, err
		}
	}

	for id, args := range GetNotifyNATS(cfg) {
		if !args.Enable {
			continue
		}
		newTarget, err := target.NewNATSTarget(id, args, doneCh, logger.LogOnceIf)
		if err != nil {
			return nil, err
		}
		if test {
			newTarget.Close()
			continue
		}
		if err = targetList.Add(newTarget); err != nil {
			return nil, err
		}
	}

	for id, args := range GetNotifyNSQ(cfg) {
		if !args.Enable {
			continue
		}
		newTarget, err := target.NewNSQTarget(id, args, doneCh, logger.LogOnceIf)
		if err != nil {
			return nil, err
		}
		if test {
			newTarget.Close()
			continue
		}
		if err = targetList.Add(newTarget); err != nil {
			return nil, err
		}
	}

	for id, args := range GetNotifyPostgres(cfg) {
		if !args.Enable {
			continue
		}
		newTarget, err := target.NewPostgreSQLTarget(id, args, doneCh, logger.LogOnceIf)
		if err != nil {
			return nil, err
		}
		if test {
			newTarget.Close()
			continue
		}
		if err = targetList.Add(newTarget); err != nil {
			return nil, err
		}
	}

	for id, args := range GetNotifyRedis(cfg) {
		if !args.Enable {
			continue
		}
		newTarget, err := target.NewRedisTarget(id, args, doneCh, logger.LogOnceIf)
		if err != nil {
			return nil, err
		}
		if test {
			newTarget.Close()
			continue
		}
		if err = targetList.Add(newTarget); err != nil {
			return nil, err
		}
	}

	for id, args := range GetNotifyWebhook(cfg, rootCAs) {
		if !args.Enable {
			continue
		}
		args.RootCAs = rootCAs
		newTarget, err := target.NewWebhookTarget(id, args, doneCh, logger.LogOnceIf)
		if err != nil {
			return nil, err
		}
		if test {
			newTarget.Close()
			continue
		}
		if err := targetList.Add(newTarget); err != nil {
			return nil, err
		}
	}

	return targetList, nil
}

func mergeTargets(cfgTargets map[string]config.KVS, envname string, defaultKVS config.KVS) map[string]config.KVS {
	newCfgTargets := make(map[string]config.KVS)
	for _, e := range env.List(envname) {
		tgt := strings.TrimPrefix(e, envname+config.Default)
		if tgt == envname {
			tgt = config.Default
		}
		newCfgTargets[tgt] = defaultKVS
	}
	for tgt, kv := range cfgTargets {
		newCfgTargets[tgt] = kv
	}
	return newCfgTargets
}

// GetNotifyKafka - returns a map of registered notification 'kafka' targets
func GetNotifyKafka(s config.Config) map[string]target.KafkaArgs {
	kafkaTargets := make(map[string]target.KafkaArgs)
	defaultKVS := config.KVS{
		target.KafkaTLSClientAuth: "0",
		target.KafkaSASLEnable:    config.StateOff,
		target.KafkaTLSEnable:     config.StateOff,
		target.KafkaTLSSkipVerify: config.StateOff,
		target.KafkaQueueLimit:    "0",
		config.State:              config.StateOff,
	}
	for k, kv := range mergeTargets(s[config.NotifyKafkaSubSys], target.EnvKafkaState, defaultKVS) {
		stateEnv := target.EnvKafkaState
		if k != config.Default {
			stateEnv = stateEnv + config.Default + k
		}
		enabled := env.Get(stateEnv, kv.Get(config.State)) == config.StateOn
		if !enabled {
			continue
		}
		var err error
		var brokers []xnet.Host
		brokersEnv := target.EnvKafkaBrokers
		if k != config.Default {
			brokersEnv = brokersEnv + config.Default + k
		}
		kafkaBrokers := env.Get(brokersEnv, kv.Get(target.KafkaBrokers))
		for _, s := range strings.Split(kafkaBrokers, config.ValueSeparator) {
			var host *xnet.Host
			host, err = xnet.ParseHost(s)
			if err != nil {
				break
			}
			brokers = append(brokers, *host)
		}
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		queueLimitEnv := target.EnvKafkaQueueLimit
		if k != config.Default {
			queueLimitEnv = queueLimitEnv + config.Default + k
		}
		queueLimit, err := strconv.ParseUint(env.Get(queueLimitEnv, kv.Get(target.KafkaQueueLimit)), 10, 64)
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		clientAuthEnv := target.EnvKafkaTLSClientAuth
		if k != config.Default {
			clientAuthEnv = clientAuthEnv + config.Default + k
		}
		clientAuth, err := strconv.Atoi(env.Get(clientAuthEnv, kv.Get(target.KafkaTLSClientAuth)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		topicEnv := target.EnvKafkaTopic
		if k != config.Default {
			topicEnv = topicEnv + config.Default + k
		}

		queueDirEnv := target.EnvKafkaQueueDir
		if k != config.Default {
			queueDirEnv = queueDirEnv + config.Default + k
		}

		kafkaArgs := target.KafkaArgs{
			Enable:     enabled,
			Brokers:    brokers,
			Topic:      env.Get(topicEnv, kv.Get(target.KafkaTopic)),
			QueueDir:   env.Get(queueDirEnv, kv.Get(target.KafkaQueueDir)),
			QueueLimit: queueLimit,
		}

		tlsEnableEnv := target.EnvKafkaTLSEnable
		if k != config.Default {
			tlsEnableEnv = tlsEnableEnv + config.Default + k
		}
		tlsSkipVerifyEnv := target.EnvKafkaTLSSkipVerify
		if k != config.Default {
			tlsSkipVerifyEnv = tlsSkipVerifyEnv + config.Default + k
		}
		kafkaArgs.TLS.Enable = env.Get(tlsEnableEnv, kv.Get(target.KafkaTLSEnable)) == config.StateOn
		kafkaArgs.TLS.SkipVerify = env.Get(tlsSkipVerifyEnv, kv.Get(target.KafkaTLSSkipVerify)) == config.StateOn
		kafkaArgs.TLS.ClientAuth = tls.ClientAuthType(clientAuth)

		saslEnableEnv := target.EnvKafkaSASLEnable
		if k != config.Default {
			saslEnableEnv = saslEnableEnv + config.Default + k
		}
		saslUsernameEnv := target.EnvKafkaSASLUsername
		if k != config.Default {
			saslUsernameEnv = saslUsernameEnv + config.Default + k
		}
		saslPasswordEnv := target.EnvKafkaSASLPassword
		if k != config.Default {
			saslPasswordEnv = saslPasswordEnv + config.Default + k
		}
		kafkaArgs.SASL.Enable = env.Get(saslEnableEnv, kv.Get(target.KafkaSASLEnable)) == config.StateOn
		kafkaArgs.SASL.User = env.Get(saslUsernameEnv, kv.Get(target.KafkaSASLUsername))
		kafkaArgs.SASL.Password = env.Get(saslPasswordEnv, kv.Get(target.KafkaSASLPassword))

		if err = kafkaArgs.Validate(); err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		kafkaTargets[k] = kafkaArgs
	}

	return kafkaTargets
}

// GetNotifyMQTT - returns a map of registered notification 'mqtt' targets
func GetNotifyMQTT(s config.Config, rootCAs *x509.CertPool) map[string]target.MQTTArgs {
	mqttTargets := make(map[string]target.MQTTArgs)
	defaultKVS := config.KVS{
		config.State:                 config.StateOff,
		target.MqttKeepAliveInterval: "0s",
		target.MqttReconnectInterval: "0s",
		target.MqttQoS:               "0",
		target.MqttQueueLimit:        "0",
	}
	for k, kv := range mergeTargets(s[config.NotifyMQTTSubSys], target.EnvMQTTState, defaultKVS) {
		stateEnv := target.EnvMQTTState
		if k != config.Default {
			stateEnv = stateEnv + config.Default + k
		}

		enabled := env.Get(stateEnv, kv.Get(config.State)) == config.StateOn
		if !enabled {
			continue
		}

		brokerEnv := target.EnvMQTTBroker
		if k != config.Default {
			brokerEnv = brokerEnv + config.Default + k
		}
		brokerURL, err := xnet.ParseURL(env.Get(brokerEnv, kv.Get(target.MqttBroker)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		reconnectIntervalEnv := target.EnvMQTTReconnectInterval
		if k != config.Default {
			reconnectIntervalEnv = reconnectIntervalEnv + config.Default + k
		}
		reconnectInterval, err := time.ParseDuration(env.Get(reconnectIntervalEnv,
			kv.Get(target.MqttReconnectInterval)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		keepAliveIntervalEnv := target.EnvMQTTKeepAliveInterval
		if k != config.Default {
			keepAliveIntervalEnv = keepAliveIntervalEnv + config.Default + k
		}
		keepAliveInterval, err := time.ParseDuration(env.Get(keepAliveIntervalEnv,
			kv.Get(target.MqttKeepAliveInterval)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		queueLimitEnv := target.EnvMQTTQueueLimit
		if k != config.Default {
			queueLimitEnv = queueLimitEnv + config.Default + k
		}
		queueLimit, err := strconv.ParseUint(env.Get(queueLimitEnv, kv.Get(target.MqttQueueLimit)), 10, 64)
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		qosEnv := target.EnvMQTTQoS
		if k != config.Default {
			qosEnv = qosEnv + config.Default + k
		}

		// Parse uint8 value
		qos, err := strconv.ParseUint(env.Get(qosEnv, kv.Get(target.MqttQoS)), 10, 8)
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		topicEnv := target.EnvMQTTTopic
		if k != config.Default {
			topicEnv = topicEnv + config.Default + k
		}

		usernameEnv := target.EnvMQTTUsername
		if k != config.Default {
			usernameEnv = usernameEnv + config.Default + k
		}

		passwordEnv := target.EnvMQTTPassword
		if k != config.Default {
			passwordEnv = passwordEnv + config.Default + k
		}

		queueDirEnv := target.EnvMQTTQueueDir
		if k != config.Default {
			queueDirEnv = queueDirEnv + config.Default + k
		}

		mqttArgs := target.MQTTArgs{
			Enable:               enabled,
			Broker:               *brokerURL,
			Topic:                env.Get(topicEnv, kv.Get(target.MqttTopic)),
			QoS:                  byte(qos),
			User:                 env.Get(usernameEnv, kv.Get(target.MqttUsername)),
			Password:             env.Get(passwordEnv, kv.Get(target.MqttPassword)),
			MaxReconnectInterval: reconnectInterval,
			KeepAlive:            keepAliveInterval,
			RootCAs:              rootCAs,
			QueueDir:             env.Get(queueDirEnv, kv.Get(target.MqttQueueDir)),
			QueueLimit:           queueLimit,
		}

		if err = mqttArgs.Validate(); err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		mqttTargets[k] = mqttArgs
	}
	return mqttTargets
}

// GetNotifyMySQL - returns a map of registered notification 'mysql' targets
func GetNotifyMySQL(s config.Config) map[string]target.MySQLArgs {
	mysqlTargets := make(map[string]target.MySQLArgs)
	defaultKVS := config.KVS{
		config.State:           config.StateOff,
		target.MySQLQueueLimit: "0",
	}
	for k, kv := range mergeTargets(s[config.NotifyMySQLSubSys], target.EnvMySQLState, defaultKVS) {
		stateEnv := target.EnvMySQLState
		if k != config.Default {
			stateEnv = stateEnv + config.Default + k
		}

		enabled := env.Get(stateEnv, kv.Get(config.State)) == config.StateOn
		if !enabled {
			continue
		}

		hostEnv := target.EnvMySQLHost
		if k != config.Default {
			hostEnv = hostEnv + config.Default + k
		}

		host, err := xnet.ParseURL(env.Get(hostEnv, kv.Get(target.MySQLHost)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		queueLimitEnv := target.EnvMySQLQueueLimit
		if k != config.Default {
			queueLimitEnv = queueLimitEnv + config.Default + k
		}
		queueLimit, err := strconv.ParseUint(env.Get(queueLimitEnv, kv.Get(target.MySQLQueueLimit)), 10, 64)
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		formatEnv := target.EnvMySQLFormat
		if k != config.Default {
			formatEnv = formatEnv + config.Default + k
		}
		dsnStringEnv := target.EnvMySQLDSNString
		if k != config.Default {
			dsnStringEnv = dsnStringEnv + config.Default + k
		}
		tableEnv := target.EnvMySQLTable
		if k != config.Default {
			tableEnv = tableEnv + config.Default + k
		}
		portEnv := target.EnvMySQLPort
		if k != config.Default {
			portEnv = portEnv + config.Default + k
		}
		usernameEnv := target.EnvMySQLUsername
		if k != config.Default {
			usernameEnv = usernameEnv + config.Default + k
		}
		passwordEnv := target.EnvMySQLPassword
		if k != config.Default {
			passwordEnv = passwordEnv + config.Default + k
		}
		databaseEnv := target.EnvMySQLDatabase
		if k != config.Default {
			databaseEnv = databaseEnv + config.Default + k
		}
		queueDirEnv := target.EnvMySQLQueueDir
		if k != config.Default {
			queueDirEnv = queueDirEnv + config.Default + k
		}
		mysqlArgs := target.MySQLArgs{
			Enable:     enabled,
			Format:     env.Get(formatEnv, kv.Get(target.MySQLFormat)),
			DSN:        env.Get(dsnStringEnv, kv.Get(target.MySQLDSNString)),
			Table:      env.Get(tableEnv, kv.Get(target.MySQLTable)),
			Host:       *host,
			Port:       env.Get(portEnv, kv.Get(target.MySQLPort)),
			User:       env.Get(usernameEnv, kv.Get(target.MySQLUsername)),
			Password:   env.Get(passwordEnv, kv.Get(target.MySQLPassword)),
			Database:   env.Get(databaseEnv, kv.Get(target.MySQLDatabase)),
			QueueDir:   env.Get(queueDirEnv, kv.Get(target.MySQLQueueDir)),
			QueueLimit: queueLimit,
		}
		if err = mysqlArgs.Validate(); err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		mysqlTargets[k] = mysqlArgs
	}
	return mysqlTargets
}

// GetNotifyNATS - returns a map of registered notification 'nats' targets
func GetNotifyNATS(s config.Config) map[string]target.NATSArgs {
	natsTargets := make(map[string]target.NATSArgs)
	defaultKVS := config.KVS{
		config.State:                           config.StateOff,
		target.NATSSecure:                      config.StateOff,
		target.NATSPingInterval:                "0",
		target.NATSQueueLimit:                  "0",
		target.NATSStreamingEnable:             config.StateOff,
		target.NATSStreamingAsync:              config.StateOff,
		target.NATSStreamingMaxPubAcksInFlight: "0",
	}
	for k, kv := range mergeTargets(s[config.NotifyNATSSubSys], target.EnvNATSState, defaultKVS) {
		stateEnv := target.EnvNATSState
		if k != config.Default {
			stateEnv = stateEnv + config.Default + k
		}
		enabled := env.Get(stateEnv, kv.Get(config.State)) == config.StateOn
		if !enabled {
			continue
		}

		addressEnv := target.EnvNATSAddress
		if k != config.Default {
			addressEnv = addressEnv + config.Default + k
		}

		address, err := xnet.ParseHost(env.Get(addressEnv, kv.Get(target.NATSAddress)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		pingIntervalEnv := target.EnvNATSPingInterval
		if k != config.Default {
			pingIntervalEnv = pingIntervalEnv + config.Default + k
		}

		pingInterval, err := strconv.ParseInt(env.Get(pingIntervalEnv, kv.Get(target.NATSPingInterval)), 10, 64)
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		queueLimitEnv := target.EnvNATSQueueLimit
		if k != config.Default {
			queueLimitEnv = queueLimitEnv + config.Default + k
		}

		queueLimit, err := strconv.ParseUint(env.Get(queueLimitEnv, kv.Get(target.NATSQueueLimit)), 10, 64)
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		secureEnv := target.EnvNATSSecure
		if k != config.Default {
			secureEnv = secureEnv + config.Default + k
		}

		subjectEnv := target.EnvNATSSubject
		if k != config.Default {
			subjectEnv = subjectEnv + config.Default + k
		}

		usernameEnv := target.EnvNATSUsername
		if k != config.Default {
			usernameEnv = usernameEnv + config.Default + k
		}

		passwordEnv := target.EnvNATSPassword
		if k != config.Default {
			passwordEnv = passwordEnv + config.Default + k
		}

		tokenEnv := target.EnvNATSToken
		if k != config.Default {
			tokenEnv = tokenEnv + config.Default + k
		}

		queueDirEnv := target.EnvNATSQueueDir
		if k != config.Default {
			queueDirEnv = queueDirEnv + config.Default + k
		}

		natsArgs := target.NATSArgs{
			Enable:       true,
			Address:      *address,
			Subject:      env.Get(subjectEnv, kv.Get(target.NATSSubject)),
			Username:     env.Get(usernameEnv, kv.Get(target.NATSUsername)),
			Password:     env.Get(passwordEnv, kv.Get(target.NATSPassword)),
			Token:        env.Get(tokenEnv, kv.Get(target.NATSToken)),
			Secure:       env.Get(secureEnv, kv.Get(target.NATSSecure)) == config.StateOn,
			PingInterval: pingInterval,
			QueueDir:     env.Get(queueDirEnv, kv.Get(target.NATSQueueDir)),
			QueueLimit:   queueLimit,
		}

		streamingEnableEnv := target.EnvNATSStreamingEnable
		if k != config.Default {
			streamingEnableEnv = streamingEnableEnv + config.Default + k
		}

		streamingEnabled := env.Get(streamingEnableEnv, kv.Get(target.NATSStreamingEnable)) == config.StateOn
		if streamingEnabled {
			asyncEnv := target.EnvNATSStreamingAsync
			if k != config.Default {
				asyncEnv = asyncEnv + config.Default + k
			}
			maxPubAcksInflightEnv := target.EnvNATSStreamingMaxPubAcksInFlight
			if k != config.Default {
				maxPubAcksInflightEnv = maxPubAcksInflightEnv + config.Default + k
			}
			maxPubAcksInflight, err := strconv.Atoi(env.Get(maxPubAcksInflightEnv,
				kv.Get(target.NATSStreamingMaxPubAcksInFlight)))
			if err != nil {
				logger.LogIf(context.Background(), err)
				continue
			}
			clusterIDEnv := target.EnvNATSStreamingClusterID
			if k != config.Default {
				clusterIDEnv = clusterIDEnv + config.Default + k
			}
			natsArgs.Streaming.Enable = streamingEnabled
			natsArgs.Streaming.ClusterID = env.Get(clusterIDEnv, kv.Get(target.NATSStreamingClusterID))
			natsArgs.Streaming.Async = env.Get(asyncEnv, kv.Get(target.NATSStreamingAsync)) == config.StateOn
			natsArgs.Streaming.MaxPubAcksInflight = maxPubAcksInflight
		}

		if err = natsArgs.Validate(); err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		natsTargets[k] = natsArgs
	}
	return natsTargets
}

// GetNotifyNSQ - returns a map of registered notification 'nsq' targets
func GetNotifyNSQ(s config.Config) map[string]target.NSQArgs {
	nsqTargets := make(map[string]target.NSQArgs)
	defaultKVS := config.KVS{
		config.State:            config.StateOff,
		target.NSQTLSEnable:     config.StateOff,
		target.NSQTLSSkipVerify: config.StateOff,
		target.NSQQueueLimit:    "0",
	}

	for k, kv := range mergeTargets(s[config.NotifyNSQSubSys], target.EnvNSQState, defaultKVS) {
		stateEnv := target.EnvNSQState
		if k != config.Default {
			stateEnv = stateEnv + config.Default + k
		}
		enabled := env.Get(stateEnv, kv.Get(config.State)) == config.StateOn
		if !enabled {
			continue
		}
		addressEnv := target.EnvNSQAddress
		if k != config.Default {
			addressEnv = addressEnv + config.Default + k
		}
		nsqdAddress, err := xnet.ParseHost(env.Get(addressEnv, kv.Get(target.NSQAddress)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		tlsEnableEnv := target.EnvNSQTLSEnable
		if k != config.Default {
			tlsEnableEnv = tlsEnableEnv + config.Default + k
		}
		tlsSkipVerifyEnv := target.EnvNSQTLSSkipVerify
		if k != config.Default {
			tlsSkipVerifyEnv = tlsSkipVerifyEnv + config.Default + k
		}

		queueLimitEnv := target.EnvNSQQueueLimit
		if k != config.Default {
			queueLimitEnv = queueLimitEnv + config.Default + k
		}
		queueLimit, err := strconv.ParseUint(env.Get(queueLimitEnv, kv.Get(target.NSQQueueLimit)), 10, 64)
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		topicEnv := target.EnvNSQTopic
		if k != config.Default {
			topicEnv = topicEnv + config.Default + k
		}
		queueDirEnv := target.EnvNSQQueueDir
		if k != config.Default {
			queueDirEnv = queueDirEnv + config.Default + k
		}

		nsqArgs := target.NSQArgs{
			Enable:      enabled,
			NSQDAddress: *nsqdAddress,
			Topic:       env.Get(topicEnv, kv.Get(target.NSQTopic)),
			QueueDir:    env.Get(queueDirEnv, kv.Get(target.NSQQueueDir)),
			QueueLimit:  queueLimit,
		}
		nsqArgs.TLS.Enable = env.Get(tlsEnableEnv, kv.Get(target.NSQTLSEnable)) == config.StateOn
		nsqArgs.TLS.SkipVerify = env.Get(tlsSkipVerifyEnv, kv.Get(target.NSQTLSSkipVerify)) == config.StateOn

		if err = nsqArgs.Validate(); err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		nsqTargets[k] = nsqArgs
	}
	return nsqTargets
}

// GetNotifyPostgres - returns a map of registered notification 'postgres' targets
func GetNotifyPostgres(s config.Config) map[string]target.PostgreSQLArgs {
	psqlTargets := make(map[string]target.PostgreSQLArgs)
	defaultKVS := config.KVS{
		config.State:              config.StateOff,
		target.PostgresQueueLimit: "0",
	}
	for k, kv := range mergeTargets(s[config.NotifyPostgresSubSys], target.EnvPostgresState, defaultKVS) {
		stateEnv := target.EnvPostgresState
		if k != config.Default {
			stateEnv = stateEnv + config.Default + k
		}
		enabled := env.Get(stateEnv, kv.Get(config.State)) == config.StateOn
		if !enabled {
			continue
		}

		hostEnv := target.EnvPostgresHost
		if k != config.Default {
			hostEnv = hostEnv + config.Default + k
		}

		host, err := xnet.ParseHost(env.Get(hostEnv, kv.Get(target.PostgresHost)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		queueLimitEnv := target.EnvPostgresQueueLimit
		if k != config.Default {
			queueLimitEnv = queueLimitEnv + config.Default + k
		}

		queueLimit, err := strconv.Atoi(env.Get(queueLimitEnv, kv.Get(target.PostgresQueueLimit)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		formatEnv := target.EnvPostgresFormat
		if k != config.Default {
			formatEnv = formatEnv + config.Default + k
		}

		connectionStringEnv := target.EnvPostgresConnectionString
		if k != config.Default {
			connectionStringEnv = connectionStringEnv + config.Default + k
		}

		tableEnv := target.EnvPostgresTable
		if k != config.Default {
			tableEnv = tableEnv + config.Default + k
		}

		portEnv := target.EnvPostgresPort
		if k != config.Default {
			portEnv = portEnv + config.Default + k
		}

		usernameEnv := target.EnvPostgresUsername
		if k != config.Default {
			usernameEnv = usernameEnv + config.Default + k
		}

		passwordEnv := target.EnvPostgresPassword
		if k != config.Default {
			passwordEnv = passwordEnv + config.Default + k
		}

		databaseEnv := target.EnvPostgresDatabase
		if k != config.Default {
			databaseEnv = databaseEnv + config.Default + k
		}

		queueDirEnv := target.EnvPostgresQueueDir
		if k != config.Default {
			queueDirEnv = queueDirEnv + config.Default + k
		}

		psqlArgs := target.PostgreSQLArgs{
			Enable:           enabled,
			Format:           env.Get(formatEnv, kv.Get(target.PostgresFormat)),
			ConnectionString: env.Get(connectionStringEnv, kv.Get(target.PostgresConnectionString)),
			Table:            env.Get(tableEnv, kv.Get(target.PostgresTable)),
			Host:             *host,
			Port:             env.Get(portEnv, kv.Get(target.PostgresPort)),
			User:             env.Get(usernameEnv, kv.Get(target.PostgresUsername)),
			Password:         env.Get(passwordEnv, kv.Get(target.PostgresPassword)),
			Database:         env.Get(databaseEnv, kv.Get(target.PostgresDatabase)),
			QueueDir:         env.Get(queueDirEnv, kv.Get(target.PostgresQueueDir)),
			QueueLimit:       uint64(queueLimit),
		}
		if err = psqlArgs.Validate(); err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		psqlTargets[k] = psqlArgs
	}
	return psqlTargets
}

// GetNotifyRedis - returns a map of registered notification 'redis' targets
func GetNotifyRedis(s config.Config) map[string]target.RedisArgs {
	redisTargets := make(map[string]target.RedisArgs)
	defaultKVS := config.KVS{
		config.State:           config.StateOff,
		target.RedisQueueLimit: "0",
	}
	for k, kv := range mergeTargets(s[config.NotifyRedisSubSys], target.EnvRedisState, defaultKVS) {
		stateEnv := target.EnvRedisState
		if k != config.Default {
			stateEnv = stateEnv + config.Default + k
		}
		enabled := env.Get(stateEnv, kv.Get(config.State)) == config.StateOn
		if !enabled {
			continue
		}
		addressEnv := target.EnvRedisAddress
		if k != config.Default {
			addressEnv = addressEnv + config.Default + k
		}
		addr, err := xnet.ParseHost(env.Get(addressEnv, kv.Get(target.RedisAddress)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		queueLimitEnv := target.EnvRedisQueueLimit
		if k != config.Default {
			queueLimitEnv = queueLimitEnv + config.Default + k
		}
		queueLimit, err := strconv.Atoi(env.Get(queueLimitEnv, kv.Get(target.RedisQueueLimit)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		formatEnv := target.EnvRedisFormat
		if k != config.Default {
			formatEnv = formatEnv + config.Default + k
		}
		passwordEnv := target.EnvRedisPassword
		if k != config.Default {
			passwordEnv = passwordEnv + config.Default + k
		}
		keyEnv := target.EnvRedisKey
		if k != config.Default {
			keyEnv = keyEnv + config.Default + k
		}
		queueDirEnv := target.EnvRedisQueueDir
		if k != config.Default {
			queueDirEnv = queueDirEnv + config.Default + k
		}
		redisArgs := target.RedisArgs{
			Enable:     enabled,
			Format:     env.Get(formatEnv, kv.Get(target.RedisFormat)),
			Addr:       *addr,
			Password:   env.Get(passwordEnv, kv.Get(target.RedisPassword)),
			Key:        env.Get(keyEnv, kv.Get(target.RedisKey)),
			QueueDir:   env.Get(queueDirEnv, kv.Get(target.RedisQueueDir)),
			QueueLimit: uint64(queueLimit),
		}
		if err = redisArgs.Validate(); err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		redisTargets[k] = redisArgs
	}
	return redisTargets
}

// GetNotifyWebhook - returns a map of registered notification 'webhook' targets
func GetNotifyWebhook(s config.Config, rootCAs *x509.CertPool) map[string]target.WebhookArgs {
	webhookTargets := make(map[string]target.WebhookArgs)
	defaultKVS := config.KVS{
		config.State:             config.StateOff,
		target.WebhookQueueLimit: "0",
	}
	for k, kv := range mergeTargets(s[config.NotifyWebhookSubSys], target.EnvWebhookState, defaultKVS) {
		stateEnv := target.EnvWebhookState
		if k != config.Default {
			stateEnv = stateEnv + config.Default + k
		}
		enabled := env.Get(stateEnv, kv.Get(config.State)) == config.StateOn
		if !enabled {
			continue
		}
		urlEnv := target.EnvWebhookEndpoint
		if k != config.Default {
			urlEnv = urlEnv + config.Default + k
		}
		url, err := xnet.ParseURL(env.Get(urlEnv, kv.Get(target.WebhookEndpoint)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		queueLimitEnv := target.EnvWebhookQueueLimit
		if k != config.Default {
			queueLimitEnv = queueLimitEnv + config.Default + k
		}
		queueLimit, err := strconv.Atoi(env.Get(queueLimitEnv, kv.Get(target.WebhookQueueLimit)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		queueDirEnv := target.EnvWebhookQueueDir
		if k != config.Default {
			queueDirEnv = queueDirEnv + config.Default + k
		}

		webhookArgs := target.WebhookArgs{
			Enable:     enabled,
			Endpoint:   *url,
			RootCAs:    rootCAs,
			QueueDir:   env.Get(queueDirEnv, kv.Get(target.WebhookQueueDir)),
			QueueLimit: uint64(queueLimit),
		}
		if err = webhookArgs.Validate(); err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		webhookTargets[k] = webhookArgs
	}
	return webhookTargets
}

// GetNotifyES - returns a map of registered notification 'elasticsearch' targets
func GetNotifyES(s config.Config) map[string]target.ElasticsearchArgs {
	esTargets := make(map[string]target.ElasticsearchArgs)
	defaultKVS := config.KVS{
		config.State:             config.StateOff,
		target.ElasticQueueLimit: "0",
	}
	for k, kv := range mergeTargets(s[config.NotifyESSubSys], target.EnvElasticState, defaultKVS) {
		stateEnv := target.EnvElasticState
		if k != config.Default {
			stateEnv = stateEnv + config.Default + k
		}
		enabled := env.Get(stateEnv, kv.Get(config.State)) == config.StateOn
		if !enabled {
			continue
		}

		urlEnv := target.EnvElasticURL
		if k != config.Default {
			urlEnv = urlEnv + config.Default + k
		}

		url, err := xnet.ParseURL(env.Get(urlEnv, kv.Get(target.ElasticURL)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		queueLimitEnv := target.EnvElasticQueueLimit
		if k != config.Default {
			queueLimitEnv = queueLimitEnv + config.Default + k
		}

		queueLimit, err := strconv.Atoi(env.Get(queueLimitEnv, kv.Get(target.ElasticQueueLimit)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}

		formatEnv := target.EnvElasticFormat
		if k != config.Default {
			formatEnv = formatEnv + config.Default + k
		}

		indexEnv := target.EnvElasticIndex
		if k != config.Default {
			indexEnv = indexEnv + config.Default + k
		}

		queueDirEnv := target.EnvElasticQueueDir
		if k != config.Default {
			queueDirEnv = queueDirEnv + config.Default + k
		}

		esArgs := target.ElasticsearchArgs{
			Enable:     enabled,
			Format:     env.Get(formatEnv, kv.Get(target.ElasticFormat)),
			URL:        *url,
			Index:      env.Get(indexEnv, kv.Get(target.ElasticIndex)),
			QueueDir:   env.Get(queueDirEnv, kv.Get(target.ElasticQueueDir)),
			QueueLimit: uint64(queueLimit),
		}
		if err = esArgs.Validate(); err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		esTargets[k] = esArgs
	}
	return esTargets
}

// GetNotifyAMQP - returns a map of registered notification 'amqp' targets
func GetNotifyAMQP(s config.Config) map[string]target.AMQPArgs {
	amqpTargets := make(map[string]target.AMQPArgs)
	defaultKVS := config.KVS{
		config.State:            config.StateOff,
		target.AmqpDeliveryMode: "0",
	}
	for k, kv := range mergeTargets(s[config.NotifyAMQPSubSys], target.EnvAMQPState, defaultKVS) {
		stateEnv := target.EnvAMQPState
		if k != config.Default {
			stateEnv = stateEnv + config.Default + k
		}
		enabled := env.Get(stateEnv, kv.Get(config.State)) == config.StateOn
		if !enabled {
			continue
		}
		urlEnv := target.EnvAMQPURL
		if k != config.Default {
			urlEnv = urlEnv + config.Default + k
		}
		url, err := xnet.ParseURL(env.Get(urlEnv, kv.Get(target.AmqpURL)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		deliveryModeEnv := target.EnvAMQPDeliveryMode
		if k != config.Default {
			deliveryModeEnv = deliveryModeEnv + config.Default + k
		}
		deliveryMode, err := strconv.Atoi(env.Get(deliveryModeEnv, kv.Get(target.AmqpDeliveryMode)))
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		exchangeEnv := target.EnvAMQPExchange
		if k != config.Default {
			exchangeEnv = exchangeEnv + config.Default + k
		}
		routingKeyEnv := target.EnvAMQPRoutingKey
		if k != config.Default {
			routingKeyEnv = routingKeyEnv + config.Default + k
		}
		exchangeTypeEnv := target.EnvAMQPExchangeType
		if k != config.Default {
			exchangeTypeEnv = exchangeTypeEnv + config.Default + k
		}
		mandatoryEnv := target.EnvAMQPMandatory
		if k != config.Default {
			mandatoryEnv = mandatoryEnv + config.Default + k
		}
		immediateEnv := target.EnvAMQPImmediate
		if k != config.Default {
			immediateEnv = immediateEnv + config.Default + k
		}
		durableEnv := target.EnvAMQPDurable
		if k != config.Default {
			durableEnv = durableEnv + config.Default + k
		}
		internalEnv := target.EnvAMQPInternal
		if k != config.Default {
			internalEnv = internalEnv + config.Default + k
		}
		noWaitEnv := target.EnvAMQPNoWait
		if k != config.Default {
			noWaitEnv = noWaitEnv + config.Default + k
		}
		autoDeletedEnv := target.EnvAMQPAutoDeleted
		if k != config.Default {
			autoDeletedEnv = autoDeletedEnv + config.Default + k
		}
		queueDirEnv := target.EnvAMQPQueueDir
		if k != config.Default {
			queueDirEnv = queueDirEnv + config.Default + k
		}
		queueLimitEnv := target.EnvAMQPQueueLimit
		if k != config.Default {
			queueLimitEnv = queueLimitEnv + config.Default + k
		}
		queueLimit, err := strconv.ParseUint(env.Get(queueLimitEnv, kv.Get(target.AmqpQueueLimit)), 10, 64)
		if err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		amqpArgs := target.AMQPArgs{
			Enable:       enabled,
			URL:          *url,
			Exchange:     env.Get(exchangeEnv, kv.Get(target.AmqpExchange)),
			RoutingKey:   env.Get(routingKeyEnv, kv.Get(target.AmqpRoutingKey)),
			ExchangeType: env.Get(exchangeTypeEnv, kv.Get(target.AmqpExchangeType)),
			DeliveryMode: uint8(deliveryMode),
			Mandatory:    env.Get(mandatoryEnv, kv.Get(target.AmqpMandatory)) == config.StateOn,
			Immediate:    env.Get(immediateEnv, kv.Get(target.AmqpImmediate)) == config.StateOn,
			Durable:      env.Get(durableEnv, kv.Get(target.AmqpDurable)) == config.StateOn,
			Internal:     env.Get(internalEnv, kv.Get(target.AmqpInternal)) == config.StateOn,
			NoWait:       env.Get(noWaitEnv, kv.Get(target.AmqpNoWait)) == config.StateOn,
			AutoDeleted:  env.Get(autoDeletedEnv, kv.Get(target.AmqpAutoDeleted)) == config.StateOn,
			QueueDir:     env.Get(queueDirEnv, kv.Get(target.AmqpQueueDir)),
			QueueLimit:   queueLimit,
		}
		if err = amqpArgs.Validate(); err != nil {
			logger.LogIf(context.Background(), err)
			continue
		}
		amqpTargets[k] = amqpArgs
	}
	return amqpTargets
}
