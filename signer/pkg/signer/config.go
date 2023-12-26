package signer

type Config struct {
	BindAddress string
	PrivateKey  string
	Contracts   ContractsConfig
	Ethereum    EthereumConfig
}

type ContractsConfig struct {
	OracleContractAddress  string
	DistKeyContractAddress string
}

type EthereumConfig struct {
	Address    string
	PrivateKey string
	ChainID    int64
}
