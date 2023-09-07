package config

const (
	ConfigPathKey = "config-path"

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

	StartCursor = "start-cursor"
	EndCursor   = "end-cursor"
	// CatchUpSleepInterval is the interval in seconds to sleep when catching up
	CatchUpSleepInterval = "catch-up-sleep-interval"

	ChainbaseSQLAPIKey = "chainbase-sql-api-key"

	EthereumRPCUrl = "ethereum-rpc-url"
	OptimismRPCUrl = "optimism-rpc-url"
	ArbitrumRPCUrl = "arbitrum-rpc-url"
	BSCRPCUrl      = "bsc-rpc-url"
	PolygonRPCUrl  = "polygon-rpc-url"
	BaseRPCUrl     = "base-rpc-url"
	GoerliRPCUrl   = "goerli-rpc-url"
)
