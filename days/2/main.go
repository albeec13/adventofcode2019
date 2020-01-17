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

    // override postions 1 and 2 with error code from program alarm in puzzle description
    positions[1] = 12
    positions[2] = 02

    const (
        OPCODE = iota
        ARGS
        HALT
    )

    add := func(args []int) (int, int) {
        return positions[args[0]] + positions[args[1]], OPCODE
    }

	multiply := func(args []int) (int, int) {
        return positions[args[0]] * positions[args[1]], OPCODE
    }

	halt := func(args []int) (int, int) {
        return 0, HALT
    }

    type opdef struct {
        exec func(args []int) (int, int)
        argcnt int
    }

    operations := map[int]opdef {
        1 : {add,3},
        2 : {multiply,3},
        99: {halt,0},
    }

    state := OPCODE
    opcode := 0
    args := make([]int, 3)
    argcnt := 0
    for _, intcode := range positions {
        switch state {
        case OPCODE:
            opcode = intcode
            argcnt = operations[intcode].argcnt
            state = ARGS
        case ARGS:
            if argcnt > 0 {
                args[3 - argcnt] = intcode
                argcnt--
            }
            if argcnt == 0 {
                var retval int
                retval, state = operations[opcode].exec(args)
                fmt.Printf("%d - %d\n", state, retval)
                if state == HALT {
                    break;
                } else {
                    positions[args[2]] = retval
                }
            }
        }
    }

    // display result in position 0
    fmt.Printf("Position 0: %d\n", positions[0])
}
