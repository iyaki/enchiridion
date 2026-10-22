---
title: "11ty (nodejs)"
notion_id: 67d9ed7c-5ae1-4208-867b-af85ccb333b1
notion_url: https://app.notion.com/p/11ty-nodejs-67d9ed7c5ae14208867baf85ccb333b1
last_edited: 2026-09-21T17:14:00.000Z
source_url: https://www.11ty.dev/
tags: ["English", "Blogging/Content Creation", "Untried", "Frontend", "Tool"]
---


## Quick Start

Eleventy requires a way to run JavaScript on your computer and we recommend Node.js (version 18 or newer). You can check whether or not you have Node.js installed by running node --version in a Terminal. (Well, wait—what is a Terminal?) If node is not found or it reports a version number below 18, you will need to install Node.js before moving on.

Now we’ll create an index.md Markdown file. You can do this in the text editor of your choice or by running one of these commands in your terminal:

```plain text
echo '# Heading' > index.md
```

```plain text
echo '# Heading' | out-file -encoding utf8 'index.md'
```

If the out-file command is not available in your Windows Terminal window (it’s PowerShell specific), use the Cross Platform method.

```plain text
echo '# Heading' | npx @11ty/create index.md
```

Learn more about @11ty/create (requires Node.js 18 or newer).

Run Eleventy using npx, an npm-provided command that is bundled with Node.js.

```plain text
npx @11ty/eleventy --serve
```

```plain text
pnpm dlx @11ty/eleventy --serve
```

pnpm is an optional alternative to npm that needs to be installed separately.

```plain text
yarn dlx @11ty/eleventy --serve
```

Yarn is an optional alternative to npm that needs to be installed separately.

Eleventy compiles any files in the current directory matching valid file extensions (md is one of many) to the _site output folder. It might look like this:

```plain text
[11ty] Writing _site/index.html from ./index.md (liquid) [11ty] Wrote 1 file in 0.03 seconds (v3.1.6) [11ty] Watching… [11ty] Server at http://localhost:8080/
```

The --serve option also starts a local development server. Open up http://localhost:8080/ in your favorite web browser to view your web site.

If you’d like to experiment further with different template syntax, edit the following sample index.md file in your browser. Front Matter, Liquid and Markdown are in use.

```plain text
--- title: Heading --- # {{ title }}
```

- Read our full Get Started guide on the docs.
- Watch 6 minutes to Build a Blog from Scratch.

## Try Eleventy in Your Browser

Next try editing one of the three files in this Eleventy project. Change the title in front matter on a blog post and watch the list update on the index page!

```plain text
--- subject: World --- # Hello {{ subject }} You can type here! - [Markdown](/docs/languages/markdown/) - [Liquid](/docs/languages/liquid/) ## Posts {%- for post in collections.posts %} - [{{ post.data.title }}]({{ post.url}}) {%- endfor %} _Built with {{ eleventy.generator }}_
```

```plain text
--- title: First blog post ⬅️ tags: posts --- # {{ title }}
```

```plain text
--- title: Second blog post tags: posts --- # {{ title }}
```

## News from the Blog

- New Sponsorship Tiers for the Build Awesome Kickstarter (2026 May 13)
- Back Build Awesome Pro and make it easier to build for the web! (2026 April 28)
- Collaborative Editing as Progressive Enhancement (2026 April 21)
- Eleventy is now Build Awesome (2026 March 03)
- Securely Publishing our Packages to npm (2025 December 03)
- The Eleventy Community Survey (2025) (2025 November 19)
- Eleventy v3.1.0 is now available — 11% faster and 22% smaller! (2025 May 13)
- Eleventy v3.0.0 is now available! (2024 October 02)
- 11ty is joining Font Awesome (2024 September 12)
- …and 67 more on the blog archives.

## Why should you use Eleventy?

