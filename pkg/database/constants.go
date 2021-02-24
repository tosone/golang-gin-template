package database

const pageSize = 100

// Options ..
type Options struct {
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
	Category string `form:"category"`
}
