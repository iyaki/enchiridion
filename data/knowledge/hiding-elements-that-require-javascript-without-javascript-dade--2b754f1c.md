---
title: "Hiding elements that require JavaScript without JavaScript :: dade"
notion_id: 2b754f1c-7d23-814d-b731-c73483b4d40a
notion_url: https://app.notion.com/p/Hiding-elements-that-require-JavaScript-without-JavaScript-dade-2b754f1c7d23814db731c73483b4d40a
last_edited: 2025-11-26T14:48:00.000Z
source_url: https://0xda.de/blog/2025/04/hiding-elements-that-require-javascript-without-javascript/
tags: ["0xda Blog", "English", "Web Development", "HTML", "Javascript", "CSS", "Article"]
---
 
• [⭐ Using noscript and a d-js-required class.](https://0xda.de/blog/2025/04/hiding-elements-that-require-javascript-without-javascript/#star-using-noscript-and-a-d-js-required-class) 

- • [⭐ Using noscript and a d-js-required class.](https://0xda.de/blog/2025/04/hiding-elements-that-require-javascript-without-javascript/#star-using-noscript-and-a-d-js-required-class)

I’ve tried my best to make sure that this site works great (or at least reasonably well) even without JavaScript, but when JavaScript isn’t available, it can be a little clunky to hide things that do require it. With a mere 7 lines (or a one-liner if you’re nasty), you can easily hide elements that require JavaScript so that you don’t end up with broken functionality visible to users who have JavaScript off.

For context, I’m working on a small share button that I can add to my posts that makes it easy to share my posts however you’d prefer. So here’s an example of what that looks like right now without JavaScript:

 Screenshot snippet showing a share SVG icon next to the permalink for the post

![image](https://0xda.de/blog/2025/04/hiding-elements-that-require-javascript-without-javascript/js-required-button.5dbba86bfdb33883d6d4783f8af61f16.PNG)

While it doesn’t look _hideous_, the link isn’t necessary when JavaScript is enabled, and the SVG isn’t necessary when JavaScript is disabled. I wanted to just show the SVG as a button you can press when JavaScript is enabled, and just show the link to make it easy to copy/paste when JavaScript is disabled.

If you’re not already familiar, the latter part is easy with the `<noscript>` HTML tag. This tag tells browsers to render its contents only when JavaScript is disabled. So making the link only show up when JavaScript is disabled is as easy as wrapping it in a `<noscript>` tag.

But if you want to make an element only appear when JavaScript is enabled, it’s a little less clear cut. There’s not an `<onlyscript>` tag, and the `<script>` tag has a very different meaning than the `<noscript>` tag.

## Using JavaScript to indicate JavaScript is enabled

In my first approach with this, I thought maybe I could just add a line of JavaScript in the `<head>` tag that, if JavaScript is enabled, would run and add `js-enabled` to the class list on the `<html>` element.

This would work fine, then I could style things in a special way if JavaScript is enabled, by creating styles like this:

```plain text
.share-button {
    display: none;
}
.js-enabled .share-button {
    display: block;
}

```

But this felt clunky. It results in needing to create two CSS rules for each element, and for my goal of just hiding things when JavaScript isn’t enabled, that seems heavy handed.

This might be a useful solution if you needed to do significantly more elaborate styling only when JavaScript is enabled, but for my purpose, it wasn’t simple enough.

## Using noscript to override specific elements

The next solution I came to was to combine the use of the `<noscript>` tag with a `<style>` tag in the head of my site. Then I could just list out elements that I don’t want to show, and mark them explicitly as `display:none;`. They would still be present in the markup, but that’s fine, they shouldn’t get in anybody’s way.

In your site’s header, you can simply link all of your CSS as you normally would, and then add the following snippet to override the fields that you want to not show when JavaScript is disabled. In my example, I was hiding the `.theme-toggle` and the `.menu-trigger` classes. These are related to features that only work if JavaScript is enabled – the three way theme switcher and the hamburger menu dropdown on mobile.

```plain text
<noscript>
    <style>
        .theme-toggle {
            display: none;
        }
        .menu-trigger {
            display:none;
        }
    </style>
</noscript>

```

This is alright. It’s explicit. It relies on existing browser behaviors to do selective style overrides when JavaScript isn’t available, so you don’t have to style both the JS-enabled and JS-disabled cases explicitly. But as I was developing new progressive enhancement features, I didn’t really like this pattern anymore. The more enhanced the website becomes, the longer this list of overrides becomes, and you have to make sure you always keep it up to date if you change class names.

## ⭐ Using noscript and a d-js-required class.

Since we can append any number of classes to an element, I realized that it makes much more sense to make this noscript-style as simple as possible. We’ll simply create a single class, `d-js-required` to indicate that JavaScript needs to be enabled for us to display this element. We can update our `<noscript><style>` block with just a single class, and then update our existing elements to also use this class.

```plain text
<noscript>
    <style>
        .d-js-required {
            display: none;
        }
    </style>
</noscript>

```

So now we can have any number of simple elements that only get displayed when JavaScript is enabled, without writing another line of JavaScript, and without needing to maintain an ever-growing list of CSS overrides.

I hope this has given you some inspiration for creative solutions to solve progressive-enhancement problems. If you have JavaScript enabled, you can click the share button below to share this page! And if not, you can just copy the link down below to share this page!
