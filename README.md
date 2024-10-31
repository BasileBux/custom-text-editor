# Custom text editor

A small personal project to write a text editor from 0 in Go and [Raylib](https://www.raylib.com/). This is a work in progress and pretext to learn a bunch of stuff. 

The current features are really basic. Move with arrows, write text, delete text, open files, write files, ... If you are wondering what I will implement next, take a look at my [todos](#TODO).

I also implemented syntax highlighting using [Tree-sitter](https://tree-sitter.github.io/tree-sitter/). For, now there is only `c`, but I want to implement other languages I know when I have time. If you want to throw a quick pull request and implement syntax highlighting for a language you like, you are more welcome to do so. (I find it pretty boring actually)

> [!WARNING]
> Unicode is not supported for now. I have an issue with raylib I need to settle. But won't affect the rest of the devlopement. So it will only accept ASCII characters. 

### Supported language syntax highlighting
- c

## TODO

- Implement scrolling both vertically and horizontally (with padding)
    - Cursor (arrows) then mouse / trackpad
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

