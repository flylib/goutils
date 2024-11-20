package nacos

type option struct {
	host           string
	port           int
	user, password string
	namespace      string
	scheme         string
}

type OptionFunc func(o *option)

func WithHost(host string, port int) OptionFunc {
	return func(o *option) {
		o.host = host
		o.port = port
	}
}

func WithAuth(username, password string) OptionFunc {
	return func(o *option) {
		o.user = username
		o.password = password
	}
}

func WithNamespace(namespace string) OptionFunc {
	return func(o *option) {
		o.namespace = namespace
	}
}

// Default:http
func WithGrpc(namespace string) OptionFunc {
	return func(o *option) {
		o.scheme = "grpc"
	}
}
