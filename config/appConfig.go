package config

// AppConfiguration main config for the app
type AppConfiguration struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Github struct {
		ClientID     string `yaml:"clientId"`
		ClientSecret string `yaml:"clientSecret"`
	} `yaml:"github"`
	JWT struct {
		Secret     string `yaml:"secret"`
		CookieName string `yaml:"cookieName"`
		Domain     string `yaml:"domain"`
		Secure     bool   `yaml:"secure"`
	} `yaml:"jwt"`
	CORS struct {
		Domains []string `yaml:"domains"`
	} `yaml:"cors"`
}

// BaseObject represents base object for controllers and services
type BaseObject struct {
	Config *AppConfiguration
}
