# Kenzan text editor

A small personal project to write a text editor from 0 in Go and [Raylib](https://www.raylib.com/). Cross-platform (only tested on Linux tho). This is a work in progress and pretext to learn a bunch of stuff.

The current features are really basic. Move with arrows, write text, delete text, open files, write files, ... If you are wondering what I will implement next, take a look at my [todos](#todo).

I also implemented syntax highlighting using [Tree-sitter](https://tree-sitter.github.io/tree-sitter/). For, now there is only `c`, but I want to implement other languages I know when I have time. If you want to throw a quick pull request and implement syntax highlighting for a language you like, you are more than welcome to do so. (I find it pretty boring actually)

> [!WARNING]
> I put the [Geist mono](https://vercel.com/font) nerd font in the project to use as default font temporarily. I do not own this font and have nothing to do with it. I just like how it looks. 

> [!WARNING]
> Unicode is not supported for now. I have solved my issue so it might come in the future but not yet. So if you type non-ascii characters, they will simply be ignored

### Supported language syntax highlighting

- c

## Usage

To build it all you need is golang installed on your machine. Clone the repo, execute `go mod tidy` in the directory and build it with go.

To run it, Just execute the program and provide the path to the file you want to edit. If you don't give a file, it will open a blank file which won't be able to be saved. The text editor cannot create a new file yet.

## Todo

- Change fonts with env vars for font folders (OS dependent good luck)
- Get monitor refresh rate and adapt fps to it
- `Delete key` normal behavior
- Change between indentation with tabs and spaces + modify sizes -> indent guides
- Text Selection
- Copy / Paste
- Mouse support (click, select, scroll)
- Change cursor shape
- ...

## Issues

- Crashing when on last line going right. This only happens when the cursor on lines above was more on the right than the last char on last line -> out of bounds array access -> other cases too

- Weird padding on certain machines ? Env vars ? Debug build ?

> [!WARNING]
> The padding error is really cryptic. However, it seems to be a really weird error tied to the monitor or some shit like that. This should be a known issue but no fix is needed.
