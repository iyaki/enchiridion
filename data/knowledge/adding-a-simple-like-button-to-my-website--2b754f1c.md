---
title: "Adding a Simple Like Button to My Website"
notion_id: 2b754f1c-7d23-816f-91f9-d20234e55e53
notion_url: https://app.notion.com/p/Adding-a-Simple-Like-Button-to-My-Website-2b754f1c7d23816f91f9d20234e55e53
last_edited: 2025-11-26T18:09:00.000Z
source_url: https://joeldare.com/adding-a-simple-like-button-to-my-website
tags: ["Web Development", "HTML", "CSS", "Productivity", "Article", "joeldare.com", "English"]
---
I’ve been reading _The Art and Business of Online Writing_ by _Nicolas Cole_ and it prompted me to try to get some feedback on my writing. So, I worked on adding a “like” button to my website.

I like to keep my website as small, light, and private as I can, so I don’t want to embed anything. What I decided was to implemented this as a simple hyperlink to an auto-saving form. Here’s an example of the link I used:

I created the form itself on [Tally.so](https://tally.so/) and then linked the button to that form. I set the form up to save automatically, even if the user doesn’t hit the submit button. My hope was that this would make the “like” a single click action. Unfortunately, Tally doesn’t save the form unless the user does _something_ on the page. Setting a hidden field does not trigger the auto save like I had hoped.

I added a hidden field to capture the url the user came from. I also added a field for additional comments if the user wants to add any.

When a user clicks on the like button they go to a Tally.so form and the URL and their like/dislike choice is automatically selected. They can, optionally, enter some comments and hit submit to record the rating.

I also tried adding a hidden field and a visible field and setting the visible field to the value passed to the hidden field. That also didn’t trigger an automatic save.

If you have experience with Tally and know how to make this more automatic, hit one of the buttons below and let me know.

Want to see my personal updates, tech projects, and business ideas? Join the mailing list.

JoelDare.com © Dare Companies Dotcom LLC
