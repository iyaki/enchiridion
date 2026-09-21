---
title: "capo.js - Get your <head> in order"
notion_id: 7346d7ec-9205-4a48-9b37-f936c9137c79
notion_url: https://app.notion.com/p/capo-js-Get-your-head-in-order-7346d7ec92054a489b37f936c9137c79
last_edited: 2024-03-25T17:32:00.000Z
source_url: https://github.com/rviscomi/capo.js
tags: ["English", "Web Development", "HTML", "Tool"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# [rviscomi/capo.js](https://github.com/rviscomi/capo.js)

_Get your __`<head>`__ in order_

Inspired by [Harry Roberts](https://twitter.com/csswizardry)' work on [ct.css](https://csswizardry.com/ct/) and [Vitaly Friedman](https://twitter.com/smashingmag)'s [Nordic.js 2022 presentation](https://youtu.be/uqLl-Yew2o8?t=2873):

![image](https://private-user-images.githubusercontent.com/1120896/239937765-2319bf3e-21b3-48dd-afcd-a1d379a1daeb.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTA5MzU4NzEsIm5iZiI6MTcxMDkzNTU3MSwicGF0aCI6Ii8xMTIwODk2LzIzOTkzNzc2NS0yMzE5YmYzZS0yMWIzLTQ4ZGQtYWZjZC1hMWQzNzlhMWRhZWIucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI0MDMyMCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNDAzMjBUMTE1MjUxWiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9ZWQwM2U5YjY0MDEyZGY4ZTU3MjUyNWJjY2E1MmJjNWUzOWZiY2I0YWZmMDhkNjRiODZlMzMyZTQ4OWNlYzg3MyZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmYWN0b3JfaWQ9MCZrZXlfaWQ9MCZyZXBvX2lkPTAifQ.q1qx6sDUJLTI2kx4YGBg6vAca39OtuQSUUj8G2EY_uc)

## Why it matters

How you order elements in the `<head>` can have an effect on the (perceived) performance of the page.

This script helps you identify which elements are out of order.

## How to use it

✨ _New: Install the _[_Capo Chrome extension_](https://chrome.google.com/webstore/detail/capo-get-your-%3Chead%3E-in-o/ohabpnaccigjhkkebjofhpmebofgpbeb) ✨

1. Copy [capo.js](https://raw.githubusercontent.com/rviscomi/capo.js/main/snippet/capo.js)
2. Run it in a new [DevTools snippet](https://developer.chrome.com/docs/devtools/javascript/snippets/), or use a [bookmarklet](https://caiorss.github.io/bookmarklet-maker/) generator
3. Explore the console logs

![image](https://private-user-images.githubusercontent.com/1120896/239570699-b29672f9-1f05-4a05-a85e-df27acd153bd.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTA5MzU4NzEsIm5iZiI6MTcxMDkzNTU3MSwicGF0aCI6Ii8xMTIwODk2LzIzOTU3MDY5OS1iMjk2NzJmOS0xZjA1LTRhMDUtYTg1ZS1kZjI3YWNkMTUzYmQucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI0MDMyMCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNDAzMjBUMTE1MjUxWiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9OGQxYTcwOGYyMWQ1NWVhYTNjODYwYzhkOGZlYjNlODJhMWMwYTE2ZjNlNTA1MzE5ZDE4YTQ5ZmExMDllMjQzZSZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmYWN0b3JfaWQ9MCZrZXlfaWQ9MCZyZXBvX2lkPTAifQ.9FfzOcNaYGsQeCezExrQUeSWqoi5jS6XWRg-0scnWHM)

For applications that add lots of dynamic content to the `<head>` on the client, it'd be more accurate to look at the server-rendered `<head>` instead.

### Chrome extension

![image](https://private-user-images.githubusercontent.com/1120896/246672425-389bcec0-567d-448f-9897-eee5ca373e6b.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTA5MzU4NzEsIm5iZiI6MTcxMDkzNTU3MSwicGF0aCI6Ii8xMTIwODk2LzI0NjY3MjQyNS0zODliY2VjMC01NjdkLTQ0OGYtOTg5Ny1lZWU1Y2EzNzNlNmIucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI0MDMyMCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNDAzMjBUMTE1MjUxWiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9YmE1ZWQxNjdhZjlkZTNkZGYxYTM3OWRiZjU0Yzc4ZDc5N2ZhNjNmM2M5M2RkYzZmMzEwM2I4OGViZTU3NmQyNCZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmYWN0b3JfaWQ9MCZrZXlfaWQ9MCZyZXBvX2lkPTAifQ.b7F2DW3bB8OuRj6T7bGT1em0GhwhqjqsTHoy_5OUjOI)

WIP see [crx/](https://github.com/rviscomi/capo.js/blob/main/crx)

### WebPageTest

You can use the [`capo`](https://github.com/rviscomi/capo.js/blob/main/webpagetest)[ WebPageTest custom metric](https://github.com/rviscomi/capo.js/blob/main/webpagetest) to evaluate only the server-rendered HTML `<head>`. Note that because this approach doesn't output to the console, we lose the visualization.

### BigQuery

You can also use the [`httparchive.fn.CAPO`](https://github.com/rviscomi/capo.js/blob/main/bigquery) function on BigQuery to process HTML response bodies in the HTTP Archive dataset. Similar to the WebPageTest approach, the output is very basic.

### Other

Alternatively, you can use local overrides in DevTools to manually inject the capo.js script into the document so that it runs before anything else, eg the first child of `<body>`. Harry Roberts also has a nifty [video](https://www.youtube.com/watch?v=UOn0b5kn3jk) showing how to use this feature. This has some drawbacks as well, for example the inline script might be blocked by CSP.

Another idea would be to use something like Cloudflare workers to inject the script into the HTML stream. To work around CSP issues, you can write the worker in such a way that it parses out the correct `nonce` and adds it to the inline script. _(Note: Not tested, but please share examples if you get it working! __😄__)_

## Summary view

The script logs two info groups to the console: the actual order of the `<head>`, and the optimal order. In this collapsed view, you can see at a glance whether there are any high impact elements out of order.

Each "weight" has a corresponding color, with red being the highest and blue/grey being the lowest. See [capo.js](https://github.com/rviscomi/capo.js/blob/main/capo.js#L1-L13) for the exact mapping.

Here are a few examples.

### [www.nytimes.com](http://www.nytimes.com/)

![image](https://private-user-images.githubusercontent.com/1120896/239564002-5c19e758-9f88-42c1-81e8-9f757a6c92be.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTA5MzU4NzEsIm5iZiI6MTcxMDkzNTU3MSwicGF0aCI6Ii8xMTIwODk2LzIzOTU2NDAwMi01YzE5ZTc1OC05Zjg4LTQyYzEtODFlOC05Zjc1N2E2YzkyYmUucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI0MDMyMCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNDAzMjBUMTE1MjUxWiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9NGU2NmRmOGJlZjk2ZWIwOWRkOWVkNDZmMjUwNGMyYzRiMTUyODllMzNkZmIwYjUwMGEzNzQyMDJmOTA0NjRiYyZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmYWN0b3JfaWQ9MCZrZXlfaWQ9MCZyZXBvX2lkPTAifQ.AA9jfWaFOGK-1eiA8s_ttsSQSEyPmRZPLNyQ4vV1q30)

### docs.github.io

![image](https://private-user-images.githubusercontent.com/1120896/239564471-798a0e99-04dd-4d27-a241-b5d77320a46e.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTA5MzU4NzEsIm5iZiI6MTcxMDkzNTU3MSwicGF0aCI6Ii8xMTIwODk2LzIzOTU2NDQ3MS03OThhMGU5OS0wNGRkLTRkMjctYTI0MS1iNWQ3NzMyMGE0NmUucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI0MDMyMCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNDAzMjBUMTE1MjUxWiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9NjBjMDA4MzJiMTAxMWRhNTEwMTVhZjk0M2I0NzBmOWJjNTNkNTAzZDczZmVjNmIyMDc5NWRlMjZiZDg3NjJkYiZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmYWN0b3JfaWQ9MCZrZXlfaWQ9MCZyZXBvX2lkPTAifQ.3Dn_HtwincxNbyEZ026I-LpDuA7y_nmVle3CgA0vOV0)

### web.dev

![image](https://private-user-images.githubusercontent.com/1120896/239565494-fe6bb67c-697a-4fdf-aa28-52429239fcf5.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTA5MzU4NzEsIm5iZiI6MTcxMDkzNTU3MSwicGF0aCI6Ii8xMTIwODk2LzIzOTU2NTQ5NC1mZTZiYjY3Yy02OTdhLTRmZGYtYWEyOC01MjQyOTIzOWZjZjUucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI0MDMyMCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNDAzMjBUMTE1MjUxWiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9MGU3ZDE4ODgxMzRhMDgyOTM0Y2Y5OGQ1MDE4YTBlNzFhYzE2ZjE0YjhlMzk0YmM5NzJjODM4ODliMWY3YzUyYSZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmYWN0b3JfaWQ9MCZrZXlfaWQ9MCZyZXBvX2lkPTAifQ.vJoXFUooi28wbKlkhzgzngkv-LN5M-nOcItSfMgLR7M)

## stackoverflow.com

![image](https://private-user-images.githubusercontent.com/1120896/239568615-8964fed2-e933-4795-ada6-79c56cbc416d.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTA5MzU4NzEsIm5iZiI6MTcxMDkzNTU3MSwicGF0aCI6Ii8xMTIwODk2LzIzOTU2ODYxNS04OTY0ZmVkMi1lOTMzLTQ3OTUtYWRhNi03OWM1NmNiYzQxNmQucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI0MDMyMCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNDAzMjBUMTE1MjUxWiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9YzM0ZTcxYTQwODA1ODY2ZWRlMmNmNTEzYTczMmUyZDJjNzhhYjYwNTZiMTMzMDgzZDJlMjEwNjY0NGM0YzkxNCZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmYWN0b3JfaWQ9MCZrZXlfaWQ9MCZyZXBvX2lkPTAifQ._PQhDBYwmP5DdCWja_SMVLRj9_R7YUVZ5mqcs69NDa4)

## Detailed view

Expanding the actual or sorted views reveals the detailed view. This includes an itemized list of each `<head>` element and its weight as well as a reference to the actual or sorted `<head>` element.

### [www.nytimes.com](http://www.nytimes.com/)

Here you can see a drilled-down view of the end of the `<head>` for the NYT site, where high impact origin trial meta elements are set too late.

![image](https://private-user-images.githubusercontent.com/1120896/239567027-c0342d54-9e23-4b91-8df1-277c251ee0c5.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTA5MzU4NzEsIm5iZiI6MTcxMDkzNTU3MSwicGF0aCI6Ii8xMTIwODk2LzIzOTU2NzAyNy1jMDM0MmQ1NC05ZTIzLTRiOTEtOGRmMS0yNzdjMjUxZWUwYzUucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI0MDMyMCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNDAzMjBUMTE1MjUxWiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9OWI2ZWMwNzAyMmIxYjM3Zjg2ZGFhYTMzNzljNWQ3Y2ZlNzYwNTI2YTEwZmQ0ZTQxNWNmNTBhMjdmYmNjNmFkNCZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmYWN0b3JfaWQ9MCZrZXlfaWQ9MCZyZXBvX2lkPTAifQ.2Rz_tqmJYcj1PoB894h_ZnRKSRmZvloxXvcdjjKpO4Y)
