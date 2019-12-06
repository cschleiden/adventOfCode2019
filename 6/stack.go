package main

// https://stackoverflow.com/a/28542256/561159
type e struct {
	object string
	path   int
}

type stack []e

func (s stack) Push(v e) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, e) {
	l := len(s)
	return s[:l-1], s[l-1]
}
