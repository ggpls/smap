# smap | Slice Mapper

type conversion utilities for slices that smaps 🖐️

## Installation

```bash
go get github.com/ggpls/smap
```

## Usage

### Number Type Conversions

Convert slices between different numeric types (int, float32, float64, etc.):

```go
package main

import "github.com/ggpls/smap"

func main() {
    // Convert []int to []float64
    islice := []int{1, 2, 3, 4, 5}
    f64slice := smap.N[int, float64](islice)

    // Convert []float32 to []int
    f32slice := []float32{1.1, 2.2, 3.3}
    islice = smap.N(f32slice)
}
```

### Byte/String Type Conversions

Convert between `[]byte` and `string` slices:

```go
package main

import "github.com/ggpls/smap"

func main() {
    // convert []string to [][]byte
    strs := []string{"hello", "world"}
    byteses := smap.S[string, []byte](strs)

    // and back
    strs = smap.S(byteses)
}
```

### Custom Function Mapping

Transform slices using custom functions with the `F` mapper:

```go
package main

import (
    "strings"
    "github.com/ggpls/smap"
)

func main() {
    // Using standard library functions
    words := []string{"hello", "world"}
    upperWords := smap.F(words, strings.ToUpper)
    // Result: ["HELLO", "WORLD"]

    // Using custom transformation functions
    numbers := []int{1, 2, 3, 4}
    squared := smap.F(numbers, func(x int) int {
        return x * x
    })
    // Result: [1, 4, 9, 16]

    // Type conversion with formatting
    ids := []int{1, 2, 3}
    formatted := smap.F(ids, func(id int) string {
        return fmt.Sprintf("ID-%d", id)
    })
    // Result: ["ID-1", "ID-2", "ID-3"]
}
```

## Type Constraints

The package uses three main approaches:

- `N`: Supports all integer and floating-point types (using `constraints.Integer | constraints.Float`)
- `S`: Supports `[]byte` and `string` types
- `F`: Supports any types that satisfy the provided transformation function

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
