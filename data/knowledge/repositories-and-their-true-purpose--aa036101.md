---
title: "Repositories and their true purpose"
notion_id: aa036101-1b73-44c6-85c3-c87b46d33901
notion_url: https://app.notion.com/p/Repositories-and-their-true-purpose-aa0361011b7344c685c3c87b46d33901
last_edited: 2023-05-26T19:29:00.000Z
source_url: https://muhammedsari.me/repositories-and-their-true-purpose
tags: ["Programming", "System Design / Software Architecture", "Domain Driven Design", "Article", "Muhammed Sarı", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

_Lately, posts and tweets regarding the __`Repository`__ pattern have made yet another resurgence. It's seemingly impossible to predict when, where or why such "spicy topics" will rear their heads... However, the spark that causes the ignition of these "hot topics" is almost always the following question (or something similar):_

_— Random Techfluencer__ __
__That's why, in this blog post, I'd like to provide some further clarity regarding this totally misunderstood software design pattern and why the #1 argument (the question above) against its use is actually insignificant and __almost irrelevant__.__ __**
**__**Defining a Repository**__ __
__First and foremost, let's start off by defining what a __`Repository`__ actually is. The __`Repository`__ pattern is defined as follows in _[_PoEAA_](https://martinfowler.com/eaaCatalog/repository.html)_:__ __ __Mediates between the domain and data mapping layers using a collection-like interface for accessing domain objects.__ __ __
__It is of paramount importance that we establish the facts below before moving on to the other sections.__ __**
**__**(...) accessing domain objects**__ __
__Domain objects are actors in the domain layer that possess an authoritative set of business capabilities for carrying out certain tasks. These capabilities or behaviors are exposed as public methods on said actors in order to make consistent state changes. Domain objects are also known as write models, entities or as aggregates in _[_DDD_](https://www.youtube.com/watch?v=8Z5IAkWcnIw)_ lingo.__ __
__Note: Aggregates and their exact purpose is out of scope for this blog post. However, I can recommend _[_this short, concise write-up by Shawn McCool_](https://shawnmc.cool/2023-05-11_aggregates-for-those-familiar-with-activerecord)_ if you'd like to learn more about them.__ __
__You've probably heard the notion "business logic" numerous times by now. Well, these models are the ones that actually determine what that "business logic" should entail.__ __**
**__**(...) collection-like interface (...)**__ __
__In an ideal world, a persistence layer for the entities would not be needed as everything can be added and removed from an in-memory collection. For example:__ __

__`final class Users
{
private array $users;

private function __construct(User ...$users)
    {
$this->users = $users;
    }

public static function empty(): self
    {
return new self();
    }

public function add(User $user): void
    {
$this->users[$user->id()->asString()] = $user;
    }

public function find(UserId $id): User
    {
return $this->users[$id->asString()] 
?? throw CouldNotFindUser::becauseItIsMissing();
    }

public function remove(User $user): void
    {
unset($this->users[$user->id()->asString()]);
    }
}
`__ __
__Unfortunately, the real world is often rather different from the ideal world. PHP has its (in)famous request-response lifecycle which results in the loss of every bit of relevant context once an incoming request has been handled and a response has been sent to the client. A __`Repository`__ assists us in approximating this ideal world by giving the __illusion__ that we can perform our operations on in-memory collections that seemingly live forever. An example __`Repository`__ could be:__ __

__`interface UserRepository
{
public function find(UserId $id): User;
public function save(User $user): void;
public function remove(User $user): void;
}
`__ __
__Please take note of the minimal signature of this interface.__ __**
**__**Collection-oriented vs. persistence-oriented**__ __
_[_Vaughn Vernon_](https://vaughnvernon.com/)_, the author of The Big Red Book (_[_iDDD_](https://kalele.io/books/)_), mentions collection-oriented and persistence-oriented __`Repository`__ implementations in Chapter 12. I'd like to briefly mention this fact, because this is the reason why you might see different "flavors" of __`Repository`__ implementations in the wild. The difference lies primarily in the semantics.__ __
__A collection-oriented design can be considered a traditional design because of adherance to an in-memory collection's standard interface.__ __

__`$users->add($user);
`__ __
__A persistence-oriented design is also known as a save-based __`Repository`__.__ __

__`$users->save($user);
`__ __
__Personally, I prefer the persistence orientation due to PHP's ephemeral nature.__ __**
**__**Authoritative collection**__ __
__A __`Repository`__ is the authoritative collection for interacting with a specific type of entity. It can be used to store, filter, retrieve and remove entities based on the application's needs. In other words, we delegate the task for remembering the existence of a certain entity to the __`Repository`__.__ __**
**__**Explained by example: publishing a post**__ __
__Let's take a look at a use case to solidify our understanding.__ __

__`final readonly class PublishPostHandler
{
public function __construct(
private PostRepository $posts,
    ) {}

public function handle(PublishPost $command): void
    {
        $post = $this->posts->find($command->id);

        $post->publish();

$this->posts->save($post);
    }
}
`__ __
__This use case assumes that a __`Post`__ entity must already exist in order to publish it. Since the __`PostRepository`__ is the authoritative collection for dealing with these __`Post`__ entities, we can ask it to provide us with a __`Post`__ entity for the given __`PostId`__:__ __

__`$post = $this->posts->find($command->id);
`__ __
__Once we've received a __`Post`__ instance, we carry on with the task we were supposed to carry out in the first place:__ __

__`$post->publish();
`__ __
__The __`publish`__ method exposes the behavior that is responsible for actually "publishing a blog post". If we dig a little deeper, we can see that it is also enforcing crucial _[_invariants_](https://codeopinion.com/aggregate-design-using-invariants-as-a-guide)_:__ __

__`public function publish(): void
{
if ($this->isPublished()) {
throw CouldNotPublish::becauseAlreadyPublished();
    } elseif ($this->summary->isEmpty()) {
throw CouldNotPublish::becauseSummaryIsMissing();
    } elseif ($this->body->isEmpty()) {
throw CouldNotPublish::becauseBodyIsMissing();
    } elseif ($this->tags->isEmpty()) {
throw CouldNotPublish::becauseTagsAreMissing();
    }

// omitted for brevity
}
`__ __
__If everything goes well, we move on and tell the __`PostRepository`__ to remember the __`Post`__ in its __current state__:__ __

__`$this->posts->save($post);
`__ __
__Next time we interact with the __`PostRepository`__ and ask for the exact same entity, we can expect to receive the __`Post`__ in this state. The __`PostRepository`__ will ensure that this condition is met at __all times__. This is a __`Repository`__'s single most important __responsibility__, after all. The __`PostRepository`__ clearly defines the boundaries around the application service which also yields a lot of benefits such as isolated testability and purposefully keeping the (core) domain oblivious to its surroundings.__ __**
**__**Persistence agnosticity**__ __
__Let's quickly recall __Random Techfluencer__'s original statement:__  __
__Random Techfluencer__ is actually discouraging the use of the __`Repository`__ because "how many times are you going to swap out data sources?".__ __
__Now, please let me make something absolutely clear. The swapping of the data source is a puny argument whether you use it to __promote__ __**or**__ __obstruct__ the use of the __`Repository`__. It does __not__ matter which camp (pro / contra) you belong to. Do you really want to think about swapping out data sources as you are designing the domain? This kind of thinking is - in my humble opinion - flawed.__ __**
**__**Ad hoc persistence swapping**__ __
__The fact that you can easily swap out data sources later on is nothing but a __bonus__ that you are awarded by carefully placing boundaries around your application. This is "boundaries in software design 101" and trying to use this as the main selling point every single time does noone any good.__ __
__You can start out your application by using simple JSON files on disk and gradually evolve towards "beefier" solutions as different needs emerge.__ __

__`final class UserRepositoryUsingJsonFilesOnDisk implements UserRepository
{
public function add(User $user): void
    {
// add a user
    }

public function find(UserId $id): User
    {
// find a user
    }

public function remove(User $user): void
    {
// remove a user... you get the point
    }
}
`__ __
__Different features can evolve independent of each other and infrastructural costs can be kept to a minimum. Why use an expensive cloud-hosted solution __for everything__ if 90% of the other features are well-suited for a storage mechnism like SQLite? Why keep using MySQL __for every single feature__ if 10% of the features are well-suited for Elastic and Riak?__ __**
**__**Testability**__ __
__In a similar vein to persistence agnosticity, testability is another __bonus__ we are awarded by carefully placing boundaries around our application. The real thing can keep using a __`DoctrinePostRepository`__ while the tests can use an __`InMemoryPostRepository`__ allowing us to have lightning fast tests.__ __
__The test for the "publishing a blog post" use case, that was mentioned previously, might look as follows:__ __

__`// Arrange
$post = $this->aPost(['id' => PostId::fromInt($id = 123)]); // draft
$repository = $this->aPostRepository([$post]); // in-memory repository
$handler = new PublishPostHandler($repository);

// Act
$handler->handle(new PublishPost($id));

// Assert
$this->assertTrue($repository->wasSaved($post));
$this->assertTrue($post->isPublished());
`__ __
__In this example, we're testing the application service represented by the command handler. We don't need to test that the repository stored the data in the database or wherever else. We need to test the specific behavior of the handler, which is to publish the __`Post`__ object and pass it into the repository to preserve its state.__ __
__"This is not a big deal", you might rightfully say, "I can just hit persistence during my tests every single time". I'm not sure if you know someone who's worked on a project whose test suite was completely shut down because it just took way too long to go through the entire thing? I do know someone and that person is unfortunately me. Integration and System / E2E tests definitely have their place, but the sheer velocity and the fast feedback loop of unit tests is still highly desirable.__ __**
**__**Alleviating performance issues**__ __
__Performance is another reason as to why a __`Repository`__ is often employed. It's not an uncommon scenario to have millions of instances of a certain entity type so we are kind of forced to offload this to an external data store.__ __
__Assume the following excerpt from an imaginary __`User`__ entity:__ __

__`public function changeEmail(Email $newEmail, Users $allUsers)
{
if ($allUsers->findByEmail($newEmail)) {
throw new CannotChangeEmail::becauseEmailIsAlreadyTaken();
    }

$this->email = $newEmail;
}
`__ __
__The __`changeEmail`__ behavior depends on a __`Users`__ collection to determine whether the new email address can be used. The (imaginary) domain experts told us that an email change may not happen as long as there is another user in possession of that new email address.__ __
__This code will work just fine until we hit a certain amount of users. The collection's sheer size will become a bottleneck for the lookups that must be performed in order to enforce invariants. We could fix this problem by injecting a __`UserRepository`__ instead of passing every single __`User`__ in existence via an in-memory __`Users`__ collections.__ __

__`public function changeEmail(Email $newEmail, UserRepository $users)
{
if ($users->findByEmail($newEmail)) {
throw new CannotChangeEmail::becauseEmailIsAlreadyTaken();
    }

$this->email = $newEmail;
}
`__ __
__This way, the domain model will still be responsible for enforcing the invariants; but we had to trade the _[_domain model's purity_](https://enterprisecraftsmanship.com/posts/domain-model-purity-completeness/)_ off against performance. Nonetheless, this is most definitely an acceptable trade-off.__ __**
**__**Command Query Responsibility Segregation**__ __
__"I thought this blog post was about the __`Repository`__ pattern? What's the deal with _[_CQRS_](https://web.archive.org/web/20120419072250/https://goodenoughsoftware.net/2012/03/02/cqrs)_ all of a sudden..?" Please let me explain.__ __**
**__**Write models (commands)**__ __
__Until now, we've seen how the __`Repository`__ helps us with dealing with the lifecycle of __domain objects__. We established the fact that these domain objects are also known as write models / entities / aggregates that are responsible for performing state changes in a consistent manner. In other words, the aggregates represent a consistency boundary that must follow the business rules and apply them at all times in order to stay consistent. Naturally, these state changes always occur as a result of a command entering an application.__ __**
**__**Read models (queries)**__ __
__We need to ask ourselves whether we actually need to perform state changes or just need some data. Why would we "just need some data"? Well... you guessed it right: for __**queries**__. _[_CQRS_](https://web.archive.org/web/20120419072250/https://goodenoughsoftware.net/2012/03/02/cqrs)_ is a dead simple pattern for separating the logical models for read and write concerns—that's it. It has nothing to do with event sourcing / eventual consistency / separated data stores etc. These buzzwords are often thrown into the mix by people who don't really know what they're talking about. Use cases that involve queries will benefit from better optimized, dedicated __read models__.__ __**
**__**Explained by example: displaying a table of invoices**__ __
__Let's take a look at a use case to solidify our understanding.__ __

__`final readonly class ViewInvoicesController
{
public function __construct(
private GetMyInvoices $query,
private Factory $view,
    ) {}

public function __invoke(Request $request): View
    {
        $invoices = $this->query->get();

return $this->view->make('view-invoices', [
'invoices' => $invoices,
        ]);
    }
}
`__ __
__This use case is responsible for __displaying__ a table of invoices to the user. All of the magic happens during this line:__ __

__`$invoices = $this->query->get();
`__ __
__The query handler __`GetMyInvoices`__ provides us with a collection of __`InvoiceSummary`__ __**read models**__ dedicated for this purpose. A single __`InvoiceSummary`__ instance might look as follows:__ __

__`final readonly class InvoiceSummary
{
public function __construct(
public int $amountOfDiscountsApplied,
public string $paymentTerms,
public string $recipient,
public int $totalAmountInCents,
    ) {}
}
`__ __
__Eagle-eyed readers may already have noticed that this is in fact a __`Data Transfer Object`__. __`DTO`__s typically contain only data and no behavior. However, this is exactly what we want: a __read model__ dedicated to the purpose of displaying some relevant data to the user. You may already have noticed that this model doesn't contain any information regarding the individual invoice line items; and this is totally on purpose! A table view cannot display individual invoice line items. Thus, our __read model__ is optimized and carefully crafted for this exact use case.__ __
__The write model might look like this (courtesy of _[_Shawn McCool_](https://shawnmc.cool/2023-05-11_aggregates-for-those-familiar-with-activerecord)_):__ __

__`final readonly class LineItem
{
public __construct(private bool $isDiscount) {}

public function isDiscount(): bool
    {
return $this->isDiscount;
    }
}

final class Invoice
{
private RecipientName $recipientName;

private LineItems $lineItems;

public function __construct(
RecipientName $recipientName
    ) {
$this->recipientName = $recipientName;
$this->lineItems = LineItems::empty();
    }

public function addLineItem($item): void
    {
if (
            $item->isDiscount()
&& $this->lineItems->hasDiscountedItem()
        ) {
throw CannotAddLineItem::multipleDiscountsForbidden($item);
        }

$this->lineItems->add($item);
    }
}
`__ __
__So to be more precise, __we went directly to the data source itself instead of trying to shoe-horn a use case into an Invoice write model that is totally not designed to fulfill a specialized query-based, read use case__. Why carry the burden of instantiating this complex write model in order to fulfill a use case that won't even need any of the line items that are defined within this write model? The write model __requires__ all of the line items in order to keep its state consistent, but the read model does not.__ __**
**__**Where does a Repository belong to: application or domain layer?**__ __
__We can consider the application layer as the specific layer within a multi-layered architecture that handles the implementation details unique to the application, such as database persistence, internet protocol knowledge (sending emails, API interactions), and more. Now, let's establish the domain layer as the layer in a multi-layered architecture that primarily deals with business rules and business logic.__ __
__Given these definitions, where exactly do our repositories fit into the picture? Let's revisit a variation of a source code example we discussed earlier:__ __

__`final class InMemoryUserRepository implements UserRepository
{
private array $users = [];

public function find(UserId $id): User
    {
return $this->users[$id->asString()]
?? throw CouldNotFindUser::becauseItIsMissing();
    }

public function remove(User $user): void
    {
unset($this->users[$user->id()->asString()]);
    }

public function save(User $user): void
    {
$this->users[$user->id()->asString()] = $user;
    }
}
`__ __
__I'm observing numerous implementation details that can be regarded as "noise". Therefore, this implementation detail belongs in the application layer. Let's remove this noise and see what we are left with:__ __

__`final class InMemoryUserRepository implements UserRepository
{
private array $users = [];

public function find(UserId $id): User
    {
    }

public function remove(User $user): void
    {
    }

public function save(User $user): void
    {
    }
}
`__ __
__Does this actually remind you of something? Perhaps this?__ __

__`interface UserRepository
{
public function find(UserId $id): User;
public function save(User $user): void;
public function remove(User $user): void;
}
`__ __
__Placing an interface at layer boundaries entails the following implication: While the interface itself can encompass domain-specific concepts, its implementation should not. In the context of repository interfaces, __**they belong to the domain layer**__. The __implementation__ of repositories belongs to the application layer. Consequently, we can freely utilize type-hinting for repositories within the domain layer, without any need for dependencies on the application layer.__ __**
**__**Various other benefits**__ __
__Below is a non-exhaustive list of various other benefits a __`Repository`__ can bring along with it:__ __ __
• __Access to the decorator pattern to add additional concerns without having to modify the domain e.g. to employ something like _[_hashids_](https://hashids.org/)_ for YouTube-like identifiers.__ __
• __The ability to implement the _[_transactional outbox pattern_](https://microservices.io/patterns/data/transactional-outbox.html)_ for mission-critical, event-driven systems.__ __
• __Centralizing access / persistence logic if the application relies on data models primarily and you'd like to migrate away.__ __
• __Automatically adding audit information alongside the persisted entity.__ __ __
__...__ __**
**__**Wrap-up**__ __
__That was a lot to go through! Thanks for sticking around until the end.__ __
__Basically, if we were to enumerate all of the benefits for using a __`Repository`__, persistence agnosticity would definitely come last or at the very least be close to being last. Therefore, I hope that we can stop taking concepts at face value and actually examine them a little deeper to unearth the actual use cases and the contexts in which they're supposed to be used.__ __ __
• __ __`Repository`__ is the authoritative actor for safely collecting and preserving entities and managing their lifecycle__ __
• __The ability to swap underlying the persistence driver is a mere __bonus__ __ __
• __The ability to easily test without an actual persistence driver is a mere __bonus__ __ __
• __Do use a __`Repository`__ for your write models__ __
• __Don't use a __`Repository`__ for your read models: go to the data source instead__ __ __
_[_Join the discussion on Twitter!_](https://twitter.com/mabdullahsari/status/1661288614263750656)_ I'd love to know what you thought about this blog post.__ __
__Thanks for reading!__  _
