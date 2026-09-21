---
title: "Easy way to keep background PHP jobs alive"
notion_id: 1d054f1c-7d23-8170-a8c9-cab7bc5b09a7
notion_url: https://app.notion.com/p/Easy-way-to-keep-background-PHP-jobs-alive-1d054f1c7d238170a8c9cab7bc5b09a7
last_edited: 2025-07-26T22:21:00.000Z
source_url: https://www.algotech.solutions/blog/php/easy-way-to-keep-background-php-jobs-alive/
tags: ["PHP", "Tutorial", "English"]
---
![image](https://www.algotech.solutions/wp-content/uploads/2016/03/keep-background-php-jobs-alive.svg)

In modern web applications, there are many use cases when you need to run background jobs. Most of the times, these are time consuming processes and you shouldn’t let the user wait until the job is completed, but rather notify them while they are probably doing other stuff. It’s even more misleading when they see a blank page reloading for several seconds. Most of the users could think that something went wrong and will try to refresh the page, which is definitely not the desired user experience.

or [deleting unused files](https://www.algotech.solutions/blog/python/deleting-unused-django-media-files/), [enhanced validation techniques](https://www.algotech.solutions/blog/javascript/how-to-create-custom-validator-directives-with-angularjs/) etc. Today, let’s talk about running background jobs in PHP. Of course, feel free to apply any of these techniques to the programming language of your choice.

## Background PHP jobs

Let’s roll up our sleeves and see some examples. In order to handle the background jobs we are going to start several instances of a _worker_ process.

We will take a simple example first. We are starting a background PHP job on a unix machine:

```shell
$ php a_job.php list of params &
```

With the PHP job looking like this:

```php
<?php
echo "I am a PHP job.\n"

for ($i = 0; $i < 1000; $i++) {
	// Do something
}
```

Nothing special so far. Let’s add a little bit of fun here:

```php
<?php
echo "I am a PHP job\n";
for ($i = 0; $i < 1000; $i++) {
	// Do something
	// Oopsss!
	if (13 === rand(1, 100)) {
		die;
	}
}
```

Oopsss! It seems that our PHP job could randomly die. In a real case the script dying is rather caused by an uncaught exception or an error occurring rather than a [die](https://php.net/manual/en/function.die.php) call, as in this example. Anyway, whatever the cause of failure is, we need to keep the worker alive, since it takes care of our background processes. And yes, we can start as many instances as we want (as many as the machine can handle), but it’s just a matter of time until all instances will die. Also, manually monitoring them and starting them on cue is not a viable solution. If you want to read more about server firepower, try [my article on AWS t2 instances](https://www.algotech.solutions/blog/engineering/how-powerful-are-the-burstable-aws-t2-instances/), but, until then, let’s focus on finding an elegant solution to the task at hand.

Another possible solution is to make sure all the exceptions are caught, there is no error that could possibly occur, and there is no other thing that could stop the job from staying alive. In my opinion, this is by far [the best solution we could apply](https://www.algotech.solutions/blog/engineering/good-code/). However, in a real-world scenario it is not always suitable, especially when the tests are missing. Furthermore, writing tests is not always feasible, particularly when it would suppose a huge amount of work.

So, what are we going to do? PHP offers us the possibility to listen for the _shutdown_ event ([register_shutdown_function](https://php.net/manual/en/function.register-shutdown-function.php)) that a script is triggering when dying.

```php
<?php

function onShutdown() {
	echo "I am shutting down...\n";
}

register_shutdown_function(onShutdown);

echo "I am a PHP job\n";

for ($i = 0; $i < 1000; $i++) {
	// Do something
	// Oopsss!
	
	if (13 === rand(1, 100)) {
		die;
	}
}
```

Ok… we know when the job is dying. Now what? Let’s try something.

```php
<?php

$_ = $_SERVER['_'];

function onShutdown() {
	global $_, $argv;
	echo "I am restarting...\n";
	pcntl_exec($_, $argv);
}

register_shutdown_function(onShutdown);

echo "I am a PHP job\n";

for ($i = 0; $i < 1000; $i++) {
	// Do something
	// Oopsss!
	if (13 === rand(1, 100)) {
		die;
	}
}
```

Guess what? The job is actually restarting automatically whenever it should have died. This is done due to [pcntl_exec](https://php.net/manual/en/function.pcntl-exec.php) magic.

What is actually happening? The variable $_ stores the path to the PHP executable (e.g. “/usr/bin/php”), and $argv has all the arguments the start job command initially had, including the path to the PHP script we started. So, pcntl_exec($_, $argv) is the same with $ php a_job.php list of params & . Well, almost the same. _pcntl_exec_ executes the program in the current process space, so the process PID won’t change after restarting. If you need to restart the job in a new process space you could use [exec](https://php.net/manual/en/function.exec.php) instead of _pcntl_exec_.

```php
<?php

$_ = $_SERVER['_'];

function onShutdown() {
	global $_, $argv;
	echo "I am restarting...\n";
	exec($_, $argv);
}

register_shutdown_function(onShutdown);

echo "I am a PHP job\n";

for ($i = 0; $i < 1000; $i++) {
	// Do something
	// Oopsss!
	if (13 === rand(1, 100)) {
		die;
	}
}
```

Before going to a real use case I want to emphasize that I am not, in any way, encouraging poor application design or allowing broken code in your application, even though these could possibly be “caught” using the method described above. I strongly disagree to using it as a solution for a problem like that.

But what can be done in such cases? If you really need your workers to stay alive all the time, you can use it as an error finder as well as along with an eventual manual inspection of the code, or test writing. If you log the exceptions/errors thrown, it will be much more easier for you to determine the exact line of source code where the problem originated from. Besides, the job will never die. But again, my suggestion is to treat the cause, not the symptom.

## Real use case: RabbitMQ consumers

In the following part I will make a short presentation of a practical use case.

Consider a system where we use _queues_ to handle different background jobs from a PHP server-side application. For this we are going to use _AMQP_ (_Advanced Message Queueing Protocol_) and _RabbitMQ_ which is an AMPQ broker supporting (among other languages) PHP (see [https://github.com/php-amqplib/php-amqplib](https://github.com/php-amqplib/php-amqplib)).

Before starting let’s define some key concepts:

- _producer_ – application endpoint that _produces_ or sends messages;
- _consumer_ – application endpoint that _consumes_ or receives messages;
- _queue_ – the place where messages are stored;
- _message_ – the piece of information that is sent from producer to consumer.

As a quick overview the system will work like this: when a background job needs to be started, the producer will create a message containing the information about what job is intended to run and perhaps some other extra parameters. There will be a number of consumers running permanently and waiting for messages to process.

Here is the producer:

<?php

use PhpAmqpLib\Connection\AMQPConnection;

use PhpAmqpLib\Message\AMQPMessage;

class Producer

{

protected function produce($job, array $args = array())

{

$connection = new AMQPConnection('host', 'port', 'user', 'password', 'vhost');

$channel = $connection->channel();

$channel->queue_declare('queue', false, true, false, false);

$channel->exchange_declare('exchange', 'direct', false, true, false);

$channel->queue_bind('queue', 'exchange');

$messageBody = json_encode(array(

'job' => 'job_name',

'args' => array(

// extra parameters go here

)

));

$message = new AMQPMessage(

$messageBody,

array(

'content_type' => 'text/plain',

'delivery_mode' => 2

)

);

$channel->basic_publish($message, 'exchange');

$channel->close();

$connection->close();

}

}

And here the consumer:

<?php

use PhpAmqpLib\Channel\AMQPChannel;

use PhpAmqpLib\Connection\AMQPConnection;

use PhpAmqpLib\Message\AMQPMessage;

class Consumer

{

public function consume()

{

$connection = new AMQPConnection('host', 'port', 'user', 'password', 'vhost');

$channel = $connection->channel();

$channel->queue_declare('queue', false, true, false, false);

$channel->exchange_declare('exchange', 'direct', false, true, false);

$channel->queue_bind('queue', 'exchange');

$channel->basic_consume('queue', 'consumer_tag', false, false, false, false, 'process');

while (count($channel->callbacks) > 0) {

$channel->wait();

}

}

}

function process(AMQPMessage $message)

{

$message->delivery_info['channel']->basic_ack($message->delivery_info['delivery_tag']);

$messageBody = json_decode($message->body, true);

$job = $messageBody['job'];

$args = $messageBody['args'];

// Run job $job

}

There are a lot of function calls and blind arguments in the source code above, but it’s not the point to explain them here.

The consumers are started as PHP background processes since they have to be alive all the time in order to receive and then process messages. You can start how many instances you want using the following command:

$ php consumer.php &

We need to keep all the consumer instances alive, no matter what. So, let’s do the trick.

<?php

use PhpAmqpLib\Channel\AMQPChannel;

use PhpAmqpLib\Connection\AMQPConnection;

use PhpAmqpLib\Message\AMQPMessage;

class Consumer

{

public function consume()

{

$connection = new AMQPConnection('host', 'port', 'user', 'password', 'vhost');

$channel = $connection->channel();

$channel->queue_declare('queue', false, true, false, false);

$channel->exchange_declare('exchange', 'direct', false, true, false);

$channel->queue_bind('queue', 'exchange');

$channel->basic_consume('queue', 'consumer_tag', false, false, false, false, 'process');

register_shutdown_function('onShutdown', $channel, $connection);

while (count($channel->callbacks) > 0) {

$channel->wait();

}

}

}

function process(AMQPMessage $message)

{

$message->delivery_info['channel']->basic_ack($message->delivery_info['delivery_tag']);

$messageBody = json_decode($message->body, true);

$job = $messageBody['job'];

$args = $messageBody['args'];

// Run job $job

}

function onShutdown(AMQPChannel $channel, AMQPConnection $connection)

{

global $argv;

$_ = $_SERVER['_'];

$channel->close();

$connection->close();

pcntl_exec($_, $argv);

}

Whenever a consumer dies it is automatically restarted. Simple as that.

There’s another thing I want to do. Let’s say you want to keep the consumers “fresh” for various reasons, so we need to let the consumers live only for a specified amount of time. There will be an _expiration time_ for every one of them, and after that they should be restarted. In order to do that the following changes need to be done:

<?php

use PhpAmqpLib\Channel\AMQPChannel;

use PhpAmqpLib\Connection\AMQPConnection;

use PhpAmqpLib\Message\AMQPMessage;

class Consumer

{

public $startTime;

public function consume()

{

$this->startTime = new DateTime();

$connection = new AMQPConnection('host', 'port', 'user', 'password', 'vhost');

$channel = $connection->channel();

$channel->queue_declare('queue', false, true, false, false);

$channel->exchange_declare('exchange', 'direct', false, true, false);

$channel->queue_bind('queue', 'exchange');

$channel->basic_consume('queue', 'consumer_tag', false, false, false, false, 'process');

register_shutdown_function('onShutdown', $channel, $connection);

while (count($channel->callbacks) > 0) {

$channel->wait();

}

}

}

function process(AMQPMessage $message)

{

$now = new DateTime();

$minutesSinceStart = ($now->getTimestamp() - $this->startTime->getTimestamp()) / 60;

if ($minutesSinceStart >= 60) {

die;

}

$message->delivery_info['channel']->basic_ack($message->delivery_info['delivery_tag']);

$messageBody = json_decode($message->body, true);

$job = $messageBody['job'];

$args = $messageBody['args'];

// Run job $job

}

function onShutdown(AMQPChannel $channel, AMQPConnection $connection)

{

global $argv;

$_ = $_SERVER['_'];

$channel->close();

$connection->close();

pcntl_exec($_, $argv);

}

There will be a simple die call if there is more then 60 minutes since the consumer process started.

First, the die call is made before the message is acknowledged. It means that it has not been already removed from the queue, so it will be reassigned to another consumer.

Second, a die call is enough since we already added the piece of code which listens to the process shutdown event and then does a restart.

And that’s a quick overview of how you can practically use the “restart on die” trick.

Let’s recap. We have discovered several ways we can leverage this trick:

- prevent background PHP jobs from dying;
- detect uncaught exceptions or broken code;
- keep the processes “fresh” by restarting them after a specified amount of time.

But overall, please keep this in mind: _treat the cause, not the symptom_.

## We transform challenges into digital experiences

Get in touch to let us know what you’re looking for. Our policy includes 14 days risk-free!

[Free project consultation](https://www.algotech.solutions/contact/)
