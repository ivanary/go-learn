package controller

import "testing"

func TestHelloWorld(t *testing.T) {
	hello := HelloWorld("ivanary")
	if hello != "Hi, ivanary" {
		t.Errorf("Testing fail")
	}

	hello = HelloWorld("ivanary ")
	if hello != "Hi, ivanary" {
		t.Errorf("Testing fail")
	}
}
