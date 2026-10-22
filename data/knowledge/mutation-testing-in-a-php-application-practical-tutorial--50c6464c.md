---
title: "Mutation testing in a PHP application – practical tutorial"
notion_id: 50c6464c-02e0-4727-a443-cfb894e108fb
notion_url: https://app.notion.com/p/Mutation-testing-in-a-PHP-application-practical-tutorial-50c6464c02e04727a443cfb894e108fb
last_edited: 2026-09-21T17:41:00.000Z
source_url: https://tsh.io/blog/mutation-testing-php
tags: ["English", "PHP", "Testing", "Tutorial", "The Software House Blog"]
---
[https://tsh.io/blog/mutation-testing-php/](https://tsh.io/blog/mutation-testing-php/)

Marek Chuchro

2021-07-08 • 7 min read

table of contents

Share the article

It’s difficult to overestimate the importance of writing bulletproof tests for the quality of your application. Most developers worth their salt know that but their everyday practice is often limited to only the most popular types such as unit or integration testing.

Today, I’m going to show why you should also practice mutation testing and how to go about it in a PHP application.

Without a shadow of a doubt, you definitely need to test. The most popular trio of tests are:

unit tests,

integration tests,

end-to-end tests.

I’m only mentioning this because we will need to take a closer look at unit tests as they’re correlated with the topic of this article. Writing well-defined and implemented unit tests can protect you from making mistakes while implementing features of the application at the same time.

Unit tests (name’s not accidental) are focused on testing small parts of code, in other words - units of code. For example, we can say that the method of a class is a unit and it is a good practice to cover it with suitable tests to avoid mistakes in the development process.

the scroll of truth meme

## Mutation tests PHP – introduction

OK, now after the quick reminder of what unit tests are, let me introduce another concept related to testing – mutation testing. Why (if at all) you need them to write good quality failproof code for your application?

Mutation tests in PHP are not mandatory, but they can significantly improve the quality of not only the aforementioned unit tests but the implementation as a whole.

My Quality Assurance colleague has explored the introduction to mutation testing in detail. Before you move forward, I suggest you catch up with the basics there.

unit test + mutation tests meme

## How does mutation testing work?

The name of mutation tests came from the biological term referring to changing attributes inside DNA. In the context of testing, it means modifying existing code in some way to test whether the unit test is reacting properly – by detecting a mutation. Each modification case is called a mutant.

What kind of changes are performed in mutation? That depends on the framework used, but mutation rules usually try to mimic typical programmer mistakes such as using “<” instead of “>” and so on. The expected positive result of a single mutation test is called killing the mutant. In a negative case, the mutation is not revealed by the test and the result of a unit test is the same as in the original unmodified code.

The result of mutation tests is the ratio between escaped mutants and the killed ones, usually expressed by percentage referred to as coverage. Typically in a project, you should try to achieve the value of 100%.

All that information is returned to stout by default but we can redirect it to the report file.

## Infection framework – installation

The configuration is stored at the root of the PHP project.

Each node contains some information:

In the Source, you declare where code is stored.

In PHPUnit, you get a path to the PHPUnit binary.

Logs allow you to set a path where the output log should be stored.

In Mutators, you can select which mutators should be used; the good practice is using the default configuration, but you can declare it manually.

Let's see some examples of mutation cases.

## Examples of mutation cases

Let’s look at more examples showing potentially dangerous situations that you can discover with mutation tests. This method calculates the final product price. If the discount is too big, it returns an exception.

It uses this custom exception defined by class:

Test covering this case when the exception is thrown looks like this:

PHPUnit tests pass successfully.

PHP unit test

However, the output from mutation tests is different. Let’s look at the result log provided by the infection framework:

log infection

As you can see, this modification was not discovered by PHPUnit tests. Let's try to fix that.

The code above is modified to pass infection tests. We added an additional assertion to make sure the message of exception is valid. With this modification, it passes infection tests successfully.

Here’s another case. Let’s assume that a programmer used this function out of class so it was public. Now, it’s not used because the code was changed and someone forgot to change the accessor of the function.

The infection log will look like this:

infection result

You’ve probably noticed that this is not used outside the class. The fix for this case is rather easy – you just change the accessor to private.

## False-positive example

Not every mutant that escaped means that there’s something wrong with your unit test. For example, you can take mutation which changes arithmetic operator multiplication for division.

Let’s say you’ve got the following operation:

$x * $this->getDivider()

And for some reason in this test, the getDivider method always returns 1. There is no difference between multiplication time 1 and dividing by 1. So if mutated code is $x / $this->getDivider(), the result will still be the same and the unit test with the mutation will also pass. It’s called a false positive.

Another example is counting the price of a product:

The formula for counting price is $productPrice - $this->getDiscount($productId). It’s possible that one test will check this formula for a product that doesn’t have a discount and the function will always return 0 for this product.

We have the product price that equals 10 and the discount is 0. In that case, the final price of the product based on the formula is 10-0=10. If the code mutation changes the formula to $productPrice + $this->getDiscount($productId), then the price of the product will be calculated as 10+0. Therefore, it will be the same as the code without mutation. The test will pass and the mutant will escape. Those situations will lower the score of mutation tests but it doesn’t mean there is a problem in the code. It's a completely viable scenario that some products won’t have a discount and that the test works as it should.

Let’s look at this example of a false positive:

Here we have a simple class that calculates the final price of a product. The private function getGlobalDiscount returns a discount which is always counted, but for now, it’s equal to 0.

If you write a test for this, it will look like this:

Cases are provided by dataProvider and are equal to:

The PHPUnit tests will pass successfully. But if you run mutation tests, not all mutants will be killed:

not all mutant killed

As you can see inside the infection framework log, the mutant wasn't killed despite the change of operator to "-". This is because the mutation doesn’t affect passing PHPUnit tests. In this case, this is not dangerous because the GLOBAL_DISCOUNT can be equal to 0.

There are many more mutators. In order to find out how they work, you can read the infection documentation.

## Result report and what it means

Let’s discuss the console output result of infection tests:

console output

As you can see, there are some useful metrics here: numbers for generated mutation killed (or not), timeout, etc.

The metrics section gives you short information about the test quality in an application and helps you to find out if your tests are good and failproof:

Mutation Score Indicator – it’s a relation between killed and not killed mutants.

Mutation Code Coverage – it’s a percentage of covered code by PHPUnit tests.

Covered Code – in this example, the score is the same as the Mutation Score Indicator, because you’ve got 100% test coverage. Otherwise, those values will differ depending on coverage.

Let's get to CI/CD integration with mutation testing PHP.

## CI/CD Integration

By default, regardless of the result score (high or low), the process will exit with successful code unless there is an error in the process. In this form, it's not usable in the pipeline.

To make it usable inside the pipeline, you need to specify certain CLI options while running the test. Inside the Infection framework, three are two available options:

--min-msi – refers to the ratio of killed and not killed mutants. We can specify the percentage, which is the minimum at which the step will be passed.

--min-covered-msi – it's the same as the previous option but it counts only covered code mutants.

Here’s an example for Bitbucket Pipelines:

If the result of the step is below the requested score, the pipeline will fail.

the pipeline fails

If the result is higher or equal to the requested pipeline, it will pass.

pipeline succeeds

Putting the mutation testing as an additional step inside the CI/CD pipeline gives you another security layer. If a newly written feature is not tested well, it will not permit another step, like deploying to a staging environment. This feature is definitely worth a try but be aware that in big projects this step can take a lot of time. In cases like this, it can even be a disadvantage.

The Infection framework offers CLI options that are used to improve the process of mutation testing. Let's see some examples from the documentation:

-> If you have a powerful machine that supports multithreading and you expect mutation tests to be fast, you can use the --threads option. Then you can specify the number of threads which will be used in the testing process. By default, it is run on a single thread and it can be a bit annoying in bigger projects. In my experience with larger application mutation testing, it lasted 10 minutes for a single thread, which is a lot. By comparison, PHPUnit took 2 seconds.

-> If you want to check only a specific class, location, or file, we can use the --filter option. For example, you got the PriceCalculator class. To run mutation tests only for this class, you can use infection --filter=PriceCalculator.php.

Many more useful options can be found here.

## Mutation testing PHP – summary

Let’s summarize all the information about mutation testing PHP that you have acquired in this article.

Mutation tests are very helpful in the software development process to create higher quality tests, detecting malicious code changes and protecting us from potentially dangerous situations. Since we are not (yet?) machines, we make mistakes, bigger or smaller. But it should not affect the production environment.

You can also use it as a kind of benchmark for existing tests inside an application. With this kind of benchmark, you’ll know metrics for present tests, but also which tests can be improved. In my opinion, it’s worth using mutation tests in every project to improve the overall quality of tests.

### Authors

- Marek ChuchroPHP developer at work, Arduino/Raspberry PI enthusiast, video gamer and... grower of hot peppers off work.
