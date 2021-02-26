# CHIP-8 Interpreter
[CHIP-8](https://en.wikipedia.org/wiki/CHIP-8) is an interpreted minimalist programming language developed in the 70’s. It enjoyed relative success during the late 1970s and early 1980s as a popular language for the development of simple video games.

![Chip8-games](https://user-images.githubusercontent.com/49096838/109256823-08185b80-7821-11eb-8ae4-041a472f3090.png)

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
  To run the program use Makefile:\
  `make run`

### Controls
Original CHIP-8 keyboard layout is mapped to PC keyboard as follows:
```
|1|2|3|C| => |1|2|3|4|
|4|5|6|D| => |Q|W|E|R|
|7|8|9|E| => |A|S|D|F|
|A|0|B|F| => |Z|X|C|V|
```
## Graphics
The graphics of games are flicker because it's the way CHIP-8 designed, but my project includes the ability to reduce flickering by simulating the behaviour of old phosphor displays. 
The pixels of phosphor displays glow for several milliseconds after being turned off which increases the appearance of ghosting.
