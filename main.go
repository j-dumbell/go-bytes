// Package bytes provides functionality for ergonomically representing and
// converting quantities of bytes in the range -2^63 bytes to 2^63 bytes.
package bytes

import (
	"fmt"
	"math"
)

// Bytes represents a signed 64-bit integer number of bytes. Inspired by time.Duration.
type Bytes int64

// Common base 10 and base 2 units of bytes.
const (
	Byte     Bytes = 1
	Kilobyte Bytes = 1_000 * Byte
	Megabyte Bytes = 1_000 * Kilobyte
	Gigabyte Bytes = 1_000 * Megabyte
	Terabyte Bytes = 1_000 * Gigabyte
	Petabyte Bytes = 1_000 * Terabyte

	Kibibyte Bytes = 1_024 * Byte
	Mebibyte Bytes = 1_024 * Kibibyte
	Gibibyte Bytes = 1_024 * Mebibyte
	Tebibyte Bytes = 1_024 * Gibibyte
	Pebibyte Bytes = 1_024 * Tebibyte
)

// Kilobytes returns the bytes as a fractional number of kilobytes.
func (bytes Bytes) Kilobytes() float64 {
	return float64(bytes) / float64(Kilobyte)
}

// Megabytes returns the bytes as a fractional number of megabytes.
func (bytes Bytes) Megabytes() float64 {
	return float64(bytes) / float64(Megabyte)
}

// Gigabytes returns the bytes as a fractional number of gigabytes.
func (bytes Bytes) Gigabytes() float64 {
	return float64(bytes) / float64(Gigabyte)
}

// Terabytes returns the bytes as a fractional number of terabytes.
func (bytes Bytes) Terabytes() float64 {
	return float64(bytes) / float64(Terabyte)
}

// Petabytes returns the bytes as a fractional number of petabytes.
func (bytes Bytes) Petabytes() float64 {
	return float64(bytes) / float64(Petabyte)
}

// Kibibytes returns the bytes as a fractional number of kibibytes.
func (bytes Bytes) Kibibytes() float64 {
	return float64(bytes) / float64(Kibibyte)
}

// Mebibytes returns the bytes as a fractional number of mebibytes.
func (bytes Bytes) Mebibytes() float64 {
	return float64(bytes) / float64(Mebibyte)
}

// Gibibytes returns the bytes as a fractional number of gibibytes.
func (bytes Bytes) Gibibytes() float64 {
	return float64(bytes) / float64(Gibibyte)
}

// Tebibytes returns the bytes as a fractional number of tebibytes.
func (bytes Bytes) Tebibytes() float64 {
	return float64(bytes) / float64(Tebibyte)
}

// Pebibytes returns the bytes as a fractional number of pebibytes.
func (bytes Bytes) Pebibytes() float64 {
	return float64(bytes) / float64(Pebibyte)
}

// Round rounds the number of bytes to the nearest multiple of precision.
func (bytes Bytes) Round(precision Bytes) Bytes {
	if precision < 0 {
		return bytes
	}

	return Bytes(math.Round(float64(bytes)/float64(precision))) * precision
}

// Floor rounds the number of bytes down to the nearest multiple of precision.
func (bytes Bytes) Floor(precision Bytes) Bytes {
	if precision < 0 {
		return bytes
	}

	return Bytes(math.Floor(float64(bytes)/float64(precision))) * precision
}

// Ceil rounds the number of bytes up to the nearest multiple of precision.
func (bytes Bytes) Ceil(precision Bytes) Bytes {
	if precision < 0 {
		return bytes
	}

	return Bytes(math.Ceil(float64(bytes)/float64(precision))) * precision
}

// String returns a string representation of the number of bytes in base 10 units.
func (bytes Bytes) String() string {
	unitIndex := 0
	fractionalBytes := float64(bytes)
	for math.Abs(fractionalBytes) > 1_000 && unitIndex < len(units)-1 {
		fractionalBytes /= 1_000
		unitIndex++
	}

	return fmt.Sprintf("%g %s", fractionalBytes, units[unitIndex])
}

var units = []string{"B", "KB", "MB", "GB", "TB", "PB"}
