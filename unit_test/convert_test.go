package unittest

import (
	"fmt"
	"gohss/utils"
	"testing"
)

func TestIntToHexFixedLength(t *testing.T) {
	tests := []struct {
		inputNum    uint64
		inputLen    int
		expectedHex string
	}{
		{123456789012, 12, "001CBE991A14"},
		{255, 6, "0000FF"},
		{0, 4, "0000"},
		{123, 8, "0000007B"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %d, Length: %d", test.inputNum, test.inputLen), func(t *testing.T) {
			result := utils.IntToHexFixedLength(test.inputNum, test.inputLen)
			if result != test.expectedHex {
				t.Errorf("Expected %s, but got %s", test.expectedHex, result)
			}
		})
	}
}
