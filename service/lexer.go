package service

type Lexer interface {
	NextChar(input string)
	Peek(input string) string
	SkipWhitespace(string)
}

type Lex struct {
	CurrChar string
	CurrPos  int
	Input    string
}

func NewLex(in string) *Lex {
	lex := Lex{
		Input: in,
	}
	return &lex
}

var _ Lexer = &Lex{}

func (lex *Lex) NextChar(input string) {
	/*
		Process the next character
	*/
	lex.CurrPos++
	if lex.CurrPos >= len(input) {
		lex.CurrChar = `\0` // EOF
	} else {
		lex.CurrChar = string(input[lex.CurrPos])
	}

}
func (lex *Lex) Peek(input string) string {
	/*
		Return the lookahead character
	*/
	if lex.CurrPos+1 >= len(input) {
		return string(`\0`)
	}

	temp := input[lex.CurrPos+1]
	return string(temp)
}

func (lex *Lex) SkipWhitespace(input string) {

}
