---
title: "Functional PHP: Functional primitives for PHP"
notion_id: 4ac02eb0-9b41-4a88-a457-253d350f6efc
notion_url: https://app.notion.com/p/Functional-PHP-Functional-primitives-for-PHP-4ac02eb09b414a88a457253d350f6efc
last_edited: 2026-09-21T17:21:00.000Z
source_url: https://github.com/lstrojny/functional-php#functional-php-functional-primitives-for-php
tags: ["Framework/Library", "English", "PHP", "Untried", "Functional programming"]
---
[https://github.com/lstrojny/functional-php#functional-php-functional-primitives-for-php](https://github.com/lstrojny/functional-php#functional-php-functional-primitives-for-php)

## Folders and files

NameNameLast commit messageLast commit date.github.githubdocsdocssrc/Functionalsrc/Functionaltests/Functionaltests/Functional.editorconfig.editorconfig.gitignore.gitignore.php_cs.dist.php_cs.dist.phpstorm.meta.php.phpstorm.meta.php.scrutinizer.yml.scrutinizer.ymlLICENSELICENSEREADME.mdREADME.mdcomposer.jsoncomposer.jsonphpcs.xml.distphpcs.xml.distphpunit.xml.distphpunit.xml.dist

- README
- MIT license

NOTE: functional-php used to come with a C extension that implemented most of the functions natively. As the performance differences weren’t that huge compared to the maintenance cost it has been removed.

A set of functional primitives for PHP, heavily inspired by Scala’s traversable collection, Dojo’s array functions and Underscore.js

- Works with arrays and everything implementing interface Traversable
- Consistent interface: for functions taking collections and callbacks, first parameter is always the collection, then the callback. Callbacks are always passed $value, $index, $collection. Strict comparison is the default but can be changed
- Calls 5.3 closures as well as usual callbacks
- All functions reside in namespace Functional to not raise conflicts with any other extension or library

## Installation

Run the following command in your project root:

```plain text
composer require lstrojny/functional-php
```

## Docs

Read the docs

## Contributing

1. Fork and git clone the project
2. Install dependencies via composer install
3. Run the tests via composer run tests
4. Write code and create a PR

## Mailing lists

- General help and development list: http://groups.google.com/group/functional-php
- Commit list: http://groups.google.com/group/functional-php-commits

## Thank you

- Richard Quadling and Pierre Joye for Windows build help
- David Soria Parra for various ideas and the userland version of Functional\flatten()
- Max Beutel for Functional\unique(), Functional\invoke_first(), Functional\invoke_last() and all the discussions
- An Phan for many great contributions
