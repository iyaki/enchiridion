---
title: "Vanilla JS"
notion_id: df3f18a6-9728-41ea-bdfa-09a5171cc7ac
notion_url: https://app.notion.com/p/Vanilla-JS-df3f18a6972841eabdfa09a5171cc7ac
last_edited: 2023-02-05T01:25:00.000Z
source_url: http://vanilla-js.com/
tags: ["Website", "Framework/Library", "Article", "English", "Web Development", "Javascript", "Frontend", "Reflection"]
---
Vanilla JS is a fast, lightweight, cross-platform frameworkfor building incredible, powerful JavaScript applications.

## Introduction

The Vanilla JS team maintains every byte of code in the framework and works hard each day to make sure it is small and intuitive. Who's using Vanilla JS? Glad you asked! Here are a few:

- Facebook
- Google
- YouTube
- Yahoo
- Wikipedia
- Windows Live
- Twitter
- Amazon
- LinkedIn
- MSN
- eBay
- Microsoft
- Tumblr
- Apple
- Pinterest
- PayPal
- Reddit
- Netflix
- Stack Overflow

In fact, Vanilla JS is already used on more websites than jQuery, Prototype JS, MooTools, YUI, and Google Web Toolkit - combined.

## Download

Ready to try Vanilla JS? Choose exactly what you need!

Core Functionality

Final size: 0 bytes uncompressed, 25 bytes gzipped. Show human-readable sizes

## Testimonials

## Getting Started

The Vanilla JS team takes pride in the fact that it is the most lightweight framework available anywhere; using our production-quality deployment strategy, your users' browsers will have Vanilla JS loaded into memory before it even requests your site.

To use Vanilla JS, just put the following code anywhere in your application's HTML:

1. <script src="path/to/vanilla.js"></script>

When you're ready to move your application to a production deployment, switch to the much faster method:

That's right - no code at all. Vanilla JS is so popular that browsers have been automatically loading it for over a decade.

## Speed Comparison

Here are a few examples of just how fast Vanilla JS really is:

### Retrieve DOM element by ID

Codeops / sec Vanilla JSdocument.getElementById('test-table');12,137,211 Dojodojo.byId('test-table');5,443,343 Prototype JS$('test-table')2,940,734 Ext JSdelete Ext.elCache['test-table']; Ext.get('test-table');997,562 jQuery$jq('#test-table');350,557 YUIYAHOO.util.Dom.get('test-table');326,534 MooToolsdocument.id('test-table');78,802

### Retrieve DOM elements by tag name

Codeops / sec Vanilla JSdocument.getElementsByTagName("span");8,280,893 Prototype JSPrototype.Selector.select('span', document);62,872 YUIYAHOO.util.Dom.getElementsBy(function(){return true;},'span');48,545 Ext JSExt.query('span');46,915 jQuery$jq('span');19,449 Dojodojo.query('span');10,335 MooToolsSlick.search(document, 'span', new Elements);5,457

## Code Examples

Here are some examples of common tasks in Vanilla JS and other frameworks:

### Fade an element out and then remove it

Vanilla JSvar s = document.getElementById('thing').style; s.opacity = 1; (function fade(){(s.opacity-=.1)<0?s.display="none":setTimeout(fade,40)})(); jQuery<script src="//ajax.googleapis.com/ajax/libs/jquery/1/jquery.min.js"></script> <script> $('#thing').fadeOut(); </script>

### Make an AJAX call

Vanilla JSvar r = new XMLHttpRequest(); r.open("POST", "path/to/api", true); r.onreadystatechange = function () { if (r.readyState != 4 || r.status != 200) return; alert("Success: " + r.responseText); }; r.send("banana=yellow"); jQuery<script src="//ajax.googleapis.com/ajax/libs/jquery/1/jquery.min.js"></script> <script> $.ajax({ type: 'POST', url: "path/to/api", data: "banana=yellow", success: function (data) { alert("Success: " + data); }, }); </script>

## Further Reading

For more information about Vanilla JS:

- check out the Vanilla JS documentation
- read some books on Vanilla JS
- or try out one of the many Vanilla JS plugins.

When powering your applications with Vanilla JS, feel free to use this handy button!
