---
title: "Process-driven REST API design"
notion_id: 41502e95-a0fa-40b9-a77a-529b6aac71ee
notion_url: https://app.notion.com/p/Process-driven-REST-API-design-41502e95a0fa40b9a77a529b6aac71ee
last_edited: 2022-12-21T17:16:00.000Z
source_url: https://hackernoon.com/process-driven-rest-api-design-75ca88917582
tags: ["REST API", "System Design / Software Architecture", "Article", "Hackernoon", "English"]
---
REST is a heavily oversimplified and massively misunderstood (but very widely used) concept for designing API’s. Just look for any forum thread that features Roy Fielding (the “founder” of REST) and you will probably find him telling the author and other participants how they misunderstood what he intended REST to be. Which is perhaps no surprise, as his original paper is too long to read for many, and the ideas are not completely trivial.

So instead of REST, we usually get something almost, but not quite, entirely unlike REST.

REST is often simplified to CRUD because contrary to REST, CRUD is simple and well-understood. And perhaps because they are both four letter abbreviations that contain an R, who knows.

With CRUD, your operations are Create, Read, Update and Delete. These can be conveniently mapped to the four main HTTP verbs: POST, GET, PUT and DELETE. Now you just need something to perform these operations on, so let’s take your data entities from your database, convert them into URLs, call them “resources”, and presto, API! And it’s RESTful because there are HTTP verbs and resources, so we can move the Trello card to Done and go for lunch.

I don’t think this is proper API design, and definitely not RESTful. You didn’t create a REST API; you just created an ORM with only four operations and the latency of HTTP. There is no Representational State Transfer (REST) going on, not really. It’s merely a simple way of leveraging HTTP to access your data, so simple even that there are frameworks that can automatically generate such an API for you from your database entities.

I call this approach data-driven API design. What I want to do however is to describe the method of designing REST API’s that I use, which is supported by quite some literature and good experiences. I feel like this is a lot more like how REST was intended to be, but because of the confusion regarding the terminology, I tend to call this process-driven REST API design.

Before we discuss this method, however, let’s first consider the downsides of data-driven API design.

### The pitfalls of data-driven API design

Consider a simple webshop. There are products and customers who can place orders for these products. When an order is placed, it is created as an order with one or more order items, with each order item representing a product. See below for the database scheme (my UML is rusty so bear with me):

