package mysql

import "regexp"

type Option func(*option)

type option struct {
	maxOpenConns int
	maxIdleConns int
}

func WithMaxOpenConns(num int) Option {
	return func(o *option) {
		o.maxOpenConns = num
	}
}

// The default max idle connections is currently 2.
func WithMaxIdleConns(num int) Option {
	return func(o *option) {
		o.maxIdleConns = num
	}
}

func WithRawParamRegexp(exp string) Option {
	return func(o *option) {
		var err error
		rawRegexp, err = regexp.Compile(exp)
		if err != nil {
			panic(err)
		}
	}
}

func WithProcedureParamRegexp(exp string) Option {
	return func(o *option) {
		var err error
		pcdRegexp, err = regexp.Compile(exp)
		if err != nil {
			panic(err)
		}
	}
}
