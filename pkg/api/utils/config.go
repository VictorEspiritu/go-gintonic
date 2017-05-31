package utils

type ServiceConfig struct {
	Server Server `json:"server"`
}

type Server struct {
	Port string `json:"port"`
	Host string `json: "host"`
}
