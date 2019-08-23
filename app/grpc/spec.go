package grpc

import "github.com/pkg/errors"

const (
	prefix       = "app.grpc"
	unmarshalKey = "grpc"
)

var errCfgInvalid = errors.New("cfg is not present or invalid")
