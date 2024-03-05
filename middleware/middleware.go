// middleware.go
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	BasicAuth struct {
		Users []User `yaml:"users"`
	} `yaml:"basic_auth"`
}

func LoadConfig() (Config, error) {
	var config Config
	file, err := os.Open("config.yaml")
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config, err := LoadConfig()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		user, pass, ok := c.Request.BasicAuth()
		if !ok {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if the provided credentials match any user in the config
		validUser := false
		for _, u := range config.BasicAuth.Users {
			if user == u.Username && pass == u.Password {
				validUser = true
				break
			}
		}

		if !validUser {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		fmt.Println("Authentication successful - ", "User:", user, "Pass:", pass)

		// Call the next handler
		c.Next()
	}
}
