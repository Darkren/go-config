// Package json encapsulates structure and methods for
// parsing and getting values from json configuration files
package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	config "github.com/Darkren/go-config"
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

	if _, ok := c.c[key]; !ok {
		return nil, fmt.Errorf("section %s not present in config", key)
	}

	if err := json.Unmarshal(*(c.c[key]), &(section.c)); err != nil {
		return nil, err
	}

	return section, nil
}

// SectionAsJSON returns config section as JSON string. Used for nested objects
// within configuration
func (c Config) SectionAsJSON(key string) (string, error) {
	sectionBytes, ok := c.c[key]
	if !ok {
		return "", fmt.Errorf("section %s not present in config", key)
	}

	return string(*sectionBytes), nil
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

// GetDuration tries to get time.Duration value by key from configuration.
// The value must be a valid string to be parsed by standard methods. Returns
// acquired value or the specified default value
func (c Config) GetDuration(key string, defaultVal time.Duration) time.Duration {
	value, err := c.getDuration(key)
	if err != nil {
		return defaultVal
	}

	return value
}

// MustGetDuration tries to get time.Duration value by key from configuration.
// The value must be a valid string to be parsed by standard methods. Returns
// acquired value or panics in case of any error
func (c Config) MustGetDuration(key string) time.Duration {
	value, err := c.getDuration(key)
	if err != nil {
		panic(err)
	}

	return value
}

// GetStringSlice tries to get the string slice value by key from configuration.
// Returns acquired value or the specified default value
func (c Config) GetStringSlice(key string, defaultVal []string) []string {
	value, err := c.getStringSlice(key)
	if err != nil {
		return defaultVal
	}

	return value
}

// GetStringSlice tries to get the string slice value by key from configuration.
// Returns acquired value or panics in case of any error
func (c Config) MustGetStringSlice(key string) []string {
	value, err := c.getStringSlice(key)
	if err != nil {
		panic(err)
	}

	return value
}

// GetBool tries to get bool value by key from configuration.
// Returns acquired value or the specified default value
func (c Config) GetBool(key string, defaultVal bool) bool {
	value, err := c.getBool(key)
	if err != nil {
		return defaultVal
	}

	return value
}

// MustGetBool tries to get bool value by key from configuration.
// Returns acquired value or panics in case of any error
func (c Config) MustGetBool(key string) bool {
	value, err := c.getBool(key)
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

func (c Config) getBool(key string) (bool, error) {
	var value bool

	if err := json.Unmarshal(*(c.c[key]), &value); err != nil {
		return false, err
	}

	return value, nil
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

func (c Config) getDuration(key string) (time.Duration, error) {
	valueStr, err := c.getString(key)
	if err != nil {
		return time.Nanosecond, err
	}

	value, err := time.ParseDuration(valueStr)
	if err != nil {
		return time.Nanosecond, err
	}

	return value, nil
}

func (c Config) getStringSlice(key string) ([]string, error) {
	var value []string

	if err := json.Unmarshal(*(c.c[key]), &value); err != nil {
		return nil, err
	}

	return value, nil
}