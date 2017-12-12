package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestSum(t *testing.T) {

	a := 20
	b := 30
	c := Sum(a, b)
	assert.Equal(t,a+b, c)

}
