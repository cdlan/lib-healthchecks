package healthchecks

import "os"

// Config is the struct used to setup the data
type Config struct {

	// Url is the url that will be used as a base to notify state
	Url string `mapstructure:"url"`
}

// LoadVarsFromEnv checks if each var exists inside os environment variables.
// If it does exist, overrides the existing value
func (C *Config) LoadVarsFromEnv() {

	BaseUrl, ok := os.LookupEnv("HC_URL")
	if ok {
		C.Url = BaseUrl
	}
}
