package signer

type Config struct {
	BindAddress string
	Contracts   ContractsConfig
	Ethereum    EthereumConfig
	Generator   string
}

type ContractsConfig struct {
	ContractAddress        string
	DistKeyContractAddress string
}

type EthereumConfig struct {
	Address    string
	PrivateKey string
	ChainID    int64
}
