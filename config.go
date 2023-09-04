package gotransaction

type Config struct {
	EndPoint   string
	PrivateKey string
}

type config struct {
	endPoint   string
	privateKey string
}

func (c *Config) config() (conf *config) {
	conf = new(config)
	conf.endPoint = c.EndPoint
	conf.privateKey = c.PrivateKey
	return
}
