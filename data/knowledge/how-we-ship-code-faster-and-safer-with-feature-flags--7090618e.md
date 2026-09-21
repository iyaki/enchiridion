---
title: "How we ship code faster and safer with feature flags"
notion_id: 7090618e-4e39-4011-b629-c26aee837a9c
notion_url: https://app.notion.com/p/How-we-ship-code-faster-and-safer-with-feature-flags-7090618e4e394011b629c26aee837a9c
last_edited: 2026-09-21T17:41:00.000Z
source_url: https://github.blog/engineering/ship-code-faster-safer-feature-flags/
tags: ["English", "Programming", "Project Management", "Product Management", "System Design / Software Architecture", "Article", "Github Blog"]
---
[https://github.blog/2021-04-27-ship-code-faster-safer-feature-flags/](https://github.blog/2021-04-27-ship-code-faster-safer-feature-flags/)

At GitHub, we’re continually working to improve existing features and shipping new ones all the time. From our launch of GitHub Discussions to the release of manual approvals for GitHub Actions—in order to ship new features and improvements faster while lowering the risk in our deployments, we have a simple but powerful tool: feature flags.

## Reducing deployment risk

We deploy changes around the clock, and we need to keep the service running with no disruptions. For these reasons, deploying needs to be an almost risk-free process. Any problem that comes up during a deployment requires a rollback of the code. During this rollback period, users could be impacted and other deployments delayed; the more time that elapses the bigger the impact.

We use feature flags to reduce the risk of deploying to production. Any potentially risky change is put behind a feature flag in the code and then, when the deployment is done, we enable the feature flag to everyone or to a percentage of actors. This way we minimize the impact of the new changes, and if something goes wrong we can disable the feature flag completely in a matter of seconds without interrupting other deployments. This is fundamental to us for the following reasons:

- It isolates and reduces the deployment risk.
- We have the ability to disable changes in seconds without rolling back a deployment, which could take minutes.

What kind of things can go wrong during a deployment, where a feature flag can help? These are some examples:

- Changes in database queries can potentially make existing functionality slower.
- Changes in the logic on existing features can produce unexpected behavior, edge cases not considered, and more.
- Changes in how we store or process information, such as saving new columns in the database can accidentally start inserting invalid values or generate too many writes to the database.

## Working on new features

Feature flags are not only useful to reduce deployment risk when fixing or improving existing features, we also use them when developing new features!

Using feature flags allows us to work on features incrementally. Only staff members working on the project have the corresponding feature flag enabled, so they can see the new feature that is under development while other users cannot. We don’t use long-lived feature branches. Instead, feature flags allow us to work on small batches, which brings us many benefits:

- Small batches are easier to review by other engineers in pull requests.
- The smaller the change, the lesser the chance to get something wrong during the production deployment.
- Not using long-lived feature branches avoids lots of potential merge conflicts and clashes with other features under development.

In order to adopt this way of building new features, you need to do a little bit of additional planning upfront to divide the work. For example, a new feature typically needs new data models or changes in existing ones. If you create a pull request only for the data model changes, then you may get blocked until those changes are merged. There are a couple of things we usually do to prevent this:

- Create a main pull request. A spike, where you put the data model changes, and continue working on other areas that require changes: UI, API, background jobs, etc. Then, start extracting those changes to smaller pull requests that your peers can review.
- Alternatively, you can create an initial pull request, and then create new branches off of that first branch, even if it’s not merged to the main branch yet. If you need to make changes to the first branch, you’ll need to rebase the other branches.

This way of working could introduce some friction due to having to ask for reviews more frequently. Sometimes it can take days until a team reviews your pull request. To prevent this from being a blocker, we follow one or more of these strategies:

Keep working even if the changes are not merged into the main branch. Either on that “spike” pull request or branching off the branch that is waiting for reviews.

Have a first responder person in each team that is focused on reviewing pull requests quickly to unblock the progress of their own team or other teams. Sometimes if a critical pull request is ready to be deployed sooner rather than later, but has outstanding non-critical feedback waiting to be addressed, we do that in follow up pull requests.

## Testing features

We have to make sure that the new functionality works as expected while also maintaining the existing functionality. In order to do that, we have many mechanisms to enable or disable feature flags at all levels as follows:

- In our development environments, we can toggle feature flags from the command line.
- In automated tests, we can enable or disable feature flags in the code.
- In our CI, we have two different builds: one that runs with all feature flags disabled by default, and another one that runs with all feature flags enabled by default. This drastically reduces the chances of not covering most code paths in automated tests properly.
- In production, we can enable or disable feature flags in the query string of a request.

## Different shipping strategies

We already mentioned that we can flag particular actors, or we can enable a feature flag for a percentage of actors. Those are the most common options, but others include:

- Individual actors: We use this mechanism to flag employees working on a feature or to enable the flag to customers that are experimenting with a feature, or if they have a problem we are trying to fix.
- Staff shipping: When a feature is almost ready to be released, we first enable it for all GitHub staff, publish an internal post for awareness, and create an internal issue for gathering feedback or possible bugs.
- Early access maintainers or other beta groups: When we are about to release an important feature that will impact the workflows of open source software (OSS) maintainers or other types of users, we test the feature with a small group first. After some time, we may interview them and gather their feedback to confirm our hypothesis, and validate the implementation.
- Percentage of actors: As mentioned earlier, we can specify a percentage of actors for which to enable the feature flag. In this case when an actor is flagged, it remains flagged, unless the feature flag is rolled back. When the percentage of actors is changed, a message is sent to our deployments channel on Slack so other engineers are aware of the change.
- Dark shipping: This allows us to enable the feature flag for a percentage of calls. This is different from the previous mechanism because an actor can get the feature enabled in one call (one request, for example) but not in the next one. This mechanism is not meant for features that are visible to the users, but for internal changes, such as a performance improvement in a query.

All these strategies and the full administration of feature flags is done through a web UI. Not all staff members have access to this UI, but most engineers do. The interface allows us to manage the shipping status of a feature flag, creation of new ones, and deletion. Additionally, we can see a history of changes made to the feature flag: who made the change, when and what changes were made, and other metadata, such as which team owns it.

## What to flag

When we say that we can flag an actor, we mean that we can flag not only users, but other entities. In particular, we can flag users, organizations, teams, enterprises, repositories, or GitHub Apps. When writing the code that checks if the flag is enabled or not, we need to specify which actor has the flag enabled or not. These are some examples:

- If the flag is purely visual, we flag the user, and we check the current logged user. Examples of this could be: dark mode, layout or navigation changes, and new components in the UI that don’t depend on data changes.
- If the flag changes the way we store data, it is usually better to flag the repository for consistency for all users using the repository. For example, if we start storing timestamps for specific actions, we don’t want some users to produce the new timestamp while others do not.
- For API changes, we also flag GitHub Apps.
- Sometimes, especially for some risky changes, we create custom, context-specific actor types, where the existing actor types just wouldn’t have worked.

Sometimes we do more sophisticated checks, such as checking multiple actors or checking if a repository is public or not. For example, when we staff ship a feature it may make a new feature available for all repositories of an employee. If the feature must not be visible to other users, we want to enable the feature only for private repositories to prevent public repositories of an employee leaking the new feature.

## The cost of a feature flag

As we have seen, using feature flags extensively changes the way we work, and requires a bit of additional planning and coordination. But there’s also additional costs associated with feature flags that we need to keep in mind. First, we have a runtime cost. When we check if a feature flag is enabled to an actor, we need to first load the metadata information of the flag, which is stored in MySQL but cached in memcached. We are now considering having the feature metadata and the list of actors always in memory to reduce this overhead. There’s an exception for this: we consider some feature flags “large.” An example of this was GitHub Actions, a feature that we rolled out to tens of thousands of users every week. For these large feature flags, when they are still not fully enabled, we need to do a per-actor query to MySQL to check if the actor is in the list of enabled actors.

For runtime, there’s also a cost in the form of technical debt. Once a feature flag is fully enabled, it leaves a lot of dead code and old tests in the repository that need to be deleted. Usually we do this manually after a feature has been rolled out for a few days or weeks depending on the case. But we are now automating the process: we have created a script that can be either invoked locally or by triggering a workflow manually.

The script uses regular expressions to find usages of the feature flag using git grep. Then for the matched Ruby files, it modifies the code using rubocop-ast. It is capable of deleting code blocks, such as if/else statements, and modifying boolean expressions, in addition to reindenting the code afterwards. When the file is a test, it deletes the code that enables the feature flag. For tests that check the behavior of the application when the feature flag is disabled, it just leaves a comment and makes the test fail, since in many cases we don’t want to just delete the test, but simply update it. As a result, even when the script requires some intervention from an engineer, it drastically reduces the manual work required in the process.

When the script is run using a workflow dispatch, it also creates a new branch and a pull request. Since we do extensive use of CODEOWNERS, the pull request will have the proper teams and people as reviewers. In the near future, we want to automate this process even more by running the script automatically when we can consider a feature flag stale.

## Conclusion

Feature flags allow us to ship code faster with higher confidence. Using them we reduce deployment risk and make code reviews and development on new features easier, because we can work on small batches. There are also associated costs, but the benefits highly exceed them.

## Written by

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
