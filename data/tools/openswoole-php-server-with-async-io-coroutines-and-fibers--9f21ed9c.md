---
title: "OpenSwoole - PHP Server with async IO, coroutines and fibers"
notion_id: 9f21ed9c-6cb7-4188-b3df-28946ddb0150
notion_url: https://app.notion.com/p/OpenSwoole-PHP-Server-with-async-IO-coroutines-and-fibers-9f21ed9c6cb74188b3df28946ddb0150
last_edited: 2026-09-21T17:21:00.000Z
source_url: https://openswoole.com/
tags: ["English", "PHP", "Untried", "Framework/Library"]
---
[https://openswoole.com/](https://openswoole.com/)

Compared with other async programming frameworks or software such as Nginx, Tornado, Node.js, Open Swoole is a complete async solution that has built-in support for async programming via fibers/coroutines, a range of multi-threaded I/O modules (HTTP Server, WebSockets, GRPC, TaskWorkers, Process Pools) and support for popular PHP clients like PDO for MySQL, Redis and CURL.

You can use sync or async, coroutine, fiber API to write the applications or create thousands of light weight fibers within one Linux process.

Open Swoole enhances the efficiency of your PHP applications and brings you out of the traditional stateless model, enabling you to focus on the development of innovative products at high scale, bringing event loops and asynchronous programming to the PHP language.

Event-driven, PHP Coroutine, PHP Fiber, Asynchronous API

Async TCP / UDP / HTTP / Websocket / HTTP2 / GRPC Client / Server Side API

IPv4 / IPv6 / UnixSocket / TCP / UDP and SSL / TLS / DTLS

Native PHP Coroutine and PHP Fiber Support

High performance, scalable, support C1000K

Milliseconds task scheduler

Multiprocessing and daemonize

PHP PSR support

GRPC server and clients

Free and Open Source (Apache 2 License)

### Get Started

- 1Installation# Linux users #!/bin/bash pecl install openswoole-26.2.0 # Mac users brew install php #!/bin/bash pecl install openswoole-26.2.0 # Install core library composer require openswoole/core:26.2.0
- 2HTTP Server<?php $server = new OpenSwoole\HTTP\Server("127.0.0.1", 9501); $server->on("start", function (OpenSwoole\Http\Server $server) { echo "OpenSwoole http server is started at http://127.0.0.1:9501\n"; }); $server->on("request", function (OpenSwoole\Http\Request $request, OpenSwoole\Http\Response $response) { $response->header("Content-Type", "text/plain"); $response->end("Hello World\n"); }); $server->start();
- 3WebSocket Server<?php $server = new OpenSwoole\Websocket\Server("127.0.0.1", 9502); $server->on('open', function($server, $req) { echo "connection open: {$req->fd}\n"; }); $server->on('message', function($server, $frame) { echo "received message: {$frame->data}\n"; $server->push($frame->fd, json_encode(["hello", "world"])); }); $server->on('close', function($server, $fd) { echo "connection close: {$fd}\n"; }); $server->start();
- 4TCP Server<?php $server = new OpenSwoole\Server("127.0.0.1", 9503); $server->on('connect', function ($server, $fd){ echo "connection open: {$fd}\n"; }); $server->on('receive', function ($server, $fd, $from_id, $data) { $server->send($fd, "OpenSwoole: {$data}"); $server->close($fd); }); $server->on('close', function ($server, $fd) { echo "connection close: {$fd}\n"; }); $server->start();
- 5Coroutine TCP Client<?php $client = new OpenSwoole\Client(\OpenSwoole\Constant::SOCK_TCP); if (!$client->connect('127.0.0.1', 9501, 0.5)) { exit("connect failed. Error: {$client->errCode}\n"); } $client->send("hello world\n"); echo $client->recv(); $client->close();
- 6Coroutine Redis / HTTP / WebSocket Client<?php // Enable coroutine support for Redis API co::set(['hook_flags' => OpenSwoole\Runtime::HOOK_ALL]); $redis = new Redis(); $redis->connect('127.0.0.1', 6379); $val = $redis->get('key'); echo $val; $http = new OpenSwoole\Coroutine\Http\Client('127.0.0.1', 80); $http->get('/index.php'); echo $http->body; $http->close();
- 7Coroutine Tasks<?php $server = new OpenSwoole\Http\Server("0.0.0.0", 9501); $server->set([ 'worker_num' => 1, 'task_worker_num' => 2, ]); $server->on('Request', function ($request, $response) use ($server) { $tasks[0] = ['time' => 0]; $tasks[1] = ['data' => 'openswoole.com', 'code' => 200]; $result = $server->taskCo($tasks, 1.5); $response->end('<pre>Task Result: '.var_export($result, true)); }); $server->on('Task', function (OpenSwoole\Server $server, $task_id, $worker_id, $data) { if ($server->worker_id == 1) { sleep(1); } $data['done'] = time(); $data['worker_id'] = $server->worker_id; return $data; }); $server->start();
- 8PSR HTTP Services<?php require __DIR__ . '/vendor/autoload.php'; use OpenSwoole\Core\Psr\Middleware\StackHandler; use OpenSwoole\Core\Psr\Response; use OpenSwoole\HTTP\Server; use Psr\Http\Message\ResponseInterface; use Psr\Http\Message\ServerRequestInterface; use Psr\Http\Server\MiddlewareInterface; use Psr\Http\Server\RequestHandlerInterface; $server = new Server('127.0.0.1', 9501); $server->on('start', function (Server $server) { echo "OpenSwoole http server is started at http://127.0.0.1:9501 "; }); class DefaultResponseMiddleware implements MiddlewareInterface { public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface { return (new Response('aaaa'))->withHeader('x-a', '1234'); } } $stack = (new StackHandler()) ->add(new DefaultResponseMiddleware()); $server->setHandler($stack); $server->start();
- 9GRPC Services<?php require __DIR__ . '/vendor/autoload.php'; use Helloworld\GreeterService; use Helloworld\StreamService; use OpenSwoole\GRPC\Middleware\LoggingMiddleware; use OpenSwoole\GRPC\Middleware\TraceMiddleware; use OpenSwoole\GRPC\Server; // enable hooks on IO clients co::set(['hook_flags' => OpenSwoole\Runtime::HOOK_ALL]); $server = (new Server('127.0.0.1', 9501)) ->register(GreeterService::class) ->register(StreamService::class) ->withWorkerContext('worker_start_time', function () { return time(); }) // use middlewares ->addMiddleware(new LoggingMiddleware()) ->addMiddleware(new TraceMiddleware()) ->set([ 'log_level' => \OpenSwoole\Constant::LOG_INFO, ]) ->start() ;
