package error_message

type Unauthorized struct {
	Message string `json:"message"`
}

func (e Unauthorized) Error() string {
	return e.Message
}

type Duplicate struct {
	Message string `json:"message"`
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
	Message string `json:"message"`
}

func (e NotFound) Error() string {
	return e.Message
}
