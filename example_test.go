package smap_test

import (
	"fmt"
	"strings"

	"github.com/ggpls/smap"
)

func ExampleN_intToFloat64() {
	ints := []int{1, 2, 3, 4, 5}
	floats64 := smap.N[int, float64](ints)
	fmt.Printf("%v -> %v\n", ints, floats64)
	// Output: [1 2 3 4 5] -> [1 2 3 4 5]
}

func ExampleN_float32ToInt() {
	floats32 := []float32{1.9, 2.2, 3.7, 4.1, 5.9}
	ints := smap.N[float32, int](floats32)
	fmt.Printf("%v -> %v\n", floats32, ints)
	// Output: [1.9 2.2 3.7 4.1 5.9] -> [1 2 3 4 5]
}

func ExampleN_int32ToInt64() {
	int32s := []int32{100, 200, 300}
	int64s := smap.N[int32, int64](int32s)
	fmt.Printf("%v -> %v\n", int32s, int64s)
	// Output: [100 200 300] -> [100 200 300]
}

func ExampleN_float64ToFloat32() {
	float64s := []float64{1.123456789, 2.123456789}
	float32s := smap.N[float64, float32](float64s)
	fmt.Printf("%.6f -> %.6f\n", float64s[0], float32s[0])
	// Output: 1.123457 -> 1.123457
}

func ExampleS_stringToBytes() {
	strings := []string{"Hello", "World"}
	bytes := smap.S[string, []byte](strings)
	// Print length of each byte slice to avoid platform-specific byte representations
	fmt.Printf("%q -> [%d, %d] (byte lengths)\n", strings, len(bytes[0]), len(bytes[1]))
	// Output: ["Hello" "World"] -> [5, 5] (byte lengths)
}

func ExampleS_bytesToString() {
	byteSlices := [][]byte{
		[]byte("Golang"),
		[]byte("Rules"),
	}
	strings := smap.S[[]byte, string](byteSlices)
	fmt.Printf("%q\n", strings)
	// Output: ["Golang" "Rules"]
}

func ExampleS_utf8Handling() {
	// Test UTF-8 round trip conversion
	original := []string{"café", "über", "🚀"}
	bytes := smap.S[string, []byte](original)
	result := smap.S[[]byte, string](bytes)
	fmt.Printf("%q -> %q\n", original, result)
	// Output: ["café" "über" "🚀"] -> ["café" "über" "🚀"]
}

func ExampleF_stringTransform() {
	// Transform strings to uppercase
	words := []string{"hello", "world"}
	upper := smap.F(words, strings.ToUpper)
	fmt.Printf("%q -> %q\n", words, upper)
	// Output: ["hello" "world"] -> ["HELLO" "WORLD"]
}

func ExampleF_customTransform() {
	// Custom transformation function
	numbers := []int{1, 2, 3, 4}
	squared := smap.F(numbers, func(x int) int {
		return x * x
	})
	fmt.Printf("%v -> %v\n", numbers, squared)
	// Output: [1 2 3 4] -> [1 4 9 16]
}

func ExampleF_typeConversion() {
	// Convert integers to their string representation
	numbers := []int{42, 100, 999}
	formatted := smap.F(numbers, func(n int) string {
		return fmt.Sprintf("Number: %d", n)
	})
	fmt.Printf("%v -> %q\n", numbers, formatted)
	// Output: [42 100 999] -> ["Number: 42" "Number: 100" "Number: 999"]
}
