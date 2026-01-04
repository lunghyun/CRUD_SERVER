package cerrors

import (
	"fmt"
	"net/http"
)

// 에러 관련 타입 및 함수들을 정의

const (
	BadRequestBind = http.StatusBadRequest
	NotFoundUser   = http.StatusNotFound

	InternalServerError = http.StatusInternalServerError
)

var errMessage = map[int64]string{
	NotFoundUser: "user not found",
}

func Errorf(code int64, args ...interface{}) error {
	message, ok := errMessage[code]
	if ok {
		return fmt.Errorf("%s : %v", message, args)
	}
	return fmt.Errorf("not found err code")
}
