package main

import "fmt"

func getNextOpcode() word {
	res := word(gameMemory[programCounter])
	res <<= 8
	res |= word(gameMemory[programCounter+1])

	programCounter += 2

	return res
}

func executeOpcode(opcode word) {
	switch opcode & 0xf000 {
	case 0x0:
		decodeOpcode00E(opcode)
	case 0x1:
		opcode1NNN(opcode)
	case 0x2:
		opcode2NNN(opcode)
	case 0x3:
		opcode3XNN(opcode)
	case 0x4:
		opcode4XNN(opcode)
	case 0x5:
		opcode5XY0(opcode)
	case 0x6:
		opcode6XNN(opcode)
	case 0x7:
		opcode7XNN(opcode)
	case 0x8:
		decodeOpcode8XY(opcode)
	case 0x9:
		opcode9XY0(opcode)
	case 0xA:
		opcodeANNN(opcode)
	case 0xB:
		opcodeBNNN(opcode)
	case 0xC:
		opcodeCXNN(opcode)
	case 0xD:
		opcodeDXYN(opcode)
	case 0xE:
		decodeOpcodeEX(opcode)
	case 0x00e0:
		opcodeEXA1(opcode)
	case 0x00e0:
		opcodeFX07(opcode)
	case 0x00e0:
		opcodeFX0A(opcode)
	case 0x00e0:
		opcodeFX15(opcode)
	case 0x00e0:
		opcodeFX18(opcode)
	case 0x00e0:
		opcodeFX1E(opcode)
	case 0x00e0:
		opcodeFX29(opcode)
	case 0x00e0:
		opcodeFX33(opcode)
	case 0x00e0:
		opcodeFX55(opcode)
	case 0x00e0:
		opcodeFX65(opcode)
	default:
		fmt.Printf("UNKOWN OPCODE: 0x%x\n", opcode)
	}
}

func decodeOpcode00E(opcode word) {
	switch opcode & 0x000f {
	case 0xe:
		opcode00EE()
	case 0x0:
		opcode00E0(opcode)
	default:
		fmt.Printf("UNKOWN OPCODE: 0x%x\n", opcode)
	}
}

func decodeOpcode8XY(opcode word) {
	switch opcode & 0x000f {
	case 0x0:
		opcode8XY0(opcode)
	case 0x1:
		opcode8XY1(opcode)
	case 0x2:
		opcode8XY2(opcode)
	case 0x3:
		opcode8XY3(opcode)
	case 0x4:
		opcode8XY4(opcode)
	case 0x5:
		opcode8XY5(opcode)
	case 0x6:
		opcode8XY6(opcode)
	case 0x7:
		opcode8XY7(opcode)
	case 0xe:
		opcode8XYE(opcode)
	default:
		fmt.Printf("UNKOWN OPCODE: 0x%x\n", opcode)
	}
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

// RECHECK
func opcodeDXYN(opcode word) {
	const scale byte = 10 // to struct ?

	regx := opcode & 0x0f00
	regx >>= 8
	xCoord := registers[regx] * scale

	regy := opcode & 0x00f0
	regy >>= 4
	yCoord := registers[regy] * scale

	height := opcode & 0x000f

	for yline := word(0); yline < height; yline++ {
		sprite := gameMemory[addressI+yline]

		for xpixel, xpixelinv := 0, 7; xpixel < 8; xpixel, xpixelinv = xpixel+1, xpixelinv+1 {
			var mask byte = 1 << xpixelinv
			if sprite&mask == 1 {
				x := xCoord + byte(xpixel)*scale
				y := yCoord + byte(yline)*scale

				var color byte
				if screenData[y][x][0] == 0 { // !=
					color = 255
					registers[0xf] = 1
				} else {
					color = 0 // remove ?
					registers[0xf] = 0
				}

				for i := byte(0); i < scale; i++ {
					for j := byte(0); j < scale; j++ {
						screenData[y+i][x+j][0] = color
						screenData[y+i][x+j][1] = color // really needed ?
						screenData[y+i][x+j][2] = color
					}
				}
			}
		}
	}
}

func opcodeEX9E(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	key := registers[regx]

	if keyState[key] == 1 {
		programCounter += 2
	}
}

func opcodeEXA1(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	key := registers[regx]

	if keyState[key] == 0 {
		programCounter += 2
	}
}

func opcodeFX07(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	registers[regx] = delayTimer
}

func opcodeFX0A(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	keyInd := pressedKey()

	if keyInd == -1 {
		programCounter -= 2
	} else {
		registers[regx] = byte(keyInd)
	}
}

func opcodeFX15(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	delayTimer = registers[regx]
}

func opcodeFX18(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	soundTimer = registers[regx]
}

func opcodeFX1E(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	addressI += word(registers[regx])
}

func opcodeFX29(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8
	addressI = word(registers[regx]) * 5
}

func opcodeFX33(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	hundreds := registers[regx] / 100
	tens := registers[regx] % 100 / 10
	ones := registers[regx] % 10

	gameMemory[addressI] = hundreds
	gameMemory[addressI+1] = tens
	gameMemory[addressI+2] = ones
}

func opcodeFX55(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	for i := word(0); i <= regx; i++ {
		gameMemory[addressI+i] = registers[i]
	}

	addressI += regx + 1 // incremnt in loop
}

func opcodeFX65(opcode word) {
	regx := opcode & 0x0f00
	regx >>= 8

	for i := word(0); i <= regx; i++ {
		registers[i] = gameMemory[addressI+i]
	}

	addressI += regx + 1 // incremnt in loop
}
