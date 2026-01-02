package types

type APIResponse struct {
	Result      int         `json:"result_code"`
	Description string      `json:"description"`
	ErrorCode   interface{} `json:"error_code"`
}

func NewAPIResponse(description string, result int, errCode interface{}) *APIResponse {
	return &APIResponse{
		Result:      result,
		Description: description,
		ErrorCode:   errCode,
	}
}
