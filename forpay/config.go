package forpay

import (
	"reflect"
	"strconv"
	"time"
)

// Config provides configurations to change api client's behaviuor.
type Config struct {
	AutoRetry    bool   `default:"true"`
	MaxRetryTime int    `default:"3"`
	UserAgent    string `default:""`
	Debug        bool   `default:"false"`
	Timeout      time.Duration
}

// DefaultConfig returns config with default configurations.
func DefaultConfig() *Config {
	config := &Config{}
	initConfigWithDefaultTag(config)
	return config
}

// EnableAutoRetry enables auto retry feature.
func (c *Config) EnableAutoRetry() *Config {
	panic("not implemented yet")

	// c.AutoRetry = true
	// return c
}

// DisableAutoRetry disables auto retry feature.
func (c *Config) DisableAutoRetry() *Config {
	c.AutoRetry = false
	return c
}

// WithMaxRetryTime sets the maximum retry time of an api call.
func (c *Config) WithMaxRetryTime(maxRetryTime int) *Config {
	c.MaxRetryTime = maxRetryTime
	return c
}

// WithTimeout sets request timeout for client.
func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}

func initConfigWithDefaultTag(src interface{}) {
	configType := reflect.TypeOf(src)

	for i := 0; i < configType.Elem().NumField(); i++ {
		field := configType.Elem().Field(i)
		defaultValue := field.Tag.Get("default")
		if defaultValue == "" {
			continue
		}

		setter := reflect.ValueOf(src).Elem().Field(i)

		switch field.Type.String() {
		case "int":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue)
		case "time.Duration":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue)
		case "string":
			setter.SetString(defaultValue)
		case "bool":
			boolValue, _ := strconv.ParseBool(defaultValue)
			setter.SetBool(boolValue)
		}
	}
}
