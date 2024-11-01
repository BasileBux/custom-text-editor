# Custom text editor

A small personal project to write a text editor from 0 in Go and [Raylib](https://www.raylib.com/). Cross-platform (only tested on Linux tho). This is a work in progress and pretext to learn a bunch of stuff. (The name is temporary until I find a nice one)

The current features are really basic. Move with arrows, write text, delete text, open files, write files, ... If you are wondering what I will implement next, take a look at my [todos](#TODO).

I also implemented syntax highlighting using [Tree-sitter](https://tree-sitter.github.io/tree-sitter/). For, now there is only `c`, but I want to implement other languages I know when I have time. If you want to throw a quick pull request and implement syntax highlighting for a language you like, you are more welcome to do so. (I find it pretty boring actually)

> [!WARNING]
> Unicode is not supported for now. I have an issue with raylib I need to settle. But won't affect the rest of the devlopement. So it will only accept ASCII characters. 

### Supported language syntax highlighting
- c

## Usage

To build it all you need is golang installed on your machine. Clone the repo, execute `go mod tidy` in the directory and build it with go. 

To run it, Just execute the program and provide the path to the file you want to edit. If you don't give a file, it will open a blank file which won't be able to be saved. The text editor cannot create a new file yet. 

## TODO

- Implement scrolling both vertically and horizontally (with padding)
    - Cursor (arrows)
- Config file -> struct
- Optimize syntax highlighting with caching (good luck)
- Ctrl+backspace deletes whole word or whole space
- Delete key normal behaviour + ctrl+del
- Change fonts with env vars for font folders
- Change between indentation with tabs and spaces + modify sizes
- Text Selection
- Copy / Paste
- Mouse support (click, select, scroll)
- Change cursor shape
- Better theme files. Remove useless things

## Issues

- Weird padding on certain machines ? Env vars ? Debug build ?
- LoadUTF8 -> fix for unicode support. I might be stupid, not sure yet

