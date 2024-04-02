package httpModul

import (
	"encoding/json"
	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/api/apiv1"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, message string, errorCode apiv1.ErrorCode, statusCode int) {
	errMsg := message
	apiError := apiv1.Error{
		Code:    errorCode,
		Message: &errMsg,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(apiError)
}
