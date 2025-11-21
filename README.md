# testo

...

## Installation

```bash
go get github.com/orsinium-labs/testo
```

## Usage

```go
body := `{"name": "aragorn", "age": 87}`
pattern := `{"name": "aragorn", "age": int}`
err := testo.ValidateJSON(body, pattern)
if err != nil {
    t.Fatalf("invalid body: %v", err)
}
```

## Syntax

The pattern syntax is JSON with a few additional keywords:

* `true`, `false`, `null`: same as in JSON.
* `any`: a value of any type.
* `string`: any string of any length.
* `int`: an integer number.
* `uint`: a non-negative integer number.
* `float`: a floating point number.
* `bool`: a boolean value (`true` or `false`).
* `object`: any object.
* `array`: any JSON array.
* `strings`: array of strings (including empty array).
* `ints`: array of integer numberss (including empty array).
* `uints`: array of unsigned integer numbers (including empty array).
* `floats`: array of floating point numbers (including empty array).
* `bools`: array of boolean values (including empty array).
* `objects`: array of objects (including empty array).
