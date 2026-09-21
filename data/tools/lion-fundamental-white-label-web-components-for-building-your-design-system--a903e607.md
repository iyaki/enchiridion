---
title: "Lion - Fundamental white label web components for building your design system"
notion_id: a903e607-1e76-4eb0-95d3-47f16c44befe
notion_url: https://app.notion.com/p/Lion-Fundamental-white-label-web-components-for-building-your-design-system-a903e6071e764eb095d347f16c44befe
last_edited: 2023-11-28T19:06:00.000Z
source_url: https://lion-web.netlify.app/
tags: ["English", "Frontend", "HTML", "Framework/Library"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

An accordion is a vertically stacked set of interactive headings that each contain a title representing a section of content. It allows users to toggle the display of sections of content.

## When to use

- Accordions are commonly used to reduce the need to scroll when presenting multiple sections of content on a single page
- Longer pages can benefit users. Accordions shorten pages and reduce scrolling, but they increase the interaction cost by requiring people to decide on topic headings.
- Accordions conserve space on mobile but they can also cause disorientation and too much scrolling.
- Accordions should be avoided when your audience needs most or all of the content on the page to answer their question. Better to show all page content at once when the use case supports it.
- Accordions are more suitable when people need only a few key pieces of content on a single page. By hiding most of the content, users can spend their time more efficiently focused on the few topics that matter.

## Features

- Content gets provided by users (slotted in)
- Handles accessibility
- Support navigation via keyboard

## How to use

### Code

1. Install

```shell
npm i --save @lion/ui
```

1. Use scoped registry

```javascript
import { html, LitElement } from 'lit';
import { ScopedElementsMixin } from '@open-wc/scoped-elements';
import { LionAccordion } from '@lion/ui/accordion.js';

class MyComponent extends ScopedElementsMixin(LitElement) {
  static get scopedElements() {
    return { 'lion-accordion': LionAccordion };
  }
  render() {
    return html`
      <lion-accordion>
        <h3 slot="invoker">
          <button>Nutritional value</button>
        </h3>
        <p slot="content">
          Orange flesh is 87% water, 12% carbohydrates, 1% protein, and contains negligible fat
          (table). In a 100 gram reference amount, orange flesh provides 47 calories, and is a rich
          source of vitamin C, providing 64% of the Daily Value. No other micronutrients are present
          in significant amounts (table).
        </p>
      </lion-accordion>
    `;
  }
}

```

1. Use html

```html
<script type="module">
  import '@lion/ui/define/lion-accordion.js';
</script>

<lion-accordion>
  <h3 slot="invoker">
    <button>Nutritional value</button>
  </h3>
  <p slot="content">
    Orange flesh is 87% water, 12% carbohydrates, 1% protein, and contains negligible fat (table).
    In a 100 gram reference amount, orange flesh provides 47 calories, and is a rich source of
    vitamin C, providing 64% of the Daily Value. No other micronutrients are present in significant
    amounts (table).
  </p>
</lion-accordion>

```
