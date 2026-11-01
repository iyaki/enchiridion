---
title: "Understanding Domain-Driven Design"
notion_id: ff89f897-6cc7-4769-aa43-071ddbd3baa9
notion_url: https://app.notion.com/p/Understanding-Domain-Driven-Design-ff89f8976cc74769aa43071ddbd3baa9
last_edited: 2024-01-29T11:22:00.000Z
source_url: https://compiler.blog/series/domain-driven-design
tags: ["English", "Domain Driven Design", "Article", "Ozan Akman"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Introduction

Hey there! Today, I’m explaining a good way of developing enterprise software. I’ve spent roughly 10 years in the industry and experienced many different architectures.

Let me be honest with you, this series won’t be useful if you’re only looking for short-term–rapid application development. I emphasize “short-term” because this is actually faster in the long run. Why? Well, as your product grows to a couple dozen of use cases (if not a hundred), you’ll find yourself fixing a lot of issues. This can slow down your progress in many ways: the release of new features, and bug fixes, onboarding of a new programmer, refactoring, and eventually, creating a bad experience for both you and your users. We’ll explore these issues in detail.

## What is “Domain” in Domain-Driven Design?

In Domain-Driven Design (DDD), “domain” means the main area or topic that a software system is about. For instance, Facebook is about the social media domain, while Duolingo is about language and education. It’s the area we want to turn into software. This domain represents the real-world concepts, rules, and processes related to that space. The goal is to understand and model the complexities of that space in software. We want to ensure the software meets the requirements and behaves correctly, according to the domain.

In short, Domain-Driven Design helps you understand your business and users, and create software around their needs. Not influenced by your favorite framework or the preferences of IT people. End-users will never say “Oh look, they’re using the latest JavaScript framework!” unless your target audience consists solely of web developers.

## Who is a Domain Expert?

Domain experts are individuals with extensive knowledge in a particular field. In healthcare, these experts could be doctors or nurses. In finance, they might be economists. These experts give useful advice during the software design process. Close collaboration with others helps software developers understand the specific knowledge required to create software, instead of just imagining how things should work. This teamwork ensures that the software aligns with real-world needs and functions optimally. If you’re involved in creating a product for a certain time and making decisions within that area, you too are a domain expert!

## Why is Domain-Driven Design Important?

Domain-Driven Design (DDD) is important because it helps people to understand key questions like, “What are we solving?”, “Why are we solving this?”, “Who is this for?”. It helps you to understand that you’re solving people’s problems, not computer problems. By abstracting the domain and using it as the foundation for a software project, we can avoid being tied to specific technical terms, frameworks, or libraries. We can change any other parts of the project, as our foundation — the abstracted domain will remain strong. This means our software reflects real-world situations. We can have a shared understanding between our business and technical teams. This method makes it easier for software architectures to adapt and be modular, letting software systems change as requirements change.

## Practical Examples: Design a Blogging Platform with Domain-Driven Design (DDD)

Throughout some examples, you’ll realize that you have more freedom than you might expect, particularly when compared to a framework-centric approach.

We’ll be considering a few different strategies, including Screaming Architecture (also known as [Use-case Driven Approach](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html)), Test-Driven Development ([TDD](https://martinfowler.com/bliki/TestDrivenDevelopment.html)), and a bit of Command Query Responsibility Segregation ([CQRS](https://martinfowler.com/bliki/CQRS.html)). Don’t worry if you don’t know these terms, they sound fancy — they’re all tools designed to help solve software design problems.

As for the examples, I’ll be using PHP, which is my language of choice. However, the beauty of DDD is that it’s not about the language or the framework. It’s about designing a software system. So feel free to use any language you’re comfortable with. But before we start, please forget about how you learned to develop software with frameworks. The first thing you will do would be to install a framework because you will need routing, authentication, or validation, right? Not really, you don’t even have one user, yet. First, you need to abstract your domain away. You don’t need any frameworks or libraries other than a testing framework. Something that helps you to prove your use case works as expected.

### File Structure

Software design isn’t just about playing with files and folders, arranging them like they’re puzzle pieces. Many developers get stuck in the habit of thinking that designing software is all about arranging files in a certain pattern. What about frameworks? They’re very flexible but they still force their own semantics. They provide a predefined structure, rules, and many different tools, no questions asked, whether you’re building a big blogging platform or a calculator app.

We’re not jumping into picking a framework right off the bat. Sure, it’s convenient to have folders like `App/Models`, `App/Controllers`, or `App/Events` for your small projects. It’s like a cozy house where everything has its place. But when it comes to expanding an enterprise-level project with hundreds of use cases, this cozy house rapidly turns into a complex maze. Can you imagine a newcomer trying to navigate through this labyrinth? The confusion would lead to endless questions and slow progress.

So, instead of starting with a framework, we’ll take a different route. We’ll start by defining our domain model and Bounded Contexts.

### Bounded Context: The Building Blocks

When we talk about Domain-Driven Design (DDD), we use the term “Bounded Context” to describe different modules of our software system. Each Bounded Context is unique, focusing on its own tasks, rules, responsibilities, and concepts.

Let’s consider designing a blogging platform. What are the main components? You might think of elements like the blog posts, the authors, comments, and social interactions such as giving a like or dislike. Each of these elements addresses specific aspects. Let’s inspect each one more closely and give them specific names.

- Content Management: This context manages all the content-related entities such as `Blog Posts` and `Tags`.
- Comment: This context handles all entities related to comments, such as `Comments`, `Comment Replies`, and `Comment Feedback (Likes and Dislikes)`.
- Author Profile: This context manages the `Author` profile page (Bio, profile picture, etc.), and `Followers`.
- Social Interaction: This context covers interactions such as `Bookmarking`, `Post Feedback (Likes and Dislikes)`, and `Subscriptions (Follow and Unfollow)`.
- User Management: This context takes care of user-related entities such as `Settings`, `Privacy`, `Notifications`, and `Premium Subscription`.

### A couple of Insights

Something is interesting, the `User` represents the same person, but we treat them differently in different Bounded Contexts. Why is that?

- In the `User Management` context, a `User` can modify their personal settings, adjust privacy preferences, and manage notification settings.
- In the `Content Management` context, the same `User` is represented as an `Author` who creates, edits or deletes blog posts.
- On the other hand, in the `Social Interaction` context, a `User` might be someone who bookmarks posts, likes posts or follows other authors.

These actors may even share the same name, but they are three distinct actors, each with their own attributes, relationships, and rules within their respective contexts.

### Visualizing Bounded Contexts: Project Structure

One major change we’re making is reconsidering how we organize files and folders. Traditionally, developers have followed specific patterns for arranging them, which is largely dictated by the chosen framework. This approach is not the way for every software project. As we’ve discussed, Domain-Driven Design (DDD) focuses on the purpose of the software, which means our file structure needs to reflect our domain model, not the other way around. So, how do we translate the idea of Bounded Contexts into a practical file structure?

Let’s begin by creating a separate folder for each Bounded Context in our blogging platform. These folders will serve as the individual homes for each context’s unique rules, tasks, and concepts.

```plain text
└── src
    ├── AuthorProfile
    ├── Comment
    ├── ContentManagement
    |   ├── Post.php
    |   ├── PostRepository.php
    ├── SocialInteraction
    └── UserManagement

```

Here, each folder represents a Bounded Context. Even from this high-level view, you can understand the domain of this application without needing to dive deep into the codebase. This is the essence of Screaming Architecture. So, rather than screaming “This is a Laravel app!”, it screams “This is a blogging platform!”.

But hold on, what about routing? Don’t we need a framework or library for that? Let’s take a step back and remember that we are still in the process of defining our domain and use cases. We are not serving any users, yet.

### DDD Layers: Domain, Application, Infrastructure

In the DDD world, usually, the applications are layered into the domain, the application, the infrastructure, and perhaps the presentation. We’ve already discussed the domain layer—it’s the heart of our software, encapsulating the business rules and entities.

The application layer is like the conductor of our software orchestra; it uses the building blocks provided by the domain layer to accomplish specific tasks. Consider a task like `PlayMusic`, the application layer coordinates the domain objects like `Music` or `Artist` to execute this task, guiding the process from start to finish.

The infrastructure layer, on the other hand, is like the stage where our software orchestra performs. It provides the instruments required by the domain and application layers to function effectively. It supports the performance without dictating the music.

The presentation layer is where you perform your gig. It's the concert hall.

Imagine a scenario where your orchestra is restricted to playing in only one place—a stage where you can't move or replace a chair or an instrument. That’s what it feels like when you start your project by choosing a framework. You’re tied to that stage, and any attempt to change a component could break the whole setup. Instead, you should be able to replace infrastructural components at any time without causing a catastrophic failure.

Now, let’s address a common misconception: while these layers are key to designing your software, they don’t necessarily need to directly translate into your folder structure.

```plain text
└── UserManagement
    ├── Application
    ├── Domain
    └── Infrastructure

```

Many developers, myself included, have understood these layers as something to be explicitly represented in the codebase, resulting in folder structures like the one above. While this structure is often more manageable than a purely framework-based one or the traditional MVC folders, it still misses the point: DDD is not about folder structures, it’s about software design. The layers should guide your software design as a principle in your mind; it doesn’t have to be a literal representation in your folders. Simply put, you should have a clear mental picture of this layered architecture; it doesn’t have to be exactly reflected in a folder structure.

Now, let’s take a moment to consider the folder structures above – whether you’re an experienced developer or not, think about which one truly makes it easier to maintain the code?

### Our First Use Case: Creating a Blog Post

The first step in designing our blogging platform using Domain-Driven Design (DDD) is to focus on the use cases. Before we dive into routing, logging, or caching, let’s remember that these are infrastructural concerns and can be easily replaced. We shouldn’t be bound to specific tools; after all, tools are meant to assist us in achieving our goals.

Let’s start by learning the blueprint of our application. The journey begins from a Bounded Context and then drills down into specific use cases. You have the option to create additional folders to separate sub-contexts, like `src/ContentManagement/Post/CreatePost`, instead of `src/ContentManagement/CreatePost`. That’s totally fine! The key here is ensuring the folder structure is easy to understand and makes sense to anyone who looks at the codebase. It should focus on the domain logic rather than being overly technical or grouped by file types. We want it to be clear and straightforward so that developers can quickly grasp how everything is organized. With this approach, we can gain a clearer view of our initial design before refining it further.

```plain text
└──📁src
    └──📁Bounded Context
        ├──📁Use Case 1
        ├──📁Use Case 2
        ├──📄Context Object 1 (e.g. Domain Entity)
        ├──📄Context Object 2 (e.g. Repository Interface)
        └──📄Context Object 3 (e.g. Aggregate Root)

# This translates into something like the following in real-world projects.

└──📁src
    └──📁ContentManagement
        ├──📁CreatePost
        |     ├──📄CreatePost.php
        |     ├──📄PostCreated.php
        ├──📁GetPost
        ├──📄Author.php
        ├──📄Post.php
        ├──📄PublishedPosts.php
        ├──📄PostRepository.php
        └──📄PostRepositoryUsingDoctrine.php

```



## **Focusing on The Actual Problem**

Every domain has its unique set of problems. Let’s name a few: the travel industry, trade industry, real estate industry, auto industry, finance and banking industry, airline industry, commerce industry, and the list goes on… They all represent complex spaces. Our job is to abstract and simplify their solutions into a digital environment. Therefore, we need to focus on what matters. We want to design our software in a way that technical problems won’t affect us.

![image](https://cdn.hashnode.com/res/hashnode/image/upload/v1705845602185/27d2b696-56c5-4b70-a0a2-1e11d41dca3f.png?auto=compress,format&format=webp)

## **Are Frameworks That Bad?**

No, frameworks are not that bad. They’re just tools. They’re okay when you use them effectively. They’re there to help you. However, they become problematic when they dictate your project structure, dependencies, and development approach. They become problematic when you become too dependent on them. You shift your attention from your domain to the framework, a technical detail. When tackling your own issues, you go to the framework documentation to find the best solution. Frameworks don’t understand your problems and needs. What if they suggest bad practices? Have you ever thought about this? That’s not the way it should be. I’ll show you another way in this series. Please, keep up with me. I promise you a mind shift to Domain-Driven Design (DDD) and it will be easy! We will try to combine the best of both worlds.

## **Use Case-Driven Approach**

When modeling our domain, we will always think about its use cases. We will stick with the real-world concepts. Ask yourself, “What problem are you solving?”. For instance, imagine buying a flight ticket, before the computer era? Let’s break down this process step by step:

1. You would visit a travel agency or airline office.
2. Ask about available flights and ticket prices.
3. Choose a flight.
4. Make the payment.
5. Get your paper ticket, or the ticket might be mailed to your address.

This process is still the same, except now the agent is a website or a mobile app. All those steps are use cases. We still solve the same problem: buying a flight ticket. You search for flights online, view available flights and ticket prices, choose a flight, make the payment, and finally receive your ticket in your inbox. We have abstracted this entire process and made it more efficient and accessible with today’s digital systems. The `GetAvailableFlights` is a use case. The `BookFlight` is a use case. The `MakePayment` is another use case. We want to focus on abstracting this particular aspect of the business, initially. Not the routing, or not the caching. These are not even important, yet. We don’t serve any users. Caching is a concern for an application that has performance issues. We’re not there, yet.

This little mind shift will also affect how we write our tests. When we want to validate our use cases, we can always provide test doubles. You don’t have to run Redis or MySQL to prove that your use case–the important, works as expected. You can create an in-house `InMemory` implementations for different concerns. This will help you to test all of your use cases, without setting up a single infrastructural component. This is how you can start a project: create one folder and open a text editor. No need to install anything yet. This is freedom. We can always replace the `InMemory` drivers with the actual implementations, in the production configuration, before we launch. We simply need to provide an interface; A connection point between the infrastructure and the domain. This is where the mental model of separation comes into play. Don’t worry if this sounds a bit complex. We will dive into this topic later.

### **Good Software Architecture**

There is an important nuance with this structure and design. You need to understand that, as a software architect, you have to keep it clean and teach others how to keep it clean. The changes made in the UI components shouldn’t affect the use cases. Also, the changes in the use cases shouldn’t affect the domain entities. If you make changes in a REST controller, they shouldn’t change how your business operates. You need to have a strong mental model of separation. The location of files, next to each other or in different folders, shouldn’t really affect this mental model. You can organize things in a way that makes sense. A good software architecture allows for changes without high costs. It’s the one where you can make changes without saying “We need to create this project from scratch”.

### **Understanding a Use Case**

A Use Case is a simple set of instructions to the computer, describing our task. You minimize your contextual complexity within each Use Case. You solve a small problem. Also, to understand the same problem, anyone just needs to study a small folder. Isn’t that amazing? Just imagine how much less explanation you need when someone joins your team. And how much less cognitive load they will have. Let’s create our first Use Case, the `CreatePost`.

### **Creating a Use Case**

```php
namespace Platform\ContentManagement\CreatePost;

use DateTimeImmutable;
use PublishMate\Platform\ContentManagement\PostRepository;

final readonly class CreatePost
{
    public function __construct(
        private PostRepository $postRepository,
    ) {
    }

    public function create(
        string $title,
        string $content,
    ): void {
        $post = new Post($title, $content);
        $this->postRepository->save($post);
    }
}

```

That’s it? Yeah, that’s it! I told you it would be simple. I can hear you saying, “Bullshit! A real app would be much more complex!”. Let’s make it a bit more like a real-world app.

```php
namespace Platform\ContentManagement\CreatePost;

use DateTimeImmutable;
use PublishMate\Platform\ContentManagement\Author;
use PublishMate\Platform\ContentManagement\MarkdownConverter; use PublishMate\Platform\ContentManagement\Post;
use PublishMate\Platform\ContentManagement\PostId;
use PublishMate\Platform\ContentManagement\PostRepository;
use PublishMate\Platform\Shared\Event\EventDispatcher;

final readonly class CreatePost
{
    public function __construct(
        private PostRepository $postRepository,
        private EventDispatcher $eventDispatcher,
        private MarkdownConverter $markdownConverter,
    ) {
    }

    public function create(
        Author $author,
        string $title,
        string $slug,
        string $contentMarkdown,
        string $summary,
        string $canonicalUrl,
        ?string $coverImageUrl = null,
        DateTimeImmutable $createdAt = new DateTimeImmutable(),
        DateTimeImmutable $publishedAt = new DateTimeImmutable(),
    ): void {
        $htmlContent = $this->markdownConverter->toHtml($markdownContent);

        $post = new Post(
            author: $author,
            title: $title,
            slug: $slug,
            contentMarkdown: $contentMarkdown,
            contentHtml: $htmlContent,
            summary: $summary,
            canonicalUrl: $canonicalUrl,
            coverImageUrl: $coverImageUrl,
            createdAt: $createdAt,
            publishedAt: $publishedAt,
        );

        $this->postRepository->save($post);

        $this->eventDispatcher->dispatch(
            new PostCreated($post->id),
        );
    }
}

```

This is more or less a real-world use case. I will simplify the other examples for the article's purposes, but you can still understand my point. Of course, your domain has many use cases, dozens if not hundreds, and they all require different complex solutions, but your code doesn’t have to be complex. Coding is a tool to simplify those complex problems.

## **Testing a Use Case**

How do you ensure this use case works as expected? You create a test case. I usually start with a test case and then build my way up to the complete use case. Why? Because this helps me design for the expected outcome. This is called [**Test-Driven Development (TDD)**](https://martinfowler.com/bliki/TestDrivenDevelopment.html). It’s a technique that is often misunderstood. It’s not just about getting your test coverage to 100%. Ironically, it’s not even about testing; it is actually a design approach that helps you understand how the computer executes your instructions. I usually follow the AAA pattern:

- **Arrange** refers to the initial setup of your use case. This could be something like “I need a repository to save a post” or “I need an author to create a post”. Essentially, it’s about the requirements to execute the task.
- **Act** refers to the actual execution.
- **Assert** means comparing the expected outcome to the actual result to ensure they align.

Let’s apply each step:

```php
public function it_creates_blog_post(): void
{
    // Arrange
    $postRepository = new InMemoryPostRepository();
    $useCase = new CreatePost($postRepository);
}

```

Now that we have arranged the requirements, we can execute the task.

```plain text
public function it_creates_blog_post(): void
{
    // Arrange
    $postRepository = new InMemoryPostRepository();
    $useCase = new CreatePost($postRepository);

    // Act
    $useCase->create('Title', 'Content');
}

```

We can finally verify if the post was actually created.

```php
public function it_creates_blog_post(): void
{
    // Arrange
    $postRepository = new InMemoryPostRepository();
    $useCase = new CreatePost($postRepository);

    // Act
    $useCase->create('Title', 'Content');

    // Assert
    $actualPost = $postRepository->getLastCreatedPost();
    $expectedPost = Post::createWithId($actualPost->id, 'Title', 'Content');

    self::assertEquals($expectedPost, $actualPost);
}

```

Here we go—simple and easy. Once we run the test, we can confirm that our use case leads to the expected outcome. It is essential to test the behavior of the use case, not the implementation details. For instance, if you’re testing a database query, you need to ensure that the result is correct when you execute it. You don't need to validate if the SQL query is correctly written. Automated testing is a complex topic, and unfortunately, this is the only explanation I can provide in this series. But don’t worry, there’s good news! I plan to create an extensive guide about automated testing, too!

### **Delaying Important Decisions**

Good architecture allows you to delay major decisions like “What database are we going to use?”, “Which ORM tool should we pick?”, “What framework should we adopt?”. These are not the questions you should ask at the beginning of your project. They are tools, and you can pick them up whenever you need them. They should be plug-and-play. They shouldn’t dictate how you operate. This model enables you to replace any infrastructure parts without creating a huge fuss. As you might have noticed, we used the `InMemoryPostRepository`, which implements the `PostRepository` interface, just to validate our use case. We delayed the decision to choose a database. The Use Case-Driven Approach lets you focus on coding the important part rather than worrying about which tool to choose.

## **Gluing Use Cases and Frameworks**

Now that we have isolated our use case, we are free to choose any framework or tool we want. The process is simple: just inject the use case as a dependency and execute with the parameters. We still need to keep frameworks at a distance; we don't want to mix tools with our business. We want to be able to replace them whenever needed. Let’s create a basic Laravel controller using our use case:

```php
namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\View\View;

class CreatePostController extends Controller
{
    public function __construct(
        private CreatePost $createPostUseCase,
    ) {
    }

    public function createPost(Request $request): View
    {
        $this->createPostUseCase->createPost(
            $request->input('title'),
            $request->input('content'),
        );

        return view('post.created');
    }
}

```



## **The Big Picture**

The Use Cases are essentially a set of instructions to our system. They’re the big picture. That’s why, most of the time, testing a use case’s behavior already reveals underlying problems without delving deeper. But, there are some cases where the big picture is not the full picture. Sometimes, we need a little more detail. Let’s think about the “buying a flight ticket” process from [**the previous part.**](https://compiler.blog/understanding-domain-driven-design-part-2#heading-use-case-driven-approach)

### **Abstracting Use Cases Into Smaller Building Blocks**

In traditional apps, you might find all use cases handled by a single controller, such as `BookingController`. This controller would manage the entire process of buying a flight ticket. Checking availability, calculating prices, making reservations, processing payments, generating tickets, sending notifications, and the list continues. When the complexity of this controller becomes hard to manage, we look for ways to split it into smaller units.

Ultimately, we should anticipate and address those potential problems before they become problems, even in DDD apps. We have to be careful, though: we don’t want to overcomplicate our solutions while trying to simplify them. This distinction is not an easy challenge. It is hard. You need experience to do that properly. I will give you a little secret, though. When you focus on the behavior of the system rather than implementation details, it gets easier. What does that mean? And how do we focus on the behavior? We create a test case explaining the desired behavior. I explained this briefly in the previous article. We can leverage Test-Driven Development (TDD) to design our system rather than testing its implementation.

### **Testing The Behavior**

Imagine, you have a Symfony console command that converts CSV file into database entries.

```php
#[AsCommand(name: 'app:booking-importer:store-csv-file')]
final class StoreCsvFileCommand extends Command
{
    // ...
    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $entry = $this->csvToDbConverter->convert($input->getArgument('csvPath'));
        $this->bookingRepository->save($entry);

        $output->writeln('CSV content successfully converted!');

        return Command::SUCCESS;
    }
}

```

Let’s create an integration test case following the framework documentation. Typically, if you don’t understand the behavior’s importance, the test becomes an implementation test.

```php
public function test_execute(): void ❌ (A)
{
    $applicationContainer = // ...;
    $command = $applicationContainer->get(StoreCsvFileCommand::class);
    $commandTester = new CommandTester($command);
    $commandTester->execute(['csvPath' => 'someFile.csv']);

    // A private assertion method to validate command name
    self::assertLastExecutedCommandName('app:importer:store-csv-file'); ❌ (B)

    $actualOutput = $commandTester->getDisplay();
    $expectedOutput = 'CSV content successfully converted!';
    self::assertSame($expectedOutput, $actualOutput); ❌ (C)
}

```

Let’s review this test case together. I’ve already marked some problems with the ❌:

- A: The name doesn’t provide any information about the use case and its behavior.
- B: We are already retrieving the command by its class name. We don’t care about its configuration name. The configuration name doesn’t validate anything about the command’s behavior.
- C: The output text of the command is irrelevant when it comes to its behavior. The behavior is to store CSV entries in the database, not to output a particular text. How would I change this test case?

```php
public function test_it_converts_csv_file_into_database_entries(): void ✅ (A)
{
    $applicationContainer = // ...;
    $command = $applicationContainer->get(StoreCsvFileCommand::class);
    $commandTester = new CommandTester($command);
    $commandTester->execute(['csvPath' => 'someFile.csv']);

    $expectedEntries = [...]; // Array of expected entries
    $actualEntries = $applicationContainer->get(ImportedBookingRepository::class)->findAll();
    self::assertEquals($expectedEntries, $actualEntries); ✅ (C)
}

```

- A: We described the behavior of the use case in the test case name.
- B: The command name is not that important. It doesn’t impact the behavior; it’s a configuration detail. Once you validate this command’s behavior, you already know you’re testing the correct command, implicitly confirming the command name.
- C: We asserted against the expected outcome, not the command output. Once again, we can change this output message at any time. It won’t affect the use case itself. The test case also plays a role in documenting the use case behavior. To better understand the use case, we can break it down into steps:

1. We need a command
2. That will take a CSV file path
3. When we run this command, there will be corresponding database entries, generated from the CSV file.

This is a simple example, but you can already see the big picture from it. It’s much easier to think about smaller abstractions when you have that. Test-Driven Development (TDD) helps us see the big picture.

### **Why Do We Even Need Smaller Abstractions?**

The short answer is you don’t always need them. Sometimes, a simple use case is more than enough. Then, there are other times when your use case gets complicated and harder to maintain. It starts to hold different types of behaviors. It gets complicated. That’s usually the signal to start abstracting it into smaller units. There is no clear line between both decisions. You, as a software architect, need to draw that line between them. Running a business is never easy, and neither is the process of abstracting its problems. We need to find ways to decompose those complex problems into smaller, simpler ones. This approach helps us in validating the behavior of each small solution individually.

![image](https://cdn.hashnode.com/res/hashnode/image/upload/v1706510808091/f980ec52-3d3b-4dad-a3b7-5b0ef4049ae7.png?auto=compress,format&format=webp)

## **Strategic Abstraction**

If your use case is simple enough, such as our creating post example, you don’t have to extract it into a service. There is no real benefit to communicating through that service. This separation adds an unnecessary layer of complexity. You’re creating unnecessary dependencies. Dependencies aren’t just about third-party libraries; you also need to manage dependencies between components. It only makes sense to abstract your problem away if it’s complex enough. If you are unable to solve it, then try abstracting it into smaller problems.

If you remember the “real-world” example of `CreatePost` from the previous article, we abstracted the part that converts Markdown to HTML there. It’s already a big problem, not necessarily what we need to solve in the `CreatePost` use case. Now, our use case depends on that abstraction. Additionally, there is one abstraction that is almost a de facto industry standard for all Domain-Driven Design projects—it’s called “Repository”. I’m pretty sure you know about it. Have you ever asked yourself why we abstract this part of our applications, almost in all projects? The answer is simple: because we benefit from it. They encapsulate data access logic and help us focus on the domain rules and behavior. They protect the domain model from knowing the details of how we store or retrieve data. When we abstract this layer of details, it becomes really easy to test our use cases. We replace that repository abstraction with an `InMemory` implementation. So we don’t need to run a whole datastore infrastructure to validate our use cases. That abstraction solves a complex problem for us. In this case, the trade-off we’re taking for abstracting that layer is paying us off. That’s what you should always focus on.

![image](https://cdn.hashnode.com/res/hashnode/image/upload/v1706510819289/17c08b52-3560-4b71-855d-81b6388786af.png?auto=compress,format&format=webp)

## Artículos originales

1. [https://compiler.blog/understanding-domain-driven-design-part-1](https://compiler.blog/understanding-domain-driven-design-part-1)
2. [https://compiler.blog/understanding-domain-driven-design-part-2](https://compiler.blog/understanding-domain-driven-design-part-2)
3. [https://compiler.blog/understanding-domain-driven-design-part-3](https://compiler.blog/understanding-domain-driven-design-part-3)


