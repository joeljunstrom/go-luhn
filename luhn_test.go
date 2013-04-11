package luhn

import (
	"math/rand"
	"testing"
	"time"
)

func TestValid(t *testing.T) {
	var test string = "1111111116"
	if !Valid(test) {
		t.Error("Expexted to be valid", test)
	}
}

func TestNotValid(t *testing.T) {
	var test string = "1111111111"
	if Valid(test) {
		t.Error("Expexted to be invalid", test)
	}
}

func TestGeneratesTheControlDigit(t *testing.T) {
	control_digit := generateControlDigit("811218987")
	if control_digit != 6 {
		t.Error("Expexted control_digit to be 6, was", control_digit)
	}
}

func TestGeneratesValidLuhn(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	var length int = rand.Intn(32) + 1
	generated := Generate(length)

	if len(generated) != length {
		t.Errorf("Expexted generated string to have length %d was %d", length, len(generated))
	}

	if !Valid(generated) {
		t.Error("Expected generated string to be valid luhn", generated)
	}
}
