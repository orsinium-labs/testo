# testo

Go testing framework for matching a JSON payload against a pattern.

The pattern is a subset of JSON with some additional keywords that is then converted into a [valdo](github.com/orsinium-labs/valdo) validator.

## Installation

```bash
go get github.com/orsinium-labs/testo
```

## Usage

```go
func TestMyCoolAPI(t *testing.T) {
    body := `{"name": "aragorn", "age": 87}`
    testo.Assert(t, body, `{"name": "aragorn", "age": int}`)
}
```

The `body` can be bytes, string, or `io.Reader` (for example, an HTTP response body).

## Syntax

The pattern syntax is a suparset JSON with a few additional features.

Keywords:

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

If a property name starts with `^`, it's interpreted as a regular expression. For example, the following pattern defines an object with non-empty unsigned integer numbers as keys and integer values:

```json
{"^[0-9]+$": int}
```

So, if you want to assert only one property of an object:

```json
{
    "name": "Aragorn",
    "^.+$": any,
}
```

An object can have any number or properties and regex properties in any combination.
