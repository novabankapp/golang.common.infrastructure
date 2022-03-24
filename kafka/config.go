package kafka

type Config struct {
	Brokers    []string    `mapstructure:"brokers"`
	GroupID    string      `mapstructure:"groupID"`
	InitTopics bool        `mapstructure:"initTopics"`
	Topics     KafkaTopics `mapstructure:"kafkaTopics"`
}

// TopicConfig kafka topic config
type TopicConfig struct {
	TopicName         string `mapstructure:"topicName"`
	Partitions        int    `mapstructure:"partitions"`
	ReplicationFactor int    `mapstructure:"replicationFactor"`
}

type KafkaTopics struct {
	UserCreate          TopicConfig `mapstructure:"userCreate"`
	UserCreated         TopicConfig `mapstructure:"userCreated"`
	UserUpdate          TopicConfig `mapstructure:"userUpdate"`
	UserUpdated         TopicConfig `mapstructure:"userUpdated"`
	UserDelete          TopicConfig `mapstructure:"userDelete"`
	UserDeleted         TopicConfig `mapstructure:"userDeleted"`
	UserPasswordChanged TopicConfig `mapstructure:"userDeleted"`
}
