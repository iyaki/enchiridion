---
title: "Modern front-end monorepos"
notion_id: 15085d92-272a-44fd-917f-a827799bc00b
notion_url: https://app.notion.com/p/Modern-front-end-monorepos-15085d92272a44fd917fa827799bc00b
last_edited: 2026-09-21T17:05:00.000Z
source_url: https://app.notion.com/p/Modern-front-end-monorepos-15085d92272a44fd917fa827799bc00b
tags: ["English", "Monorepositories", "Frontend", "Article", "madewithlove Blog"]
---
# Índice

<!-- unsupported block: table_of_contents -->

# [Managing dependencies and sharing code](https://madewithlove.com/blog/software-engineering/modern-front-end-monorepos-part-1-managing-dependencies-and-sharing-code-2/)

In short

In this series, we will explore setting up modern monorepos for the front-end that contain multiple packages and/or applications.

I’ve really learned to love a good monorepo setup, a repository that contains multiple packages and/or applications.

Being able to make changes across applications or packages in 1 pull request (PR), having the option to centralize and reuse code over applications, and unifying documentation and processes greatly simplifies the daily task of finding your way across multiple projects.

A monorepo comes with its own level of complexity and prompts questions such as:

- How do we handle dependencies?
- How do we run scripts?
- Where do we define build pipelines?
- How do we publish and manage packages?
- How do we deploy our applications?

A lot of questions.

