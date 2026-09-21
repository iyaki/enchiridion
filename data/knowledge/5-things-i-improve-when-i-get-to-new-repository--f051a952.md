---
title: "5 Things I Improve when I Get to new Repository"
notion_id: f051a952-736b-4a86-85b9-b60938f03af9
notion_url: https://app.notion.com/p/5-Things-I-Improve-when-I-Get-to-new-Repository-f051a952736b4a8685b9b60938f03af9
last_edited: 2023-04-20T19:41:00.000Z
source_url: https://tomasvotruba.com/blog/2019/12/23/5-things-i-improve-when-i-get-to-new-repository
tags: ["Article", "Tomas Votruba Blog", "English", "Programming", "Productivity", "Untried"]
---
I started to write this post as follow up for [clean and sustainable code](https://tomasvotruba.com/blog/2019/12/16/8-steps-you-can-make-before-huge-upgrade-to-make-it-faster-cheaper-and-more-stable/) post. In the middle of writing, I've realized I have this approach to ever repository I meet.

Imagine it like a working desk. But not your usual stable place where you work every day. **Instead, you are assigned to a new desk of a former employee, who worked in the company for 5 years and as a bonus - it was the CTO**. For you, it's a mess.

What is the first thing we do? We'll **prepare it for hard work**.

I get to 2-3 new projects/week and during the last couple of years I've noticed **I repeat the same preparing process before work itself**. It makes me much more effective and creates a very intuitive environment to work in.

## 1. Set 1 Spacing Rule

Without this file, every file has a different number of spaces, tabs, line-endings... and everything else we can't see.

Well, until you have errors like:

- _YAML syntax error_
- _Mix of tabs and spaces_

Or _creatively_ structured code like:

```plain text
<?php

class SomeClass
{
public function someMethod()
{
}
}

```

We don't need these problems, it's the computers' job.

**This takes 1 minute to set up and commit**:

```plain text
# .editorconfig
root = true

[*]
charset = utf-8
end_of_line = lf
insert_final_newline = true
trim_trailing_whitespace = true
indent_style = space
indent_size = 4

```

✅

## 2. Make Sure vendor is in `vendor`

The easier miss-location, that works as default on some operation systems, is this:

```plain text
{
"config": {
"vendor-dir": "Vendor"
    }
}

```

It might be a great feeling to be creative for various reasons, but **most PHP tools are not ready for this**, e.g. PHPStan and Rector fail here.

Unless there is some supercritical issue, I always make it a standard way:

```plain text
-{
-    "config": {
-        "vendor-dir": "Vendor"
-    }
-}

```

✅

## 3. Move Code to `/src`

In times old so I can't remember, code was located randomly. Then PSR-0 came, then PSR-4. [Many tools depend on PSR-4](https://tomasvotruba.com/blog/2019/12/16/8-steps-you-can-make-before-huge-upgrade-to-make-it-faster-cheaper-and-more-stable/#1-psr-4-standard), so another standard naturally came to my toolset.

> Keep source code in

```plain text
/src
```

That means only source code needed for production, so:

- **no migrations**
- **no fixtures**
- **no test helpers**
- **no coding standard utils**
- **no Rector utils**
- **no templates**
- **no translations**
- **no configs**
- **no bin files**
- **no helpers bash scripts**
- **no cool git repository tricks**

Once we rule in place know that every command, that works with PHP code will get only one argument: `src`

```plain text
vendor/bin/ecs c src
vendor/bin/phpstan a src
vendor/bin/rector p src

```

✅

## 4. Directory Name = Directory Content

Saying the one above, I apply the same for other content:

- binary files? → `bin`
- files used in continues integration? → `ci`
- configs? → `configs`
- templates? → `templates`
- translations? → `translations`
- database migrations? → `migration`

...and so on.

You may know it as:

> 1 level architecture.

It's very intuitive to use, based on UX, DX and well... **human brain. We tend to choose simpler solutions over complex ones**. Often it leads to crappy application design.

I have a special case for Rector, coding standards, PHPStan rules, utils that helps in development, but aren't part of the project itself → `utils/<project>/src`

✅

## 5. `fs`, `ps`, `pu`... 2 Chars Shortcuts for Tools that Help Me

```plain text
fs
ps
pu

```

6 characters, even "characters" has 10 characters.

I can't imagine to code without them. What are they?

```plain text
vendor/bin/ecs check src tests --fix
vendor/bin/phpstan analyse src tests
vendor/bin/phpunit

```

Well, now you know [I use aliases in my bash](https://tomasvotruba.com/blog/2019/11/25/the-single-best-skill-to-master-command-line/). It's **the ultimate **_**skill**_, because your brain gets much more space to think.

> "But every project has different directories.

That's right, in one project it is:

```plain text
vendor/bin/ecs check src tests --fix

```

in another it is:

```plain text
vendor/bin/ecs check src packages tests --fix

```

Well, this how my bash aliases look like:

```plain text
alias fs="composer fix-cs"
alias ps="composer phpstan"
alias pu="vendor/bin/phpunit"

```

**I never change them**. So where is the dynamic part?

Have you heard of [composer scripts](https://blog.martinhujer.cz/have-you-tried-composer-scripts)?

In every project I came to, I set up dev dependencies and scripts first:

```plain text
{
"require-dev": {
"symplify/easy-coding-standard": "^7.1",
"phpstan/phpstan": "^0.12",
"phpunit/phpunit": "^8.5"
    },
"scripts": {
"fix-cs": "vendor/bin/ecs check bin src tests --fix --ansi",
"phpstan": "vendor/bin/phpstan analyse bin src tests --ansi --error-format symplify"
    }
}

```

That way I can modify directories right in `composer.json`.

So **when I do any change in the code**:

- I open 3 terminals in PHPStorm console
- I run 3 scripts in parallel = **it's faster and I can focus better on 1 tool**
- I know what came wrong and re-run only the broken part

```plain text
fs
ps
pu

```

As a side benefit, continuous integration is easier to set up and maintain:

```plain text
# travis.yml
jobs:
include:
        -
stage: test
name: ECS
script:
                - composer check-cs

        -
name: PHPStan
script:
                - composer phpstan

```

Is there one new directory `tests` to check? Just update `composer.json`:

```plain text
 {
     "require-dev": {
         "symplify/easy-coding-standard": "^7.1",
         "phpstan/phpstan": "^0.12",
         "phpunit/phpunit": "^8.5"
     },
     "scripts": {
-        "fix-cs": "vendor/bin/ecs check src --fix --ansi",
+        "fix-cs": "vendor/bin/ecs check src tests --fix --ansi",
-        "phpstan": "vendor/bin/phpstan analyse src --ansi"
+        "phpstan": "vendor/bin/phpstan analyse src tests --ansi"
    }
}

```

This way I also see, **what directories contain PHP code, that needs to be checked**.

✅

> "If you want to go quickly, go alone.

Happy Xmass Coding!
