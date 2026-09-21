---
title: "URL as UI"
notion_id: 3eb4c29b-e263-4806-8d40-f969137279d4
notion_url: https://app.notion.com/p/URL-as-UI-3eb4c29be26348068d40f969137279d4
last_edited: 2023-11-27T19:56:00.000Z
source_url: https://www.nngroup.com/articles/url-as-ui/
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

**Summary:**

**Topics:**

- [Web Usability](https://www.nngroup.com/topic/web-usability/)

Share this article:

The URL will continue to be part of the Web user interface for several more years, so a usable site requires:

- a ** domain name ** that is easy to remember and easy to spell
- **short ** URLs
- **easy-to-type ** URLs
- URLs that visualize the ** site structure**
- URLs that are "**hackable**" to allow users to move to higher levels of the information architecture by hacking off the end of the URL
- **persistent URLs ** that don't change

In principle, users should not need to know about URLs which are a machine-level addressing scheme. In practice, users often go to websites or individual pages through mechanisms that involve ** exposure to raw URLs**:

- people ** guess the domain name ** of sites they have not visited before: if possible, secure the name of your company and main brands as domain names
- even when people have been to a site before, they will often try to guess or remember the site name instead of using a bookmark or history list: have ** memorable domain names ** that are easy to spell
- the social interface to the Web relies on ** email ** when users want to recommend Web pages to each other, and email is the second-most common way users get to new sites (search engines being the most common): make sure that all URLs on your site are ** less than 78 characters ** long so that they will not wrap across a line feed
- shorter URLs are better since people often type them manually
- **do not use MiXeD case ** text in URLs since people can't remember the difference between upper-case and lower-case characters: all-lowercase URLs are usually preferred (domain names are less of a problem since they are case-insensitive - usability would increase if webservers would ignore case in resolving URLs)
- use a ** spelling-checking ** webserver to minimize the damage caused by the inevitable typos

## Persistent URLs Attract Links

Links from other websites are the third-most common way people find sites (after search engines and email recommendations), so build your site to make it easy to attract inbound links:

- [Linkrot](https://www.nngroup.com/articles/fighting-linkrot/) equals lost business: make sure all URLs live forever and continue to point to relevant pages.
- **Do not move pages around ** but keep them at the same URL: it is very annoying for authors of other sites when their links either stop working or turn into pointers to something different because the original page has been moved and replaced by something new. There can be reasons to reserve a special URL for the _ current _ edition of a column or other special content, but the article should be stored at a permanent URL from the start and this URL should be listed on the page that is accessed through the temporary or varying URL.

## Should Domains End in .com?

The most frequently asked question on my recent lecture tour to Iceland was whether it is better to get a domain ending in .com or to use the country's own domain (.is).

Unfortunately, many users have been trained to view ".com" as the standard ending for commercial websites: this is an artifact of the early American dominance on the Web and of the completion algorithm in several popular browsers which automatically add .com to any name. Because of this situation, my advice is:

- for a site that ** uses English ** _ and _ is clearly ** world-wide ** in its appeal and user base: get a .com domain
- for a site that uses ** any other language**: use the appropriate country domain ending
- for a site that has mainly ** local appeal**, covers mainly local issues, or sells mainly local products: use the country domain, no matter what language is used on the site

I recommend use of the local domain for local sites because it is misleading to use the "international" domain ending .com for such sites. As ecommerce and other uses of the Web grow around the world, people will start to expect local domains for local sites and they will not think to type .com for local service. Since the ability to provide great local service is a major selling point, sites are better off by staying with their own country's domain name unless they deliberately want to be seen as disembodied cyberspace entities.

## Domain Names May Die

It is likely that domain names only have 3–5 years left as a major way of finding sites on the Web. In the long term, it is not appropriate to require unique words to identify every single entity in the world. That's not how human language works.

The proposals to open up ** new top-level domains ** like .shop are a poor solution from a usability perspective since there is no easy way to remember which domain ending is associated with which site. The only new TLD that's useful is .sex which would allow very simple ways of filtering content that's undesired (or desired, as the case may be).

New addressing schemes are likely to be introduced with better support for ambiguity and the ability to find things without knowing the exact spelling. Search engines and directories are an early attempt, but we can surely do better.

Because of the [ conservatism of Web users](https://www.nngroup.com/articles/the-increasing-conservatism-of-web-users/), we will have to cater to old browsers, old software, and old habits for many years, so good domain names will continue to be important for many years. The remaining useful life of a domain name may be as much as ten years.

See [ reader comments](https://www.nngroup.com/articles/readers-comments-on-url-as-ui/) on this Alertbox and a sidebar on [ compound domain names](https://www.nngroup.com/articles/compound-domain-names/).

**Update added 2005:**

Domain names are still in use, and it looks like they may have several more years to go. But my prediction that they would become a less important way to find sites has come true: it is now much more common for users to type a company name, a site name, or even a domain name into a search engine than it is for them to guess the URL. Even so, enough users continue to enter their destination directly into the browser's URL field that it's still important to have an easily guessable domain name and easily typable URLs.

URLs also impact several of the [ design guidelines for confirmation email](https://www.nngroup.com/reports/ecommerce-transactional-email-confirmation-message/), because these messages often need to refer the user back to the originating website.

**Update added 2007:**

Edward Cutrell and Zhiwei Guan from Microsoft Research have conducted an [ eyetracking study of search engine use](http://research.microsoft.com/apps/pubs/default.aspx?id=70395) (warning: PDF) that found that people spend ** 24% of their gaze time looking at the URLs ** in the search results.

MSR used Microsoft's own search engine (fair enough), but their results match what we found in [ our eyetracking research](https://www.nngroup.com/topic/eyetracking/) which included the current market leader as well as the #2 search engine in addition to MSN.

Users have evolved a firm model of search behavior which they apply across search engines, which is why it's probably a lost cause to make a non-standard search user interface.

We found that searchers are particularly interested in the URL when they are assessing the credibility of a destination. If the URL looks like garbage, people are less likely to click on that search hit. On the other hand, if the URL looks like the page will address the user's question, they are more likely to click.
