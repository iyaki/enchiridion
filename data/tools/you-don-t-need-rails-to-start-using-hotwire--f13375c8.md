---
title: "You Don't Need Rails to Start Using Hotwire"
notion_id: f13375c8-9712-4756-ba61-a429334fd9c0
notion_url: https://app.notion.com/p/You-Don-t-Need-Rails-to-Start-Using-Hotwire-f13375c897124756ba61a429334fd9c0
last_edited: 2023-08-15T11:15:00.000Z
source_url: https://www.akshaykhot.com/using-hotwire-without-rails/
tags: ["English", "Frontend", "Javascript", "HTML", "Untried", "Tutorial", "Framework/Library", "Akshay's Blog"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

> Update:

So far on this blog, we've learned about [Hotwire](https://www.akshaykhot.com/introduction-to-hotwire/) and built a [hotwired to-do list](https://www.akshaykhot.com/building-to-do-list-using-hotwire-and-stimulus/) powered by Rails. We've also seen how you can [iteratively build and progressively enhance an application using Hotwire](https://www.akshaykhot.com/progressive-application-development-with-hotwire/). All these projects used Hotwire with Ruby on Rails.

However, you don't really have to use Rails (or Ruby) to get most of the benefits of Hotwire. Most static websites can just drop-in the [Turbo](https://github.com/hotwired/turbo?ref=akshaykhot.com) library to behave like responsive single-page applications, without incurring any of the costs and complexities associated with the SPA frameworks. And if you have an existing app written in PHP, Go, Rust, or even Java, you can start using Hotwire, right now.

**All you have to do is follow certain **[**conventions**](https://rubyonrails.org/doctrine?ref=akshaykhot.com#convention-over-configuration)**.**

This article is divided into two parts:

1. First one shows how you can use the first two Hotwire techniques (Turbo Drive and Turbo Frames) in a simple static website, to fetch and update entire web pages or parts of the page, without fully-reloading the browser. We'll use a simple static site to demo this.
2. Second part shows how to tweak your existing back-end code to send Turbo Streams to update multiple parts on your website dynamically, in response to form submissions. Since Turbo Streams work with form submissions, I'll use Sinatra to demo this.

Here're the topics we'll cover in this article.

- [How to set up a simple website](https://www.akshaykhot.com/using-hotwire-without-rails/#set-up-a-simple-website)
- [Using an HTTP server to serve static pages](https://www.akshaykhot.com/using-hotwire-without-rails/#using-an-http-server-to-serve-static-pages)
- [How to install Turbo](https://www.akshaykhot.com/using-hotwire-without-rails/#how-to-install-turbo)
- [Faster navigation with Turbo Drive](https://www.akshaykhot.com/using-hotwire-without-rails/#faster-navigation-with-turbo-drive)
- [Dynamic page updates with Turbo Frames](https://www.akshaykhot.com/using-hotwire-without-rails/#dynamic-page-updates-with-turbo-frames)
- [How Turbo Frames Work?](https://www.akshaykhot.com/using-hotwire-without-rails/#how-turbo-frames-work)
- [Step One: Wrap Component in Turbo Frame](https://www.akshaykhot.com/using-hotwire-without-rails/#step-one-wrap-the-component-in-a-turbo-frame)
- [Step Two: Wrap Response in a Turbo Frame with same ID](https://www.akshaykhot.com/using-hotwire-without-rails/#step-two-wrap-the-response-in-a-turbo-frame-with-same-id)
- [Working with Turbo Streams](https://www.akshaykhot.com/using-hotwire-without-rails/#working-with-turbo-streams)
- [Target Multiple Elements with Turbo Streams](https://www.akshaykhot.com/using-hotwire-without-rails/#target-multiple-elements-with-turbo-streams)
- [Step One: Use Turbo HTTP Header](https://www.akshaykhot.com/using-hotwire-without-rails/#step-one-use-turbo-http-header)
- [Step Two: Send Turbo Streams in Response](https://www.akshaykhot.com/using-hotwire-without-rails/#step-two-send-turbo-streams-in-response)
- [Where to go from here](https://www.akshaykhot.com/using-hotwire-without-rails/#where-to-go-from-here)

💡

For those of you who don't enjoy reading long, rambling posts that try to explain every nook and cranny of the topic at hand, I've uploaded [the finished project](https://github.com/akshayKhot/wireframe?ref=akshaykhot.com) on my GitHub account, which also contains the instructions to get started. Just clone the [repository](https://github.com/akshayKhot/wireframe?ref=akshaykhot.com), run `npm run launch`, and you're good to go.

**Prerequisites:** If you don't know what [Hotwire](https://hotwired.dev/?ref=akshaykhot.com) is, I suggest you check out the following article, written by yours truly. It briefly introduces Hotwire and explains the problems it solves.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Set up a Simple Website

> In this section, we'll set up a simple website that serves static files. Both Turbo Frames and Turbo Drive don't need a backend server, so a simple static website should be simple and barebones enough to explain the basic concepts.

Create a brand new directory for your website and `cd` into it. I'll call mine `wireframe`.

```plain text
➜ mkdir wireframe
➜ cd wireframe
```

Run the [`npm init`](https://docs.npmjs.com/cli/v9/commands/npm-init?ref=akshaykhot.com) command to set up a new project. It will ask you a bunch of questions and then create a `package.json` in the current directory.

```plain text
➜ wireframe npm init

package name: (wireframe)
version: (1.0.0)
description: Using Hotwire without Ruby on Rails
entry point: (index.js)
test command:
git repository:
keywords: hotwire, turbo
author: AK
license: (ISC) MIT
```

Here's the resulting `package.json` file

```plain text
{
  "name": "wireframe",
  "version": "1.0.0",
  "description": "Using Hotwire without Ruby on Rails",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "keywords": [
    "hotwire",
    "turbo"
  ],
  "author": "AK",
  "license": "MIT"
}

```

Now open the directory in your favorite editor. We're ready to start coding.

**Add a simple HTML file**

Let's create a new folder named `public` with an `index.html` HTML file in it. The HTML file will have the following content. Feel free to copy and paste.

To make our website look a bit nice, I am using [SimpleCSS](https://simplecss.org/?ref=akshaykhot.com), a simple, classless CSS framework. It's so simple that I don't even have to explain it. Basically, it makes semantic HTML look good, that's it. No classes required.

If you open the `index.html` file directly in the browser, it should look like this. Pretty neat, right?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

index.html in browser

Notice that the nav bar says `File` , instead of `HTTP`, because chrome is directly serving the file. To make it look like a real website served with the HTTP protocol, we'll need an HTTP server that will serve the HTML.

> But why?

[this StackOverflow question](https://stackoverflow.com/questions/40204913/difference-between-localhost-and-opening-html-file?ref=akshaykhot.com)

## Using an HTTP Server to Serve Static Pages

I am going to use a simple static HTTP server called [`http-server`](https://www.npmjs.com/package/http-server?ref=akshaykhot.com) which is more than sufficient for our needs.

> http-server

Did you know that you can run the package without installing it first, using the `npx` command?

Run the following command from the `wireframe` directory.

Without any arguments, the above command serves the `index.html` file in the `public` directory when you visit [`http://127.0.0.1:8080`](http://127.0.0.1:8080/?ref=akshaykhot.com) or [`localhost:8080`](http://127.0.0.1:8080/?ref=akshaykhot.com). This is why I'd asked you to create a `public/index.html` file for your project.

You can also add the above command as a `script` in the `package.json` file. This will allow you to launch the website using the `npm run launch` command.

Now that your server is up and running, visit the [`http://127.0.0.1:8080`](http://127.0.0.1:8080/?ref=akshaykhot.com) or [`http://localhost:8080`](http://127.0.0.1:8080/?ref=akshaykhot.com) URL in the browser.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

HTTP Browser

Now that our website is up and running, we're ready to install Turbo.

## How to Install Turbo

We are going to use the pre-compiled, optimized NPM package from [skypack.dev](https://www.skypack.dev/?ref=akshaykhot.com) using the `<script>` tag, just like it's 2007. For other installation methods, check out the [Installing Turbo](https://turbo.hotwired.dev/handbook/installing?ref=akshaykhot.com) documentation.

**Step 1: **Add the following `script` tag just above the `<title>` tag in your HTML.

**Step 2:** There's no step 2. ;)

If you're curious about how the above snippet works, I highly recommend you read the [MDN documentation on JavaScript Modules](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Modules?ref=akshaykhot.com).

That's it. Our little website is using Turbo.

To verify, reload the browser, open the DevTools window, go to the `Console` tab, and type `Turbo` in it. If it doesn't throw an error, you're good to go.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Installing Turbo

Now that we've successfully installed Turbo, we're ready to use it. We'll start with the first big component in Turbo, called [Turbo Drive](https://turbo.hotwired.dev/handbook/drive?ref=akshaykhot.com).

## Faster Navigation with Turbo Drive

> It just works out-of-box.

The best thing about Turbo Drive is that you get it for free. Yes, you heard that right. You don't have to do anything to get the benefits of Turbo Drive.

**But how does it work?**

When you click a link or submit a form (to the same domain), Turbo Drive does the following:

1. Prevent the browser from following the link,
2. Change the browser URL using the [History API](https://developer.mozilla.org/en-US/docs/Web/API/History?ref=akshaykhot.com),
3. Request the new page using a [`fetch`](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API?ref=akshaykhot.com) request
4. Render the response HTML by replacing the current `<body>` element with the response and merging the `<head>`.

The JavaScript `window` and `document` objects as well as the `<html>` element persist from one rendering to the next.

The same goes for an HTML form. **Turbo Drive intercepts and converts Form submissions into fetch requests. Then it follows the redirect and renders the HTML response.**

As a result, your browser doesn’t have to reload, and the website feels much faster and more responsive, just like a single-page application.

### Let's Add a Contact Page

To see how Turbo Drive works, we need to set up another page on our website that we'll add a link to.

We've already added the links to the _**Contact **_and _**About**_ pages when we wrote the initial HTML, so let's go ahead and add a _**Contact**_ page. To keep it really simple, I'll just copy and paste the `index.html` page, changing the filename and a little content to make it unique.

> This is only for demo. Your back-end framework or static-site generator uses a templating system to extract all the duplicate HTML.

You should see the following page when you go to `/contact` page.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

contact page

Now go ahead and click back and forth between the _**Home**_ and _**Contact **_links.

**No?** let me give you a hint. Comment out the `<script>` tags that load the Turbo library on both pages.

_Don't forget to comment it on both pages, okay?_

Now clear the cache and hard reload the browser by pressing and holding the reload button while the DevTools window is open. This removes the `Turbo` library you loaded earlier from the website.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

clear cache and hard reload

Now click between _**Home **_and _**Contact **_links.

See something different?

I'm sure you must have figured it out by now. When you navigate between multiple pages, the browser is doing a full reload. You can see this by noticing the earth icon in the tab, which spins for a quick second when you go to a different page.

Now uncomment the script from both pages and reload the browser again. No need to clear the cache this time. The website should fetch the Turbo library without any issues.

Go ahead, and click between the pages.

Get it?

**The earth is not spinning anymore! **The pages are updating without a full browser reload. How cool is that?

This is the power of Turbo Drive. Without any extra effort on your part, your website has instantly become more responsive and dynamic.

To recap, here's what happened when you clicked on the _**Contact **_link.

1. Turbo intercepted that click, prevented the browser from following it, made a `fetch` request to get the content of the _**Contact**_ page.
2. Upon receiving the HTTP response, Turbo then replaced the current body of the web page with the body of the result.
3. Additionally, it merged the contents of the `<head>` tag if there's new stuff here, like new `<meta>` tags or new JavaScript. In our case, there wasn't any new stuff, so it left the head tag as it is.

### Displaying a Progress Bar

You can improve the perceived responsiveness of your website by displaying a progress bar while Turbo fetches the new page. Simply add the following CSS that targets the `.turbo-progress-bar` element.

It might be hard to see it, as the navigation is so fast. You can [throttle the network](https://css-tricks.com/throttling-the-network/?ref=akshaykhot.com) to _Slow 3G_ to see the progress bar.

That's the essence of how Turbo Drive works. **You get a bulk of the benefits of modern single-page applications, with a fraction of the complexity associated with the complicated SPA frameworks.**

You can check out the [documentation](https://turbo.hotwired.dev/handbook/drive?ref=akshaykhot.com) to learn more about the advanced features of Turbo Drive. But for now, let's move on to Turbo Frames.

## Dynamic Page Updates with Turbo Frames

We've seen how Turbo Drive can make your website responsive by replacing the current body element with the response body.

For most websites (that are not web applications), this is absolutely enough to get the majority of performance boost and to give that SPA-like feel without any added complexity.

However, sometimes you have a website that only needs to update a small section on the page while leaving the whole page intact.

Imagine a blog with comments enabled (just like this blog you're reading this post on). When someone adds a comment to my post, I only want to update the comments section, without updating the whole blog post. Replacing the whole body doesn't make sense here.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This is the appeal behind single-page applications, where most of the page remains as it is, and only sections on the page are updated independently.

**What if you could just send the specific HTML that changed, i.e. the comments section, without touching the rest of the page, i.e. the blog post? The response would be much smaller, and the rest of the HTML could be easily cached, making the application even more responsive.**

For this, we need to bring out the next weapon in our arsenal: **Turbo Frames. **Turbo Frames allows us to do the exactly same thing.

The only way Turbo Frames differ from SPA JavaScript frameworks is this: **the part of the page that's updated is retrieved from the response HTML, instead of making an API call to retrieve the JSON response.**

## What are Turbo Frames?

Turbo Frames allow you to dynamically update sections on the page in response to some action, such as clicking a link or submitting a form.

I will demonstrate Turbo Frames by building a simple gallery on our home page. Here's how the resulting page will look.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Clicking on the _**Next **_link dynamically updates the image without reloading the browser.

Let's get started. First, update the `index.html` page's `<main>` tag with the following HTML. Just copy and paste it, I'll explain what's going on soon. All I did is added an image and the link, and wrapped them in a `<turbo-frame>` tag.

Then create a new folder named `gallery` in the `public` directory. It contains three HTML files named `forest.html`, `mountains.html`, and `ocean.html`. Here's their content.

I've also added three new images in the `public/images` directory that I grabbed from the [w3schools](https://www.w3schools.com/css/css_image_gallery.asp?ref=akshaykhot.com) website (or you can also find them in this project's[ GitHub repository](https://github.com/akshayKhot/wireframe/tree/main/public/gallery?ref=akshaykhot.com)).

That's all needed for our image gallery. Clicking on the _**Next **_link updates the picture without reloading the page. Rest of the page, like the `<header>` content doesn't change.

**Wait, what just happened?**

## How Turbo Frames Work?

You may have noticed the new HTML element named `<turbo-frame>`. It's a [custom HTML element](https://developer.mozilla.org/en-US/docs/Web/Web_Components/Using_custom_elements?ref=akshaykhot.com) provided by Turbo. This element allows you to divide your website into independent components that need to be changed independently. Let's see how they work, in three simple steps:

### Step one: Wrap the component in a turbo frame.

You wrap the section on the page that you want to update in response to link clicks or form submissions inside a `<turbo-frame>` element and give it a sensible `id`.

In our example, I only want to update the image and the link below it, so I wrapped it inside a turbo frame with the ID `gallery`.

### Step two: Wrap the response in a turbo frame with same ID.

Any HTML that you want to send from the server, you wrap it in a turbo frame and give it the same ID as the original component. This is how Turbo figures out which frame to update on the page.

For example, here is the forest page that contains the new image and the new link, both wrapped in a turbo frame with the ID `gallery`.

**Step three: **[**There's no step three ;)**](https://www.youtube.com/watch?v=2iyMf3tlKpU&ref=akshaykhot.com)

When you click the link and the response from the server arrives, Turbo finds the `<turbo-frame>` element with the matching ID, and replaces the current `<turbo-frame>` element with the one from the response.

In the above example, when you click the Next button, the `<turbo-frame id="gallery">` element on the `index.html` page is replaced with the matching turbo frame element that arrives from the response. That's how the image + link is replaced with a new image + link.

You can have multiple `<turbo-frame>` elements on the page. Each one should have its own, unique ID. That's how Turbo knows which frame to replace when the response arrives from the server.

However, keep in mind that **at any given request-response cycle, only one Turbo Frame will be swapped. **If you need to update multiple components on the site, you'll have to use Turbo Streams, which we'll explore next.

## Working with Turbo Streams

> Turbo Streams

```plain text
<turbo-stream>
```

I'll demo the Turbo Streams with a different example, since you need to handle form submissions for Turbo Streams, and our static site can't do that. So I'll use [Sinatra](https://sinatrarb.com/intro?ref=akshaykhot.com), an elegant Ruby web framework.

### Create a Sinatra Project

First, let's install the Sinatra framework and Puma web server using the `gem install` command.

Create a new directory for the project. I'll call mine `wirestream`. Navigate into it and open it in your favorite browser.

Create a new Ruby script called `app.rb` that adds a route for the home page.

Now run that script just like any other Ruby script:

Sinatra is up and running and serving your web application at `localhost:8000`.

### Render a Template

Before we use Turbo Streams, let's set up a proper HTML template, just like Rails.

Passing the `:index` symbol tells Sinatra to look for a `index.erb` template in the `views` directory. Just copy + paste the following HTML.

Let's restart Sinatra and reload the page. As you can see, we've added a simple newsletter form and a subscriber list.

Whenever someone enters their name + email and hits "Subscribe", we want to add their name to the "Subscriber List", and also show a notification-like header at the top. Like this:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Stream Update

Let's see how you'd accomplish this using Turbo Streams.

## Target Multiple Elements with Turbo Streams

The very first thing we'll need is a new route to handle form submissions. Let's add a new `/subscribe` route that handles a form POST submission.

All it's doing is accepting a POST request to `/subscribe`, getting the entered name and saving it to an instance variable (just like Rails) and return a `views/subscribe.erb` template. Since our form already submits to `/subscribe`, this route will handle the form submission.

Next, let's add a simple `subscribe.erb` template under the `views` directory. For now, it says that new user has subscribed.

Restart the app, and submit the form after entering the name and email. You should see the following page.

Now, instead of rendering a separate `/subscribe` page, we want to send a Turbo Stream that update multiple elements at the same time. For this, let's first add the Turbo library to our app, just like we did earlier. Under the `<head>` tag, add the following code.

Now that we've Turbo, we'll accomplish the multiple dynamic Turbo Stream updates in two simple steps:

### Step One: Use Turbo HTTP Header

For the Turbo JavaScript library to identify the Turbo HTTP response from the server, the response needs to have [a special HTTP header that indicates the Content Type](https://github.com/hotwired/turbo/blob/0291998ba465bf4660b860e40744dfa24818fbee/src/core/streams/stream_message.ts?ref=akshaykhot.com#L5), as follows:

As long as the response has this header, the Turbo library will treat it as a Turbo Stream response and treat it accordingly.

Adding a new header in Sinatra is very simple.

### Step Two: Send Turbo Streams in Response

Replace the existing content of the `subscribe.erb` template with the following code, which contains two separate Turbo Stream responses containing following action attributes:

1. **replace: **to swap the existing content of an element with the ID `subscriber-notification` with the content inside the `<template>` tag.
2. **append: **to append the contents inside the `<template>` tag to the element with the ID `subscriber-list`.

If you notice the HTML that you copy+pasted earlier in the `index.erb` template, it contains these two elements:

So, in its essence, we are instructing Turbo to replace the empty notification div element with another div containing the subscriber name, and to append the name of the new subscriber to the existing list.

To learn more about various Turbo Stream actions, check out its [documentation](https://turbo.hotwired.dev/handbook/streams?ref=akshaykhot.com).

That's it. You're all set. Restart the app, fill out the form, and hit submit. You should see the green notification as well as the name of the new subscriber.

## Where to go from here

This wraps up our exploration into using Hotwire (without Rails) on a static website as well as a web app not using Rails.

Over the last few months, I've written multiple articles on Hotwire, as I've found it an excellent way to build single-page web apps without incurring the complexities associated with so-called modern SPA frameworks like React or Vue.

Check them out:

If you enjoy writing Ruby, PHP, Python, Go, Rust, or any other back-end language (including JavaScript via node), and keep your front-end as simple as possible while still retaining the interactivity, I highly recommend you check it out.

That's a wrap. I hope you liked this article and you learned something new. If you're new to the blog, check out the [full archive](https://www.akshaykhot.com/page/2/) to see all the posts I've written so far or the [ favorites page](https://www.akshaykhot.com/favorites/) for the most popular articles on this blog.

As always, if you have any questions or feedback, didn't understand something, or found a mistake, please leave a comment below or [send me an email](mailto:akshay.khot@hey.com?ref=akshays-blog). I look forward to hearing from you.

If you'd like to receive future articles directly in your email, please [subscribe to my blog](https://www.akshaykhot.com/#/portal/signup). If you're already a subscriber, thank you.
