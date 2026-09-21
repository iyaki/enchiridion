---
title: "Special interest infodump: restful is a scam, but browsers are cool\""
notion_id: b1e85482-cb9c-4356-80c0-b3f2bafbbee0
notion_url: https://app.notion.com/p/Special-interest-infodump-restful-is-a-scam-but-browsers-are-cool-b1e85482cb9c435680c0b3f2bafbbee0
last_edited: 2023-07-03T11:29:00.000Z
source_url: https://cohost.org/tef/post/1794038-special-interest-inf
tags: ["@tef (cohost)", "English", "Web Development", "REST API", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

here's the deal

- restful is a scam
- (optional) the thesis is confusing: hypermedia as the whatsit of thingy (optional)
- rest is how the web (used to) work
- other things that aren't rest
- an actual restful api: screen scraping

once again, this is a "post your drafts" post, my apologies

- 

anyone who says "this is a restful api" is almost certainly making shit up. absolute scam

- 

people who yell "your thing should be restful" usually mean "do it the way i like doing it"

- 

"the way i like doing it" is often a mix of "use http well" with some other shit

- 

the other shit is always a bunch of weird superstitions and naming conventions, neither of which matter in real actual-rest

- 

in other words, actual-rest isn't about being crud or exposing a filesystem, and isn't really about the naming scheme for urls, or using patch over post

- 

actual rest is about "how does web browser work", and the comes from a thesis about a wholly other thing entirely—defining architectural styles through constraints

- 

the thesis isn't actually that helpful in understanding how rest works, but it's worth covering anyway

### the thesis is confusing: hypermedia as the whatsit of thingy

- 

the thesis is rather light on how http, html, and urls all fit together but it does decompose the design of a "distributed hypermedia system" (a web browser) into a list of constraints

- 

the easy constraints "Client-Server", "Stateless", "Cache", "Layered System" just mean things like "A web browser uses HTTP because load balancers, caches, and proxies don't really care what sort of website you have or what kind of files you serve"

- 

the "Code-On-Demand" constraint that everyone ignores is "websites can have javascript in them (optional)"

-  

the hardest to understand constraint, "uniform interface", is actually four constraints in a trench coat: identification, self description, representations, hypermedia

- 

"identification of resources" means "different things have different urls" and "different things have different content types"

- 

"self-descriptive messages" means "web servers return caching headers" and "load balancers don't have to make sticky sessions" and "everything speaks http", but it can also mean "web servers return html, and html has links and forms"

- 

"manipulation of resources through representations" means things like "things like html can describe their own interfaces in terms of links and forms" but also "things at a url can be described in different forms, with different content types"

- 

"hypermedia as the engine of application state" is possibly the most confusing thing, and it means several things at once: "the user uses links and forms to navigate the website", "each part of your website has a different url", and "if you opened the same link in two different tabs, you'd be in the same part of the website, twice over".

- 

in other words, links and forms (hypermedia) are the way in which users change the url (state) of the web browser (application)

- 

when you put this last constraint together with all the rest, you start to define a style, which gets named "representational state transfer"

- 

the name comes from how a "representation of state" (urls) is transferred from server to client (webpages with links and forms), and then how this state is transferred back to the server (when the user clicks a link)

### rest is how the web (used to) work

- 

it isn't about clients and services, it's about web browsers and web servers. a distributed hypermedia system.

-  

it isn't about apis, it's about how the transport (http), the interface description (html),

and application state (urls) all fit together

- 

it's about how using something like http lets you use load balancers, caches, or proxies near invisibly

- 

it's about how using something like (html) lets a service can expose an interface, in a platform neutral way, over the network.

- 

it's about storing application state inside a url, and manipulating the url through things like links and forms, rather than keeping state client or server-side.

- 

it's about why the server can change all the urls and everything still works, as long as the links are updated.

- 

rest is about opening "next" in two new tabs, and getting the same page back twice, over getting back pages 2 and 3.

- 

it's about is about clicking `next` on a webpage, as opposed to a user typing `/collection?offset=80` to to go the next page

- 

and it's also about using standard means for doing so, such that any web browser can access and interact with any web server too

- 

it in other words, rest is about how you can use your web browser to interact with webmail, without your browser knowing anything about email

- 

the closest thing to actual rest that people understand is the "richardson maturity model", which breaks down "being rest" into three easy steps.

1. play nice with load balancers: different things have different urls
2. play nice with caches: get can be cached, post can't be cached
3. play nice with browsers. actually serve html. now drawl the fucking owl.

- if there isn't a browser, there isn't any rest, but just having a browser in and of itself does not make things restful

### other things that aren't rest

- 

to understand how the web works, it's worth looking at things that aren't the web

-  

take gopher, for example. compared to http, gopher comes up a little lacking:

- starts with a similar request/response like protocol
- but middleware can't really cache results without hardcoding exceptions
- and there's fixed list of built in types, rather than a content-type field
- only the menu type and directory type can link to other files
- one special command: search action. any other special commands have to be made as protocol extensions
- it's a browser, of sorts, but it really isn't restful
- if the web worked this way, you'd struggle to write an email service for any gopher client
-  

rest doesn't mean "filesystem" either, take 9P for example

- 9p is remote filesystem protocol, but instead of passing filenames around, it uses file-descriptor
- a client authenticates, and opens a given path as a root directory for the session
- clients can call `walk` on a file-descriptor for a directory to get a list of files
- only directories can link to other files, in terms of the protocol
- there's no typed files, everything is just a blob
-  

9p just isn't very web like, even if you can browse a filesystem:

- there's no special commands: everything must be implemented in terms of file operations, like say, echoing `"now"` into `/service/next_reboot`
- there's no hypermedia, but there are similar idioms. a plan9 service will expose some file that has a filename inside. for example `/net/mux` returns `1` and so the client decides to read from `/net/1/ctl`
- although any 9p client can connect to any 9p service, the way in which services offer other features varies wildly from one service to the next
- you can expose an email service over the filesystem, but you end up having to write an email client atop, almost like a browser plugin
- and it's worth noting that plan9's 9p isn't that easy to write middleware for—a proxy, cache, or load balancer, has to keep track of how file-descriptors match to files
- but making it more http like wouldn't be enough
-  

the web is more than just "a file system over http" too

- sure, you get caching, load balancers, and proxies, but that's only the beginning
- the big problem is this: where's the hypermedia, and what's the engine of application state
- with webdav, with filesystems, the client has to know the exact urls of the things it wants, and on the web, you'd use the name of a link instead
- any special operation still has to be implemented in terms of reading/writing files, and every application built atop will do something different.
- by comparison, a web session can carry state from one request to the next inside the url, because the client navigates by names
-  

and json-rpc over http isn't much restful usually, either.

- sure enough the caching, load balancer, and proxy bit come for free with using http, as your api doesn't need to be filesystem like to benefit from middleware, but that's about it for restful-ness
- and sometimes there is a little bit of state passing, much like the plan9 thing of returning file names
- if you have a pagination api, and it involves passing a token from one reply into the next request, then you're getting pretty close to using urls to model application state
- the problem is? instead of html, they use a schema, and that schema has to be downloaded in advance
- clients can't read from a schema the same way they read from html,
- if a web browser worked the same way, you'd be downloading electron apps for every website (welp)
-  

it's almost worth a summary list

- rest is about how html, http, and urls combine to form a distributed hypermedia system
- it's about a system where you can cache, proxy, or load balance requests without new client code
- it's about a system where application state is encoded in urls
- it's about a system where application state is changed by clicking on links or forms
- 

so after all that, you might be asking, why the fuck would anyone want to make a restful api?

- 

you might even be asking, could a restful api even exist?

- 

the answer is yes, restful apis do exist. and yes, they looks like a web browser

### an actual restful api: screen scraping

-  

here is what an actual restful client might looks like, for a given restful api:

```plain text
import browser

api = browser.Open("http://localhost:1729/service")

users = api.Click("admin").Click("users")

result = users.Submit("create", {"name": "jeff"})

if result.Error != nil  { // etc }

```

- 

yes, it's screen scraping

-  

yes, i am being 100% serious here, restful apis are screen scraping:

- it uses http, so fulfils the whole set of caching, load balancer, proxy constraints
- in theory, the library could support javascript, so code-on-demand doesn't have to be optional
- the uniform interface is there:
- the same client can be used to access any website or service speaking html
- the client only cares about one url, the starting point (or endpoint or bookmark)
- every other url is navigated to by clicking on links and submitting forms
- although the underlying protocol is making GET or POST requests, the client doesn't care
- and the representation of application state (urls) is being transferred
- 

and if your api doesn't look like this, well, it probably isn't restful.

- 

in theory, it could speak something other than html, but there's a catch, the html replacement has to be generic enough to describe different types of interfaces. sometimes people will claim they've made a restful system, but instead of using one language (html) to describe every interface, each interface is a unique type of resource unto itself.

- 

so now we know it's possible, it's a good time to ask why anyone would do anything like this

- 

maybe i'll cover that in another post, but for now i'd hope the idea of "you can use the same client for lots of different services" would be pretty cool, along with "you can carry state from one request to the next without explicitly asking the client to do so"

- 

but maybe the coolest bit? with rest, by using html, we can express a far richer set of operations than http contains, in an interoperable way with existing middleware and services, without having every client know things up front and in advance
