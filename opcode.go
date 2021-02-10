package main

func getNextOpcode() word {
	res := word(gameMemory[programCounter])
	res <<= 8
	res |= word(gameMemory[programCounter+1])

	programCounter += 2

	return res
}

func opcode00E0(opcode word) {
	// clear screen
}

func opcode00EE() {
	programCounter = gameStack.pop()
}

func opcode1NNN(opcode word) {
	programCounter = opcode & 0x0fff
}

func opcode2NNN(opcode word) {
	gameStack.push(programCounter)
	programCounter = opcode & 0x0fff
}

func opcode3XNN(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	nn := opcode & 0x00ff

	if registers[regx] == byte(nn) {
		programCounter += 2
	}
}

func opcode4XNN(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	nn := opcode & 0x00ff

	if registers[regx] != byte(nn) {
		programCounter += 2
	}
}

func opcode5XY0(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	if registers[regx] == registers[regy] {
		programCounter += 2
	}
}
