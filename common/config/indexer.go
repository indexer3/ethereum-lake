package config

type IndexerConfig struct {
	IndexerName                string `mapstructure:"name" json:"name" yaml:"name"`
	StartCursor                uint64 `mapstructure:"start_cursor" json:"start_cursor" yaml:"start_cursor"`
	EndCursor                  uint64 `mapstructure:"end_cursor" json:"end_cursor" yaml:"end_cursor"`
	CatchUpSleepIntervalSecond uint64 `mapstructure:"catch_up_sleep_interval_second" json:"catch_up_sleep_interval_second" yaml:"catch_up_sleep_interval_second"`

	RPCConfig     RPCConfig     `mapstructure:"rpc" json:"rpc" yaml:"rpc"`
	MetricsConfig MetricsConfig `mapstructure:"metrics" json:"metrics" yaml:"metrics"`

	ClickHouseConfig ClickHouseConfig `mapstructure:"clickhouse" json:"clickhouse" yaml:"clickhouse"`
	MySQLConfig      MySQLConfig      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PostgresConfig   PostgresConfig   `mapstructure:"postgres" json:"postgres" yaml:"postgres"`
}

var IndexerConf IndexerConfig

type RPCConfig struct {
	RPCMaxRetry            uint64 `mapstructure:"rpc_max_retry" json:"rpc_max_retry" yaml:"rpc_max_retry"`
	RPCRetryIntervalSecond uint64 `mapstructure:"rpc_retry_interval_second" json:"rpc_retry_interval_second" yaml:"rpc_retry_interval_second"`
	RPCUrl                 string `mapstructure:"rpc_url" json:"rpc_url" yaml:"rpc_url"`
}

type MetricsConfig struct {
	Enable                bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	CollectURL            string `mapstructure:"collect_url" json:"collect_url" yaml:"collect_url"`
	CollectIntervalSecond uint64 `mapstructure:"collect_interval_second" json:"collect_interval_second" yaml:"collect_interval_second"`
}

type ClickHouseConfig struct {
	Enable         bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Host           string `mapstructure:"host" json:"host" yaml:"host"`
	Port           string `mapstructure:"port" json:"port" yaml:"port"`
	User           string `mapstructure:"user" json:"user" yaml:"user"`
	Password       string `mapstructure:"password" json:"password" yaml:"password"`
	DataBase       string `mapstructure:"database" json:"database" yaml:"database"`
	BatchWriteSize uint64 `mapstructure:"batch_write_size" json:"batch_write_size" yaml:"batch_write_size"`
}

type MySQLConfig struct {
	Enable         bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Host           string `mapstructure:"host" json:"host" yaml:"host"`
	Port           string `mapstructure:"port" json:"port" yaml:"port"`
	User           string `mapstructure:"user" json:"user" yaml:"user"`
	Password       string `mapstructure:"password" json:"password" yaml:"password"`
	DataBase       string `mapstructure:"database" json:"database" yaml:"database"`
	BatchWriteSize uint64 `mapstructure:"batch_write_size" json:"batch_write_size" yaml:"batch_write_size"`
}

type PostgresConfig struct {
	Enable         bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Host           string `mapstructure:"host" json:"host" yaml:"host"`
	Port           string `mapstructure:"port" json:"port" yaml:"port"`
	User           string `mapstructure:"user" json:"user" yaml:"user"`
	Password       string `mapstructure:"password" json:"password" yaml:"password"`
	DataBase       string `mapstructure:"database" json:"database" yaml:"database"`
	BatchWriteSize uint64 `mapstructure:"batch_write_size" json:"batch_write_size" yaml:"batch_write_size"`
}
