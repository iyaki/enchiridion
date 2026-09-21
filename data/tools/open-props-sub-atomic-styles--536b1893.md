---
title: "Open Props: sub-atomic styles"
notion_id: 536b1893-4ae0-416c-9cd1-e4a7415b7223
notion_url: https://app.notion.com/p/Open-Props-sub-atomic-styles-536b18934ae0416c9cd1e4a7415b7223
last_edited: 2023-03-04T02:49:00.000Z
source_url: https://open-props.style/
tags: ["English", "CSS", "Untried", "Framework/Library"]
---
- Expertly crafted web design tokens
- Create consistent components
- Useful in any framework

```css
@import "https://unpkg.com/open-props";
```

```css
.card {
  border-radius: var(--radius-2);
  padding: var(--size-fluid-3);
  box-shadow: var(--shadow-2);

  &:hover {
    box-shadow: var(--shadow-3);
  }

  @media (--motionOK) {
    animation: var(--animation-fade-in);
  }
}
```



## Why Use Open Props?

It's **non-prescriptive**.

  Design Consistently  Incidentally harmonious.     Predictable  Thanks to consistent naming conventions.     Incrementally Adoptable  Grab all the props, props as JS, or only what you need.     Customizable  Map the props from JS or customize builds from the command line.
