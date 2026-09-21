---
title: "How to make your images in Markdown on GitHub adjust for dark mode and light mode"
notion_id: 1d954f1c-7d23-8117-83c3-c173a5ac5a9d
notion_url: https://app.notion.com/p/How-to-make-your-images-in-Markdown-on-GitHub-adjust-for-dark-mode-and-light-mode-1d954f1c7d23811783c3c173a5ac5a9d
last_edited: 2025-08-06T21:47:00.000Z
source_url: https://github.blog/developer-skills/github/how-to-make-your-images-in-markdown-on-github-adjust-for-dark-mode-and-light-mode/
tags: ["Github Blog", "English", "Producer (Individual Contributor)", "Guide"]
---
![image](https://github.blog/wp-content/uploads/2025/04/darklight.png?w=1600)

GitHub supports dark mode and light mode, and as developers, we can make our README images look great in both themes. Here’s a quick guide to using the `<picture>` element in your GitHub Markdown files to dynamically switch images based on the user’s color scheme.

When [developers switch to GitHub’s dark mode (or vice versa)](https://docs.github.com/en/get-started/accessibility/managing-your-theme-settings), standard images can look out of place, with bright backgrounds or clashing colors.

Instead of forcing a one-size-fits-all image, you can tailor your visuals to blend seamlessly with the theme. It’s a small change, but it can make your project look much more polished.

## **One snippet, two themes!**

Here’s the magic snippet you can copy into your README (or any Markdown file):

```html
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="dark-mode-image.png">
  <source media="(prefers-color-scheme: light)" srcset="light-mode-image.png">
  <img alt="Fallback image description" src="default-image.png">
</picture>
```

Now, we say it’s magic, but let’s take a peek behind the curtain to show how it works:

- The `<picture>` tag lets you define multiple image sources for different scenarios.
- The `<source media="...">` attribute matches the user’s color scheme. 
- When `media="(prefers-color-scheme: dark)"`, the browser loads the `srcset` image when GitHub is in dark mode.
- Similarly, when `media="(prefers-color-scheme: light)"`, the browser loads the `srcset` image when GitHub is in light mode.
- If the browser doesn’t support the `<picture>` element, or the user’s system doesn’t match any defined media queries, the fallback `<img>` tag will be used.

You can use this approach in your repo README files, documentation hosted on GitHub, and any other Markdown files rendered on GitHub.com!

## **Demo**

What’s better than a demo to help you get started? Here’s what this looks like in practice:

[https://github.blog/wp-content/uploads/2025/04/Toggle-Dark-and-Light-Mode-on-GitHub-.mp4#t=0.001?_=1](https://github.blog/wp-content/uploads/2025/04/Toggle-Dark-and-Light-Mode-on-GitHub-.mp4#t=0.001?_=1)
