package day23

import (
	"time"

	intcodecomputer "github.com/sogard-dev/advent-of-code-2019/intcode"
	"github.com/sogard-dev/advent-of-code-2019/utils"
)

func part1(input string) int {
	aggOutput := make(chan []int, 2^16)

	inputChannels := make([]chan int, 50)
	for id := range 50 {
		iccOutput := make(chan int, 2^16)
		iccInput := make(chan int, 2^16)
		inputChannels[id] = iccInput
		iccInput <- id

		go func(c chan int) {
			messages := []int{id}
			for msg := range c {
				messages = append(messages, msg)
				if len(messages) == 4 {
					aggOutput <- messages
					messages = []int{id}
				}
			}
		}(iccOutput)

		go func() {
			var program string = input
			startComputer(program, iccOutput, iccInput)
		}()
	}

	for msg := range aggOutput {
		src := msg[0]
		dst := msg[1]
		println("Msg from", src, "to", dst, "with data", msg[2], msg[3])
		if dst == 255 {
			return msg[3]
		}
		inputChannels[dst] <- msg[2]
		inputChannels[dst] <- msg[3]
	}

	return 0
}

func part2(input string) int {
	aggOutput := make(chan []int, 2^16)

	inputChannels := make([]chan int, 50)
	iccs := make([]func() int, 50)
	for id := range 50 {
		iccOutput := make(chan int, 2^16)
		iccInput := make(chan int, 2^16)
		inputChannels[id] = iccInput
		iccInput <- id

		go func(c chan int) {
			messages := []int{id}
			for msg := range c {
				messages = append(messages, msg)
				if len(messages) == 4 {
					aggOutput <- messages
					messages = []int{id}
				}
			}
		}(iccOutput)

		go func() {
			icc, f := startComputer(input, iccOutput, iccInput)
			iccs[id] = f
			icc.ExecuteUntilHalt()
		}()
	}

	nat := []int{0, 0}

	go func() {
		for range time.Tick(time.Second * 1) {
			idle := true
			for _, icc := range iccs {
				if icc == nil || icc() < 10 {
					idle = false
				}
			}

			if idle {
				inputChannels[0] <- nat[0]
				inputChannels[0] <- nat[1]
				println("Sending ", nat[0], nat[1])
			}
		}
	}()

	for msg := range aggOutput {
		// src := msg[0]
		dst := msg[1]
		//println("Msg from", src, "to", dst, "with data", msg[2], msg[3])
		if dst == 255 {
			nat = msg[2:4]
		} else {
			inputChannels[dst] <- msg[2]
			inputChannels[dst] <- msg[3]
		}
	}

	return 0
}

func startComputer(program string, output chan int, input chan int) (intcodecomputer.IntCodeComputer, func() int) {
	icc := intcodecomputer.New(utils.GetAllNumbers(program))

	idleCounter := 0
	icc.SetInputter(func() int {
		select {
		case x := <-input:
			idleCounter = 0
			return x
		default:
			idleCounter += 1
			return -1
		}
	})
	icc.SetOutputter(func(n int) {
		output <- n
	})
	return icc, func() int { return idleCounter }
}
