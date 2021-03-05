package main

import (
	"fmt"
	"time"
)

func (c *chip8) getNextOpcode() uint16 {
	res := uint16(c.gameMemory[c.programCounter])
	res <<= 8
	res |= uint16(c.gameMemory[c.programCounter+1])

	c.programCounter += 2

	return res
}

func (c *chip8) executeOpcode(opcode uint16) {
	switch opcode & 0xf000 {
	case 0x0000:
		c.decodeOpcode00E(opcode)
	case 0x1000:
		c.opcode1NNN(opcode)
	case 0x2000:
		c.opcode2NNN(opcode)
	case 0x3000:
		c.opcode3XNN(opcode)
	case 0x4000:
		c.opcode4XNN(opcode)
	case 0x5000:
		c.opcode5XY0(opcode)
	case 0x6000:
		c.opcode6XNN(opcode)
	case 0x7000:
		c.opcode7XNN(opcode)
	case 0x8000:
		c.decodeOpcode8XY(opcode)
	case 0x9000:
		c.opcode9XY0(opcode)
	case 0xA000:
		c.opcodeANNN(opcode)
	case 0xB000:
		c.opcodeBNNN(opcode)
	case 0xC000:
		c.opcodeCXNN(opcode)
	case 0xD000:
		c.opcodeDXYN(opcode)
	case 0xE000:
		c.decodeOpcodeEX(opcode)
	case 0xF000:
		c.decodeOpcodeFX(opcode)
	default:
		fmt.Printf("UNKOWN OPCODE: 0x%x\n", opcode) // panic or fatal
	}
}

func (c *chip8) decodeOpcode00E(opcode uint16) {
	switch opcode & 0x000f {
	case 0xe:
		c.opcode00EE()
	case 0x0:
		c.opcode00E0(opcode)
	default:
		fmt.Printf("UNKOWN OPCODE: 0x%x\n", opcode)
	}
}

func (c *chip8) decodeOpcode8XY(opcode uint16) {
	switch opcode & 0x000f {
	case 0x0:
		c.opcode8XY0(opcode)
	case 0x1:
		c.opcode8XY1(opcode)
	case 0x2:
		c.opcode8XY2(opcode)
	case 0x3:
		c.opcode8XY3(opcode)
	case 0x4:
		c.opcode8XY4(opcode)
	case 0x5:
		c.opcode8XY5(opcode)
	case 0x6:
		c.opcode8XY6(opcode)
	case 0x7:
		c.opcode8XY7(opcode)
	case 0xe:
		c.opcode8XYE(opcode)
	default:
		fmt.Printf("UNKOWN OPCODE: 0x%x\n", opcode)
	}
}

func (c *chip8) decodeOpcodeEX(opcode uint16) {
	switch opcode & 0x00ff {
	case 0x9e:
		c.opcodeEX9E(opcode)
	case 0xa1:
		c.opcodeEXA1(opcode)
	default:
		fmt.Printf("UNKOWN OPCODE: 0x%x\n", opcode)
	}
}

func (c *chip8) decodeOpcodeFX(opcode uint16) {
	switch opcode & 0x00ff {
	case 0x07:
		c.opcodeFX07(opcode)
	case 0x0A:
		c.opcodeFX0A(opcode)
	case 0x15:
		c.opcodeFX15(opcode)
	case 0x18:
		c.opcodeFX18(opcode)
	case 0x1e:
		c.opcodeFX1E(opcode)
	case 0x29:
		c.opcodeFX29(opcode)
	case 0x33:
		c.opcodeFX33(opcode)
	case 0x55:
		c.opcodeFX55(opcode)
	case 0x65:
		c.opcodeFX65(opcode)
	default:
		fmt.Printf("UNKOWN OPCODE: 0x%x\n", opcode)
	}
}

func (c *chip8) opcode00E0(opcode uint16) { // clear display
	for y := range c.screen.pixels {
		for x := range c.screen.pixels[y] {
			c.screen.pixels[y][x] = 0
		}
	}
}

func (c *chip8) opcode00EE() {
	c.programCounter = c.memStack.pop()
}

func (c *chip8) opcode1NNN(opcode uint16) {
	c.programCounter = opcode & 0x0fff
	// programCounter = opcode & 0x0fff - 2
}

func (c *chip8) opcode2NNN(opcode uint16) {
	c.memStack.push(c.programCounter)
	c.programCounter = opcode & 0x0fff
	// programCounter = (opcode & 0x0fff) -2
}

func (c *chip8) opcode3XNN(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	nn := opcode & 0x00ff

	if c.registers[regx] == byte(nn) {
		c.programCounter += 2
	}
}

func (c *chip8) opcode4XNN(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	nn := opcode & 0x00ff

	if c.registers[regx] != byte(nn) {
		c.programCounter += 2
	}
}

func (c *chip8) opcode5XY0(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	if c.registers[regx] == c.registers[regy] {
		c.programCounter += 2
	}
}

func (c *chip8) opcode6XNN(opcode uint16) {
	regx := opcode & 0x0f00 // >> 8
	regx >>= 8

	nn := byte(opcode & 0x00ff)
	c.registers[regx] = nn
}

func (c *chip8) opcode7XNN(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	nn := byte(opcode & 0x00ff)

	c.registers[regx] += nn
}

