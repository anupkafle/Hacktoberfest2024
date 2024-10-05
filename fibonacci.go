package main

        import "fmt"

        func fibonacciRecursive(n int) int {
            if n <= 1 {
                return n
            }
            return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
        }

        func main() {
            fmt.Println(fibonacciRecursive(9))  // Output: 34
        }
        
