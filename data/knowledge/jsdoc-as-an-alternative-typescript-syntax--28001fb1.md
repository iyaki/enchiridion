---
title: "JSDoc as an alternative TypeScript syntax"
notion_id: 28001fb1-c877-46c7-abc6-bb9e0bb7d193
notion_url: https://app.notion.com/p/JSDoc-as-an-alternative-TypeScript-syntax-28001fb1c87746c7abc6bb9e0bb7d193
last_edited: 2024-03-11T14:06:00.000Z
source_url: https://alexharri.com/blog/jsdoc-as-an-alternative-typescript-syntax
tags: ["Alex Harri", "English", "Javascript", "Article"]
---
As web development has embraced static typing during the past decade, TypeScript has become the default language of choice. I think this is great—I love working with TypeScript!

But what if you can't use TypeScript? You may encounter circumstances where you need to work in plain JavaScript, be it tooling constraints or a team member who does not like static typing.

Under these circumstances, look to JSDoc for salvation:

```plain text
/** * @param {number} a * @param {number} b * @returns {number} */function add(a, b) {return a + b;}
```

I was surprised when I learned that the TypeScript compiler actually understands JSDoc comments. This fact allows you to type your entire codebase without creating a single `.ts` file.

Think of this post as your crash course in using JSDoc as an alternative syntax for TypeScript. We'll cover all the important TypeScript-related features JSDoc has to offer—and their limitations.

## JSDoc

JSDoc is expressed through block comments in the form `/** */`, which may contain _block tags_ such as `@param` and `@type`. Normal `//` and `/* */` comments don't work:

```plain text
// @type {number}let a; // Doesn't work/* @type {number} */let b; // Doesn't work/*** @type {number} */let c; // Doesn't work/** @type {number} */let d; // Works!
```

The majority of your JSDoc block tags will be used for typing variables, arguments, and return types. The block tags for those are `@type`, `@param`, and `@returns`.

```plain text
/** * @param {string} message * @returns {number} */function len(message) {return message.length;}/** @type {{ name: string, age: number }} */const user = {name: "Alex",age: 26,};
```

### Type casting

Type casting in TypeScript can be done using `expression as T` or `<T>expression`:

Type casting in JSDoc is done by wrapping the expression in parentheses and adding a preceding `@type` comment:

The parentheses are required. If they are missing the cast will not work:

Missing parentheses are a really easy mistake to make, which can easily lead to bugs when casting from `any`. Be careful with casts in JSDoc!

### Const assertions

