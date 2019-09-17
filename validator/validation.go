package validator

import (
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/leebenson/conform"
	"github.com/thedevsaddam/govalidator"
)

type (
	// Rules interface
	Rules interface {
		Rules() map[string][]string
	}

	// Messages interface
	Messages interface {
		Messages() map[string][]string
	}

	// ValidatableRequest interface
	ValidatableRequest interface {
		Rules
		Messages
	}
)

// ValidateRequestData ...
func ValidateRequestData() RPCRequestValidationFunc {
	return func(r *rpc.RequestInfo, i interface{}) error {
		if err := conform.Strings(i); err != nil {
			return &json2.Error{
				Code:    json2.E_INTERNAL,
				Message: err.Error(),
			}
		}

		if vr, ok := i.(ValidatableRequest); ok {
			v := govalidator.New(govalidator.Options{
				Data:     i,
				Rules:    vr.Rules(),
				Messages: vr.Messages(),
			})
			if e := v.ValidateStruct(); len(e) > 0 {
				return &json2.Error{
					Code:    http.StatusUnprocessableEntity,
					Message: "Validation Error",
					Data:    e,
				}
			}
		}

		return nil
	}
}
