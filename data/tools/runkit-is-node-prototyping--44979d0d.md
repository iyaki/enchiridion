---
title: "RunKit is Node prototyping"
notion_id: 44979d0d-ff39-40ba-99fa-3891650c5daf
notion_url: https://app.notion.com/p/RunKit-is-Node-prototyping-44979d0dff3940ba99fa3891650c5daf
last_edited: 2023-06-22T11:16:00.000Z
source_url: https://runkit.com/home
tags: ["Tool", "Framework/Library", "English", "Programming", "Javascript", "Untried"]
---
[https://static.runkitcdn.com/assets/videos/demo.mp4?v=runkit](https://static.runkitcdn.com/assets/videos/demo.mp4?v=runkit)

## Prototype and explore your ideas

RunKit notebooks completely remove the friction of trying new ideas. With one click you'll have a sandboxed JavaScript environment where you can instantly switch node versions, use every npm module without having to wait to install it, and even visualize your results. No more configuration, just straight to coding.

**connected**

Every version of every package on npm pre-installed. Built in search to help you find it.

**Visualize your data**

From graphs and maps to low level hexadecimal inspectors, you can pick and choose the best way to look at your data after creating it.

**Time Traveling Debugging**

RunKit allows you to rewind your work to a previous point, even filesystem changes are rewound!

**Async Friendly**

Instead of fiddling with chains of callbacks, you can use async/await.

**Download**

All notebooks are just node modules, so you can download them and run them on your own setup with no changes.

**Share**

You can link anyone to your notebook and it serves as a frozen example. Great for bug reports and code samples!

### The Global Library at Your Fingertips

RunKit removes all barriers to finding existing libraries that already do what you need. In fact, you can even require multiple versions of the same package side by side in a document, providing unique debugging opportunities.

[API Diff Example](https://runkit.com/runkit/api-diff-example)

### Next Generation Time Traveling

RunKit takes time traveling debugging further using a revolutionary new technology called CRIU, which allows notebooks to snapshot the entire environment. That means you can rewind subprocesses, the filesystem, really just about anything in your session.

[Learn More About Time Traveling](http://blog.runkit.com/2015/09/10/time-traveling-in-node.js-notebooks.html)

## Instant API.

Create an API without worrying about servers or configuration. Just export a endpoint function and your notebook automatically becomes an HTTPS endpoint, accessible from any app. Great for prototyping iOS and Android backends, or creating microservices.

exports.endpoint = function(req, res) { res.end("Hello, World!"); }

# Terminal

```plain text
$ curl -L https://runkit.io/runkit/hello-world-api/1.0.0
Hello, World!
$
```

Try it yourself with `$ curl -L `[`https://runkit.io/runkit/hello-world-api/1.0.0`](https://runkit.io/runkit/hello-world-api/1.0.0)

[Full Documentation](https://runkit.com/docs/endpoint)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Bring your code to life.

RunKit makes it easy to let your users run the sample code in your blog posts and documentation right on your website.

# index.html

1 2 3 4 5 6 7 8 9 10

```plain text
<!-- Insert this line in your HTML. -->
<script src="https://embed.runkit.com"
         data-element-id="my-element"></script>

<!-- anywhere else on your page -->
<div id="my-element">
    console.log("hello world");
</div>
```

your-great-site.com         console.log("hello world");   Loading…

[Full Documentation](https://runkit.com/docs/embed)

## Case Studies

## [The Fieldbook](https://fieldbook.com/developers)

Fieldbook lets you create a database as easily as a spreadsheet, and uses RunKit Embed to let their users try their API right from their site. Power users can use JavaScript to interact with Fieldbook databases right in the browser!

## [Gitbook](https://www.gitbook.com/)

The RunKit Gitbook plugin lets you embed live examples into your books. Documentation, tutorials, and educational materials can benefit from working examples!
