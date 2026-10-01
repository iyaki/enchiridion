---
title: "Locality of Behavior"
notion_id: c73f0341-327a-4dd4-a903-49156de2f48e
notion_url: https://app.notion.com/p/Locality-of-Behavior-c73f0341327a4dd4a90349156de2f48e
last_edited: 2023-02-09T01:23:00.000Z
source_url: https://www.eloquentarchitecture.com/locality-of-behavior/
tags: ["Article", "English", "System Design / Software Architecture", "Programming"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The documentation for [htmx](https://htmx.org/) refers to something called "[Locality of Behavior](https://htmx.org/essays/locality-of-behaviour/)".

> The behaviour of a unit of code should be as obvious as possible by looking only at that unit of code

> The LoB principle is a simple prescriptive formulation of the quoted statement from

[Richard Gabriel](https://www.dreamsongs.com/)

The author was talking about how htmx demonstrates this principle well when compared to [jQuery](https://jquery.com/). And indeed it does, to an extent. Most modern view libraries are all great examples of this, to varying degrees. [React](https://reactjs.org/), [Vue](https://vuejs.org/), [Stimulus](https://stimulus.hotwired.dev/), [Alpine](https://alpinejs.dev/), [htmx](https://htmx.org/), and many others, each have their events, actions, and targets defined in the html to different extents, therefore making it easier to know that something special is happening here.

But, is simply placing these things next to each other enough to make a unit of code "obvious"? I think there's more to the discussion.

**Obvious **(ob·vi·ous)
Easily perceived or understood; apparent.

## There's something missing

The notion of pairing these concerns together, on its own, doesn't guarantee that it is obvious by looking only at that unit code. Each of these libraries comes with the possibility of obscuring what's really going on behind the curtain. We've all seen code like this.

React.js example

Stimulus.js example

...and so on. These examples could be doing anything. They are all doing something, of that we can be sure. And while this a step in the right direction, still far from what many would consider to be obvious.

It's much easier to get away with this when you have the rest of the code in the same file like you do with React and Vue's Single-File Components. Scrolling's not that hard.

## Naming things _is_ hard

These libraries all provide a way to declare what is happening, when it occurs, and where the underlying logic can be found. All you, as the author, have to do is to name things well enough so that the reader might be given some insight without going source diving.

This is a minor gripe. Too many of these libraries include examples on their front page like the ones above that don't convey anything about anything. Sure, it's just an example. I get it. But this is where beginners go to learn.

Stimulus, on the other hand, provides a great example for beginners on their front page. 👏

```plain text
<div data-controller="hello">
    <input data-hello-target="name" type="text">

    <button data-action="click->hello#greet">
        Greet
    </button>

    <span data-hello-target="output">
    </span>
</div>
```

And again and again throughout their documentation.

```plain text
<form data-controller="search">
    <input data-action="input->search#loadResults">

    <ul data-search-target="results"></ul>
</form>
```

They get it - pairing actions and events with your markup doesn't guarantee transparency, insight, or understanding by itself. For that you need more.

## Stimulus isn't the only one who gets it

Take a look at this controller found within [Laravel Fortify](https://laravel.com/docs/9.x/fortify) and take note of how long it takes you to fully wrap your head around what it is responsible for.

```plain text
class PasswordController extends Controller
{
    public function update(Request $request, UpdatesUserPasswords $updater)
    {
        $updater->update($request->user(), $request->all());

        return app(PasswordUpdateResponse::class);
    }
}
```

You have to admit that this code satisfies the LoB principle pretty damn well. There's more to see, but you can probably guess what it does. Laravel gets it too.

## Conclusion

One can only achieve behavior transparency to a certain extent without also taking a minute to name things well. Hopefully, bringing this missing part of the equation into the discussion will increase the efficacy of the LoB prescription.

> The function of good software is to make the complex appear to be simple. ~ Grady Booch
