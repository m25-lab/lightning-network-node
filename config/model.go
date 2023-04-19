package config

type Config struct {
	Database     Database
	Node         Node
	Telegram     Telegram
	LNode        LNode
	Kafka        Kafka
	Corechain    CoreChain
	CryptoEngine CryptoEngine
}

type Database struct {
	Host     string
	User     string
	Password string
	Timeout  int
	Dbname   string
}

type Node struct {
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

type Kafka struct {
	Brokers []string
}

type CoreChain struct {
	Endpoint string
}

type CryptoEngine struct {
	SuperKey string
}
