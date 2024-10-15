# Custom text editor

Using [Raylib](https://www.raylib.com/)

> [!WARNING]
> Unicode is not supported for now. I have an issue with raylib I need to settle. But won't affect the rest of the developpement. 

## TODO
- Treesitter synthax highlighting
- Add theme files
- Open, read / write files
- Change fonts with env vars for font folders
- Config file -> struct

LoadUTF8 -> fix for unicode support. I am just stupid my bad

## Treesitter steps

- Reading tree correctly and printing with basic colors
- Detect language with file extension and switch case for tree init (list available langs)
- Print with colors with raylib funcs
    - This requires care for optimization. Buffer and print all at once ??