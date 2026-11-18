---
title: "How to store your app's entire state in the url"
notion_id: d679e37c-3bda-4448-a7d5-8289d4c23b34
notion_url: https://app.notion.com/p/How-to-store-your-app-s-entire-state-in-the-url-d679e37c3bda4448a7d58289d4c23b34
last_edited: 2023-01-21T19:23:00.000Z
source_url: https://www.scottantipa.com/store-app-state-in-urls
tags: ["Article", "English", "Web Development"]
---
I'm working on [a flowchart editor](https://www.knotend.com/) that runs in the browser, and I wanted a way for people to use it without having to sign in, or store any data on our server. I wanted to give them control over their data and to be able to store it locally to open and edit later. And also easily share it with other people. It's easy to do this by supporting file upload/download, but I wanted something simpler, like the ability to share by sending a url. I also didn't want to store anything on the backend (at least for the free tier).

I decied to encode the entire application state as a Base64 encoded string in the hashmark of the url. For example, a url would look like (note its truncated since they are very long):

`knotend.com/g/a#N4IgzgpgTglghgGxgLwnARgiAxA9lAWxAC5QA7X...`

Everything after the ``/g/a#`` is a stringified version of a json object that contains all the information about the flowchart. It gets stringified, then compressed, then Base64 encoded. I update the url on every graph edit, so copying the graph state is as simple as copying the url in your browser bar.

Here's the pseudo code for creating the url, and then later reading it:

```javascript
const stateString = JSON.stringify(appState); // appState is a json object
const compressed = compress(stateString);
const encoded = Base64.encode(compressed);
// Push that `encoded` string to the url
// ... Later, on page load or on undo/redo we read the url and
// do the following
const decoded = Base64.decode(encoded); // same encoded as above, but read from url
const uncompressed = uncompress(decoded);
const newState = JSON.parse(uncompressed);
// Now load your application with the newState
```

There are several options for implementing the compress/uncompress functions, such as [lz-stirng](https://github.com/pieroxy/lz-string) or [pako](https://github.com/nodeca/pako).

Since I update it on every graph edit, I get something major for free -- undo/redo. The browser's history stack becomes my undo/redo functionality. The user can hit the browser back/forward buttons, or Command-Z,Command-Shift-Z which I map to history pop and push. This is a major win for something which is a free product that I wanted to ship quickly.

Another great benefit is that these urls can be embedded. That means the user can put their graph on any ewb page that supports embedding. I see people typically do this with wikis like Notion, which means you can share with a team without anyone needing an account on my site.

You can see how it works by checking out [knotend](https://www.knotend.com/), the keyboard-centric flowchart editor that I'mw working on.

### Prior work and thank yous.

I'm not the first one to take this approach. I've seen atleast [mermaidjs](https://mermaid.live/) do this before, and I'm sure there are others.

Thank you to [this comment by redleader55](https://news.ycombinator.com/item?id=34313074) on hacker news for pointing out that using window.location.hashmark is better for storing longer urls since some browsers will truncate the url when sending it over http. But that this doesn't apply to the hashmark, which stays client side.
