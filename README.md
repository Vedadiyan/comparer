# Go Comparer Library

A flexible and type-safe comparison library for Go that enables comparing values across different numeric types and strings.

## Features

- Generic comparison function that works with multiple data types
- Type-safe comparisons between different numeric types
- Fallback string comparison for unsupported types
- Support for all major numeric types:
  - Integers (int, int8, int16, int32, int64)
  - Unsigned integers (uint, uint8/byte, uint16, uint32, uint64)
  - Floating-point numbers (float32, float64)
- String comparison support
- Zero external dependencies

## Installation

```bash
go get github.com/vedadiyan/comparer
```

## Usage

### Basic Comparison

The library provides a `Compare` function that returns:
- 0 if values are equal
- 1 if the first value is greater
- -1 if the first value is smaller

```go
import "github.com/vedadiyan/comparer"

func main() {
    // Compare integers
    result := comparer.Compare(5, 3)     // returns 1
    result = comparer.Compare(5, 5)      // returns 0
    result = comparer.Compare(3, 5)      // returns -1
    
    // Compare different numeric types
    result = comparer.Compare(int32(5), uint64(3))  // returns 1
    result = comparer.Compare(float32(5.0), int(5)) // returns 0
}
```

### Type-Specific Comparisons

You can also use type-specific comparers for more explicit comparisons:

```go
import "github.com/vedadiyan/comparer"

func main() {
    // Using Int type
    value := comparer.Int(5)
    result := value.Compare(3)           // returns 1
    result = value.Compare(uint32(5))    // returns 0
    result = value.Compare(float64(7.0)) // returns -1
    
    // Compare with strings
    result = value.Compare("5")          // returns 0
}
```

## Supported Type Conversions

The library handles comparisons between different numeric types by converting them to the base type before comparison. Here's the compatibility matrix:

| From Type | Can Compare With |
|-----------|-----------------|
| int       | All numeric types, string |
| int8/16/32/64 | All numeric types, string |
| uint/8/16/32/64 | All numeric types, string |
| float32/64 | All numeric types, string |

## Implementation Details

- The library uses type assertions and switches to handle different data types
- When comparing with unsupported types, it falls back to string comparison
- All numeric comparisons maintain type safety while allowing cross-type comparisons

## Error Handling

The library handles type mismatches gracefully by falling back to string comparison when types are incompatible or unknown. This ensures that the comparison function never panics and always returns a meaningful result.

## Performance Considerations

- The library uses type switches for type detection, which is more performant than reflection
- Direct type comparisons are used when possible
- String conversion is only used as a fallback for unsupported types

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

[Your chosen license]

## Author

[Your name/organization]