Luckily, tooling has improved leaps and bounds over the last couple of years. Features like Yarn workspaces and tools like [Lerna](https://github.com/lerna/lerna), [NX](https://nx.dev/), and [Turborepo](https://turborepo.org/) (which was [acquired by Vercel](https://vercel.com/blog/vercel-acquires-turborepo)) are in a pretty stable phase and can help you address those issues.

**There has never been a better time to jump in than now.**

Over the course of multiple articles, I’ll guide you through the process of setting up a monorepo using mentioned tools. Let’s start with managing dependencies and sharing code in part 1.

## 1. Initializing a new monorepo with yarn

Create a folder called ‘monorepo-101’, ‘**cd**‘ into it, run ‘**yarn init**‘, then walk through the initialization steps.

Initialize a git repository by running ‘**git init**‘ and make sure you have a ‘.gitignore’ file containing (at least) the following lines.

**.gitignore**

```plain text
.next
node_modules
yarn-error.log
```

This step is important if we want to add Turborepo later. Create the following folder structure.

```plain text
/apps
    /admin
    /client
```

We will create Next.js applications in ‘admin’ and ‘client’; each application will have its own ‘package.json’, listing the dependencies.

Usually, this would resolve in multiple lock files and ‘node_modules’ folders. By enabling [Yarn workspaces](https://classic.yarnpkg.com/lang/en/docs/workspaces/), Yarn moves all dependencies to the root ‘node_modules’ folder, manages a single lockfile, and deduplicates dependencies if possible.

Add the following fields.

**package.json**

```plain text
...

"private": true,
"workspaces": [
    "apps/*",
],

...
```

The ‘private’ field indicates that you can’t publish this package to a registry. It’s required for Yarn workspaces to work and acts as a safeguard.

‘**cd**‘ into the ‘client’ directory and run ‘**yarn init**‘. Install the packages we need by running ‘**yarn react react-dom next**‘.

You’ll notice that the ‘node_modules’ folder remains empty (there might be a ‘.bin’ folder, but it shouldn’t contain any dependencies). When installing dependencies in a specific workspace, Yarn moves them to the root of your repository and manages the lockfile there.

You can install dependencies in the root of a workspace, too. We might want to install a dependency that is not tied to a specific workspace. Prettier is a great example.

Install prettier as a root dependency by adding the ‘-W’ flag: ‘**yarn add prettier -W -D**‘ (we’ve also added the -D flag because it’s also a devDependency).

Let’s verify that everything works as expected. Navigate back to the root of the project and remove all installed dependencies by running ‘**rm -rf node_modules**‘.

Now, run ‘**yarn**‘ (in the root of the repository), Yarn will go over all workspaces and install the dependencies, based on the lockfile.

Nice!

## 2. Set up two applications in the same monorepo

Let’s set up some applications. From this point onwards, I’m assuming you have some (very basic) knowledge of Next.js. If not, you can follow their excellent [getting started](https://nextjs.org/docs/getting-started) tutorial.

Run ‘**yarn init**‘ in ‘apps/client’, walk through the steps, and create a homepage that renders a header.

**apps/client/pages/index.js**

```plain text
const Home = () => {
  return <h2>Hello client</h2>
}
export default Home;
```

For convenience, let’s add a ‘dev’ script in the generated ‘package.json’. We’re also defining a port to have some control over which port links to which application.

**apps/client/package.json**

```plain text
...

"scripts": {
    "dev": "next dev -p 3000"
},

...
```

You can copy ‘package.json’ and the ‘pages’ folder to the ‘admin’ workspace to create the admin application.

Make sure to edit the following items after doing so:

- the name field in ‘package.json’ (to “admin”)
- the copy in ‘pages/index.js’ (to “hello admin”)
- the port in the dev script (to “3001”)

Run the applications by running ‘**yarn dev**‘ in both workspaces. You’ll need two terminal windows, for now; we’ll revisit this when adding Turborepo.

You should be able to visit the client application on [http://localhost:3000](http://localhost:3000/) and the admin application on [http://localhost:3001](http://localhost:3001/).

Yarn can also help you with sharing code (as a package) over applications.

Let’s introduce a UI library to create some visual consistency. Create a ‘packages’ folder in the root of your repository, and add it to the workspaces array.

**package.json**

```plain text
...

"workspaces": [
    ...
    "packages/*",
    ...
],

...
```

Create a ‘ui’ workspace (folder) in the ‘packages’ folder and run ‘**yarn init**‘ there.

```plain text
/packages
    /ui
```

Let’s kick off our UI library with a simple Header component.

**packages/ui/Header.js**

```plain text
const Header = ({ children }) => {
  return <h2>hello {children}</h2>;
};

export { Header };
```

Export all named exports from our Header module in an index file.

**packages/ui/index.js**

And, finally, edit ‘package.json’ to link everything up.

**packages/ui/package.json**

## 3. Create a package by transpiling modules

We’ll be using [next-transpile-modules](https://github.com/martpie/next-transpile-modules) to avoid having a Babel setup in our packages. If we want to publish this package to a registry (NPM) later, we might have to reconsider this (which we will when adding Lerna).

Add ‘next-transpile-modules’ as a devDependency in BOTH workspaces running.

```plain text
yarn add -D next-transpile-modules
```

Create a Next.js configuration file in both applications, and configure ‘next-transpile-modules’ to transpile the ‘ui’ package when needed.

**packages/client/next.config.js**

**packages/admin/next.config.js**

```plain text
const withTranspileModules = require('next-transpile-modules')(['ui']);

module.exports = withTranspileModules({
  reactStrictMode: true,
});
```

Run **yarn** in the root of our repository to link everything up, and you should be good to go. Now we can use this package in our applications by referring to it as ‘ui’.

Finally, import the Header component and use it on our applications.

**apps/client/pages/index.js**

```plain text
import { Header } from 'ui';

const Home = () => {
  return <Header>client</Header>
}

export default Home;
```

Use the same component in our admin application. When making changes to the component, both applications will refresh.

You’ve successfully created a monorepo with 2 apps and a package. In the next article, we’ll create build pipelines and enable caching by adding Turborepo. Here’s some confetti to celebrate.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# [Running scripts](https://madewithlove.com/blog/software-engineering/modern-front-end-monorepos-part-2-running-scripts/)

In the [previous article](https://madewithlove.com/blog/software-engineering/modern-front-end-monorepos-part-1-managing-dependencies-and-sharing-code-2/), we set up a monorepo project with 2 applications and a package. We also enabled [Yarn workspaces](https://classic.yarnpkg.com/lang/en/docs/workspaces/) to manage the dependencies and linking of the workspaces.

One inconvenience so far was having to start 2 separate terminal windows to run the applications. This isn’t a huge issue at the moment, but it doesn’t scale well. Let’s introduce [Turborepo](https://turborepo.org/) to handle this; it also adds some extra benefits I’ll talk about later in this article.

## 1. Start the monorepo with just 1 command

Install Turborepo as a root devDependency.

```plain text
yarn add turbo -D -W
```

And add .turbo to our ‘.gitignore’ file.

**.gitignore**

Let’s start with the ‘**dev**‘ script. Add a dev script in our root ‘package.json’ that uses Turborepo to run all workspace dev scripts.

**package.json**

Run ‘**yarn dev**‘ in the root of your project, you’ll get the following error.

```plain text
Could not find turbo.json. Follow directions at https://turborepo.org/docs/getting-started to create one
```

Let’s create a configuration file (called ‘turbo.json’) containing an empty object for now.

**turbo.json**

Run ‘**yarn dev**‘, again.

```plain text
task `dev` not found in turbo pipeline in package.json. Are you sure you added it?
```

We’re making progress. Turborepo is looking for a dev pipeline, one that is currently not defined in our configuration file. Let’s create a minimal configuration file.

**turbo.json**

```plain text
{
  "$schema": "https://turborepo.org/schema.json",
  "baseBranch": "origin/main",
  "pipeline": {
    "dev": {}
  }
}
```

We’re defining a pipeline called ‘dev’ which currently has no additional configuration. We’re also providing a link to the schema (for better IDE integration) and provide a baseBranch.

Run ‘**yarn dev**‘, again.

This time, Turborepo runs all dev scripts in the workspaces defined in ‘package.json’. If a workspace has no ‘dev’ script, it is ignored.

There is a warning, though, which is easy to fix.

```plain text
Did not find "packageManager" in your package.json. Please run "npx @turbo/codemod add-package-manager"
```

Run the suggested command to fill in the ‘**packageManager**‘ key in your root ‘package.json’. Commit or stash your changes before doing so or run it with the –force flag. This key will help Turborepo identify the package manager you’re currently using.

## 2. Run scripts for one application in a monorepo

Sometimes, you only want to run scripts in a specific workspace. Define the scope using the flag. Here, ‘admin’ is the name of the application (package.json name property) you want to target.

```plain text
yarn dev --scope=admin
```

Let’s add an extra build script. Add build scripts in the ‘package.json’ file of each application (‘admin’ & ‘client’).

**apps/client/package.json**

Do the same for the admin application (apps/admin/package.json); make sure you define a different port (3001).

Add a build script in the root ‘package.json’

**package.json**

```plain text
"build": "turbo run build"
```

Define a pipeline in our configuration file. Since the build script is only dependent on its own dependencies, we define the dependsOn as ‘^build’. Since Next.js outputs a ‘.next’ folder when creating a build, make sure you mark it as the output of the ‘build’ pipeline.

**turbo.json**

Trigger the build process by running ‘**yarn build**‘. Wait for it to finish. Now run it again. Something interesting happened.

Since we didn’t make any changes between the previous and the current build, Turborepo replays the cache and goes into **“FULL TURBO”** mode, saving you some precious time. We will make use of this mechanism (through remote caching) in our deployment flow later.

## 3. Configure Turborepo cache settings

For some pipelines, we can disable caching. There is no need to cache our dev output, let’s disable it.

**turbo.json**

Now we have two scripts, one that runs our applications in development mode and another that builds our applications.

With this knowledge, try to set up a ‘start’ pipeline. ‘**next start**‘ runs the production build process for a single Next.js project, so we need to make sure it is built first.

Make sure you:

- add the correct scripts to the applications
- create a script in the root ‘package.json’
- edit ‘turbo.json’ (what should ‘dependsOn’ contain?)

One final thing we can enable is [remote caching](https://turborepo.org/docs/features/remote-caching). This will make it possible to share the Turborepo cache between several developers. This is a beta feature that is currently supported by [Vercel](https://turborepo.org/docs/features/remote-caching#vercel).

Login to Vercel using ‘**npx turbo login**‘ and link your cache via ‘**npx turbo link**‘.

Let’s verify if this works; run a build (‘**yarn build**‘). You should spot ‘**Remote computation caching enabled (experimental)**‘ in the log.

Remove your local cache by running

```plain text
rm -rf ./node_modules/.cache/turbo
```

And run your build again. Now it should use the remote cache. You can also set up your own caching server (more information [here](https://turborepo.org/docs/features/remote-caching#custom-remote-caches)).

That’s it for now. You’ve successfully configured your scripts in a monorepo. In the next article, I’ll show you how we deploy the applications in our monorepo to Vercel using our current scripts whilst leveraging the remote cache.

Here are some party balloons to celebrate.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->
