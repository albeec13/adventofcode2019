package main

import (
   "fmt"
   "log"
   "strings"
   "strconv"
)

func main() {
    if !(len(input) > 0) {
        log.Fatal("Bad input file.")
    }

    fuelTotal := 0
    modules := strings.Split(input, "\n")

    for _,module := range modules {
        if moduleInt,err := strconv.Atoi(module); err != nil {
            log.Fatal(err)
        } else {
            fuelTotal += (moduleInt / 3 - 2)
        }
    }
    fmt.Println("Total Fuel: ", fuelTotal)
}
