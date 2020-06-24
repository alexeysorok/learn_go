package main

import (
	"strings"
	"bufio"
	"testing"
)

var testOk := '1
2
3
4
5
'

func TestOk(t *testing.T) {

	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq()
	if err != nil {
		t.Errorf("test for Ok Failed")
	}
	if out.String() != testOkResult{
		t.Errorf("test for OK Failed - result not match\n %v %v")
	}
}