func (c *chip8) opcode8XY0(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	c.registers[regx] = c.registers[regy]
}

func (c *chip8) opcode8XY1(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	c.registers[regx] |= c.registers[regy]
}

func (c *chip8) opcode8XY2(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	c.registers[regx] &= c.registers[regy]
}

func (c *chip8) opcode8XY3(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	c.registers[regx] ^= c.registers[regy]
}

func (c *chip8) opcode8XY4(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	c.registers[0xf] = byte((uint16(c.registers[regx]) + uint16(c.registers[regy])>>8))
	c.registers[regx] += c.registers[regy]
}

func (c *chip8) opcode8XY5(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	if c.registers[regx] < c.registers[regy] {
		c.registers[0xf] = 0
	} else {
		c.registers[0xf] = 1
	}

	c.registers[regx] = c.registers[regx] - c.registers[regy]
}

func (c *chip8) opcode8XY6(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	c.registers[0xf] = c.registers[regx] & 0x1

	c.registers[regx] >>= 1
}

func (c *chip8) opcode8XY7(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	if c.registers[regy] < c.registers[regx] {
		c.registers[0xf] = 0
	} else {
		c.registers[0xf] = 1
	}

	c.registers[regx] = c.registers[regy] - c.registers[regx]
}

func (c *chip8) opcode8XYE(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	c.registers[0xf] = c.registers[regx] >> 7

	c.registers[regx] <<= 1
}

func (c *chip8) opcode9XY0(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	regy := opcode & 0x00f0
	regy >>= 4

	if c.registers[regx] != c.registers[regy] {
		c.programCounter += 2
	}
}

func (c *chip8) opcodeANNN(opcode uint16) {
	c.addressI = opcode & 0x0fff
}

func (c *chip8) opcodeBNNN(opcode uint16) {
	nnn := opcode & 0x0fff
	c.programCounter = nnn + uint16(c.registers[0])
}

func (c *chip8) opcodeCXNN(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	nn := byte(opcode & 0x00ff)

	c.registers[regx] = randByte() & nn
}

// RECHECK
func (c *chip8) opcodeDXYN(opcode uint16) {
	c.registers[0xf] = 0

	regx := opcode & 0x0f00
	regx >>= 8
	xCoord := c.registers[regx]

	regy := opcode & 0x00f0
	regy >>= 4
	yCoord := c.registers[regy]

	height := byte(opcode & 0x000f)

	for row := byte(0); row < height; row++ {
		for i := byte(0); i < 8; i++ {
			y := (yCoord + row)
			y %= byte(len(c.screen.pixels)) // handling out of range

			x := (xCoord + i)
			x %= byte(len(c.screen.pixels[0]))

			sprite := c.gameMemory[c.addressI+uint16(row)]
			spriteBit := sprite & (128 >> i)

			// If any 'on' pixels are going to be flipped, then set
			// VF to 1 per the spec
			if spriteBit > 0 {
				if c.screen.pixels[y][x] > 0 {
					c.registers[0xf] = 1
					c.screen.pixels[y][x] = 0
				} else {
					c.screen.pixels[y][x] = 1
				}
			}
		}
	}
}

func (c *chip8) opcodeEX9E(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	key := c.registers[regx]

	if c.keyState[key] {
		c.programCounter += 2
	}

	c.keyState[key] = false
}

func (c *chip8) opcodeEXA1(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	key := c.registers[regx]

	if !c.keyState[key] {
		c.programCounter += 2
	}

	c.keyState[key] = false
}

func (c *chip8) opcodeFX07(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	c.registers[regx] = c.delayTimer
}

func (c *chip8) opcodeFX0A(opcode uint16) {
	c.running = false

	regx := opcode & 0x0f00
	regx >>= 8

	for {
		keyInd := c.pressedKey()

		if keyInd != -1 {
			c.registers[regx] = byte(keyInd)
			c.running = true
			break
		}

		time.Sleep(5 * time.Millisecond)
	}
}

func (c *chip8) pressedKey() int {
	for i, k := range c.keyState {
		if k {
			return i
		}
	}

	return -1
}

func (c *chip8) opcodeFX15(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	c.delayTimer = c.registers[regx]
}

func (c *chip8) opcodeFX18(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	c.soundTimer = c.registers[regx]
}

func (c *chip8) opcodeFX1E(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	c.addressI += uint16(c.registers[regx])
}

func (c *chip8) opcodeFX29(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8
	c.addressI = uint16(c.registers[regx]) * 5
}

func (c *chip8) opcodeFX33(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	hundreds := c.registers[regx] / 100
	tens := (c.registers[regx] / 10) % 10
	ones := c.registers[regx] % 10

	c.gameMemory[c.addressI] = hundreds
	c.gameMemory[c.addressI+1] = tens
	c.gameMemory[c.addressI+2] = ones
}

func (c *chip8) opcodeFX55(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	for i := uint16(0); i <= regx; i++ {
		c.gameMemory[c.addressI+i] = c.registers[i]
	}

	// addressI += regx + 1 // incremnt in loop
}

func (c *chip8) opcodeFX65(opcode uint16) {
	regx := opcode & 0x0f00
	regx >>= 8

	for i := uint16(0); i <= regx; i++ {
		c.registers[i] = c.gameMemory[c.addressI+i]
	}

	// addressI += regx + 1 // incremnt in loop
}