- Eleventy has fast builds and even faster web sites. Name Building ×4000 Markdown Files Eleventy 1.93s == 🏁 Astro 22.90s ======================= 🏁 Gatsby 29.05s ============================== 🏁 Next.js 70.65s ======================================================================= 🏁
- Eleventy is production ready and trusted by: Eleventy has been downloaded 21,245,498 times and is used on 88,000+ repositories on GitHub.
- Eleventy offers full control over your project’s output. We don’t inject our own markup into your pages.
- Eleventy has a lovely community of folks that really care about what they build. What is Andy Bell saying about Eleventy?
- Eleventy is stable. We’ve shipped 226 releases going back to the first version in December 2017 and only three of those releases have had Eleventy-specific changes requiring developer changes. a11yproject.com launched with version 1.0.0 of Eleventy. […] It's been a little under three years and I haven't had to make any adjustments to its dependencies, and it can still install and run from a cold start with no complications. When I update the site to use version 2.0.0 I'll actually be removing dependencies, and not adding more. […] That's rare and special.— Eric Bailey
- Eleventy doesn’t track you. We don’t have or use telemetry nor require you to opt-out of data collection.
- Eleventy is zero-config to start and can be extended with flexible configuration options.
- Eleventy works with multiple template languages. You can pick one or use them all together in a single project: HTML *.html Markdown *.md WebC *.webc JavaScript *.11ty.js Liquid *.liquid Nunjucks *.njk Handlebars *.hbs Mustache *.mustache EJS *.ejs Haml *.haml Pug *.pug TypeScript *.ts JSX *.jsx MDX *.mdx Sass *.scss Custom *.*
- Eleventy uses independent template languages. We don’t want to hold your content hostage with a custom format. If you decide to use a different syntax later, having your content decoupled in this way will make migration easier.
- Eleventy does not require that you use a JavaScript framework—that means zero client-side JavaScript by default across the board. We’re thinking long-term to opt-out of the framework rat race. The tool chain, modules, and components you use in your front end stack are decoupled from this tool. Work from a solid foundation of pre-rendered templates that suit your project’s progressive enhancement baseline requirements.
- Eleventy works with your project’s existing directory structure. The tool doesn’t require an app directory or a pages directory. Use the structure that you want.
- Eleventy allows incremental adoption. We only look for the files and directories you specify. Further, with even more precision you can opt-out or ignore specific files in your project. You don’t need to start an Eleventy project from scratch. Eleventy is flexible enough to allow conversion of only a few templates at a time. Migrate as fast or as slow as you’d like.

Get started with Eleventy today! or read more about Eleventy’s project goals.

Documentation

Todd and Bruce said this button should be bigger and as you can see they were right.

## Built With Eleventy

A random sample taken from 881 authors. Check out the fastest of their 1269 web sites.

## Don’t take my word for it Rainbow

Listen to what these happy developers are saying about Eleventy:

> “Eleventy is absolutely wonderful. It’s by far the nicest static site generator I’ve used in what feels like forever.” —Addy Osmani

> “I looked into and actively tried using various static site generators for this project. Eleventy was the only one I could find that gave me the fine-grained control I needed at blazingly fast build times.” —Mathias Bynens

> “Eleventy + Netlify have become my new workflow for static sites. I think I'm in love.” —Mina Markham

> “Eleventy is a killer static site generator. That’s all.” —Sara Soueidan

> “2022 winner of the Google Open Source Peer Bonus Award” —Google

> “Don’t tell Zach I said it but Eleventy is seeming fresh as hell so far” —Mat Marquis

> “I use Eleventy on almost every project at this point and I love it.” —Lea Verou

…and many more!

## Alternatives

This project aims to directly compete with all other static site generators. We encourage you to try out a few others:

- Jekyll (Ruby)
- Hugo (Go)
- Hexo (JavaScript)
- Gatsby (JavaScript using React)
- Nuxt (JavaScript using Vue)
- Next.js (JavaScript using React)
- Bridgetown (Ruby)
- Astro (JavaScript)
- Remix (JavaScript using React)
- SvelteKit (JavaScript using Svelte)
- More at jamstack.org
