---
title: "Doctrine 2 ORM Best Practices"
notion_id: c7d81d56-5bf0-4cc1-be2b-aa414f3d6dc8
notion_url: https://app.notion.com/p/Doctrine-2-ORM-Best-Practices-c7d81d565bf04cc1be2baa414f3d6dc8
last_edited: 2023-10-30T14:20:00.000Z
source_url: https://ocramius.github.io/doctrine-best-practices/#/
tags: ["Slides", "Marco Pivetta (Ocramius)", "English", "PHP", "Object Oriented Programming"]
---
## Hi!

## I'm Marco!

## Ocramius

## Doctrine Project

A group of persistence-oriented libraries for PHP

## We are looking for contributors!

(Must be prepared to take unpopular decisions)

A JSR-317 / Hibernate inspired Object Relational Mapper

## Doctrine 2 ORM Best Practices

## Know your enemy

## Know your enemy tools

## Knowledge is Power

## When is an ORM the appropriate tool?

## Who is Doctrine ORM for?

## OLTP

## DDD

## Fast prototyping

## OO-first

## Who is Doctrine ORM NOT for?

## Dynamic data structures

## Reporting

That's what SQL is for

It's a QUERY language, after all

## Where do we start?

## Entities (of course)

Entities should work without the ORM

Entities should work without the DB

## Entities (mostly) represent your domain

## The Database is just saving things

## Design entities first

## Define the Database after modeling your domain

## Define Mappings after designing the Entities

## Designing Entities

```plain text
class User { private $username; private $passwordHash; public function getUsername() :string { return $this->username; } public function setUsername(string $username) { $this->username = $username; } public function getPasswordHash() : string { return $this->passwordHash; } public function setPasswordHash(string $passwordHash) { $this->passwordHash = $passwordHash; } }
```

## Entities are NOT typed arrays

## Entities have Behavior

## Entity internal State design is Irrelevant

You can deal with state after designing the API

Designing State-first leads to terrible coupling

```plain text
class User { private $banned; private $username; private $passwordHash; public function toNickname() : string { return $this->username; } public function authenticate( string $pass, callable $checkHash ) : bool { return $checkHash($pass, $this->passwordHash) && ! $this->hasActiveBans(); } public function changePass(string $pass, callable $hash) { $this->passwordHash = $hash($pass); } }
```

## No behavior = no need for the ORM

## Respect the LOD

```plain text
class User { // ... public function hasAccessTo(Resource $resource) : bool { return (bool) array_filter( $this->role->getAccessLevels(), function (AccessLevel $acl) use ($resource) : bool { return $acl->canAccess($resource) } ); } }
```

```plain text
class User { // ... public function hasAccessTo(Resource $resource) : bool { return $this->role->allowsAccessTo($resource); } }
```

More expressive

Easier to test

Less coupling

More flexible

Easier to refactor

## Disallow Collection access from outside the Entity

```plain text
class User { private $bans; public function getBans() : Collection { return $this->bans; } }
```

```plain text
public function banUser(Uuid $userId) { $user = $this->repository->find($userId); $user->getBans()->add(new Ban($user)); }
```

## Keep Collections hidden in your Entities

```plain text
class User { private $bans; public function ban() { $this->bans[] = new Ban($this); } }
```

```plain text
public function banUser(Uuid $userId) { $user = $this->repository->find($userId); $user->ban(); }
```

## Entity Validity

## Entities should always be valid

## Invalid state should be in a different object

(You may need a DTO)

(Also applies to Temporary State)

## Stay valid after __construct

(Regardless of the DB)

Named constructors are OK

## Avoid setters

## Avoid coupling with the Application Layer

```plain text
class UserController { // form reads from/writes to user entity (bad) public function registerAction() { $this->userForm->bind(new User()); } }
```

```plain text
class UserController { // coupling between form and user (bad) public function registerAction() { $this->em->persist(User::fromFormData($this->form)); } }
```

## Form components break Entity Validity

## Both Symfony\Form and Zend\Form are terrible

(For this use-case)

Use a DTO instead

