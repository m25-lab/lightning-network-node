package config

type Config struct {
	Database DatabaseConfig
	Node     NodeConfig
	Telegram Telegram
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Timeout  int
	Dbname   string
}

type NodeConfig struct {
	ChainId       string
	Endpoint      string
	CoinType      uint64
	PrefixAddress string
	TokenSymbol   string
}

type Telegram struct {
	BotId string
}
