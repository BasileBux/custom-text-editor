# Settings docs

The settings are in json as it is really common amongst other editors and I don't mind it.
You can access the defaults in ~/.config/kenzan/default.json to look up the global structure.

This file structure is copied from the [zed docs](https://zed.dev/docs/configuring-zed)

## Ui elements

### Padding

- Description: Space between text and edge of the window.
- Setting: `padding_[sub-setting]`
  - Sub-settings: `top`, `right`, `bottom`, `left`
- Default: 13

#### Options

`integer` values

```toml
"padding": {
      "top": 13,
      "right": 13,
      ...
    },
...
```

### Font Family

- Description: Font family to use in the editor
- Setting: `font_family`
- Default: "GeistMonoNerdFont-Regular"

#### Options

`string` values

### Font size

- Description: Font size to use in the editor
- Setting: `font_size`
- Default: 30

#### Options

`integers` values

### Font spacing

- Description: Space between characters
- Setting: `font_spacing`
- Default: 1

#### Options

`integer` values

### Scroll padding

- Description: Number of characters and lines to keep visible around the cursor when scrolling, maintaining a buffer in all directions.
- Setting: `scroll_padding`
- Default: 5

#### Options

positive `integer` values (can be 0)

### Cursor ratio

- Description: Ratio of cursor height to text height.
- Setting: `cursor_ratio`
- Default: 1

#### Options

`integers` values between 0 and 1 with 0 not included = ]0;1]

### Theme

- Description: Editor color theme name
- Setting: `theme`
- Default: "Tokyo-night-storm"

#### Options

`string` values. Theme name which corresponds to the name of the theme file in ~/.config/kenzan/themes/
