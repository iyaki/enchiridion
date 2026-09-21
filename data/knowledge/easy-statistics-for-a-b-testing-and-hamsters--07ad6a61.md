---
title: "Easy statistics for A/B testing and hamsters"
notion_id: 07ad6a61-77bf-45e1-b176-d7a9def1b9be
notion_url: https://app.notion.com/p/Easy-statistics-for-A-B-testing-and-hamsters-07ad6a6177bf45e1b176d7a9def1b9be
last_edited: 2024-02-21T00:13:00.000Z
source_url: https://longform.asmartbear.com/ab-testing-statistics/
tags: ["English", "Product Management", "Entrepreneurship", "Project Management", "Article", "A Smart Bear: Longform"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

So you’ve got your AdWords test all set up: Will people go for the headline “Code Review Tools” or “Tools for Code Review?”

Gee they’re both so exciting! Who could choose! I know, I know, settle down. Welcome to A/B testing.

Anyway, the next day you have this result:

|  | Variant A  (“Code Review Tools”) | Variant B  (“Tools for Code Review”) |
| --- | --- | --- |
| **Clicks:** | 31 | 19 |

**Is this conclusive?** Has A won? Or should you let the test run longer? Or should you try completely different text?

**The answer matters.** If you wait too long between tests, you’re wasting time. If you don’t wait long enough for _statistically conclusive_ results, you might _think_ a variant is better and use that false assumption to create a new variant, and so forth, all on a wild goose chase! That’s not just a waste of time, it also prevents you from doing the _correct_ thing, which is to come up with a _completely new test_.

**Normally a formal statistical treatment would be too difficult, but I’m here to rescue you** with a statistically sound yet incredibly simple formula that determines whether your A/B test results really are significant.

I’ll get to it in a minute, but I can’t help but include a more entertaining example than AdWords. Meet Hammy the Hamster, the probably-biased-but-incredibly-lovable tester of organic produce:

In the movie, Hammy chooses the organic produce **8 times** and the conventional **4 times**. This is an A/B test, just like with AdWords… but healthier.

If you’re like me, you probably think “organic” is the clear-cut winner—after all Hammy chose it _twice as often_ as conventional veggies. But, as so often happens with probability and statistics, **you’d be wrong**.

That’s because human beings are notoriously bad at guessing these things from gut feel. Most people are more afraid of dying in a plane crash than a car crash, even though the latter is [1000x more likely](https://injuryfacts.nsc.org/all-injuries/preventable-death-overview/odds-of-dying/?utm_source=longform.asmartbear.com&utm_campaign=longform.asmartbear.com&utm_medium=post). On the other hand, we’re amazed when CNN “calls the election” for a governor with a mere 1% of the state ballots reporting in. We also [can’t distinguish between patterns and noise](https://longform.asmartbear.com/pattern-seeking-fallacy/).

Okay okay, we suck at math. So what’s the answer? Here’s the bit you’ve been waiting for1:

> 

1

## Determining whether an A/B test is statistically significant

1. Define NN as “the number of trials.”  For Hammy, N=8+4=12N=8+4=12  For AdWords, N=31+19=50N=31+19=50
2. Define DD as “half the difference between the ‘winner’ and the ’loser.’”  For Hammy, D=8−42=2D=28−4 =2   For AdWords, D=31−192=6D=231−19=6
3. The test result is statistically significant only if D2>ND2>N.  For Hammy, D2=4D2=4, which is _not_ bigger than 1212, so it is _not significant_.  For AdWords, D2=36D2=36, which is _not_ bigger than 5050, so it is _not significant_.

So your AdWords test isn’t statistically significant yet. But you let the test continue to run. The next day you find 31 more clicks for variant A, and 19 more clicks for B. Rerunning the test, the measured difference is now significant:

|  | Variant  A | Variant  B | NN | DD | D2D2 | Stat  Sig? |
| --- | --- | --- | --- | --- | --- | --- |
| **Day one:** | 31 | 19 | 50 | 6 | 36 | No |
| **Day two:** | 62 | 38 | 100 | 12 | 144 | Yes |

A lot of times, though, you keep running the test and it’s still not significant. That’s when you realize you’re not learning anything new; the variants you picked are not meaningfully different for your readers. That means it’s time to come up with something new.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

When you start applying the formula to real-world examples, you’ll notice that **when N is small, it is difficult—or even impossible—to be statistically significant**. For example, say you’ve got one ad with 6 clicks and the other with 1. That’s N=7;D=2.5;D2=6.25N=7;D=2.5;D2=6.25. So the test is still inconclusive, even though A is beating B six-to-one. Trust the math here—with only a few data points, you really don’t know anything yet.

The smaller the NN, the bigger the difference needs to be, to be detectable. Specifically, results are significant only if the ratio between `A` and `B` is larger than N+2N−2N−2N+2. So for example, if N=50N=50, as it was in our “day one A/B test” example, the winning variant needs to be almost double the number of clicks as the losing variant in order to have a detectable difference, e.g. a conversation rate of 10% versus 5%. This is a huge difference; it’s great if you find something so dramatically better, but this is rare, and therefore you can almost never find a significant A/B test given only 50 clicks to analyze.

When N=100N=100, the winner needs to be at least 50% higher than the loser (which it was by the second day in our example). It takes N=1800N=1800 to detect the case where the winner is only 10% larger than the loser.

And this is bad news for A/B tests, because often one variant isn’t better than the other by more than 10%, e.g. a “2.4% conversion rate” versus a “2.2% conversion rate.” What does this mean, especially if you don’t have large N? It means **you need to be seeking big differences**, not subtle ones. Test wildly different designs, rather than tweaks. Tweaks can only be tested when NN is enormous.

I hope this formula will help you make the right choices when running A/B tests. It’s simple enough that you have no excuse not to apply it! Human intuition sucks when it comes to these things, and A/B testing tools often use misleading or incorrect math, so let this formula help you draw the right conclusions.

_**☞ If you're enjoying this, please **_[_**subscribe**_](https://longform.asmartbear.com/subscribe/)_** and share this article! ☜**_

### [Subscribe](https://longform.asmartbear.com/subscribe) for more, andThank you for sharing!

## Endnote for the mathematically inclined: The derivation

The [null-hypothesis](http://en.wikipedia.org/wiki/Null_hypothesis?utm_source=longform.asmartbear.com&utm_campaign=longform.asmartbear.com&utm_medium=post) is that the results of the A/B test are due to chance alone. The statistical test we need is [Pearson’s chi-squared](http://en.wikipedia.org/wiki/Pearson%27s_chi-square_test?utm_source=longform.asmartbear.com&utm_campaign=longform.asmartbear.com&utm_medium=post)2.

> 

2

The definition of the χ2χ2 statistic follows, where:

mm = number of possible outcomes;

OkOk = observed quantity of results in category kk;

EkEk = expected quantity of results in category kk:

χ2=∑k=1m(Ok−Ek)2Ekχ2=k=1∑mEk(Ok−Ek)2

In the simple case of a two-variant A/B test, _m_ = 2. O1O1 and O2O2 are the observed results, and definitionally N=O1+O2N=O1+O2. The expected result under the null-hypothesis is that the quantities fall equally into each category, therefore E1=E2=N/2E1=E2=N/2.

Plugging this into the definition:

χ2=(O1−N2)2N2+(O2−N2)2N2χ2=2N(O1−2N)2+2N(O2−2N)2

The first numerator can be rewritten in terms of O1O1 and O2O2 by substituting N=O1+O2N=O1+O2, and this results in our variable D2D2 as defined in the main text:

(O1−N2)2=(O1−O1+O22)2=(2O1−O1−O22)2=(O1−O22)2=D2(O1−2N)2=(O1−2O1+O2)2=(22O1−O1−O2)2=(2O1−O2)2=D2

We can repeat with the second numerator, and so the expression simplifies:

χ2=(O1−N2)2N2+(O2−N2)2N2=D2N2+D2N2=2N(2D2)=4D2Nχ2=2N(O1−2N)2+2N(O2−2N)2=2ND2+2ND2=N2(2D2)=N4D2

Now that we have a simple formula for the chi-squared statistic, we refer to the chi-squared distribution to determine statistical significance. Specifically: What is the probability this result would have happened by chance alone?

Looking at [the distribution](http://www.itl.nist.gov/div898/handbook/eda/section3/eda3674.htm?utm_source=longform.asmartbear.com&utm_campaign=longform.asmartbear.com&utm_medium=post) at 1 degree of freedom, we must exceed 3.8 for 95% confidence and 6.6 for 99% confidence. For this simplified rule-of-thumb formula, I selected 4 as the critical threshold. Solving for D2D2 completes the derivation:

χ2>4χ2>4 4D2N>4N4D2>4 D2>ND2>N

□

(And if D2D2 is more than double NN, you’re well past the 99% confidence level.)

Deriving the other statement in the article—that the ratio between the two variants needs to exceed a certain threshold to be significant—start with the boundary condition of being significant, and derive the values of AA and BB in that case:
