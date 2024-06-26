package generator

import "gorm.io/gen"

type (
	Option           func(config *option)
	FileNameHandler  func(tableName string) (fileName string)
	ModelNameHandler func(tableName string) (modelName string)
)

type option struct {
	gen.Config
	FileNameHandler
	ModelNameHandler
	genTables []string
	mapTypes  map[string]string
}

// default: all
func GenTables(table ...string) Option {
	return func(c *option) {
		c.genTables = table
	}
}

// default: ./model
func WithOutPath(path string) Option {
	return func(c *option) {
		c.OutPath = path
	}
}

// default: xxx.gen.go
func WithOutFileSuffixName(path string) Option {
	return func(c *option) {
		c.OutPath = path
	}
}

func WithOutFileNameHandler(handler FileNameHandler) Option {
	return func(c *option) {
		c.FileNameHandler = handler
	}
}
func WithOutModelNameHandler(handler ModelNameHandler) Option {
	return func(c *option) {
		c.ModelNameHandler = handler
	}
}

func WithFieldNullable(yes bool) Option {
	return func(c *option) {
		c.FieldNullable = yes
	}
}

// generate pointer when field has default value
func WithFieldCoverable(yes bool) Option {
	return func(c *option) {
		c.FieldCoverable = yes
	}
}

/*
	【default】

"bool": {  "bool" },
"tinyint":  {  "int8" },
"smallint": {  "int16" },
"mediumint": {  "int32" },
"bigint":   {  "int64" },
"int":      {  "int" },
*/
func WithMapGoTypes(mapTypes map[string]string) Option {
	return func(c *option) {
		c.mapTypes = mapTypes
	}
}
