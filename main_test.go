package bytes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes_Ceil(t *testing.T) {
	testCases := []struct {
		Input     Bytes
		Precision Bytes
		Expected  Bytes
	}{
		{
			Input:     150 * Megabyte,
			Precision: Gigabyte,
			Expected:  Gigabyte,
		},
		{
			Input:     251 * Megabyte,
			Precision: Kilobyte,
			Expected:  251 * Megabyte,
		},
		{
			Input:     0,
			Precision: Kilobyte,
			Expected:  0,
		},
		{
			Input:     1400 * Terabyte,
			Precision: Petabyte,
			Expected:  2 * Petabyte,
		},
		{
			Input:     -2250 * Gigabyte,
			Precision: Terabyte,
			Expected:  -2 * Terabyte,
		},
	}

	for _, tCase := range testCases {
		testName := fmt.Sprintf("input=%d; precision=%d", tCase.Input, tCase.Precision)
		t.Run(testName, func(t *testing.T) {
			actual := tCase.Input.Ceil(tCase.Precision)
			assert.Equal(t, tCase.Expected, actual)
		})
	}
}

func TestBytes_Floor(t *testing.T) {
	testCases := []struct {
		Input     Bytes
		Precision Bytes
		Expected  Bytes
	}{
		{
			Input:     150 * Megabyte,
			Precision: Gigabyte,
			Expected:  0,
		},
		{
			Input:     251 * Megabyte,
			Precision: Kilobyte,
			Expected:  251 * Megabyte,
		},
		{
			Input:     0,
			Precision: Kilobyte,
			Expected:  0,
		},
		{
			Input:     1400 * Terabyte,
			Precision: Petabyte,
			Expected:  1 * Petabyte,
		},
		{
			Input:     -2250 * Gigabyte,
			Precision: Terabyte,
			Expected:  -3 * Terabyte,
		},
	}

	for _, tCase := range testCases {
		testName := fmt.Sprintf("input=%d; precision=%d", tCase.Input, tCase.Precision)
		t.Run(testName, func(t *testing.T) {
			actual := tCase.Input.Floor(tCase.Precision)
			assert.Equal(t, tCase.Expected, actual)
		})
	}
}

func TestBytes_Round(t *testing.T) {
	testCases := []struct {
		Input     Bytes
		Precision Bytes
		Expected  Bytes
	}{
		{
			Input:     150 * Megabyte,
			Precision: Gigabyte,
			Expected:  0,
		},
		{
			Input:     251 * Megabyte,
			Precision: Kilobyte,
			Expected:  251 * Megabyte,
		},
		{
			Input:     0,
			Precision: Kilobyte,
			Expected:  0,
		},
		{
			Input:     1400 * Terabyte,
			Precision: Petabyte,
			Expected:  1 * Petabyte,
		},
		{
			Input:     500 * Kilobyte,
			Precision: Megabyte,
			Expected:  1 * Megabyte,
		},
		{
			Input:     -2250 * Gigabyte,
			Precision: Terabyte,
			Expected:  -2 * Terabyte,
		},
	}

	for _, tCase := range testCases {
		testName := fmt.Sprintf("input=%d; precision=%d", tCase.Input, tCase.Precision)
		t.Run(testName, func(t *testing.T) {
			actual := tCase.Input.Round(tCase.Precision)
			assert.Equal(t, tCase.Expected, actual)
		})
	}
}

func TestBytes_String(t *testing.T) {
	testCases := []struct {
		Input    Bytes
		Expected string
	}{
		{
			Input:    0,
			Expected: "0 B",
		},
		{
			Input:    1250 * Gigabyte,
			Expected: "1.25 TB",
		},
		{
			Input:    3_743_500 * Byte,
			Expected: "3.7435 MB",
		},
		{
			Input:    -3_743_500 * Byte,
			Expected: "-3.7435 MB",
		},
	}

	for _, tCase := range testCases {
		testName := fmt.Sprintf("input=%d", tCase.Input)
		t.Run(testName, func(t *testing.T) {
			actual := tCase.Input.String()
			assert.Equal(t, tCase.Expected, actual)
		})
	}
}
