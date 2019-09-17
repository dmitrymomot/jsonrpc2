package validator

import (
	"github.com/dmitrymomot/go-jwt"
	"github.com/dmitrymomot/jsonrpc2/request"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

// JWT validation func
func JWT(ji jwt.Interactor) RPCRequestValidationFunc {
	return func(r *rpc.RequestInfo, i interface{}) error {
		if ar, ok := i.(request.Auth); ok && ar != nil {
			if ar.GetAccessToken() == "" {
				return &json2.Error{
					Code:    json2.E_INVALID_REQ,
					Message: "missed access token",
				}
			}

			claims := new(jwt.DefaultClaims)
			if err := ji.Parse(ar.GetAccessToken(), claims); err != nil {
				return &json2.Error{
					Code:    json2.E_INVALID_REQ,
					Message: err.Error(),
				}
			}

			ar.SetClaims(claims)
			ar.SetUserID(claims.UserID)
			ar.SetUserRole(claims.Role)
			ar.SetAppID(claims.ApplicationID)
		}
		return nil
	}
}
