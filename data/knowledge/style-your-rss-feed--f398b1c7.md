---
title: "Style your RSS feed"
notion_id: f398b1c7-2ee1-4f24-a625-30b07790554d
notion_url: https://app.notion.com/p/Style-your-RSS-feed-f398b1c72ee14f24a62530b07790554d
last_edited: 2023-06-23T11:52:00.000Z
source_url: https://darekkay.com/blog/rss-styling/
tags: ["Web Development", "Blogging/Content Creation", "Article", "Guide", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[RSS](https://en.wikipedia.org/wiki/RSS) is not dead. It is not _mainstream_, but it's still a thriving protocol, especially among tech users. However, many people do not know what RSS feeds are or how to use them. Most browsers render RSS as raw XML files, which doesn't help users understand what it's all about:

In this post, I'll explain how to style RSS feeds and educate readers at the same time.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## XSL(T) to the rescue

This is how the [RSS feed for this blog](https://darekkay.com/atom.xml) looks like:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

To style a raw XML file in a browser, you have to provide styling information. You can do that by [attaching](https://www.w3.org/TR/2010/REC-xml-stylesheet-20101028/) an `xml-stylesheet` processing instruction within the RSS feed:

```plain text
<?xml version="1.0" encoding="utf-8"?>
<?xml-stylesheet href="/rss.xsl" type="text/xsl"?>
<feed xmlns="http://www.w3.org/2005/Atom"
      xmlns:media="http://search.yahoo.com/mrss/">
  ...
</feed>
```

The `href` attribute specifies a URL to a valid [XSL](https://www.w3.org/Style/XSL/WhatIsXSL.html) file. You can check out [mine](https://darekkay.com/assets/xsl/rss-style.xsl) for inspiration. Here's the gist:

```plain text
<?xml version="1.0" encoding="utf-8"?>
<xsl:stylesheet version="3.0"
                xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
                xmlns:atom="http://www.w3.org/2005/Atom">
  <xsl:output method="html" version="1.0" encoding="UTF-8" indent="yes"/>
  <xsl:template match="/">
  <html xmlns="http://www.w3.org/1999/xhtml" lang="en">
    <head>
      <title>
        RSS Feed | <xsl:value-of select="/atom:feed/atom:title"/>
      </title>
      <link rel="stylesheet" href="/assets/styles.css"/>
    </head>
    <body>
      <p>
        This is an RSS feed. Visit
        <a href="https://aboutfeeds.com">About Feeds</a>
        to learn more and get started. It’s free.
      </p>
      <h1>Recent blog posts</h1>
      <xsl:for-each select="/atom:feed/atom:entry">
        <a>
          <xsl:attribute name="href">
            <xsl:value-of select="atom:link/@href"/>
          </xsl:attribute>
          <xsl:value-of select="atom:title"/>
        </a>
        Last updated:
        <xsl:value-of select="substring(atom:updated, 0, 11)" />
      </xsl:for-each>
    </body>
    </html>
  </xsl:template>
</xsl:stylesheet>
```

This code is inspired by [pretty-feed-v3](https://github.com/genmon/aboutfeeds/blob/main/tools/pretty-feed-v3.xsl), but I've adopted it for the Atom specification.

The XSL file transforms the XML feed into a valid HTML document that any browser can render. You can use any information included in the XML file and put it into a new markup structure. You can even include content that is not part of the XML feed. A perfect use case is to include a note about what RSS feeds are and how to use them (as shown in the example above).

You can define inline CSS styles with a regular `<style>` element, but it's also possible to import external CSS files. I've imported my blog's main CSS stylesheet, so I didn't have to write any new code to make the RSS feed match my blog's design.

It's also possible to use [XSLT functions](https://www.w3.org/TR/xpath-functions-30/) to alter values. Here's how I truncate the timestamp to display only the date:

```plain text
<!-- Full timestamp -->
<xsl:value-of select="atom:updated" />

<!-- Date only -->
<xsl:value-of select="substring(atom:updated, 0, 11)" />
```

The [browser support](https://caniuse.com/?search=xslt) for XSL transformations is great; unsupported browsers will fall back to the default behavior (progressive enhancement).

## Examples

Here are some styled RSS feeds you can check for inspiration:

- [This blog](https://darekkay.com/atom.xml)
- [My photography website](https://photos.darekkay.com/atom.xml)
- [BBC News](https://feeds.bbci.co.uk/news/world/europe/rss.xml)
- [Hsiaoming Yang](https://lepture.com/feed.xml)
- [Dave Rupert](https://daverupert.com/atom.xml)
- [Oli Warner](https://thepcspy.com/feeds/full.xml)

The XSL files are public, so you can check how those pages have been implemented.

## Sitemap

You can apply XSL files to _any_ XML file. Another use case for styled XML files is a [website sitemap](https://www.sitemaps.org/). Although sitemaps are meant to be consumed by machines (e.g. crawlers), you can style them with little effort. See [my sitemap](https://darekkay.com/sitemap.xml) and the [respective XSL file](https://darekkay.com/assets/xsl/sitemap-style.xsl) as an example. Here's the gist:

```plain text
<?xml version="1.0" encoding="utf-8"?>
<xsl:stylesheet version="3.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
                xmlns:sitemap="http://www.sitemaps.org/schemas/sitemap/0.9">
  <xsl:output method="html" version="1.0" encoding="UTF-8" indent="yes"/>
  <xsl:template match="/">
    <html xmlns="http://www.w3.org/1999/xhtml" lang="en">
      <body>
        <h1>Sitemap</h1>
        <xsl:for-each select="/sitemap:urlset/sitemap:url">
          <a>
            <xsl:attribute name="href">
              <xsl:value-of select="sitemap:loc"/>
            </xsl:attribute>
            <xsl:value-of select="sitemap:loc"/>
          </a>
          Last updated:
          <xsl:value-of select="substring(sitemap:lastmod, 0, 11)" />
        </xsl:for-each>
      </body>
    </html>
  </xsl:template>
</xsl:stylesheet>
```

## Conclusion

I love RSS, and — according to [my stats](https://darekkay.com/blog/rss-subscriber-count) — many of my readers love it, too. With all the social media walled gardens, I hope for RSS to become more popular. By providing a custom XSL file, you can make your RSS feed look nice _and_ include some information about what RSS actually is.
