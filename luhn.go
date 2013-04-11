package luhn

import (
	"math/rand"
	"strconv"
	"strings"
)

func Valid(luhn string) bool {
	checksum := calculateChecksum(luhn, false)

	return (checksum % 10) == 0
}

func Generate(target_size int) string {
	var random_size = target_size - 1
	var source = make([]int, random_size)

	for i := 0; i < random_size; i++ {
		source[i] = rand.Intn(9)
	}

	temp := integersToString(source)

	control_digit := generateControlDigit(temp)

	generated := temp + strconv.Itoa(control_digit)

	return generated
}

func generateControlDigit(luhn string) int {
	checksum := calculateChecksum(luhn, true)
	control_digit := checksum % 10

	if control_digit != 0 {
		control_digit = 10 - control_digit
	}

	return control_digit
}

func calculateChecksum(luhn string, double bool) int {
	source := strings.Split(luhn, "")
	var checksum int = 0
	var length = len(luhn) - 1

	for i := length; i > -1; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)

		if double {
			n = n * 2
		}
		double = !double

		if n >= 10 {
			n = n - 9
		}

		checksum += n
	}

	return checksum
}

func integersToString(integers []int) string {
	length := len(integers)
	var result = make([]string, length)

	for i := 0; i < length; i++ {
		result[i] = strconv.Itoa(integers[i])
	}

	return strings.Join(result, "")
}
