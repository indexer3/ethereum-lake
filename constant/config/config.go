package config

const (
	RPCEndpoint = "rpc-endpoint"

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

	StartCursor     = "start-cursor"
	EndCursor       = "end-cursor"
	TaskChannelSize = "task-channel-size"
)
