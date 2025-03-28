package lang

type Lexer struct {
	tokens []Token
	start uint
	current uint
	line uint

}

func (l *Lexer) ScanTokens() []Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}

	l.tokens = append(l.tokens, Token{})
	return l.tokens
}

func (l *Lexer) scanToken() {
	panic("not implemneted")

}

func (l Lexer) isAtEnd() bool {
	return true
}
