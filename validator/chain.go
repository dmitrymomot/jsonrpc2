package validator

import "github.com/gorilla/rpc/v2"

// RPCRequestValidationFunc ...
type RPCRequestValidationFunc func(r *rpc.RequestInfo, i interface{}) error

// WrapFunc is a wrapper to run chain of rpc request validation functions
func WrapFunc(fn ...RPCRequestValidationFunc) RPCRequestValidationFunc {
	return func(r *rpc.RequestInfo, i interface{}) error {
		for _, f := range fn {
			if err := f(r, i); err != nil {
				return err
			}
		}
		return nil
	}
}
