---
title: "Trees, from The Pierre Computer Company - A file tree rendering library"
notion_id: 35e54f1c-7d23-81a4-b017-dd162b7e76b6
notion_url: https://app.notion.com/p/Trees-from-The-Pierre-Computer-Company-A-file-tree-rendering-library-35e54f1c7d2381a4b017dd162b7e76b6
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://trees.software/
tags: ["Tool", "Article", "The Pierre Computer Company", "English", "Web Development", "Frontend", "React", "Open Source", "User Experience", "Accessibility", "Javascript"]
---
[Home](https://trees.software/)[Docs](https://trees.software/docs)[Diffs](https://diffs.com/)[Theme](https://diffs.com/theme)

# A file tree rendering library

`@pierre/trees` is an open source file tree rendering library. It's built for performance and flexibility, is super customizable, and comes packed with features. Made with love by [The Pierre Computer Company](https://pierre.computer/).

Currently v1.0.0-beta.3

## 

## 

Use the [`gitStatus`](https://trees.software/docs#show-git-status-and-row-annotations) option to show status badges for added, modified, deleted, renamed, untracked, and ignored files. Ignored items inherit their styling without rendering an indicator while folders with changed descendants get a dot indicator automatically.

## 

Render your own custom context menu with [`composition.contextMenu`](https://trees.software/docs#rename-drag-and-trigger-item-actions-add-a-context-menu-as-an-optional-command-surface) and the React `renderContextMenu` prop. This demo exposes trigger modes for right-click, trigger button, or both, and menu actions for new files, new folders, rename, and delete. This demo uses Shadcn UI components for the context menu as an example. Your app can use the menus that you already have.

## 

Move files and folders by dragging them onto other folders, flattened folders, or the root with `dragAndDrop: true`. Drop targets open automatically when you hover, and dragging is disabled while search is active. Pass a `canDrag` callback to prevent specific paths from being dragged. Learn more in the [item actions guide](https://trees.software/docs#rename-drag-and-trigger-item-actions-move-items-with-drag-and-drop).

## 

Filter the tree by typing in the search field. Search across file paths and names. Trees includes three [`fileTreeSearchMode`](https://trees.software/docs#shared-concepts-search-mode-semantics) options to control how non-matching items are shown. All three demos below start with search prepopulated to show the different modes.

## 

Trees with tens of thousands of items render instantly with built-in and automatic virtualization. Only visible rows are mounted. The tree below contains **2,956 files** with every folder expanded. Shown with `stickyFolders` enabled.

## 

With built-in keyboard navigation, focus management, and ARIA roles (tree, treeitem) plus aria-level, aria-posinset, and aria-setsize attributes, Trees are immediately accessible to all users. We've designed Trees to align with WCAG 2.1 guidance.

## 

Choose between the shipped `minimal`, `standard`, and `complete` icon tiers. Each tier is cumulative. Override the built-in palette with CSS variables like `--trees-file-icon-color-javascript`, or fall back to a fully custom sprite. See the [`FileTreeIconConfig`](https://trees.software/docs#icons-configuration-shape)[ reference](https://trees.software/docs#icons-configuration-shape) for the full API.

## 

pierre-lightpierre-dark

Love the Pierre themes? [Install our Pierre Theme pack](https://trees.software/theme) with light and dark flavors, or learn how to build your own Shiki themes.

## Style with CSS variables

Modify CSS custom properties via the `style` prop to override UI and theme colors. For example, below are three examples—custom light, dark, and Synthwave '84—that override our default values and the CSS we use to style the tree. See the [styling and theming reference](https://trees.software/docs#styling-and-theming) for more info.

We’re using [OKLCH colors](https://oklch.com/) here—a modern color space that allows for more uniform colors and more consistent palettes.

## 

## With love from The Pierre Computer Company

Collectively, our team brings over 150 years of expertise designing, building, and scaling the world's largest distributed systems at Cloudflare, Coinbase, Discord, GitHub, Reddit, Stripe, X, and others.

© 2026 The Pierre Computer Co.
