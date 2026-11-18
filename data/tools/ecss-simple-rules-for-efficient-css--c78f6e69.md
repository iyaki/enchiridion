---
title: "ECSS - Simple Rules for efficient CSS"
notion_id: c78f6e69-6137-43db-acbb-ef5940434e59
notion_url: https://app.notion.com/p/ECSS-Simple-Rules-for-efficient-CSS-c78f6e69613743dbacbbef5940434e59
last_edited: 2024-06-11T16:32:00.000Z
source_url: https://ecss.info/en/
tags: ["Website", "Tool", "English", "CSS", "Frontend"]
---
## Values

### Intentionality

**Clear objectives, rational decisions.**

Every parcel of CSS has an objective. This objective should be clear to anyone reading the code. Be it HTML or CSS. Selectors are the perfect vehicle for intentionality. ECSS incentivizes the use of rational selectors.

### Consistency

**Repeating patterns save time & energy.**

Flexible naming guidelines give the codebase a consistency encouraging standardization and reuse. Prefixes and usage guidelines enforced through linting insures everyone on your team will follow ECSS rules.

### Simplicity

**“Simple” is writing less code & and limiting dependencies.**

Class soup & divitis aren’t simple. Untold numbers of `@media` queries aren’t either. But modern CSS is. By accepting CSS for what it is (a designing language!) you can do _a lot_ of styling mileage with very little code.

### Expressivity

**Code that speaks for itself is actually achievable.**

“thumbnail as-circle with-border” is instantly understandable while “h-10 w-10 bdr-50 br-1 overh” is not. Code should communicate information. The clearer the information, the easier it is to understand the system. Mistaken semantics should give way to expressiveness. _Expressive best practices work._

### Predictability

**Using simple but consistent authoring rules leads to predictability.**

It feels good to dive into unknown code… while already knowing what it looks like. Consistency, repeating patterns, simplicity (yes, again) all coalesce into predictable code. And predictable code means less stress, less friction and happier teams!

### Sustainability

**Vanilla CSS is future ready.**

In practice it means being at the forefront of progress and evolution. No need to wait for third parties to implement new features. You can use `clamp()` as of now and say goodbye to 80% of your media queries. No more `sm-this md-that`. Only clean and modern code.

## Guiding principles

### CSS is a _designing_ language, not a programming one.

CSS authoring should reflect _design_ patterns, concepts and perspectives. With global (sorry) scales, tokens, rules, etc. you limit repetition and facilitate future styling extension.

### Empowering designers to refine their work autonomously is efficient.

Making the needed parts of CSS being intelligible (code-savvy or not!) to designers cuts stress, time and promotes good allocation of human resources. Refining design in the browser is a great vector of economic profitability.

### Not all CSS properties are born equal.

The box-model properties (`display`, `position`, `width`, etc.) are extremely sensitive and potentially way more disruptive than, say, color or typographic ones. The former are the stuff of professional CSS authors and the latter are the ones designers should be able to use, modify and experiment with in context.

### Properties & units have specific roles and usage.

Limiting their use to these roles enhances clarity and promotes consistency as well as predictability. The same goes for units. Some suggest type, others, images or rhythm. Use them intentionally & rationally to communicate meaning.

### Components and patterns _need_ structure.

Trying to abstract it away is vain. Instead, documentation and dilligence are key. HTML is semantically rich, so choose the right tags and refrain from changing them for superficial reasons. One would not remove or change `class="header"` without understanding its purpose first, then one should not change or remove `<header></header>` on a whim either.

### CSS Selectors are vehicles for intention.

