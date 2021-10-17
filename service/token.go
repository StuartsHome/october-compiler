package service

type tokener interface {
	GetToken() error
}

type token struct {
}

var _ tokener = &token{}

// var _ token = tokener{}

func NewToken() *token {
	tk := token{}
	return &tk
}

func (tk *token) GetToken() error {
	return nil
}
