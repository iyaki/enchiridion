---
title: "Modern PHP Cheat Sheet"
notion_id: 3a4005f7-1396-4802-bfd4-7e9796d27ed0
notion_url: https://app.notion.com/p/Modern-PHP-Cheat-Sheet-3a4005f713964802bfd47e9796d27ed0
last_edited: 2022-12-21T14:18:00.000Z
source_url: https://front-line-php.com/cheat-sheet
tags: ["CheatSheet", "English", "PHP"]
---
This sample is an excerpt from the ebook Front Line PHP. Check out the entire book to learn how to build modern applications in PHP 8.3.

## # Arrays

### # Destructuring 7.1

You can destructure arrays to pull out several elements into separate variables:

```plain text
$array = [1, 2, 3]; // Using the list syntax: list($a, $b, $c) = $array; // Or the shorthand syntax: [$a, $b, $c] = $array;
```

You can skip elements:

```plain text
[, , $c] = $array; // $c = 3
```

As well as destructure based on keys:

```plain text
$array = [ 'a' => 1, 'b' => 2, 'c' => 3, ]; ['c' => $c, 'a' => $a] = $array;
```

### # Rest and Spread Operators 5.6

Arrays can be spread into functions:

```plain text
$array = [1, 2]; function foo(int $a, int $b) { /* … */ } foo(...$array);
```

Functions can automatically collect the rest of the variables using the same operator:

```plain text
function foo($first, ...$other) { /* … */ } foo('a', 'b', 'c', 'd', …);
```

Rest parameters can even be type hinted:

```plain text
function foo($first, string ...$other) { /* … */ } foo('a', 'b', 'c', 'd', …);
```

7.4 Arrays with numerical keys can also be spread into a new array:

```plain text
$a = [1, 2]; $b = [3, 4]; $result = [...$a, ...$b]; // [1, 2, 3, 4]
```

8.1 You can also spread arrays with textual keys starting from PHP 8.2

## # Attributes 8.0

Make your own by tagging a class with #[Attribute]

```plain text
#[Attribute] class ListensTo { public string $event; public function __construct(string $event) { $this->event = $event; } }
```

You can add them to properties, (anonymous) classes, functions, constants, closures, function arguments:

```plain text
#[ Route(Http::POST, '/products/create'), Autowire, ] class ProductsCreateController { public function __invoke() { /* … */ } }
```

Use reflection to get them, you can pass in optional arguments to `

getAttributes

` in order to filter the result:

```plain text
$attributes = $reflectionClass->getAttributes( ContainerAttribute::class, ReflectionAttribute::IS_INSTANCEOF );
```

## # Closures

## # Short Closures 7.4

Short closures have automatic access to the outer scope, and only allow a single expression which is automatically returned:

```plain text
array_map( fn($x) => $x * $this->modifier, $numbers );
```

## # First class callables 8.1

You can now make a closure from a callable by calling that callable and passing it '...' as its argument

```plain text
function foo(int $a, int $b) { /* … */ } $foo = foo(...); $foo(a: 1, b: 2);
```

## # Cosmetics

### # Class Names 8.0

As of PHP 8, you can use ::class on objects as well:

```plain text
Order::class; $object::class;
```

### # Numeric Values 7.4

Use the _ operator to format numeric values:

```plain text
$price = 100_10; // $100 and 10 cents
```

### # Trailing Commas 8.0

Trailing commas are allowed in several places:

- Arrays
- Function calls
- Function definitions
- Closure use statements

## # Enums

### # Declaring Enums 8.1

Enums are built-in in the language:

```plain text
enum Status { case DRAFT; case PUBLISHED; case ARCHIVED; }
```

### # Enum methods 8.1

Enums can have methods, as well as have a string or integer value per case:

```plain text
enum Status: int { case DRAFT = 1; case PUBLISHED = 2; case ARCHIVED = 3; public function color(): string { return match($this) { Status::DRAFT => 'grey', Status::PUBLISHED => 'green', Status::ARCHIVED => 'red', }; } }
```

## # Exceptions 8.0

