package luhn

import (
	"testing"
)

func TestValid(t *testing.T) {
	test := "1111111116"

	if !Valid(test) {
		t.Errorf("Expexted ”%s” to be valid according to the Luhn algorithm", test)
	}

	test = "1111111111"

	if Valid(test) {
		t.Errorf("Expexted ”%s” to be invalid according to the Luhn algorithm", test)
	}
}

func TestGeneratesTheControlDigit(t *testing.T) {
	control_digit := generateControlDigit("811218987")

	if control_digit != 6 {
		t.Error("Expexted control_digit to equal 6, was", control_digit)
	}
}

func TestGeneratesValidLuhn(t *testing.T) {
	length := 32
	generated := Generate(length)

	if len(generated) != length {
		t.Errorf("Expexted generated string to have length %d was %d", length, len(generated))
	}

	if !Valid(generated) {
		t.Error("Expected generated string to be valid luhn", generated)
	}
}

func TestGeneratesLuhnWithPrefix(t *testing.T) {
	length := 10
	prefix := "12345"
	generated := GenerateWithPrefix(length, prefix)

	if len(generated) != length {
		t.Errorf("Expected generated string to have length %d was %d", length, len(generated))
	}

	generated_prefix := generated[0:5]
	if generated_prefix != prefix {
		t.Errorf("Expected prefix to equal %s, was %s", prefix, generated_prefix)
	}

	if !Valid(generated) {
		t.Error("Expected generated string to be valid luhn", generated)
	}
}
