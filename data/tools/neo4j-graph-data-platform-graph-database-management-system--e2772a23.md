---
title: "Neo4j Graph Data Platform - Graph Database Management System"
notion_id: e2772a23-65c8-49a1-9bb3-c70aa11a276d
notion_url: https://app.notion.com/p/Neo4j-Graph-Data-Platform-Graph-Database-Management-System-e2772a2365c849a19bb3c70aa11a276d
last_edited: 2023-01-18T14:01:00.000Z
source_url: https://neo4j.com/
tags: ["Tool", "Service", "English", "Databases", "Untried"]
---
# Blazing-Fast Graph, Petabyte Scale

With proven [trillion+ entity performance](https://neo4j.com/product/neo4j-graph-database/scalability/), developers, data scientists, and enterprises rely on Neo4j as the top choice for high-performance, scalable analytics, intelligent app development, and advanced AI/ML pipelines.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

["[Neo4j] is ](https://neo4j.com/case-studies/the-international-consortium-of-investigative-journalists-icij/)[**a revolutionary discovery tool**](https://neo4j.com/case-studies/the-international-consortium-of-investigative-journalists-icij/)[ that's transformed our investigative journalism process."](https://neo4j.com/case-studies/the-international-consortium-of-investigative-journalists-icij/)

["Neo4j allows us to see relationships...that simply ](https://neo4j.com/case-studies/boston-scientific/)[**would not be possible**](https://neo4j.com/case-studies/boston-scientific/)[ with other tools."](https://neo4j.com/case-studies/boston-scientific/)

[Neo4j helped NASA ](https://neo4j.com/users/nasa/)[**save millions of dollars and two years**](https://neo4j.com/users/nasa/)[ in their work on the Mission to Mars.](https://neo4j.com/users/nasa/)

["[Neo4j] is ](https://neo4j.com/case-studies/the-international-consortium-of-investigative-journalists-icij/)[**a revolutionary discovery tool**](https://neo4j.com/case-studies/the-international-consortium-of-investigative-journalists-icij/)[ that's transformed our investigative journalism process."](https://neo4j.com/case-studies/the-international-consortium-of-investigative-journalists-icij/)

["Neo4j allows us to see relationships...that simply ](https://neo4j.com/case-studies/boston-scientific/)[**would not be possible**](https://neo4j.com/case-studies/boston-scientific/)[ with other tools."](https://neo4j.com/case-studies/boston-scientific/)

[Neo4j is “](https://neo4j.com/case-studies/ebay/)[**thousands of times faster**](https://neo4j.com/case-studies/ebay/)[ than our prior MySQL solution, with queries that require ](https://neo4j.com/case-studies/ebay/)[**10-100 times less code**](https://neo4j.com/case-studies/ebay/)[.”](https://neo4j.com/case-studies/ebay/)

[“If I were to tell our investigators today that we were doing away with Neo4j, there would be a ](https://neo4j.com/case-studies/zurich-insurance/)[**huge outcry**](https://neo4j.com/case-studies/zurich-insurance/)[.”](https://neo4j.com/case-studies/zurich-insurance/)

[Meredith turned 14B anonymous users into 163M unique profiles, creating a better user experience and leading to ](https://neo4j.com/case-studies/meredith/)[**612% more web traffic**](https://neo4j.com/case-studies/meredith/)[.](https://neo4j.com/case-studies/meredith/)

[Neo4j is “](https://neo4j.com/case-studies/ebay/)[**thousands of times faster**](https://neo4j.com/case-studies/ebay/)[ than our prior MySQL solution, with queries that require ](https://neo4j.com/case-studies/ebay/)[**10-100 times less code**](https://neo4j.com/case-studies/ebay/)[.”](https://neo4j.com/case-studies/ebay/)

[“If I were to tell our investigators today that we were doing away with Neo4j, there would be a ](https://neo4j.com/case-studies/zurich-insurance/)[**huge outcry**](https://neo4j.com/case-studies/zurich-insurance/)[.”](https://neo4j.com/case-studies/zurich-insurance/)

[Meredith turned 14B anonymous users into 163M unique profiles, creating a better user experience and leading to ](https://neo4j.com/case-studies/meredith/)[**612% more web traffic**](https://neo4j.com/case-studies/meredith/)[.](https://neo4j.com/case-studies/meredith/)

Cypher is a powerful, intuitive, graph-optimized query language that understands, and takes advantage of, data connections. It’s user-friendly, easy to learn, and follows connections – in any direction – to reveal previously unknown relationships and clusters.

When trying to find patterns or insights within data, Cypher queries are much simpler and easier to write than massive SQL joins. Since Neo4j doesn’t have tables, there are no joins to worry about. Compare the Cypher query at the left with its equivalent in SQL.

[Learn more about Cypher](https://neo4j.com/developer/guide-sql-to-cypher/?ref=home)

```plain text
MATCH (p:Product)-[:CATEGORY]->(l:ProductCategory)-[:PARENT*0..]->(:ProductCategory {name:"Dairy Products"})
RETURN p.name
```

```plain text
SELECT p.ProductName
FROM Product AS p
JOIN ProductCategory pc ON (p.CategoryID = pc.CategoryID AND pc.CategoryName = "Dairy Products")

JOIN ProductCategory pc1 ON (p.CategoryID = pc1.CategoryID)
JOIN ProductCategory pc2 ON (pc1.ParentID = pc2.CategoryID AND pc2.CategoryName = "Dairy Products")

JOIN ProductCategory pc3 ON (p.CategoryID = pc3.CategoryID)
JOIN ProductCategory pc4 ON (pc3.ParentID = pc4.CategoryID)
JOIN ProductCategory pc5 ON (pc4.ParentID = pc5.CategoryID AND pc5.CategoryName = "Dairy Products");
```

### Use Your Favorite Programming Languages

We aim to make the Neo4j experience fast, natural, and fun for developers. Neo4j supports GraphQL and drivers for .Net, Java, Node.js, Python, and more. Our community of contributors provide many more drivers, including PHP, Ruby, R, Erlang, and Clojure.

[Learn more about Drivers](https://neo4j.com/developer/language-guides/?ref=home)

```plain text
// npm install --save neo4j-driver
// node example.js
const neo4j = require("neo4j-driver");
const driver = neo4j.driver("bolt://<HOST>:<BOLTPORT>", neo4j.auth.basic("<USERNAME>", "<PASSWORD>"), {
	/* encrypted: 'ENCRYPTION_OFF' */
});

const query = `
  MATCH (p:Product)-[:PART_OF]->(:Category)-[:PARENT*0..]->
  (:Category {categoryName:$category})
  RETURN p.productName as product
  `;

const params = { category: "Dairy Products" };

const session = driver.session({ database: "neo4j" });

session
	.run(query, params)
	.then((result) => {
		result.records.forEach((record) => {
			console.log(record.get("product"));
		});
		session.close();
		driver.close();
	})
	.catch((error) => {
		console.error(error);
	});

```

### Helpful Tools for Modern App & Web Development

Neo4j provides an array of tools, libraries, and frameworks to make development faster and easier. Developer tools like Neo4j Desktop, Browser, and Sandbox make it simple to learn and develop graph apps.

The new Neo4j GraphQL Library translates GraphQL queries into Cypher, making it easier for GraphQL users to use Neo4j. It also streamlines integration of Neo4j with React, Vue, and other open source frameworks.
