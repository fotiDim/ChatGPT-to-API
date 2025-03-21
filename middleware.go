package main

import (
	"bufio"
	"os"
	"strings"

	gin "github.com/gin-gonic/gin"
)

var ADMIN_PASSWORD string
var API_KEYS map[string]bool

func init() {
	ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")
	if ADMIN_PASSWORD == "" {
		ADMIN_PASSWORD = "TotallySecurePassword"
	}

	API_KEYS = make(map[string]bool)

	// Read API keys from environment variable
	if envAPIKeys := os.Getenv("API_KEYS"); envAPIKeys != "" {
		for _, key := range strings.Split(envAPIKeys, ",") {
			if key != "" {
				API_KEYS["Bearer "+strings.TrimSpace(key)] = true
			}
		}
	}

	// Read API keys from file (existing functionality)
	if _, err := os.Stat("api_keys.txt"); err == nil {
		file, _ := os.Open("api_keys.txt")
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			key := scanner.Text()
			if key != "" {
				API_KEYS["Bearer "+key] = true
			}
		}
	}
}

func adminCheck(c *gin.Context) {
	password := c.Request.Header.Get("Authorization")
	if password != ADMIN_PASSWORD {
		c.String(401, "Unauthorized")
		c.Abort()
		return
	}
	c.Next()
}

func cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Next()
}

func Authorization(c *gin.Context) {
	if len(API_KEYS) != 0 && !API_KEYS[c.Request.Header.Get("Authorization")] {
		if c.Request.Header.Get("Authorization") == "" {
			c.JSON(401, gin.H{"error": "No API key provided. Please provide an API key as part of the Authorization header."})
		} else if strings.HasPrefix(c.Request.Header.Get("Authorization"), "Bearer sk-") {
			c.JSON(401, gin.H{"error": "You tried to use the official API key which is not supported."})
		} else if strings.HasPrefix(c.Request.Header.Get("Authorization"), "Bearer eyJhbGciOiJSUzI1NiI") {
			return
		} else {
			c.JSON(401, gin.H{"error": "Invalid API key."})
		}
		c.Abort()
		return
	}
	c.Next()
}
