---
title: "Alpine 3.x Tips and Tricks"
notion_id: 575111d2-6fbc-4804-b0b1-11fa5b4f30dc
notion_url: https://app.notion.com/p/Alpine-3-x-Tips-and-Tricks-575111d26fbc4804b0b111fa5b4f30dc
last_edited: 2023-01-13T16:45:00.000Z
source_url: https://ryangjchandler.co.uk/posts/alpine-3-tips-and-tricks
tags: ["Article", "Tutorial", "English", "Frontend", "Javascript", "HTML"]
---
At the time of writing this article, Alpine Day has finished and [Alpine 3.x](https://github.com/alpinejs/alpine-next) has only been out a few hours. Some of the cool things that came with the new release are:

- And much, much more...

I've only had my hands on it for a few hours, just like everybody else, but I've already had a big dive through the source code and here are a couple of things I've found.

## 1. Automatic `init` function calls

The new documentation states that components powered by `Alpine.data` declarations that contain an `init` function will automatically have the `init` function called. This is great and saves you some time since you don't need to hook up the `x-init` directive.

After reading through the source code, it turns out you can actually add an `init` function to any data object and Alpine will still call that function automatically.

```plain text
<div x-data="{
    init() {
        console.log('Here we go!')
    }
}"></div>

```

When this component is initialized, the `init` function will be called automatically and the `console.log` will appear in your DevTools.

## 2. Clean up after yourself with `destroy`

Anybody familiar with PHP will likely know about the `__destruct` magic method. This method is called when an object is garbage collected, allowing you to clean up or free other resources manually.

Alpine 3.x also introduces a similar pattern for components through the use of `destroy` method. This method will be called when the component is removed from the DOM,

A good example might be a carousel or image gallery library that modifies the DOM outside of your component, or isn't handled directly by Alpine.

```plain text
<div x-data="{
    init() {
        carouselLibrary.create();
    },
    destroy() {
        carouselLibrary.delete();
    }
}"></div>

```

Alpine will call the `destroy` method, allowing you to clean up after yourself.

**This feature hasn't been documented yet. Please use with caution.**

## 3. Interact with global stores from external JavaScript

This is something that Spruce supported to, so if you used that with Alpine 2.x you'll be familiar with this one.

Creating a store with `Alpine.store` allows you to access global state in your components using a `$store` property. That same `Alpine.store` method can be used to retrieve a store in your external scripts.

```plain text
Alpine.store('counter', {
    count: 0
})

```

```plain text
// In a different file or area.
Alpine.store('counter').count += 1

```

Calling `Alpine.store` with a single argument will return the `Proxy` instance for that particular store. Awesome, right?

## 4. Independent `x-init` directives

This one is mentioned [in the documentation](https://alpinejs.dev/directives/init#standalone-x-init), but you can declare `x-init` on its own, inside or outside of an Alpine component.

This was a question that got asked _so_ much previously -- people were getting frustrated with adding `x-init` and nothing happening.

```plain text
<div x-data>
    <p x-init="console.log('I am ready!')"></p>
</div>

<img src="..." x-init="doSomeMagicHere()">

```

This is great for elements that don't need any state and just need to call another method or third-party library on initialisation.

## 5. Unfurl / unwrap `Proxy` with `Alpine.raw`

When debugging an issue in a component, 9 times out of 10 we turn to `console.log()`. This is fine in most cases. Other times, people get confused by the appearance of a `Proxy` in their console window.

This is somewhat annoying since expanding the object will give you `[[Target]]` and `[[Handler]]` which can be confusing.

To save you some confusion, you can use `Alpine.raw` inside of your `console.log()` calls. This method will unfurl / unwrap the `Proxy` instance created by Alpine and return the underlying value, i.e. the plain object, array, etc.

```plain text
<div x-data="{ user: { name: 'Ryan' } }" x-init="console.log(user)">
<!-- This produces a `Proxy` in the console -->
</div>

<div x-data="{ user: { name: 'Ryan' } }" x-init="console.log(Alpine.raw(user))">
<!-- This produces the "real" `user` object in the console -->
</div>

```

And that's it from me, at least for now. I'll be working hard over the next couple of weeks trying to find other tips and tricks. In the meantime you can find me on [Twitter](https://twitter.com/ryangjchandler) and [GitHub](https://github.com/ryangjchandler).
