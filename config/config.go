package config

// YamlConfig represents the YAML configuration read from a yaml file
type YamlConfig struct {
	Server                   Server      `yaml:"server"`
	DataSource               DataSource  `yaml:"datasource"`
	Swagger                  Swagger     `yaml:"swagger"`
}

type Swagger struct {
	Host     string `yaml:"host"`
	Basepath string `yaml:"basepath"`
}
type Server struct {
	Address string `yaml:"address"`
	Port    int64  `yaml:"port"`
}

type DataSource struct {
	Host         string `yaml:"host"`
	Port         int64  `yaml:"port"`
	DatabaseName string `yaml:"dbname"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	AuthSource   string `yaml:"authsource"`
}

func GetServiceConfig(profile string) *YamlConfig {
	return &YamlConfig{
		Server:     Server{
			Address: "",
			Port:    8088,
		},
		DataSource: DataSource{
			Host:         "localhost",
			Port:         5542,
			DatabaseName: "wednesday_db",
			Username:     "wednesday",
			Password:     "wednesday",
			AuthSource:   "",
		},
		Swagger:    Swagger{
			Host:     "localhost:8088",
			Basepath: "/",
		},
	}
}
