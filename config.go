// Package config holds interface definition for configuration
package config

import "time"

// Config is an interface for configuration implementations
type Config interface {
	Section(key string) (Config, error)
	GetString(key string, defaultVal string) string
	MustGetString(key string) string
	GetInt(key string, defaultVal int) int
	MustGetInt(key string) int
	GetTime(key string, defaultVal time.Time) time.Time
	MustGetTime(key string) time.Time
}
