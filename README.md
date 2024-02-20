## Hexabloom

Hexabloom is a lightweight Golang package to check if an address is within a bloom filter.

An example code would be:
```
package main

import (
	"fmt"

	"github.com/hexagateinc/hexabloom"
)

func main() {
	bloomFilterClient, _ := hexabloom.NewBloomFilterClientFromFile("bloom.bin")
	fmt.Println(bloomFilterClient.ContainsStr("0x910cbd523d972eb0a6f4cae4618ad62622b39dbf")) // Example usage of containsStr
}
```

`$ go run main.go`
