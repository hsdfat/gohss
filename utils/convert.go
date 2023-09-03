package utils

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	SQN_LENGTH = 12
)

func strictHex(s string, n int) string {
	l := len(s)
	if l < n {
		return fmt.Sprintf(strings.Repeat("0", n-l) + s)
	} else {
		return s[l-n : l]
	}
}

func hexToInt(hexString string, sqnLength int) (uint64, error) {
	strictHex(hexString, sqnLength)
	intValue, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	return uint64(intValue), nil
}

func IntToHexFixedLength(num uint64, length int) string {
	// Convert the integer to a hexadecimal string
	hexString := fmt.Sprintf("%012X", num)

	// If the hex string length exceeds the desired length, truncate it
	if len(hexString) > length {
		hexString = hexString[len(hexString)-length:]
	}

	// If the hex string length is less than the desired length, pad it with zeroes
	if len(hexString) < length {
		padding := strings.Repeat("0", length-len(hexString))
		hexString = padding + hexString
	}

	return hexString
}
