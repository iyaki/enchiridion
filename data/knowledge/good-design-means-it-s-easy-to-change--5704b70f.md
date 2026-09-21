---
title: "Good design means it's easy-to-change"
notion_id: 5704b70f-f2d0-4163-aebe-d45c22c65e30
notion_url: https://app.notion.com/p/Good-design-means-it-s-easy-to-change-5704b70ff2d04163aebed45c22c65e30
last_edited: 2023-01-25T19:53:00.000Z
source_url: https://matthiasnoback.nl/2022/09/good-design-means-easy-to-change/
tags: ["English", "System Design / Software Architecture", "Article", "Matthias Noback Blog"]
---
Software development seems to be about _change_: the business changes and we need to reflect those changes, so the requirements or specifications change, frameworks and libraries change, so we have to change our integrations with them, etc. Changing the code base accordingly is often quite painful, because we made it _resistant to change_ in many ways.

## Code that resists change

I find that not every developer notices the "pain level" of a change. As an example, I consider it very painful if I can't rename a class, or change its namespace. One reason could be that some classes aren't auto-loaded with Composer, but are still manually loaded with `require` statements. Another reason could be that the framework expects the class to have a certain name, be in a certain namespace, and so on. This may be something you personally don't consider painful, since you can avert the pain by simply not considering to rename or move classes.

Still, in the end, you know that a code base that resists change like this is going to be considered "a case of severe legacy code". That's because you can't resist change in the software development world, and eventually it will be time to make that change, and then you can experience the pain that you've been postponing for so long.

Software can resist change in many ways, here are just a few examples that come to mind:

- Classes have to be in a certain directory/namespace and methods have to have certain names in order to be picked up by the framework
- There's another application using the same database, which can't deal with extra columns that it doesn't know about
- Class auto-loading only works if you call `session_start()` first
- And so on... If you have another cool example, please add it as a comment below!

## Socially-established change aversion

Change-aversion can also be socially established. As an example, the team may use a rule that says "If you create a class, you also have to create a unit test for it". Which is very bad, because you can use multiple classes in a test and still call it a unit-test, so the every-class-is-a-unit assumption is plain wrong. More importantly, you can't unit-test all types of classes; some will require integrated tests. Anyway, let's not get carried away ;) My point is, if you have such a rule you'll make it harder for developers to add a new class, since they fear the additional (often pointless) work of creating a test for it. In a sense, developers start to resist change. The code base itself will resist change as well, because unit tests are often too close to the implementation, making a change in the design really hard.

Unit tests are my favorite example, but there are other socially-established practices that get in the way of change. Like, "don't change this code, because 5 years ago Leo touched it and we had to work until midnight to fix production". Or "we asked the manager for some time to work on this, but we didn't get it".

## Make changing things easy

From these, and many more - indeed - painful experiences, I have come to the conclusion that a very powerful way to judge the quality of code and design is to answer the question: is this easy to change? The change can be about a function name, the location of a file, installing a Composer dependency, injecting an additional constructor dependency, and so on.

However, it's sometimes really hard to perform this evaluation yourself, since as a long-time developer you may already be used to quite some "pain". You may be jumping through hoops to make a change, and not even realize that it's silly and should be much easier. This is where pair or mob/ensemble programming can be really useful: working together on the same computer will expose all the changes that you avoid:

- "Hey, let's rename that class!"
- "Well, I'm not sure that we can, let's save this for another time."
- "Now let's inject that new service as a constructor argument."
- "Sorry, we can't use dependency injection in this part of the code base."

That's why I usually go all-in on ensemble programming, so we can have a clear view on all the changes that the team averts. We look the monster in the eyes.

## Addendum: when changes break stuff

A partial reason for change aversion in developers is the risk that the change may break other things. If you rename a method, you should rename all the relevant method calls too. Luckily, we have static reflection these days, which will tell you about any call sites that you missed. And of course, the IDE can safely make the change for you in most cases. Unfortunately, this is not always the case. A lot of code in this world has the following issue: if you rename a file/class/method/etc., you won't be able to find all the places that you'd have to update. For instance:

- The class name might be used to derive a string name from (e.g. an `EmailValidator` can be used by dynamically invoking the `'email'` validator). Renaming the class breaks the validation.
- The controller action has to be called `public function [name]Action()`, or it can't be invoked. Renaming the class makes an entire route or endpoint unreachable.
- A model class has to be in `src/App/Entity` or it won't be added to the database schema.
- And so on!

The danger is, you can't see that you broke something unless you run the full application (or its tests, if you have full coverage). This is very inconvenient. It's why the general rule is "Good code is easy to change", and part of "easy to change" is that a change that wasn't fully propagated through the system will "explode" early. Besides an early warning, preferably a static error (as opposed to a runtime error), it would be great if the error message we get after breaking something is clear and helps us find the problem (the change that we just made). This is by far not always the case in real-world projects!

## Building a safety net to prevent breaking changes

One way to make your project safer-to-change is to look for these errors. When you make a change, and you learn (too late) about some issue it causes, make sure it can never happen again. As an example, recently I was working on a Symfony console command. I wanted to use the `QuestionHelper` which can be retrieved by calling `$this->getHelper('question')`. Being a dependency for one of my own services, I didn't want to get this helper on the spot, but I wanted to be properly set up as a dependency in the constructor, along with the other service dependencies. Unfortunately, the `'question'` helper is not available in the constructor, and you only learn this when you run the command in the terminal. To prevent the issue from happening again I added an integration test that verifies you can run the command and it will at least do something (get past the constructor). That way, the build will let me know when I accidentally broke the command.

Adding a test is a way of preventing a change-related issue at runtime. Maybe a better option is to prevent it _analysis time_, e.g. by writing a [PHPStan rule](https://github.com/matthiasnoback/php-ast-inspector/blob/main/utils/PHPStan/src/GetHelperRule.php) that triggers an error whenever you call `$this->getHelper()` in the constructor of a command:

```php
/**
 * @implements Rule<MethodCall>
 */
final class GetHelperRule implements Rule
{
    public function getNodeType(): string
    {
        return MethodCall::class;
    }

    /**
     * @param MethodCall $node
     */
    public function processNode(Node $node, Scope $scope): array
    {
        // ...

        if ($node->name->name !== 'getHelper') {
            // This is not a call to getHelper()
            return [];
        }

        // ...

        if (! $scope->getFunction()->getDeclaringClass()
            ->isSubclassOf(Command::class)) {
            // This is not a command class
            return [];
        }

        if ($scope->getFunctionName() !== '__construct') {
            // The call happens outside the constructor
            return [];
        }

        return [
            RuleErrorBuilder::message('getHelper() should not be called ...')
                ->build()
        ];
    }
}

```

Now whenever someone makes the mistake of calling `getHelper()` inside the constructor, they'll get this nice and useful error:

```shell
 15/15 [▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓] 100%

 ------ -------------------------------------------------------------
  Line   src/PhpAstInspector/Console/InspectCommand.php
 ------ -------------------------------------------------------------
  38     getHelper() should not be called in the constructor because
         helpers have not been registered at that point
 ------ -------------------------------------------------------------

```

By the way, if you want to become a _master at writing custom PHPStan rules_ that will make your code safer to change, get my latest book [Recipes for Decoupling](https://leanpub.com/recipes-for-decoupling/).

## Conclusion

In short, keep an eye out for the following situations:

1. You want to make a change but you can't because (legacy) reason X. Don't ignore this, or work around it, do something about it!
2. You make a change, but the change breaks something in a surprising way. Make sure you can't make the mistake again (add a test, or a static analysis rule).
