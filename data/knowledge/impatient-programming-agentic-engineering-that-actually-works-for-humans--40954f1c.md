---
title: "Impatient Programming · Agentic engineering that actually works for humans."
notion_id: 40954f1c-7d23-81d7-a9f4-cdd9c1d89627
notion_url: https://app.notion.com/p/Impatient-Programming-Agentic-engineering-that-actually-works-for-humans-40954f1c7d2381d7a9f4cdd9c1d89627
last_edited: 2026-10-30T04:19:00.000Z
source_url: https://impatientprogramming.org/
tags: ["English", "Agile", "Product Management", "Software Development", "Decision Making", "Systems Design / Software Architecture", "Programming", "Article", "Guide", "impatientprogramming.org"]
---
Agentic engineering that actually works for humans.

A method for engineering teams that build with coding agents. It keeps the values of Extreme Programming (XP) and rewrites its twenty-nine rules for a world where typing is nearly free and the hard part is keeping control of the code.

[Read the intro](https://impatientprogramming.org/xp-revisited/)[Watch the explainer](https://impatientprogramming.org/#explainer)

![image](https://impatientprogramming.org/assets/huff.svg)

## The conditions that birthed XP are back

> Our industry is kind of like the cartoon character who ran off the cliff and hasn't looked down yet.

### Big Design Up Front

Y2K Teams doing "waterfall" spent months on requirements that were already out of date by the time they were tossed "over the wall" to the programmers.

2026 The new BDUF is machine-generated plans that nobody reads. **Confluslop.** Waterfall development for the TikTok generation.

### No process at all

Y2K Most software in the 90s and 2000s got built with _no process at all_. It was the age of cowboys.

2026 The new wild west of no-software-process-development is nothing other than **vibe coding**.

The good news is that despite the confusion everywhere, XP can inspire _what comes next_. We call it **Impatient Programming**. IMP for short. Practitioners are Imps, and we keep our clankers on tight leashes.

## How an IMP team works

Six of the twenty-nine rules, each in two sentences.

[All twenty-nine rules](https://impatientprogramming.org/rules/)

Impatient Programming in two minutes.

[Rules 13 and 20](https://impatientprogramming.org/rules/#rule-13)

### Agree on the Contract

As you work out what the feature must do, the agent writes each obligation into the Contract. On the exhale, you make sure it's complete.

[Rule 12](https://impatientprogramming.org/rules/#rule-12)

### Breathe in, breathe out

On the inhale, the agent sprawls until the feature shows what it wants to be. On the exhale, you fold what you learned back into what already exists, in the same PR.

[Rule 26](https://impatientprogramming.org/rules/#rule-26)

### Make every test earn its place

Agents will happily write thousands of automated tests along with your code. Imps run them through a kill matrix using mutation testing, and keep a test only if it catches something no other test catches.

[Rule 27](https://impatientprogramming.org/rules/#rule-27)

### Let green merge itself

As many pull requests as possible ship on automated checks and adversarial model review. Human attention is saved for the reviews that guard against catastrophic risks.

[Rule 24](https://impatientprogramming.org/rules/#rule-24)

### Build guardrails

A guardrail is software, and building one is programming. Think of bowling with kids, when you put the bumpers up so the ball reaches the pins.

[Rule 4](https://impatientprogramming.org/rules/#rule-4)

### Show finished work every week

IMP has no timeboxed iterations, and the rework happens inside every pull request. Once a week the whole company holds a showcase, and presenters can only show something they've finished.

One pull request breathing in and out above the Contract
    
    
    
  **Breathe in.** The agent sprawls until the feature shows what it wants to be.
    **Breathe out.** You fold what you learned back into what already exists.
    **Merge.** Same pull request, with the Contract underneath the whole time.

![image](https://impatientprogramming.org/assets/img/zar-logo.webp)

## Born at ZAR

XP was born on Chrysler's C3 payroll project. IMP was born at ZAR, where Obie is CTO. We give people in places like Pakistan, Bangladesh, and Nigeria access to dollars, debit cards, stocks and gold. We're a financial technology company. People depend on what we build, so it has to be right.

IMP describes how our engineering team already works, rule by rule.

[zarpay.app](https://zarpay.app/)

[Rule 4](https://impatientprogramming.org/rules/#rule-4) The weekly showcase
            Once a week the whole company gathers and shows what shipped. Nothing in progress counts, and product work is demoed in production. People outside engineering present too.

[Rule 2](https://impatientprogramming.org/rules/#rule-2) One shot a week
            Each engineer names one focus for the week, with something to demo by Friday.

[Rule 18](https://impatientprogramming.org/rules/#rule-18) The customer codes
            The whole company is expected to code. Our CEO isn't a programmer, and some of our best systems started as prototypes he built himself.

[Rule 6](https://impatientprogramming.org/rules/#rule-6) Agents in the open
            People and agents work side by side in public chat channels. Every Claude Code session feeds a shared knowledge graph, so what one session learns, the whole team knows.

[Rules 19 and 24](https://impatientprogramming.org/rules/#rule-19) Guardrails in machines
            Our conventions are written down, and machines check each change against them, from lint rules up to model reviewers.

[Rule 26](https://impatientprogramming.org/rules/#rule-26) Tests that earn their place
            Mutation testing runs in CI. A test that catches nothing of its own fails the build.
[Rules 13, 17 and 20](https://impatientprogramming.org/rules/#rule-13) The Contract
            In September 2026 Obie deleted components of Agora, our agent orchestration platform. Claude Code regenerated each one from a written intent, a set of decision records and an executable Contract, without access to the old code.
One obligation in that Contract says an operator's Stop landing mid-reboot wins. Writing it down exposed a bug in production, where a Stop clicked during a Reboot was getting lost. The new obligation went in first, and three examples went red against the unchanged code. Then the code changed to match.
[Read the case study](https://impatientprogramming.org/case-studies/born-at-zar/)

> Keeping the codebase pristine at each atomic change is so radical and so appealing.

## Twenty-nine rules, 1999 next to 2026

In 1999 Don Wells put the rules of Extreme Programming on a single web page. Obie went through all twenty-nine and asked the same question of each one. What does it look like when typing is nearly free and you have a team of agents at your side?

[Read all of them](https://impatientprogramming.org/rules/)

## Lazy, impatient and proud

The name comes from Larry Wall, one of Obie's earliest programming heroes. In Programming Perl he listed the three great virtues of a programmer: _laziness, impatience and hubris_.

[The values](https://impatientprogramming.org/values/)

Impatience kept honest by feedbackThis one is in the name. We're impatient. We want to ship as much as possible, as quickly as possible, so speed is a core value.


          Laziness kept honest by simplicityWall's lazy programmer goes to great lengths to save effort overall. In IMP that means putting an enormous amount of effort into continuous (and ideally automated!) guardrail engineering, so that no coding agent ever makes the same mistake more than once.


          Hubris kept honest by courage and respectWall's proud programmer writes code that other people won't want to say bad things about. With great AI powers comes great responsibility, to make sure that our messes do not derail others on the team.

XP's own values carry over untouched: communication, simplicity, feedback, courage and respect. Underneath them is the heart.

Obie's friend Jerry Weinberg said that no matter how it looks at first, it's always a people problem. We've found that's true, and that the reverse is true too. When we win, it's always because the people involved are amazing individuals.

## Where to start

[Read the intro](https://impatientprogramming.org/xp-revisited/)The long version of the argument, with the history, in Obie's words.

[Brief your agent](https://impatientprogramming.org/primer.md)If you work with a coding agent, give it the primer for agents at `/primer.md`.

[Get the newsletter](https://impatientprogramming.org/newsletter/)We write when a new essay or case study goes up, or when a rule changes.
