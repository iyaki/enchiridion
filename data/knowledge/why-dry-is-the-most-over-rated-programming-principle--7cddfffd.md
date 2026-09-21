---
title: "Why DRY is the most over-rated programming principle"
notion_id: 7cddfffd-dfc5-48e2-a921-86237306b944
notion_url: https://app.notion.com/p/Why-DRY-is-the-most-over-rated-programming-principle-7cddfffddfc548e2a92186237306b944
last_edited: 2026-09-21T17:35:00.000Z
source_url: https://gordonc.bearblog.dev/dry-most-over-rated-programming-principle/
tags: ["Article", "Thoughts on Software Blog (Gordon Cassie)", "English", "Programming"]
---
[https://gordonc.bearblog.dev/dry-most-over-rated-programming-principle/](https://gordonc.bearblog.dev/dry-most-over-rated-programming-principle/)

07 Jul, 2022

I figured I'd kick off my new blog with the most click baity thing I could think of. I suspect any developer reading this is aware of the DRY principle because it is just so ubiquitous. If not though, you just need to know that it stands for "Don't Repeat Yourself" and is generally invoked when advising people to not copy and paste snippets of code all over the place and instead consolidate logic into a central place.

DRY was the first programming principle I encountered and probably the only one I was aware of for the first year that I was a developer. It's also probably one of the simplest principles to understand. If you see two things in your code that are the same, maybe they should just be one thing. Hard to argue with that. But, I think that DRY is just like every other principle out there - it has its place, but it's best taken in moderation. And I think that, due to its ubiquity and simplicity, we tend to take DRY too far, far too often.

So without further ado, let's dive into my three criticisms of DRY.

## 1. DRY is misused to eliminate coincidental repetition

Sometimes things happen to be the same, but it's just a coincidence. For instance, consider some python code that requests a pizza from a fictional API.

```plain text
def make_hawaiian_pizza(): payload = { crust: "thin", sauce: "tomato", cheese: "regular", toppings: ["ham", "pineapple"] } requests.post(PIZZA_URL, payload) def make_pepperoni_pizza(): payload = { crust: "thin", sauce: "tomato", cheese: "regular", toppings: ["pepperoni"] } requests.post(PIZZA_URL, payload)
```

Quite a lot of repetition going on in these payloads. Really the only difference between the pizzas is the toppings. One would be very tempted to "DRY it up" and make the following refactor:

```plain text
def make_pizza(toppings): payload = { crust: "thin", sauce: "tomato", cheese: "regular", toppings: toppings } requests.post(PIZZA_URL, payload) def make_pepperoni_pizza(): make_pizza(["pepperoni"]) def make_hawaiian_pizza(): make_pizza(["ham", "pineapple"])
```

The problem is that these two pizzas just happen to have the same crust, sauce and cheese. Had we started out with two pizza types that have different crust/sauce/cheese, we never would have made this refactor. Instead of our code being architected around the concept of how pizzas are made in the abstract, its architecture is tightly coupled to the specific needs of these two pizzas that we happened to be dealing with. The chance that we will be putting this code back the way it was is extremely high.

## 2. DRY creates a presumption of reusability

Imagine we are at a company with a large codebase where multiple product areas want to integrate ordering pizza. Rather than every product writing their own make_pizza() function, why not stick it in a common library that any product can import and call?

So we go down this path and we end up with 5 products each calling make_pizza() with different arrays of arguments for the various types of pizzas they want.

Now along comes some cutting edge product team and they really want to start making pizzas that are half Hawaiian, half pepperoni. The developers on this team are all about DRY and know that there is a great shared pizza function so they go to use it. The only problem is, it can't take split pizza orders. Some modifications will have to be made.

```plain text
# cool_product/pizza.py left_toppings = ["beef"] right_toppings = [] make_pizza(left_toppings, right_toppings) # this will be a very funny pizza # common/make_pizza.py def make_pizza(*args): payload = { crust: "original", sauce: "tomato", cheese: "regular", } if len(args) == 2: payload["toppings_left"] = args[0] payload["toppings_right"] = args[1] else: payload["toppings_left"] = args[0] payload["toppings_right"] = args[0] return requests.post(PIZZA_URL, payload)
```

This works, and it doesn't require changing every existing usage of the API. Hopefully though, you can agree that it is not good™. Having the meaning of the first argument change because you passed an optional second argument is very odd. There are many other ways of doing this refactor but I would assert that any change that doesn't modify the existing calls of make_pizza or make a totally separate function for split topping pizzas (not DRY) will be some level of bad.

You might think that legit reasonable developers would not actually do something like this and would instead go back to the existing invocations and modify them to get a nice solution, but I've seen this happen all over the place. Overzealous use of DRY puts us into a mindset where we are always looking to reuse code, even when it is so obviously taking us down a bad path. We end up with a presumption of reusability when really we should have a presumption of repetition.

## 3. DRY is a gateway drug to unnecessary complexity

If you are a 10x developer you probably at this point have a long list of potential solutions for the problems I've highlighted. You might say that I am making my examples deliberately obtuse to win my points and that there are actually ways I could fix these problems.

To solve my sauce issue, maybe I could use an OOP style and have a PizzaOrderer class that can be subclassed for each pizza type, allowing each type to override sensible sauce/crust defaults. Or maybe I could use a class to represent a Pizza and have methods like add_toppping()/add_topping_left()/add_topping_right() so consumers can add toppings quickly if making a whole pizza but also opt into granularity for split pizzas. There are many other tricks that you could suggest.

All these ideas are great. But remember that the fundamental goal here, is to send a POST request with a single JSON object. That is a very, very simple thing to do. Now we are talking about all kinds of fancy programming stuff to try to solve problems that only exist because we don't want to repeat the same 6 line snippet in a handful of different places because DRY tells us that's bad.

What's happening is that our adherence to DRY is leading us down a garden path to building an unnecessarily complex application that could be written very simply. This occurs far too often. Copying and pasting a few lines of code takes almost zero thought and no time. Find and replace are very good at finding repeating things later if we start to care. As soon as we start the thought process of thinking how to avoid a copy paste and refactor instead, we are losing the complexity battle.

## So what?

Well, obviously I'm not saying we should throw DRY completely out the window. I'm not sure it would actually be possible to write code that "never doesn't repeat itself". But I do think we should tone down knee jerk reactions to PRs that contain several repetitions of a block of code. There are at least a few cases where that might be the exact right thing to do.

I hope you enjoyed the post and please hit me up on Twitter to discuss!

Subscribe to my blog via email or RSS feed.
