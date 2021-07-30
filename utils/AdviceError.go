package utils

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func ErroAPIHandle(err error) Error {
	erro := Error{}
	erro.Message = err.Error()
	return erro
}
