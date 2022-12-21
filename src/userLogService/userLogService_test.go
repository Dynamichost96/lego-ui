package userLogService

import (
	"testing"
)

var sample = New()

func defaultServiceMustNotFail(t *testing.T) {
	sample.Write("hello") 
}