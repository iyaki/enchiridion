---
title: "onmount - Safe, reliable, idempotent and testable behaviors for DOM nodes"
notion_id: 6828912d-a206-4c1e-ba73-10eb5871a1b5
notion_url: https://app.notion.com/p/onmount-Safe-reliable-idempotent-and-testable-behaviors-for-DOM-nodes-6828912da2064c1eba7310eb5871a1b5
last_edited: 2023-02-22T18:29:00.000Z
source_url: https://github.com/rstacruz/onmount
tags: ["Framework/Library", "English", "Javascript", "Frontend", "Untried"]
---
# onmount

Run something when a DOM element appears and when it exits.
 No dependencies. Legacy IE compatible. 1kb .min.gz.

## Overview

### Detecting elements

Run something to initialize an element on its first appearance.

```plain text
onmount = require('onmount')

onmount('.push-button', function() {
 $(this).on('click', function() {
 alert('working...')
 })
})
```

> _See: _[_Premise_](https://github.com/rstacruz/onmount/blob/master/docs/premise.md)

### Using with React

If you're looking to use Onmount to mount React components, check out [Remount](https://github.com/rstacruz/remount) instead.

### Polling for changes

Call `$.onmount()` everytime your code changes.

```plain text
$('<button class="push-button">Do something</button>').appendTo('body')

$.onmount()

$('.push-button').click() //=> 'working...'
```

> _See: _[_Idempotency_](https://github.com/rstacruz/onmount/blob/master/docs/idempotency.md)

### jQuery integration

jQuery is optional; use it to poll on popular events.

```plain text
$(document).on('ready show.bs closed.bs load page:change', function() {
 $.onmount()
})
```

> _See: _[_API_](https://github.com/rstacruz/onmount/blob/master/docs/api.md)

### Cleanups

Supply a 2nd function to _onmount()_ to execute something when the node is first detached.

```plain text
$.onmount(
 '.push-button',
 function() {
 /*...*/
 },
 function() {
 alert('button was removed')
 }
)

document.body.innerHTML = ''

$.onmount() //=> 'button was removed'
```

> _See: _[_Cleanups_](https://github.com/rstacruz/onmount/blob/master/docs/cleanup.md)

## What for?

Onmount is a safe, reliable, idempotent, and testable way to attach JavaScript behaviors to DOM nodes. It's great for common websites that are not Single-Page Apps. Read more on its [premise and motivation](https://github.com/rstacruz/onmount/blob/master/docs/premise.md).

[rsjs](https://github.com/rstacruz/rsjs) (Reasonable System for JavaScript Structure) is a great standard that onmount fits perfectly into.

## Usage

Onmount is available via [npm](https://www.npmjs.com/package/onmount) and Bower.

```plain text
npm install onmount
bower install onmount

```

It can be used as a CommonJS module or on its own. It doesn't require jQuery, but if jQuery is found, it'll attach itself to it as `$.onmount`.

```plain text
onmount = require('onmount') // With CommonJS (ie, Browserify)
window.onmount // with no module loaders:
$.onmount // with jQuery
```

## API

> _See: _[_API_](https://github.com/rstacruz/onmount/blob/master/docs/api.md)

## Browser compatibility

All modern browsers and IE8+. For legacy IE, use it with jQuery 1.x.


