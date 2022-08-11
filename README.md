# go-memoize
Impements a function memoization in Golang (1.18+)

# Installation

```sh
go get github.com/AlexandreChamard/go-memoize
```

# Examples

```golang
import (
    "math/rand"

    "github.com/AlexandreChamard/go-memoize"
)

var mfib func(int64) int64
fib := func(i int64) int64 {
    if i <= 0 {
        return 0
    }
    if i < 3 {
        return 1
    }
    return mfib(i-1) + mfib(i-2)
}
mfib = memoize.Memoize1(fib)

mfib(30) // expected 832040 in relatively low time
```

```golang
import (
    "math/rand"

    "github.com/AlexandreChamard/go-memoize"
)

var mfib func(int64) int64
fib := func(i int64) int64 {
    if i <= 0 {
        return 0
    }
    if i < 3 {
        return 1
    }
    return mfib(i-1) + mfib(i-2)
}

// It will cache a maximum of 100k items (least resut used will be delete in case it reach full)
mfib = memoize.Memoize1(fib, NewMemoizeOption().MaxMemoize(1<<20))

mfib(30) // expected 832040 in relatively low time
```
