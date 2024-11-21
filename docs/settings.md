# Settings docs

## Ui elements

All the following config fields will have to be after the following header

```toml
[ui]
```

### Padding

- Description: Space between text and edge of the window.
- Setting: `padding_[sub-setting]`
  - Sub-settings: `top`, `right`, `bottom`, `left`
- Value type: integer
- Default: 13

```toml
padding_top = 13
padding_right = 13
...
```

### Font Family

- Description: Font family to use in the editor
- Setting: `font_family`
- Value type: string
- Default: "GeistMonoNerdFont-Regular"

### Font size

- Description: Font size to use in the editor
- Setting: `font_size`
- Value type: integer
- Default: 30

### Font spacing

- Description: Space between characters
- Setting: `font_spacing`
- Value type: integer
- Default: 1

### Scroll padding

- Description: Number of characters and lines to keep visible around the cursor when scrolling, maintaining a buffer in all directions.
- Setting: `scroll_padding`
- Value type: positive integer
- Default: 5

### Cursor ratio

- Description: Ratio of cursor height to text height.
- Setting: `cursor_ratio`
- Value type: integers between ]0;1]
- Default: 1

### Theme

- Description: Editor color theme name
- Setting: `theme`
- Value type: string
- Default: "Tokyo-night-storm"
