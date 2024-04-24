# cache

## Example #1

```go
package main

import (
	"fmt"

	"github.com/YerzhanAkhmetov/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)

	userId := cache.Get("userId")

	fmt.Println(userId)

	//cache.Delete("userId")
	cache.Set("Yerzahn", 35)
	userId = cache.Get("userId")
	cache.Show()
	fmt.Println(userId)
}

```