## Avoid Lifecycle Callbacks

## Lifecycle Callbacks are a Persistence Hack

Lifecycle Callbacks are supposed to be the ORM-specific serialize and unserialize

Don't use Lifecycle Callbacks for Business Logic/Events

## Avoid auto-generated identifiers

Your db operations will block each other

You are denying bulk inserts

You cannot make multi-request transactions

Your object is invalid until saved

Your object does not work without the DB

## Use UUIDs instead

Don't forget that a UUID is just a 128 bit integer!

```plain text
public function __construct() { $this->id = Uuid::uuid4(); }
```

## auto_increment is abused for sorting

Are you looking for a DATETIME field instead?

## Avoid derived primary keys

You are just normalizing for the sake of it

Does your domain really NEED it?

## Avoid composite primary keys

Any reason to not use an unique constraint instead?

Do they make a difference in your domain?

## Favour immutable entities

Or append-only data-structures

```plain text
class PrivateMessage { private $from; private $to; private $message; private $read = []; public function __construct( User $from, User $to, string $message ) { // ... } public function read(User $user) { $this->read[] = new MessageRead($user, $this); } }
```

```plain text
class MessageRead { private $user; private $message; public function __construct(User $user, Message $message) { $this->id = Uuid::uuid4(); $this->user = $user; $this->message = $message; } }
```

Immutable data is simple

Immutable data is cacheable (forever)

Immutable data is predictable

Immutable data enables historical analysis

You may want to look at Event Sourcing...

## Avoid Soft-Deletes

## Soft Deletes are a broken idea

Soft Deletes come from an era where keeping everything in a single DB was required

## Soft Deletes break immutability

## Soft Deletes break Data integrity

(and therefore validity)

Soft Deletes can usually be replaced with more specific domain concepts

## Mapping driver choice

## Use Annotations in private packages

## Use XML Mappings in public packages

## Lazy or Eager?

## Eager Loading is Useless

## Extra Lazy indicates high risk areas

We reach the limits of the ORM

Careful about performance/transaction size!

## Avoid bi-directional associations

## bi-directional associations are overhead

Code only what you need for your domain logic to work

Hack complex DQL queries instead of making them simpler with bi-directionality

## Use custom repositories for improved expressiveness

```plain text
final class UserRepository { public function findUsersThatHaveAMonthlySubscription() { // ... INSERT DQL/SQL HELL HERE ... } }
```

## Query Functions are better than repositories

```plain text
final class UsersThatHaveAMonthlySubscription { public function __construct(EntityManagerInterface $em) { // ... } public function __invoke() : Traversable { // ... INSERT DQL/SQL HELL HERE ... } }
```

## Repositories are Services

So are Query Functions

Avoid ObjectManager#getRepository()

It is a ServiceLocator

It causes the same problems of the ServiceLocator

## Inject Repositories instead

Separate MyRepository#get() and MyRepository#find()

MyRepository#find() can return null

MyRepository#get() cannot return null

```plain text
final class BlogPostRepository { // ... public function getBySlug($slug) : BlogPost { $found = $this->findOneBy(['slug' => (string) $slug]); if (! $found) { throw BlogPostNotFoundException::bySlug($slug); } return $found; } }
```

Using a get() method that throws, you can simplify error logic

## Avoid 2pc

## Keep transactions unrelated

Use ObjectManager#clear() between different ObjectManager#flush() calls

## Different boundaries = Different transactions

Communicate between boundaries via identifiers, not object references

## Keep Normalization under control

## Keep Normalization (freaks) under control

You may need to gag your DBAs...

Or get them to understand your needs

Academic and practical knowledge may differ

## What about Performance?

## Know how the ORM is structured

EntityManager

UnitOfWork

Metadata Drivers

DQL

Repositories

Second Level Cache

## Profile these hotspots

Measuring is the only way

There are other talks about this...

## RECAP

## Domain First

## Do not normalize without a need for it

Otherwise, you're digging your own grave!

## Consider using Separate DBs

Transactional Data != Reporting Data

## Thanks!
