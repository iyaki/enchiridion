---
title: "Write GitHub Actions for Gitlab Too"
notion_id: cf1cafc7-9458-4101-a956-b2cb37bce03e
notion_url: https://app.notion.com/p/Write-GitHub-Actions-for-Gitlab-Too-cf1cafc794584101a956b2cb37bce03e
last_edited: 2023-01-13T16:42:00.000Z
source_url: https://tomasvotruba.com/blog/write-github-actions-for-gitlab-too/
tags: ["English", "DevOps", "Continuous Integration/Continuous Delivery", "Article", "Tomas Votruba Blog"]
---
In a recent post [How can We use GitHub Actions in Gitlab?](https://tomasvotruba.com/blog/how-can-we-use-github-actions-in-gitlab), we looked at the idea, how both services could **the use same CI recipe**. As a Gitlab CI user, you can use some GitHub Actions to do the work for you.

Today we look at **how to write such action** to provide an excellent developer experience for both.

> "If you want to go fast, go alone.

## Let's Go Together

5 years ago [I gave a talk](https://www.youtube.com/watch?v=D827D5ILfh8) [Czech only] about 2 frameworks. These 2 frameworks were two different groups that didn't like each other. Both frameworks were written in PHP, both MVC and both were used by PHP developers. I was wondering, why not go together?

My talk was about how to **convert this aversion into a healthy competition, friendship, and mutual learning**.

What movie does that resemble?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

I would love to see a similar story between "Old Gitlabhand" and "GitHubtou" :). While researching Gitlab and GitHub in this theme, I found a post from 2018 called [GitLab CI/CD for GitHub](https://blog.anoff.io/2018-03-30-gitlab-ci-for-github/) by Andreas Offenhauser. The idea in the post is amazingly simple.

Back in 2018, GitHub didn't have a proper CI yet. We used more matured Travis, Circle CI, etc. Andreas suggests, **we can use the best of both worlds** - hook in GitHub to Gitlab CI pipelines and let GitHub collect Gitlab CI feedback.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The rest is history, but the underlying philosophy goes beyond time.

## How they Place the Input?

We already know [how to write ](https://tomasvotruba.com/blog/how-can-we-use-github-actions-in-gitlab#2-from-pseudo-syntax-to-gitlab-ci-syntax)[`.gitlab-ci.yml`](https://tomasvotruba.com/blog/how-can-we-use-github-actions-in-gitlab#2-from-pseudo-syntax-to-gitlab-ci-syntax) so we can use GitHub Action. But how can we write GitHub Action without writing two scripts - one for GitHub and another for Gitlab?

First, we need to use the Docker approach. **The GitHub docs suggest using Docker with arguments**:

```plain text
docker run some-image $ARGUMENTS
```

Then work with arguments **by their order** - $1, $2 etc. In practise the GitHub workflow might look like this:

```plain text
# .github/workflows/monorepo_split.yaml
# ...
with:
    from-package: 'packages/easy-coding-standard'
    to-repository: 'https://github.com/symplify/easy-coding-standard'
```

The philosophy of **Gitlab is a bit different**. They are closer to Docker ideology, so instead of argument order, they promote ENV variables.

```plain text
# .gitlab-ci.yml
env:
    FROM_PACKAGE: "packages/easy-coding-standard"
    TO_REPOSITORY: "https://github.com/symplify/easy-coding-standard"
```

You can find this approach in [Postgres Docker](https://hub.docker.com/_/postgres), [Mysql Docker](https://hub.docker.com/_/mysql), and many others.

So Github script works with arguments order and Gitlab with ENV variables. That's a pickle. What now? Do we have to write two scripts to make our GitHub Action work for both?

## What is the Shared Way?

When I spoke about Nette and Symfony to both communities, I **focused on values they share** - active community, simple controller architecture, creative solutions in small packages. This way, we could find understanding from each other.

What would be the **shared** path here?

The default way of doing things suggests there are also alternative ways. You can use magic facades in Laravel by default, or you can use [constructor injection](https://tomasvotruba.com/blog/2019/03/04/how-to-turn-laravel-from-static-to-dependency-injection-in-one-day/) alternative. I was looking for such an alternative for a while in GitHub and Gitlab, so they **can bridge together**.

## Research, Explore, Doubt

After a couple of days of frustration, I came across [Create custom Github Action in 4 steps](https://www.philschmid.de/create-custom-github-action-in-4-steps) post.

There was one **very important sentence**:

- `inputs`: defines the input parameters you can pass into your bash script.
- You can access them with `$INPUT_{Variable}` in our example `$INPUT_POKEMON_ID`

So does that mean that this GitHub Action input:

```plain text
with:
    from-package: 'packages/easy-coding-standard'
    to-repository: 'https://github.com/symplify/easy-coding-standard'
```

is also:

```plain text
env:
    INPUT_FROM_PACKAGE: 'packages/easy-coding-standard'
    INPUT_TO_REPOSITORY: 'https://github.com/symplify/easy-coding-standard'
```

Can you see the shared pattern? GitHub is using ENV, Gitlab is using ENV...

**Heureka! We've found it!**

## How to Write it Once?

The final solution is straightforward:

- we name the input variables with the single name
- on GitHub, we prefixed them with `INPUT_`

```plain text
$env = getenv();

$ciPlatform = '...'; // detect via known ci-based ENV variables

$envPrefix = $ciPlatform === 'GITHUB' ? 'INPUT_' : '';

// shared input
$packageDirectory = $env[$envPrefix . 'FROM_DIRECTORY'];
$toRepository = $env[$envPrefix . 'TO_REPOSITORY'];
```

In the end, the Docker image [has one script for both](https://github.com/symplify/monorepo-split-github-action/pull/10). You would not even notice what CI service it was written for originally.

So next time you'll be writing a GitHub Action, **think of your friends in Gitlab and write a Docker image for them too**. Thank you.

Happy coding!
