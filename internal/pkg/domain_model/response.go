package domain_model

// Response response model
type Response struct {
	Result interface{} `json:"result"`
}

// BuildResponse response builder
func BuildResponse(result interface{}) Response {
	return Response{
		Result: result,
	}
}
