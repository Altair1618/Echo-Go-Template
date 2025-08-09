package response

import (
	"errors"
	"net/http"

	appErrors "github.com/Altair1618/Echo-Go-Template/internal/shared/errors"
)

func MapAppError(err error, isDev bool) (int, Response[any]) {
	var appErr *appErrors.AppError
	if !errors.As(err, &appErr) {
		resp := Response[any]{
			Status:  StatusError,
			Message: "An unexpected error occurred. Please try again later.",
		}

		if isDev {
			resp.Error = err.Error()
		}

		return http.StatusInternalServerError, resp
	}

	if appErr.Code < 500 {
		return appErr.Code, Response[any]{
			Status:  StatusFailed,
			Message: appErr.Message,
			Error:   appErr.Error(),
		}
	}

	if isDev {
		return appErr.Code, Response[any]{
			Status:  StatusError,
			Message: appErr.Message,
			Error:   appErr.Error(),
		}
	} else {
		return http.StatusInternalServerError, Response[any]{
			Status:  StatusError,
			Message: "An unexpected error occurred. Please try again later.",
		}
	}
}
