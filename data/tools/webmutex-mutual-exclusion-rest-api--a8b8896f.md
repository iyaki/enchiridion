---
title: "WebMutex - Mutual Exclusion REST API"
notion_id: a8b8896f-734f-49c5-8177-916b3d92af4a
notion_url: https://app.notion.com/p/WebMutex-Mutual-Exclusion-REST-API-a8b8896f734f49c58177916b3d92af4a
last_edited: 2023-04-22T01:16:00.000Z
source_url: https://webmutex.io/
tags: ["Programming", "Network", "REST API", "Untried", "Service"]
---
## API

Parameters can be passed as json in the body of the request, or as http form parameters.

All requests also accept the mutex id via `/request/<id>/`

Service is free for now until I get overwhelmed by requests and I have to upgrade from the cheapest Linode instance, which will happen probably never.

### Reserve a mutex: `/reserve`

Allocates a new mutex.

Parameters: (none)

Returns:

- `id`
- `status`

`POST` only

### Grab an existing mutex: `/grab`

Tries to grab an existing mutex if it’s currently free.

Parameters:

- `id`

Returns:

- `id`
- `token`
- `status`

`POST` only

### Release a mutex: `/release`

Releases the mutex, using the token received from `/grab`.

Parameters:

- `id`
- `token`

Returns:

- `status`

`POST` only

### Checks mutex status: `/status`

Checks the current status of the mutex.

Parameters:

- `id`

Returns:

- `status`
- `in_use`

`GET` and `POST`

### Grabs as soon as possible: `/subscribe_and_grab` (**NOT IMPLEMENTED YET**)

Contacts the endpoint when the mutex has been grabbed.

Multiple users can subscribe to the same mutex. As soon as the mutex is available, one will be woken up.

Parameters:

- `id`
- `endpoint`

Returns:

- `status`

`POST` only

## Example use

Multiple tests in your CI are running concurrently. You want to save their results in some database but you can’t be arsed to set one up properly, so you figure you’ll just use a sqlite file and sync it to artifactory, as sqlite can sort the concurrency aspect out for you.

You proceed to naively wrap each test in

```plain text
curl -O $artifactory_url/test_results.db
./build/mytest > test_results.json
./test_result_extractor test_results.json test_results.db
curl -X PUT $artifactory_url --data-binary @test_results.db
```

Now you realise that each test is running on its own instance of your CI, so they’re all grabbing the db and using their own local copy. There is no concurrency aspect, and when things get pushed to artifactory they overwrite each other. You screwed up, _again_, like you always do. This is why nobody loves you.

But you figure that not all hope is lost. You reach this service. Now you just need to grab a mutex before fetching your db.

## COMING SOON

- the subscription thingy
- auto expiration
- usage limits
- paid enterprise tier because why not
- ???
- nice graphs? idk
- drop-in replacement library for std::mutex (see [nft_ptr](https://github.com/zhuowei/nft_ptr))

[Source](https://github.com/izabera/webmutex)
