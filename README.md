# Kenzan text editor

A small personal project to write a text editor from 0 in Go and [Raylib](https://www.raylib.com/). Cross-platform (only tested on Linux tho). This is a work in progress and pretext to learn a bunch of stuff. 

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

- Fix scrolling bug. For now, I put arbitrary numbers to add to the offset. This number changes with fontsize. Make it not magic numbers
- Clean scrolling code (currently really ugly and unmaintainable code)
- Implement config file (store it in a struct)
- Optimize syntax highlighting with caching (good luck)
- Optimize positions (text, cursor, ...) with caching (calculations are pretty big atp)
- Ctrl+backspace deletes whole word or whole space
- Delete key normal behavior + ctrl+del
- Change fonts with env vars for font folders (OS dependent good luck)
- Change between indentation with tabs and spaces + modify sizes
- Text Selection
- Copy / Paste
- Mouse support (click, select, scroll)
- Change cursor shape
- Better theme files. Remove useless things
- ...

## Issues

- Weird padding on certain machines ? Env vars ? Debug build ?
> [!WARNING]
> The padding error is really cryptic. However, it seems to be a really weird error tied to the monitor or some shit like that. This should be a known issue but no fix is needed.

- When launching the program and the window not rescaled, there is a scroll when writing in an empty file. This causes the characters to be drawn out of frame. This is fixed after a window resize. The problem is probably initialization of the scrolling or something like that. Shouldn't be too hard to fix. 

- LoadUTF8 -> fix for unicode support. I might be stupid, not sure yet

