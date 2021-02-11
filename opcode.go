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

func opcode6XNN(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	nn := byte(opcode & 0x00ff)
	registers[regx] = nn
}

func opcode7XNN(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	nn := byte(opcode & 0x00ff)

	registers[regx] += nn
}

func opcode8XY0(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	registers[regx] = registers[regy]
}

func opcode8XY1(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	registers[regx] |= registers[regy]
}

func opcode8XY2(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	registers[regx] &= registers[regy]
}

func opcode8XY3(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	registers[regx] ^= registers[regy]
}

func opcode8XY4(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	sum := registers[regx] + registers[regy]
	registers[regx] = sum
	if sum > 255 {
		registers[0xf] = 1
	}
}

func opcode8XY5(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	if registers[regx] < registers[regy] {
		registers[0xf] = 0
	}

	registers[regx] = registers[regx] - registers[regy]
}

func opcode8XY6(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	registers[0xf] = registers[regx] & 0x1

	registers[regx] >>= 1
}

func opcode8XY7(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	if registers[regy] < registers[regx] {
		registers[0xf] = 0
	} else {
		registers[0xf] = 1
	}

	registers[regx] = registers[regy] - registers[regx]
}

func opcode8XYE(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	registers[0xf] = registers[regx] >> 7

	registers[regx] <<= 1
}

func opcode9XY0(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	if registers[regx] != registers[regy] {
		programCounter += 2
	}
}

func opcodeANNN(opcode word) {
	addressI = opcode & 0x0fff
}

func opcodeBNNN(opcode word) {
	nnn := opcode & 0x0fff
	programCounter = nnn + word(registers[0])
}

func opcodeCXNN(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	nn := byte(opcode & 0x00ff)

	registers[regx] = randByte() & nn
}
