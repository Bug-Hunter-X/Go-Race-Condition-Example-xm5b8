```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex // Add mutex for synchronization

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mu.Lock() // Lock before accessing shared variable
                        count += i
                        mu.Unlock() // Unlock after accessing shared variable
                }(i)
        }

        wg.Wait()
        fmt.Println(count)
}
```