package error_message

type Unauthorized struct {
	Message string
}

func (e Unauthorized) Error() string {
	return e.Message
}

type Duplicate struct {
	Message string
}

func (e Duplicate) Error() string {
	return e.Message
}

type Failed struct {
	Message string
}

func (e Failed) Error() string {
	return e.Message
}

type NotFound struct {
	Message string
}

func (e NotFound) Error() string {
	return e.Message
}
