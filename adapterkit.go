package adapterkit

import "io"

type Context struct {
	Logger io.Writer // zap
}

type Action interface {
	// split in steps? SetInput(), etc?
	Run(ctx Context, input interface{}) (output interface{}, err error)
}