Every selector particle must have a purpose and if not, must be discarded with. This extraneous `div` in `.card div h2` is what makes CSS brittle. These are _not_ [“best practices”](https://adamwathan.me/css-utility-classes-and-separation-of-concerns/).

### Selector types all have a function.

Each one of them may be of use in a certain context. No absolute usage rule should be edicted. It’s not “always” or “never” but “it depends” and “why”. Absolute rules absolutely corrupt. Or something like that.

### Selection specificity should be harnessed, not rejected.

Nevertheless, “always” (ahem) keep it as low as possible. Yes, in 2015 it may have not been easy. But it’s 2023 and [new](https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Cascade_layers) [features](https://developer.mozilla.org/en-US/docs/Web/CSS/:where) are now widely supported. Use them.

### Global scope in CSS is not a sin.

Remember, _CSS is a designing language_ and the global scope fits perfectly in design’s fundamentally layered, dimensional approach. Programming best practices are not necessarily designing best practices. Embrace design best practices.

### Global design rules must be largely anonymous & unqualified.

Only a few declarations per ruleset only. We are talking rhythm, typography, color; not `display` or `position`. _This_ is the way of limiting repetition, bloat & complexity.

### One should wait until a concept is clearly understood before naming it.

Until then use HTML elements. Premature abstraction or [“solving problems not yet encountered”](https://en.wikipedia.org/wiki/Overengineering) are bigger problems than the use of `header` in CSS. Yes, these are programming best practices. But who told you CSS development was pure design?

### When more than 3 rules are repeated more than 3 times, a concept _is_ emerging.

Rules should then be concentrated by naming and using said concept instead of the bare rules. Reusability, expressivity and simplicity are then greatly improved.

### Named concepts must be autonomous and self-sufficient.

Named concepts should not be defined inside another named concept. If it ever is necessary, the child concept must be internal to the parent concept and is _not_ to be reused anywhere else.

### Manipulating named concepts throughout multiple files is strictly prohibited.

All rules for a named concept must reside in the same, unique CSS file. For authors and maintainers to have confidence in the system, there must a single [source of truth](https://en.wikipedia.org/wiki/Single_source_of_truth) for any concept.

### State, variant or structural are child concepts and must be namespaced & prefixed.

These concepts may be represented as child or combined classes. Never to be used alone, always alongside the parent concept. Yes, here are both “never” _and_ “always” in the same sentence. [“Double jeopardy, we are fine”](https://www.youtube.com/watch?v=Z58eTP2gcw0).

### Namespaced classes should generally be implemented with specificity reducing selectors.

Unless otherwise needed, one should strive to keep specificity at 21 or below. With a preference for around 10 to 12. Modern and widely supported pseudo-selectors must here be used.

### One may not come up at first with the efficient way to select interface elements.

One’s intention may not be clear at the beginning of a styling job. Instead of fruitlessly trying and trying, a code “quarantine” file should be temporarily used until the right selection manifests itself. Still, no quarantine file should ever be published publicly.

### HTML should be as simple and expressive as can be.

Refrain from over-containing, don’t use `<div>` where you could use `<aside>`. Dont wrap your single-level navigation in unordered lists. And yes, simple navigation is [accessible](https://dockyard.com/blog/2019/11/29/using-nav-without-a-list-element). Keep it simple… Suzy.

### Any HTML element should endorse one role only.

As per the famous programming [best practice](https://en.wikipedia.org/wiki/Single-responsibility_principle) (again, programming in designing), semantic tags are for… semantics whilst `<div>` or `<span>` tags are for graphical or logical division. Any type of grid should be implemented with `<div>` tags.

### Any `@media` query adaptation must be included in its related concept’s file.

Not as as a separate file, not as suffixed classes. If one uses utility classes, the provided rules should not be query dependent but rather be universally needed, in every media configuration.

### Whole stylesheets are better used when global.

Media query scopes should be used in HTML `<link>` tags to promote reuse, performance and optimization concerns. Each fundamental design layer should be autonomous, independent and removable.

### Component styling should only be served with live components.

Not as a big minified file in the `<head>` tag of the document. This way, the infamous problem of loading unused CSS is practically solved, without any technological debt. First render is also faster since only what is used is processed by the browser. As a bonus, you get the CSS file path in component file.

### Let the browser do as much of the rendering work as it can.

By embracing fluidity, adaptativity is mostly done by the browser. Less rules are needed to insure correct rendering in the infinite number of usage contexts.

### Make full use of freely available development tools by going with the grain of the Web platform.

By using “just-in-time” rulesets, design tokens, the cascade and intentional selection, the debugging & refining workflow is simpler, lighter and clearer. Devtools are our friends!

### Eschewing technological abstractions like high-level frameworks spurs lighter, simpler native code.

By writing native CSS, one uses CSS better; by writing native HTML, one writes HTML better. With better CSS and better HTML comes a better user experience.

## Authoring rules

### All component selectors _must_ start with its filename.

### Principles

- [Autonomous concepts](https://ecss.info/en/#autonomous-concept)
- [Concept manipulation](https://ecss.info/en/#concept-manipulation)

### Notes

To exclude a file from this rule in the **Stylelint** config, prefix the filename with a digit or “x-” as in “x-quarantine.css”

**Example**

### All paddings should first be uniform and then be adjusted if necessary.

### Principles

- [Specific role](https://ecss.info/en/#specific-role)

### Notes

Non-uniform paddings may be used in particular instances for `overflow` reasons or for added pseudo-content.

**Example**

### Paddings are to be applied on containing and interactive elements only.

### Principles

- [One role only](https://ecss.info/en/#one-role)

### Notes

Text elements should only host text styles while container elements can host theming styles. In some instances, container elements could host text styles to be inherited by children.

**Example**

### Horizontal margins should not be applied to text elements.

### Principles

- [Specific roles and usage](https://ecss.info/en/#specific-role)
- [One role only](https://ecss.info/en/#one-role)

### Notes

Horizontal margins can be applied to text elements in the case of `inline` elements such as icons or tags & on indented elements like `blockquote`.

**Example**

### Typographic elements should use inheritance when possible.

### Principles

- [Specific roles and usage](https://ecss.info/en/#specific-role)
- [Global scope](https://ecss.info/en/#global-scope)

**Example**

### Design tokens should be used for most numerical values.

### Principles

- [Designing language](https://ecss.info/en/#designing-language)
- [Designer empowerment](https://ecss.info/en/#design-empowerment)
- [Development tools](https://ecss.info/en/#devtools)

### Notes

Numerical values should only be used for exceptional alignments. These should be commented to communicate the intent. Even then, why not create a bespoke property meaningfully named?

**Example**

### Repeating default rules is discouraged.

### Principles

- [Selector intention](https://ecss.info/en/#selector-intention)
- [Selector function](https://ecss.info/en/#selector-function)

### Notes

Any instance of repeating default _must_ be justified in a comment.

**Example**

### All class entities besides components _must_ be prefixed.

### Principles

- [Child concepts](https://ecss.info/en/#child-concepts)
- [Selector intention](https://ecss.info/en/#selector-intention)
- [Selector function](https://ecss.info/en/#selector-function)
- [Autonomous concepts](https://ecss.info/en/#autonomous-concept)

### Notes

We propose this nomenclature of prefixes:

- `as-`: generic graphical form.
- `of-`: group of elements.
- `is-` interactive state of the system.
- `on-`: user interaction.
- `in-`: internal concept.
- `with-`: functional rule.
- `from-*-to-`: adaptive patterns.

**Example**

### Nesting is restricted to one level only.

### Principles

- [Child concepts](https://ecss.info/en/#child-concepts)
- [Selector intention](https://ecss.info/en/#selector-intention)

### Notes

Nesting can encourage authoring on autopilot, without thinking about our selection intention. That’s one reason why we will ultimately write `.card div h2` instead of `.card h2`, fragilizing our CSS. One level of nesting should be enough almost all cases.

**Example**

### The use of specific dimensions should be avoided.

### Principles

- [Properties are not equal](https://ecss.info/en/#properties-not-equal)
- [Specific role](https://ecss.info/en/#specific-role)
- [Browser rendering](https://ecss.info/en/#browser-rendering)

### Notes

The Web is fluid and authoring should embrace this fluidity. By choosing to limit dimensions of our components instead of dictating them, we delegate work to the browser, letting it do its job. Fixed dimensions should be only be applied to graphical elements like pictures, videos and icons.

**Example**

### Use attribute selectors to convey unicity.

### Principles

- [Selector intention](https://ecss.info/en/#selector-intention)
- [Specific role](https://ecss.info/en/#specific-role)
- [Harness specificity](https://ecss.info/en/#harness-specificity)

### Notes

Ids convey meaning but their CSS selector is a specificity nightmare. Use attribute selector to keep the meaning (unicity!) while keeping specificiy at sane levels. Other use cases includes form controls & inputs and js managed states (ie: `[data-state="active"]`)

**Example**

### The use of positioning and display rules should minimized.

### Principles

- [Properties are not equal](https://ecss.info/en/#properties-not-equal)
- [Browser rendering](https://ecss.info/en/#browser-rendering)

### Notes

Use these properties only on elements needing them. Use the narrowest selectors that makes sense in the context.

**Example**

### Overqualified selectors are discouraged.

### Principles

- [Harness specificity](https://ecss.info/en/#harness-specificity)
- [Selector intention](https://ecss.info/en/#selector-intention)

### Notes

Class entities should not be combined with tag selectors. Generic selectors can be used as a selection piece but not as qualifiers. Exceptionally, certain specific tags such as `html` or `body` could be used to restrict the application of a class entity.

**Example**

### Components cannot exert outside influence.

### Principles

- [Autonomous concepts](https://ecss.info/en/#autonomous-concept)
- [Properties are not equal](https://ecss.info/en/#properties-not-equal)
- [Specific roles and usage](https://ecss.info/en/#specific-role)

### Notes

Components are made to fit into “holes”. These holes are managed by another component or a higher-level utility (e.g. a grid or carousel). No component in this sense should incorporate margins. These are applied as a rhythm by a parent. Thus, the components are versatile and reusable in various contexts.

**Example**

### Overflow should not be hidden.

### Principles

- [Browser rendering](https://ecss.info/en/#browser-rendering)
- [Properties are not equal](https://ecss.info/en/#properties-not-equal)

### Notes

By hiding the way browsers treat overflowing items, we may hide content without knowing. By using `auto` instead of `hidden`, the oveflowing content will be accessible _and_ the scrollbar will signal the design problem instead of hiding it. Hiding scrollbars may be necessary in some particular composition (see this very site’s portrait adaptation for an example!).

**Example**

### Numbers in class names are to be avoided.

### Principles

- [Selector intention](https://ecss.info/en/#selector-intention)

### Notes

Numbers may convey intention in grid system cases as in, for example, `grid-10` or `col-3`. Still, the utility of these systems is low in modern CSS times. Leaving them aside is suggested.

**Example**

### Magic numbers _must_ be avoided or explained in a line comment.

### Principles

- [Designing language](https://ecss.info/en/#designing-language)
- [Designer empowerment](https://ecss.info/en/#design-empowerment)
- [Specific roles and usage](https://ecss.info/en/#specific-role)

### Notes

Value and unit choices should be dictated by the graphic language and its design tokens. Absolute values in numerical form should be reserved to special cases like bespoke icon alignment or bullet list padding. The intention _must_ be clear in itself or commented if not.

**Example**

### Rhythm _must_ be applied in one direction only, preferably top.

### Principles

- [Designing language](https://ecss.info/en/#designing-language)
- [Specific roles and usage](https://ecss.info/en/#specific-role)

### Notes

Spacing is rhythm, a fundamental design concept. Rhythm occurs _between_ graphical elements and should be applied as such in CSS. The [lobotomized owl](https://alistapart.com/article/axiomatic-css-and-lobotomized-owls/) selector (_+_) is tailor made for this task. When using a flex container, the `gap` property does it all for you!

**Example**

### The use of problematic units is discouraged.

### Principles

- [Specific roles and usage](https://ecss.info/en/#specific-role)

### Notes

Static viewport units (`vh`, `vw`, `vmin`, `vmax`) are famously problematic, particularly when used as `100vh` or `100vw`. [Dynamic viewport units](https://caniuse.com/viewport-unit-variants) are better suited for this task.

**Example**

### Tag selectors should be leveraged inside components.

### Principles

- [Component structure](https://ecss.info/en/#component-structure)
- [Expressive HTML](https://ecss.info/en/#expressive-html)

### Notes

Components need to prescribe a certain HTML structure. Using said structure through intentional tag selectors is the best way to insure structural persistence, to keep specificity and complexity low and prevent naming fatigue. Why rename that which is already named?

**Example**

### One entity per file with a soft cap of about a hundred lines.

### Principles

- [Autonomous concepts](https://ecss.info/en/#autonomous-concept)
- [Concept manipulation](https://ecss.info/en/#concept-manipulation)

### Notes

This is one of the most important rules for achieving sustainable CSS. One autonomous concept per file prevents the multiplication of overrides and redefinition across the codebase. It’s even easy to enforce with the [ECSS Stylelint config](https://www.npmjs.com/package/@efficientcss/stylelint-config-ecss).

**Example**

### Selectors _must_ strictly include only necessary selection parts.

### Principles

- [Component structure](https://ecss.info/en/#component-structure)
- [Selector intention](https://ecss.info/en/#selector-intention)
- [Harness specificity](https://ecss.info/en/#harness-specificity)

### Notes

This mistake, being passed as “best practices”, is the source of many self-induced CSS headaches. Fortunately, the solution is simple: _do not use selection parts that you could remove without changing the original intention_.

**Example**

### Duplication of a set of more than 3 rules is discouraged.

### Principles

- [Rule repetition](https://ecss.info/en/#rule-repetition)
- [Autonomous concepts](https://ecss.info/en/#autonomous-concept)

### Notes

Repeating a pattern instead of naming it and grouping its rules inside a unique declaration block is a classic CSS mistake. And it can be hard to find and target instances of this problem. Tools like PostCSS plugins can help by identifying repetition meeting a defined threshold throughout a codebase.

**Example**

### The use of the global scope is encouraged for all fundamental design layers.

### Principles

- [Designing language](https://ecss.info/en/#designing-language)
- [Designer empowerment](https://ecss.info/en/#design-empowerment)
- [Global scope](https://ecss.info/en/#global-scope)
- [Global rules](https://ecss.info/en/#global-rules)
- [Harness specificity](https://ecss.info/en/#harness-specificity)

**Example**

### HTML structure should be as flat as possible and bereft of single tag nesting.

### Principles

- [Component structure](https://ecss.info/en/#component-structure)
- [Expressive HTML](https://ecss.info/en/#expressive-html)

### Notes

Sometimes it could make sense to add a div even if it’s the only child. The case of a content grid comes to mind. See this page’s source for an example.

**Example**

### Specificity _must_ be kept as low as possible.

### Principles

- [Harness specificity](https://ecss.info/en/#harness-specificity)
- [Child specificity](https://ecss.info/en/#child-specificity)

### Notes

Sometimes it makes sense to add specificity. But don’t do it until proven necessary.

**Example**

## Related tools

- [ECSS Stylelint config](https://www.npmjs.com/package/@efficientcss/stylelint-config-ecss) (get live authoring tips!)
- [Scaffolding & code library](https://github.com/efficientcss/scaffolding) (use it… or not!)
- [Stylelint](https://stylelint.io/) (lint and inform)
- [CSSCSS](https://github.com/zmoazeni/csscss) (find rule duplication in your codebase)
- [PurifyCSS](https://github.com/purifycss/purifycss), [PurgeCSS](https://github.com/FullHuman/purgecss) or Chrome Coverage devtool
