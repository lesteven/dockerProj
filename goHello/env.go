package main
import (
    "fmt"
    "os"
)

func main() {
    val := os.Getenv("VAL")
    fmt.Printf("hello " + val + "\n")
}
