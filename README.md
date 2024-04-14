# Hexabloom

Hexabloom is a lightweight cross-language package to check if an address is within a bloom filter.
Addresses are lowercased strings, like one in the following example:

An example code would be:
```go
package main

import (
	"fmt"

	"github.com/hexagateinc/hexabloom/go"
)

func main() {
	bloomFilterClient, _ := hexabloom.NewBloomFilterClientFromFile("bloom.bin")
	fmt.Println(bloomFilterClient.ContainsStr("0x910cbd523d972eb0a6f4cae4618ad62622b39dbf")) // Example usage of containsStr
}
```

