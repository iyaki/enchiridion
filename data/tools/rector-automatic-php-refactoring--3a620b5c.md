---
title: "rector - automatic PHP refactoring"
notion_id: 3a620b5c-74a6-43d7-8b3b-5c6d94f12594
notion_url: https://app.notion.com/p/rector-automatic-PHP-refactoring-3a620b5c74a643d78b3b5c6d94f12594
last_edited: 2026-09-21T17:21:00.000Z
source_url: https://getrector.com/
tags: ["English", "PHP", "Tool", "Framework/Library"]
---
[https://getrector.org/](https://getrector.org/)

## Upgrade your PHP codebase & your mindset.

We help successful, growing companies get the most out of the code they already have. Reduce maintenance cost, make feature delivery cheaper, and turn legacy into a foundation you can build on.

Legacy After Rector src/Order.php

1<?php

2

3class Order

4{

5 /** @var Item[] */

6 private $items = array();

7

8 public function hasDiscount()

9 {

10 foreach ($this->items as $item) {

11 if ($item->discount > 0) {

12 return true;

13 }

14 }

15 return false;

16 }

17}

1<?php

2

3declare(strict_types=1);

4

5final readonly class Order

6{

7 /** @param Item[] $items */

8 public function __construct(

9 private array $items = [],

10 ) {}

11

12 public function hasDiscount(): bool

13 {

14 return array_any(

15 $this->items,

16 fn (Item $item): bool => $item->discount > 0,

17 );

18 }

19}

Clean, reviewable diffs

Trusted by engineers at

## A PHP tool you run on any project for an instant upgrade or automated refactoring

In the hands of an expert, Rector massively reduces work-time. An upgrade that would take 3 years takes us 6 months.

↑

### PHP & framework upgrades

Move from PHP 5.3 to 8.5, or jump major Symfony, Laravel and CakePHP versions - safely and incrementally.

⇆

### In-house framework migrations

Retire a custom framework and move to Symfony or Laravel, one safe step at a time.

✦

### Higher code quality

Ship features faster than the competition with a codebase your team actually wants to work in.

★

### Delegate the hard upgrade, powered by our Go engine

A big legacy jump you'd rather not own - Symfony 2, Zend 1, CakePHP 2? Our team takes the whole upgrade off your plate while you keep shipping features. We run it on Reco, our internal Go engine that applies the same rules up to 20× faster than Rector.

## Legacy stacks we've upgraded

No version is too old. We've taken projects off frameworks that were declared dead years ago and brought them to the latest stable release.

### Migrated to Symfony

CakePHP CodeIgniter Zend Nette PhalconPHP Yii

## Every change is a reviewable diff - strict, repeatable, no black boxes

Rector applies battle-tested rules and hands back a clean, line-by-line diff your team reviews like any other pull request. From type coverage to dead code, modern PHP syntax to privatization - every change is small, explainable and safe to merge.

Types Dead code PHP Privatization

final class Product

{

- private $name;

- private $price;

+ private string $name;

+ private float $price;

- public function discounted($rate)

+ public function discounted(float $rate): float

{

return $this->price - $this->price * $rate;

}

- public function label()

+ public function label(): ?string

{

return $this->name;

}

}

public function total(): int

{

- $legacy = $this->oldFlag;

$sum = 0;

foreach ($this->items as $item) {

- if (false) {

- continue;

- }

$sum += $item->price;

}

return $sum;

-

- $this->log('unreachable');

}

public function rank(string $type, ?string $note): int

{

- if (strpos((string) $note, 'urgent') !== false) {

+ if (str_contains((string) $note, 'urgent')) {

return 9;

}

- switch ($type) {

- case 'low': return 1;

- case 'high': return 3;

- default: return 0;

- }

+ return match ($type) {

+ 'low' => 1,

+ 'high' => 3,

+ default => 0,

+ };

}

-class Report

+final class Report

{

- public $cache = [];

+ private array $cache = [];

- public function __construct($name)

+ public function __construct(private readonly string $name)

{

- $this->name = $name;

}

- public function buildRow(): string

+ private function buildRow(): string

{

// used only inside this class

}

}

## Let agents be creative. Let Rector be deterministic - and save the costs

Let your agents handle the creative, hard-to-figure-out problems - the high-complexity, ambiguous solutions. Rector handles all the mechanical, repeatable changes, with feedback in CI minutes on every commit - far cheaper than tokens.

Your agents

### Creative & ambiguous work

Hard-to-figure-out problems and high-complexity, judgment-heavy solutions - exactly where a model's reasoning earns its tokens.

Rector

### Predictable & repeatable changes

Every rename, type declaration and upgrade rule - applied the same way, every time, with a reviewable diff. No guesswork, no wasted runs.

## We upgrade your legacy code in parallel - your team keeps shipping features

01

### Intro analysis

A focused 2-week analysis maps your codebase, surfaces the risks, and delivers a detailed upgrade estimate.

02

### Automated upgrade

We run Rector with custom rules, working in parallel to your ongoing development - no slowdown.

03

### Self-sufficient team

We wire Rector into your CI and hand over best practices. Your next upgrade takes days, on your own.

## What CTOs say after we upgraded their legacy PHP

> “In my time as Head of software engineering at i6, I found in Rector the perfect core and enabling team to tackle technical debt, unblock critical upgrades and build strong foundations for leading aviation software project.”

Dom Graziano

Head of Software Engineering at i6 Group

> “I'm extremely pleased with the progress we are making.It's really come a long way.”

William Adam Gleiss

VP of Technology at aRes Travel

> “Tomas' deep expertise and tools-driven approach allowed us to modernize a significantly large legacy PHP codebase while teaching us best practices. Rector team worked around our schedule to ensure minimal disruption. I highly recommend Rector team for legacy PHP code modernization!”

Eric Molitor

CTO at Curve

> “Thanks to Rector, we were able to quite simply refactor the core of our API, which saved a lot of work that our developers would otherwise have to do manually.”

Milan Mimra

CTO at Spaceflow

> “From upgrading our legacy project and improving team's productivity, to faster and easier code reviews, Rector is in the center of our PHP ecosystem.”

Nathan Page

Technical Lead at EONX

## Frequently Asked Questions

Can you upgrade PHP 5.3 code?

Yes. PHP 5.3 or older, an open-source framework, an in-house one, or pure spaghetti - we've upgraded it before.

How long does an upgrade take?

Most upgrades land in 6-12 months. You get a detailed, project-specific estimate in the 2-week intro analysis.

Can we continue developing new features during the upgrade?

Absolutely. Our upgrade process runs parallel to your ongoing development, ensuring no slowdown in your business growth.

Will we need to re-hire your team for future upgrades?

No. Part of our work is to make your team self-sufficient. We get your code quality the highest possible level, get Rector to your CI working for you and then, next upgrade will be a matter of days on your own.

We have a custom framework we want get rid off. Can you migrate it to an open-source one?

Yes. We specialize in framework migration, mostly to Symfony or Laravel frameworks.

## Turn legacy into leverage

Start with a focused 2-week intro analysis - then watch the debt shrink release after release, no frozen sprints required.
