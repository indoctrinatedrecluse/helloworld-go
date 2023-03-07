package helloworld-go

import "fmt"
import "rsc.io/quote"

func helloworld() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
}