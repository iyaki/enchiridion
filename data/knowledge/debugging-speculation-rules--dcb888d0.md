---
title: "Debugging speculation rules"
notion_id: dcb888d0-30ca-4a27-8488-7764df82fdd0
notion_url: https://app.notion.com/p/Debugging-speculation-rules-dcb888d030ca4a2784887764df82fdd0
last_edited: 2023-09-01T15:45:00.000Z
source_url: https://developer.chrome.com/en/blog/debugging-speculation-rules
tags: ["English", "Web Development", "Guide", "Chrome Developers"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Speculation rules can be used to prefetch and prerender next page navigations as detailed in [the previous blog post](https://developer.chrome.com/blog/prerender-pages/). This can allow for much quicker—or even instant—page loads, greatly improving [Core Web Vitals](https://web.dev/vitals/) for these additional page navigations.

Debugging speculation rules can be tricky. This is particularly true for prerendered pages, as these pages are rendered in a separate renderer—kind of like a hidden background tab that replaces the current tab when activated. Therefore, the usual DevTools options cannot always be used to debug issues.

The Chrome team has been working hard to enhance DevTools support for speculation rules debugging. In this post, you'll see all the various ways of using these tools to understand a page's speculation rules, why they may not be working, and when developers can use the more familiar DevTools options—and when not.

## [#](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#explanation-of-pre-terms) Explanation of "pre-" terms

There's a lot of "pre-" terms that are easily confused, so let's start with an explanation of these:

- **Prefetch**: fetching a resource or document in advance to improve future performance. This post covers prefetching documents using the Speculation Rules API, rather than the similar, but older `<link rel="prefetch">` option often used for prefetching subresources.
- **Prerender**: this goes a step beyond prefetching and actually renders the whole page as if the user had navigated to it, but keeps it in a hidden background renderer process ready to be used if the user actually navigates there. Again, this document is concerned with the newer Speculation Rules API version of this, rather than the older `<link rel="prerender">` option (which [no longer does a full prerender](https://developer.chrome.com/blog/nostate-prefetch/)).
- **Navigational preloading**: the collective term for the new prefetch and prerender options triggered by speculation rules.
- **Preloading**: an overloaded term that can refer to a number of technologies and processes including `<link rel="preload">`, the [preload scanner](https://web.dev/preload-scanner/), and [service worker navigation preloads](https://web.dev/navigation-preload/). These items will not be covered here, but the term is included to clearly differentiate those from the "navigational preloading" term above.

For more details see [this preloading tech landscape](https://docs.google.com/document/d/1FNLyXW0Q5Fi5-kEOtjfOdmtoDxRtUR_wX4dl5Fz9c7o/edit) document.

## [#](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#speculation-rules-for-prefetch) Speculation rules for `prefetch`

Speculation rules can be used to prefetch the next navigation's document. For example, when inserting the following JSON into a page, `next.html` and `next2.html` will be prefetched:

```plain text
<script type="speculationrules">
  {
    "prefetch": [
      {
        "source": "list",
        "urls": ["next.html", "next2.html"]
      }
    ]
  }
</script>
```

Speculation rules are used for navigation prefetches have some advantages over the older `<link rel="prefetch">` syntax, such as a more expressive API and the results being stored in memory cache rather than the HTTP disk cache.

### [#](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#debugging-prefetch-speculation-rules) Debugging `prefetch` speculation rules

Prefetches triggered by speculation rules can be seen in the **Network** panel in the same way as other fetches:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The two requests highlighted in red are the prefetched resources, as can be seen by the **Type** column. These are fetched at `Lowest` priority as they are for future navigations and Chrome prioritizes the current page's resources.

Clicking on one of the rows also shows the `Sec-Purpose: prefetch` HTTP header, which is how these requests can be identified on the server side:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### [#](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#debugging-prefetch-with-the-preloading-panes) Debugging `prefetch` with the Preloading panes

A new **Preloading** section has been added in the **Application** panel of Chrome DevTools to help aid in debugging speculation rules:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

There are three panes available in this section:

- **Speculation Rules** which lists all the rule sets found on the current page.
- **Preloads** which lists all the prefetched and prerendered URLs from the rule sets.
- **This Page** which lists the prerendered status of the current page.

The **Speculation Rules** pane is shown above, and we can see this example page has a single set of speculation rules for prefetching 3 pages. Two of those prefetches succeeded and one failed. The icon besides the **Rule set** can be clicked to take you to the source of the rule set in the **Elements** panel. Alternatively, the **Status** link can be clicked to take you to the **Preloads** pane filtered to that ruleset.

The **Preloads** pane lists all the target URLs, along with the action (prefetch or prerender), which rule set they came from (as there may be multiple on a page), and the status of each preload:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Above the URLs, a drop down can be used to show URLs from all the rule sets, or only URLs from a particular rule set. Below that, all the URLs are listed. Clicking on a URL gives you more detailed information.

In this screenshot, we can see the failure reason for the `next3.html` page (which does not exist and therefore returns a 404, which is a non-2xx HTTP status code).

The final pane, **This Page**, shows a **Preloading status** report to show whether a prefetch or prerender was used for this page or not.

For a prefetched page, you should see a successful message when that page is navigated to:

### [#](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#unmatched-preloads) Unmatched preloads

When a navigation happens from a page with speculation rules that does not result in a preload, an additional section of the pane will show more details of why the URL did not match any of the preloaded URLs. This is useful for spotting typos in your speculation rules.

For example, here we navigated to `next4.html`, but only `next.html`, `next2.html`, or `next3.html` are prefetches, so we can see this doesn't quite match any of those three rules.

As this feedback is primarily intended for debugging speculation rules, this URL matching can look a little confusing when navigating to a completely different page that is not included in speculation rules, and is therefore not expected to be preloaded:

The Chrome team is [continuing to iterate on this feature](https://bugs.chromium.org/p/chromium/issues/detail?id=1473861) to improve the user interface in these cases.

The **Preloading** panes are very useful for debugging the speculation rules themselves, and finding any syntax errors in the JSON.

As for the prefetches themselves, the **Network** panel is likely a more familiar place. For the prefetch failure example, you can see the 404 for the prefetch here:

However, the **Preloading** panes become much more useful for prerendering speculation rules, which are covered next.

## [#](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#speculation-rules-for-prerender) Speculation rules for `prerender`

Prerender speculation rules follow the same syntax as prefetch speculation rules. For example:

```plain text
<script type="speculationrules">
  {
    "prerender": [
      {
        "source": "list",
        "urls": ["next.html", "next2.html"]
      }
    ]
  }
</script>
```

This rule set triggers a full load and render of the specified pages (subject to certain restrictions). This can provide an instant loading experience—albeit with extra resource costs.

Unlike prefetches however, these are not available to be seen in the **Network** panel, as these are fetched and rendered in a separate rendering process in Chrome. This makes the **Preloading** panes more important to debug prerender speculation rules.

### [#](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#debugging-prerender-with-the-preloading-panes) Debugging `prerender` with the Preloading panes

The same **Preloading** section introduced above can be used for prerender speculation rules as demonstrated with a similar demo page that attempts to prerender, instead of prefetching the three pages:

Here we see again that one of the 3 preloads failed to prerender, and developers can get the details per URL in the **Preloads** pane by clicking on the **2 Ready, 1 Failure** link.

We are currently experimenting with [document speculation rules](https://github.com/WICG/nav-speculation/blob/main/chrome-2023q1-experiment-overview.md#automatic-link-finding) (available as [an origin trial](https://developer.chrome.com/origintrials/#/view_trial/705939241590325249)), which—rather than listing a specific set of URLs—allows the browser to pick these up from same origin links on the page:

```plain text
<script type="speculationrules">
{
  "prerender": [
    {
      "source": "document",
      "where": {
        "and": [
          {"href_matches": "/*\\?*#*"},
          {"not": { "href_matches": "/not-safe-to-prerender/*\\?*#*"}}
        ]
      },
      "eagerness": "moderate"
    }
  ]
}
</script>
```

The above example selects all same origin links (including those with URL params or fragments—that is what the `\\?*#*` part means), except those beginning with `/not-safe-to-prerender `as prerender candidates.

It also sets the prerender `eagerness` to `moderate` which means the navigations are prerendered when the link is hovered, or clicked.

There are similar rules like this on [developer.chrome.com](https://developer.chrome.com/) itself, and using the new **Preloading** section on this site shows the usefulness of this new pane as all the eligible URLs the browser found on the page are listed:

The **Status** is **Not triggered** as the prerender process for these has not started. However, as we hover over the links, we see the status change as each URL is prerendered:

Chrome currently [limits a page to 10 prerenders](https://bugs.chromium.org/p/chromium/issues/detail?id=1464021), so after hovering over the 11th link, we see the failure reason for that URL:

### [#](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#debugging-prerender-with-the-other-devtools-panels) Debugging `prerender` with the other DevTools panels

Unlike prefetches, pages that have been prerendered will not show up in the current rendering processes in DevTools panels like the **Network** panel, because they are rendered in their own behind-the-scenes renderer.

However, it is now possible to switch the renderer used by the DevTools panels with the drop down menu in the top right drop down, or by selecting a URL in the top part of the panel and selecting **Inspect**:

This drop down (and the value selected) is shared across all the other panels too, such as the **Network** panel, where you can see the page being requested is the prerendered one:

Looking at the HTTP headers for these resources we can see they will all be set with the `Sec-Purpose: prefetch;prerender` header:

Or the **Elements** panel, where you can see the page contents, like in below screenshot where we see the `<h1>` element is for the prerendered page:

Or the **Console panel**, where you can see console logs emitted by the prerendered page:

It's even possible to see the **Preload** pane for a prerendered frame to see if it had speculation rules:

Note that speculation rules themselves are not actioned until the prerendered page is activated. These pages will always be in the `Not triggered` status.

### [#](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#debugging-speculation-rules-on-the-prerendered-page) Debugging speculation rules on the prerendered page

The previous sections discuss how to debug prerendered pages on the page which initiates the prerendering. However, it's also possible for the prerendered pages themselves to provide debugging information, either by making analytics calls or logging to the console (which is viewable as described above).

Additionally, once a prerendered page is activated by the user navigating to it, the **This Page** pane will show this status, and whether it was successfully prerendered or not. If it could not be prerender an explanation as to why that was the case is provided:

Additionally—[as is the case for prefetches](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#unmatched-preloads)—navigating from a page with speculation rules that did not match the current page will attempt to show you why the URLs did not match those covered in the previous page's speculation rules in the **This Page** pane:

## [#](https://developer.chrome.com/en/blog/debugging-speculation-rules/?utm_source=tldrwebdev#conclusion) Conclusion

In this post, we have shown the various ways developers can debug prefetch and prerender speculation rules. The team is continuing to work on tooling for speculation rules, and we would love to hear suggestions from the developers as to what other ways would be helpful for debugging this exciting new API. We encourage developers to raise an issue on the [Chrome issue tracker](https://bugs.chromium.org/p/chromium/issues/entry?template=Defect&components=Internals-Preload) for any feature requests or bugs spotted.


