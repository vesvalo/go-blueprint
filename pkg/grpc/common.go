package grpc

import "github.com/pkg/errors"

const (
	Prefix       = "app.grpc"
	UnmarshalKey = "grpc"
)

var errCfgInvalid = errors.New("cfg is not present or invalid")
