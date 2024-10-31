# Custom text editor

Using [Raylib](https://www.raylib.com/)

> [!WARNING]
> Unicode is not supported for now. I have an issue with raylib I need to settle. But won't affect the rest of the developpement.

### Supported languages
- c

## TODO

- Open, read / write files
- Implement scrolling both vertically and horizontally (with padding)
- Ctrl+backspace deletes whole word or whole space
- Delete key normal behaviour + ctrl+del
- Change fonts with env vars for font folders
- Config file -> struct
- Change between indentation with tabs and spaces + modify sizes
- Copy / Paste
- Mouse support
- Change cursor shape
- Better theme files. Remove useless things

## Issues

- Weird padding on certain machines ? Env vars ? debug build ?
- LoadUTF8 -> fix for unicode support. I might be stupid, not sure yet

