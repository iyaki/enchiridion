---
title: "Query Objects instead of Repositories"
notion_id: 681afd43-0007-48c0-a834-9d1bf5fb2bf4
notion_url: https://app.notion.com/p/Query-Objects-instead-of-Repositories-681afd43000748c0a8349d1bf5fb2bf4
last_edited: 2023-10-23T11:24:00.000Z
source_url: https://codeopinion.com/query-objects-instead-of-repositories/
tags: ["Object Oriented Programming", "Domain Driven Design", "System Design / Software Architecture", "Article", "CodeOpinion", "English"]
---
> 

_**Sponsor: **__Do you build complex software systems? See how NServiceBus makes it easier to design, build, and manage software systems that use message queues to achieve loose coupling. _[_Get started for free_](https://go.particular.net/codeopinion)_._

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The [repository pattern](http://martinfowler.com/eaaCatalog/repository.html) is often used to encapsulate simple and sometimes rather complex query logic. However, it has also been morphed into handling persistence and is often used as another layer of abstraction from your data mapping layer. This blog post show you how to slim down and simplify your repositories or possibly eliminate them all together by using query objects. A typical repository will look something like this:

```plain text
public interface IProductRepository
{
	void Insert(Product product);
	void Delete(Product product);
	IEnumerable<Product> GetById(Guid id);
	IEnumerable<Product> GetAllActive();
	IEnumerable<Product> FindByName(string name);
	IEnumerable<Product> FindBySku(string name);
	IEnumerable<Product> Find(string keyword, int limit, int page);
	IEnumerable<Product> GetRelated(Guid id);
}

```

Each of the Get/Find methods implemented above would encapsulate a specific a query. This type of repository would most likely be used to then transform the data returned from a set of methods into a View Model which would then be passed to our view or serialized and sent back to the caller (browser). As an example of passing the model to a view in ASP.NET MVC, it would look something like this:

```plain text
public ViewResult ProductDetails(Guid productId)
{
	var product = _productRepository.GetById(productId);
	var relatedProducts = _productRepository.GetRelated(productId);

	var model = new ProductDetailsModel
	{
		Id = product.Id,
		Name = product.Name,
		Price = product.Price,
		PriceFormatted = product.Price.ToString("C"),
		RecommendedProducts = (from x in relatedProducts select new ProductDetailModel.RecommendedProducts {
				ProductId = x.RecommendedProductId,
				Name = x.Name,
				Price = x.Price,
				PriceFormatted = x.Price.ToString("C")
			})
	};

	return View(model);
}

```

As I’ve mentioned [before](https://codeopinion.com/simplify-your-code/), I believe you should think of your MVC framework as an HTTP interface to your application. Regardless if you are returning HTML or JSON, the generation of your View Model should not be coupled to your MVC framework.

### Query Objects

In the example above, I want to extract the generation of my view model info a query object. A query object is similar to a command object for processing behavior in our domain. Our query object will now look like this:

```plain text
public class ProductDetailsQuery
{
	public Guid ProductID { get; private set; }

	public ProductDetailQuery(Guid productId)
	{
		ProductId = productId
	}
}
```

In order to execute our query, we will implement in a query handler.

```plain text
public class ProductDetailQueryHandler
{
	private DbContext _db;

	public ProductDetailQueryHandler(DbContext db)
	{
		_db = db;
	}

	public ProductDetailModel Handle(ProductDetailQuery query)
	{
		var product = (from p in _db.Products where p.ProductId == query.ProductId).SingleOrDefault();
		if (product == null) {
			throw new InvalidOperationException("Product does not exist.");
		}

		var relatedProducts = (from p in _db.RecommendedProducts where p.PrimaryProductId == query.ProductId);

		return new ProductDetailsModel
		{
			Id = product.Id,
			Name = product.Name,
			Price = product.Price,
			PriceFormatted = product.Price.ToString("C"),
			RecommendedProducts = (from x in relatedProducts select new ProductDetailModel.RecommendedProducts {
				ProductId = x.RecommendedProductId,
				Name = x.Name,
				Price = x.Price,
				PriceFormatted = x.Price.ToString("C")
			})
		};
	}
}

```

Now we have encapsulated the generation of our view model which can be used in our controller.

```plain text
public ViewResult ProductDetails(ProductDetailQuery query)
{
	var model = _queryHandler(query);
	return View(model);
}

```

Now our controller is responsible for delegating the call to generate the model and action result or serialization. In my next post I will take this a step further by introducing a common interface for our query handler in order to accept multiple query objects and return types.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->
