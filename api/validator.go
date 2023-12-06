package api

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ReadRequest(ctxBind func() error, ctxReq func(), req interface{}) error {
	if err := ctxBind(); err != nil {
		return err
	}

	if err := validate.StructCtx(ctxReq(), req); err != nil {
		return err
	}
	return validate.StructCtx(ctxReq(), req)
}
