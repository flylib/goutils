package configger

type Option func(c *Config)

// if (isReadFromEnv==true) { config.{env}.yaml } else { config.env.yaml}
func Environment(isReadFromEnv bool, env string) Option {
	return func(c *Config) {
		c.isReadFromEnvironment = isReadFromEnv
		c.env = env
	}
}

// Default current program directory
func SearchPath(path string) Option {
	return func(c *Config) {
		c.path = path
	}
}

func WatchFileChange(f func(config any, err error)) Option {
	return func(c *Config) {
		c.configChangeCallBack = f
	}
}
