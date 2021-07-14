# Blazinly fast PRNG (normal distributed)

> Generate 8 bytes (64 bit) PRNG identifiers blazingly fast

## Usage

```go
import "github.com/zerjioang/fastprng"
```

## Example Program

> Example program that creates a PRNG UUID of 64 bits encoded as hex string

```go
package main

import (
	"fmt"
	"github.com/zerjioang/prng"
)

func main(){
	uuid := prng.New()
	fmt.Println(uuid)
}
```

## Performance

**Using standard Go package**

```go
func gen() string {
	dst := make([]byte, 8)
	for i:=0; i<8;i++ {
		dst[i] = byte(rand.Intn(256))
	}
	return hex.EncodeToString(dst[:])
}
// Output: 336159b54b9e2839
```

**Using `fastprng` package**

```go
func gen() string {
    return fastprng.New()
}
// Output: dca90c93b8848229
```

## Performance Results

> Note: always run your own performance test on your hardware

```bash
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz

BenchmarkRng-0-12  41574139  56.28 ns/op  17.77 MB/s  0 B/op  0 allocs/op
BenchmarkRng-0-12  43105999  55.83 ns/op  17.91 MB/s  0 B/op  0 allocs/op
BenchmarkRng-0-12  42470292  55.33 ns/op  18.07 MB/s  0 B/op  0 allocs/op
BenchmarkRng-0-12  44294082  55.39 ns/op  18.05 MB/s  0 B/op  0 allocs/op
BenchmarkRng-0-12  43065757  55.16 ns/op  18.13 MB/s  0 B/op  0 allocs/op
BenchmarkRng-0-12  43489963  54.92 ns/op  18.21 MB/s  0 B/op  0 allocs/op
BenchmarkRng-0-12  40353843  55.29 ns/op  18.09 MB/s  0 B/op  0 allocs/op
BenchmarkRng-0-12  44173886  55.38 ns/op  18.06 MB/s  0 B/op  0 allocs/op
BenchmarkRng-0-12  44519041  54.78 ns/op  18.26 MB/s  0 B/op  0 allocs/op
BenchmarkRng-0-12  43383703  55.69 ns/op  17.96 MB/s  0 B/op  0 allocs/op
```

### Speedup calculation

**Spoiler alert**: +970.22% CPU and zero allocations

```bash
go test -v -run=^$ -bench=^BenchmarkRng/\$ -benchtime=2s -count=10 > old.txt
```

Make some changes

```bash
go test -v -run=^$ -bench=^BenchmarkRng/\$ -benchtime=2s -count=10 > new.txt
```

And then, compare results with `benchstat`

```bash
benstat old.txt new.txt
```

```bash
name           old time/op    new time/op     delta
Rng/uuid-0-12     593ns ± 2%       55ns ± 2%   -90.66%  (p=0.000 n=9+10)

name           old speed      new speed       delta
Rng/uuid-0-12  1.69MB/s ± 2%  18.05MB/s ± 2%  +970.22%  (p=0.000 n=9+10)

name           old alloc/op   new alloc/op    delta
Rng/uuid-0-12     32.0B ± 0%       0.0B       -100.00%  (p=0.000 n=10+10)

name           old allocs/op  new allocs/op   delta
Rng/uuid-0-12      2.00 ± 0%       0.00       -100.00%  (p=0.000 n=10+10)
```
