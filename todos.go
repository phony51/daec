package daec

import (
	"strconv"
	"unicode"
)

func mustParseNumber(input string, pDst *string, pPosAfter *int) {

	strnum := make(string, 0)
	posAfter :=
	for i, char := range input {
		if unicode.IsDigit(char) {
			strnum = append(strnum, char)
		}
	}
	*pPosAfter += len(strnum)
	*pDst, _ = strconv.Atoi(string(strnum))
}
func InfixToPostfix(infix string) string {
	var (
		st      StringStack
		postfix string
	)
	strchar := string(char)
	for i, char := range infix {

		if char >= '0' && char <= '9' {
			postfix += mustParseNumber(infix[:i])
		} else if char == '(' {
			st.PushItoP(strchar)
		} else if char == ')' {
			for st.StringTop() != "(" {
				postfix += st.StringTop()
				st.Pop()
			}
			st.StringTop()
		} else {
			for !st.IsEmpty() && precedence(strchar) <= precedence(st.StringTop()) {
				postfix += st.StringTop()
				st.Pop()
			}
			st.PushItoP(strchar)
		}
	}
	for !st.IsEmpty() {
		postfix += st.StringTop()
		st.Pop()
	}
	return postfix
}
