package configure

type IConfigure interface {
	Env() string
	Scan(v any) error
}
