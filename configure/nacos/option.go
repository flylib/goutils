package nacos

type option struct {
	host             string
	port             int
	user, password   string
	namespace, group string
	scheme           string
	onChange         func(filed string, err error)
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

func WithNamespaceId(namespace string) OptionFunc {
	return func(o *option) {
		o.namespace = namespace
	}
}

// Default:http
func WithGrpc() OptionFunc {
	return func(o *option) {
		o.scheme = "grpc"
	}
}

// Default:DEFAULT_GROUP
func WithGroup(group string) OptionFunc {
	return func(o *option) {
		o.group = group
	}
}

func WithOnchange(f func(filed string, err error)) OptionFunc {
	return func(o *option) {
		o.onChange = f
	}
}
