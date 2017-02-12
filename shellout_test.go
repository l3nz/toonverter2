package main

import (
	"fmt"
	"testing"
)

func TestShelloutTokenExpander(t *testing.T) {

	in := []string{
		"A", "B", "C",
	}

	val := expand_tokens(in)
	exp := "[A B C]"
	vals := fmt.Sprintf("%+v", val)

	if vals != exp {
		t.Errorf("Got %s expected %s", vals, exp)
	}
}

func TestShelloutTokenExpander_SingleToken(t *testing.T) {

	in := []string{
		"A", "B", "C", TOKEN_PREFIX + "DD", "E",
	}

	val := expand_tokens(in)
	exp := "[A B C DD E]"
	vals := fmt.Sprintf("%+v", val)

	if vals != exp {
		t.Errorf("Got %s expected %s", vals, exp)
	}
}

func TestShelloutTokenExpander_MultiToken(t *testing.T) {

	in := []string{
		"A", TOKEN_PREFIX + "BA BB", "C", TOKEN_PREFIX + "DD DE", "E",
	}

	val := expand_tokens(in)
	exp := "[A BA BB C DD DE E]"
	vals := fmt.Sprintf("%+v", val)

	if vals != exp {
		t.Errorf("Got %s expected %s", vals, exp)
	}
}
