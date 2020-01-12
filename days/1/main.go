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

    for _,mass := range modules {
        if massInt,err := strconv.Atoi(mass); err != nil {
            log.Fatal(err)
        } else {
            fuelTotal += getFuel(massInt)
        }
    }
    fmt.Println("Total Fuel: ", fuelTotal)
}

func getFuel(mass int) int {
    fuel := mass / 3 - 2

    if (fuel > 0) {
        fuel += getFuel(fuel)
    } else {
        fuel = 0
    }

    return fuel
}
