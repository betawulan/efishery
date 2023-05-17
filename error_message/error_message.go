package error_message

type Unauthorized struct {
	Message string `json:"message" example:"unauthorized"`
}

func (e Unauthorized) Error() string {
	return e.Message
}

type Duplicate struct {
	Message string `json:"message" example:"the phone already exist"`
}

func (e Duplicate) Error() string {
	return e.Message
}

type Failed struct {
	Message string `json:"message"`
}

func (e Failed) Error() string {
	return e.Message
}

type NotFound struct {
	Message string `json:"message" example:"phone or password incorrect"`
}

func (e NotFound) Error() string {
	return e.Message
}
