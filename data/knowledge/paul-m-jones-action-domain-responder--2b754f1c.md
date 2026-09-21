---
title: "Paul M. Jones | Action Domain Responder"
notion_id: 2b754f1c-7d23-81aa-b92d-c371bdf9f330
notion_url: https://app.notion.com/p/Paul-M-Jones-Action-Domain-Responder-2b754f1c7d2381aab92dc371bdf9f330
last_edited: 2025-11-26T14:48:00.000Z
source_url: https://pmjones.io/adr/
tags: ["Article", "Paul M. Jones", "English", "System Design / Software Architecture", "Web Development", "PHP"]
---
organizes a single user interface interaction between an HTTP client and a HTTP server-side application into three distinct roles.

![image](https://pmjones.io/adr/adr.png)

## Components

_Action_ is the logic to connect the _Domain_ and _Responder_. It invokes the _Domain_ with inputs collected from the HTTP Request, then invokes the _Responder_ with the data needed to build an HTTP Response.

_Domain_ is an entry point to the domain logic forming the core of the application. It may be a _Transaction Script_, _Service Layer_, _Application Service_, or something similar.

_Responder_ is the presentation logic to build an HTTP Response using data it receives from the _Action_. It deals with status codes, headers and cookies, content, formatting and transformation, templates and views, and so on.

## Collaborations

1. 
2. 
3. 
4. 
5. 

## Reading

- [_Model View Controller_](https://github.com/pmjones/adr/blob/master/MVC-MODEL-2.md)[ and "Model 2"](https://github.com/pmjones/adr/blob/master/MVC-MODEL-2.md)
- [Comparing "Model 2" MVC to ADR](https://github.com/pmjones/adr/blob/master/ADR.md)
- [Tradeoffs in ADR](https://github.com/pmjones/adr/blob/master/TRADEOFFS.md)
- [Objections to ADR](https://github.com/pmjones/adr/blob/master/OBJECTIONS.md)

## Resources

- [Refactoring from "Model 2" MVC to ADR](https://github.com/pmjones/adr/blob/master/REFACTORING.md)
- [Implementation notes and advice](https://github.com/pmjones/adr/blob/master/IMPLEMENTATION.md)
- [ADR discussions, mentions, implementations, etc.](https://github.com/pmjones/adr/blob/master/MENTIONS.md)

Example code resides in the [`example`](https://github.com/pmjones/adr-example) repository.

You can also find out about [the history of this pattern](https://github.com/pmjones/adr/blob/master/HISTORY.md), [the research bibliography](https://github.com/pmjones/adr/blob/master/BIBLIO.md), and [the author's acknowledgments](https://github.com/pmjones/adr/blob/master/ACKNOWLEDGEMENTS.md).
