package configure

type Option func(c *Config)

// if ($env) { config.{env}(environment).yaml } else { config.env.yaml}
func WithEnvironment(env string) Option {
	return func(c *Config) {
		c.env = env
	}
}

// Default current program directory
func WithSearchPath(path string) Option {
	return func(c *Config) {
		c.path = path
	}
}

func WithFileChangeCallBack(f func(config any, err error)) Option {
	return func(c *Config) {
		c.onChangeCallBack = f
	}
}
