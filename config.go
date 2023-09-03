package gotransaction

type Config struct {
	Rpc        string
	PrivateKey string
}

type config struct {
	rpc        string
	privateKey string
}

func (c *Config) config() (conf *config) {
	conf = new(config)
	conf.rpc = c.Rpc
	conf.privateKey = c.PrivateKey
	return
}
