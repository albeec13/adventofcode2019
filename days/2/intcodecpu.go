package main 

const (
    OPCODE = iota
    ARGS
    HALT
)

// private structs
type instruction struct {
    exec func(args []int) (int, int)
    argcnt int
}

type Intcodecpu struct {
    memory []int
    instTable map[int]instruction
}

func (i *Intcodecpu) Init(meminput []int) {
    i.instTable = map[int] instruction {
        1 : {i.add, 3},
        2 : {i.multiply, 3},
        99: {i.halt, 0},
    }

    i.memory = make([]int, len(meminput))
    copy(i.memory, meminput)
}

func (i *Intcodecpu) add(args []int) (int, int) {
        return i.memory[args[0]] + i.memory[args[1]], OPCODE
}

func (i *Intcodecpu)	multiply(args []int) (int, int) {
        return i.memory[args[0]] * i.memory[args[1]], OPCODE
}

func (i *Intcodecpu)	halt(args []int) (int, int) {
        return 0, HALT
}

func (i *Intcodecpu) Process() int {
    state := OPCODE
    opcode := 0
    args := make([]int, 3)
    argcnt := 0

    for _, inst := range i.memory {
        switch state {
        case OPCODE:
            opcode = inst
            argcnt = i.instTable[inst].argcnt
            state = ARGS
        case ARGS:
            if argcnt > 0 {
                args[3 - argcnt] = inst
                argcnt--
            }
            if argcnt == 0 {
                var retval int
                retval, state = i.instTable[opcode].exec(args)
                if state == HALT {
                    break;
                } else {
                    i.memory[args[2]] = retval
                }
            }
        }
    }

    return i.memory[0]
}
