---
title: "The infamous Repository Pattern in PHP"
notion_id: 42e4aeda-39b5-43f8-bbc0-b3e1d61c56d7
notion_url: https://app.notion.com/p/The-infamous-Repository-Pattern-in-PHP-42e4aeda39b543f8bbc0b3e1d61c56d7
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://doeken.org/blog/repository-pattern
tags: ["Doeke Norg Blog", "English", "PHP", "System Design / Software Architecture", "Domain Driven Design", "Article"]
---
   Photo by [Susan Q Yin](https://unsplash.com/photos/2JIvboGLeho)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The repository pattern is easily one of the most talked about patterns. Some love it; others... not so much. In this post we're going to look at the repository pattern; what it is and what it is _not_. We'll also look at the differences between read and write repositories and some conventions.

## What is the Repository Pattern?

The repository pattern is a data source abstraction that behaves like a collection; containing domain (value) objects. This abstraction is called a _Repository_ and is often introduced as an interface. The repository is considered to be the single source of truth for retrieving and storing domain objects or entities.

Let's look at a quick example repository for retrieving `Book` entities.

```php
interface BookRepository
{
	public function get(BookId $id): Book;

	public function save(Book $book): void;

	public function delete(Book $book): void;
}
```

This is pretty much all that is needed for a functional repository; a way to `get`, `save` and `delete` a book from the collection.

Notice the absence of an `add` and `update` method. These methods are of no use to a repository, because it communicaties knowledge about the underlying data source (e.g. whether the `Book` is already stored). The `save` method will guide the process of adding or updating the `Book` in data source.

> **Note:** Since a Repository is the single source of truth, it should be the only service to persist an entity. This makes the repository pattern somewhat incompatible with the _active record pattern_, because then the entities could persist themselves. To use active record, make sure you only call the entity's `save` method inside the repository's `save` method.

## A repository is not a factory

As you might have noticed; the repository does not have a `create` method. This is because the purpose of a repository is to contain objects. Just as a library (mostly) does not write and publish books, a repository does not create the objects it contains. If the creation of the objects is a complex matter (e.g. it takes many dependencies to create the object) this job is probably more suitable for a _factory_.

This doesn't mean you cannot combine a repository and a factory in a single service. You absolutely can; but the question is if you should. This is a question I cannot answer for your situation. Personally, I like to keep my services separate, so they can evolve independently.

## A repository is not an ORM

Object Relation Mapping (ORM) is abstraction layer on top of a (relational) database. It provides tools to retrieve a row from the database, and map its columns to the properties of an entity object. An ORM is often equipped to handle multiple databases types, like MySQL, PostgreSQL or Redis. The goal of the ORM is to make interaction with the database an object-oriented approach. You update the columns as properties on the entity, and the ORM does the translation back to and from the database.

Often, an ORM has a Query Builder; which is another object-oriented approach for retrieving data as entity objects from the database. You construct your `WHERE` clauses with simple `->where()` methods, and the ORM will again translate these query objects to _real_ database queries and return you the results as mapped entities.

A repository however is not concerned with specifics on how the entity came to be, where it came from, or how it was located. The underlying data _might_ be located and stored using an ORM, or read and written into a file or even retrieved via a search engine. But those are only the _technical implementations_ of the repository. The repository only represents these entities as if it was a _collection_.

> **Caution:** When creating a repository in conjunction with an ORM, it is tempting to use Query Builder specific arguments, or even to return the Query Builder to retrieve the information. _Do not do this._ In that case you are coupling the Query Builder to your repository, defeating the purpose of the repository.

When it comes to advantages of a Repository; the ability to "easily replace the database layer" is often thought of to be a non-argument and is ridiculed by saying it is very unlikely to replace your database implementation, and that (as we have seen) the ORM already takes care of that.

However, a repository abstracts away the complicated queries (even if you use a Query Builder) or services needed behind the scenes to retrieve the required domain objects. This means that you can start off with an implementation of the repository that fully relies on a database like MySQL (even through an ORM), while being able to replace certain logic in the future with a service like a search engine. This is especially useful when using _Read Repositories_ (more on those later).

> **Note:** The repository pattern can be mixed with the [Decorator Pattern](https://doeken.org/blog/decorator-vs-proxy-pattern) pretty well. For example: you can decorate your repository with a cacheable implementation which could reduce the amount of round trips to underlying data source.

Another advantage of a repository is you can easily replace it with an in-memory variant like this, which makes for easy (and maybe even more important: _fast_) unit testing.

```php
class InMemoryBookRepository implements BookRepository
{
	/** @var Book[] */
	private array $collection = [];

	public function get(BookId $id): Book
	{
		return $this->collection[$id->id] ?? throw new BookNotFound();
	}

	public function save(Book $book): void
	{
		$this->collection[$book->id->id] = $book;
	}

	public function delete(Book $book): void
	{
		unset($this->collection[$book->id->id]);
	}
}
```

> **Note:** Some developers are against using an in-memory repository and encourage the use of a mock; because an in-memory class adds the need for another test for this implementation. And while this _might_ be true (I would not test this class) let's be honest; it is not _that_ much work, and you get the benefit of using it in multiple places. It also has the exact same interaction, without the need to add assertions and instructions on what to return on which method calls. I personally would advise [against using a mock](https://doeken.org/blog/stop-mocking-about-event-dispatcher).

## Naming conventions (find vs. get)

You'll often come across different names for methods on a repository. `delete` might very well be called `remove`, and `save` could be called `store`. But their underlying purpose does not change. There is however a small nuance when it comes to the naming of query methods (prefixed with) `get` and `find`.

In general, the use of `get`-methods implies the client expects the object to exist and be returned. In the case of `find`-methods, there is an expectation that the object might not exist. Therefore, in most cases `get` will throw a `NotFound` exception when the object cannot be located, while `find` will return `null` or an empty collection.

`get` is also predominantly used to retrieve a single object, while `find` is used to locate multiple objects based on a certain criteria. This is why read repositories will often have methods like `findBySomeValue($value)` for retrieving multiple objects, and `findOneBySomeValue($value)` to implicate the result of a single object.

> **Note:** It is custom for `find` methods to return an empty collection, when there are no matches, instead of a `null` result.

```php
// Implies the existence of a Book.
public function get(BookId $book_id): Book;

// Implies the possibility of not finding the Book.
public function findOneById(BookId $book_id): ?Book;

// Expects to find multiple Books from a given Author.
public function findByAuthor(Author $author): BookCollection;
```

This being said; it is a convention and not a hard rule. If you are more comfortable using `get`-methods without an exception, or even `find`-methods with; just go for it. There is no repository-police (that I know of).

## Read vs. Write repositories

There is a famous line by a statistician named George Box, who wrote: "All models are wrong, some are useful." When it comes to domain value objects, it highly depends on the context what information they need to contain.

Let's continue with our `Book` example. A book might be viewed within different contexts. When searching for a book, it is useful to know the title, author, isbn, category and an excerpt of what the book is about. When it comes to the context of shipping (of physical books); this is useless information. With shipping we are more interested in the dimensions and the weight to be able to figure out the size of the packaging and the cost of the shipping.

Both of these models are "wrong" (they do not portray the entirety of the object), but they are useful in their context.

### Write repositories

Within your application there might be a database table that contains _all_ the information for the books. This information needs to be written (at least once) and maybe updated once in a while. For these cases we use a _Write Repository_. This is the repository that stores, updates and deletes the underlying data.

Because it is very likely you are writing to only one entity (at a time), a write repository is often a very small interface, like the `BookRepository` at the beginning of this post. You probably already have the `BookId` available, so there is no need for additional `find` methods. After you `get` your entity, you either store the (updated) entity, or remove it.

### Read repositories

In a lot of applications it is more common to _read_ than it is to write data. Again continuing with our `Book` example; it is likely you would need a page that displays the _details_ about a book. It will show: the cover, the title, author, long excerpt, and a lot of specifics like page count, book binding, etc.

It's also very likely you have a page that _lists_ books. This list will probably only contain an image of the cover, the title, author and a tiny excerpt, and a link to the detail page. This list can likely also be filtered by category, author or a search phrase.

Just looking at this example, we already see two different contexts: Listing and Details (there might be better terms; but this is just an example). And we also see that the listing has a few specifics filter options, while the details page does not. We can also determine that the amount of information needed for the listing page is way less than the detail page.

### Retrieving entities

Since the details page contains just about every piece of information we have about a book (in this example) it might make sense to use the write repository to retrieve the `Book` entity. We can then use it to show all the required details.

However, note that this will probably require you to add a bunch of helper methods to return certain information in a specific format. Instead, we might want to create a specific "View Model" and let _it_ contain all the information and helper methods.

### Retrieving context specific domain objects (View Models)

In the case of the listing, it makes sense to create a specific object that represents a book within the listing context. Such a domain object is often called a "View Model". It is a context specific representation of the actual book. For this context we can create a specific _Read Repository_ like this:

```php
namespace App\Books\Listing;

interface BookRepository
{
	public function findByAuthor(Author $author): BookCollection;

	public function findByCategory(Category $category): BookCollection;

	public function findBySearchPhrase(string $search_phrase): BookCollection;
}
```

The `BookCollection` here is also a specific collection in the `Listing` context, and it contains only `App\Books\Listing\Book` objects (or view models).

As you can see, a read repository (almost) only contains `find` methods. And the methods are only catered to the actual needs of the context.

> **Criteria Pattern / Specification pattern** As an alternative to having a bunch of `findByX` methods, you can also incorporate the Criteria or Specification pattern. In this pattern you create a single `findBy(Criteria $criteria)` method which retrieves all objects that satisfy the criteria.

While this may look like a great one-size-fits-all solution, you have to be careful not to be creating a Query Builder. Because for every `Criteria` you create, you need an implementation that is compatible with your repository's underlying data source. This can get cumbersome real quick. My advice would be to keep read repositories tiny and only catered to the actual requirements; not to possible needs in the future.

### Retrieving from projections

When your application gets bigger, and the amount of data gets larger, certain queries might become very slow. In our example we might be retrieving the name of the author of a book and other information through a relation to the `authors` table or other join- or pivot-tables. The amount of relations can make these queries very slow.

Since we only need certain information we can create a specific table containing all the information we need. This table can be filled by running the slow query once, and update the table every time a book gets updated. Such a table is called a "projection". The queries from these projections are very fast, since it no longer requires the joins.

This is a big advantage of the repository pattern. The only place where you need to change this logic is in the repository. The repository will still return the same objects with the same information, only the technical implementation has changed.

### Retrieving from external services

Another example of a technical implementation is the use of search engines. The `findBySearchPhrase(string $search_phrase)` method can initially be implemented by some simple ORM or database logic. But these abilities are limited to what the database can do. A search engine is beter catered to the indexing of information, and can even locate data based on mistyped words.

So when your application matures; you can replace the implementation of these methods to use a proper external service, and map the returning data onto the View models. Again without changing the usage of the repository.

## When _not_ to use the Repository Pattern in PHP

As with most patterns; the repository pattern isn't a "once-size-fits-all-solution". So lets look at a few situations in which the repository pattern probably isn't a good fit.

### Small CRUD applications

If you have a small app that is primarily CRUD, using the repository pattern is probably not necessary. When your entities are small, and you only use them with an ORM, the repository pattern will only be a wrapper around the ORM; which will not provide many benefits. Of course there is the testability aspect; but if your framework already supports testing for the ORM, it won't add much.

### Large reports

When your application creates big reports using complex custom queries, or if your result sets are so large you need to paginate the results for displaying purposes; it would make more sense to use an approach like [CQRS (Command Query Responsibility Segregation)](https://martinfowler.com/bliki/CQRS.html) . In that case you would create a context specific query object to retrieve the domain objects according to your criteria.

## Summary & Links

In this (rather lengthy) post we've explored what the Repository pattern is and what it is _not_, and also when you might (not) want to use it. This post, like the rest of the Patterns series is meant to be informative only, even if my personal likings might shine through.

To summarise some things we explored:

- A repository is a data source abstraction that acts like a collection to retrieve and update domain objects.
- It does not know about any technical implementation, be it a database / ORM or in-memory variant, and is therefor completely decoupled from them.
- It is not meant to replace your ORM, or meant to be able to replace the database layer; but it can work together _with_ an ORM.
- Write repositories are used to get, store or delete domain objects like entities.
- Read repositories are used in different contexts to find context specific models of your domain objects for reading purposes.

Thank you for reading! If you found a mistake, or have questions I might be able to answer; please let me know in the comments below. I'm always out to learn new things myself and improve any (false) believes I might hold.

You might also consider [following me on Twitter / X](https://twitter.com/intent/follow?screen_name=doekenorg) as that is the social media I'm most active on.

Here is a short list of the links that are mentioned in this post, or might be interesting to you.

- [Decorator vs. Proxy Pattern](https://doeken.org/blog/decorator-vs-proxy-pattern) in the [Patterns for the Rest of Us](https://doeken.org/blog/categories/patterns-for-the-rest-of-us)series
- [Repositories and their true purpose](https://muhammedsari.me/repositories-and-their-true-purpose) by Muhammed Sari
- [CQRS](https://martinfowler.com/bliki/CQRS.html) as explained by Martin Fowler
- [Stop Mocking About](https://doeken.org/blog/stop-mocking-about-event-dispatcher)

Syntax highlighting by [Torchlight](https://torchlight.dev/).
