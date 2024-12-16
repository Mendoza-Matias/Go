# Go

- Start the module manager:
```bash
  go mod init <module-name>
```
- Download a module:
```bash
    go get rsc.io/quote@latest
```
- Run go:
```bash
    go run main.go
```
## Data Types in Go :
  Go provides various data types to handle different needs in your program. Below is a summary of the main ones:
### Integer
- int8, int16, int32, int64: Signed integers, ranges depend on the size.
- Example: int8 (-128 to 127).
- uint: Unsigned integers, only positive numbers.
### Float
- float32: For smaller decimal numbers.
- float64: For larger and more precise decimal numbers.
### Boolean
- bool: Represents true (true) or false (false) values.
### String
- string: Used for text data.
### Byte
- byte: Alias for uint8, commonly used to represent ASCII characters.
Rune
### rune: 
- Alias for int32, used to represent Unicode characters.