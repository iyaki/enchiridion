---
title: "Scripts to Rule Them All"
notion_id: 7565fc40-81c8-45c1-8a7c-8e73c118fb35
notion_url: https://app.notion.com/p/Scripts-to-Rule-Them-All-7565fc4081c845c18a7c8e73c118fb35
last_edited: 2023-04-25T14:00:00.000Z
source_url: https://github.blog/2015-06-30-scripts-to-rule-them-all/
tags: ["English", "Programming", "DevOps", "Article", "Github Blog"]
---
[https://github.blog/2015-06-30-scripts-to-rule-them-all/](https://github.blog/2015-06-30-scripts-to-rule-them-all/)

At GitHub we have a lot of software running our product and company. We also have a lot of potential contributing members. Being able to get from git clone to an up-and-running project in a development environment is imperative for fast, reliable contributions. A consistent bootstrapping experience across all our projects reduces friction and encourages contribution.

With practically every software project, developers need to perform the following tasks:

- bootstrap
- run tests
- run continuous integration
- start app

At GitHub, we have normalized on a set of script names for all of our projects that individual contributors will be familiar with the second after they clone a project. We call them “Scripts to Rule Them All”.

Here’s a quick mapping of what our scripts are named and what they’re responsible for doing:

- script/bootstrap – installs/updates all dependencies
- script/setup – sets up a project to be used for the first time
- script/update – updates a project to run at its current version
- script/server – starts app
- script/test – runs tests
- script/cibuild – invoked by continuous integration servers to run tests
- script/console – opens a console

Each of these scripts is responsible for a unit of work. This makes them composable, easy to call from other scripts, and easy to understand.

For example, script/bootstrap is only responsible for dependency management. script/setup initially calls script/bootstrap and then has its own code that will set a project to an initial state. script/test can be called on its own in a development environment to run tests, but is also called by script/cibuild by our CI server. script/console will load a development console on its own, but if appended with an environment name, will load the console for that environment.

Another advantage of consistent script names is language agnosticism. This means the scripts themselves can be written in whichever language is appropriate for the maintainers or the project. It also means the conventions can work for projects of varying languages and frameworks. This ensures an individual contributor can do things like bootstrap or run tests without knowing how to do them for a wide range of project types.

For example, script/bootstrap might call bundle install, npm install, carthage bootstrap, or git submodule update.

Normalizing on script names not only minimizes duplicated effort, it means contributors can do the things they need to do without having an extensive fundamental knowledge of how the project works. Lowering friction like this is key to faster and happier contributions.

### Read More

We’ve created github/scripts-to-rule-them-all as a home for this pattern. In this repository, you’ll find working examples of scripts that make use of this technique as well as more detailed responsibilities of each script.

## Written by

## Related posts

## Explore more from GitHub

### Docs

Everything you need to master GitHub, all in one place.

Go to Docs

### GitHub

Build what’s next on GitHub, the place for anyone from anywhere to build anything.

Start building

### Customer stories

Meet the companies and engineering teams that build with GitHub.

Learn more

### GitHub Universe 2026

Join us October 28-29 in San Francisco or online for GitHub Universe, our flagship developer event uniting people, agents, and the world’s code.

Register now
