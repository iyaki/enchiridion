---
title: "Modern CSS Code Snippets | modern.css"
notion_id: 30d54f1c-7d23-813d-aa25-ce2f4dd6e082
notion_url: https://app.notion.com/p/Modern-CSS-Code-Snippets-modern-css-30d54f1c7d23813daa25ce2f4dd6e082
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://modern-css.com/
tags: ["Tool", "Article", "modern.css", "English", "Web Development", "CSS", "Frontend", "Animation"]
---
## All comparisons

22 snippets

Browser compatibility: All Newly available Widely available Limited

[   Animation Intermediate](https://modern-css.com/custom-easing-without-cubic-bezier-guessing/)

### [Custom easing curves without cubic-bezier guessing](https://modern-css.com/custom-easing-without-cubic-bezier-guessing/)

[Old /* JS animation library */
anime({ targets: el,
 easing: 'easeOutBounce' });   see modern →   87% →](https://modern-css.com/custom-easing-without-cubic-bezier-guessing/)

[Layout Beginner  ](https://modern-css.com/preventing-layout-shift-from-scrollbar/)[**Preventing layout shift from scrollbar appearance**](https://modern-css.com/preventing-layout-shift-from-scrollbar/)[    Old body { overflow-y: scroll; }
/* or hardcode the scrollbar width */
body { padding-right: 17px; }   see modern →   90% →](https://modern-css.com/preventing-layout-shift-from-scrollbar/)

[Layout Beginner  ](https://modern-css.com/scrollbar-styling-without-webkit-pseudo-elements/)[**Scrollbar styling without -webkit- pseudo-elements**](https://modern-css.com/scrollbar-styling-without-webkit-pseudo-elements/)[    Old /* webkit only */
::-webkit-scrollbar { width: 8px; }
::-webkit-scrollbar-thumb { background: #888; }   see modern →   75% →](https://modern-css.com/scrollbar-styling-without-webkit-pseudo-elements/)

[Selector Beginner  ](https://modern-css.com/form-validation-styles-without-javascript/)[**Form validation styles without JavaScript**](https://modern-css.com/form-validation-styles-without-javascript/)[    Old // JS: add .touched on blur
el.addEventListener('blur', () =>
 el.classList.add('touched'))
/* .touched:invalid { color: red } */   see modern →   85% →](https://modern-css.com/form-validation-styles-without-javascript/)

[Layout Beginner  ](https://modern-css.com/auto-growing-textarea-without-javascript/)[**Auto-growing textarea without JavaScript**](https://modern-css.com/auto-growing-textarea-without-javascript/)[    Old // JS: resize on every keystroke
el.addEventListener('input', () => {
 el.style.height = 'auto';
 el.style.height = el.scrollHeight + 'px'; })   see modern →   73% →](https://modern-css.com/auto-growing-textarea-without-javascript/)

[Animation Beginner  ](https://modern-css.com/smooth-height-auto-animations-without-javascript/)[**Smooth height auto animations without JavaScript**](https://modern-css.com/smooth-height-auto-animations-without-javascript/)[    Old // measure, set px, then snap to auto
el.style.height = el.scrollHeight + 'px';
el.addEventListener('transitionend', ...)   see modern →   69% →](https://modern-css.com/smooth-height-auto-animations-without-javascript/)

[Color Intermediate  ](https://modern-css.com/vivid-colors-beyond-srgb/)[**Vivid colors beyond sRGB**](https://modern-css.com/vivid-colors-beyond-srgb/)[    Old .hero {
 color: rgb(200, 80, 50);
}
/* sRGB only, washed on P3 */   see modern →   90% →](https://modern-css.com/vivid-colors-beyond-srgb/)

[Color Advanced  ](https://modern-css.com/color-variants-without-sass-functions/)[**Color variants without Sass functions**](https://modern-css.com/color-variants-without-sass-functions/)[    Old /* Sass: lighten($brand, 20%), darken($brand, 10%) */
.btn { background: #e0e0e0; }   see modern →   87% →](https://modern-css.com/color-variants-without-sass-functions/)

[Typography Beginner  ](https://modern-css.com/drop-caps-without-float-hacks/)[**Drop caps without float hacks**](https://modern-css.com/drop-caps-without-float-hacks/)[    Old .drop-cap::first-letter {
 float: left;
 font-size: 3em; line-height: 1;
}   see modern →   91% →](https://modern-css.com/drop-caps-without-float-hacks/)

[Workflow Intermediate  ](https://modern-css.com/lazy-rendering-without-intersection-observer/)[**Lazy rendering without IntersectionObserver**](https://modern-css.com/lazy-rendering-without-intersection-observer/)[    Old // JS IntersectionObserver
new IntersectionObserver(
 (entries) => { /* render */ }
).observe(el);   see modern →   93% →](https://modern-css.com/lazy-rendering-without-intersection-observer/)

[Layout Beginner  ](https://modern-css.com/dropdown-menus-without-javascript-toggles/)[**Dropdown menus without JavaScript toggles**](https://modern-css.com/dropdown-menus-without-javascript-toggles/)[    Old .menu { display: none; }
.menu.open { display: block; }
/* + JS: click, clickOutside, ESC, aria */   see modern →   86% →](https://modern-css.com/dropdown-menus-without-javascript-toggles/)

[Workflow Advanced  ](https://modern-css.com/scoped-styles-without-bem-naming/)[**Scoped styles without BEM naming**](https://modern-css.com/scoped-styles-without-bem-naming/)[    Old // BEM: .card__title, .card__body
.card__title { â€¦ }
.card__body { â€¦ }
// or CSS Modules / styled-components */   see modern →   84% →](https://modern-css.com/scoped-styles-without-bem-naming/)

[Workflow Advanced  ](https://modern-css.com/typed-custom-properties-without-javascript/)[**Typed custom properties without JavaScript**](https://modern-css.com/typed-custom-properties-without-javascript/)[    Old // --hue was a string, no animation
:root { --hue: 0; }
hsl(var(--hue), â€¦) /* no interpolation */   see modern →   92% →](https://modern-css.com/typed-custom-properties-without-javascript/)

[Animation Intermediate  ](https://modern-css.com/animating-display-none-without-workarounds/)[**Animating display none without workarounds**](https://modern-css.com/animating-display-none-without-workarounds/)[    Old // wait for transitionend then display:none
el.addEventListener('transitionend', â€¦)
visibility + opacity + pointer-events   see modern →   85% →](https://modern-css.com/animating-display-none-without-workarounds/)

[Animation Intermediate  ](https://modern-css.com/entry-animations-without-javascript-timing/)[**Entry animations without JavaScript timing**](https://modern-css.com/entry-animations-without-javascript-timing/)[    Old // add class after paint
requestAnimationFrame(() => {
 el.classList.add('visible');
});   see modern →   85% →](https://modern-css.com/entry-animations-without-javascript-timing/)

[Animation Advanced  ](https://modern-css.com/page-transitions-without-a-framework/)[**Page transitions without a framework**](https://modern-css.com/page-transitions-without-a-framework/)[    Old // Barba.js or React Transition Group
Barba.init({ â€¦ })
transition hooks + duration state   see modern →   89% →](https://modern-css.com/page-transitions-without-a-framework/)

[Typography Beginner  ](https://modern-css.com/balanced-headlines-without-manual-line-breaks/)[**Balanced headlines without manual line breaks**](https://modern-css.com/balanced-headlines-without-manual-line-breaks/)[    Old // manual <br> or Balance-Text.js
h1 { text-align: center; }
.balance-text /* JS lib */   see modern →   87% →](https://modern-css.com/balanced-headlines-without-manual-line-breaks/)

[Color Intermediate  ](https://modern-css.com/dark-mode-colors-without-duplicating-values/)[**Dark mode colors without duplicating values**](https://modern-css.com/dark-mode-colors-without-duplicating-values/)[    Old @media (prefers-color-scheme: dark) {
 color: #eee;
}   see modern →   83% →](https://modern-css.com/dark-mode-colors-without-duplicating-values/)

[Layout Advanced  ](https://modern-css.com/aligning-nested-grids-without-duplicating-tracks/)[**Aligning nested grids without duplicating tracks**](https://modern-css.com/aligning-nested-grids-without-duplicating-tracks/)[    Old .child-grid {
 grid-template-columns: 1fr 1fr 1fr;
/* duplicate parent tracks */
}   see modern →   88% →](https://modern-css.com/aligning-nested-grids-without-duplicating-tracks/)

[Workflow Beginner  ](https://modern-css.com/nesting-selectors-without-sass-or-less/)[**Nesting selectors without Sass or Less**](https://modern-css.com/nesting-selectors-without-sass-or-less/)[    Old // requires Sass compiler
.nav {
 & a { color: #888; }
}   see modern →   91% →](https://modern-css.com/nesting-selectors-without-sass-or-less/)

[Colors Intermediate  ](https://modern-css.com/mixing-colors-without-a-preprocessor/)[**Mixing colors without a preprocessor**](https://modern-css.com/mixing-colors-without-a-preprocessor/)[    Old // Sass required
$blend: mix(
 $blue, $pink, 60%);   see modern →   89% →](https://modern-css.com/mixing-colors-without-a-preprocessor/)

[Selectors Intermediate  ](https://modern-css.com/selecting-parent-elements-without-javascript/)[**Selecting parent elements without JavaScript**](https://modern-css.com/selecting-parent-elements-without-javascript/)[    Old // JavaScript required
el.closest('.parent')
 .classList.add(…)   see modern →   94% →](https://modern-css.com/selecting-parent-elements-without-javascript/)