TypeScript supports [const assertions](https://www.typescriptlang.org/docs/handbook/release-notes/typescript-3-4.html#const-assertions), which can be quite useful.

You can also use const assertions in JSDoc, they're just a type cast:

### Declaring types

In TypeScript, you can declare types using the `type` or `interface` keywords:

In JSDoc, types are declared using the `@typedef` keyword:

Having declared a type with `@typedef`, you can reference it like any other TypeScript type:

An alternative way to declare the properties of an object type is using `@property`:

Nested properties can be specified using `.` as a separator:

There's no syntax for exporting types in JSDoc. Instead, types defined using `@typedef` are **exported by default**. This auto-exporting applies to all types declared at the top level of a module.

As someone who cares a lot about the interfaces of modules, I strongly dislike this feature.

You can avoid the auto-exporting by declaring types in the scope that they're needed in:

One thing worth mentioning is that types declared in JavaScript modules using JSDoc can be imported from TypeScript modules.

### Importing types

In TypeScript, you can reference types from other modules via `import` statements or `import("./path").Type`:

Note: In TypeScript modules you can declare that the import is for a type via `import type { Foo }` or `import { type Foo }`

JSDoc only allows you to use `import("./path")`:

This can get quite verbose for long module paths and type names, so you can "fake" normal imports using `@typedef`:

But keep the auto-exporting footgun in mind! As mentioned in [Exporting types](https://alexharri.com/blog/jsdoc-as-an-alternative-typescript-syntax#exporting-types), types defined via `@typedef` are auto-exported, which means that the type is re-exported.

### Non-null assertions

We've arrived at my largest gripe with JSDoc: **it doesn't support non-null assertions**.

Take the `Map<K, V>` data structure as an example. The return type of `Map<K, V>.get` is `V | undefined`, which can be frustrating when you know for certain that a value is non-null.

In TypeScript, you can use `!` after an expression to assert that it is non-nullable.

There is no equivalent `@nonnull` tag or syntax in JSDoc.

One possible workaround is to use a type cast like so:

The problem with type casts is that they can become incorrect as the code evolves. Imagine that `map` is updated to store `string | number` instead of just `number`:

Type casting `string | number | undefined` to `number` is valid, so we get no type error. **The type cast masks the type error**, which would have not happened using non-null assertions.

There is one safe way to express non-nullability in JSDoc, which is using the `NonNullable` type in conjunction with `typeof`. Any expression `expr` can be declared non-nullable by casting it to `NonNullable<typeof expr>`, though this can be quite verbose.

We can make this more readable like so:

But this is still terribly noisy! This would be much cleaner if `@nonnull` were supported:

This issue is being tracked in [#23405 in microsoft/TypeScript](https://github.com/microsoft/TypeScript/issues/23405). Let us pray that `@nonnull` will be added at some point.

### Optional parameters

Parameters can be marked as optional in TypeScript using `?`:

In JSDoc, you can mark parameters as optional by wrapping their name in `[]`:

A parameter can also be implicitly marked as optional by providing a default argument, just like in TypeScript:

There is an alternative syntax for marking parameters as optional where `=` is placed after the type:

I find this syntax a bit weird, but hey, it's supported.

### Generic type parameters

Declaring a generic type parameter is done using `@template`:

Expressing an `extends` constraint is done like so:

The `@template` block tag can also be used for type definitions, classes, methods, and more.

### Class properties

In TypeScript, you can declare class properties using the `public`/`private` keywords for constructor arguments or by explicitly declaring the properties.

Since JavaScript does not support the `public`/`private` keywords, we need to take the latter approach and assign manually:

Unlike TypeScript, we don't need to explicitly declare `x` and `y` as properties in JavaScript modules. They are implicitly declared by assigning to them in the constructor.

However, I would argue that it's good practice to explicitly declare the types of class properties to avoid possible implicit `any`s.

### Class `implements`

In TypeScript you can declare that a class `implements` a certain interface:

It won't come as a surprise to hear that JSDoc has `@implements` for this purpose:

### Public and private properties and methods

Properties and methods can be declared as `public` and `private` via `@public` and `@private`:

### Typing `this`

TypeScript enables you to type the `this` argument for a function or method:

JSDoc contains an `@this` keyword for this purpose:

### `@ts-*` comments

All of your normal `@ts-*` comments, such as `@ts-ignore`, work as expected:

## Practical matters

We've now gone through all the major (in my opinion) TypeScript-related features in JSDoc. They should cover the vast majority of TypeScript features you'll ever need in JSDoc.

We'll now cover some practical things to know if you intend to use JSDoc with TypeScript.

### Enable `checkJs`

As we've seen, type annotations in JSDoc comments are used as type information in `.js` files.

However, `checkJs` needs to be enabled in your `tsconfig.json` for type errors to be emitted. If you don't enable `checkJs`, your JSDoc comments will only be used for IDE annotations—not type checking. Be sure to enable it!

### TypeScript interop

If you type a function using JSDoc in a `.js` module, you can import that function in a `.ts` module without any issues. This also works the other way: you can import things from `.ts` modules and use them in `.js` modules.

Generally, interop between `.js` modules using JSDoc and `.ts` modules "just works".

### JSDoc does not work in TypeScript modules

You can't use JSDoc for type annotations in `.ts` modules. This can make migrating from JSDoc to TypeScript a bit frustrating, especially for larger modules.

## Conclusion

I worked in a JSDoc codebase for a significant amount of time, and have gone through the process of migrating a lot of that codebase to TypeScript. JSDoc definitely has flaws, such as its clunky and verbose syntax, but it's still a perfectly viable way to go about typing your codebase.

If you're not able to use TypeScript for some reason, then consider giving JSDoc a shot. It's better than no types.
