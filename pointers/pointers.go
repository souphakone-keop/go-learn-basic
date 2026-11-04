package pointers

import "fmt"

// Config represents the application configuration
type Config struct {
	LogLevel string
	Port     int
}

func UpdateChange(c *Config) {
	c.LogLevel = "end"
	c.Port = 3000
}

// UpdateConfig modifies the provided configuration
func Pointers() {
	configs := Config{
		LogLevel: "start",
		Port:     8080,
	}
	fmt.Println(configs)
	UpdateChange(&configs)
	fmt.Println(configs)
}
