# CHIP-8 Interpreter
CHIP-8 is an interpreted minimalist programming language developed in the 70’s for microcomputers [COSMAC VIP](https://en.wikipedia.org/wiki/COSMAC_VIP) и [Telmac 1800](https://en.wikipedia.org/wiki/Telmac_1800). 
It enjoyed relative success during the late 1970s and early 1980s as a popular language for the development of simple video games. Games written in the CHIP-8 language could easily run on systems that had a CHIP-8 interpreter, often mislabeled as "emulator".

*Screenshots*

## How it works
### Instructions
Basically interpreter reads intructions loaded from specified ROM and executes them. Each CHIP-8 instruction is two bytes in length and is represented using four hexadecimal digits.

For example:
- 00E0 - instruction for clearing the screen of all graphics data.
- 1NNN - Jump to address NNN.
- FX0A - Wait for a keypress and store the result in register VX.
- etc...

### Graphics
The graphics of games are flicker because it's the way CHIP-8 designed, but my project includes the ability to reduce flickering by simulating the behaviour of old phosphor displays. 
The pixels of phosphor displays glow for several milliseconds after being turned off which increases the appearance of ghosting.

*Gifs*

## Getting Started
### Requirements
Below is some commands that can be used to install the required [SDL2](http://libsdl.org/download-2.0.php) package.

On __Ubuntu 14.04 and above__:\
`apt install libsdl2-dev`

On __macOS__:\
`brew install sdl2 pkg-config`

### Usage
  Before start the interpreter put your preferenced values in '.env' file. 
  ```
    WINDOW_WIDTH=640
    WINDOW_HEIGHT=320

    GAME_ROM_PATH=roms/PONG2

    OPCODES_PER_SECOND=300
  ```
  To run the programm use Makefile:
  ```
  make run 
  ```

### Controls
Original CHIP-8 keyboard layout is mapped to PC keyboard as follows:
```
|1|2|3|C| => |1|2|3|4|
|4|5|6|D| => |Q|W|E|R|
|7|8|9|E| => |A|S|D|F|
|A|0|B|F| => |Z|X|C|V|
```
