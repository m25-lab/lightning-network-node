package crypto

type IEngine interface {
	Hash(secret interface{}) (string, error)
	Compare(secret interface{}, witness string) (bool, error)
}

type Crypto struct {
	SuperKey string
}

var CryptoEngine Crypto

func (cryptoEngine *Crypto) Hash(secret interface{}) (string, error) {
	return "SecretHashNotImplement", nil
}

func (cryptoEngine *Crypto) Compare(secret interface{}, witness string) (bool, error) {
	return true, nil
}
