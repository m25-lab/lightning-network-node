package config

type Config struct {
	Database DatabaseConfig
	Node     NodeConfig
	Telegram Telegram
	LNode    LNode
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

type LNode struct {
	External string
}
