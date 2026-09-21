---
title: "Using php-fpm as a simple built-in async queue"
notion_id: ad26e863-ff31-4efd-a890-8dc242c92a3c
notion_url: https://app.notion.com/p/Using-php-fpm-as-a-simple-built-in-async-queue-ad26e863ff314efda8908dc242c92a3c
last_edited: 2023-04-25T14:16:00.000Z
source_url: https://tideways.com/profiler/blog/using-php-fpm-as-a-simple-built-in-async-queue
tags: ["Article", "tideways Blog", "English", "PHP"]
---
[https://tideways.com/profiler/blog/using-php-fpm-as-a-simple-built-in-async-queue](https://tideways.com/profiler/blog/using-php-fpm-as-a-simple-built-in-async-queue)

PHP Performance11.08.2017

There are many tasks that a web-request should not perform directly so the user doesn’t have to wait many seconds for a response. Examples include sending emails, uploading images to a CDN, resizing images or making expensive call to external services and many more depending on your use-case.

The usual advice you find on the internet is to setup a queue such as RabbitMQ, Redis, Kafka, Gearman or Beanstalkd. But this means another service that you need to install, setup, maintain and monitor. With some of the queue systems operating them includes a steep learning phase that requires time and money for additional hardware.

But maybe you just need a poor mans version of an asynchronous queue without all the overhead? Then why not just use PHP-FPM itself?

Maybe you are hosting on a platform that either runs only web-requests or makes it difficult to start and communicate with workers?

Granted, this is more of an experimental approach, but I think it is a perfectly valid way to use PHP-FPM sockets.

PHP-FPM already acts as a queue for Nginx/Apache FastCGI clients. While your web-request is running you can just send another FastCGI request to the same PHP-FPM socket asynchronously and non-blocking. This request is immediately executed in another php-fpm process in parallel and you could wait for it to complete or just fire and forget.

The good news: There is already a PHP library that support you with this task, so that you can skip learning about the FastCGI protocol, how to do asynchronous streaming in PHP and get going immediately with just a few lines of code.

The author Holger Woltersdorf has additional blog posts that combine his library with Redis or RabbitMQ, my experimental solution skips these middleman and talks directly to PHP-FPM without requiring another queue.

```plain text
composer require hollodotme/fast-cgi-client:^1.0
```

Then in a web-request just the client to send an async request:

```plain text
<?php // /var/www/index.php use hollodotmeFastCGIClient; use hollodotmeFastCGISocketConnectionsUnixDomainSocket; use hollodotmeFastCGIRequestsPostRequest; $connection = new UnixDomainSocket( '/var/run/php/php7.0-fpm.sock', 5000, 5000 ); $client = new Client($connection); $content = http_build_query([ 'task' => 'SendMail', 'payload' => json_encode('...') ]); $request = new PostRequest('/var/www/worker.php', $content); $client->sendAsyncRequest($request);
```

In the worker.php script we are calling, we can unpack the payload and delegate it to a handler:

```plain text
<?php // /var/www/worker.php namespace MyProjectCommands; class SendMailCommand { public function handle(array $payload) { // send mail here } } $class = sprintf('MyProjectCommands%sCommand', $_POST['task']); $command = new $class(json_decode($_POST['payload'], true)); $worker->handle($command);
```

Voila! We have built ourselves a simple asynchronous queue.

The code is intentionally simple, you can add a layer of abstraction around this if you want to improve the code, add error handling and more.

There are obvious downsides to this approach that I don’t want to sweep under the rug:

- Most important, this is an experimental approach, since this is not the primary way of using php-fpm.
- Your long(er) running worker jobs can block legimate web-requests from execution and take away resources from your webserver unless you increase simoultaneous requests per pool, setup another FPM pool for this or talk to a remote PHP-FPM.
- Jobs are not persistent and there is no retry for failed tasks or built-in metrics for monitoring. You can however build all this into your setup yourself. For monitoring, Tideways works out of the box with this approach and can monitor response, error-rates and collect traces for you.

For simple asynchronous tasks such as sending a registration or password forgotten mail, sending requests to third party analytics services or uploading images to a CDN this approach is a nice workaround before starting with more complex queuing setups. If you are willing to live with the downsides then you already have a “perfect” queue with php-fpm.

### About the author

Benjamin Founder & CEO

I am the founder and CEO of Tideways. I started the company over 10 years ago with the mission to move the PHP ecosystem forward, starting with performance. As managing director, I work across product, strategy, and the day-to-day of building a developer-focused SaaS business.

I’m a core contributor to the Doctrine open-source project and a founding board member of the PHP Foundation, which reflects my long-standing commitment to the PHP ecosystem. I particularly enjoy working at the intersection of application performance, developer experience, and the open-source community that makes PHP what it is. Outside of work, I enjoy board games, hiking, and coffee.
