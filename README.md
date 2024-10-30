# Custom text editor

Using [Raylib](https://www.raylib.com/)

> [!WARNING]
> Unicode is not supported for now. I have an issue with raylib I need to settle. But won't affect the rest of the developpement.

## Features

- Finish syntax highlighting for C
- Open, read / write files
	- File selector exactly like the vim one
- Implement scrolling both vertically and horizontally
- Ctrl+backspace deletes whole word or whole space
- Change fonts with env vars for font folders
- Config file -> struct
- Change between indentation with tabs and spaces + modify sizes

## Issues

- Weird alignment issues -> on c example, when adding spaces on second line, it moves the third line ??? Same behavior with base miss alignment on line "Vec2* v;"
- Weird padding on certain machines ? Env vars ? debug build ?
- LoadUTF8 -> fix for unicode support. I might be stupid, not sure yet

### Supported languages
- c
