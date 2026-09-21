---
title: "Extremly defensive PHP"
notion_id: 5d3e70cf-e4af-4b49-9226-ef8e4bc80272
notion_url: https://app.notion.com/p/Extremly-defensive-PHP-5d3e70cfe4af4b499226ef8e4bc80272
last_edited: 2026-09-21T17:33:00.000Z
source_url: https://ocramius.github.io/extremely-defensive-php/#/
tags: ["English", "Programming", "PHP", "Slides"]
---
[https://ocramius.github.io/extremely-defensive-php/#/](https://ocramius.github.io/extremely-defensive-php/#/)

## Hi!

## I'm Marco!

## Ocramius

> @Ocramius @asgrim @RoaveTeam Imma start calling you Gandalphp. — Lee Davis (@leedavis81) 9 July 2015

I'm an Extremist.

I became an Extremist.

I maintain too many projects that allow too much.

Everyone wants a shiny new feature.

This talk is about aggressive practices.

These practices work well for long-lived projects.

Every OSS project should probably adopt them.

## Know your Limits

If you don't like these practices, don't apply them ...

... but please be considerate about what you release

## Let's start in a bold way

## You are dumb

## No, I do not mean You

(Sorry!)

## I am dumb

## We are dumb

## Everyone is dumb

Those who don't admit it, are the dumbest!

## Do you know Defensive Driving?

Defensive driving = assuming that other drivers will make mistakes

## Defensive coding is like Defensive Driving

## Be Cautious about Everyone's Code

## Be Cautious about Your Code

## Some references

## Poka Yoke

Avoid Mistakes

Fool Proofing

## Only one way to plug it in

## Maybe NOT such a good design

## Let's make one concept clear:

## Code is Reusable

## Code is Reusable

## Code is NOT Reusable

## Code is Crap

> ninety percent of everything is crap. Theodore Sturgeon

## Abstractions are Reusable

## Don't write Crap

## Don't write Code

## Code for your Use-case

## NOT for your re-use-case

## Good Practices

## Assuming SOLID code

## Assuming some Object Calisthenics

## Make State Immutable

Greatly reduces bugs and cognitive load

PSR-7 got this right!

## What about performance?

Profile!

Profile!

Profile!

> O(log(n))

## No Setters

## Setters don't mean Anything, anyway

## The Constructor is the ONLY injection point

## No un-initialized properties

## Optional Dependencies?

```plain text
class DbConnection { public function __construct(...) { // ... } public function setLogger(Logger $logger = null) { $this->logger = $logger; } }
```

## Optional dependencies Do Not Exist!

```plain text
class DbConnection { public function __construct(..., Logger $logger) { // ... } // look ma, no setters! }
```

```plain text
$dbConnection = new DbConnection(..., new FakeLogger());
```

## Avoid unnecessary public methods

## Unnecessary = you can find another way

> A public method is like a child: once you've written it, you are going to maintain it for the rest of its life! Stefan Priebsch

## Avoid Logic Switch Parameters

```plain text
class Spammer { public function spam($email, $template, $optOutLink = false) { // yes, this is a really bad spammer } }
```

```plain text
class Spammer { public function sendIllegalSpam($email, $template) { // without opt-out link } public function sendApparentlyLegalSpam($email, $template) { // with opt-out link } }
```

## All State should be Encapsulated

```plain text
class BankAccount { // ... public function setLastRefresh(DateTime $lastRefresh) { $this->lastRefresh = $lastRefresh; } }
```

(examples are purposely simplified/silly!)

```plain text
$currentTime = new DateTime(); $bankAccount1->setLastRefresh($currentTime); $bankAccount2->setLastRefresh($currentTime);
```

```plain text
// ... stuff
```

```plain text
$currentTime->setTimestamp($aTimestamp);
```

```plain text
class BankAccount { // ... public function setLastRefresh(DateTime $lastRefresh) { $this->lastRefresh = clone $lastRefresh; } }
```

```plain text
class BankAccount { // ... public function setLastRefresh(DateTimeImmutable $lastRefresh) { $this->lastRefresh = $lastRefresh; } }
```

```plain text
class MoneyTransfer { public function __construct(Money $amount, DateTime $transferDate) { $this->amount = $amount; $this->transferDate = $transferDate; } }
```

```plain text
class MoneyTransfer { public function __construct(Money $amount, DateTime $transferDate) { $this->amount = $amount; $this->transferDate = clone $transferDate; } }
```

## External Access includes Subclasses

## Make State Private

## Private properties

BY DEFAULT!

## A Public Method is a Transaction

```plain text
public function addMoney(Money $money) { $this->money[] = $money; } public function markBillPaidWithMoney(Bill $bill, Money $money) { $this->bills[] = $bill->paid($money); }
```

```plain text
$bankAccount->addMoney($money); $bankAccount->markBillPaidWithMoney($bill, $money);
```

These methods are part of a single interaction, but are exposed as separate public APIs

```plain text
private function addMoney(Money $money) { $this->money[] = $money; } private function markBillPaidWithMoney(Bill $bill, Money $money) { $this->bills[] = $bill->paid($money); } public function payBill(Bill $bill, Money $money) { $this->addMoney($money); $this->markBillPaidWithMoney($bill, $money); }
```

```plain text
$bankAccount->payBill($bill, $money);
```

Single public API endpoint wrapping around state change

## DO NOT assume Method Idempotence

```plain text
$userId = $controller->request()->get('userId'); $userRoles = $controller->request()->get('userRoles');
```

Subsequent method calls are assuming that the request will be the same

(apart from being foolish API)

```plain text
$request = $controller->request(); $userId = $request->get('userId'); $userRoles = $request->get('userRoles');
```

No assumption on what request() does: we don't trust it

## Assert

## Enforce Strict Invariants

```plain text
class Train { public function __construct(array $wagons) { $this->wagons = (function (Wagon ...$wagons) { return $wagons; })(...$wagons); } }
```

Yes, it's horrible, but it works

(on PHP 7)

## beberlei/assert

Just use it!

## No Mixed Parameter Types

```plain text
class PrisonerTransferRequest { /** * @param mixed $accessLevel * - false if none * - true if guards are required * - null if to be decided * - 10 if special cargo is needed * - 20 if high security is needed */ public function approve($securityLevel) { // ... } }
```

## WTFBBQ?!

```plain text
class PrisonerTransferRequest { public function approve(PrisonerSecurityLevel $securityLevel) { // ... } }
```

## Value Objects FTW

## Value Objects ensure Consistency

## Value Objects simplify Data Validation

```plain text
register(new EmailAddress('ocramius@gmail.com'));
```

No need to re-validate inside register()!

## No Mixed Return Types

## Value Objects, again!

## Avoid Fluent Interfaces

```plain text
$thing ->doFoo() ->doBar();
```

```plain text
$thing = $thing->doFoo(); $thing = $thing->doBar();
```

```plain text
$thing->doFoo(); $thing->doBar();
```

## POKA YOKE !!!!!!!1!!ONE

Only ONE way!

```plain text
$thing->doFoo(); $thing->doBar();
```

(There's a blogpost for that!)

## Make Classes Final By Default

## An Extension is a different Use-Case

## Does Extending the Interface add Any Meaning?

If not, then do not use extends.

You don't even need extends.

(There's a blogpost for that!)

## Traits

## Traits

Do you REALLY need to re-use that code?

DRY or just hidden coupling?

## Ensure Deep Cloning

```plain text
class LoginRequest { public function __clone() { $this->time = clone $this->time; } }
```

Cloning requires encapsulated state understanding

... which is a contradiction!

## So ...

## ... disable Cloning?

```plain text
class Human { public function __clone() { throw new \DomainException( 'Why would you even clone me?!' . 'It\'s currently illegal!' ); } }
```

## Same for Serialization!

## Is it really Serializable?

Do you know the lifecycle of every dependency?

Is your object even supposed to be stored somewhere?

```plain text
class MyAwesomeThing { public function __sleep() { throw new BadMethodCallException( 'NO! MY THING! CANNOT HAZ!' ); } }
```

## Maybe serialize() only if it is a Serializable?

## UNIT TESTING

## Test all scenarios

YES, ALL OF THEM!

Bring your CRAP score <= 2

It will make you cry.

You will implore for less features.

You will love having less features.

```plain text
interface AuthService { /** * @return bool * * @throws InvalidArgumentException * @throws IncompatibleCredentialsException * @throws UnavailableBackendException */ public function authenticate(Credentials $credentials); }
```

```plain text
class MyLogin { public function __construct(AuthService $auth) { $this->auth = $auth; } public function login($username, $password) { if ($this->auth->authenticate(new BasicCredentials($username, $password)) { // ... return true; } return false; } }
```

What can happen here?

A boolean true is returned

A boolean false is returned

A InvalidArgumentException is thrown

A IncompatibleCredentialsException is thrown

A UnavailableBackendException is thrown

Such little code, 5 tests already!

## Recap

## These are suggestions

## Trust NO Code

## Reduce features to the bare minimum

## YAGNI

## Strict invariants

Await strict

Return strict

## Thanks!
