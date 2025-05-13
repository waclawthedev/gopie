<div align="center">

<img height="190" src="https://github.com/waclawthedev/gopie/blob/main/gopie_logo.png">

<h1 align="center">GoPie</h1>

Useful functions ported to Go (map, reduce, filter etc.)

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/79a87ed9999b4682b8409e98d8e6b482)](https://app.codacy.com/gh/waclawthedev/gopie/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![codecov](https://codecov.io/gh/waclawthedev/gopie/graph/badge.svg?token=DQHPG80YRS)](https://codecov.io/gh/waclawthedev/gopie)

</div>

## Quick start

### Install

```sh
go get -u github.com/waclawthedev/gopie
```

### Import & use

```go
package main

import (
    "fmt"

    "github.com/waclawthedev/gopie"
)

func main() {
    set := gopie.NewSet[string]()

    set.Add("one")
    set.Add("two")

    set.Remove("one")
    fmt.Println(set.Contains("hello"))
}
```

## Functions list

✅ Don't hesitate to [create issue](https://github.com/waclawthedev/gopie/issues/new) with request of desired function

### Related to slices

*  Map - works in similar way as map function in Python
*  Filter - works in similar way as filter function in Python
*  Reduce - works in similar way as reduce function in Python
*  Flatten - works in similar way as flatten function in Python
*  SplitByN - splits slice by slices of size N

### Related to strings

*  TruncateString - returns first N runes of provided string with "..." at the end. Works properly even with emojis etc.
*  IsVisible - returns true if character is visible. Control characters considered as invisible. unicode.IsPrint & unicode.IsGraphic avoided to bypass bug described here: [https://github.com/golang/go/issues/73673](https://github.com/golang/go/issues/73673)
*  WithoutInvisibleChars - returns string without invisible chars

### Related to pointers

*  PtrTo - returns point to provided value. Useful to get pointer if you want to avoid variable declaration.
*  PtrValueOrDefault - safely gets value by pointer or returns provided default value if nil.
*  PtrValueOrZero - safely gets value by pointer or returns zero value according to type

### Related to cryptography

*  MD5 - returns MD5 hash of string

### Related to functions

*  Must - Accepts results with err from another functions and calls panic if err!=nil. Example: `gopie.Must(json.Marshal(...))`

### Set implementation

*  NewSet - creates set of selected type. Attention: it is thread unsafe implementation.
*  Add - adds element to set. Does nothing if such element already exists
*  Remove - removes element from set. Does nothing if such element doesn't exist
*  Contains - returns true if set contains provided element

## How to contribute

1. Fork this repository
2. Create a new branch with prefix in name "feat/" or "fix/"
3. Commit your changes
4. Push to your fork
5. Open a Pull Request to this original repository
