---
title: "Understanding this single concept helped me write more reusable code: separating layout from content"
notion_id: 4f3927fc-5fd3-4bbb-9b55-0eb0e4bd2a8b
notion_url: https://app.notion.com/p/Understanding-this-single-concept-helped-me-write-more-reusable-code-separating-layout-from-content-4f3927fc5fd34bbb9b550eb0e4bd2a8b
last_edited: 2025-02-10T18:07:00.000Z
source_url: https://www.reddit.com/r/webdev/comments/1f1epje/understanding_this_single_concept_helped_me_write
tags: ["Article", "English", "System Design / Software Architecture"]
---
When developing UI components, separating layout from content can greatly enhance code reusability and maintainability.

Let’s illustrate this with a simple example – a button with supporting text:

In this case, the button has a fixed width of **15em** and a **margin-bottom** to provide spacing between itself and the text below.

# What's wrong with this approach?

While the above approach works for this particular scenario, the moment we want to reuse the button in a different location we run into the following issues:

**1. It has a fixed width**

Perhaps the design will call for a width other than **15em**, such as a full-width button.

**2. It has a fixed bottom-margin**

If the supporting text isn’t there for any reason, users will see additional spacing which will make the design less polished.

# How do we address this?

Here's where we separate the layout from the content. In this case the button is the content, and the width & margin would be considered the layout.

Let's update the code by creating an extra **div** to hold those properties:

Updated code to separate the content from the layout

By adding a **.layout** class we ensure that the button’s layout is handled by its parent container. This makes the button reusable across different contexts with varying layout requirements.

The button’s styling has also been simplified – we set its width to **100%**, and let the parent decide if it should be smaller. We can now use the same button in various places and it’ll display differently based on the parent element’s layout styling.

# There are two things to remember with this approach.

**1. Layout elements aren’t concerned about their children.**

Really, anything could live inside a layout element, whether that’s a button, text or an image. It’s main focus is alignment, spacing, content width, and other layout concerns.

**2. Content elements aren’t concerned where they’re placed.**

A content element should not care about itself in relation to other elements. Properties such as margins and width should be determined by its parent (the layout element).

# OK, it’s time for a design change

If the design requires 2 buttons to be added, that’s now much easier:

A new button added, along with its updated code

We only added another layout element with a **.btn-container** class, and added another button element inside it. Easy.

# Where this approach really shines

This gets even better when combined with Tailwind as you don’t need to create & name different classes for containers. Also, if you’re using a JS library (eg. Svelte, React, etc.) to build components, then you’re really flying:

Updated code using JS framework & Tailwind

# Conclusion

This simple example shows how separating layout from content can improve your code maintainability. Imagine using this approach on a larger project with many components and sections.

Like anything, there are times where this method doesn’t work, but thinking about these two different concerns will help you write reusable code. It’ll also mean you’ll hate yourself less when needing to dive into an old project if you’ve used this method… I’m speaking from first-hand experience.
