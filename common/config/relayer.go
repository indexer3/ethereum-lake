package config

type RelayerConfig struct {
	Port             string           `mapstructure:"port" json:"port" yaml:"port"`
	ChainbaseAPIKey  string           `mapstructure:"chainbase_api_key" json:"chainbase_api_key" yaml:"chainbase_api_key"`
	NetworkRPCConfig NetworkRPCConfig `mapstructure:"network_rpc_config" json:"network_rpc_config" yaml:"network_rpc_config"`
}

type NetworkRPCConfig struct {
	EthereumRPCUrl string `mapstructure:"ethereum_rpc_url" json:"ethereum_rpc_url" yaml:"ethereum_rpc_url"`
	PolygonRPCUrl  string `mapstructure:"polygon_rpc_url" json:"polygon_rpc_url" yaml:"polygon_rpc_url"`
	ArbitrumRPCUrl string `mapstructure:"arbitrum_rpc_url" json:"arbitrum_rpc_url" yaml:"arbitrum_rpc_url"`
	OptimismRPCUrl string `mapstructure:"optimism_rpc_url" json:"optimism_rpc_url" yaml:"optimism_rpc_url"`
	BSCRPCUrl      string `mapstructure:"bsc_rpc_url" json:"bsc_rpc_url" yaml:"bsc_rpc_url"`
	BaseRPCUrl     string `mapstructure:"base_rpc_url" json:"base_rpc_url" yaml:"base_rpc_url"`
}

var RelayerConf RelayerConfig
