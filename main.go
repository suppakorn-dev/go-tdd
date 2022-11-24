package main

import (
	"fmt"
	"go-tdd/captcha"
	"math/rand"
	"time"
)

const (
	maxPatternNumber     = 2
	maxOperatorNumber    = 3
	maxSideOperandNumber = 9
)

type captchaInput struct {
	pattern      int
	leftOperand  int
	operator     int
	rightOperand int
}

func randomNumber(max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	min := 1

	return r1.Intn(max-min+1) + min
}

func main() {
	input := captchaInput{
		pattern:      randomNumber(maxPatternNumber),
		leftOperand:  randomNumber(maxSideOperandNumber),
		operator:     randomNumber(maxOperatorNumber),
		rightOperand: randomNumber(maxSideOperandNumber),
	}

	captcha := captcha.New(input.pattern, input.leftOperand, input.operator, input.rightOperand)
	printInput(input)
	printExpression(captcha)
}

func printInput(c captchaInput) {
	fmt.Printf("pattern: %d, %d %d %d = ", c.pattern, c.leftOperand, c.operator, c.rightOperand)
}

func printExpression(c captcha.Captcha) {
	fmt.Printf("%s %s %s\n", c.LeftOperand(), c.Operator(), c.RightOperand())
}
