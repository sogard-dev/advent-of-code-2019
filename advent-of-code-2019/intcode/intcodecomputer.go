package intcodecomputer

type IntCodeComputer struct {
	memory       map[int]int
	ip           int
	relativeBase int
	inputter     func() int
	outputter    func(int)
}

func New(initialMemory []int) IntCodeComputer {
	mem := map[int]int{}
	for i, v := range initialMemory {
		mem[i] = v
	}
	return IntCodeComputer{
		memory: mem,
	}
}

type op interface {
	execute(*IntCodeComputer)
}

type halt struct{}
type oneArg struct {
	a int
}

type twoArgs struct {
	a, b int
}

type threeArgs struct {
	a, b, c int
}

type input oneArg
type output oneArg
type adjustRelativeBase oneArg
type jmpIfTrue twoArgs
type jmpIfFalse twoArgs
type addition threeArgs
type multiplication threeArgs
type lessThan threeArgs
type equals threeArgs

func (h halt) execute(icc *IntCodeComputer) {
}

func (i addition) execute(icc *IntCodeComputer) {
	icc.memory[i.c] = i.a + i.b
	icc.ip += 4
}

func (i multiplication) execute(icc *IntCodeComputer) {
	icc.memory[i.c] = i.a * i.b
	icc.ip += 4
}

func (i adjustRelativeBase) execute(icc *IntCodeComputer) {
	icc.relativeBase += i.a
	icc.ip += 2
}

func (i input) execute(icc *IntCodeComputer) {
	icc.memory[i.a] = icc.inputter()
	icc.ip += 2
}

func (i output) execute(icc *IntCodeComputer) {
	icc.outputter(icc.memory[i.a])
	icc.ip += 2
}

func (i jmpIfTrue) execute(icc *IntCodeComputer) {
	if i.a != 0 {
		icc.ip = i.b
	} else {
		icc.ip += 3
	}
}

func (i jmpIfFalse) execute(icc *IntCodeComputer) {
	if i.a == 0 {
		icc.ip = i.b
	} else {
		icc.ip += 3
	}
}

func (i lessThan) execute(icc *IntCodeComputer) {
	if i.a < i.b {
		icc.memory[i.c] = 1
	} else {
		icc.memory[i.c] = 0
	}
	icc.ip += 4
}

func (i equals) execute(icc *IntCodeComputer) {
	if i.a == i.b {
		icc.memory[i.c] = 1
	} else {
		icc.memory[i.c] = 0
	}
	icc.ip += 4
}

func nextOp(icc *IntCodeComputer) op {
	opcode := icc.memory[icc.ip]

	isImmediateArg := func(arg int) bool {
		if arg == 1 {
			return (opcode%1000)/100 == 1
		}
		if arg == 2 {
			return (opcode%10000)/1000 == 1
		}
		if arg == 3 {
			return (opcode%100000)/10000 == 1
		}
		panic("Unknown pos")
	}

	isRelativeArg := func(arg int) bool {
		if arg == 1 {
			return (opcode%1000)/100 == 2
		}
		if arg == 2 {
			return (opcode%10000)/1000 == 2
		}
		if arg == 3 {
			return (opcode%100000)/10000 == 2
		}
		panic("Unknown pos")
	}

	getInputArgValue := func(arg int) int {
		argAddress := icc.ip + arg

		v := icc.memory[argAddress]
		if isImmediateArg(arg) {
			return v
		}

		if isRelativeArg(arg) {
			return icc.memory[icc.relativeBase+v]
		}

		return icc.memory[v]
	}

	getOutputAddress := func(arg int) int {
		argAddress := icc.ip + arg
		if isImmediateArg(arg) {
			return argAddress
		}

		v := icc.memory[argAddress]
		if isRelativeArg(arg) {
			return icc.relativeBase + v
		}

		return v
	}

	instruction := opcode % 100
	switch instruction {
	case 1:
		return addition{
			a: getInputArgValue(1),
			b: getInputArgValue(2),
			c: getOutputAddress(3),
		}
	case 2:
		return multiplication{
			a: getInputArgValue(1),
			b: getInputArgValue(2),
			c: getOutputAddress(3),
		}
	case 3:
		return input{a: getOutputAddress(1)}
	case 4:
		return output{a: getOutputAddress(1)}
	case 5:
		return jmpIfTrue{
			a: getInputArgValue(1),
			b: getInputArgValue(2),
		}
	case 6:
		return jmpIfFalse{
			a: getInputArgValue(1),
			b: getInputArgValue(2),
		}
	case 7:
		return lessThan{
			a: getInputArgValue(1),
			b: getInputArgValue(2),
			c: getOutputAddress(3),
		}
	case 8:
		return equals{
			a: getInputArgValue(1),
			b: getInputArgValue(2),
			c: getOutputAddress(3),
		}
	case 9:
		return adjustRelativeBase{a: getInputArgValue(1)}
	case 99:
		return halt{}
	default:
		panic("Unknown op")
	}
}

func (icc *IntCodeComputer) ExecuteUntilHalt() {
	for {
		op := nextOp(icc)
		switch op.(type) {
		case halt:
			return
		default:
			op.execute(icc)
		}
	}
}

func (icc *IntCodeComputer) GetMemory0() int {
	return icc.memory[0]
}

func (icc *IntCodeComputer) SetInput(n []int) {
	ptr := 0
	icc.inputter = func() int {
		ptr += 1
		return n[ptr-1]
	}
}

func (icc *IntCodeComputer) SetInputter(n func() int) {
	icc.inputter = n
}

func (icc *IntCodeComputer) SetOutputter(n func(int)) {
	icc.outputter = n
}
