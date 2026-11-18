---
title: "A Complete Course of the Raku programming language"
notion_id: a5df940d-21a4-4c5e-b0ce-2cac9e763519
notion_url: https://app.notion.com/p/A-Complete-Course-of-the-Raku-programming-language-a5df940d21a44c5eb0ce2cac9e763519
last_edited: 2023-09-13T14:40:00.000Z
source_url: https://course.raku.org/
tags: ["English", "Raku", "Guide"]
---
## Part 1. Raku essentials

### Basic introduction to Raku and its compiler

- [What is Raku](https://course.raku.org/essentials/what-is-raku)
- [Raku vs. Rakudo](https://course.raku.org/essentials/raku-vs-rakudo)
- [How to install Rakudo](https://course.raku.org/essentials/how-to-install-rakudo)
- [Editors and IDEs](https://course.raku.org/essentials/editors-and-ides)
- [Conventional file extensions](https://course.raku.org/essentials/conventional-file-extensions)
- [Hello, World!](https://course.raku.org/essentials/hello-world)
- [Notes on using Unicode](https://course.raku.org/essentials/on-unicode)
- [Running programs](https://course.raku.org/essentials/running-programs) 
- [Running from command line](https://course.raku.org/essentials/running-programs/from-command-line)
- [Running from REPL](https://course.raku.org/essentials/running-programs/from-repl)
- [Running from IDE](https://course.raku.org/essentials/running-programs/from-ide)
- [Using docker](https://course.raku.org/essentials/running-programs/using-docker)
- [Using online services](https://course.raku.org/essentials/running-programs/using-online-services)
- [Simple input and output](https://course.raku.org/essentials/simple-input-output) 
- [Output with ](https://course.raku.org/essentials/simple-input-output/say)[`say`](https://course.raku.org/essentials/simple-input-output/say)
- [Input with ](https://course.raku.org/essentials/simple-input-output/prompt)[`prompt`](https://course.raku.org/essentials/simple-input-output/prompt)
- [Comments](https://course.raku.org/essentials/comments) 
- [Single-line comments](https://course.raku.org/essentials/comments/single-line)
- [Multi-line comments](https://course.raku.org/essentials/comments/multi-line)
- [Embedded comments](https://course.raku.org/essentials/comments/embedded)

### Variables and data types essentials

- [Scalar variables](https://course.raku.org/essentials/scalar-variables) 
- [Declaring a variable](https://course.raku.org/essentials/scalar-variables/declaring-a-variable)
- [Assigning a value](https://course.raku.org/essentials/scalar-variables/assigning-a-value)
- [Declaration with initialization](https://course.raku.org/essentials/scalar-variables/declaration-with-initialization)
- [The defined-or operator](https://course.raku.org/essentials/scalar-variables/defined-or-operator)
- [Names of the variables](https://course.raku.org/essentials/scalar-variables/identifiers)
- [Strings](https://course.raku.org/essentials/strings) 
- [String concatenation](https://course.raku.org/essentials/strings/string-concatenation)
- [Variable interpolation](https://course.raku.org/essentials/strings/variable-interpolation)
- [Code interpolation](https://course.raku.org/essentials/strings/code-interpolation)
- [Escaping special characters](https://course.raku.org/essentials/strings/escaping-special-characters)
- [String length](https://course.raku.org/essentials/strings/string-length)
- [Numbers](https://course.raku.org/essentials/numbers) 
- [Integer numbers](https://course.raku.org/essentials/numbers/integers)
- [Rational numbers](https://course.raku.org/essentials/numbers/rats)
- [Floating-point numbers](https://course.raku.org/essentials/numbers/numeric)
- [Operations with numbers](https://course.raku.org/essentials/numbers/operations)
- [Boolean type](https://course.raku.org/essentials/booleans) 
- [Boolean operations](https://course.raku.org/essentials/booleans/operations)
- [Boolean operations with other types](https://course.raku.org/essentials/booleans/boolean-operations-other-types)
- [Ranges](https://course.raku.org/essentials/ranges) 
- [Excluding endpoints](https://course.raku.org/essentials/ranges/excluding-endpoints)
- [Matching against a range](https://course.raku.org/essentials/ranges/matching-against-a-range)

### Control flow essentials

- [Code blocks](https://course.raku.org/essentials/code-blocks) 
- [Lexical scope](https://course.raku.org/essentials/code-blocks/lexical-scope)
- [Global variables](https://course.raku.org/essentials/code-blocks/global-variables)
- [Local variables](https://course.raku.org/essentials/code-blocks/local-variables)
- [Conditional checks](https://course.raku.org/essentials/conditional-checks) 
- [`if`](https://course.raku.org/essentials/conditional-checks/if)[ blocks](https://course.raku.org/essentials/conditional-checks/if)
- [`else`](https://course.raku.org/essentials/conditional-checks/else)[ blocks](https://course.raku.org/essentials/conditional-checks/else)
- [Using ](https://course.raku.org/essentials/conditional-checks/elsif)[`elsif`](https://course.raku.org/essentials/conditional-checks/elsif)
- [Using ](https://course.raku.org/essentials/conditional-checks/unless)[`unless`](https://course.raku.org/essentials/conditional-checks/unless)
- [`if`](https://course.raku.org/essentials/conditional-checks/modifiers)[ and ](https://course.raku.org/essentials/conditional-checks/modifiers)[`unless`](https://course.raku.org/essentials/conditional-checks/modifiers)[ as statement modifiers](https://course.raku.org/essentials/conditional-checks/modifiers)
- [Comparing numbers](https://course.raku.org/essentials/conditional-checks/comparing-numbers)
- [Comparing strings](https://course.raku.org/essentials/conditional-checks/comparing-strings)
- [Ternary operator](https://course.raku.org/essentials/conditional-checks/ternary-operator)
- [Loops](https://course.raku.org/essentials/loops) 
- [Using ](https://course.raku.org/essentials/loops/while)[`while`](https://course.raku.org/essentials/loops/while)
- [Using ](https://course.raku.org/essentials/loops/until)[`until`](https://course.raku.org/essentials/loops/until)
- [Using ](https://course.raku.org/essentials/loops/repeat)[`repeat`](https://course.raku.org/essentials/loops/repeat)
- [`while`](https://course.raku.org/essentials/loops/modifiers)[ and ](https://course.raku.org/essentials/loops/modifiers)[`until`](https://course.raku.org/essentials/loops/modifiers)[ as statement modifiers](https://course.raku.org/essentials/loops/modifiers)
- [Three-statement ](https://course.raku.org/essentials/loops/loop)[`loop`](https://course.raku.org/essentials/loops/loop)
- [Infinite loops](https://course.raku.org/essentials/loops/infinite-loops)
- [`for`](https://course.raku.org/essentials/loops/for)[ loops](https://course.raku.org/essentials/loops/for)
- [Topic variable](https://course.raku.org/essentials/loops/topic)
- [Postfix form of ](https://course.raku.org/essentials/loops/postfix-for)[`for`](https://course.raku.org/essentials/loops/postfix-for)

### More about types

- [Data type conversion](https://course.raku.org/essentials/coercion) 
- [Introspection with ](https://course.raku.org/essentials/coercion/what)[`WHAT`](https://course.raku.org/essentials/coercion/what)
- [Converting types using type coercion methods](https://course.raku.org/essentials/coercion/methods)
- [Converting types using constructor forms](https://course.raku.org/essentials/coercion/new)
- [Converting types with prefix operators](https://course.raku.org/essentials/coercion/prefixes)
- [Typed variables](https://course.raku.org/essentials/typed-variables) 
- [Type constraints](https://course.raku.org/essentials/typed-variables/type-constraints)
- [Using ](https://course.raku.org/essentials/typed-variables/using-of)[`of`](https://course.raku.org/essentials/typed-variables/using-of)
- [Type conversion for typed variables](https://course.raku.org/essentials/typed-variables/coercion)
- [Allomorphs](https://course.raku.org/essentials/typed-variables/allomorphs)
- [Positional data types](https://course.raku.org/essentials/positionals) 
- [Arrays](https://course.raku.org/essentials/positionals/arrays)
- [Nested arrays](https://course.raku.org/essentials/positionals/nested-arrays)
- [The ](https://course.raku.org/essentials/positionals/args-array)[`@*ARGS`](https://course.raku.org/essentials/positionals/args-array)[ array](https://course.raku.org/essentials/positionals/args-array)
- [Interpolating arrays](https://course.raku.org/essentials/positionals/interpolating-arrays)
- [Lists](https://course.raku.org/essentials/positionals/lists)
- [Quoting string arrays](https://course.raku.org/essentials/positionals/quoting-string-arrays)
- [Subscripting ranges](https://course.raku.org/essentials/positionals/subscripting-ranges)
- [Associative data types](https://course.raku.org/essentials/associatives) 
- [Pairs](https://course.raku.org/essentials/associatives/pairs)
- [Hashes](https://course.raku.org/essentials/associatives/hashes)
- [Nested hashes](https://course.raku.org/essentials/associatives/nested-hashes)
- [Interpolating hashes](https://course.raku.org/essentials/associatives/interpolating-hashes)

### Functions essentials

- [Creating and calling functions](https://course.raku.org/essentials/functions) 
- [Function names](https://course.raku.org/essentials/functions/identifiers)
- [Function parameters](https://course.raku.org/essentials/functions/function-parameters)
- [Returning the result](https://course.raku.org/essentials/functions/return-result)
- [Positional parameters](https://course.raku.org/essentials/functions/positional-parameters)
- [Named parameters](https://course.raku.org/essentials/functions/named-parameters)
- [Default values](https://course.raku.org/essentials/functions/default-values)
- [More about functions](https://course.raku.org/essentials/more-on-functions) 
- [Mind the space](https://course.raku.org/essentials/more-on-functions/mind-the-space)
- [Typed parameters](https://course.raku.org/essentials/more-on-functions/typed-parameters)
- [Return type](https://course.raku.org/essentials/more-on-functions/return-type)
- [Multi-functions](https://course.raku.org/essentials/more-on-functions/multi-functions)
- [Built-in functions for printing](https://course.raku.org/essentials/built-in-functions-for-printing) 
- [`say`](https://course.raku.org/essentials/built-in-functions-for-printing/say)
- [`print`](https://course.raku.org/essentials/built-in-functions-for-printing/print)
- [`put`](https://course.raku.org/essentials/built-in-functions-for-printing/put)
- [`note`](https://course.raku.org/essentials/built-in-functions-for-printing/note)
- [The ](https://course.raku.org/essentials/the-main-function)[`MAIN`](https://course.raku.org/essentials/the-main-function)[ function](https://course.raku.org/essentials/the-main-function) 
- [Reading command-line arguments](https://course.raku.org/essentials/the-main-function/reading-command-line-arguments)
- [Multiple ](https://course.raku.org/essentials/the-main-function/multi-main-functions)[`MAIN`](https://course.raku.org/essentials/the-main-function/multi-main-functions)[ functions](https://course.raku.org/essentials/the-main-function/multi-main-functions)
- [Using ](https://course.raku.org/essentials/the-main-function/using-unit-sub)[`unit sub`](https://course.raku.org/essentials/the-main-function/using-unit-sub)

## Part 2. Advanced Raku subjects

### Containers

- Understanding Raku containers
- Integers
- Strings
- Date and Time built-in support
- Other data times
- Understanding a sequence

### Operators

- Types of Raku operators (infix, prefix, etc.)
- Overview of operators in Raku
- Meta-operators
- User-defined operators

### Control flow

- Phasers
- Block-related phasers (e.g. `LEAVE`)
- Other options (e.g., `gather`, `given`)

### Subroutines

- Signature
- Multiple dispatch
- More on MAIN subroutines
- Nested subroutines
- Anonymous subroutines

### Modules

- Creating modules
- Using modules
- Different types of importing (`import`, `need`, etc.)
- Introspections
- Installing modules from web

## Part 3

### Object-oriented programming

- Classes in Raku
- Attributes
- Methods
- Class methods
- Subroutines vs methods
- Inheritance
- Roles
- Introspection

### Input and output

- Standard input, output, and errors
- Working with files
- Working with directories
- File streams

### Exceptions

- The `try` block
- What is a soft failure
- The `CATCH` phaser
- Exception objects
- Failure objects
- Multiple dispatch in handling exceptions
- Custom exceptions

## Part 4

### Regexes

- Literals and character classes
- Regexp matching
- Quantifiers
- Captures
- Alternations
- Anchors
- Forward and backward assertions
- Adverbs (such as `:g` etc.)
- String substitution and replacement

### Grammars

- What is a grammar
- Creating grammars
- Rules
- Tokens
- Grammars vs. classes and inheritance
- AST (Abstract syntax tree), `make` and `made`
- Actions
- Inline actions vs. action class

## Part 5

### Functional programming

- Recursion
- Reduction
- Higher-order functions
- Lambdas
- Data feeds
- Iterators
- Lazy and infinite sequences

### Concurrent programming

- Junctions
- Threads
- Promises
- Channels

### Reactive programming

- Supplies
- Live and on-demand supplies
- Understanding `react`
- Understanding `whenever`
- Understanding `await`

### Web programming

- Making remote connections
- Simple HTTP client
- Simple HTTP server
- Cro 101

## Appendix

- List of quizzes
- List of exercises
- 

© 2021
