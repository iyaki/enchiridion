---
title: "Global Variables and States: Why So Much Hate?"
notion_id: 884e837a-87bf-4b53-81ef-a0ee8afaa696
notion_url: https://app.notion.com/p/Global-Variables-and-States-Why-So-Much-Hate-884e837a87bf4b5381efa0ee8afaa696
last_edited: 2022-12-19T14:15:00.000Z
source_url: https://thevaluable.dev/global-variable-explained/
tags: ["Article", "The Valuable Dev", "English", "Programming"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Global variables. Global states. These words induce fear and pain in every developer’s heart who had the unfortunate experience to deal with them.

Did you already fight against applications behaving unexpectedly, without knowing exactly why, like a poor knight would try to kill an Hydra with too many heads to deal with?

Did you end up in the infinite loop of tries and errors, _guessing_ 90% of the time what was happening?

This could be the annoying consequences of globals: hidden variables changing their states in unknown places, depending on things you don’t yet understand.

Do you like the feeling to be totally in the dark when you try to modify an application?

Of course, you don’t. Fortunately, I have some candles for you:

- First, I will describe what we, most of the time, call globals. This name is lacking some precision, that’s why we need to clarify it.
- Then, let see together why globals are bad for your codebase.
- Thirdly, I explain how to shrink the scope of globals to make them local.
- Finally, I speak about encapsulation in general, and why the crusade against globals is only part of a bigger problem.

I hope this article will brush over everything you need to know about global states. If you think it’s not everything, that I forgot a lot of points and therefore you hate me and you don’t want to see me ever again, please leave a comment to complete the knowledge provided by this article. It would be nice for me, for my readers and for everybody who unexpectedly ended up on this page.

Are you ready, dear reader, on your horse, to understand what our enemy is? Let’s go find these globals, and let them taste the cutting edges of our swords!

## What’s a State?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Let’s begin by the basics, to be sure we, developers, understand each other.

A state is the definition of a system or an entity. States can be found in real life:

- When your computer is off, the state of the computer is `off`.
- When a cup of tea is hot, the state of the cup of tea is `hot`.

In software development, some construct (like variables) can have states. For instance, the string `"hello"` or the integer `11` are not considered as states. They are _values_. They become a state when they are attached to variables and stocked in memory.

```plain text
<?php

echo "hello"; // No state here!
$variable = "hello"; // The variable $lala has the state 'hello'.

```

We can distinguish two sorts of state:

First, the **mutable states**: they can change at run time (during the execution of your application) after their initialization, at any point in time.

```plain text
<?php

$variable = "hello"; // Initialisation of the variable.
$variable = "bye"; // The state of the variable $lala can be changed at runtime.

```

Second, the **immutable states**: they can’t change them at run time. You assign to your variable a first state and this value will never change afterward. Constants are the colloquial examples for immutable states:

```plain text
<?php

define("GREETING", "hello"); // Constant definition.
echo GREETING;
GREETING = "bye"; // This line will produce an error!

```

Now, let’s listen to this hypothetical conversation between Dave and Harry, your colleague developers:

- Dave! You create _global variables_ everywhere! It’s impossible to change them without breaking everything! I will burn your chair!
- You’re wrong Harry! My _global states_ are beautiful! They are my life! My masterpieces! I love my _globals_ so much!

Most of the time, developers will call `global states`, `global variables` or `globals` what they should call instead `global mutable states`. That is, a state which can be modified (`mutable`) in the biggest [scope](https://users.cs.cf.ac.uk/Dave.Marshall/PERL/node52.html) you can find: the whole application.

When a variable doesn’t have the whole application as scope, we speak about `local variables`, or `locals`. They are variables which exist in a defined scope, smaller than the application’s scope.

```plain text
<?php

namespace App\Ecommerce;

$global = "I'm a mutable global variable!"; // global variable

class Shipment
{
    public $warehouse; // local variable existing in the whole class

    public function __construct()
    {
        $info = "You're creating a shipment object!"; // local variable bound to the constructor scope
        echo $info;
    }
}

class Product
{
    public function __construct()
    {
        global $global;
        $global = "I change the state now 'cause I can!";
        echo "You're creating a product object!"; // no state here
    }
}


```

At that point, you might think the following: how handy to have variables I can access and change everywhere! I can share states from one part of the application to another! No need to pass them through functions and write so many lines of code! Glory to the _global mutable state_!

If you think that, I urge you to keep reading.

## Are Global Variables Worse than Plague and Cholera?

### The Biggest Mental Model of All

Here’s a fact: you will have an easier way to build an accurate mental model of your application if it contains only locals with small and defined scopes, without any global.

Why is that?

Imagine you have a big application with global variables. Each time you need to change something, you need to:

- Remember that these mutable global states exist.
- Try to guess if they might affect the scope you need to change.

You won’t care, most of the time, about the local variables which are in different scopes that the scope you’re modifying, but you will always have to let some space in your tired brain for global mutable states, whatever you do, since they can potentially _affect all scopes_.

On top of that, your global mutable states can _change anywhere_ in your application. Most of the time, you will have to guess their states. It means looking in the whole application, trying to guess the value of your globals in the scope you modify.

That’s not all. If you need to change the state of your globals, you won’t have any clue what scope will be affected. Does another class, method or function will behave unexpectedly because of the change? Good luck to find out.

> If n objects all know about each others, then a change to just one object can result in the other n-1 objects needing changes.

**The Pragmatic Programmer**

In short, you will couple every class, method and function which use these global states together. Don’t forget: [dependencies](https://thevaluable.dev/cohesion-coupling-guide-examples/) bring a lot of [complexity](https://thevaluable.dev/kiss-principle-explained/#the-dependency-problem).

Does it afraid you? It should. **Small defined scopes** are very useful: you don’t have to put your whole application in your brains, but only the scopes you’re working on.

Humans are not good to take a lot of information into consideration at once. When we try to do so, our mental energy will go down, our focus as well, and we will create bugs and nonsense.

That’s exactly why reasoning in the global scope of your application is a pain.

### Global Name Collisions

Another thing to consider: using third party libraries. Let’s imagine you want to use this super cool library which color each letter of your outputs randomly with a nice blinking effect. The dream of every developer!

If this third party library use globals as well, with the same names that your own global variables, you will experience the pleasure of name collisions. Your application will crash and you will wonder why, possibly for a while:

- First, you will need to find out that your library use global variables.
- Second, you will need to figure out what global variables are used at run time: yours or the library one? Difficult to know, they have the same name!
- Third, since you can’t change the library itself, you will need to rename your own global mutable variables. If it’s used everywhere in your application, it will bring you burning tears.

At each step, expect to lose some hairs from your rage and desperation. Your hairdresser will be out of work quickly. Nobody wants that.

Some of you might remember Javascript libraries like Mootools, Underscore and jQuery always colliding which each other, if you would not wrap them in smaller scopes than the global one. Ah, the infamous global jQuery `$` object!

### Testing Will be Your New Nightmare

If you’re still not convinced, let’s look at the problem from the unit test point of view: how the heck do you write tests when you have global variables?

Since your tests might change your globals, you will never be sure what are their states in any test.

Your tests should be isolated from each other, global states will tie them together.

Did you ever have tests succeeding in isolation, but failing when you were executing the whole test suit? No? I did. Pain is always vivid when I think about it.

### Concurrency Issues

Mutable global states can bring many problems if you need to do some concurrency. It makes sense: if you change the state of your globals on multiple threads, expect to take a big [race condition](https://en.wikipedia.org/wiki/Race_condition#Example) in your face.

If you’re a PHP developer, you won’t really care about that, except if you’re using some library which allows you to do concurrency.

However, when you will [learn a new programming language](https://thevaluable.dev/how-to-learn-a-programming-language/) where you can do easily some concurrency, I hope you’ll remember my prose.

## Avoiding Global Mutable States

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Even though global mutable states can bring a lot of problems, it sometime feels that it’s difficult to avoid.

Let’s take a REST API: endpoints will receive some HTTP requests with parameters and send back a response.

These HTTP parameters sent to the server could be needed in multiple layers of your application. It’s very tempting at that point to make these parameters global when you receive the HTTP request, modifying them along the way before sending the response back. Add concurrency for each request on top, and you will have a good recipe for disasters.

Global mutable states can as well be backed directly in the language implementation. For example, PHP has [superglobals](https://www.php.net/manual/en/language.variables.superglobals.php).

Whenever your globals come from, how to manage them? How to refactor the application of Dave, your colleague developer, who created globals everywhere because he didn’t read anything about development in twenty years?

### Function Arguments

The simplest way to avoid globals all together is to simply pass your variables using function arguments.

Let’s take a simple example:

```plain text
<?php

namespace App;

use Router\HttpRequest;
use App\Product\ProductData;
use App\Exceptions;

class ProductController
{
    public function createAction(HttpRequest $httpReq)
    {
        $productData = $httpReq->get("productData");

        if (!$this->productModel->validateProduct($productData)) {
            return ValidationException(sprintf("The product %d is not valid", $productData["id"]));
        }

        $product = $this->productModel->createProduct($productData);
    }
}

class Product
{
    public function createProduct(array $productData): Product
    {
        $productData["name"] = "SuperProduct".$productData["name"]; // This is not what you should do; I talk about it later in the article.

        try {
            $product = $this->productDao->find($productData["id"]);
            return product;
        } catch (NotFoundException $e) {
            $product = $this->productDao->save($productData);
            return $product;
        }
    }
}

class ProductDao
{
    private $db;

    public function find(int $id): array
    {
        return $this->db->find(['product' => $id]);
    }

    public function save(array $productData): array
    {
        return $this->db->saveProduct($productData);
    }
}

```

As you can see, the `$productData` array from the controller (via HTTP request) goes through different layer:

1. The controller receives the HTTP request.
2. The parameters are passed to the model.
3. The parameters are saved in the application’s database.

We could have made this array of parameters global, when we extract it from the HTTP request. It seems easier: no need to pass the data in four different function.

However, passing these parameters as argument to your functions will:

- Make clear that these functions use the `$productData` array.
- Make clear what function use what parameter. You can see that `ProductDao::find` only needs the `$id` from the `$productData` array, not everything.

Using globals would make the code less understandable and couple the methods to each other, which is a very high price for almost no benefit.

You can already see Dave, your colleague developer, objecting: “What about functions which have three or more arguments already? If you need to add even more, the complexity of the function will rise! And what about variables, objects or whatever construct you need almost everywhere? You pass them to every function in your application?”

These questions are legitimate, dear reader. As a [good developer](https://thevaluable.dev/better-web-developer/), you should explain to Dave, with your famous communication skills, the following:

“Dave, if your functions have too many arguments, the problem might be the functions themselves. Chances are that they do too many things, that they are _responsible_ for too many things. Did you consider splitting them into multiple functions?”

Feeling as an orator on Athene’s Acropolis, you continue:

“It’s a problem if you need variables in many scopes, and we will speak about that soon. However, if you really need them, passing them via function argument is not a problem, is it? You will have to type them, yes, but we are developers, it’s our job to type code.”

It can seem more complex to have more function arguments (arguably, it is), but again, the advantages outweigh the disadvantages: better to be as clear as possible, instead of using hidden global mutable states.

### Context Objects

Context objects are basically objects which contains data defined by some context. Most of the time, the data will be stored in a key / pair data construct, like an associative array in PHP, for example.

It’s a dumb object, that is, an object which has no behavior, only data. It’s similar to a [value object](https://de.wikipedia.org/wiki/Value_Object).

The context object can replace any global mutable state. To illustrate, consider the last code example above. Instead of passing the data of the request through our layers, we could use an object which encapsulate this data.

The context would be the request itself: another request, another context, another set of data. The context object is then passed to every method which needs this data.

You might ask: that’s super cool and all, but what are the advantages?

- The data is encapsulated in an object. The goal will be, most of the time, to make the data immutable, that is, you won’t be able to change the state, the value of the object’s data after initialization.
- It’s clear that the context needs the context object’s data, since it’s passed to every function (or method) which needs it.
- It solves the concurrency problem: if each request has its own context object, you can safely write or read them on their own thread.

However, everything has a cost in development. The context object can be harmful:

- Looking at the function arguments, you won’t know what data is in the context object.
- You can put anything in your context object. Careful not to put too much, like an entire user session for example, or even a big part of the data of your application. You can end up with things like `$context->getSession()->getUser()->getProfil()->getUsername()`. Violating the [Law of Demeter](https://en.wikipedia.org/wiki/Law_of_Demeter) and bringing insane complexity will be your curse.
- Larger the context object, trickier it is to know what data is used, in what scope.

At the end, I would avoid using context objects as much as I can. As seen above, it can bring a lot of doubts. Making the data immutable is a big plus, but its disadvantages need to be taken into account. If you use it, make sure that is stays pretty small and pass them in small and well-defined scope.

In case you don’t have any clue before run time how many states will be passed to your functions (parameters from an HTTP request for example), a context object can be useful. That’s why some framework use it: think about the `Request` object in Symfony, for example.

## Dependency injection

Another good alternative of global mutable states is to inject what you need in your object directly at their creation. This is the definition of dependency injection: a set of techniques to inject objects in your components (classes).

### Why dependency injection?

The goal of injecting your dependency is, again, to restrict the use of your variables, objects and whatever construct you want to use, in a delimited scope. When you have dependencies which are injected and, therefore, can only act in the scope of an object, it will be easier to know in what context they’re used, why and how. Good bye, headaches and frustrations!

It separates two important phases of your application lifetime:

1. Building your application objects and inject their dependencies.
2. Use your objects to achieve the functionality, task and whatnot you need to perform.

It makes things clearer than instantiating everything at random places, or even worse, using global objects everywhere.

Many frameworks use dependency injection, sometimes in a complex way, with configuration files and a Dependency Injection Container (DIC). It doesn’t need to be fancy and complex, though. You can simply create your dependencies on the same layer and inject them on the layer below.

In Golang, for example, I don’t know anybody who’s using a DIC. You just create your dependencies in the main code file (`main.go`), then you pass them to the layer below. You can instantiate everything in a different package as well, to make clear that the “dependency injection phase"™ should only happen on this precise layer.

Package scopes make things easier in Golang than in PHP, where DICs are heavily used in every framework I know, including Symfony and Laravel.

### Constructor Injection vs Setter Injection

There are two ways to inject your dependencies: via constructor or setters. I would advise you to use constructor injection as much as you can:

- If you want to know the dependencies of a class, you just have to find the constructor. You don’t have to search disparate setter methods throughout the whole class.
- Setting your dependencies during instantiation give you the assurance it’s safe to use your object.

Let’s stop a bit on this last point: it’s called enforcing your invariant. By instantiating your object and inject its dependencies, you know that, whatever your object needs, it’s correctly set.

If you use setters, how do you know that your dependencies are set when you want to use your object? You can go up the stack, trying to find if your setters was called, but I’m sure you’ll be happier if you can avoid it.

## Global States and Encapsulation

At the end, the only difference between local and global states is their scopes. Local states have limited scopes, global states have the whole application as scope.

However, you can stumble upon the same problems you have with global states using local states. Why is that?

### You Said Encapsulation?

At the end, using global states will violate encapsulation, like you can violate encapsulation with local states.

Let’s begin by the beginning. What’s the definition of encapsulation according to [Wikipedia](https://en.wikipedia.org/wiki/Encapsulation_%28computer_programming%29)?

> A language mechanism for restricting direct access to some of the object’s components.

Restricting accesses? Why?

Well, as we saw before, reasoning in the boundaries of a local scope is way easier than reasoning in the global scope.

Global mutable states, by definition, are available everywhere, which is the _contrary_ of encapsulation! No access restriction whatsoever.

### Growing Scope and Leaking States

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Let’s imagine together a state which is in its own little scope, a nice and shiny local state.

Unfortunately, as time goes by, as the application grows, this local state begins to be passed, as function argument, all over the application. Your local is used now in many scopes. Its direct access is authorized in all these scopes.

At that point, it will be difficult to know the precise state of your local without looking in every scope it exists, where it could be potentially modified. A problem we saw with global mutable states already.

Let’s take an example: an [Anemic Domain Model](https://thevaluable.dev/anemic-domain-model/) can enlarge the scope of your mutable models.

The Anemic Domain Model basically strips apart data and behavior of your domain objects in two different groups: `models` (objects with only data), and `services` (objects with only behaviors).

Most of the time, these `models` will be used everywhere in `services`. Therefore, for some generic model, chances are that their scopes will grow, and grow, and grow. You won’t have any clue what `model` is used in what context, their states will change, the same problems will jump at your face like angry Facehungers.

The message I want to convey here is important: it’s not because you avoid using global mutable states that you can relax, cocktail in one hand, typing with the other, enjoying life and your fabulous code. First, because you won’t code [very efficiently with alcohol](https://res.cloudinary.com/practicaldev/image/fetch/s--CWGA3Ax4--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://imgs.xkcd.com/comics/ballmer_peak.png), and second, because the real danger is Walter White the lack of encapsulation, that is, mutable states which begin to leak into multiple layers of your application.

If you have to refactor these leaky states, try to contain them in defined scope. Exactly like you would try to contain global mutable states.

How to define these scopes? Follow the business model of your company. Your application should be the mirror of a business, therefore the scopes should include states and behaviors which mimic the business model.

For example, for an e-commerce, having a scope which include everything related to a warehouse could be a good idea. It could be, but it could be as well a huge mistake; every business is different.

There is no magic formula here, except this one: get as much knowledge as you can about the business you work for. You will then be able to encapsulate the data and behavior of your application correctly, in a way which makes sense.

To know more about this topic, I encourage you to read what [Domain Driven Design](https://de.wikipedia.org/wiki/Domain-driven_Design) is about.

### The Power Of Copying States

Copying your states without modifying them directly can be a good solution in many cases. Let’s go back to our `Product` example above, specifically on this method:

```plain text
<?php

class Product
{
    public function createProduct(array $productData): Product
    {
        $productData["name"] = "SuperProduct".$productData["name"]; // This is not what you should do; I talk about it later in the article.

        try {
            $product = $this->productDao->find($productData["id"]);
            return product;
        } catch (NotFoundException $e) {
            $product = $this->productDao->save($productData);
            return $product;
        }
    }
}

```

It would be better for the array `$productData` to stay immutable. If you modify directly its state, and then pass it to different functions, you won’t be able to know what state this array has after a while.

In this example, the value of `name` is only modified once, which is not a big deal. However, what will happen when your colleague developers will see your code? They might think it’s okay to change directly the states of mutable variable. In this context, expect a big mess of unknown states.

When we write code, we write as well guidelines for everybody who will follow us. This is an important thing to consider.

A better solution would be:

```plain text
<?php

class Product
{
    public function createProduct(array $productData): Product
    {
        // Since $productData is passed to other veriable, it has to be immutable.
        $name = "SuperProduct".$productData["name"];

        try {
            $product = $this->productDao->find($productData["id"]);
            return product;
        } catch (NotFoundException $e) {
            $product = $this->productDao->save($name, $productData);
            return $product;
        }
    }
}

```

You make clear that the product name is not the same as the original product name from the `$productData` array. You make clear that the state is modified. If you need to pass this `$productData` to any other method, you will know that it will _always_ contains the original data from your HTTP request.

You could even isolate this change of state in a private method, to make it even clearer: “I’m changing this state now, be careful”.

## What About Global Immutable States?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Using global immutable states is pretty common in any application. A constant, for example, has a global immutable state.

Is it safe to use?

- You don’t have to worry about the actual state of your global; it will be exactly the same state all over your application.
- As a consequence, you only need to look at the first assignation (preferably following the declaration) of your immutable variables to know their states.

However, they are not totally harmless either. You have to be careful to make clear what part of the application is using what global immutable states.

For example, a constant `ShipmentDelay` will be hopefully used only where the shipment logic is implemented.

If Dave, your colleague developer, begin to use `ShipmentDelay` to any delay which has nothing to do with the shipments, your global will be used in situation which doesn’t make sense. Sounds stupid to you? I saw a lot of developers doing this kind of weird thing, on the name of the holy [DRY principle](https://thevaluable.dev/dry-principle-explained/).

If you have global immutable states used in many places, all over the application, in a way which makes sense, it could be the sign that behaviors which should be together are spread apart. Refactoring is the way to go in that case: bringing back what belongs together in a class, a set of class, a package or whatever construct you see fit.

Therefore, when using constants or other global immutable states, keep your two eyes open and make sure they don’t spread everywhere.

It’s not because you can access global immutable constructs everywhere that you should access them everywhere.

## Should We Really Hunt and Burn Global States?

As always, if you develop a small application (a plugin or a library, for example), and you know that it won’t grow, you can use global mutable states as much as you like. Since the scope of the application is (and will stay) tiny anyway, you won’t have difficulties to know where the globals are used, where their states are modified, and so on.

However, as we all know, application are, most of the time, meant to grow. To summarize what we should keep in mind in that case:

- Global mutable states should be avoided as much as possible.
- Using function arguments, dependency injection and context objects can reduce the scope of any global mutable state. Reducing their scopes is what you want.
- Global variables are just the tip of the icebergs: what is really important to respect is the general principle of encapsulation, and there are many ways to break it. Global mutable states are just one (extreme) example.
- Global immutable states are less harmful, but it’s better not to use them everywhere either.

Nothing is totally wrong or good in software development. Everything can be useful depending on the use case, and it’s one of the most difficult thing to do: taking the right decision for your precise needs. Experimenting, prototyping and be careful of the consequences would be your best bet.

However, I would argue that global mutable state do more harm than goods in most cases.

A last word: when you think about it, the root cause of the problems I described in this article are the states themselves, local or global, and their mutability.

If everything would be immutable, all these problems would not pop up as often as it does in most applications.

Obviously, we need to change states for our applications to work as expected. That’s why you should keep in mind that it’s always better to avoid mutability as much as you can!

> A large fraction of the flaws in software development are due to programmers not fully understanding all the possible states their code may execute in.

**John Carmack**

### Let's Connect

You'll receive **each month** the last article with additional resources and updates.

[Here's how it looks](https://buttondown.email/thevaluabledev/archive/the-valuable-dev-new-article-about-vim-and-many/)

You can reply to any email if you have questions, problems, or feedback. I'll write back as soon as I can.

Share Your Knowledge
