---
title: "Keep Your ORM Out of my Controller"
notion_id: 5d7b08aa-447b-43e8-a042-5f6159bb08ae
notion_url: https://app.notion.com/p/Keep-Your-ORM-Out-of-my-Controller-5d7b08aa447b43e8a0425f6159bb08ae
last_edited: 2023-04-25T13:36:00.000Z
source_url: https://ferd.ca/keep-your-orm-out-of-my-controller.html
tags: ["English", "Programming", "Object Oriented Programming", "Domain Driven Design", "Article", "My bad opinions (Fred Hebert)"]
---
[https://ferd.ca/keep-your-orm-out-of-my-controller.html](https://ferd.ca/keep-your-orm-out-of-my-controller.html)

One thing I've learned by working on a few legacy server applications is how a very small leak in an abstraction can be enough to make you lose weeks of work when tiptoeing around it to add or modify a feature, years later.

ORM usage as demonstrated in tutorials online (and therefore in many real life applications) frequently suffers from such leaks. This is a rant about it.

### ORMs Are Middleware, Not an Abstraction

Object-Relational Mappers are all about abstracting away how you fetch data: they hide SQL, they hide implementation specific idiosyncrasiesm etc. That's a fine goal in itself, but the problem is how it is being used instead.

I frequently see them used straight in the controller to access data from the model layer. After all, most ORMs let you model your object transparently over the database table and often support all the CRUD idioms you might need. Adding more feels like abstraction for the sake of it and a loss of time.

As a perfect study case, let's take the following example from the 5th chapter of the Django Book. The following snippet is named The "Dumb" Way to Do Database Queries in Views (views are Django's equivalent to controllers):

```plain text
def book_list(request): db = MySQLdb.connect(user='me', db='mydb', passwd='secret', host='localhost') cursor = db.cursor() cursor.execute('SELECT name FROM books ORDER BY name') names = [row[0] for row in cursor.fetchall()] db.close() return render_to_response('book_list.html', {'names': names})
```

Now, I agree this isn't the place to put SQL queries. The following solution is instead suggested:

```plain text
def book_list(request): books = Book.objects.order_by('name') return render_to_response('book_list.html', {'books': books})
```

That's much better, really. The book_list controller now has a quick access to the objects it needs and is reduced to a much shorter two lines of code.

The problem I see with this is that the really important details are still there. I can still see "Give me a list of all books ordered by name." Okay, that doesn't sound too bad. That's because the query is very simple. Imagine that your boss comes in and asks you to have a block on a page that returns a list of all books having more than 150 pages, written by British authors who were born in the past century. Moreover, it should be possible to limit the books to a specific category (say, biographies) and it'd be nice to sort them by title.

Suddenly, the query isn't as simple anymore. You begrudgingly replace this and add the query to your controller. It now looks something like (pardon the pseudo-code ORM, no longer Django's):

```plain text
Book.objects.filter('pages', larger_than=150).filter('author.nationality', 'British').filter('author.birthdate', larger_than=1900, smaller_than=2000).filter('category', 'Biography').order_by('name')
```

Oh. Well this is still nicer than the underlying SQL statement to be sure. No handling of joins, no handling of query building.

Now your boss has decided he wants this example search on 6 different pages (that require 6 different controllers). What do you do?

1. Copy/paste the query within all 6 controllers
2. Create a class with a method (or just a function) to handle it

Chances are the second solution is the best one. Not that we hate the ORM, but the query is pretty complex. What we really want might be something like "OldBritishBiographies()" and use that call in the 6 controllers instead. Now, why is it okay to use the ORM directly in the controller when the query is short, but we all cringe when it's long and ugly? Shouldn't we use only one way or the other?

I argue that all ORM calls should be wrapped in functions or methods. Otherwise, the programmer reading the code now has to understand something about 'author.birthdate', 'author.nationality' or yet 'name' (why not title?), etc. which are usually database column names.

This means that we're in the controller but we still have to grasp what the hell is going on in the database. The abstraction is leaking! We still need to know how the database works in order to fetch data from our models. That is because ORMs don't abstract away any applicative logic. All they abstract away is 'how you get the data', not 'what you are trying to get'. ORMs are middleware.

Any code base you have that will last a few years and scale up a bit will require a change in technology at some point. You want to move the book inventory to a web service, based on RESTful HTTP calls. Suddenly, none of the ORM calls above work without either tweaking the ORM a whole lot, or changing the calls to books in every controller that deals with them.

Except maybe that call about old British authors. It turns out this one only needs to be changed once, in its implementation. I suggest that this is because adding that descriptive function call completed the abstraction: you can get your results without caring about whatever the underlying implementation is. It could be done with or without an ORM, straight from text files, or even from a physical scanner with OCR. We just don't care, it needs to be implemented once and that's it.

The ORM is middleware between you and a database. A proper model is an abstraction between the data you get and whatever business logic is hidden behind it.

### But wait, there's more!

To add to this all, keeping the ORM outside of the controller and inside the model as suggested above makes it infinitely much easier to do any kind of unit or integration testing.

It becomes easy to use whatever integration testing method you want to make sure your code fetches and stores data correctly from and to the database. It also becomes easy to build mock objects or functions to pass to the controller to test it independently from the model implementation.

Conversely, with ORM queries in the controller, verifying that everything goes fine is limited to testing the ORM itself or the controller and the underlying model at once. There is no easy way to test only one or the other without some copy/paste abuse because the ORM queries only exist as part of the controller.

So on one hand you have code that is easier to modify (because you need to change it at only one place) and to test. On the other hand you have code that is harder to maintain (because it is spread across controllers) and harder to test, just to add a kick to the nuts of the developer responsible of doing it.

Of course, using an ORM in the controller is still leagues better than what I'm having to maintain as described earlier. However, you should not assume your application won't live for more than a few years or that you do not need things such as 'being easier to test'. Please, for the sake of programmers doing maintenance on your code in the future, do your abstractions right. A small leak today is a flood tomorrow.
