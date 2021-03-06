package main

import (
	"testing"
	"time"

	"github.com/borischen0203/litclock/cmd"
	"github.com/borischen0203/litclock/services"
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

//Should return "Invalid input" when input empty string.
func TestInputEmptyString(t *testing.T) {
	input := [...]string{""}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input empty string")
}

//Should return current human time text when input without parameter.
func TestWithoutInput(t *testing.T) {
	input := [...]string{}

	current := time.Now()
	expected := services.TimeToWords(current.Hour(), current.Minute())
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input without parameter")
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

//Should return correct human text when input minute less than quater.
func TestInputMinuteLess15(t *testing.T) {
	input := [...]string{"13:14"}

	expected := "Fourteen past one"
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is less than quater")
}

//Should return correct human text when input minute more than quater.
func TestInputMinuteLess16(t *testing.T) {
	input := [...]string{"13:16"}

	expected := "Sixteen past one"
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input minute is more than quater")
}

//Should return the same human text when input hour in 1 and 13.
func TestInputHour13and01(t *testing.T) {
	input01 := [...]string{"01:01"}
	input13 := [...]string{"13:01"}

	expected := cmd.NumericToText(input01[:]...)
	actual := cmd.NumericToText(input13[:]...)

	assert.Equal(t, expected, actual, "error in input hour in 1 and 13")
}

//Should return the correct human text when input hour is 12 and 00.
func TestInputHour12and00(t *testing.T) {
	input12 := [...]string{"12:01"}
	input00 := [...]string{"00:01"}

	expected := cmd.NumericToText(input12[:]...)
	actual := cmd.NumericToText(input00[:]...)

	assert.Equal(t, expected, actual, "error in input hour is 12 and 00")
}

//Should return the correct human text when input hour is 12 and 0.
func TestInputHour12and0(t *testing.T) {
	input12 := [...]string{"12:00"}
	input0 := [...]string{"0:00"}

	expected := cmd.NumericToText(input12[:]...)
	actual := cmd.NumericToText(input0[:]...)

	assert.Equal(t, expected, actual, "error in input hour in 12 and 0")
}

//Should return the "Invalid input" when input min 0.
func TestInputMin0(t *testing.T) {
	input := [...]string{"0:0"}

	expected := InvalidInput
	actual := cmd.NumericToText(input[:]...)

	assert.Equal(t, expected, actual, "error in input min 0")
}

//Should return the correct human text when input hour is 01 and 1.
func TestInputHour01and1(t *testing.T) {
	input01 := [...]string{"01:00"}
	input1 := [...]string{"1:00"}

	expected := cmd.NumericToText(input01[:]...)
	actual := cmd.NumericToText(input1[:]...)

	assert.Equal(t, expected, actual, "error in input hour in 01 and 1")
}
