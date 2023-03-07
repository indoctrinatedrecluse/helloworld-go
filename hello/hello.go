package main

import (
    "fmt"
    "github.com/helloworld-go/greetings"
)
    
func main() {
    // import from another package
    message := greetings.Hello("indoctrinatedrecluse")
    fmt.Println(message)
}