Throwing an exception is an expression now, which means there are more places you can throw from, such as short closures or as a null coalescing fallback:

```plain text
$error = fn($message) => throw new Error($message); $input = $data['input'] ?? throw new Exception('Input not set');
```

You also don't have to catch an exception with a variable anymore:

```plain text
try { // … } catch (SpecificException) { throw new OtherException(); }
```

## # Match 8.0

Similar to switch, but with strong type checks, no break keyword, combined arms and it returns a value:

```plain text
$message = match ($statusCode) { 200, 300 => null, 400 => 'not found', 500 => 'server error', default => 'unknown status code', };
```

## # Dealing with null

### # Null Coalescing 7.0

Use the null coalescing operator to provide a fallback when a property is null:

```plain text
$paymentDate = $invoice->paymentDate ?? Date::now();
```

It also works nested:

```plain text
$input = $data['few']['levels']['deep'] ?? 'foo';
```

7.4

You can use the null coalescing assignment operator to write the value into the original variable when it's null:

```plain text
$temporaryPaymentDate = $invoice->paymentDate ??= Date::now(); // $invoice->paymentDate is now also set
```

### # Nullsafe Operator 8.0

Chain methods that possibly return null:

```plain text
$invoice ->getPaymentDate() ?->format('Y-m-d');
```

The nullsafe operator can also be chained multiple times:

```plain text
$object ?->methodA() ?->methodB() ?->methodC();
```

## # Named Arguments 8.0

Pass in arguments by name instead of their position:

```plain text
setcookie( name: 'test', expires: time() + 60 * 60 * 2, );
```

Named arguments also support array spreading:

```plain text
$data = [ 'name' => 'test', 'expires' => time() + 60 * 60 * 2, ]; setcookie(...$data);
```

## # Performance

### # The JIT 8.0

Enable the JIT by specifying a buffer size in your ini settings; you can switch the jit mode between function or tracing:

```plain text
opcache.jit_buffer_size=100M opcache.jit=function ; opcache.jit=tracing
```

### # Preloading 7.4

Add opcache.preload to your ini settings:

```plain text
opcache.preload=/path/to/project/preload.php
```

Every file that's loaded in the preload script will be preloaded into memory until server restart.

## # Properties

## # Using new in initializers 8.1

PHP 8.2 allows you to use the new keyword in function definitions as a default parameter, as well as in attribute arguments and other places.

```plain text
class MyController { public function __construct( private Logger $logger = new NullLogger(), ) {} }
```

This also means that nested attributes are a thing now:

```plain text
#[Assert\All( new Assert\NotNull, new Assert\Length(max: 6), )]
```

## # Read only properties 8.1

Properties can be marked readonly:

```plain text
class Offer { public readonly ?string $offerNumber = null; public readonly Money $totalPrice; }
```

Once a readonly property is set, it cannot change ever again.

8.2

Classes can be completely readonly:

```plain text
readonly class Offer { public ?string $offerNumber = null; public Money $totalPrice; }
```

All properties of a class will become readonly, adding properties dynamically is impossible.

## # Types

### # Built-in Types

During the PHP 7.x releases and with PHP 8, several new built-in types were added:

- void: a return type indicating nothing's returned 7.1
- static: a return type representing the current class or its children 8.0
- object: anything that is an object 7.2
- mixed: anything 8.0

### # Union Types 8.0

Combine several types into one union, which means that whatever input must match one of the given types:

```plain text
interface Repository { public function find(int|string $id); }
```

### # Typed properties 7.4

Add types to your class properties:

```plain text
class Offer { public ?string $offerNumber = null; public Money $totalPrice; }
```

Beware of the uninitialized state, which is only checked when reading a property, and not when constructing the class:

```plain text
$offer = new Offer(); $offer->totalPrice; // Error
```

### # Intersection Types 8.1

Combine several types into one intersection, which means that whatever input must match all of the given types

```plain text
public function url(WithId&WithSlug $obj): string;
```

### # DNF Types 8.2

You can create a union of intersection types, useful for creating a nullable intersection type:

```plain text
public function url((WithId&WithSlug)|null $obj): string;
```
