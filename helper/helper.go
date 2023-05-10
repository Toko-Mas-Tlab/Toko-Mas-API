package helper

type Respon struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponseJson(message string, code int, status string, data interface{}) Respon {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonrespon := Respon{
		Meta: meta,
		Data: data,
	}
	return jsonrespon
}
