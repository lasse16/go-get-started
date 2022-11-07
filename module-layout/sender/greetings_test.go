package sender

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greet("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greet("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Greet("")
	if msg != "" || err == nil {
		t.Fatalf(`Greet("") = %q, %v, want "", error`, msg, err)
	}
}

func TestFail(t *testing.T) {
	msg, err := Greet("")
	t.Fatalf(`Greet("") = %q, %v, want "", error`, msg, err)
}
