package config

type Config struct {
	Templates []Template `yaml:"templates"`
}

type Template struct {
	Source      string `yaml:"src"`
	Destination string `yaml:"dest"`
}
