---
title: "Open Color - open-source color scheme"
notion_id: 2d79c9f0-d47c-40a9-8429-ff493abd62b0
notion_url: https://app.notion.com/p/Open-Color-open-source-color-scheme-2d79c9f0d47c40a98429ff493abd62b0
last_edited: 2023-12-19T14:04:00.000Z
source_url: https://yeun.github.io/open-color/
tags: ["English", "Frontend", "Graphic Design", "Framework/Library", "Website", "Tool"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# **Open color is an open-source color scheme**

**Optimized for UI like font, background, border, etc.**



# **Goals**

- All colors shall have adequate use
- Provide general color for UI design
- All colors will be beautiful in itself and harmonious
- At the same brightness level, the perceived brightness will be constant

# **Installation**

`$ npm install open-color`

# **Variable convention**

- Sass, SCSS:

`$oc-(color)-(number)`

- Less:

`@oc-(color)-(number)`

- Stylus:

`oc-(color)-(number)`

- CSS:
- `-oc-(color)-(number)`
- **oc** Abbreviation for Open color
- **(color)** Color name like gray, red, lime ect.
- **(number)** 0 to 9. 0 to 9. Brightness spectrum.

# **How to use**

Import the file to your project and use the variables.**Example for Sass, SCSS**

```scss
  @import 'path/open-color';

  .body {
    background-color: $oc-gray-0;
    color: $oc-gray-7;
  }

  a {
    color: $oc-teal-7;

    &:hover,
    &:focus,
    &:active {
      color: $oc-indigo-7;
    }
  }
```

**Example for Tailwind CSS**

```javascript
// tailwind.config.js      
module.exports = {
  presets: [require("./open-color.js")],
  purge: [],
  mode: "jit",
  darkMode: false,
  theme: {
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
```

**Example for Less**

```less
  @import 'path/open-color';

  .body {
    background-color: @oc-gray-0;
    color: @oc-gray-7;
  }

  a {
    color: @oc-teal-7;

    &:hover,
    &:focus,
    &:active {
      color: @oc-indigo-7;
    }
  }
```

**Example for Stylus**

```css
  @import 'path/open-color.styl'

  .body
    background-color: oc-gray-0
    color: oc-gray-7

  a
    color: oc-teal-7;

    &:hover
    &:focus
    &:active
      color: oc-indigo-7
```

**Example for CSS**

```css
  @import 'path/open-color.css';

  .body {
    background-color: var(--oc-gray-0);
    color: var(--oc-gray-7);
  }

  a {
    color: var(--oc-teal-7);
  }

  a:hover,
  a:focus,
  a:active {
    color: var(--oc-indigo-7);
  }
```




