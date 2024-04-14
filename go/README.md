# Hexabloom - Golang

An example code would be:
```go
package main

import (
	"fmt"

	"github.com/hexagateinc/hexabloom/go"
)

func main() {
	bloomFilterClient, _ := hexabloom.NewBloomFilterClientFromFile("bloom.bin")
	fmt.Println(bloomFilterClient.ContainsStr("0x910cbd523d972eb0a6f4cae4618ad62622b39dbf")) // Example of containsStr, should print 'true'
}
```

Now you can get the dependencies and run your module with:

```bash
go mod init bloom
go mod tidy
go run main.go
```
