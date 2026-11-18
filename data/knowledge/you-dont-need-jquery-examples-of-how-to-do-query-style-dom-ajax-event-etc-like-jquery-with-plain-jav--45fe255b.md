---
title: "You-Dont-Need-jQuery - Examples of how to do query, style, dom, ajax, event etc like jQuery with plain javascript"
notion_id: 45fe255b-050c-4ffb-9d37-da1374f4570d
notion_url: https://app.notion.com/p/You-Dont-Need-jQuery-Examples-of-how-to-do-query-style-dom-ajax-event-etc-like-jQuery-with-pla-45fe255b050c4ffb9d37da1374f4570d
last_edited: 2023-01-18T14:10:00.000Z
source_url: https://github.com/camsong/You-Dont-Need-jQuery#you-might-dont-need-jquery
tags: ["English", "Español", "Others", "Frontend", "Javascript", "Book"]
---
## You (Might) Don't Need jQuery

Frontend environments evolve rapidly nowadays and modern browsers have already implemented a great deal of DOM/BOM APIs which are good enough for production use. We don't have to learn jQuery from scratch for DOM manipulation or event handling. In the meantime, thanks to the spread of frontend libraries such as React, Angular and Vue, manipulating the DOM directly becomes anti-pattern, so that jQuery usage has never been less important. This project summarizes most of the alternatives in native Javascript implementation to jQuery methods, with IE 10+ support.

### **Notice**

1. jQuery is still a great library and has many valid use cases. Don’t migrate away if you don’t want to!
2. The alternatives are not completely equivalent in all scenarios, and it is recommended that you test it before using it.

## Table of Contents

