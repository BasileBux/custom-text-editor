# Settings docs

The settings are in json as it is really common amongst other editors and I don't mind it.
You can access the defaults in ~/.config/kenzan/default.json to look up the global structure.

This file structure is copied from the [zed docs](https://zed.dev/docs/configuring-zed)


## Padding

- Description: Space between text and edge of the window.
- Setting: `padding`
- Default: 13, 13, 13, 2

#### Options

`integer` values\
Sub-settings: `top`, `right`, `bottom`, `left`

```toml
"padding": {
    "top": 13,
    "right": 13,
	"bottom": 13,
	"left": 2
},
```

## Font Family

- Description: Font family to use in the editor
- Setting: `font_family`
- Default: "GeistMonoNerdFont-Regular"

#### Options

`string` values

## Font size

- Description: Font size to use in the editor
- Setting: `font_size`
- Default: 30

#### Options

`integers` values

## Font spacing

- Description: Space between characters
- Setting: `font_spacing`
- Default: 1

#### Options

`integer` values

## Scroll padding

- Description: Number of characters and lines to keep visible around the cursor when scrolling, maintaining a buffer in all directions.
- Setting: `scroll_padding`
- Default: 5

#### Options

positive `integer` values (can be 0)

## Cursor ratio

- Description: Ratio of cursor height to text height.
- Setting: `cursor_ratio`
- Default: 1

#### Options

`integers` values between 0 and 1 with 0 not included = ]0;1]

## Theme

- Description: Editor color theme name
- Setting: `theme`
- Default: "ayu-light"

#### Options

`string` values. Theme name which corresponds to the name of the theme file in ~/.config/kenzan/themes/

## Line numbers

- Description: Section which handles line numbers
- Setting: `line_numbers`

### Show

- Description: Show the line numbers or not
- Setting: `show`
- Default: true

#### Options

`boolean`

### Relative

- Description: Set line numbers as relative or absolute
- Setting: `relative`
- Default: false

#### Options

`boolean`

### Padding left

- Description: Space from left window border to line numbers
- Setting: `padding_left`
- Default: 24

#### Options

Positive `integer` values. If the value is too small or too big, it will just look ugly

### Padding right

- Description: Space from line numbers to text
- Setting: `padding_right`
- Default: 10

#### Options

Positive `integer` values. If the value is too small or too big, it will just look ugly

### Line width

- Description: Width of the line separating line numbers from text
- Setting: `line_width`
- Default: 0

### Offset current

- Description: Aligns the current line number to the right to 3 digits
- Settings: `offset_current`
- Default: true

#### Options

`boolean`

#### Options

Positive `integer` values. Set to 0 to remove line

## Line highlight

- Description: Toggle highlight on current line
- Setting: `line_highlight`
- Default: true

#### Options

`boolean`

## High dpi

- Description: Enable high dpi mode
- Setting: `high_dpi`
- Default: true

#### Options

`boolean` values. True is activated. 

## Performance display

- Description: Enable performance display
- Setting: `performance_display`
- Default: false

#### Options

`boolean`
