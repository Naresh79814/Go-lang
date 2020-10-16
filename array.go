package main
import "fmt"

func main() {  
    var numbers [3] string  
    numbers[0] = "One"
    numbers[1] = "Two"
    numbers[2] = "Three"
    fmt.Println(numbers[1]) 
    fmt.Println(len(numbers)) 
    fmt.Println(numbers) 

    directions := [...] int {1,2,3,4,5}  
    fmt.Println(directions) 
    fmt.Println(len(directions))
}