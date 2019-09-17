package validator

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/dmitrymomot/jsonrpc2/request"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

// RBAC validation func
func RBAC(model, policy string) RPCRequestValidationFunc {
	enforcer := casbin.NewEnforcer(model, policy)
	return RBACCustomEnforcer(enforcer)
}

// RBACCustomEnforcer validation func
func RBACCustomEnforcer(e *casbin.Enforcer) RPCRequestValidationFunc {
	return func(r *rpc.RequestInfo, i interface{}) error {
		if ar, ok := i.(request.Auth); ok && ar != nil {
			if !e.Enforce(ar.GetUserRole(), r.Method) {
				return &json2.Error{
					Code:    http.StatusForbidden,
					Message: http.StatusText(http.StatusForbidden),
				}
			}
		}
		return nil
	}
}
