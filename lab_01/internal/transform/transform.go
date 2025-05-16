package transform

import (
	"strings"
	"unicode"
)

const (
	maxPriority = 4
)

// insertConcatOperators восстанавливает
func insertConcatOperators(infix string) string {
	var result strings.Builder
	n := len(infix)

	for i := 0; i < n; i++ {
		result.WriteByte(infix[i])

		if i+1 < n && isImplicitConcat(infix[i], infix[i+1]) {
			result.WriteByte('.')
		}
	}

	return result.String()
}

// isImplicitConcat нужна ли неявная конкатенация
func isImplicitConcat(a, b byte) bool {
	return (isAlphanumeric(a) && isAlphanumeric(b)) ||
		(isAlphanumeric(a) && b == '(') ||
		(a == ')' && isAlphanumeric(b)) ||
		((a == '*' || a == '+') && (isAlphanumeric(b) || b == '(')) || (a == ')' && b == '(')
}

func isAlphanumeric(c byte) bool {
	r := rune(c)
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func getOpPrecedence(r rune) int {
	specialCharsPriorityMap := map[rune]int{
		'(': 1,
		'|': 2,
		'.': 3,
		'?': 4,
		'*': 4,
		'+': 4,
	}

	if priority, ok := specialCharsPriorityMap[r]; ok {
		return priority
	}
	return maxPriority + 1
}

func InfixToPostfix(infix string) string {
	infix = insertConcatOperators(infix)

	var (
		postfix []rune
		stack   []rune
	)

	for _, r := range infix {
		switch r {
		case '(':
			stack = append(stack, r)
		case ')':
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			for len(stack) > 0 && getOpPrecedence(stack[len(stack)-1]) >= getOpPrecedence(r) {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, r)
		}

	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return string(postfix)
}
