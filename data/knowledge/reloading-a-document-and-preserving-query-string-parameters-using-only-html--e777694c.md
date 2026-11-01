---
title: "Reloading a Document (and Preserving Query String Parameters) Using Only HTML"
notion_id: e777694c-bd13-45b5-8ea6-3a63bf5f7b11
notion_url: https://app.notion.com/p/Reloading-a-Document-and-Preserving-Query-String-Parameters-Using-Only-HTML-e777694cbd1345b58ea63a63bf5f7b11
last_edited: 2023-08-29T11:54:00.000Z
source_url: https://blog.jim-nielsen.com/2023/reloading-document-in-html-and-preserve-query-params/
tags: ["Web Development", "Article", "Tutorial", "English"]
---
**tl;dr**: an empty string for your link, e.g. `<a href="">Reload</a>`

The other day I was trying to write some HTML to give the user the ability to reload the document in its exact state by clicking on a link (same functionality as if they hit `CMD` + `R` on their keyboard, or clicked “reload” in the browser UI).

My first attempt was to use a relative reference to the document, e.g.

`<a href=".">Reload</a>`

But then I thought, “I don’t know if this will preserve the existing query parameters in the URL…” Turns out, it doesn’t.

“What about a `<form>`?” I thought.

`<form><button type='submit'>Refresh</button></form>`

Nope. Doesn’t work either.

So I started searching:

“how do you reload an html document using the `<a>` tag and preserve query parameters?”

That gave me a bunch of answers on how to do it _with JavaScript_, so I had to add “without JavaScript” to my query.

But I could not find an answer.

So I turned to ChatGPT, who told me it was not possible (and then recommended I use JavaScript)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

(Funny aside: this is a familiar feeling when searching for answers about how to do stuff on the web. You ask, “How do I do _x_ without using JavaScript?” And all the answers are: “Here’s how to do it with JavaScript”. Or, for anybody who grew up in the jQuery age, “How do I do _x_ without jQuery?” And all the answers were, “Here’s how to do it with jQuery.”)

But I digress.

At that point, it seemed pretty inconceivable to me that there was really no way to reload a document using only HTML that preserved the state of the URL.

So I turned to Mastodon and Twitter for help.

Fortunately, Ryan Florence (who knows a lot about routing on the web) had [the answer](https://twitter.com/ryanflorence/status/1693648034939469827?s=20): use an `href` with an empty string.

`<a href="">Reload</a>`

Sure enough, that worked. If you hover a link like that in the browser, you’ll see the little link preview show the exact same URL as your current document.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Cheers to anybody trying to do this without JavaScript!
