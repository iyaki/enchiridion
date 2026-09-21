---
title: "A Real “Buildless” Modern Web Development Workflow? Oh Yes!"
notion_id: 5ca4e262-30ff-44bb-9efa-cf2e8da6938f
notion_url: https://app.notion.com/p/A-Real-Buildless-Modern-Web-Development-Workflow-Oh-Yes-5ca4e26230ff44bb9efacf2e8da6938f
last_edited: 2023-03-04T02:42:00.000Z
source_url: https://www.spicyweb.dev/buildless-modern-development-workflows-are-this-close-to-a-reality/
tags: ["English", "Frontend", "Javascript", "CSS", "HTML", "Article", "The Spicy Web - Blog"]
---
It's 👉this close👈 to becoming a reality, and you can get a sneak peek today.

**tl;dr** Show it to me! [Demo](https://buildless-static.onrender.com/) [Repo](https://github.com/jaredcwhite/buildless-static-site)

For the entire history of the web, we’ve _never_ had a modern web development workflow which didn’t absolutely and unconditionally require build tooling—either as a one-time process (what’s become known as SSG) or continuously upon request (what’s become known as SSR).

But the times they are a-changing. The era of buildless is nearly upon us. What does that even mean? First, a quick definition of terms:

Buildless The files you edit _are_ the files delivered to the web browser. There’s no intermediate build step. Modern Development Workflow Reusable units of structure (HTML), styles (CSS), and behavior (JS) are encapsulated and scoped via components.

Essentially we’ve grown accustomed to the latter in all of our projects so matter how small or minimalist…yet even in those simple cases (say, a single landing page, or a small 5-page website, or a weekend side-project), there’s been no way to ditch all the extra trappings of the typical development workflow and _just edit a freakin’ HTML file on the filesystem_.

While I’m not about to advocate for the bad old days of cowboy-coding where we zipped up folders and FTP’d them over to a server somewhere (that’s fun the first time and then it’s a total PITA ever after!), there’s certainly a great deal of appeal to “this website is literally this folder with some files in it.” No commands to run (other than a bare-bones localhost web server to test with). No tools to installs. No configuration to tweak. No dependencies to break. Just _HTML_, _CSS_, and _JavaScript_. **As vanilla as vanilla can get.**

So…what is it exactly which purports to turn the tide now? It’s a constellation of recent and emerging web specs that—when combined—offer a whole new way of packaging up and delivering the nuts and bolts of a website…build free! I’ll run through them one at a time, and then _show you_ what’s possible using these new technologies.

### Custom Elements (aka Web Components)

The ability to add arbitrary tags anywhere in your HTML markup is revolutionizing the industry. Suddenly if you want to add tabs or a sidebar or a dropdown menu, you don’t need a high-level template language and gobs of `<div>` and `<span>` tags everywhere as in the days of old. Just write `<nifty-tabs>` or `<spiffy-sidebar>` or `<sl-dropdown>` (hey, [that’s a real thing!](https://shoelace.style/components/dropdown)) in your HTML content, and then define those tags using the web component spec or source them from a third-party. Mind. Blown. 🤯

### Import Maps

In order to write useful web components, or anything useful really in your JavaScript code these days, you need to be able to import at least a small handful of key libraries. But without a build step, how can you import, say `lit`? Or `dayjs`? And I don’t mean the “old-school” way where everything’s global and hanging off of `window`…I mean the new way: using [ES modules](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Modules).

The answer is…you can! That is, _if_ you use import maps. They create a literal “map” between the bare module specifiers (`lit`, `dayjs`, etc.) and ES module packages hosted either locally or, more likely on simple projects, on CDNs.

### CSS Module Scripts + Constructable Stylesheets

This part might be optional for some people, but I for one _like_ authoring my styles in real honest-to-goodness `.css` files. Problem was there was no easy way to import these CSS files such that they could be used within the “shadow DOM” of web components. So instead of your `<a>` styling remaining contained within the single component you’re trying to author, it’d bust out and change `<a>` tags everywhere…all while _not_ working your shadow DOM! D’oh! 🤦‍♂️

Thankfully, the oddly-worded but definitely cool one-two punch of CSS module scripts plus constructable stylesheets lets us import a `.css` file directly into a `.js` file, apply those styles to a web component, and _Bob’s your uncle!_ **It. Just. Works.** 😎

### And Lest We Forget…CSS Itself Has Gotten Freakishly Great

All the aforementioned whiz-bang goodness of component authoring wouldn’t help us all that much if we were still saddled with a frustrating & limited & hacky vanilla styling experience.

Thankfully, all is well now on that front. The CSS of 2022 is so far superior to the CSS of 2012 that it might as well be a whole new language. It’s such a dramatic leap forward that **I’m writing a whole Spicy Web course on it**. Oh yeah. 🙌

But for now, let’s just for the sake of argument take it as a given that vanilla CSS is fantastic, and move on to the demonstration of how all these puzzle pieces fit together.

### First, Tokens and a Polyfill

(If you want to skip down to the demo link, go for it. Otherwise keep reading to implement this yourself.)

Let’s save the simplest HTML page imaginable:

```html
<!doctype html>
<html>
  <head>
    <title>Hello World</title>
  </head>
  <body>
    <h1>Hello Buildless World!</h1>

    <p>The future is now.</p>
  </body>
</html>

```

Fire up a web server of your choosing and you’ll see the page appear in all its retro Times New Roman glory.

Let’s modernize things a bit. We’ll add a `styles/index.css` file, import some design tokens from [Open Props](https://open-props.style/), and make a few improvements:

```css
@import "https://unpkg.com/open-props";

body {
  background: var(--yellow-0);
  font-family: var(--font-sans);
}

p {
  font-size: var(--font-size-3);
}

```

Now we’ll add a `link` tag to our HTML `<head>`:

```html
<link href="/styles/index.css" rel="stylesheet" />

```

Refresh and you should see a more streamlined sans-serif font and a creamy background color.

Now, before we get into the JavaScript side of things, we’ll need to add a couple of polyfills—unless you _only_ care about Chromium-based browsers. I personally use Safari and Firefox and occasionally Edge (I run a Chrome-free household here!), so it’s important to me to use technologies which will work cross-browser.

We’ll add a couple of polyfills to our `<head>`: one for import maps + module scripts, and one for constructable stylesheets:

```html
<!-- Polyfill CSS/JSON module imports -->
<script>window.esmsInitOptions = { polyfillEnable: ['css-modules', 'json-modules'] }</script>
<script async src="https://ga.jspm.io/npm:es-module-shims@1.5.4/dist/es-module-shims.js"></script>
<script async src="https://unpkg.com/construct-style-sheets-polyfill@3.1.0/dist/adoptedStyleSheets.js"></script>

```

I know it looks odd to use two different CDNs, but I wasn’t able to get these working with only one or another. Go figure.

### Here Comes the Import Map

With that out of the way, we can go ahead and create an import map. I’m using [Lit](https://lit.dev/), a fast web component base library, to demonstrate how easy it is to write components with a “modern” DX, but if you wanted to write purely vanilla web components you wouldn’t even need this particular map. Anyway, let’s try adding the following map to `<head>` right below the polyfills:

```javascript
<!-- Add import for Lit, etc. -->
<script type="importmap">
  {
    "imports": {
      "lit": "https://ga.jspm.io/npm:lit@2.2.3/index.js"
    },
    "scopes": {
      "https://ga.jspm.io/": {
        "@lit/reactive-element": "https://ga.jspm.io/npm:@lit/reactive-element@1.3.2/development/reactive-element.js",
        "lit-element/lit-element.js": "https://ga.jspm.io/npm:lit-element@3.2.0/development/lit-element.js",
        "lit-html": "https://ga.jspm.io/npm:lit-html@2.2.3/development/lit-html.js"
      }
    }
  }
</script>

```

Don’t worry if that all looks confusing. There are [tools to help you generate import maps](https://jspm.org/docs/workflows) just by typing in your NPM dependency name(s).

Let’s try writing a very simple web component! Create a `components/index.js` file:

```javascript
import "./example-component.js"

```

and the `components/example-component.js` file to go with it:

```javascript
import { LitElement, html } from "lit"

class ExampleComponent extends LitElement {
  render() {
    return html`
      <p>Rendering a Lit component. This is cool! <slot></slot></p>
    `
  }
}

customElements.define("example-component", ExampleComponent)

```

Time to add our script tag to the bottom of `<head>`:

```html
<script src="/components/index.js" type="module"></script>

```

And then add the custom element somewhere in our HTML `<body>`:

```html
<example-component>🎉</example-component>

```

Refresh the page and you should see a paragraph which reads: **Rendering a Lit component. This is cool!** 🎉

### Construct Those Stylesheets

If you’re a big fan of SFCs as seen in other front-end frameworks such as Vue & Svelte, you’ll probably appreciate that you can write “vanilla” CSS right inside of tagged template literals inside of Lit components:

```javascript
import { LitElement, css, html } from "lit"

class ExampleComponent extends LitElement {
  static styles = css`
    p {
      color: var(--cyan-8);
    }
  `

  render() {
    return html`
      <p>Rendering a Lit component. This is cool! <slot></slot></p>
    `
  }
}

customElements.define("example-component", ExampleComponent)

```

In case you’re wondering: with the right IDE plugin (like [lit-plugin for VSCode](https://marketplace.visualstudio.com/items?itemName=runem.lit-plugin)), you get syntax highlighting and all the niceties you’d expect. (And of course because you’re using the shadow DOM, the `<p>` style here will _only_ affect markup inside of `example-component` and nowhere else on your page.)

_However_, as I’ve mentioned before, I really do like keeping my component CSS in `.css` files—what are sometimes referred to as sidecar stylesheets. For larger, more complex components, I definitely appreciate being able to maintain the structure + behavior of the component as a separate concern from the styling. Also if I’m working with designers on a team or developers with past experience with Sass and so forth, it’s a huge win.

So we should just be able to import styles like this, right?

```javascript
import "./example-component.css"

```

LOL no. That simply does not work at all, which has been a problem for the longest time. _Thankfully_, we have a savior on the horizon: **CSS module scripts** to the rescue!

I really think that’s kind of a confusing term for this, but nobody asked me. Also it’s totally different from the prior art known as [CSS Modules](https://css-tricks.com/css-modules-part-1-need/), though in the end it sort of accomplishes the same goal. Go figure.

Anyway, by using a default import name combined with a special `assert { type: "css" }` ending, we receive a `CSSStyleSheet` object representing the styles in those files. We can then apply those styles to our component, courtesy of the new **constructable stylesheets** spec. Lit supports this already, so we can simply rewrite our component like so:

```javascript
import { LitElement, html } from "lit"
import sheet from "./example-component.css" assert { type: "css" }

class ExampleComponent extends LitElement {
  static styles = [sheet]

  render() {
    return html`
      <p>Rendering a Lit component. This is cool! <slot></slot></p>
    `
  }
}

customElements.define("example-component", ExampleComponent)

```

We can then add a true sidecar stylesheet in the `example-component.css` file:

```css
p {
  color: var(--cyan-8);
}

```

_Are you getting it yet?_ Using nothing but the browser (polyfills notwithstanding) with no build tools required, you can author encapsulated components with scoped styles, save those `.js`+`.css` file combos right there in your project folders, import them into your HTML as ES modules, and **It. Just. Works.** 🙌

The only thing keeping this from being truly gobsmackingly fantastic is we still don’t have native “HTML imports”. So for example if we wanted to have a few HTML pages in our project folder to share a common header or nav bar or footer, we couldn’t do it unless we encapsulated all of those inside of web components—which I honestly can’t recommend at all for several reasons, not the least of which is your page layout will be horribly broken without client-side JavaScript enabled and executed.

So…in the end…we really do need _some_ kind of build-time or server-side tooling to offer a template syntax of some sort so it’s easy to share template partials across pages. And we probably also want a way to author content in some nice format like Markdown. And…yeah. Perhaps we’re not _quite_ ready to embrace this “buildless future”.

Nevertheless, we’re getting close. _Reeeeally_ close. Perhaps once we arrive at the 40th anniversary of the web in 2030, we’ll have a way to import a chunk of HTML into another chunk of HTML! 😆🤪🤨🤷‍♂️

All joking aside, there _are_ some [promising discussions along these lines](https://github.com/whatwg/html/issues/2791), as well as various custom elements ([example](https://shoelace.style/components/include), [another example](https://github.com/justinfagnani/html-include-element)) to provide workarounds…given the aforementioned caveats. It really does feel like a concept that’s long overdue.

### See This All In Action

[Demo](https://buildless-static.onrender.com/) [Repo](https://github.com/jaredcwhite/buildless-static-site)

Here’s a simple project which shows all of this working as well as a vanilla web component sitting alongside a Lit-based web component. The Node.js web server is only there to be an utterly basic web server. As you can see by inspecting the `package.json` file, there are _zero_ front-end dependencies. Our import map is what allows Lit to work when the page loads.

By the way, if you want just a smidge more tooling at your disposal, check out [Web Dev Server](https://modern-web.dev/docs/dev-server/overview/). It will handle rewriting bare module import specifiers for you automatically and even let you use local `node_modules` without needing to add import maps at all. I wouldn’t necessarily recommend this as a soup-to-nuts buildless workflow as it results in a sort of dependency on using this particular server (hence the name Web _Dev_ Server I suppose). Still, it’s an astonishingly simple way to spin up a new web project to kick the tires on some new library or feature approach.

So there you have it folks: **a sneak peak at the future of buildless modern web development**. We’re not quite there yet, I’ll be the first to admit. But perhaps for the first time in forever, **I’m reaching for my shades**. Because…the future. It’s…bright. _Keep up people!_ 😎
