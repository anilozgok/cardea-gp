package config

func Default() *Config {
	return &Config{
		Secrets: new(Secrets),
	}
}

type Secrets struct {
	CardeaDBCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"cardeaDBCredentials"`
}

type Config struct {
	Secrets  *Secrets
	CardeaDB struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
	} `json:"cardeaDB"`
}
