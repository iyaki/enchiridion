---
title: "Hosting all your PHP packages together in a monorepo"
notion_id: b0b1ee88-69e2-4575-9bfc-9f5b329ab32a
notion_url: https://app.notion.com/p/Hosting-all-your-PHP-packages-together-in-a-monorepo-b0b1ee8869e245759bfc9f5b329ab32a
last_edited: 2023-04-25T14:27:00.000Z
source_url: https://blog.logrocket.com/hosting-all-your-php-packages-together-in-a-monorepo/
tags: ["Article", "LogRocket Blog", "English", "Monorepositories", "PHP"]
---
[https://blog.logrocket.com/hosting-all-your-php-packages-together-in-a-monorepo/](https://blog.logrocket.com/hosting-all-your-php-packages-together-in-a-monorepo/)

When a PHP project gets big and complex, it becomes difficult to manage.

In this situation, we’d split the project into independent packages and use Composer to import all packages into the project. Then, different functionalities can be implemented and maintained by different teams and can be reused by other projects, too.

Composer uses the Packagist registry to distribute PHP packages. Packagist requires us to provide a repository URL when publishing a new package.

As a consequence, splitting a project into packages also affects how they are hosted: from a single repository hosting the whole code to a multitude of repositories to host the code for every package.

So we have solved the problem of managing the project’s code, but at the expense of creating a new problem: now we have to manage the hosting of the code.

## The problem with decentralized package hosting

Our packages will be versioned, and every version of the package will depend on some specific version of another package, which will itself depend on some other version of some other package, and so on.

This becomes a problem when submitting a pull request for your project; most likely, you will also need to modify the code in some package, so you need to create a new branch for that package and point to it in your composer.json.

Then, if that package depends on some other package that must also be modified, you need to create a new branch for it and update the first package’s composer.json to point to it.

And if that package depends on some other package… You get the point.

Then, once you approve the pull request, you need to undo all modifications in all composer.json files to point to the newly published version of the package.

This all becomes so difficult to achieve that you may quite likely stop altogether using feature branches and publish straight to master, so you won’t be able to track a change across packages. Then if, in the future, you need to revert the change, good luck finding all the pieces of code, across all packages, that were modified.

What can we do about it?

## Intro to the monorepo

This is where the monorepo comes to save the day. Instead of having our code distributed across a multitude of repositories, we can have all packages hosted in a single repository.

The monorepo allows us to version-control all our packages together, so that creating a new branch and submitting a pull request will be done in a single place, including the code for all packages that may be affected by it.

However, we’re still bound by the constraints of Packagist: for distribution purposes, every package needs to live under its own repository.

What do we do now?

### Addressing the Packagist constraints

The solution is to decouple development and distribution of the code:

- Use a monorepo to develop the code
- Use a multitude of repositories (one repo per package) to distribute it (the famous “[READ ONLY]” repos)

Then, we must keep all the source and distribution repositories synchronized.

When developing the code in the monorepo, after a new pull request is merged, the new code for every package must be copied over to its own repository, from which it can be distributed.

This is called splitting the monorepo.

## How to split the monorepo

A simple solution is to create a script using git subtree split and then synchronize the package code into its own repo.

A better solution is to use a tool to do exactly this so we can avoid doing it manually. There are several tools to choose from:

- Git Subtree Splitter (splitsh/lite)
- Git Subsplit (dflydev/git-subsplit)
- Monorepo builder (symplify/monorepo-builder)

From these, I’ve chosen to use the Monorepo builder because it’s written in PHP, so I can extend it with custom functionality. (In contrast, splitsh/lite is written in Go, and dflydev/git-subsplit is a Bash script.)

> N.B., the Monorepo builder works for PHP packages only. If you need to manage JavaScript packages or anything else, you must use another tool.

## Organizing the monorepo structure

You must create a structure to organize the code in the monorepo. In the simplest case, you can have a root packages/ folder and add each package there in its own subfolder.

If your code is more complex, containing not only packages, but also bundles, or contracts, or others, you can create a multilevel structure.

Symfony, for instance, uses the following structure in its monorepo symfony/symfony:

Symfony monorepo structure

In my own case, I only recently set up a monorepo to host all my projects together. (The reason being that I had a potential contributor who could not manage to set up the development environment, and so he went away 😢.)

My overall project encompasses multiple layers: the GraphQL API for WordPress plugin sits on top of the server GatoGraphQL, which sits on top of the framework PoP.

And while these are related, they are also independent: we can use PoP to power other applications, not only Gato GraphQL; and Gato GraphQL can power any CMS, not just WordPress.

Hence, my decision was to treat these as “layers,” where each layer might see and use another one, but not others.

When creating the monorepo structure, I replicated this idea by distributing the code over two levels: layers/ first, and only then packages/ (and, for one specific case, also plugins/):

PoP monorepo structure

Instead of creating a new repository, I decided to reuse the one from PoP, under leoloso/PoP, because it was the foundation of the whole code (and also because I didn’t want to lose the stars it had been given 😁).

Once you have defined the monorepo structure, you can migrate the code from each package’s repository.

## Importing code, including the Git history

If you’re starting the monorepo from scratch, you can run monorepo-builder init to set it up and also create a new repository for each of your new packages. Otherwise, if you have been developing your packages in their own repositories, you will need to port them over to the monorepo.

Most likely, when migrating the packages, you will also want to port their Git histories and commit hashes to keep browsing them as documentation and keep track of who did what, when, and why.

The Monorepo builder will not help you with this task. So, you need to use another tool:

- Multi- To Mono-repository (hraban/tomono)
- Shopsys Monorepo Tools (shopsys/monorepo-tools)

After you have migrated the code, you can start managing it with the Monorepo builder as explained in its README.

## A single composer.json to rule them all

Every PHP package has its own composer.json file defining what dependencies it has.

The monorepo will also have its own composer.json file, containing all the dependencies for all PHP packages. This way, we can run PHPUnit tests, PHPStan static analysis, or anything else for all code from all packages by executing a single command from the monorepo root.

For this, PHP packages must contain the same version for the same dependency! Then, if package A requires PHPUnit 7.5, and package B requires PHPUnit 9.3, it will not work.

Monorepo builder provides the following commands:

- monorepo-builder validate checks that dependencies in all composer.json do not conflict
- monorepo-builder merge extracts all dependencies (and other information) from all composer.json, and merges them into the monorepo’s own composer.json

What took me a bit of time to realize is that then, you must not manually edit the root composer.json! Because this file is automatically generated, you can lose your custom changes if they were not added via the tool’s configuration file.

Funnily enough, this is the case for dealing with the Monorepo builder itself. To install this library in your project, you can run composer require symplify/monorepo-builder --dev in the monorepo root, as usual. But immediately after, you should recreate the dependency in the config file monorepo-builder.php:

```plain text
return static function (ContainerConfigurator $containerConfigurator): void { $parameters = $containerConfigurator->parameters(); $parameters->set(Option::DATA_TO_APPEND, [ 'require-dev' => [ 'symplify/monorepo-builder' => '^9.0', ] ]); }
```

## Splitting the monorepo

So you have merged a pull request. Now it’s time to synchronize the new code into the package repositories. This is called splitting.

If you’re hosting your monorepo on GitHub, you can just create an action to be triggered on the push event of the master (or main) branch to execute the GitHub Action for Monorepo Split, indicating which is the source package directory and which repository to copy the contents to:

```plain text
name: 'Monorepo Split' on: push: branches: - master jobs: monorepo_split_test: runs-on: ubuntu-latest steps: - uses: actions/checkout@v2 with: fetch-depth: 0 - uses: "symplify/monorepo-split-github-action@master" env: GITHUB_TOKEN: ${{ secrets.ACCESS_TOKEN }} with: # ↓ split "packages/your-package-name" directory package-directory: 'packages/your-package-name' # ↓ into https://github.com/your-organization/your-package-name repository split-repository-organization: 'your-organization' split-repository-name: 'your-package-name' # ↓ the user signed under the split commit user-name: "your-github-username" user-email: "[email protected]"
```

To have this working, you also need to create a new access token with scopes “repo” and “workflow,” as explained here, and set up this token under secret ACCESS_TOKEN, as explained here.

The example above works for splitting a single package. How do we manage to split multiple packages? Do we have to declare a workflow for each of them?

Of course not. GitHub actions support defining a matrix of different job configurations. So we can define a matrix to launch many runner instances in parallel, with one runner per package to split:

```plain text
jobs: provide_packages_json: runs-on: ubuntu-latest steps: - uses: actions/checkout@v2 - uses: shivammathur/setup-php@v2 with: php-version: 7.4 coverage: none - uses: "ramsey/composer-install@v1" # get package json list - id: output_data run: echo "::set-output name=matrix::$(vendor/bin/monorepo-builder packages-json)" outputs: matrix: ${{ steps.output_data.outputs.matrix }} split_monorepo: needs: provide_packages_json runs-on: ubuntu-latest strategy: fail-fast: false matrix: package: ${{fromJson(needs.provide_packages_json.outputs.matrix)}} steps: - uses: actions/checkout@v2 - name: Monorepo Split of ${{ matrix.package }} uses: symplify/github-action-monorepo-split@master env: GITHUB_TOKEN: ${{ secrets.ACCESS_TOKEN }} with: package-directory: 'packages/${{ matrix.package }}' split-repository-organization: 'your-organization' split-repository-name: '${{ matrix.package }}' user-name: "your-github-username" user-email: "[email protected]"
```

Now, the package name is no longer hardcoded, but comes from the matrix (“the reality is, the spoon does not exist”).

Moreover, since the list of packages is provided via the monorepo-builder.php configuration file, we can just extract it from there. That is accomplished by executing the command vendor/bin/monorepo-builder packages-json, which produces a stringified JSON output containing all the packages:

Retrieving the list of packages.

## Releasing a new version (for all packages)

The monorepo is kept simple by versioning all packages together, using the same version for all of them. Thus, package A with version 0.7 will depend on package B with version 0.7, and so on.

This means we will be tagging packages even if no code has changed in them. For instance, if package A has been modified, it will be tagged as 0.7, but so will package B, even though it contains no modifications.

The Monorepo builder makes it very easy to tag all packages. We first need to have a workflow to split the monorepo whenever tagged (it’s basically the same workflow from above, plus passing the tag to symplify/github-action-monorepo-split).

Then, we tag the monorepo to version 0.7 by running this command:

```plain text
vendor/bin/monorepo-builder release "0.7"
```

Executing this command does real magic. It first releases the code for production:

- Bump mutual dependencies across packages to 0.7
- Tag the monorepo with 0.7
- Do a git push with tag 0.7

And then, it reverts the code back for development:

- Update the branch alias for dev-master in all packages to 0.8-dev
- Bump mutual dependencies to 0.8-dev
- Do a git push

Watching it in action never stops fascinating me. Check how, upon executing a command, the whole environment seems to take a life of its own:

## Removing workflows from packages

Even though we are running PHPUnit in our monorepo for all packages, we might still want to run PHPUnit on each package in its own repository after it has been split, if only to show a success badge.

However, we can’t do this anymore. Or at least, not so easily.

The fact that all packages are versioned together and released at the same time, and that the new release for every package takes a bit of time to become available on Packagist — say, five minutes — means that dependencies may not be available when running composer install, causing the PHPUnit workflow to fail.

For instance, if package A depends on package B, tagging them with version 0.3 means that package A’s version 0.3 will depend on package B’s version 0.3. However, because both are split and tagged at the same time, when package A runs an action triggered by pushing to master, package B’s version 0.3 won’t be available yet, and the workflow will fail.

In conclusion: you will need to remove running these workflows from every package’s repository, and rely only on the workflows from the monorepo.

Or, if you really want that success badge, find some hack for it (such as delaying 10 minutes the execution of the workflow).

## Conclusion

A monorepo helps manage the complexity of a big codebase. It makes it easy to maintain a coherent snapshot or state for the whole project, allows for submitting a pull request that involves code from multiple packages, and welcomes first-time contributors to set up the project without hiccups.

All these characteristics can be obtained using a multitude of repositories, too, but in practice, they are very difficult to execute.

A monorepo must itself be managed. Concerning PHP packages, we can do this through the Monorepo builder library. In this article we learned how to set up this tool, configure it, and release our packages with it.

## Get set up with LogRocket's modern error tracking in minutes:

1. Visit https://logrocket.com/signup/ to get an app ID
2. Install LogRocket via npm or script tag. LogRocket.init() must be called client-side, not server-side npm Script tag $ npm i --save logrocket // Code: import LogRocket from 'logrocket'; LogRocket.init('app/id'); // Add to your HTML: <script src="https://cdn.lr-ingest.com/LogRocket.min.js"></script> <script>window.LogRocket && window.LogRocket.init('app/id');</script>
3. (Optional) Install plugins for deeper integrations with your stack: Redux middleware NgRx middleware Vuex plugin

Get started now
