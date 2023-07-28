package config

const (
	RPCEndpoints = "rpc-endpoints"

	ClickHouseHost           = "clickhouse-host"
	ClickHousePort           = "clickhouse-port"
	ClickHouseUser           = "clickhouse-user"
	ClickHousePassword       = "clickhouse-password"
	ClickHouseBatchWriteSize = "clickhouse-batch-write-size"

	MySQLHost           = "mysql-host"
	MySQLPort           = "mysql-port"
	MySQLUser           = "mysql-user"
	MySQLPassword       = "mysql-password"
	MySQLBatchWriteSize = "mysql-batch-write-size"

	PostgresHost           = "postgres-host"
	PostgresPort           = "postgres-port"
	PostgresUser           = "postgres-user"
	PostgresPassword       = "postgres-password"
	PostgresBatchWriteSize = "postgres-batch-write-size"

	RPCMaxRetry = "rpc-max-retry"

	MetricsCollectorURL    = "metrics-collector-url"
	MetricsCollectInterval = "metrics-collect-interval"

	StartCursor     = "start-cursor"
	EndCursor       = "end-cursor"
	TaskChannelSize = "task-channel-size"
)
