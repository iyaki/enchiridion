---
title: "The Open Graph protocol"
notion_id: c962549a-966c-4a78-8732-9e4bb997c361
notion_url: https://app.notion.com/p/The-Open-Graph-protocol-c962549a966c4a7887329e4bb997c361
last_edited: 2023-01-18T17:57:00.000Z
source_url: https://ogp.me/
tags: ["English", "Web Development", "Website"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The [Open Graph protocol](https://ogp.me/) enables any web page to become a rich object in a social graph. For instance, this is used on Facebook to allow any web page to have the same functionality as any other object on Facebook.

While many different technologies and schemas exist and could be combined together, there isn't a single technology which provides enough information to richly represent any web page within the social graph. The Open Graph protocol builds on these existing technologies and gives developers one thing to implement. Developer simplicity is a key goal of the Open Graph protocol which has informed many of [the technical design decisions](https://www.scribd.com/doc/30715288/The-Open-Graph-Protocol-Design-Decisions).

To turn your web pages into graph objects, you need to add basic metadata to your page. We've based the initial version of the protocol on [RDFa](https://en.wikipedia.org/wiki/RDFa) which means that you'll place additional `<meta>` tags in the `<head>` of your web page. The four required properties for every page are:

- `og:title` - The title of your object as it should appear within the graph, e.g., "The Rock".
- `og:type` - The [type](https://ogp.me/#types) of your object, e.g., "video.movie". Depending on the type you specify, other properties may also be required.
- `og:image` - An image URL which should represent your object within the graph.
- `og:url` - The canonical URL of your object that will be used as its permanent ID in the graph, e.g., "https://www.imdb.com/title/tt0117500/".

As an example, the following is the Open Graph protocol markup for [The Rock on IMDB](https://www.imdb.com/title/tt0117500/):

```plain text
<html prefix="og: https://ogp.me/ns#">
<head>
<title>The Rock (1996)</title>
<meta property="og:title" content="The Rock" />
<meta property="og:type" content="video.movie" />
<meta property="og:url" content="https://www.imdb.com/title/tt0117500/" />
<meta property="og:image" content="https://ia.media-imdb.com/images/rock.jpg" />
...
</head>
...
</html>

```

The following properties are optional for any object and are generally recommended:

- `og:audio` - A URL to an audio file to accompany this object.
- `og:description` - A one to two sentence description of your object.
- `og:determiner` - The word that appears before this object's title in a sentence. An [enum](https://ogp.me/#enum) of (a, an, the, "", auto). If `auto` is chosen, the consumer of your data should chose between "a" or "an". Default is "" (blank).
- `og:locale` - The locale these tags are marked up in. Of the format `language_TERRITORY`. Default is `en_US`.
- `og:locale:alternate` - An [array](https://ogp.me/#array) of other locales this page is available in.
- `og:site_name` - If your object is part of a larger web site, the name which should be displayed for the overall site. e.g., "IMDb".
- `og:video` - A URL to a video file that complements this object.

For example (line-break solely for display purposes):

```plain text
<meta property="og:audio" content="https://example.com/bond/theme.mp3" />
<meta property="og:description"
  content="Sean Connery found fame and fortune as the
           suave, sophisticated British agent, James Bond." />
<meta property="og:determiner" content="the" />
<meta property="og:locale" content="en_GB" />
<meta property="og:locale:alternate" content="fr_FR" />
<meta property="og:locale:alternate" content="es_ES" />
<meta property="og:site_name" content="IMDb" />
<meta property="og:video" content="https://example.com/bond/trailer.swf" />

```

The RDF schema (in [Turtle](https://en.wikipedia.org/wiki/Turtle_(syntax))) can be found at [ogp.me/ns](https://ogp.me/ns/ogp.me.ttl).

Some properties can have extra metadata attached to them. These are specified in the same way as other metadata with `property` and `content`, but the `property` will have extra `:`.

The `og:image` property has some optional structured properties:

- `og:image:url` - Identical to `og:image`.
- `og:image:secure_url` - An alternate url to use if the webpage requires HTTPS.
- `og:image:width` - The number of pixels wide.
- `og:image:height` - The number of pixels high.
- `og:image:alt` - A description of what is in the image (not a caption). If the page specifies an og:image it should specify `og:image:alt`.

A full image example:

```plain text
<meta property="og:image" content="https://example.com/ogp.jpg" />
<meta property="og:image:secure_url" content="https://secure.example.com/ogp.jpg" />
<meta property="og:image:type" content="image/jpeg" />
<meta property="og:image:width" content="400" />
<meta property="og:image:height" content="300" />
<meta property="og:image:alt" content="A shiny red apple with a bite taken out" />

```

The `og:video` tag has the identical tags as `og:image`. Here is an example:

```plain text
<meta property="og:video" content="https://example.com/movie.swf" />
<meta property="og:video:secure_url" content="https://secure.example.com/movie.swf" />
<meta property="og:video:type" content="application/x-shockwave-flash" />
<meta property="og:video:width" content="400" />
<meta property="og:video:height" content="300" />

```

The `og:audio` tag only has the first 3 properties available (since size doesn't make sense for sound):

```plain text
<meta property="og:audio" content="https://example.com/sound.mp3" />
<meta property="og:audio:secure_url" content="https://secure.example.com/sound.mp3" />
<meta property="og:audio:type" content="audio/mpeg" />

```

If a tag can have multiple values, just put multiple versions of the same `<meta>` tag on your page. The first tag (from top to bottom) is given preference during conflicts.

```plain text
<meta property="og:image" content="https://example.com/rock.jpg" />
<meta property="og:image" content="https://example.com/rock2.jpg" />

```

Put structured properties after you declare their root tag. Whenever another root element is parsed, that structured property is considered to be done and another one is started.

For example:

```plain text
<meta property="og:image" content="https://example.com/rock.jpg" />
<meta property="og:image:width" content="300" />
<meta property="og:image:height" content="300" />
<meta property="og:image" content="https://example.com/rock2.jpg" />
<meta property="og:image" content="https://example.com/rock3.jpg" />
<meta property="og:image:height" content="1000" />

```

means there are 3 images on this page, the first image is `300x300`, the middle one has unspecified dimensions, and the last one is `1000`px tall.

In order for your object to be represented within the graph, you need to specify its type. This is done using the `og:type` property:

```plain text
<meta property="og:type" content="website" />

```

When the community agrees on the schema for a type, it is added to the list of global types. All other objects in the type system are [CURIEs](https://en.wikipedia.org/wiki/CURIE) of the form

```plain text
<head prefix="my_namespace: https://example.com/ns#">
<meta property="og:type" content="my_namespace:my_type" />

```

The global types are grouped into verticals. Each vertical has its own namespace. The `og:type` values for a namespace are always prefixed with the namespace and then a period. This is to reduce confusion with user-defined namespaced types which always have colons in them.

`og:type` values:

- `music:duration` - [integer](https://ogp.me/#integer) >=1 - The song's length in seconds.
- `music:album:disc` - [integer](https://ogp.me/#integer) >=1 - Which disc of the album this song is on.
- `music:album:track` - [integer](https://ogp.me/#integer) >=1 - Which track this song is.
- `music:song` - [music.song](https://ogp.me/#type_music.song) - The song on this album.
- `music:song:disc` - [integer](https://ogp.me/#integer) >=1 - The same as `music:album:disc` but in reverse.
- `music:song:track` - [integer](https://ogp.me/#integer) >=1 - The same as `music:album:track` but in reverse.
- `music:musician` - [profile](https://ogp.me/#type_profile) - The musician that made this song.
- `music:release_date` - [datetime](https://ogp.me/#datetime) - The date the album was released.
- `music:song:disc`
- `music:song:track`
- `music:creator` - [profile](https://ogp.me/#type_profile) - The creator of this playlist.
- `music:creator` - [profile](https://ogp.me/#type_profile) - The creator of this station.

`og:type` values:

- `video:actor:role` - [string](https://ogp.me/#string) - The role they played.
- `video:duration` - [integer](https://ogp.me/#integer) >=1 - The movie's length in seconds.
- `video:release_date` - [datetime](https://ogp.me/#datetime) - The date the movie was released.
- `video:actor:role`
- `video:director`
- `video:writer`
- `video:duration`
- `video:release_date`
- `video:tag`
- `video:series` - [video.tv_show](https://ogp.me/#type_video.tv_show) - Which series this episode belongs to.

A multi-episode TV show. The metadata is identical to [video.movie](https://ogp.me/#type_video.movie).

A video that doesn't belong in any other category. The metadata is identical to [video.movie](https://ogp.me/#type_video.movie).

These are globally defined objects that just don't fit into a vertical but yet are broadly used and agreed upon.

`og:type` values:

- `article:published_time` - [datetime](https://ogp.me/#datetime) - When the article was first published.
- `article:modified_time` - [datetime](https://ogp.me/#datetime) - When the article was last changed.
- `article:expiration_time` - [datetime](https://ogp.me/#datetime) - When the article is out of date after.

[`book`](https://ogp.me/#type_book) - Namespace URI: [`https://ogp.me/ns/book#`](https://ogp.me/ns/book)

- `profile:first_name` - [string](https://ogp.me/#string) - A name normally given to an individual by a parent or self-chosen.
- `profile:last_name` - [string](https://ogp.me/#string) - A name inherited from a family or marriage and by which the individual is commonly known.
- `profile:username` - [string](https://ogp.me/#string) - A short unique string to identify them.
- `profile:gender` - [enum](https://ogp.me/#enum)(male, female) - Their gender.

No additional properties other than the basic ones. Any non-marked up webpage should be treated as `og:type` website.

The following types are used when defining attributes in Open Graph protocol.

| **Type** | **Description** | **Literals** |
| --- | --- | --- |
| [Boolean](https://ogp.me/#bool) | A Boolean represents a true or false value | true, false, 1, 0 |
| [DateTime](https://ogp.me/#datetime) | A DateTime represents a temporal value composed of a date (year, month, day) and an optional time component (hours, minutes) | [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) |
| [Enum](https://ogp.me/#enum) | A type consisting of bounded set of constant string values (enumeration members). | A string value that is a member of the enumeration |
| [Float](https://ogp.me/#float) | A 64-bit signed floating point number | All literals that conform to the following formats: 1.234 -1.234 1.2e3 -1.2e3 7E-10 |
| [Integer](https://ogp.me/#integer) | A 32-bit signed integer. In many languages integers over 32-bits become floats, so we limit Open Graph protocol for easy multi-language use. | All literals that conform to the following formats: 1234 -123 |
| [String](https://ogp.me/#string) | A sequence of Unicode characters | All literals composed of Unicode characters with no escape characters |
| [URL](https://ogp.me/#url) | A sequence of Unicode characters that identify an Internet resource. | All valid URLs that utilize the https:// or https:// protocols |

You can discuss the Open Graph Protocol in [the Facebook group](https://www.facebook.com/groups/opengraph/) or on [the developer mailing list](https://groups.google.com/group/open-graph-protocol). It is currently being consumed by Facebook ([see their documentation](https://developers.facebook.com/docs/opengraph/)), Google ([see their documentation](https://developers.google.com/+/web/+1button/#plus-snippet)), and [mixi](https://developer.mixi.co.jp/en/connect/mixi_plugin/mixi_check/spec_mixi_check/). It is being published by IMDb, Microsoft, NHL, Posterous, Rotten Tomatoes, TIME, Yelp, and many many others.

The open source community has developed a number of parsers and publishing tools. Let the Facebook group know if you've built something awesome too!

- [Alternate WordPress OGP plugin](https://wordpress.org/plugins/wp-facebook-open-graph-protocol/) - A simple lightweight WordPress plugin which adds Open Graph metadata to WordPress powered sites.

The Open Graph protocol was originally created at Facebook and is inspired by [Dublin Core](https://en.wikipedia.org/wiki/Dublin_Core), [link-rel canonical](https://googlewebmastercentral.blogspot.com/2009/02/specify-your-canonical.html), [Microformats](https://microformats.org/), and [RDFa](https://en.wikipedia.org/wiki/RDFa). The specification described on this page is available under the [Open Web Foundation Agreement, Version 0.9](https://openwebfoundation.org/legal/the-0-9-agreements---necessary-claims). This website is [Open Source](https://github.com/facebook/open-graph-protocol).
