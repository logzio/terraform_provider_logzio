package logzio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripWhitespaceReturnsDesiredResult(t *testing.T) {
	input := "test string with	tabs"
	expected := "teststringwithtabs"

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}

func TestStripWhitespaceReturnsDesiredResultWhenInputStringIsEmpty(t *testing.T) {
	input := ""
	expected := ""

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}

func TestStripWhitespaceReturnsDesiredResultWhenInputStringIsJustSpaces(t *testing.T) {
	input := "        "
	expected := ""

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}


func TestStripWhitespaceReturnsDesiredResultWhenInputStringIsJustTabs(t *testing.T) {
	input := "				"
	expected := ""

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}

func TestStripWhitespaceReturnsDesiredResultWhenInputStringIsJustCharacters(t *testing.T) {
	input := "abcdefghijklmnoplqrstuvwxyz"
	expected := "abcdefghijklmnoplqrstuvwxyz"

	assert.EqualValues(t, expected, stripAllWhitespace(input))
}