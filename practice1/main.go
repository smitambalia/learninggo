package main

import (
    "fmt"
    "strconv"
)

// Function that takes a pointer to an int and modifies the original value
func changeValue(val *int) {
    fmt.Printf("Inside changeValue - Before change, val (address): %p, *val (value): %d\n", val, *val)
    *val = 100 // Dereference val and assign a new value
    fmt.Printf("Inside changeValue - After change, val (address): %p, *val (value): %d\n", val, *val)
}


/* func repeatingString(text string) {
    fmt.Println(text)
} */
// To run the server, use the command: go run main.go
func main() {
    /* myNumber := 42
    fmt.Printf("In main - Before calling changeValue, myNumber: %d, &myNumber (address): %p\n", myNumber, &myNumber)

    // Pass the memory address of myNumber to the function
    changeValue(&myNumber)

    // titleString := "### Hello, World! in Go language ###"
    fmt.Printf("In main - After calling changeValue, myNumber: %d\n", myNumber) // myNumber will now be 100
 */
     demoSliceOfStrong  := []string{"Welcome" , "to the world" }

    for _,item := range demoSliceOfStrong {
        fmt.Println(item)
    }

    var (
        firstName  string  = "Smit"
        lastName   string  = "Ambalia"
        age        int     = 25
        isEmployed bool    = true
    )

    fmt.Println(firstName, lastName, age, isEmployed)
    // array of strings

    // var listOfStrings = []string {titleString, "1. Start learning Go language", "2. Learn the basics of Go language", "3. Make to-do list in Go language", "4. Learn about Go modules", "5. Learn about Go packages", "6. Learn about Go functions", "7. Learn about Go methods", "8. Learn about Go interfaces", "9. Learn about Go structs", "10. Learn about Go maps", "11. Learn about Go slices", "12. Learn about Go arrays", "13. Learn about Go channels", "14. Learn about Go goroutines", "15. Learn about Go concurrency", "16. Learn about Go error handling", "17. Learn about Go testing", "18. Learn about Go debugging", "19. Learn about Go profiling", "20. Learn about Go performance tuning", "21. Learn about Go memory management", "22. Learn about Go garbage collection", "23. Learn about Go reflection", "24. Learn about Go code generation", "25. Learn about Go code analysis", "26. Learn about Go code optimization", "27. Learn about Go code review", "28. Learn about Go code style", "29. Learn about Go code conventions", "30. Learn about Go code standards"}
    // 
   /*  for _, str := range listOfStrings {
        fmt.Println(str)
    }
 */
    // for i:=0 ; i < len(listOfStrings); i++ {
    //     fmt.Println(listOfStrings[i])
    // }

    // fmt.Println(listOfStrings)
    /* fmt.Println("2. Learn the basics of Go language")
    fmt.Println("3. Make to-do list in Go language")
    fmt.Println("4. Learn about Go modules")
    fmt.Println("5. Learn about Go packages")
    fmt.Println("6. Learn about Go functions")
    fmt.Println("7. Learn about Go methods")
    fmt.Println("8. Learn about Go interfaces")
    fmt.Println("9. Learn about Go structs")
    fmt.Println("10. Learn about Go maps")
    fmt.Println("11. Learn about Go slices")
    fmt.Println("12. Learn about Go arrays")
    fmt.Println("13. Learn about Go channels")
    fmt.Println("14. Learn about Go goroutines")
    fmt.Println("15. Learn about Go concurrency")
    fmt.Println("16. Learn about Go error handling")
    fmt.Println("17. Learn about Go testing")
    fmt.Println("18. Learn about Go debugging")
    fmt.Println("19. Learn about Go profiling")
    fmt.Println("20. Learn about Go performance tuning")
    fmt.Println("21. Learn about Go memory management")
    fmt.Println("22. Learn about Go garbage collection")
    fmt.Println("23. Learn about Go reflection")
    fmt.Println("24. Learn about Go code generation")
    fmt.Println("25. Learn about Go code analysis")
    fmt.Println("26. Learn about Go code optimization")
    fmt.Println("27. Learn about Go code review")
    fmt.Println("28. Learn about Go code style")
    fmt.Println("29. Learn about Go code conventions")
    fmt.Println("30. Learn about Go code standards")
    fmt.Println("31. Learn about Go code best practices")
    fmt.Println("32. Learn about Go code patterns")
    fmt.Println("33. Learn about Go code architecture")
    fmt.Println("34. Learn about Go code design")
    fmt.Println("35. Learn about Go code tools")
    fmt.Println("36. Learn about Go code libraries")
    fmt.Println("37. Learn about Go code frameworks")
    fmt.Println("38. Learn about Go code platforms")
    fmt.Println("39. Learn about Go code services") */
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello, World!")
    // })

    // fmt.Println((`Started the server on port 8080`))
    // err := http.ListenAndServe(":8080", nil)
    // if err != nil {
    //     log.Fatal(err)
    // }

    var i int = 99
    var j string

    j = strconv.Itoa(i) // Convert int to string

    fmt.Printf("Value of i is %v  and %T", i , i)
    fmt.Println("")
    fmt.Printf("Value of j is %v  and %T", j, j)

}

