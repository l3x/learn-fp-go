package server

import (
	. "utils"
	"errors"
)

type ServerOption func(*options) error

func MaxNumber(n int) ServerOption {
	return func(o *options) error {
		o.maxNumber = n
		return nil
	}
}

func MaxConcurrentConnections(n int) ServerOption {
	return func(o *options) error {
		if n > Config.MaxConcurrentConnections {
			return errors.New("error setting MaxConcurrentConnections")
		}
		o.maxConcurrentConnections = n
		return nil
	}
}

type convert func(int) (string, error)

func FormatNumber(fn convert) ServerOption {
	return func(o *options) (err error) {
		o.convertFn = fn
		return
	}
}

func UseNumberHandler(b bool) ServerOption {
	return func(o *options) error  {
		o.useNumberHandler = b
		return nil
	}
}

