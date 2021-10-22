package main

import (
	"testing"

	"github.com/borischen0203/litclock/cmd"
	"github.com/stretchr/testify/assert"
)

var InvalidInput = "Invalid input"

//Should return "Invalid input" when parameter amount is over one.
func TestInputOverOneParameter(t *testing.T) {
	input := [...]string{"11:11", "0"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in parameter amount is over one")
}

//Should return "Invalid input" when input hour is over 23.
func TestInputInvalidTime1(t *testing.T) {
	input := [...]string{"24:00"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input hour is over 23")
}

//Should return "Invalid input" when input hour is less 00.
func TestInputInvalidTime2(t *testing.T) {
	input := [...]string{"-01:00"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input hour is less 00")
}

//Should return "Invalid input" when input minute is over 59.
func TestInputInvalidTime3(t *testing.T) {
	input := [...]string{"1:60"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is over 59")
}

//Should return "Invalid input" when input minute is less 00.
func TestInputInvalidTime4(t *testing.T) {
	input := [...]string{"1:-1"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is less 00")
}

//Should return "Invalid input" when input hour and minute are invalid.
func TestInputInvalidTime5(t *testing.T) {
	input := [...]string{"24:60"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input hour and minute are invalid")
}

//Should return "Invalid input" when when input second field.
func TestInputInvalidTime6(t *testing.T) {
	input := [...]string{"01:01:05"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input second field")
}

//Should return "Invalid input" when when input uses the operator.
func TestInputInvalidTime7(t *testing.T) {
	input := [...]string{"1+1:1+1"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input uses the operator ")
}

//Should return "Invalid input" when input is not time type.
func TestInputInvalidType1(t *testing.T) {
	input := [...]string{"Hello World"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input is not time type")
}

//Should return "Invalid input" when input is symbol.
func TestInputInvalidType2(t *testing.T) {
	input := [...]string{"!@#$%^&*()"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input is symbol")
}

//Should return "Invalid input" when parameter is not numeric time.
func TestInvalidTime(t *testing.T) {
	input := [...]string{""}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input is not numeric time")
}

//Should return correct time text when input empty string.
func TestWithoutInput(t *testing.T) {
	input := [...]string{""}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input empty string")
}
