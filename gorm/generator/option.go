package generator

import "gorm.io/gen"

type (
	Option           func(config *option)
	FileNameHandler  func(tableName string) (fileName string)
	ModelNameHandler func(tableName string) (fileName string)
)

type option struct {
	gen.Config
	FileNameHandler
	ModelNameHandler
	genTables []string
}

// default: all
func GenTables(table ...string) Option {
	return func(c *option) {
		c.genTables = table
	}
}

// default: ./model
func OutPath(path string) Option {
	return func(c *option) {
		c.OutPath = path
	}
}

// default: xxx.gen.go
func OutFileSuffixName(path string) Option {
	return func(c *option) {
		c.OutPath = path
	}
}

func OutFileNameHandler(handler FileNameHandler) Option {
	return func(c *option) {
		c.FileNameHandler = handler
	}
}
func OutModelNameHandler(handler ModelNameHandler) Option {
	return func(c *option) {
		c.ModelNameHandler = handler
	}
}

func FieldNullable(yes bool) Option {
	return func(c *option) {
		c.FieldNullable = yes
	}
}

// generate pointer when field has default value
func FieldCoverable(yes bool) Option {
	return func(c *option) {
		c.FieldCoverable = yes
	}
}

func GenerateMode(mode gen.GenerateMode) Option {
	return func(c *option) {
		c.Mode = mode
	}
}
