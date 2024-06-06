package pkg

import "os"

// test TODO

func GetEnvWithDefault(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