1. [Translations](https://github.com/camsong/You-Dont-Need-jQuery#translations)
2. [Query Selector](https://github.com/camsong/You-Dont-Need-jQuery#query-selector)
3. [CSS & Style](https://github.com/camsong/You-Dont-Need-jQuery#css--style)
4. [DOM Manipulation](https://github.com/camsong/You-Dont-Need-jQuery#dom-manipulation)
5. [Ajax](https://github.com/camsong/You-Dont-Need-jQuery#ajax)
6. [Events](https://github.com/camsong/You-Dont-Need-jQuery#events)
7. [Utilities](https://github.com/camsong/You-Dont-Need-jQuery#utilities)
8. [Promises](https://github.com/camsong/You-Dont-Need-jQuery#promises)
9. [Animation](https://github.com/camsong/You-Dont-Need-jQuery#animation)
10. [Alternatives](https://github.com/camsong/You-Dont-Need-jQuery#alternatives)
11. [Browser Support](https://github.com/camsong/You-Dont-Need-jQuery#browser-support)

## Translations

- [한국어](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README.ko-KR.md)
- [正體中文](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README.zh-TW.md)
- [简体中文](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README.zh-CN.md)
- [Bahasa Melayu](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-my.md)
- [Bahasa Indonesia](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-id.md)
- [Português(PT-BR)](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README.pt-BR.md)
- [Tiếng Việt Nam](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-vi.md)
- [Español](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-es.md)
- [Русский](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-ru.md)
- [Кыргызча](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-kg.md)
- [Türkçe](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-tr.md)
- [Italiano](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-it.md)
- [Français](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-fr.md)
- [日本語](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-ja.md)
- [Polski](https://github.com/camsong/You-Dont-Need-jQuery/blob/master/README-pl.md)

## Query Selector

In place of common selectors like class, id or attribute we can use `document.querySelector` or `document.querySelectorAll` for substitution. The differences lie in:

- `document.querySelector` returns the first matched element
- `document.querySelectorAll` returns all matched elements as NodeList. It can be converted to Array using `Array.prototype.slice.call(document.querySelectorAll(selector));` or any of the methods outlined in [makeArray](https://github.com/camsong/You-Dont-Need-jQuery#makeArray)
- If there are no elements matched, jQuery and `document.querySelectorAll` will return `[]`, whereas `document.querySelector` will return `null`.

> 

Notice: `document.querySelector` and `document.querySelectorAll` are quite **SLOW**, thus try to use `document.getElementById`, `document.getElementsByClassName` or `document.getElementsByTagName` if you want to get a performance bonus.

-  

[1.0](https://github.com/camsong/You-Dont-Need-jQuery#1.0)  Query by selector

```plain text
// jQuery
$('selector');

// Native
document.querySelectorAll('selector');
```

-  

[1.1](https://github.com/camsong/You-Dont-Need-jQuery#1.1)  Query by class

```plain text
// jQuery
$('.class');

// Native
document.querySelectorAll('.class');

// or
document.getElementsByClassName('class');
```

-  

[1.2](https://github.com/camsong/You-Dont-Need-jQuery#1.2)  Query by id

```plain text
// jQuery
$('#id');

// Native
document.querySelector('#id');

// or
document.getElementById('id');

// or
window['id']
```

-  

[1.3](https://github.com/camsong/You-Dont-Need-jQuery#1.3)  Query by attribute

```plain text
// jQuery
$('a[target=_blank]');

// Native
document.querySelectorAll('a[target=_blank]');
```

-  

[1.4](https://github.com/camsong/You-Dont-Need-jQuery#1.4)  Query in descendants

```plain text
// jQuery
$el.find('li');

// Native
el.querySelectorAll('li');
```

-  

[1.5](https://github.com/camsong/You-Dont-Need-jQuery#1.5)  Sibling/Previous/Next Elements

-  

All siblings

```plain text
// jQuery
$el.siblings();

// Native - latest, Edge13+
[...el.parentNode.children].filter((child) =>
 child !== el
);
// Native (alternative) - latest, Edge13+
Array.from(el.parentNode.children).filter((child) =>
 child !== el
);
// Native - IE10+
Array.prototype.filter.call(el.parentNode.children, (child) =>
 child !== el
);
```

-  

Previous sibling

```plain text
// jQuery
$el.prev();

// Native
el.previousElementSibling;
```

-  

Next sibling

```plain text
// jQuery
$el.next();

// Native
el.nextElementSibling;
```

-  

All previous siblings

```plain text
// jQuery (optional filter selector)
$el.prevAll($filter);

// Native (optional filter function)
function getPreviousSiblings(elem, filter) {
 var sibs = [];
 while (elem = elem.previousSibling) {
 if (elem.nodeType === 3) continue; // ignore text nodes
 if (!filter || filter(elem)) sibs.push(elem);
 }
 return sibs;
}
```

-  

All next siblings

```plain text
// jQuery (optional selector filter)
$el.nextAll($filter);

// Native (optional filter function)
function getNextSiblings(elem, filter) {
 var sibs = [];
 var nextElem = elem.parentNode.firstChild;
 do {
 if (nextElem.nodeType === 3) continue; // ignore text nodes
 if (nextElem === elem) continue; // ignore elem of target
 if (nextElem === elem.nextElementSibling) {
 if (!filter || filter(elem)) {
 sibs.push(nextElem);
 elem = nextElem;
 }
 }
 } while(nextElem = nextElem.nextSibling)
 return sibs;
 }
```

An example of filter function:

```plain text
function exampleFilter(elem) {
 switch (elem.nodeName.toUpperCase()) {
 case 'DIV':
 return true;
 case 'SPAN':
 return true;
 default:
 return false;
 }
}
```

-   

[1.6](https://github.com/camsong/You-Dont-Need-jQuery#1.6)  Closest

Return the first matched element by provided selector, traversing from current element up through its ancestors in the DOM tree.

```plain text
// jQuery
$el.closest(selector);

// Native - Only latest, NO IE
el.closest(selector);

// Native - IE10+
function closest(el, selector) {
 const matchesSelector = el.matches || el.webkitMatchesSelector || el.mozMatchesSelector || el.msMatchesSelector;

 while (el) {
 if (matchesSelector.call(el, selector)) {
 return el;
 } else {
 el = el.parentElement;
 }
 }
 return null;
}
```

-   

[1.7](https://github.com/camsong/You-Dont-Need-jQuery#1.7)  Parents Until

Get the ancestors of each element in the current set of matched elements, up to but not including the element matched by the selector, DOM node, or jQuery object.

```plain text
// jQuery
$el.parentsUntil(selector, filter);

// Native
function parentsUntil(el, selector, filter) {
 const result = [];
 const matchesSelector = el.matches || el.webkitMatchesSelector || el.mozMatchesSelector || el.msMatchesSelector;

 // match start from parent
 el = el.parentElement;
 while (el && !matchesSelector.call(el, selector)) {
 if (!filter) {
 result.push(el);
 } else {
 if (matchesSelector.call(el, filter)) {
 result.push(el);
 }
 }
 el = el.parentElement;
 }
 return result;
}
```

-  

[1.8](https://github.com/camsong/You-Dont-Need-jQuery#1.8)  Form

-  

Input/Textarea

```plain text
// jQuery
$('#my-input').val();

// Native
document.querySelector('#my-input').value;
```

-  

Get index of e.currentTarget between `.radio`

```plain text
// jQuery
$('.radio').index(e.currentTarget);

// Native
Array.from(document.querySelectorAll('.radio')).indexOf(e.currentTarget);
or
Array.prototype.indexOf.call(document.querySelectorAll('.radio'), e.currentTarget);
```

-   

[1.9](https://github.com/camsong/You-Dont-Need-jQuery#1.9)  Iframe Contents

`$('iframe').contents()` returns `contentDocument` for this specific iframe

-  

Iframe contents

```plain text
// jQuery
$iframe.contents();

// Native
iframe.contentDocument;
```

-  

Iframe Query

```plain text
// jQuery
$iframe.contents().find('.css');

// Native
iframe.contentDocument.querySelectorAll('.css');
```

-  

[1.10](https://github.com/camsong/You-Dont-Need-jQuery#1.10)  Get body

```plain text
// jQuery
$('body');

// Native
document.body;
```

-  

[1.11](https://github.com/camsong/You-Dont-Need-jQuery#1.11)  Attribute getter and setter

-  

Get an attribute

```plain text
// jQuery
$el.attr('foo');

// Native
el.getAttribute('foo');
```

-  

Set an attribute

```plain text
// jQuery
$el.attr('foo', 'bar');

// Native
el.setAttribute('foo', 'bar');
```

-  

Get a `data-` attribute

```plain text
// jQuery
$el.data('foo');

// Native (use `getAttribute`)
el.getAttribute('data-foo');

// Native (use `dataset` if only need to support IE 11+)
el.dataset['foo'];
```

-  

[1.12](https://github.com/camsong/You-Dont-Need-jQuery#1.12)  Selector containing string (case-sensitive)

```plain text
// jQuery
$("selector:contains('text')");

// Native
function contains(selector, text) {
 var elements = document.querySelectorAll(selector);
 return Array.from(elements).filter(function(element) {
 return RegExp(text).test(element.textContent);
 });
}
```

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[**back to top**](https://github.com/camsong/You-Dont-Need-jQuery#table-of-contents)

## CSS & Style

-  

[2.1](https://github.com/camsong/You-Dont-Need-jQuery#2.1)  CSS

-  

Get style

```plain text
// jQuery
$el.css('color');

// Native
// NOTE: Known bug, will return 'auto' if style value is 'auto'
const win = el.ownerDocument.defaultView;

// null means not to return pseudo styles
win.getComputedStyle(el, null).color;
```

-  

Set style

```plain text
// jQuery
$el.css({ color: '#f01' });

// Native
el.style.color = '#f01';
```

-  

Get/Set Styles

Note that if you want to set multiple styles once, you could refer to [setStyles](https://github.com/oneuijs/oui-dom-utils/blob/master/src/index.js#L194) method in oui-dom-utils package.

-  

Add class

```plain text
// jQuery
$el.addClass(className);

// Native
el.classList.add(className);
```

-  

Remove class

```plain text
// jQuery
$el.removeClass(className);

// Native
el.classList.remove(className);
```

-  

has class

```plain text
// jQuery
$el.hasClass(className);

// Native
el.classList.contains(className);
```

-  

Toggle class

```plain text
// jQuery
$el.toggleClass(className);

// Native
el.classList.toggle(className);
```

-   

[2.2](https://github.com/camsong/You-Dont-Need-jQuery#2.2)  Width & Height

Width and Height are theoretically identical, take Height as example:

-  

Window height

```plain text
// window height
$(window).height();

// without scrollbar, behaves like jQuery
window.document.documentElement.clientHeight;

// with scrollbar
window.innerHeight;
```

-  

Document height

```plain text
// jQuery
$(document).height();

// Native
const body = document.body;
const html = document.documentElement;
const height = Math.max(
 body.offsetHeight,
 body.scrollHeight,
 html.clientHeight,
 html.offsetHeight,
 html.scrollHeight
);
```

-  

Element height

```plain text
// jQuery
$el.height();

// Native
function getHeight(el) {
 const styles = window.getComputedStyle(el);
 const height = el.offsetHeight;
 const borderTopWidth = parseFloat(styles.borderTopWidth);
 const borderBottomWidth = parseFloat(styles.borderBottomWidth);
 const paddingTop = parseFloat(styles.paddingTop);
 const paddingBottom = parseFloat(styles.paddingBottom);
 return height - borderBottomWidth - borderTopWidth - paddingTop - paddingBottom;
}

// accurate to integer（when `border-box`, it's `height - border`; when `content-box`, it's `height + padding`）
el.clientHeight;

// accurate to decimal（when `border-box`, it's `height`; when `content-box`, it's `height + padding + border`）
el.getBoundingClientRect().height;
```

-  

[2.3](https://github.com/camsong/You-Dont-Need-jQuery#2.3)  Position & Offset

-   

Position

Get the current coordinates of the element relative to the offset parent.

```plain text
// jQuery
$el.position();

// Native
{ left: el.offsetLeft, top: el.offsetTop }
```

-   

Offset

Get the current coordinates of the element relative to the document.

```plain text
// jQuery
$el.offset();

// Native
function getOffset (el) {
 const box = el.getBoundingClientRect();

 return {
 top: box.top + window.pageYOffset - document.documentElement.clientTop,
 left: box.left + window.pageXOffset - document.documentElement.clientLeft
 };
}
```

-   

[2.4](https://github.com/camsong/You-Dont-Need-jQuery#2.4)  Scroll Top

Get the current vertical position of the scroll bar for the element.

```plain text
// jQuery
$(window).scrollTop();

// Native
(document.documentElement && document.documentElement.scrollTop) || document.body.scrollTop;
```

[**back to top**](https://github.com/camsong/You-Dont-Need-jQuery#table-of-contents)

## DOM Manipulation

-   

[3.1](https://github.com/camsong/You-Dont-Need-jQuery#3.1)  Remove

Remove the element from the DOM.

```plain text
// jQuery
$el.remove();

// Native
el.parentNode.removeChild(el);
```

-  

[3.2](https://github.com/camsong/You-Dont-Need-jQuery#3.2)  Text

-   

Get text

Get the combined text contents of the element including their descendants,

```plain text
// jQuery
$el.text();

// Native
el.textContent;
```

-   

Set text

Set the content of the element to the specified text.

```plain text
// jQuery
$el.text(string);

// Native
el.textContent = string;
```

-  

[3.3](https://github.com/camsong/You-Dont-Need-jQuery#3.3)  HTML

-  

Get HTML

```plain text
// jQuery
$el.html();

// Native
el.innerHTML;
```

-  

Set HTML

```plain text
// jQuery
$el.html(htmlString);

// Native
el.innerHTML = htmlString;
```

-   

[3.4](https://github.com/camsong/You-Dont-Need-jQuery#3.4)  Append

Append child element after the last child of parent element

```plain text
// jQuery: unified syntax for DOMString and Node objects
$parent.append(newEl | '<div id="container">Hello World</div>');

// Native: different syntax
parent.insertAdjacentHTML('beforeend', '<div id="container">Hello World</div>');
parent.appendChild(newEl);

// Native (ES6-way): unified syntax
parent.append(newEl | '<div id="container">Hello World</div>');
```

-  

[3.5](https://github.com/camsong/You-Dont-Need-jQuery#3.5)  Prepend

```plain text
// jQuery: unified syntax for DOMString and Node objects
$parent.prepend(newEl | '<div id="container">Hello World</div>');

// Native: different syntax
parent.insertAdjacentHTML('afterbegin', '<div id="container">Hello World</div>');
parent.insertBefore(newEl, parent.firstChild);

// Native (ES6-way): unified syntax
parent.prepend(newEl | '<div id="container">Hello World</div>');
```

-   

[3.6](https://github.com/camsong/You-Dont-Need-jQuery#3.6)  insertBefore

Insert a new node before the selected elements

```plain text
// jQuery
$newEl.insertBefore(selector);

// Native (HTML string)
el.insertAdjacentHTML('beforebegin ', '<div id="container">Hello World</div>');

// Native (Element)
const el = document.querySelector(selector);
if (el.parentNode) {
 el.parentNode.insertBefore(newEl, el);
}
```

-   

[3.7](https://github.com/camsong/You-Dont-Need-jQuery#3.7)  insertAfter

Insert a new node after the selected elements

```plain text
// jQuery
$newEl.insertAfter(selector);

// Native (HTML string)
el.insertAdjacentHTML('afterend', '<div id="container">Hello World</div>');

// Native (Element)
const el = document.querySelector(selector);
if (el.parentNode) {
 el.parentNode.insertBefore(newEl, el.nextSibling);
}
```

-   

[3.8](https://github.com/camsong/You-Dont-Need-jQuery#3.8)  is

Return `true` if it matches the query selector

```plain text
// jQuery - Notice `is` also works with a function, an existing jQuery object or a DOM element, which are not of concern here
$el.is(selector);

// Native
el.matches(selector);
```

-   

[3.9](https://github.com/camsong/You-Dont-Need-jQuery#3.9)  clone

Create a deep copy of an element: it copies the matched element as well as all of its descendant elements and text nodes.

```plain text
// jQuery. Sets parameter as `true` to indicate that event handlers should be copied along with the element.
$el.clone();

// Native
el.cloneNode();
```

-   

[3.10](https://github.com/camsong/You-Dont-Need-jQuery#3.10)  empty

Remove all child nodes

```plain text
// jQuery
$el.empty();

// Native
el.innerHTML = null;
```

-   

[3.11](https://github.com/camsong/You-Dont-Need-jQuery#3.11)  wrap

Wrap an HTML structure around each element

```plain text
// jQuery
$('.inner').wrap('<div class="wrapper"></div>');

// Native
Array.from(document.querySelectorAll('.inner')).forEach((el) => {
 const wrapper = document.createElement('div');
 wrapper.className = 'wrapper';
 el.parentNode.insertBefore(wrapper, el);
 wrapper.appendChild(el);
});
```

-   

[3.12](https://github.com/camsong/You-Dont-Need-jQuery#3.12)  unwrap

Remove the parents of the set of matched elements from the DOM

```plain text
// jQuery
$('.inner').unwrap();

// Native
Array.from(document.querySelectorAll('.inner')).forEach((el) => {
 let elParentNode = el.parentNode;

 if(elParentNode !== document.body) {
 elParentNode.parentNode.insertBefore(el, elParentNode);
 elParentNode.parentNode.removeChild(elParentNode);
 }
});
```

-   

[3.13](https://github.com/camsong/You-Dont-Need-jQuery#3.13)  replaceWith

Replace each element in the set of matched elements with the provided new content

```plain text
// jQuery
$('.inner').replaceWith('<div class="outer"></div>');

// Native (alternative) - latest, Edge17+
Array.from(document.querySelectorAll('.inner')).forEach((el) => {
 const outer = document.createElement('div');
 outer.className = 'outer';
 el.replaceWith(outer);
});

// Native
Array.from(document.querySelectorAll('.inner')).forEach((el) => {
 const outer = document.createElement('div');
 outer.className = 'outer';
 el.parentNode.replaceChild(outer, el);
});
```

-   

[3.14](https://github.com/camsong/You-Dont-Need-jQuery#3.14)  simple parse

Parse a string into HTML/SVG/XML

```plain text
// jQuery
$(`<ol>
 <li>a</li>
 <li>b</li>
</ol>
<ol>
 <li>c</li>
 <li>d</li>
</ol>`);

// Native
range = document.createRange();
parse = range.createContextualFragment.bind(range);

parse(`<ol>
 <li>a</li>
 <li>b</li>
</ol>
<ol>
 <li>c</li>
 <li>d</li>
</ol>`);
```

[**back to top**](https://github.com/camsong/You-Dont-Need-jQuery#table-of-contents)

## Ajax

[Fetch API](https://fetch.spec.whatwg.org/) is the new standard to replace XMLHttpRequest to do ajax. It works on Chrome and Firefox, you can use polyfills to make it work on legacy browsers.

Try [github/fetch](http://github.com/github/fetch) on IE9+ or [fetch-ie8](https://github.com/camsong/fetch-ie8/) on IE8+, [fetch-jsonp](https://github.com/camsong/fetch-jsonp) to make JSONP requests.

-  

[4.1](https://github.com/camsong/You-Dont-Need-jQuery#4.1)  Load data from the server and place the returned HTML into the matched element.

```plain text
// jQuery
$(selector).load(url, completeCallback)

// Native
fetch(url).then(data => data.text()).then(data => {
 document.querySelector(selector).innerHTML = data
}).then(completeCallback)
```

[**back to top**](https://github.com/camsong/You-Dont-Need-jQuery#table-of-contents)

## Events

For a complete replacement with namespace and delegation, refer to [https://github.com/oneuijs/oui-dom-events](https://github.com/oneuijs/oui-dom-events)

-  

[5.0](https://github.com/camsong/You-Dont-Need-jQuery#5.0)  Document ready by `DOMContentLoaded`

```plain text
 // jQuery
 $(document).ready(eventHandler);

 // Native
 // Check if the DOMContentLoaded has already been completed
 if (document.readyState !== 'loading') {
 eventHandler();
 } else {
 document.addEventListener('DOMContentLoaded', eventHandler);
 }

 // Native
 // Example 2 - Ternary Operator - Async
 // Check if the DOMContentLoaded has already been completed
 (async function() {
 (document.readyState !== 'loading') ?
 eventHandler() : document.addEventListener('DOMContentLoaded',
 function() {
 eventHandler(); // EventHandler
 });
 })();

 // Native
 // Example 3 - Ternary Operator - Non Async
 // Check if the DOMContentLoaded has already been completed
 (function() {
 (document.readyState !== 'loading') ?
 eventHandler() : document.addEventListener('DOMContentLoaded',
 function() {
 eventHandler(); // EventHandler
 });
 })();
```

-  

[5.1](https://github.com/camsong/You-Dont-Need-jQuery#5.1)  Bind an event with on

```plain text
// jQuery
$el.on(eventName, eventHandler);

// Native
el.addEventListener(eventName, eventHandler);
```

-  

[5.2](https://github.com/camsong/You-Dont-Need-jQuery#5.2)  Unbind an event with off

```plain text
// jQuery
$el.off(eventName, eventHandler);

// Native
el.removeEventListener(eventName, eventHandler);
```

-  

[5.3](https://github.com/camsong/You-Dont-Need-jQuery#5.3)  Trigger

```plain text
// jQuery
$(el).trigger('custom-event', {key1: 'data'});

// Native
if (window.CustomEvent) {
 const event = new CustomEvent('custom-event', {detail: {key1: 'data'}});
} else {
 const event = document.createEvent('CustomEvent');
 event.initCustomEvent('custom-event', true, true, {key1: 'data'});
}

el.dispatchEvent(event);
```

[**back to top**](https://github.com/camsong/You-Dont-Need-jQuery#table-of-contents)

## Utilities

Most of jQuery utilities are also found in the native API. Other advanced functions could be chosen from better utilities libraries, focusing on consistency and performance. [Lodash](https://lodash.com/) is a recommended replacement.

-                                                    

[6.1](https://github.com/camsong/You-Dont-Need-jQuery#6.1)  Basic utilities

- isArray

Determine whether the argument is an array.

```plain text
// jQuery
$.isArray(array);

// Native
Array.isArray(array);
```

- isWindow

Determine whether the argument is a window.

```plain text
// jQuery
$.isWindow(obj);

// Native
function isWindow(obj) {
 return obj !== null && obj !== undefined && obj === obj.window;
}
```

- inArray

Search for a specified value within an array and return its index (or -1 if not found).

```plain text
// jQuery
$.inArray(item, array);

// Native
array.indexOf(item) > -1;

// ES6-way
array.includes(item);
```

- isNumeric

Determine if the argument passed is numerical. Use `typeof` to decide the type or the `type` example for better accuracy.

```plain text
// jQuery
$.isNumeric(item);

// Native
function isNumeric(n) {
 return !isNaN(parseFloat(n)) && isFinite(n);
}
```

- isFunction

Determine if the argument passed is a JavaScript function object.

```plain text
// jQuery
$.isFunction(item);

// Native
function isFunction(item) {
 if (typeof item === 'function') {
 return true;
 }
 var type = Object.prototype.toString.call(item);
 return type === '[object Function]' || type === '[object GeneratorFunction]';
}
```

- isEmptyObject

Check to see if an object is empty (contains no enumerable properties).

```plain text
// jQuery
$.isEmptyObject(obj);

// Native
function isEmptyObject(obj) {
 return Object.keys(obj).length === 0;
}
```

- isPlainObject

Check to see if an object is a plain object (created using “{}” or “new Object”).

```plain text
// jQuery
$.isPlainObject(obj);

// Native
function isPlainObject(obj) {
 if (typeof (obj) !== 'object' || obj.nodeType || obj !== null && obj !== undefined && obj === obj.window) {
 return false;
 }

 if (obj.constructor &&
 !Object.prototype.hasOwnProperty.call(obj.constructor.prototype, 'isPrototypeOf')) {
 return false;
 }

 return true;
}
```

- extend

Merge the contents of two or more objects together into a new object, without modifying either argument. object.assign is part of ES6 API, and you could also use a [polyfill](https://github.com/ljharb/object.assign).

```plain text
// jQuery
$.extend({}, object1, object2);

// Native
Object.assign({}, object1, object2);
```

- trim

Remove the white-space from the beginning and end of a string.

```plain text
// jQuery
$.trim(string);

// Native
string.trim();
```

- map

Translate all items in an array or object to new array of items.

```plain text
// jQuery
$.map(array, (value, index) => {
});

// Native
array.map((value, index) => {
});
```

- each

A generic iterator function, which can be used to seamlessly iterate over both objects and arrays.

```plain text
// jQuery
$.each(array, (index, value) => {
});

// Native
array.forEach((value, index) => {
});
```

- grep

Finds the elements of an array which satisfy a filter function.

```plain text
// jQuery
$.grep(array, (value, index) => {
});

// Native
array.filter((value, index) => {
});
```

- type

Determine the internal JavaScript [Class] of an object.

```plain text
// jQuery
$.type(obj);

// Native
function type(item) {
 const reTypeOf = /(?:^\[object\s(.*?)\]$)/;
 return Object.prototype.toString.call(item)
 .replace(reTypeOf, '$1')
 .toLowerCase();
}
```

- merge

Merge the contents of two arrays together into the first array.

```plain text
// jQuery, doesn't remove duplicate items
$.merge(array1, array2);

// Native, doesn't remove duplicate items
function merge(...args) {
 return [].concat(...args)
}

// ES6-way, doesn't remove duplicate items
array1 = [...array1, ...array2]

// Set version, does remove duplicate items
function merge(...args) {
 return Array.from(new Set([].concat(...args)))
}
```

- now

Return a number representing the current time.

```plain text
// jQuery
$.now();

// Native
Date.now();
```

- proxy

Takes a function and returns a new one that will always have a particular context.

```plain text
// jQuery
$.proxy(fn, context);

// Native
fn.bind(context);
```

+ makeArray

Convert an array-like object into a true JavaScript array.

```plain text
// jQuery
$.makeArray(arrayLike);

// Native
Array.prototype.slice.call(arrayLike);

// ES6-way: Array.from() method
Array.from(arrayLike);

// ES6-way: spread operator
[...arrayLike];
```

-   

[6.2](https://github.com/camsong/You-Dont-Need-jQuery#6.2)  Contains

Check to see if a DOM element is a descendant of another DOM element.

```plain text
// jQuery
$.contains(el, child);

// Native
el !== child && el.contains(child);
```

-   

[6.3](https://github.com/camsong/You-Dont-Need-jQuery#6.3)  Globaleval

Execute some JavaScript code globally.

```plain text
// jQuery
$.globaleval(code);

// Native
function Globaleval(code) {
 const script = document.createElement('script');
 script.text = code;

 document.head.appendChild(script).parentNode.removeChild(script);
}

// Use eval, but context of eval is current, context of $.Globaleval is global.
eval(code);
```

-    

[6.4](https://github.com/camsong/You-Dont-Need-jQuery#6.4)  parse

- parseHTML

Parses a string into an array of DOM nodes.

```plain text
// jQuery
$.parseHTML(htmlString);

// Native
function parseHTML(string) {
 const context = document.implementation.createHTMLDocument();

 // Set the base href for the created document so any parsed elements with URLs
 // are based on the document's URL
 const base = context.createElement('base');
 base.href = document.location.href;
 context.head.appendChild(base);

 context.body.innerHTML = string;
 return context.body.children;
}
```

- 

[6.5](https://github.com/camsong/You-Dont-Need-jQuery#6.4)  exists

-   

exists

Check if an element exists in the DOM

```plain text
// jQuery
if ($('selector').length) {
 // exists
}

// Native
var element = document.getElementById('elementId');
if (typeof(element) != 'undefined' && element != null)
{
 // exists
}
```

[**back to top**](https://github.com/camsong/You-Dont-Need-jQuery#table-of-contents)

## Promises

A promise represents the eventual result of an asynchronous operation. jQuery has its own way to handle promises. Native JavaScript implements a thin and minimal API to handle promises according to the [Promises/A+](http://promises-aplus.github.io/promises-spec/) specification.

-   

[7.1](https://github.com/camsong/You-Dont-Need-jQuery#7.1)  done, fail, always

`done` is called when promise is resolved, `fail` is called when promise is rejected, `always` is called when promise is either resolved or rejected.

```plain text
// jQuery
$promise.done(doneCallback).fail(failCallback).always(alwaysCallback)

// Native
promise.then(doneCallback, failCallback).then(alwaysCallback, alwaysCallback)
```

-   

[7.2](https://github.com/camsong/You-Dont-Need-jQuery#7.2)  when

`when` is used to handle multiple promises. It will resolve when all promises are resolved, and reject if either one is rejected.

```plain text
// jQuery
$.when($promise1, $promise2).done((promise1Result, promise2Result) => {
});

// Native
Promise.all([$promise1, $promise2]).then([promise1Result, promise2Result] => {});
```

-   

[7.3](https://github.com/camsong/You-Dont-Need-jQuery#7.3)  Deferred

Deferred is a way to create promises.

```plain text
// jQuery
function asyncFunc() {
 const defer = new $.Deferred();
 setTimeout(() => {
 if(true) {
 defer.resolve('some_value_computed_asynchronously');
 } else {
 defer.reject('failed');
 }
 }, 1000);

 return defer.promise();
}

// Native
function asyncFunc() {
 return new Promise((resolve, reject) => {
 setTimeout(() => {
 if (true) {
 resolve('some_value_computed_asynchronously');
 } else {
 reject('failed');
 }
 }, 1000);
 });
}

// Deferred way
function defer() {
 const deferred = {};
 const promise = new Promise((resolve, reject) => {
 deferred.resolve = resolve;
 deferred.reject = reject;
 });

 deferred.promise = () => {
 return promise;
 };

 return deferred;
}

function asyncFunc() {
 const defer = defer();
 setTimeout(() => {
 if(true) {
 defer.resolve('some_value_computed_asynchronously');
 } else {
 defer.reject('failed');
 }
 }, 1000);

 return defer.promise();
}
```

[**back to top**](https://github.com/camsong/You-Dont-Need-jQuery#table-of-contents)

## Animation

-  

[8.1](https://github.com/camsong/You-Dont-Need-jQuery#8.1)  Show & Hide

```plain text
// jQuery
$el.show();
$el.hide();

// Native
// More detail about show method, please refer to https://github.com/oneuijs/oui-dom-utils/blob/master/src/index.js#L363
el.style.display = ''|'inline'|'inline-block'|'inline-table'|'block';
el.style.display = 'none';
```

-   

[8.2](https://github.com/camsong/You-Dont-Need-jQuery#8.2)  Toggle

Display or hide the element.

```plain text
// jQuery
$el.toggle();

// Native
if (el.ownerDocument.defaultView.getComputedStyle(el, null).display === 'none') {
 el.style.display = ''|'inline'|'inline-block'|'inline-table'|'block';
} else {
 el.style.display = 'none';
}
```

-  

[8.3](https://github.com/camsong/You-Dont-Need-jQuery#8.3)  FadeIn & FadeOut

```plain text
// jQuery
$el.fadeIn(3000);
$el.fadeOut(3000);

// Native fadeOut
function fadeOut(el, ms) {
 if (ms) {
 el.style.transition = `opacity ${ms} ms`;
 el.addEventListener(
 'transitionend',
 function(event) {
 el.style.display = 'none';
 },
 false
 );
 }
 el.style.opacity = '0';
}

// Native fadeIn
function fadeIn(elem, ms) {
 elem.style.opacity = 0;

 if (ms) {
 let opacity = 0;
 const timer = setInterval(function() {
 opacity += 50 / ms;
 if (opacity >= 1) {
 clearInterval(timer);
 opacity = 1;
 }
 elem.style.opacity = opacity;
 }, 50);
 } else {
 elem.style.opacity = 1;
 }
}
```

-   

[8.4](https://github.com/camsong/You-Dont-Need-jQuery#8.4)  FadeTo

Adjust the opacity of the element.

```plain text
// jQuery
$el.fadeTo('slow',0.15);
// Native
el.style.transition = 'opacity 3s'; // assume 'slow' equals 3 seconds
el.style.opacity = '0.15';
```

-   

[8.5](https://github.com/camsong/You-Dont-Need-jQuery#8.5)  FadeToggle

Display or hide the element by animating their opacity.

```plain text
// jQuery
$el.fadeToggle();

// Native
el.style.transition = 'opacity 3s';
const { opacity } = el.ownerDocument.defaultView.getComputedStyle(el, null);
if (opacity === '1') {
 el.style.opacity = '0';
} else {
 el.style.opacity = '1';
}
```

-  

[8.6](https://github.com/camsong/You-Dont-Need-jQuery#8.6)  SlideUp & SlideDown

```plain text
// jQuery
$el.slideUp();
$el.slideDown();

// Native
const originHeight = '100px';
el.style.transition = 'height 3s';
// slideUp
el.style.height = '0px';
// slideDown
el.style.height = originHeight;
```

-   

[8.7](https://github.com/camsong/You-Dont-Need-jQuery#8.7)  SlideToggle

Display or hide the element with a sliding motion.

```plain text
// jQuery
$el.slideToggle();

// Native
const originHeight = '100px';
el.style.transition = 'height 3s';
const { height } = el.ownerDocument.defaultView.getComputedStyle(el, null);
if (parseInt(height, 10) === 0) {
 el.style.height = originHeight;
} else {
 el.style.height = '0px';
}
```

-   

[8.8](https://github.com/camsong/You-Dont-Need-jQuery#8.8)  Animate

Perform a custom animation of a set of CSS properties.

```plain text
// jQuery
$el.animate({ params }, speed);

// Native
el.style.transition = 'all ' + speed;
Object.keys(params).forEach((key) => {
 el.style[key] = params[key];
});
```

## Alternatives

- [You Might Not Need jQuery](http://youmightnotneedjquery.com/) - Examples of how to do common event, element, ajax etc with plain javascript.
- [npm-dom](http://github.com/npm-dom) and [webmodules](http://github.com/webmodules) - Organizations you can find individual DOM modules on NPM

## Browser Support

<!-- unsupported block: child_database -->

# License

MIT
