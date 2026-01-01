package types

type APIResponse struct {
	Result      int    `json:"result"`
	Description string `json:"description"`
}

func NewAPIResponse(description string, result int) *APIResponse {
	return &APIResponse{
		Result:      result,
		Description: description,
	}
}
