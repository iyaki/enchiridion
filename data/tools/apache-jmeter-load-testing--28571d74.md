---
title: "Apache JMeter - Load testing"
notion_id: 28571d74-b3a3-4ef2-8424-432275de9aef
notion_url: https://app.notion.com/p/Apache-JMeter-Load-testing-28571d74b3a34ef28424432275de9aef
last_edited: 2023-04-25T13:41:00.000Z
source_url: https://jmeter.apache.org/
tags: ["Web Development", "Testing", "REST API", "Tool", "English"]
---
[https://jmeter.apache.org/](https://jmeter.apache.org/)

The Apache JMeter™ application is open source software, a 100% pure Java application designed to load test functional behavior and measure performance. It was originally designed for testing Web Applications but has since expanded to other test functions.

## What can I do with it?

Apache JMeter may be used to test performance both on static and dynamic resources, Web dynamic applications. It can be used to simulate a heavy load on a server, group of servers, network or object to test its strength or to analyze overall performance under different load types.

Apache JMeter features include:

- Ability to load and performance test many different applications/server/protocol types: Web - HTTP, HTTPS (Java, NodeJS, PHP, ASP.NET, …) SOAP / REST Webservices FTP Database via JDBC LDAP Message-oriented middleware (MOM) via JMS Mail - SMTP(S), POP3(S) and IMAP(S) Native commands or shell scripts TCP Java Objects
- Full featured Test IDE that allows fast Test Plan recording (from Browsers or native applications), building and debugging.
- CLI mode (Command-line mode (previously called Non GUI) / headless mode) to load test from any Java compatible OS (Linux, Windows, Mac OSX, …)
- A complete and ready to present dynamic HTML report
- Easy correlation through ability to extract data from most popular response formats, HTML, JSON , XML or any textual format
- Complete portability and 100% Java purity.
- Full multi-threading framework allows concurrent sampling by many threads and simultaneous sampling of different functions by separate thread groups.
- Caching and offline analysis/replaying of test results.
- Highly Extensible core: Pluggable Samplers allow unlimited testing capabilities. Scriptable Samplers (JSR223-compatible languages like Groovy and BeanShell) Several load statistics may be chosen with pluggable timers. Data analysis and visualization plugins allow great extensibility as well as personalization. Functions can be used to provide dynamic input to a test or provide data manipulation. Easy Continuous Integration through 3rd party Open Source libraries for Maven, Gradle and Jenkins.

## How do I do it?

- Using JMeter to understand how to use it
- Component reference to have detailed information for every Test element
- Functions reference to have detailed information and examples for every function
- Properties reference for all properties that allow you to customize JMeter
- Javadoc API documentation
- JMeter FAQ (Wiki)
- JMeter Wiki
- Building JMeter and Add-Ons for advanced usage

## JMeter is not a browser

JMeter is not a browser, it works at protocol level. As far as web-services and remote services are concerned, JMeter looks like a browser (or rather, multiple browsers); however JMeter does not perform all the actions supported by browsers. In particular, JMeter does not execute the Javascript found in HTML pages. Nor does it render the HTML pages as a browser does (it's possible to view the response as HTML etc., but the timings are not included in any samples, and only one sample in one thread is ever displayed at a time).

## Tutorials

- Distributed Testing
- Recording Tests
- JUnit Sampler
- Access Log Sampler
- Extending JMeter

## Further Information About JMeter

- Changes List
- Read about existing Issues (Bugs or Enhancements) or reporting new ones (please do it !)
- License
- Mailing Lists
- Source Repositories
- Contributors

Go to top