![image](data:image/svg+xml,%3csvg%20xmlns=%27http://www.w3.org/2000/svg%27%20version=%271.1%27%20width=%27637%27%20height=%27122%27/%3e)

![image](https://hackernoon.com/_next/image?url=https%3A%2F%2Fcdn.hackernoon.com%2Fhn-images%2F1*puvVWbskJR4lBCR9KHV4Vw.png&w=1920&q=75)

Database scheme for a very simple webshop. I know, the UML hurts my eyes too.

Through the data-driven approach, you will get the following resources for retrieving product information:

GET /products/{uuid}Response:

- uuid (string)
- name (string)
- price (float)

GET /productsResponse:

- products (Product entities)

For placing the order, though a purely data-driven approach, you would get the following resources:

POST /order-itemsRequest body:

- 

product uuid (string)Response:

- order-item

POST /ordersRequest body:

- 

order-item uuid (string)Response:

- order

But this cannot work with our database scheme: Order items cannot exist in without being part of an order. So let’s modify it to be as follows:

POST /pre-order-itemsRequest body:

- product uuid (string)

POST /ordersRequest body:

- pre-order-item uuid (string)

And introduce a new database entity:

![image](data:image/svg+xml,%3csvg%20xmlns=%27http://www.w3.org/2000/svg%27%20version=%271.1%27%20width=%27158%27%20height=%2774%27/%3e)

![image](https://hackernoon.com/_next/image?url=https%3A%2F%2Fcdn.hackernoon.com%2Fhn-images%2F1*mtODVFy3EqGuH_rtqXGx2Q.png&w=384&q=75)

Let’s now build a client application for this API. Through the /products resource, you retrieve the products. You allow the customer to put products into a shopping basket (which is persisted on the client application). To perform the checkout through the API, the client application will now perform a POST on /pre-order-items for every product, and when that is done, it will perform a POST on /orders providing all pre-order-item UUID’s that were returned when creating the pre-order-items.

Yeah, don’t do this.

For placing a single order, you are now performing 1 + n HTTP requests, with n being the number of products. If a customer orders 10 products, it will take around 10 seconds for the order to be placed. Furthermore, what do you do if creating one of the order items fails? How do you handle a rollback? And what happens with pre-order-items that are never and will never be linked to any order?

Perhaps this approach might seem silly to you. That’s a good thing. I’ve seen a design such as this more than once in practice, with similar properties regarding performance and data consistency. It is the result of an extraordinary focus on CRUD and database tables, so if you were nodding along with this approach, it’s time for a pretty radical shift in thinking.

Anyway, back to designing the API.

Now a new requirement comes in. There should be shipping costs included, but only for orders under €40,-. No worries, you can just inspect the shopping basket, look at the total, and if it is less than €40,-, you display the shipping costs. If it’s more, you don’t. Right?

No! Stop! Please don’t do this. You’ve just encoded business logic in your API client. This might work if you have only one client, but why are you painstakingly designing an API if you are only going to have a single client? What if you have 30 clients, do you expect them all to implement this logic? And what if it changes, and the cutoff becomes €50,-? Or worse, what if the shipping costs will become dependent on the content of the shopping basket?

Let’s work through the data-driven API solution, and immediately suggest an improvement to the pre-order-items approach of above.

Instead of creating pre-order-items in isolation, they could be grouped by implementing a shopping basket entity on the server that the clients can add items to (with PUT and such, REST ahoy!). The basket can then be used to place the order.

PUT /basketRequest body:

- 

product uuids (list of strings)Response:

- product uuids (list of strings)

GET /basket/{uuid}Response:

- product uuids (list of strings)

POST /ordersRequest body:

- 

basket uuid (string)Response:

- order

Tada, now your API has suddenly become responsible for storing client-specific data. And as this data is transient in nature, it has also become responsible for cleaning it up, or it has to provide mechanisms for the client to do so. And worse, the client has no way to implement another process that doesn’t include a shopping basket, for example, a simple “order again directly” button, without having to use the shopping basket entity.

As I hope I’ve illustrated, a data-driven approach leads to performance issues (because it’s always going to be at best a very inefficient ORM), to business logic in your clients, and to inflexible processes in your API. Furthermore, solving these problems in a data-driven way merely results in other problems and a more complicated API.

Let’s try something else.

### Process-driven REST API design

![image](data:image/svg+xml,%3csvg%20xmlns=%27http://www.w3.org/2000/svg%27%20version=%271.1%27%20width=%27650%27%20height=%27469%27/%3e)

![image](https://hackernoon.com/_next/image?url=https%3A%2F%2Fcdn.hackernoon.com%2Fhn-images%2F1*saNB2GwsO9CnmaAnfq6-wA.jpeg&w=1920&q=75)

Don’t enter the data-driven tunnel. Enjoy the process-driven freedom instead.

Process-driven design advocates starting with the process, instead of with the data.

You might have seem this coming.

In fact, we basically ignore our data entities for designing our resources. Instead, we first design the processes that we want our API to support, and then we base our resources on the _steps in the process_.

The method of designing the process doesn’t matter much; you can use BPMN, a flowchart, use case descriptions or whatever you like. The goal is to find the states your API can be in. With this in hand, we can consider how best to represent the states themselves as well as the transfer between the states. REST: Representational State Transfer, remember?

Let’s again consider the webshop from before, but let’s now first design and draw the processes we’ve discussed. See below for a version using BPMN (it’s simplified quite a bit, but it should serve our example).

![image](data:image/svg+xml,%3csvg%20xmlns=%27http://www.w3.org/2000/svg%27%20version=%271.1%27%20width=%27800%27%20height=%27361.20805369127515%27/%3e)

![image](https://hackernoon.com/_next/image?url=https%3A%2F%2Fcdn.hackernoon.com%2Fhn-images%2F1*ifuTg0lUqogm8fnkd46d8Q.png&w=1920&q=75)

Simplified process descriptions in BPMN of our simple webshop.

For every step in the process, we consider 1) whether the step needs access to data on the server or not, and 2) whether the step contains business logic that is relevant for all client applications. If either of those criteria is true, we will design an API resource for the process step. If not, the process step is simply a client operation that the client will be responsible for.

Let’s now do this for the process steps in the processes drawn above.

_Search products_We look for zero or more products based on some given criteria. This requires data on the server, so a resource is warranted.

GET /products?{criteria}Response:

- products (Product entities)

_Add to basket_This is a client operation as no data on the server is necessary. No resource is required.

_View basket and costs_Now, this is an interesting one. The purpose of this step is to, given a set of products (a “basket”), provide an overview of what you have to pay. This is clearly business logic and as such it is the responsibility of the API to provide a resource to facilitate this process step.

We could introduce a “bill” resource for this purpose:

GET /bill?{product-list}Response:

- product total (float)
- shipping costs (float)
- total (float)

_Side note: If the product list is complex or long, POST might be better suited here so that a request body can be used._

This resource represents the business logic of calculating and presenting the right totals and costs to the user. It has no database representation; in programming terms, it would be a function.

This is a resource that wouldn’t appear in a data-driven approach at all, but very organically appears when considering the process: of _course_ you want to see an overview of the costs! That’s an integral part of every webshop, so of course our API should support it.

_Place order_Finally, by providing one or more products, an order can be placed. We can model this as follows:

POST /ordersRequest body:

- 

product uuids (list of strings)Response:

- order

And that’s it! This gives us a set of three resources which directly model the processes we want to support.

Let’s evaluate the result and contrast it with the result of the data-driven approach.

The first observation is the complete absence of order items or their equivalent, while they seemed so important in the data-driven approach. And that makes sense; nowhere in the process are they relevant for the customer or the client application, so why would they enter the domain of our API? This, of course, makes the API significantly more straightforward.

Another thing of note is that we didn’t introduce a resource to access a single product (/products/{uuid}). While by itself that might not seem very earth-shattering, it is important that we made the _conscious decision_ to do so. Designing an elegant and functional REST API requires thought and deliberation. The process-driven method stimulates this by making you think about every step, and helps you avoid unnecessary complexities.

The final interesting observation is that adding the requirement for the shipping costs as discussed earlier _does not change_ the designed process; the “view basket and costs” step still exists, it will merely be expanded with more data. And this is actually one of the greatest properties of this approach. I’ve found that, because the processes your API has to support very rarely change fundamentally, API’s designed with this method are very good at adapting to changing requirements. They can easily grow to support new situations, while still allowing for plenty of flexibility, and without introducing unnecessary resources or being tied to your application internals.

### Wrapping up

Process-driven REST API design leads to more flexible, simpler and clearer API’s which are more close to being truly RESTful API’s. I have applied this method in multiple projects and while it takes more effort up front, it pays off significantly in the long run.

And of course, this is just a part of REST API design. In this article, I’ve largely glossed over the contents of (POST and PUT) requests, choosing the right HTTP verbs, the use of hypermedia and standardized media types, and plenty of other things that make up a good API. I’ll write more on those topics in the future.

I did however give a presentation on the above topics at PyGrunn. My slides for that can be found [here](http://larsderidder.github.io/rest-design-talk/?ref=hackernoon.com).
