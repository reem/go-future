# Future [![Build Status](https://travis-ci.org/reem/go-future.svg?branch=master)](https://travis-ci.org/reem/go-future)

> A condition-variable based Future type for synchronization in Go.

Controls an accompanying pointer, like Mutex or Cond, rather than trying
to hack around the lack of generics.

## Example

```go
package main

import future "github.com/reem/go-future"
import "fmt"

type Data struct {
    int x
}

func main() {
    data := &Data{0}
    producer, consumer := future.Pair()

    go func() {
        data.x = 12
        producer.Complete()
    }()

    go func() {
        consumer.Await()
        fmt.Println("Received data, data.x ==", data.x)
    }()
}
```

## Author

[Jonathan Reem](https://medium.com/@jreem) is the primary author and maintainer of future.

## License

MIT

