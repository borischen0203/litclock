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

//Should return current human time text when input empty string.
func TestWithoutInput(t *testing.T) {
	input := [...]string{}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input empty string")
}

//Should return correct human text when input minute is 00.
func TestInputMinuteZero(t *testing.T) {
	input := [...]string{"13:00"}

	expected := "One o'clock"
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is 00")
}

//Should return correct human text when input minute is 1.
func TestInputMinuteOne(t *testing.T) {
	input := [...]string{"13:01"}

	expected := "One past one"
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is 1")
}

//Should return correct human text when input minute 59.
func TestInputMinute59(t *testing.T) {
	input := [...]string{"13:59"}

	expected := "One to two"
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is 59")
}

//Should return correct human text when input minute 15.
func TestInputMinute15(t *testing.T) {
	input := [...]string{"13:15"}

	expected := "Quarter past one"
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is 15")
}

//Should return correct human text when input minute 30.
func TestInputMinute30(t *testing.T) {
	input := [...]string{"13:30"}

	expected := "Half past one"
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is 30")
}

//Should return correct human text when input minute 45.
func TestInputMinutes45(t *testing.T) {
	input := [...]string{"13:45"}

	expected := "Quarter to two"
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is 45")
}

//Should return correct human text when input minute more than 30.
func TestInputMinuteMoreThan30(t *testing.T) {
	input := [...]string{"13:31"}

	expected := "Twenty nine to two"
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is 31")
}

//Should return correct human text when input minute less than 29.
func TestInputMinuteLess30(t *testing.T) {
	input := [...]string{"13:29"}

	expected := "Twenty nine past one"
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is 29")
}
