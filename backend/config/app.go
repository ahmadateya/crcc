package config

type AppConfiguration struct {
	Docker Docker
}
type Docker struct {
	Host string
	Port string
}
