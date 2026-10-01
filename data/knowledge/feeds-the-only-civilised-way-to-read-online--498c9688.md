---
title: "Feeds: The Only Civilised Way to Read Online"
notion_id: 498c9688-7e61-44a2-83bc-cdd300fbe874
notion_url: https://app.notion.com/p/Feeds-The-Only-Civilised-Way-to-Read-Online-498c96887e6144a283bccdd300fbe874
last_edited: 2023-01-13T19:10:00.000Z
source_url: https://felixcrux.com/blog/feeds-the-only-civilised-way-to-read-online
tags: ["Article", "English", "Learning", "Communication"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

How do you follow your favourite online creators and their blogs, articles, videos, podcasts, comics, or what have you? How do you choose what to read when you have a few minutes to kill on the bus, or when you want to get caught up in the morning with a cup of coffee?

Most people hand this choice over to social media, inviting along the whole associated host of problems like clickbait; outrage amplification; snooping targeted advertising; radicalising rabbit-holes; echo-chambers and filter bubbles; algorithms choosing what to show you based on “engagement” rather than what you’d want for yourself; and on and on.

There’s a better way — and there has been for decades! Amazingly, it seems underused even within tech circles, and almost completely unknown to the general public. It’s super easy to use, actually _more_ convenient than social media apps, and leaves you in complete control of what you see.

I’m talking, of course, about RSS/Atom web feeds, and I contend that they are not only a better alternative, but in fact I’d go so far as to say that a feed reader is the only tolerable and civilised way to read online! The system works really well and more in line with what (I think) most people actually want; it minimizes the use of harmful social media platforms; and it helps foster a more vibrant, independent, creative, and non-commerical Web. So drop your non-chronological algorithmically-obscured sponsored timeline, and let’s have a whirlwind overview of what feeds are and how to use them!

### What are feeds and feed readers?

The core of the problem is this: You might have hundreds of people whose work you’re potentially interested in, but you can’t realistically regularly check whether each one has published something new recently. The solution (first cooked up around 1999!) is a _web feed_ or _news feed_ (sometimes _RSS feed_ or _Atom feed_, [for uninteresting technical reasons](https://felixcrux.com/blog/feeds-the-only-civilised-way-to-read-online#appendix-names)), combined with a program called a _feed reader_ (or a _news reader_ or _aggregator_). I strongly suspect that the wide range of inconsistent names has been a big hindrance to these tools being widely adopted, but in practice which name you use doesn’t matter: I’m just going to call them _feeds_ and _feed readers_ from now on.

A feed is a machine-readable version of new posts on a site. For a blog, it might contain, say, the last 10 blog posts, much like the human-readable “front page” of the web version would. A feed reader is a program that monitors all of the feeds from all of the sites that you’re interested in, and aggregates new updates in one place for you to read at your leisure.

So, if there’s a blog or podcast or something else you’d like to follow, you simply subscribe to it in your feed reader, and now all new updates will be automagically conveyed to you as soon as they come out. Your feed reader keeps track of what you’ve not yet read, and displays all posts in a simple chronological timeline, so you never miss anything. You can catch up whenever you want; there are no algorithms boosting controversial or sponsored nonsense; and no ads.

Feed readers would traditionally have been desktop programs that downloaded all the feeds you were monitoring onto your computer. That model still works just fine (and in fact may already be available on your computer thanks to it being a built-in feature of the [Thunderbird](https://support.mozilla.org/en-US/kb/how-subscribe-news-feeds-and-blogs) and [Outlook](https://support.microsoft.com/en-us/office/subscribe-to-an-rss-feed-73c6e717-7815-4594-98e5-81fa369e951c) email clients), but nowadays most people prefer to use online services so that their read/unread status for individual posts can be synchronized between their desktop and their phone. Articles can then be read in the browser or in a smartphone app, and you’ll always be up-to-date with what you’ve already seen and what you’ve not yet read.

Which feed reader you choose doesn’t matter much. Some popular web app options include [NewsBlur](https://newsblur.com/), [Feedly](https://feedly.com/), [Feedbin](https://feedbin.com/) and [Inoreader](https://www.inoreader.com/), or if you’re up for self-hosting, [Tiny Tiny RSS](https://tt-rss.org/). Desktop applications include [NetNewsWire](https://netnewswire.com/), [FeedReader](https://jangernert.github.io/FeedReader/), [Newsboat](https://newsboat.org/), and [Fluent Reader](https://hyliu.me/fluent-reader/). There are many more; these are just a sample. All of the (non self-hosted) web services are paid products — that’s the tradeoff for not being profiled and advertised at. Fortunately, most have free tiers or free trial periods, so you can try them all out and see which you like best. Many try to differentiate themselves by offering extra features such as optional filtering or even machine-learning-based prioritization of your feeds, but in my case that’s what I’m trying to get _away_ from so I just turn those features off and don’t use them. I personally use NewsBlur, which has worked just fine (along with their mobile app) for years.

So, you pick a feed reader to try out, and set up an account. Then, you find the feeds published by sites you’d like to follow, and subscribe to them in your feed reader. You can then check in whenever you’d like, and any new posts will be there waiting for you in your reader. No more random outrage-bait or advertising, just a chronological feed of what _you_ chose to follow.

### But does anybody actually publish feeds?

Surprisingly, yes, almost every blog or publishing platform out there does in fact publish a feed (often automatically, without the author needing to do anything special to set it up). This includes publishing systems like Medium, Substack, Wordpress, and Blogger, as well as smaller self-hosted site generators. You can follow feeds from an astonishing variety of sources, including YouTube channels, [Environment Canada weather warnings](https://www.canada.ca/en/environment-climate-change/services/weather-general-tools-resources/weatheroffice-online-services/data-services.html), [CNN](https://www.cnn.com/services/rss/) or [the BBC](https://www.bbc.co.uk/news/10628494), [GitHub releases](https://github.com/felixc/rexiv2/releases.atom), [status pages for online services](https://status.circleci.com/history.atom), and of course, the myriad of small blogs out there on the web. The entire podcasting ecosystem is secretly built on top of feeds — that’s how new episodes get delivered to you! And, naturally, [my feed is right here](https://felixcrux.com/blog/rss.xml).

To find a feed for a particular site, you can almost always just put in the address of the website into your feed reader, and it will auto-discover the feed. If not, first look for the ubiquitous orange and white feed icon: , perhaps alongside words such as “feed”, “RSS”, or “Atom”. There will be a link, which you can put into your feed reader to subscribe. Web browsers used to highlight the existence of feeds with an icon in the address bar, and you can restore that functionality to Firefox with [this addon](https://addons.mozilla.org/en-US/firefox/addon/want-my-rss/).

### What about discovering new authors?

This is where some forms of social media perhaps have a defensible role to play. Potential sources for content discovery can be thought of as lying along a spectrum that at one end has high quality content but low frequency of discovering new sources, and at the other end has the incredible torrent of material that can be found online, which is usually of low value.

The most reliable, but slowest, way of finding new feeds to follow is to wait for friends, family, or coworkers to share links to things they found interesting. Rather than just read them and move on, take a minute to poke around the site and see if you think you’d like to see more from that creator. If so, check whether they have a feed, and subscribe to it. If it turns out they never post again, there’s no harm done. If it turns out their future posts aren’t very interesting, that’s ok: just unsubscribe.

A higher-volume but perhaps more inconsistent source of feeds to follow might be topic/community-focused public aggregator sites. These are a form of social media, but they center on particular subjects or communities, rather than whatever anyone would like to share. For example, technologists often follow [Hacker News](https://news.ycombinator.com/), [Lobsters](https://lobste.rs/), or [the Programming subreddit](https://www.reddit.com/r/programming/). Like with links shared with you directly, if you find something here you like, poke around the author’s homepage and consider following them for more. As a bonus, if you care about accruing karma/points/reputation on these social sites, this gives you an edge in finding new content to post to them!

To get a head start, here are two exported data files of the [tech-focused](https://felixcrux.com/files/feeds-tech.opml) and [general](https://felixcrux.com/files/feeds-general.opml) feeds I follow. You should be able to import them into your feed reader of choice — hopefully you find something you like! Note though that I don’t necessarily _endorse_ any of these sites or authors; they may just have published one good thing years ago and I have no idea what they’ve been up to since.

### Ultimately, why should I bother with this?

Like I said at the start, for me it comes down to three reasons:

Firstly, it just works really well — better than alternative options. By that I simply mean that a feed reader really is the best way to keep up with blogs and posts! A favourite part of my day is having my morning coffee and catching up on feeds. I can pick them up when I have a few minutes to spare, and what I’ve read or not read is automatically synced between my devices. Easy, pleasant, simple, convenient, and I’m in control of what I read.

The second reason is maybe not universally accepted by everyone, but in my opinion, it’s a healthier alternative to social media. I feel those platforms should be used only lightly and warily, if at all. Their very structure inherently promotes controversy and pile-ons rather than discussion; and fake “influencer” distortions over reality. While they may be a reasonable way to keep up with a small group of friends and family, opening the door to strangers being algorithmically boosted into your timeline, and the associated advertising and psychological manipulation (whether intentional or merely a side-effect) doesn’t feel healthy to me. Following blogs directly is a way to minimize or outright bypass these platforms and still benefit from the wide range of content out there on the web.

And finally, the third reason is that I hope that directly following blogs and creators promotes the older, more independent, non-commercial version of what the Web could be. Rather than merely a more interactive cable TV channel controlled by big platforms, we can encourage anyone who has an interesting story to tell to share it online, and they can build up a dedicated audience independently, without having to pander and clickbait in the hope of algorithmically going viral. I want to support _that_ version of the Web by choosing what to read, following small blogs, and enjoying reading long blog posts written simply because the author had something to say.

### Appendix: Why all the inconsistent names?

An unfortunate complication when evangelizing or adopting feeds is the proliferation of opaque names that all mean essentially the same thing. Why is it like this?

The _concept_ of what we’re trying to publish is a “feed”. The original standardized _format_ for publishing feeds was “RSS”: “Really Simple Syndication”. Several years later there was a schism in the community and a separate group produced a new standard format called “Atom”. Today, the choice of format is largely irrelevant, and essentially all feed readers support both. However, the _name_ “RSS” caught on reasonably well, and so sometimes people will call the generic concept “RSS feeds” — confusingly even when the actual format in use might be Atom and not RSS!

However, ultimately, none of these details matter in the slightest to users. The only reason to know that RSS, Atom, news feed, and web feed are all possible names is so you can find the link to the feed to plonk it into your feed reader, which will happily work with any format or name.
