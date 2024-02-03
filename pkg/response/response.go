package response

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

type ResponseError struct {
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"error,omitempty"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}
