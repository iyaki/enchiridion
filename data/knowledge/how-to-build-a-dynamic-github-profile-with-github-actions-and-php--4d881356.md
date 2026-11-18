---
title: "How to Build a Dynamic GitHub Profile with GitHub Actions and PHP"
notion_id: 4d881356-015f-4e75-903b-73557c262ffd
notion_url: https://app.notion.com/p/How-to-Build-a-Dynamic-GitHub-Profile-with-GitHub-Actions-and-PHP-4d881356015f4e75903b73557c262ffd
last_edited: 2026-09-21T17:18:00.000Z
source_url: https://hackernoon.com/how-to-build-a-dynamic-github-profile-with-github-actions-and-php-h5g34cr
tags: ["Producer (Individual Contributor)", "Article", "Tutorial", "Hackernoon", "English"]
---
[https://hackernoon.com/how-to-build-a-dynamic-github-profile-with-github-actions-and-php-h5g34cr](https://hackernoon.com/how-to-build-a-dynamic-github-profile-with-github-actions-and-php-h5g34cr)

Last year, GitHub quietly released a feature that was quickly noticed by the community – profile READMEs. A profile README is a global

```plain text
README
```

file for your GitHub profile, which you can set up by creating a public repository whose name is identical to your GitHub username. For instance, as my username is

```plain text
osteel
```

, I created the

```plain text
osteel/osteel
```

repository.

A little box like this one should appear while you add your own:

Once the repository is created, add a

```plain text
README
```

file with a short description explaining how great you are, and your GitHub profile page will display its content by default:

Neat and simple.

As I was browsing examples for some inspiration, I stumbled upon Simon Willison's version, which features some dynamic content like recent work and blog publications. He explained how he used a combination of GitHub Actions and Python to achieve this in a blog post, and I decided to do something similar with PHP.

## The placeholder

The first thing to do is to create a placeholder in the

```plain text
README
```

file where the dynamic content will go. Since I wanted to automatically insert the latest publications of my blog, I used the following tags:

```plain text
<!-- posts --><!-- /posts -->
```

You might recognise this format; since Markdown files also support HTML, I used some HTML comment tags to make sure they wouldn't show up on my profile page.

## The PHP script

I can't remember the last time I wrote some PHP without a framework; as a result, I had to do a quick search just to get started with a basic PHP script and some Composer dependencies.

Turn out it's quite simple! The first step is to initialise the project with the following command:

```plain text
$ composer init
```

From there, I installed a lightweight library to parse my blog's RSS feed:

```plain text
$ composer require dg/rss-php
```

I then added a

```plain text
posts.php
```

file at the root of the project, with the following content:

Nothing too complicated here – Composer's autoload is required at the top, allowing me to load the RSS parser to generate a list of blog posts as a string, in Markdown format.

The existing content of the

```plain text
README
```

file is then loaded into the

```plain text
$content
```

variable, and the Markdown string is inserted between its

```plain text
<!-- posts -->
```

and

```plain text
<!-- /posts -->
```

tags with

```plain text
preg_replace
```

.

Finally, the file's entire content is replaced with the new one, using the

```plain text
file_put_contents
```

function.

## The GitHub action

GitHub Actions are a fairly recent addition to GitHub, allowing developers to automate various CI/CD tasks, like running test suites or deploying web services.

They must be defined using YAML format in a

```plain text
.github/workflows
```

folder at the root of the project, and contain a list of steps that they are to execute.

Here's mine, which I named

```plain text
posts.yml
```

:

Again, nothing too complicated. We first give the action a name, and then define a list of events that should trigger it – pushing some code to the repository, a manual trigger from the interface (

```plain text
workflow_dispatch
```

), or periodically like a cron job (here, every day at midnight).

We then indicate that the action should run on an Ubuntu image, where it will:

- clone the repository;
- instal PHP 7.4;
- instal the Composer dependencies;
- run the PHP script;
- commit and push the changes, if any.

That's it! My GitHub profile is now automatically updated every time I publish a new article.

## Conclusion

This was a quick experiment aiming at exploring GitHub Actions, which I expect to use more and more in the future. It was also fun to use PHP as a simple scripting language again, in a procedural way.

I've voluntarily left out a few things in order to keep this article short and simple – please refer to the repository for implementation details.

This story was originally published on tech.osteel.me.

## Resources

- Managing your profile README
- GitHub Actions
- This article’s repository
- GitHub Action: Checkout V2
- GitHub Action: Setup PHP
- GitHub Action: Git Auto Commit
