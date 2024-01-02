package signer

type Config struct {
	BindAddress string
	Contracts   ContractsConfig
	Ethereum    EthereumConfig
	Generator   string
}

type ContractsConfig struct {
	RegistryContractAddress string
	SakaiContractAddress    string
	IBSASContractAddress  string
}

type EthereumConfig struct {
	Address    string
	PrivateKey string
	ChainID    int64
}
