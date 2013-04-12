package luhn

import (
	"math/rand"
	"strconv"
	"strings"
)

func Valid(luhn string) bool {
	checksumMod := calculateChecksum(luhn, false) % 10

	return checksumMod == 0
}

func Generate(targetSize int) string {
	targetSize--

	random := randomString(targetSize)
	controlDigit := strconv.Itoa(generateControlDigit(random))

	return random + controlDigit
}

func GenerateWithPrefix(targetSize int, prefix string) string {
	targetSize = targetSize - 1 - len(prefix)

	random := prefix + randomString(targetSize)
	controlDigit := strconv.Itoa(generateControlDigit(random))

	return random + controlDigit
}

func randomString(targetSize int) string {
	source := make([]int, targetSize)

	for i := 0; i < targetSize; i++ {
		source[i] = rand.Intn(9)
	}

	return integersToString(source)
}

func generateControlDigit(luhn string) int {
	controlDigit := calculateChecksum(luhn, true) % 10

	if controlDigit != 0 {
		controlDigit = 10 - controlDigit
	}

	return controlDigit
}

func calculateChecksum(luhn string, double bool) int {
	source := strings.Split(luhn, "")
	checksum := 0

	for i := len(luhn) - 1; i > -1; i-- {
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
	result := make([]string, len(integers))

	for i, number := range integers {
		result[i] = strconv.Itoa(number)
	}

	return strings.Join(result, "")
}
