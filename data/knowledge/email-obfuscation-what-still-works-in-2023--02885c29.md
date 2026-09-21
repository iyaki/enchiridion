---
title: "Email obfuscation: What still works in 2023?"
notion_id: 02885c29-c4c8-4375-9a4f-e4722387ed93
notion_url: https://app.notion.com/p/Email-obfuscation-What-still-works-in-2023-02885c29c4c843759a4fe4722387ed93
last_edited: 2023-11-28T19:03:00.000Z
source_url: https://spencermortensen.com/articles/email-obfuscation/
tags: ["Spencer Mortensen", "English", "Programming", "Web Development", "Email", "Article"]
---
- 1 [Clickable link](https://spencermortensen.com/articles/email-obfuscation/#link) 
- 1.1 [URL encoding](https://spencermortensen.com/articles/email-obfuscation/#link-url): blocked 100% of spam
- 1.2 [HTML entities](https://spencermortensen.com/articles/email-obfuscation/#link-entity): blocked 50% of spam
- 1.3 [No protection](https://spencermortensen.com/articles/email-obfuscation/#link-none): blocked 17% of spam
- 1.4 [Concatenation](https://spencermortensen.com/articles/email-obfuscation/#link-concatenation): _still testing_ JS
- 1.5 [Rot18](https://spencermortensen.com/articles/email-obfuscation/#link-rot18): _still testing_ JS
- 1.6 [Xor](https://spencermortensen.com/articles/email-obfuscation/#link-xor): _still testing_ JS
- 1.7 [Append host](https://spencermortensen.com/articles/email-obfuscation/#link-append): _still testing_ JS
- 1.8 [User interaction](https://spencermortensen.com/articles/email-obfuscation/#link-events): _still testing_ JS
- 2 [Plain text](https://spencermortensen.com/articles/email-obfuscation/#text) 
- 2.1 [Display none](https://spencermortensen.com/articles/email-obfuscation/#text-display): blocked 100% of spam CSS
- 2.2 [HTML comment](https://spencermortensen.com/articles/email-obfuscation/#text-comment): blocked 83% of spam
- 2.3 [HTML entities](https://spencermortensen.com/articles/email-obfuscation/#text-entity): blocked 67% of spam
- 2.4 [No protection](https://spencermortensen.com/articles/email-obfuscation/#text-none): blocked 0% of spam
- 2.5 [Concatenation](https://spencermortensen.com/articles/email-obfuscation/#text-concatenation): _still testing_ JS
- 2.6 [Rot18](https://spencermortensen.com/articles/email-obfuscation/#text-rot18): _still testing_ JS
- 2.7 [Xor](https://spencermortensen.com/articles/email-obfuscation/#text-xor): _still testing_ JS
- 2.8 [Append host](https://spencermortensen.com/articles/email-obfuscation/#text-append): _still testing_ JS
- 2.9 [Symbol substition](https://spencermortensen.com/articles/email-obfuscation/#text-substitution): blocked 100% of spam Issues
- 2.10 [CSS content](https://spencermortensen.com/articles/email-obfuscation/#text-content): blocked 100% of spam CSS Issues
- 2.11 [Text direction](https://spencermortensen.com/articles/email-obfuscation/#text-direction): blocked 100% of spam CSS Issues
- 2.12 [Image](https://spencermortensen.com/articles/email-obfuscation/#text-image): blocked 100% of spam Issues
- 3 [Methodology](https://spencermortensen.com/articles/email-obfuscation/#methodology)

## 1 Clickable link

These techniques protect a clickable link that will open the user’s mail client (e.g. [email](mailto:email@example.com)). The link text is expected to be generic; if the _email address_ appears in the link text, then the email address is additionally exposed as plain text, and you’ll need to layer on the [plain-text](https://spencermortensen.com/articles/email-obfuscation/#text) obfuscation techniques.

### 1.1 URL encoding

Blocked 100% of spam

`<a href="mailto:%65%6d%61%69%6c%40%65%78%61%6d%70%6c%65%2e%63%6f%6d">email</a>`

This is based on a small sample size: just six bots that were observed over a one-year period. Of those bots, _half_ were able to decode HTML entitites, and URL-decoding is just as easy, so caution is advised.

### 1.2 HTML entities

Blocked 50% of spam

`<a href="&#109;&#97;&#105;&#108;&#116;&#111;&#58;&#101;&#109;&#97;&#105;&#108;&#64;&#101;&#120;&#97;&#109;&#112;&#108;&#101;&#46;&#99;&#111;&#109;">email</a>`

### 1.3 No protection

Blocked 17% of spam

`<a href="mailto:email@example.com">email</a>`

Surprisingly, the unprotected email address appears to have blocked a spam email. Either that message wasn’t received, or an extra message was sent to one of the protected email addresses.

### 1.4 Concatenation

_Still testing_ JS

`<script>document.write('<a href="mailto:'+'e'+'m'+'a'+'i'+'l'+'@'+'e'+'x'+'a'+'m'+'p'+'l'+'e'+'.'+'c'+'o'+'m'+'">email</a>');</script>`

This is convenient because it has no external dependencies, but it is easily broken, even by bots that don’t interpret JavaScript.

### 1.5 Rot18

_Still testing_ JS

`<head>
	<script src="js/`[`email.js`](https://spencermortensen.com/articles/email-obfuscation/src/link-rot18.js)`" defer></script>
</head>` `<a class="email" href="znvygb:rznvy@rknzcyr.pbz">email</a>`

This has an external dependency and can _still_ be broken by bots that don’t interpret JavaScript. At the very least, you should change the lengths of the rotation cycles (see the source code).

### 1.6 Xor

_Still testing_ JS

`<head>
	<script src="js/`[`email.js`](https://spencermortensen.com/articles/email-obfuscation/src/link-xor.js)`" defer></script>
</head>` `<a class="email" href="7179757d7854716c75796478713a777b79" rel="nofollow, noindex">email</a>`

This can only be defeated by a bot with full JavaScript support.

### 1.7 Append host

_Still testing_ JS

`<head>
	<script src="js/`[`email.js`](https://spencermortensen.com/articles/email-obfuscation/src/link-append.js)`" defer></script>
</head>` `<a class="email" href="mailto:email">email</a>`

### 1.8 User interaction

_Still testing_ JS

`<head>
	<script src="js/`[`email.js`](https://spencermortensen.com/articles/email-obfuscation/src/link-events.js)`" defer></script>
</head>` `<a class="email" href="mailto:email">email</a>`

## 2 Plain text

### 2.1 Display none

Blocked 100% of spam CSS

`span.email b {
	display: none;
}` `<span class="email">email@example<b>.example</b>.com</span>`

### 2.2 HTML comment

Blocked 83% of spam

`email@example<!--.example-->.com`

### 2.3 HTML entities

Blocked 67% of spam

`&#101;&#109;&#97;&#105;&#108;&#64;&#101;&#120;&#97;&#109;&#112;&#108;&#101;&#46;&#99;&#111;&#109;`

### 2.4 No protection

Blocked 0% of spam

`email@example.com`

### 2.5 Concatenation

_Still testing_ JS

`<script>document.write('e'+'m'+'a'+'i'+'l'+'@'+'e'+'x'+'a'+'m'+'p'+'l'+'e'+'.'+'c'+'o'+'m');</script>`

This is convenient because it has no external dependencies, but it is easily broken, even by bots that don’t interpret JavaScript.

### 2.6 Rot18

_Still testing_ JS

`<head>
	<script src="js/`[`email.js`](https://spencermortensen.com/articles/email-obfuscation/src/text-rot18.js)`" defer></script>
</head>` `<span class="email">rznvy@rknzcyr.pbz</span>`

This can be broken even by bots that don’t interpret JavaScript. At the very least, you should change the lengths of the rotation cycles (see the source code).

### 2.7 Xor

_Still testing_ JS

`<head>
	<script src="js/`[`email.js`](https://spencermortensen.com/articles/email-obfuscation/src/text-xor.js)`" defer></script>
</head>` `<span class="email">7179757d7854716c75796478713a777b79</span>`

### 2.8 Append host

_Still testing_ JS

`<head>
	<script src="js/`[`email.js`](https://spencermortensen.com/articles/email-obfuscation/src/text-append.js)`" defer></script>
</head>` `<span class="email">email</span>`

### 2.9 Symbol substitution

Blocked 100% of spam Issues

`email AT example DOT com`

**Breaks usability.** The user has to extract the email address from the surrounding text and then undo each symbol substitution before they can send their email.

### 2.10 CSS content

Blocked 100% of spam CSS Issues

`span.email::after {
	content: "email@example.com";
}` `<span class="email"></span>`

**Breaks usability.** The email address can be seen, but not selected. The user is forced to type out the full email address by hand.

### 2.11 Text direction

Blocked 100% of spam CSS Issues

`span.email {
	unicode-bidi: bidi-override;
	direction: rtl;
}` `<span class="email">moc.elpmaxe@liame</span>`

**Breaks usability.** The email address can be copied, but the text is reversed—which is confusing, and then frustrating. The user is forced to type out the full email address by hand.

### 2.12 Image

Blocked 100% of spam Issues

`<img src="images/email.jpg" width="216" height="18" alt="email address">`

**Breaks usability.** The user is forced to type out the full email address by hand.

## 3 Methodology

I set up a live example for each obfuscation technique, using a different email address for each technique.
