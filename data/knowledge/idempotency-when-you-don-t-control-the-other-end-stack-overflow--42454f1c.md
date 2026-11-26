---
title: "Idempotency when you don't control the other end - Stack Overflow"
notion_id: 42454f1c-7d23-8108-80b3-cb7ce9e13220
notion_url: https://app.notion.com/p/Idempotency-when-you-don-t-control-the-other-end-Stack-Overflow-42454f1c7d23810880b3cb7ce9e13220
last_edited: 2026-11-26T06:48:00.000Z
source_url: https://stackoverflow.blog/2026/11/24/idempotency-when-you-don-t-control-the-other-end/
tags: ["Stack Overflow Blog", "English", "APIs", "Communication", "Error Handling", "Performance", "Productivity", "Article", "Note"]
---
Idempotency looks like a solved problem. If you build the API, you accept an `Idempotency-Key` header, store it with the response, and when the same key comes back you return the stored response instead of doing the work again. Stripe explained this years ago, and most of us have copied it at some point.

The catch is that I sit on both sides. I work at [maesn](https://www.maesn.com/), where we build a unified API: our customers call us, and we pass their data on to 30+ accounting and ERP systems that other companies built. On our own API we can do what Stripe does. On the other side, we are the caller. When a request to one of those systems times out, I can't look into their database. I have to decide whether to send it again without knowing what happened, and many of those systems give me no key to make that safe.

Most of my work is bookkeeping: getting journal entries into accounting systems. Many of those APIs accept entries one at a time, and many also take whole batches. On some platforms a tax advisor reviews what arrives before it counts. On others the entries go straight into the books. Either way, once the entries are locked, you can't simply change or delete them anymore. You fix it with a correction entry, and somebody has to explain that correction later.

Many of the systems we work with offer no idempotency key, and that turns every timeout into a guessing game for whoever is calling. A dropped connection or a client that gives up too early, and the response is gone. The caller can't tell whether the entries were created: either the request never arrived, or it arrived and the response got lost on the way back. From the outside, those two cases look identical. Retry, and the entries may be booked twice. Hold back, and they may be missing.

Batches make it worse. On one platform we work with, journal entries arrive as a batch file, so a single upload carries many postings at once. If the caller loses the response and sends the upload again, the platform creates the second batch exactly like the first.

![image](https://cdn.stackoverflow.co/images/jo7n4k8s/production/ceff20e9b6bf45d0a7e43e45720fa7ff94ec97fa-2047x1743.jpg?auto=format)

So I tried it. I uploaded a small batch of journal entries, then uploaded the exact same file again. Both uploads came back successful. Nothing on the API side hinted that the second one was a copy.

The platform only noticed at import. Each posting in the copy got a notice saying that an entry with the same amount, document number and date already existed on that account, with a pointer to the original. A yellow notice, sitting in the import log between unrelated hints about tax codes. One click on OK and the copy was in the books as a second batch, same name and same period, right below the original.

So the platform does check for duplicates. It just tells the tax advisor, at the very end, in a list that is easy to scroll past. The caller who sent the duplicate never hears a thing.

And that platform is the forgiving case. On some systems we post to, entries are booked the moment they arrive, with no review step and no notice. There, nothing catches the duplicate.

If the system offered an idempotency key, a lost response would be harmless: the caller sends the same key again and gets the original answer back. Many of the systems we talk to don't offer one, so callers fall back on three workarounds. I have seen each of them break.

The first thing everyone tries is a lookup. Most journal entries carry a document number, usually the supplier's invoice number. So after a timeout you ask the accounting system's API whether an entry with that number already exists. If you find it, the first request went through. If not, you send it again.

The catch is that a match doesn't prove your request landed. It only proves that some entry with that number exists. The number was only ever unique for the supplier who wrote it. Two suppliers can both send invoice 1001, and many restart their numbering every January, so last year's invoice 17 is still in the books when this year's invoice 17 arrives. Your lookup finds the old one, concludes the new one went through and skips it.

So you make the search stricter and compare the account, amount and date as well. That is roughly what the platform in my test checks, and it catches far more. But two entries can match on all of that and still be two different things. I once found the same posting twice in a customer's data, on the same account and for the same 4,616.50. Maybe it had been exported twice. Maybe someone had bought two identical assets on the same day and booked each one. Nothing in the two rows could tell me which it was. Only the trial balance could.

I suspect that is why the platform only raises a notice instead of rejecting anything: its own match can be wrong, and it knows it. For the caller, a wrong "found" is worse than a wrong "not found". A duplicate at least shows up as two lines in the ledger. A real entry skipped as a duplicate shows up as nothing, until someone wonders why the numbers don't add up.

Now assume you found a field that really is unique. Checking for it before every write still has a weak spot, and this one has nothing to do with the data. It's timing. Here is how it plays out when the accounting system is having a slow day:

```plain text
10:00:00  attempt 1  search INV-1001   -> nothing found
10:00:00  attempt 1  create INV-1001   -> still processing...
10:00:30  HTTP client gives up after 30 seconds and retries
10:00:30  attempt 2  search INV-1001   -> nothing found, attempt 1 isn't done yet
10:00:30  attempt 2  create INV-1001
10:00:41  attempt 1 finishes          -> entry #1
10:00:43  attempt 2 finishes          -> entry #2
```

The check only works if the first write has finished before the second search runs, and nothing guarantees that. In the timeline, the retry comes from your own HTTP client, while the first request is still being processed on the other side. A queue does the same thing when a worker loses its lock halfway through a job and another worker picks it up. Two attempts at the same entry overlap, and both searches come back empty.

Some systems stretch the overlap much further. A record you have just created doesn't always show up in search straight away, because search reads from an index that is updated in the background. Until that index catches up, every check comes back empty, even after the first write has finished.

And the check isn't free. The search is a request to the accounting system too, so every journal entry now costs two calls instead of one. Many of these systems cap how many requests you can send per minute. Half of that allowance now goes to searching, so you can post half as many entries, and the extra traffic makes timeouts more likely. The check ends up feeding the problem it was meant to solve.

The third workaround gives up on preventing duplicates and cleans them up afterwards. You write, you retry whenever you're unsure, and a job runs later, nightly or weekly, that compares what you sent with what the accounting system now holds and flags anything that appears twice.

The weak spot is the time between the posting and that job. Until the job runs, the duplicate is a real entry in someone's books, and taking it out again is harder than it sounds. On some systems we post to there is no delete for journal entries at all. The only way to undo one is to send a second entry that reverses it. So the clean-up becomes another write into the books, with its own timeouts and its own chance of booking the reversal twice. Of the three workarounds, this is the one I trust least.

It gets worse at month-end. When the accountant closes a month, its entries are locked. From then on, even a system that normally lets you delete an entry won't let you touch it, and the only way to fix a duplicate is a correction entry in the next month. If the VAT return for the closed month has already been filed, the duplicate is in it. The correction then shows up in next month's return, and someone has to explain to the tax advisor why.

The job is also slower than it sounds. To find a duplicate, it has to pull back every entry the accounting system holds for the month, one page at a time, and compare each one with what you sent. For a customer with thousands of entries, that runs for a long time, and the books keep changing underneath it. By the time the job finishes, it has checked a month that no longer looks the way it did when it started.

We tried fixing it from the middle first. Our customers can now send us a key, and we keep it together with the result. A double click or a lost response between them and us now comes back as the original answer, without a second write. But that doesn't touch the last step. Our own request to the accounting system can time out exactly like theirs could, and then we are the caller who doesn't know. A key in the middle moves the uncertainty one step down the chain. Only the system that stores the entry can remove it.

I don't think providers skip idempotency out of carelessness. Most accounting APIs were built to get data out, or to mirror what a person does in the UI, one document at a time. A person rarely submits a form twice, and when they do, they notice. An integration retrying at 3 a.m. does it routinely, and nobody notices until the accountant does.

Even the providers that do offer a key each do it differently. On one system the key goes in a header, on another it goes in the URL. Some forget the key after a while, and if your retry arrives after that, it is treated as a new request. Some only accept the key on certain endpoints, which may not include the one you need. So "we support idempotency" in the documentation tells me very little. I test it myself: send the same request twice with the same key, check whether one entry was created or two, and repeat that after their next release.

Getting a key right is real work: the provider stores keys with their responses, defines what counts as the same request and documents how long a key lives. The details decide whether callers can trust the result. A second request whose key is still in flight should get a 409, and the same key with a different body should get a 422, because that is a bug on the caller's side and it needs to surface. A failed request should leave its key free. Store the error against the key, and one transient outage blocks the operation until the key expires.

If full keys sound like too much work, there is a simpler way to get the same safety for the case that hurts most, creating entries. Give every entry a field where the caller puts its own reference, and refuse a second entry with the same reference for the same company. When a duplicate arrives, answer with a 409 and the entry that already exists. A retry then either creates the entry, because the first attempt never arrived, or gets back the one that did. Either way, there is exactly one.

![image](https://cdn.stackoverflow.co/images/jo7n4k8s/production/31bfa27f04997ea3c505f78c7383bb0b42c5c277-2046x1835.jpg?auto=format)

For the provider, that is one rule in the database and one status code. For the caller, it replaces all three workarounds with a retry that is always safe. So if you are building an API that writes into someone's books, give your callers an idempotency key, or at the very least that field. Otherwise they will build the three workarounds themselves, and each of them breaks somewhere.
