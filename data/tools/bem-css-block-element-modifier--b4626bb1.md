---
title: "BEM CSS - Block Element Modifier"
notion_id: b4626bb1-bd18-4572-885c-d273db35279b
notion_url: https://app.notion.com/p/BEM-CSS-Block-Element-Modifier-b4626bb1bd184572885cd273db35279b
last_edited: 2023-09-04T18:32:00.000Z
source_url: https://getbem.com/
tags: ["English", "CSS", "Website", "Guide"]
---
[ Learn why it's better](https://getbem.com/introduction)

```css
#opinions_box h1 {
margin: 0 0 8px 0;
text-align: center;
}

#opinions_box {
p.more_pp {
a {
text-decoration: underline;
        }
    }

input[type="text"] {
border: 1px solid #ccc!important;
    }
}
```

```css
.opinions_box {
margin: 0 0 8px 0;
text-align: center;

&__view-more {
text-decoration: underline;
    }

&__text-input {
border: 1px solid #ccc;
    }

&--is-inactive {
color: gray;
    }
}
```
