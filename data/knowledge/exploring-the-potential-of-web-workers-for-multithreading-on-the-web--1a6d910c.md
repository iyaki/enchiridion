---
title: "Exploring The Potential Of Web Workers For Multithreading On The Web"
notion_id: 1a6d910c-5986-402b-aae7-863663140bfb
notion_url: https://app.notion.com/p/Exploring-The-Potential-Of-Web-Workers-For-Multithreading-On-The-Web-1a6d910c5986402baae7863663140bfb
last_edited: 2023-04-26T00:23:00.000Z
source_url: https://www.smashingmagazine.com/2023/04/potential-web-workers-multithreading-web/
tags: ["Article", "Guide", "Smashing Magazine", "English", "Frontend", "Javascript"]
---
Web Workers are a powerful feature of modern web development and were introduced as part of the [HTML5 specification in 2009](https://en.wikipedia.org/wiki/Web_worker). They were designed to provide a way to execute JavaScript code in the background, separate from the main execution thread of a web page, in order to improve performance and responsiveness.

The main thread is the single execution context that is responsible for rendering the UI, executing JavaScript code, and handling user interactions. In other words, [JavaScript is “single-threaded”](https://developer.mozilla.org/en-US/docs/Glossary/Main_thread). This means that any time-consuming task, such as complex calculations or data processing that is executed, would block the main thread and cause the UI to freeze and become unresponsive.

This is where Web Workers come in.

Web Workers were implemented as a way to address this problem by allowing time-consuming tasks to be executed in a separate thread, called a **worker thread**. This enabled JavaScript code to be executed in the background without blocking the main thread and causing the page to become unresponsive.

Creating a web worker in JavaScript is not much of a complicated task. The following steps provide a starting point for integrating a web worker into your application:

1. Create a new JavaScript file that contains the code you want to run in the worker thread. This file should not contain any references to the DOM, as it will not have access to it.
2. In your main JavaScript file, create a new worker object using the `Worker` constructor. This constructor takes a single argument, which is the URL of the JavaScript file you created in step 1.

```javascript
const worker = new Worker('worker.js');

```

1. Add event listeners to the worker object to handle messages sent between the main thread and the worker thread. The `onmessage` event handler is used to handle messages sent from the worker thread, while the `postMessage` method is used to send messages to the worker thread.

```javascript
worker.onmessage = function(event) {
  console.log('Worker said: ' + event.data);
};
worker.postMessage('Hello, worker!');

```

1. In your worker JavaScript file, add an event listener to handle messages sent from the main thread using the `onmessage` property of the `self` object. You can access the data sent with the message using the `event.data` property.

```javascript
self.onmessage = function(event) {
  console.log('Main thread said: ' + event.data);
  self.postMessage('Hello, main thread!');
};

```

Now let’s run the web application and test the worker. We should see messages printed to the console indicating that messages were sent and received between the main thread and the worker thread.

([Large preview](https://files.smashing.media/articles/potential-web-workers-multithreading-web/1-messages-console-between-main-worker-threads.png))

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

One key difference between Web Workers and the main thread is that Web Workers have **no access to the DOM or the UI**. This means that they cannot directly manipulate the HTML elements on the page or interact with the user.

> Web Workers are designed to perform tasks that do not require direct access to the UI, such as data processing, image manipulation, or calculations.

Another important difference is that Web Workers are designed to run in a **sandboxed environment**, separate from the main thread, which means that they have limited access to system resources and cannot access certain APIs, such as the [`localStorage`](https://www.smashingmagazine.com/2010/10/local-storage-and-how-to-use-it/) or [`sessionStorage`](https://developer.mozilla.org/en-US/docs/Web/API/Window/sessionStorage) APIs. However, they can communicate with the main thread through a messaging system, allowing data to be exchanged between the two threads.

## Importance And Benefits Of Web Workers For Multithreading On The Web

Web Workers provide a way for web developers to achieve multithreading on the web, which is crucial for building high-performance web applications. By enabling time-consuming tasks to be executed in the background, separate from the main thread, Web Workers improve the overall responsiveness of web pages and allow for a more seamless user experience. The following are some of the importance and benefits of Web Workers for multithreading on the Web.

### Improved Resource Utilization

By allowing time-consuming tasks to be executed in the background, Web Workers make more efficient use of system resources, enabling faster and more efficient processing of data and improving overall performance. This is especially important for web applications that involve large amounts of data processing or image manipulation, as Web Workers can perform these tasks without impacting the user interface.

### Increased Stability And Reliability

By isolating time-consuming tasks in separate worker threads, Web Workers help to prevent crashes and errors that can occur when executing large amounts of code on the main thread. This makes it easier for developers to write stable and reliable web applications, reducing the likelihood of user frustration or loss of data.

### Enhanced Security

Web Workers run in a sandboxed environment that is separate from the main thread, which helps to enhance the security of web applications. This isolation prevents malicious code from accessing or modifying data in the main thread or other Web Workers, reducing the risk of data breaches or other security vulnerabilities.

### Better Resource Utilization

Web Workers can help to improve resource utilization by freeing up the main thread to handle user input and other tasks while the Web Workers handle time-consuming computations in the background. This can help to improve overall system performance and reduce the likelihood of crashes or errors. Additionally, by leveraging multiple CPU cores, Web Workers can make more efficient use of system resources, enabling faster and more efficient processing of data.

Web Workers also enable better [load balancing](https://www.ibm.com/topics/load-balancing) and scaling of web applications. By allowing tasks to be executed in parallel across multiple worker threads, Web Workers can help **distribute the workload evenly across multiple cores or processors**, enabling faster and more efficient processing of data. This is particularly important for web applications that experience high traffic or demand, as Web Workers can help to ensure that the application can handle an increased load without impacting performance.

## Practical Applications Of Web Workers

Let us explore some of the most common and useful applications of Web Workers. Whether you’re building a complex web application or a simple website, understanding how to leverage Web Workers can help you improve performance and provide a better user experience.

### Offloading CPU-Intensive Work

Suppose we have a web application that needs to perform a large, CPU-intensive computation. If we perform this computation in the main thread, the user interface will become unresponsive, and the user experience will suffer. To avoid this, we can use a Web Worker to perform the computation in the background.

```javascript
// Create a new Web Worker.
const worker = new Worker('worker.js');

// Define a function to handle messages from the worker.
worker.onmessage = function(event) {
  const result = event.data;
  console.log(result);
};

// Send a message to the worker to start the computation.
worker.postMessage({ num: 1000000 });

// In worker.js:

// Define a function to perform the computation.
function compute(num) {
  let sum = 0;
  for (let i = 0; i < num; i++) {
    sum += i;
  }
  return sum;
}

// Define a function to handle messages from the main thread.
onmessage = function(event) {
  const num = event.data.num;
  const result = compute(num);
  postMessage(result);
};

```

In this example, we create a new Web Worker and define a function to handle messages from the worker. We then send a message to the worker with a parameter (`num`) that specifies the number of iterations to perform in the computation. The worker receives this message and performs the computation in the background. When the computation is complete, the worker sends a message back to the main thread with the result. The main thread receives this message and logs the result to the console.

This task involves adding up all the numbers from `0` to a given number. While this task is relatively simple and straightforward for small numbers, it can become computationally intensive for very large numbers.

In the example code we used above, we passed the number `1000000` to the `compute()` function in the Web Worker. This means that the compute function will need to add up all the numbers from 0 to one million. This involves a large number of additional operations and can take a significant amount of time to complete, especially if the code is running on a slower computer or in a browser tab that is already busy with other tasks.

By offloading this task to a Web Worker, the main thread of the application can continue to run smoothly without being blocked by the computationally intensive task. This allows the user interface to remain responsive and ensures that other tasks, such as user input or animations, can be handled without delay.

### Handling Network Requests

Let us consider a scenario where a web application needs to initiate a significant number of network requests. Performing these requests within the main thread could cause the user interface to become unresponsive and result in a poor user experience. In order to prevent this issue, we can utilize Web Workers to handle these requests in the background. By doing so, the main thread remains free to execute other tasks while the Web Worker handles the network requests simultaneously, resulting in improved performance and a better user experience.

```javascript
// Create a new Web Worker.
const worker = new Worker('worker.js');

// Define a function to handle messages from the worker.
worker.onmessage = function(event) {
  const response = event.data;
  console.log(response);
};

// Send a message to the worker to start the requests.
worker.postMessage({ urls: ['https://api.example.com/foo', 'https://api.example.com/bar'] });

// In worker.js:

// Define a function to handle network requests.
function request(url) {
  return fetch(url).then(response => response.json());
}

// Define a function to handle messages from the main thread.
onmessage = async function(event) {
  const urls = event.data.urls;
  const results = await Promise.all(urls.map(request));
  postMessage(results);
};

```

In this example, we create a new Web Worker and define a function to handle messages from the worker. We then send a message to the worker with an array of URLs to request. The worker receives this message and performs the requests in the background using the [`fetch`](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API)[ API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API). When all requests are complete, the worker sends a message back to the main thread with the results. The main thread receives this message and logs the results to the console.

### Parallel Processing

Suppose we have a web application that needs to perform a large number of independent computations. If we perform these computations in sequence in the main thread, the user interface will become unresponsive, and the user experience will suffer. To avoid this, we can use a Web Worker to perform the computations in parallel.

```javascript
// Create a new Web Worker.
const worker = new Worker('worker.js');

// Define a function to handle messages from the worker.
worker.onmessage = function(event) {
  const result = event.data;
  console.log(result);
};

// Send a message to the worker to start the computations.
worker.postMessage({ nums: [1000000, 2000000, 3000000] });

// In worker.js:

// Define a function to perform a single computation.
function compute(num) {
  let sum = 0;
  for (let i = 0; i < num; i++) {
    sum += i;
}
  return sum;
}

// Define a function to handle messages from the main thread.
onmessage = function(event) {
  const nums = event.data.nums;
  const results = nums.map(compute);
  postMessage(results);
};

```

In this example, we create a new Web Worker and define a function to handle messages from the worker. We then send a message to the worker with an array of numbers to compute. The worker receives this message and performs the computations in parallel using the map method. When all computations are complete, the worker sends a message back to the main thread with the results. The main thread receives this message and logs the results to the console.

## Limitations And Considerations

Web workers are a powerful tool for improving the performance and responsiveness of web applications, but they also have some limitations and considerations that you should keep in mind when using them. Here are some of the most important ones:

### Browser Support

Web workers are supported in all major browsers, including Chrome, Firefox, Safari, and Edge. However, there are still some other browsers that do not support web workers or may have limited support.

For a more extensive look at browser support, see [Can I Use](https://caniuse.com/webworkers).

It is important that you check out the browser support for any feature before using them in production code and test your application thoroughly to ensure compatibility.

### Limited Access To The DOM

Web workers run in a separate thread and do not have access to the DOM or other global objects in the main thread. This means you **cannot directly manipulate the DOM from a web worker or access global objects** like windows or documents.

To work around this limitation, you can use the `postMessage` method to communicate with the main thread and update the DOM or access global objects indirectly. For example, you can send data to the main thread using `postMessage` and then update the DOM or global objects in response to the message.

Alternatively, there are some libraries that help solve this issue. For example, the [WorkerDOM](https://github.com/ampproject/worker-dom) library enables you to run the DOM in a web worker, allowing for faster page rendering and improved performance.

### Communication Overhead

Web workers communicate with the main thread using the [`postMessage`](https://developer.mozilla.org/en-US/docs/Web/API/Worker/postMessage) method, and as a result, could introduce communication overhead, which refers to the amount of time and resources required to establish and maintain communication between two or more computing systems, such as between a Web Worker and the main thread in a web application. This could result in a delay in processing messages and potentially slow down the application. To minimize this overhead, you should **only send essential data** between threads and **avoid sending large amounts of data or frequent messages**.

### Limited Debugging Tools

Debugging Web Workers can be more challenging than debugging code in the main thread, as there are fewer debugging tools available. To make debugging easier, you can use the `console` API to log messages from the worker thread and use [browser developer tools](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Tools_and_setup/What_are_browser_developer_tools) to inspect messages sent between threads.

### Code Complexity

Using Web Workers can increase the complexity of your code, as you need to manage communication between threads and ensure that data is passed correctly. This can make it more difficult to write, debug, and maintain your code, so you should carefully consider whether using web workers is necessary for your application.

## Strategies For Mitigating Potential Issues With Web Workers

Web Workers are a powerful tool for improving the performance and responsiveness of web applications. However, when using Web Workers, there are several potential issues that can arise. Here are some strategies for mitigating these issues:

### Minimize Communication Overhead With Message Batching

Message batching involves grouping multiple messages into a single batch message, which can be more efficient than sending individual messages separately. This approach reduces the number of round-trips between the main thread and Web Workers. It can help to minimize communication overhead and improve the overall performance of your web application.

To implement message batching, you can **use a queue to accumulate messages and send them together as a batch** when the queue reaches a certain threshold or after a set period of time. Here’s an example of how you can implement message batching in your Web Worker:

```javascript
// Create a message queue to accumulate messages.
const messageQueue = [];

// Create a function to add messages to the queue.
function addToQueue(message) {
  messageQueue.push(message);

  // Check if the queue has reached the threshold size.
  if (messageQueue.length >= 10) {
    // If so, send the batched messages to the main thread.
    postMessage(messageQueue);

    // Clear the message queue.
    messageQueue.length = 0;
  }
}

// Add a message to the queue.
addToQueue({type: 'log', message: 'Hello, world!'});

// Add another message to the queue.
addToQueue({type: 'error', message: 'An error occurred.'});

```

In this example, we create a message queue to accumulate messages that need to be sent to the main thread. Whenever a message is added to the queue using the `addToQueue` function, we check if the queue has reached the threshold size (in this case, ten messages). If so, we send the batched messages to the main thread using the `postMessage` method. Finally, we clear the message queue to prepare it for the next batch.

By batching messages in this way, we can reduce the overall number of messages sent between the main thread and Web Workers,

### Avoid Synchronous Methods

These are JavaScript functions or operations that block the execution of other code until they are complete. Synchronous methods can block the main thread and cause your application to become unresponsive. To avoid this, you should avoid using synchronous methods in your Web Worker code. Instead, use **asynchronous methods** such as `setTimeout()` or `etInterval()` to perform long-running computations.

Here is a little demonstration:

```javascript
// In the worker
self.addEventListener('message', (event) => {
  if (event.data.action === 'start') {
    // Use a setTimeout to perform some computation asynchronously.
    setTimeout(() => {
      const result = doSomeComputation(event.data.data);

      // Send the result back to the main thread.
      self.postMessage({ action: 'result', data: result });
    }, 0);
  }
});

```

### Be Mindful Of Memory Usage

Web Workers have their own memory space, which can be limited depending on the user’s device and browser settings. To avoid memory issues, you should be mindful of the amount of memory your Web Worker code is using and avoid creating large objects unnecessarily. For example:

In this code, the Web Worker processes an array of data and returns the result to the main thread using the `postMessage` method. However, the `for` loop used to process the data may be time-consuming.

The reason for this is that the code is processing an entire array of data at once, meaning that all the data must be loaded into memory at the same time. If the data set is very large, this can cause the Web Worker to consume a significant amount of memory, potentially exceeding the memory limit allocated to the Web Worker by the browser.

To mitigate this issue, you can consider using built-in JavaScript methods like [`forEach`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/forEach) or [`reduce`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reduce), which can process data one item at a time and avoid the need to load the entire array into memory at once.

### Browser Compatibility

Web Workers are supported in most modern browsers, but some older browsers may not support them. To ensure compatibility with a wide range of browsers, you should test your Web Worker code in different browsers and versions. You can also use **feature detection** to check if Web Workers are supported before using them in your code, like this:

```javascript
if (typeof Worker !== 'undefined') {
  // Web Workers are supported.
  const worker = new Worker('worker.js');
} else {
  // Web Workers are not supported.
  console.log('Web Workers are not supported in this browser.');
}

```

This code checks if Web Workers are supported in the current browser and creates a new Web Worker if they are supported. If Web Workers are not supported, the code logs a message to the console indicating that Web Workers are not supported in the browser.

By following these strategies, you can ensure that your Web Worker code is efficient, responsive, and compatible with a wide range of browsers.

## Conclusion

As web applications become increasingly complex and demanding, the importance of efficient multithreading techniques — such as Web Workers — is likely to increase. Web Workers are an essential feature of modern web development that allows developers to offload CPU-intensive tasks to separate threads, improving application performance and responsiveness. However, there are significant limitations and considerations to keep in mind when working with Web Workers, such as the lack of access to the DOM and limitations on the types of data that can be passed between threads.

To mitigate these potential issues, developers can follow strategies as mentioned earlier, such as using asynchronous methods and being mindful of the complexity of the task being offloaded.

Multithreading with Web Workers is likely to remain an important technique for improving web application performance and responsiveness in the future. While there are other techniques for achieving multithreading in JavaScript, such as using [WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket) or [`SharedArrayBuffer`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/SharedArrayBuffer), Web Workers have several advantages that make them a powerful tool for developers.

Adopting more recent technology such as [WebAssembly](https://developer.mozilla.org/en-US/docs/WebAssembly) may open up new opportunities for using Web Workers to offload even more complex and computationally-intensive tasks. Overall, Web Workers are likely to continue to evolve and improve in the coming years, helping developers create more efficient and responsive web applications.

Additionally, many libraries and tools exist to help developers work with Web Workers. For example, [Comlink](https://github.com/GoogleChromeLabs/comlink) and [Workerize](https://github.com/developit/workerize) provide a simplified API for communicating with Web Workers. These libraries abstract away some of the complexity of managing Web Workers, making it easier to leverage their benefits.

Hopefully, this article has given you a good understanding of the potential of web workers for multithreading and how to use them in your own code.
