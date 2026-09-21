---
title: "Framework X – Simple and fast PHP micro framework. Async made easy"
notion_id: 4a9f938b-534a-4714-8e46-9a831127dce9
notion_url: https://app.notion.com/p/Framework-X-Simple-and-fast-PHP-micro-framework-Async-made-easy-4a9f938b534a47148e469a831127dce9
last_edited: 2023-01-13T01:34:00.000Z
source_url: https://framework-x.org/
tags: ["Framework/Library", "English", "PHP", "Untried"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Sneak a peek at the code!

Here’s what using Framework X in code looks like.

```plain text
<?php

require __DIR__ . '/../vendor/autoload.php';

$app = new FrameworkX\App();

$app->get('/', function () {
    return React\Http\Message\Response::plaintext(
        "Hello wörld!\n"
    );
});

$app->get('/users/{name}', function (Psr\Http\Message\ServerRequestInterface $request) {
    return React\Http\Message\Response::plaintext(
        "Hello " . $request->getAttribute('name') . "!\n"
    );
});

$app->run();

```

But that’s not all! Take a look at [our documentation](https://framework-x.org/docs/) to get a deep understanding of how everything works and how to build most common use-cases.

> 

[ As a side note, this is my first time using @x_framework in a project and I gotta say it's the most convenient framework I've ever used. ](https://twitter.com/josemmooo/status/1591771258521546752)

Framework X combines years of experience from existing real-world projects.

### Simple Web APIs

Building web APIs should be simple, lightweight and efficient. That’s what Framework X is all about. No matter if you’re building a RESTful or JSON-based HTTP API or the next big cat rating platform, we’ve got you covered. Framework X doesn’t get in your way, it paves your way.

 Framework X is super easy to get started with: If you’ve used PHP and standard PSR-7 HTTP messages before, you’ll feel at home immediately. Install the package with a single command and you’re ready to go.         No matter if you’re building a RESTful or JSON-based HTTP API or the next big cat rating platform, we’ve got you covered. Processing images, XML data or large CSV files? No problemo. 

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

```plain text
$app->get('/users', function () {
    $users = [['name' => 'Alice'], ['name' => 'Bob']];

    return React\Http\Message\Response::json(
        $users
    );
});

```

### Lightning fast

Framework X is super efficient. It’s a micro framework that is optimized for you as human (_to build fast_) and the computer (_to run fast_). Because Framework X is so efficient, it makes it much easier to scale.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

```plain text
$ ab -n100000 -c50 -k http://localhost:8080/
…
Concurrency Level:      50
Time taken for tests:   5.508 seconds
Complete requests:      100000
Failed requests:        0
Keep-Alive requests:    100000
Total transferred:      13300000 bytes
HTML transferred:       1300000 bytes
Requests per second:    18155.95 [#/sec] (mean)
Time per request:       2.754 [ms] (mean)
Time per request:       0.055 [ms] (mean, across all concurrent requests)
Transfer rate:          2358.15 [Kbytes/sec] received

```

### Supports async

Thanks to async programming, Framework X allows you to do multiple things at once. On top of this, its built-in web server also utilizes the async execution model to process multiple incoming connections at once. We support fibers, coroutines, and promises. Your choice!

 Imagine sending two HTTP requests to external APIs while already saving the request in the database at the same time. Async, non-blocking APIs allow much faster response times by doing multiple things at once, instead of having to do one thing after another.         Framework X helps you being prepared to scale to the next level. By utilizing async programming techniques, the built-in web server achieve maximum performance per CPU core. Parallel worker processes and server clusters allow even better performance.         To get best performance, we recommend using async programming. Additionally, Framework X provides the tools and guidelines so you can keep using existing, blocking code and integrate this into a non-blocking structure. This makes it easy to get started right away. 

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

```plain text
$app->get('/book/{isbn}', function (Psr\Http\Message\ServerRequestInterface $request) use ($db) {
    $isbn = $request->getAttribute('isbn');
    $result = await($db->query(
        'SELECT title FROM book WHERE isbn = ?',
        [$isbn]
    ));

    assert($result instanceof React\MySQL\QueryResult);
    $data = $result->resultRows[0]['title'];

    return React\Http\Message\Response::plaintext(
        $data
    );
});

```

### Reactive and streaming

Framework X has built-in support for reactive applications that are responsive, resilient, elastic and message-driven. It enables streaming protocols to transport structured and binary data over HTTP.

 The HTML5 `EventSource` API allows you to receive live updates in real-time as they happen on the server side. Because it uses standard HTTP requests, it works well on most stacks. When using the built-in web server, you can utilize efficient, long-lived connections. 

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

```plain text
$app->get('/users', function () use ($redis) {
    // stream messages received from Redis PubSub channel
    $stream = new React\Stream\ThroughStream();
    $redis->on('message', function ($message) use ($stream) {
        $stream->write("data: $message\n\n");
    });

    return new React\Http\Message\Response(
        React\Http\Message\Response::STATUS_OK,
        [
            'Content-Type' => 'text/event-stream'
        ],
        $stream
    );
});

```

### Runs anywhere

Framework X is scalable from the ground up. We provide a simple API that is easy to get started with for the smallest of projects, yet allow to grow to complex, production-ready setups. We're ready for PHP 8.1, but also support PHP 7.1+ with no extensions required!

 We support running behind existing PHP stacks using Apache, nginx or any other web server using CGI / FastCGI (PHP-FPM). This means you can create a simple web application in minutes and run it anywhere from shared hosting to your own clustered server setup in the clouds.         Framework X ships its own efficient web server implementation written in pure PHP. This uses an event-driven architecture to allow you to get the most out of Framework X. With no changes required, you can run the built-in web server with the exact same code base on the command line.         With the built-in web server, we provide a non-blocking implementation that can handle thousands of incoming connections and provide a much better user experience in high-load scenarios. 

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

```plain text
$ php -S localhost:8081 examples/index.php
PHP Development Server (http://localhost:8081) started

$ php examples/index.php
Listening on http://127.0.0.1:8080

$ curl http://localhost:8081/
Hello world!

$ curl http://localhost:8080/
Hello world!

```

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# Community

Framework X is so much more than _yet another framework_, it is the way to introduce asynchronous programming to the broader PHP community. We believe it's essential to involve the whole community in all our processes and provide you with the opportunity to bring up your awesome ideas.

Share your question with the community on Twitter or introduce your awesome project you created with Framework X. Just drop in, get in touch and connect with our community on Twitter.

We are where the open-source community is. [GitHub discussions](https://github.com/clue/framework-x/discussions) and the public [issue tracker](https://github.com/clue/framework-x/issues) are a great way to keep track of discussions and follow the project progress and upcoming enhancements.

You need help to include Framework X in your project or want to create a new awesome project? Get professional support for your dev team by [clue·engineering](https://clue.engineering/), the company behind Framework X.

## Quickstart

Get started in seconds

We can start by taking a look at a simple example application. You can use this example to get started by creating a new `public/` directory with an `index.php` file inside:

```plain text
<?php

require __DIR__ . '/../vendor/autoload.php';

$app = new FrameworkX\App();

$app->get('/', function () {
    return React\Http\Message\Response::plaintext(
        "Hello wörld!\n"
    );
});

$app->get('/users/{name}', function (Psr\Http\Message\ServerRequestInterface $request) {
    return React\Http\Message\Response::plaintext(
        "Hello " . $request->getAttribute('name') . "!\n"
    );
});

$app->run();

```

That's it already! You can now run the above example using any existing web server – or even more impressive: using the built-in web server like this:

```plain text
$ php public/index.php

```

And that’s all it takes! X will now listen for incoming requests using its efficient built-in web server. Can’t wait to try it out?

Framework X is released as open-source under the permissive MIT license. This means it is free as in _free speech_ and as in _free beer_.

We believe in open source and made a conscious decision to take this path. Being open-source means we can foster a community to focus on building the best possible framework together. Framework X builds on top of existing open-source projects and we want to give back to this community of awesome engineers and developers. Being open to outside contributions means we can guarantee interoperability with a vivid ecosystem and ensure the longevity of the project.

> 

I love open-source! ❤️ Open-source software ensures we can overcome market constraints and build software systems based on trust and mutual cooperations instead.

## Values

What you can expect

We ensure a long maintenance of the project with regular updates. This is an elementary part of our philosophy and is accompanied by our high quality standards for our projects.

We speak plainly, nothing happens behind closed doors. Framework X creates a high level of transparency about the current project status and its progress.

We use Framework X ourselves on a daily basis. Our firsthand experience ensures we have an interest in making sure everything works smoothly.

### Passionate and experienced

We are software engineers at heart. Framework X is not only a hobby project, it is our biggest project to solve problems and perhaps have a lasting impact on the larger PHP ecosystem.

We provide ways to allow users of Framework X to easily exchange information and ideas with each other. We offer a [discussion and chat platform](https://framework-x.org/#community) to provide direct community support.

Framework X is released under the permissive MIT license. You can have an impact on the further development of Framework X at any time. Join in, tell us about your ideas and leave your fingerprint.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

We are always interested in mutual cooperations that generate win-win situations for both sides. Think we have a match? Just approach us!

## Frequently asked questions

 Fair question, who likes reinventing the wheel, right? Unless of course, you’re sure you can build a better wheel. That’s what we’re doing with Framework X. We believe our [concepts and features](https://framework-x.org/#features) are somewhat unique.      Because like it or not, but PHP runs the world! It’s a huge ecosystem with a big demand for better solutions. This is why one of our main goals is to make sure Framework X [runs anywhere](https://framework-x.org/#runs-anywhere).      Building something like Framework X probably wouldn’t be possible if not for the excellent open-source projects that it builds on top of. In particular, it uses [ReactPHP](https://reactphp.org/) for its fast, event-driven architecture. Framework X is standing on the shoulders of giants.      It’s done when it’s done (_sorry_)! More seriously though, the code is already pretty complete and released as public beta. We will use this to process incoming feedback and push Framework X to be even better. This shouldn’t take too long (_pinky swear_!).      If you’re happy, so are we. But a lot of people think there’s room for improvement. That’s what we’re aiming for with our [concepts and features](https://framework-x.org/#features). We think Framework X is the best solution if you’re looking for simple web APIs, lightning fast execution, reactive features, and async support.      We love ReactPHP, its concepts and its versatility! At the same time, we realize it can be somewhat hard to wrap your head around without any structure. That’s why think Framework X can be the perfect _frame_·work if you want to build [simple web APIs](https://framework-x.org/#simple). 
