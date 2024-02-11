package api

import (
	"context"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func initValidator() {
	validate = validator.New()
}

func ReadRequest(ctxBind func() error, ctxReq func() context.Context, req interface{}) error {
	initValidator()

	if err := ctxBind(); err != nil {
		return err
	}

	ctx := ctxReq()

	if err := validate.StructCtx(ctx, req); err != nil {
		return err
	}
	return nil
}
