// Package json encapsulates structure and methods for
// parsing and getting values from json configuration files
package json

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/Darkren/weatherservice/config"
)

// Config represents data type for configuration
// parsed from JSON
type Config struct {
	c map[string]*json.RawMessage
}

// New parses JSON string and gets config structure
func New(jsonStr string) (config.Config, error) {
	return new([]byte(jsonStr))
}

// Load reads file from filePath, parses JSON and
// gets config structure
func Load(filePath string) (config.Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return new(data)
}

// Section returns config section by key. Used for nested objects
// within configuration
func (c Config) Section(key string) (config.Config, error) {
	section := Config{}

	if err := json.Unmarshal(*(c.c[key]), &(section.c)); err != nil {
		return nil, err
	}

	return section, nil
}

// GetString tries to get string value by key from configuration.
// Returns acquired value or the specified default value
func (c Config) GetString(key string, defaultVal string) string {
	value, err := c.getString(key)
	if err != nil {
		return defaultVal
	}

	return value
}

// MustGetString tries to get string value by key from configuration.
// Returns acquired value or panics in case of any error
func (c Config) MustGetString(key string) string {
	value, err := c.getString(key)
	if err != nil {
		panic(err)
	}

	return value
}

// GetInt tries to get int value by key from configuration.
// Returns acquired value or the specified default value
func (c Config) GetInt(key string, defaultVal int) int {
	value, err := c.getInt(key)
	if err != nil {
		return defaultVal
	}

	return value
}

// MustGetInt tries to get int value by key from configuration.
// Returns acquired value or panics in case of any error
func (c Config) MustGetInt(key string) int {
	value, err := c.getInt(key)
	if err != nil {
		panic(err)
	}

	return value
}

// GetTime tries to get time.Time value by key from configuration.
// Returns acquired value or the specified default value
func (c Config) GetTime(key string, defaultVal time.Time) time.Time {
	value, err := c.getTime(key)
	if err != nil {
		return defaultVal
	}

	return value
}

// MustGetTime tries to get time.Time value by key from configuration.
// Returns acquired value or panics in case of any error
func (c Config) MustGetTime(key string) time.Time {
	value, err := c.getTime(key)
	if err != nil {
		panic(err)
	}

	return value
}

func new(jsonData []byte) (config.Config, error) {
	config := Config{}

	if err := json.Unmarshal(jsonData, &(config.c)); err != nil {
		return nil, err
	}

	return config, nil
}

func (c Config) getString(key string) (string, error) {
	var value string

	if err := json.Unmarshal(*(c.c[key]), &value); err != nil {
		return "", err
	}

	return value, nil
}

func (c Config) getInt(key string) (int, error) {
	var value int

	if err := json.Unmarshal(*(c.c[key]), &value); err != nil {
		return 0, err
	}

	return value, nil
}

func (c Config) getTime(key string) (time.Time, error) {
	valueStr, err := c.getString(key)
	if err != nil {
		return time.Now(), err
	}

	value, err := time.Parse("2.1.2006", valueStr)
	if err != nil {
		return time.Now(), err
	}

	return value, nil
}
