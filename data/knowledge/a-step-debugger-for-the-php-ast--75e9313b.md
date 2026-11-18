---
title: "A step-debugger for the PHP AST"
notion_id: 75e9313b-4d2a-4b14-b5b2-b46a1e8e292e
notion_url: https://app.notion.com/p/A-step-debugger-for-the-PHP-AST-75e9313b4d2a4b14b5b2b46a1e8e292e
last_edited: 2023-01-13T17:08:00.000Z
source_url: https://matthiasnoback.nl/2022/09/a-step-debugger-for-the-php-ast/
tags: ["Article", "Matthias Noback Blog", "English", "PHP"]
---
When you're learning to write custom rules for [PHPStan](https://phpstan.org/developing-extensions/rules) or [Rector](https://github.com/rectorphp/rector/blob/main/docs/create_own_rule.md), you'll have to learn more about the PHP programming language as well. To be more precise, about the way the interpreter parses PHP code. The result of parsing PHP code is a tree of nodes which represents the structure of the code, e.g. you'll have a Class definition node, a Method definition node, and within those method Statement nodes, and so on. Each node can be checked for errors (with PHPStan), or automatically refactored in some way (with Rector).

The tree of nodes is called _Abstract Syntax Tree_, and a successful PHPStan or Rector rule starts with selecting the right nodes from the tree and "subscribing" your rule to these nodes. A common approach for this is to start `var_dump`-ing or `echo`-ing nodes inside your new rule, but I've found this to be quite tedious. Which is why I've created a simple command-line tool that lets you inspect the nodes of any given PHP file.

The tool is called [AST Inspector](https://github.com/matthiasnoback/php-ast-inspector/) and is available on GitHub.

Install it with Composer:

```plain text
composer require --dev matthiasnoback/php-ast-inspector

```

Then run:

```plain text
vendor/bin/ast-inspect inspect [file.php]

```

You'll see something similar to this output:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

You can navigate through the tree by going to the next or previous node, or jumping into the subnodes of the selected node. Navigation conveniently uses the `a,s,d,w` keys.

Currently the project uses the PHP-Parser library for parsing. Since PHPStan adds additional virtual nodes to the AST, it will be useful to show them in this tool as well, but that requires some additional work. Another interesting addition would be to show the types that PHPStan derives for variables in the inspected code. That will also require some more work...

For now, please give this program a try, and let me know what you think! I'm happy to add more features to it, as long as it makes the learning curve for these amazing tools less steep. And if you're looking for an in-depth exploration of writing your own PHPStan or Rector rules, check out the documentation linked above or one of my books ([Recipes for Decoupling](https://leanpub.com/recipes-for-decoupling), which shows how to create PHPStan rules, and [Rector - The Power of Automated Refactoring](https://leanpub.com/rector-the-power-of-automated-refactoring/), which does the same for Rector).
