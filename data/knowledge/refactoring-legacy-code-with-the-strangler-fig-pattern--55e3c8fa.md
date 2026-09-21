---
title: "Refactoring Legacy Code with the Strangler Fig Pattern"
notion_id: 55e3c8fa-29f3-4f83-822c-11052102c34c
notion_url: https://app.notion.com/p/Refactoring-Legacy-Code-with-the-Strangler-Fig-Pattern-55e3c8fa29f34f83822c11052102c34c
last_edited: 2026-09-21T17:40:00.000Z
source_url: https://shopify.engineering/refactoring-legacy-code-strangler-fig-pattern
tags: ["System Design / Software Architecture", "Product Management", "Project Management", "Article", "Shopify Engineering", "English"]
---
[https://shopify.engineering/refactoring-legacy-code-strangler-fig-pattern](https://shopify.engineering/refactoring-legacy-code-strangler-fig-pattern)

Large objects are a code smell: overloaded with responsibilities and dependencies, as they continue to grow, it becomes more difficult to define what exactly they’re responsible for. Large objects are harder to reuse and slower to test. Even worse, they cost developers additional time and mental effort to understand, increasing the chance of introducing bugs. Unchecked, large objects risk turning the rest of your codebase into a ball of mud, but fear not! There are strategies for reducing the size and responsibilities of large objects. Here’s one that worked for us at Shopify, an all-in-one commerce platform supporting over one million merchants across the globe.

As you can imagine, one of the most critical areas in Shopify’s Ruby on Rails codebase is the Shop model. Shop is a hefty class with well over 3000 lines of code, and its responsibilities are numerous. When Shopify was a smaller company with a smaller codebase, Shop’s purpose was clearer: it represented an online store hosted on our platform. Today, Shopify is far more complex, and the business intentions of the Shop model are murkier. It can be described as a God Object: a class that knows and does too much.

My team, Kernel Architecture Patterns, is responsible for enforcing clean, efficient, scalable architecture in the Shopify codebase. Over the past few years, we invested a huge effort into componentizing Shopify’s monolithic codebase (see Deconstructing the Monolith) with the goal of establishing well-defined boundaries between different domains of the Shopify platform.

Not only is creating boundaries at the component-level important, but establishing boundaries between objects within a component is critical as well. It’s important that the business subdomain modelled by an object is clearly defined. This ensures that classes have clear boundaries and well-defined sets of responsibilities.

Shop’s definition is unclear, and its semantic boundaries are weak. Unfortunately, this makes it an easy target for the addition of new features and complexities. As advocates for clean, well-modelled code, it was evident that the team needed to start addressing the Shop model and move some of its business processes into more appropriate objects or components.

## Using the ABC Code Metric to Determine Code Quality

Knowing where to start refactoring can be a challenge, especially with a large class like Shop. One way to find a starting point is to use a code metric tool. It doesn’t really matter which one you choose, as long as it makes sense for your codebase. Our team opted to use Flog, which uses a score based on the number of assignments, branches and calls in each area of the code to understand where code quality is suffering the most. Running Flog identified a particularly disordered portion in Shop: store settings, which contains numerous “global attributes” related to a Shopify store.

## Refactoring Shop with the Strangler Fig Pattern

Extracting store settings into more appropriate components offered a number of benefits, notably better cohesion and comprehension in Shop and the decoupling of unrelated code from the Shop model. Refactoring Shop was a daunting task—most of these settings were referenced in various places throughout the codebase, often in components that the team was unfamiliar with. We knew we’d potentially make incorrect assumptions about where these settings should be moved to. We wanted to ensure that the extraction process was well laid out, and that any steps taken were easily reversible in case we changed our minds about a modelling decision or made a mistake. Guaranteeing no downtime for Shopify was also a critical requirement, and moving from a legacy system to an entirely new system in one go seemed like a recipe for disaster.

## What is the Strangler Fig Pattern?

The solution? Martin Fowler’s Strangler Fig Pattern. Don’t let the name intimidate you! The Strangler Fig Pattern offers an incremental, reliable process for refactoring code. It describes a method whereby a new system slowly grows over top of an old system until the old system is “strangled” and can simply be removed. The great thing about this approach is that changes can be incremental, monitored at all times, and the chances of something breaking unexpectedly are fairly low. The old system remains in place until we’re confident that the new system is operating as expected, and then it’s a simple matter of removing all the legacy code.

That’s a relatively vague description of the Strangler Fig Pattern, so let’s break down the 7-step process we created as we worked to extract settings from the Shop model. The following is a macro-level view of the refactor.

Macro-level view of the Strangler Fig Pattern

We’ll dive into exactly what is involved in each step, so don’t worry if this diagram is a bit overwhelming to begin with.

## Step 1: Define an Interface for the Thing That Needs to Be Extracted

Define the public interface by adding methods to an existing class, or by defining a new model entirely

The first step in the refactoring process is to define the public interface for the thing being extracted. This might involve adding methods to an existing class, or it may involve defining a new model entirely. This first step is just about defining the new interface; we’ll depend on the existing interface for reading data during this step. In this example, we’ll be depending on an existing Shop object and will continue to access data from the shops database table.

Let’s look at an example involving Shopify Capital, Shopify’s finance program. Shopify Capital offers cash advances and loans to merchants to help them kick-start their business or pursue their next big goal. When a merchant is approved for financing, a boolean attribute, locked_settings, is set to true on their store. This indicates that certain functionality on the store is locked while the merchant is taking advantage of a capital loan. The locked_settings attribute is being used by the following methods in the Shop class:

We already have a pretty clear idea of the methods that need to be involved in the new interface based on the existing methods that are in the Shop class. Let’s define an interface in a new class, SettingsToLock, inside the Capital component.

As previously mentioned, we’re still reading from and writing to a Shop object at this point. Of course, it’s critical that we supply tests for the new interface as well.

We’ve clearly defined the interface for the new system. Now, clients can start using this new interface to interact with Capital settings rather than going through Shop.

## Step 2: Change Calls to the Old System to Use the New System Instead

Replace calls to the existing “host” interface with calls to the new system instead

Now that we have an interface to work with, the next step in the Strangler Fig Pattern is to replace calls to the existing “host” interface with calls to the new system instead. Any objects sending messages to Shop to ask about locked settings will now direct their messages to the methods we’ve defined in Capital::SettingsToLock.

In a controller for the admin section of Shopify, we have the following method:

This can be changed to:

A simple change, but now this controller is making use of the new interface rather than going directly to the Shop object to lock settings.

## Step 3: Make a New Data Source for the New System If It Requires Writing

New data source

If data is written as a part of the new interface, it should be written to a more appropriate data source. This might be a new column in an existing table, or may require the creation of a new table entirely.

Continuing on with our existing example, it seems like this data should belong in a new table. There are no existing tables in the Capital component relevant to locked settings, and we’ve created a new class to hold the business logic—these are both clues that we need a new data source.

The shops table currently looks like this in db/schema.rb

We create a new table, capital_shop_settings_locks, with a column locked_settings and a reference to a shop.

The creation of this new table marks the end of this step.

## Step 4: Implement Writers in the New Model to Write to the New Data Source

Implement writers in the new model to write data to the new data source and existing data source

The next step in the Strangler Fig Pattern is a bit more involved. We need to implement writers in the new model to write data to the new data source while also writing to the existing data source.

It’s important to note that while we have a new class, Capital::SettingsToLock, and a new table, capital_shop_settings_locks, these aren’t connected at the moment. The class defining the new interface is a plain old Ruby object and solely houses business logic. We are aiming to create a separation between the business logic of store settings and the persistence (or infrastructure) logic. If you’re certain that your model’s business logic is going to stay small and uncomplicated, feel free to use a single Active Record. However, you may find that starting with a Ruby class separate from your infrastructure is simpler and faster to test and change.

At this point, we introduce a record object at the persistence layer. It will be used by the Capital::SettingsToLock class to read data from and write data to the new table. Note that the record class will effectively be kept private to the business logic class.

We accomplish this by creating a subclass of ApplicationRecord. Its responsibility is to interact with the capital_shop_settings_locks table we’ve defined. We define a class Capital::SettingsToLockRecord, map it to the table we’ve created, and add some validations on the attributes.

Let’s add some tests to ensure that the validations we’ve specified on the record model work as intended:

Now that we have Capital::SettingsToLockRecord to read from and write to the table, we need to set up Capital::SettingsToLock to access the new data source via this record class. We can start by modifying the constructor to take a repository parameter that defaults to the record class:

Next, let’s define a private getter, record. It performs find_or_initialize_by on the record model, Capital::SettingsToLockRecord, using shop_id as an argument to return an object for the specified shop.

Now, we complete this step in the Strangler Fig Pattern by starting to write to the new table. Since we’re still reading data from the original data source, we‘ll need to write to both sources in tandem until the new data source is written to and has been backfilled with the existing data. To ensure that the two data sources are always in sync, we’ll perform the writes within transactions. Let’s refresh our memories on the methods in Capital::SettingsToLock that are currently performing writes.

After duplicating the writes and wrapping these double writes in transactions, we have the following:

The last thing to do is to add tests that ensure that lock and unlock are indeed persisting data to the new table. We control the output of SettingsToLockRecord’s find_or_initialize_by, stubbing the method call to return a mock record.

At this point, we are successfully writing to both sources. That concludes the work for this step.

## Step 5: Backfill the New Data Source with Existing Data

Backfill the data

The next step in the Strangler Fig Pattern involves backfilling data to the new data source from the old data source. While we’re writing new data to the new table, we need to ensure that all of the existing data in the shops table for locked_settings is ported over to capital_shop_settings_locks.

In order to backfill data to the new table, we’ll need a job that iterates over all shops and creates record objects from the data on each one. Shopify developed an open-source iteration API as an extension to Active Job. It offers safer iterations over collections of objects and is ideal for a scenario like this. There are two key methods in the iteration API: build_enumerator specifies the collection of items to be iterated over, and each_iteration defines the actions to be taken out on each object in the collection. In the backfill task, we specify that we’d like to iterate over every shop record, and each_iteration contains the logic for creating or updating a Capital::SettingsToLockRecord object given a store. The alternative is to make use of Rails’ Active Job framework and write a simple job that iterates over the Shop collection.

Some comments about the backfill task: the first is that we’re placing a pessimistic lock on the Shop object prior to updating the settings record object. This is done to ensure data consistency across the old and new tables in a scenario where a double write occurs at the same time as a row update in the backfill task. The second thing to note is the use of a logger to output information in the case of a persistence failure when updating the settings record object. Logging is extremely helpful in pinpointing the cause of persistence failures in a backfill task such as this one, should they occur.

We include some tests for the job as well. The first tests the happy path and ensures that we're creating and updating settings records for every Shop object. The other tests the unhappy path in which a settings record update fails and ensures that the appropriate logs are generated

After writing the backfill task, we enqueue it via a Rails migration:

Once the task has run successfully, we celebrate that the old and new data sources are in sync. It’s wise to compare the data from both tables to ensure that the two data sources are indeed in sync and that the backfill hasn’t failed anywhere.

## Step 6: Change the Methods in the Newly Defined Interface to Read Data from the New Source

Change the reader methods to use the new data source

The remaining steps of the Strangler Fig Pattern are fairly straightforward. Now that we have a new data source that is up to date with the old data source and is being written to reliably, we can change the reader methods in the business logic class to use the new data source via the record object. With our existing example, we only have one reader method:

It’s as simple as changing this method to go through the record object to access locked_settings:

## Step 7: Stop Writing to the Old Source and Delete Legacy Code

Remove the now-unused, “strangled” code from the codebase

We’ve made it to the final step in our code strangling! At this point, all objects are accessing locked_settings through the Capital::SettingsToLock interface, and this interface is reading from and writing to the new data source via the Capital::SettingsToLockRecord model. The only thing left to do is remove the now-unused, “strangled” code from the codebase.

In Capital::SettingsToLock, we remove the writes to the old data source in lock and unlock and get rid of the getter for shop. Let’s review what Capital::SettingsToLock looks like.

After the changes, it looks like this:

We can remove the tests in Capital::SettingsToLockTest that assert that lock and unlock write to the shops table as well.

Last but not least, we remove the old code from the Shop model, and drop the column from the shops table.

With that, we’ve successfully extracted a store settings column from the Shop model using the Strangler Fig Pattern! The new system is in place, and all remnants of the old system are gone.

## Takeaways

In summary, we’ve followed a clear 7-step process known as the Strangler Fig Pattern to extract a portion of business logic and data from one model and move it into another:

1. We defined the interface for the new system.
2. We incrementally replaced reads to the old system with reads to the new interface.
3. We defined a new table to hold the data and created a record for the business logic model to use to interface with the database.
4. We began writing to the new data source from the new system.
5. We backfilled the new data source with existing data from the old data source.
6. We changed the readers in the new business logic model to read data from the new table.
7. Finally, we stopped writing to the old data source and deleted the remaining legacy code.

The appeal of the Strangler Fig Pattern is evident. It reduces the complexity of the refactoring journey by offering an incremental, well-defined execution plan for replacing a legacy system with new code. This incremental migration to a new system allows for constant monitoring and minimizes the chances of something breaking mid-process. With each step, developers can confidently move towards a refactored architecture while ensuring that the application is still up and tests are green. We encourage you to try out the Strangler Fig Pattern with a small system that already has good test coverage in place. Best of luck in future code-strangling endeavors!
