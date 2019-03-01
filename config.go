// Package config holds interface definition for configuration
package config

import "time"

// Config is an interface for configuration implementations
type Config interface {
	Section(key string) (Config, error)
	SectionAsJSON(key string) (string, error)
	UnmarshalSection(key string, dest interface{}) error
	GetString(key string, defaultVal string) string
	MustGetString(key string) string
	GetInt(key string, defaultVal int) int
	MustGetInt(key string) int
	GetUint64(key string, defaultVal uint64) uint64
	MustGetUint64(key string) uint64
	GetTime(key string, defaultVal time.Time) time.Time
	MustGetTime(key string) time.Time
	GetDuration(key string, defaultVal time.Duration) time.Duration
	MustGetDuration(key string) time.Duration
	GetStringSlice(key string, defaultVal []string) []string
	MustGetStringSlice(key string) []string
	GetBool(key string, defaultVal bool) bool
	MustGetBool(key string) bool
}
