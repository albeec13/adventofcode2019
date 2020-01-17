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

    inputs := strings.Split(input, ",")
    positions := make([]int, len(inputs))

    // convert input string array to a int slice
    for i, input := range inputs {
        var err error
        if positions[i],err = strconv.Atoi(input); err != nil {
            log.Fatal(err)
        }
    }

    noun := 0
    verb := 0
    for noun = 0; noun < 100; noun++ {
        result := 0
        for verb = 0; verb < 100; verb++ {
            // override postions 1 and 2 with error code from program alarm in puzzle description
            positions[1] = noun
            positions[2] = verb

            var cpu Intcodecpu
            cpu.Init(positions)
            result = cpu.Process()

            if result == 19690720 {
                break;
            }
        }
        if result == 19690720 {
            fmt.Printf("Initital condition: %d\n", noun * 100 + verb)
            break;
        }
    }

}
