package config

type Database struct {
	Name     string
	userColl string
}

type appConfig struct {
	TokenLifetime int
	DB            Database
}

func NewAppConfig() *appConfig {
	return &appConfig{}
}
