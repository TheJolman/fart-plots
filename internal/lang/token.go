package lang

type Token int

const (
	Uncategorized Token = iota
	// single char tokens
	LeftBrace
	RightBrace
	Comma

	// literals
	Identifier
	String

	// keywords
	Plot
	Bar
	Scatter
	Box
	Histogram
)
