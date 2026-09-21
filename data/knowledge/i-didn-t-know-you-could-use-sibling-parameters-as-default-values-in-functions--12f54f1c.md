---
title: "I didn't know you could use sibling parameters as default values in functions"
notion_id: 12f54f1c-7d23-814f-b0d3-ea21aa34b893
notion_url: https://app.notion.com/p/I-didn-t-know-you-could-use-sibling-parameters-as-default-values-in-functions-12f54f1c7d23814fb0d3ea21aa34b893
last_edited: 2024-11-13T20:17:00.000Z
source_url: https://macarthur.me/posts/sibling-parameters/
tags: ["Alex MacArthur", "English", "Javascript", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

JavaScript has supported default parameter values since ES2015. You know this. I know this. What I _didn't_ know was that you can use _sibling_ _parameters_ as the default values themselves. (Or maybe "adjacent positional parameters"? Not sure what to call these.)

```javascript
function myFunc(arg1, arg2 = arg1) {
  console.log(arg1, arg2);
}

myFunc("arg1!");
// "arg1!" "arg1!"
```

This works in class constructors too – something I found to be quite helpful in making some [PicPerf.io](https://picperf.io/?ref=cms.macarthur.me) code more testable. It's common to see simple dependency injection used to that end. Let's explore it a bit.

## [A Scenario](https://macarthur.me/posts/sibling-parameters/#a-scenario)

Keeping with the image optimization theme, say you have an `OptimizedImage` class. Provide an image URL to its constructor, and you can retrieve either a freshly optimized buffer of the image or a cached version.

```javascript
class OptimizedImage {
  constructor(
    imageUrl: string,
    cacheService = new CacheService(),
    optimizeService = new OptimizeService()
  ) {
    this.imageUrl = imageUrl;
    this.cacheService = cacheService;
    this.optimizeService = optimizeService;
  }

  async get() {
    const cached = this.cacheService.get(this.imageUrl);

    // Return the previously optimized image.
    if (cached) return cached;

    const optimizedImage = await this.optimizeService
      .optimize(this.imageUrl);

    // Cache the optimized image for next time.
    return this.cacheService.put(this.imageUrl, optimizedImage);
  }
}

const instance = new OptimizedImage('https://macarthur.me/me.jpg');
const imgBuffer = await instance.get();
```

The only constructor parameter used in production is `imageUrl`, but injecting `CacheService` and `OptimizeService` enables easier unit test with mocks:

```javascript
import { it, expect, vi } from 'vitest';
import { OptimizedImage } from './main';

it('returns freshly optimized image', async function () {
  const fakeImageBuffer = new ArrayBuffer('image!');
  const mockCacheService = {
    get: (url) => null,
    put: vi.fn().mockResolvedValue(fakeImageBuffer),
  };

  const mockOptimizeService = {
    optimize: (url) => fakeImageBuffer,
  };

  const optimizedImage = new OptimizedImage(
    'https://test.jpg',
    mockCacheService,
    mockOptimizeService
  );

  const result = await optimizedImage.get();

  expect(result).toEqual(fakeImageBuffer);
  expect(mockCacheService.put).toHaveBeenCalledWith(
    'https://test.jpg',
    'optimized image'
  );
});
```

## [Making It More Complicated](https://macarthur.me/posts/sibling-parameters/#making-it-more-complicated)

In that example, both of those service classes use `imageUrl` only when particular methods are invoked. But imagine if they required it to be passed into their own constructors instead. You might be tempted to pull instantiation into `OptimizedImage`'s constructor (I was):

```javascript
class OptimizedImage {
  constructor(
    imageUrl: string
  ) {
    this.imageUrl = imageUrl;
    this.cacheService = new CacheService(imageUrl);
    this.optimizeService = new OptimizeService(imageUrl);
  }
```

That’d work, but now `OptimizedImage` is fully responsible for service instantiation, and testing becomes more of a hassle too. It's not so easy to pass in mocks for service instances.

You could get around this by passing in mock class definitions, but you'd then need create mock versions of those classes with their own constructors, making testing more tedious. Fortunately, there's another option: use the `imageUrl` parameter in the rest of your argument list.

## [Sharing Sibling Parameters](https://macarthur.me/posts/sibling-parameters/#sharing-sibling-parameters)

I wasn't aware this was even possible until a little while ago. Here's how it'd look:

```javascript
export class OptimizedImage {
  constructor(
    imageUrl: string,
    // Use the same `imageUrl` in both dependencies.
    cacheService = new CacheService(imageUrl),
    optimizeService = new OptimizeService(imageUrl)
  ) {
    this.cacheService = cacheService;
    this.optimizeService = optimizeService;
  }

  async get() {
    const cached = this.cacheService.get();

    // Return the previously optimized image.
    if (cached) return cached;

    const optimizedImage = await this.optimizeService.optimize();

    // Cache the optimized image for next time.
    return this.cacheService.put(optimizedImage);
  }
}
```

With this setup, you're able to mock those instances just as easily as before, and the rest of the class doesn't even need to hold onto an instance of `imageUrl` itself. Instantiation, of course, still remains simple:

```javascript
const instance = new OptimizedImage('https://macarthur.me/me.jpg');

const img = await instance.get();
```

The same testing approach remains in tact as well:

```javascript
import { it, expect, vi } from 'vitest';
import { OptimizedImage } from './main';

it('returns freshly optimized image', async function () {
  const mockCacheService = {
    get: () => null,
    put: vi.fn().mockResolvedValue('optimized image'),
  };

  const mockOptimizeService = {
    optimize: () => 'optimized image',
  };

  const optimizedImage = new OptimizedImage(
    'https://test.jpg',
    mockCacheService,
    mockOptimizeService
  );

  const result = await optimizedImage.get();

  expect(result).toEqual('optimized image');
  expect(mockCacheService.put).toHaveBeenCalledWith('optimized image');
});
```

Nothing groundbreaking here – just a small feature that made my life a little more ergonomically pleasing. I'm hoping to come across more gems like this in the future.
