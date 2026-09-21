---
title: "To test or not to test, a technical perspective"
notion_id: 6c41a81a-a724-4589-ad35-e3404fe35ea6
notion_url: https://app.notion.com/p/To-test-or-not-to-test-a-technical-perspective-6c41a81aa7244589ad35e3404fe35ea6
last_edited: 2023-09-04T14:07:00.000Z
source_url: https://web.dev/ta-what-to-test
tags: ["English", "Testing", "Article", "web.dev"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->









<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## 



### 





<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->





### 



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



### 











### 







### 







## 



### 



- 
- 
- 
- 
- 
- 

| Do ✅ | Don't ❌ |
| --- | --- |
| Keep the tests as small as possible. Test one thing per test case. | Write tests over large units. |
| Always keep tests isolated and mock the things you need which are outside your unit. | Include other components or services. |
| Keep tests independent. | Rely on previous tests or share test data. |
| Cover [different scenarios and paths](https://web.dev/ta-test-cases/#test-paths-typical-kinds-of-test-cases). | Limit yourself to the happy path or negative tests at maximum. |
| Use descriptive test titles, so you can immediately see what your test is about. | Test by function name only, not being descriptive enough as a result: `testBuildFoo()` or `testGetId()`. |
| Aim for good code coverage or a broader range of test cases, especially at this stage. | Test from every class down to database (I/O) level. |

### 



- 
- 
- 

| Do ✅ | Don't ❌ |
| --- | --- |
| Test the integration points: verify that each unit works together gracefully when integrated with each other. | Test each unit in isolation—that's what unit tests are for. |
| Test real-world scenarios: use test data derived from real-world data. | Use repetitive auto-generated test data or other data which doesn't reflect real-world use cases. |
| Use mocks and stubs for external dependencies to maintain control of your complete test. | Create dependencies on third-party services, for example, network requests to outside services. |
| Use a clean-up routine before and after each test. | Forget to use clean-up measures inside your tests, otherwise this can lead to test failures or false positives, due to lack of proper test isolation. |

### 



- 
- 
- 
- 
- 
- 

| Do ✅ | Don't ❌ |
| --- | --- |
| Use API-driven shortcuts. [Learn more](https://docs.cypress.io/guides/references/best-practices#Organizing-Tests-Logging-In-Controlling-State). | Use UI interactions for every step, including the `beforeEach` hook. |
| Use a clean-up routine before each test. Take even more care of test isolation than you do in unit and integration tests because there's a higher risk of side effects here. | Forget to clean up after each test. If you don't clean up the leftover state, data or side effects, they will affect other tests executed later. |
| Regard end-to-end tests as system tests. This means you need to test the whole application stack. | Test each unit in isolation—that's what unit tests are for. |
| Use minimal or no mocking inside the test. Consider carefully if you want to mock external dependencies. | Rely heavily on mocks. |
| Consider performance and workload by, for example, not over-testing large scenarios in the same test. | Cover large workflows without using shortcuts. |
