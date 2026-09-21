---
title: "shoelace - A collection of professionally designed, every day UI components built on Web standards. Works with all framework as well as regular HTML/CSS/JS"
notion_id: 13eb8e85-d796-493e-864b-59ab18eec568
notion_url: https://app.notion.com/p/shoelace-A-collection-of-professionally-designed-every-day-UI-components-built-on-Web-standards--13eb8e85d796493e864b59ab18eec568
last_edited: 2023-10-12T18:26:00.000Z
source_url: https://github.com/shoelace-style/shoelace
tags: ["English", "UI/UX", "Frontend", "Framework/Library", "CheatSheet"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# [Shoelace](https://github.com/shoelace-style/shoelace?utm_source=tldrwebdev#shoelace)

A forward-thinking library of web components.

- Works with all frameworks 🧩
- Works with CDNs 🚛
- Fully customizable with CSS 🎨
- Includes an official dark theme 🌛
- Built with accessibility in mind ♿️
- Open source 😸

Designed in New Hampshire by [Cory LaViska](https://twitter.com/claviska).

Documentation: [shoelace.style](https://shoelace.style/)

Source: [github.com/shoelace-style/shoelace](https://github.com/shoelace-style/shoelace)

Twitter: [@shoelace_style](https://twitter.com/shoelace_style)

## [Shoemakers 🥾](https://github.com/shoelace-style/shoelace?utm_source=tldrwebdev#shoemakers-)

Shoemakers, or "Shoelace developers," can use this documentation to learn how to build Shoelace from source. You will need Node >= 14.17 to build and run the project locally.

**You don't need to do any of this to use Shoelace!** This page is for people who want to contribute to the project, tinker with the source, or create a custom build of Shoelace.

If that's not what you're trying to do, the [documentation website](https://shoelace.style/) is where you want to be.

### [What are you using to build Shoelace?](https://github.com/shoelace-style/shoelace?utm_source=tldrwebdev#what-are-you-using-to-build-shoelace)

Components are built with [LitElement](https://lit-element.polymer-project.org/), a custom elements base class that provides an intuitive API and reactive data binding. The build is a custom script with bundling powered by [esbuild](https://esbuild.github.io/).

### [Forking the Repo](https://github.com/shoelace-style/shoelace?utm_source=tldrwebdev#forking-the-repo)

Start by [forking the repo](https://github.com/shoelace-style/shoelace/fork) on GitHub, then clone it locally and install dependencies.

```plain text
git clone https://github.com/YOUR_GITHUB_USERNAME/shoelace
cd shoelace
npm install
```

### [Developing](https://github.com/shoelace-style/shoelace?utm_source=tldrwebdev#developing)

Once you've cloned the repo, run the following command.

```plain text
npm start
```

This will spin up the dev server. After the initial build, a browser will open automatically. There is currently no hot module reloading (HMR), as browser's don't provide a way to reregister custom elements, but most changes to the source will reload the browser automatically.

### [Building](https://github.com/shoelace-style/shoelace?utm_source=tldrwebdev#building)

To generate a production build, run the following command.

```plain text
npm run build
```

### [Creating New Components](https://github.com/shoelace-style/shoelace?utm_source=tldrwebdev#creating-new-components)

To scaffold a new component, run the following command, replacing `sl-tag-name` with the desired tag name.

```plain text
npm run create sl-tag-name
```

This will generate a source file, a stylesheet, and a docs page for you. When you start the dev server, you'll find the new component in the "Components" section of the sidebar.

### [Contributing](https://github.com/shoelace-style/shoelace?utm_source=tldrwebdev#contributing)

Shoelace is an open source project and contributions are encouraged! If you're interesting in contributing, please review the [contribution guidelines](https://github.com/shoelace-style/shoelace/blob/next/CONTRIBUTING.md) first.

## [License](https://github.com/shoelace-style/shoelace?utm_source=tldrwebdev#license)

Shoelace is designed in New Hampshire by [Cory LaViska](https://twitter.com/claviska). It’s available under the terms of the MIT license.

Designing, developing, and supporting this library requires a lot of time, effort, and skill. I’d like to keep it open source so everyone can use it, but that doesn’t provide me with any income.

**Therefore, if you’re using my software to make a profit,** I respectfully ask that you help [fund its development](https://github.com/sponsors/claviska) by becoming a sponsor. There are multiple tiers to choose from with benefits at every level, including prioritized support, bug fixes, feature requests, and advertising.

👇 Your support is very much appreciated! 👇

- [Become a sponsor](https://github.com/sponsors/claviska)
- [Star on GitHub](https://github.com/shoelace-style/shoelace/stargazers)
- [Follow on Twitter](https://twitter.com/shoelace_style)

Whether you're building Shoelace or building something _with_ Shoelace — have fun creating! 🥾
