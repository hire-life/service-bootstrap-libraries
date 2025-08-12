package response

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/zoondengine/hl/service-bootstrap-libraries/pkg/validation"
	"net/http"
)

var (
	ErrInvalidRequest = errors.New("invalid Request")
	ErrNotSupported   = errors.New("not supported")
	ErrOk             = errors.New("ok")
)

type HttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HttpValidationError struct {
	HttpResponse
	Errors []validation.Error
}

func Ok(message string) HttpResponse {
	return HttpResponse{
		Code:    http.StatusOK,
		Message: message,
	}
}

func BadRequest(message string) HttpResponse {
	return HttpResponse{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func InternalServerError(message string) HttpResponse {
	return HttpResponse{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func Unauthorized(message string) HttpResponse {
	return HttpResponse{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func MakeHttpValidationErrorResponse(e error) *HttpValidationError {
	var validationErrors validator.ValidationErrors
	if !errors.As(e, &validationErrors) {
		return nil
	}

	errs := make([]validation.Error, len(validationErrors))
	for i, err := range validationErrors {
		if err.Field() != "" {
			errs[i] = validation.Error{Field: err.Field(), Rule: err.Tag(), Message: err.Error()}
		}
	}

	return &HttpValidationError{
		HttpResponse: HttpResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
		},
		Errors: errs,
	}
}
