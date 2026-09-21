---
title: "4 Rules for Intuitive UX"
notion_id: fa6ca36d-70cd-4257-bceb-ff52557a89ef
notion_url: https://app.notion.com/p/4-Rules-for-Intuitive-UX-fa6ca36d70cd4257bcebff52557a89ef
last_edited: 2024-09-10T19:54:00.000Z
source_url: https://www.learnui.design/blog/4-rules-intuitive-ux.html
tags: ["Article", "LEARN UI DESIGN", "English", "UI/UX", "Product Management"]
---
This is my advice on improving the UX of your designs WITHOUT hours of user research sessions, paper prototyping playtime, or any other trendy UX buzzwords 😎

(Seriously, search this page “design thinking”. 0 results. _Nailed it_!)

Who’s this article for? Easy:

- **Developers**. You created your own app, but every time someone downloads it, they struggle to use it. And if someone’s actually _told_ you this, then it’s _really_ bad.
- **Graphic designers**. Looking to make the transition into digital, but trying to learn UX by reading articles online is… a very painful way to die 😬
- **PMs**. Your job is already like 25% UX designer. Would be nice to level up those skills.
- **And the hustlers**. Anyone working on digital side projects nights/weekends. This one’s for you too 🍻

If you’re already a UX designer, I don’t expect this article to go over super well with you. I’m basically skipping over entire chunks of our field in favor of focusing entirely on **the single most lacking skill in aspiring UX designers** (or UX-adjacent folks who find themselves designing screens).

I call it “**speaking interface**”.

When I started as a professional UX designer, I was _shocked_ how many times my clients would hand me the initial wireframes (or the living, breathing, in-browser MVP) and there’d be _**completely obvious UX mistakes**_ all over them. I’m not talking about things you need hours of research and A/B testing to discover. I’m talking, like, _dead simple_ mistakes.

For lack of a better example:

> Somewhere out there, there's a team that knows HTML, but doesn't know the difference between a radio button and a checkbox. pic.twitter.com/VBwk8Jxekd

[May 24, 2017](https://twitter.com/erikdkennedy/status/867412999115464705?ref_src=twsrc%5Etfw)

Now my clients weren’t _this_ bad, but _look _- you don’t need to be Bret Victor to understand that if you can only select ONE thing from a list, you need RADIO BUTTONS, not checkboxes. To understand _that_, you just need to be able to _**speak interface**_. And that’s the craziest thing to me. Interface fluency is something _anyone_ can achieve.

Frankly, you just need the **presence of mind** to (A) pause every single time you’re confused or frustrated by some app, (B) verbalize _what_ about the interface makes you confused/frustrated/etc., and then (C) figure out how you could avoid that specific snafu that in your own designs.

Rinse and repeat that non-stop and you’ll be a pro in no time.

What I want to talk about today is **four little rules** that will help eliminate these pain points in your own designs. They’re the heuristics that are a level or two deeper than “use radio buttons if the user can only select one thing”. But, if you can remember to obey the things in this checklist, you’ll be that much closer to creating designs that your users can use easily right off the bat, freeing up your time for other, more important things.

(That’s when the _ooother_ UX designers can lecture you on the newest academic user research methodologies 😉)

Here’s what we’ll be covering:

1. [Obey the Law of Locality](https://www.learnui.design/blog/4-rules-intuitive-ux.html#1-obey-the-law-of-locality)
2. [ABD: Anything But Dropdowns](https://www.learnui.design/blog/4-rules-intuitive-ux.html#2-abd-anything-but-dropdowns)
3. [Pass the Squint Test](https://www.learnui.design/blog/4-rules-intuitive-ux.html#3-pass-the-squint-test)
4. [Teach by example](https://www.learnui.design/blog/4-rules-intuitive-ux.html#4-teach-by-example)

Any questions? Let’s dive right in.

## 1. Obey the Law of Locality

_Put interface elements where they effect change._

All else being equal, you should put the elements in your interface _near where they effect change_. This is because, when a user wants to make a change to the system, **they will unwittingly **_**glance at where that change will happen**_.

So, let’s say you have a list of _things_. Where do you put the “ADD A NEW THING” button?

![image](https://www.learnui.design/img/4-rules/lol-where.png)

law of locality in playlist illustration

Q: Well, where does the _change_ happen?

A: At the end of the list.

Great, put the button at the end of the list.

WAIT. You’d think this would be pretty simple. But there’s a _temptation_.

**The temptation is to just put it **_**where we have space for it**_**.**

For instance, if you have a menu, maybe you’d think “We have a menu! Why not just put it in the menu!?”

![image](https://www.learnui.design/img/4-rules/lol-menu-no.png)

the law of locality violated in a music UI

The answer is, of course, because **users won’t look for it there**.

(And the ultimate answer is that having a place where “we just put things” will ultimately render your app an unusable mess that people will abandon the first chance they see a half-viable alternative)

Don’t think I’m joking. Have you ever noted _this_ interface?

![image](https://www.learnui.design/img/4-rules/lol-evernote.png)

the law of locality violated in evernote's interface

An equally-bad/common alternative is to take something you saw A Respected Tech Company do and steal it for your own app, without any thought as to _if it makes sense for you_. “We need an ‘Add’ button? I’ve seen one of those. Hold my beer!”

![image](https://www.learnui.design/img/4-rules/lol-fab-no.png)

the law of locality violated with a floating action button

Look. Another button in a place users _**will never look for it**_. To compound things, users will suspect _this_ button actually adds a new whatever-is-currently-displayed-on-the-big-blank-white-space. Because that’s where the control _is_.

Your users _want_ you to follow the Law of Locality.

So, now that we know it, let’s use it.

![image](https://www.learnui.design/img/4-rules/lol-bottom-of-list.png)

the law of locality in a list of music playlists

Bam.

But maybe you’re a born UX designer and **you always visualize what happens when there’s 1000 items instead of 5** and you realize: _there’s still an issue here_. If the user creates a TON of playlists, this button will disappear hundreds of pixels offscreen!

So maybe you could anchor the button _near the bottom of the list_, but have it _always be visible_, no matter how many hundreds of playlists the user has created.

![image](https://www.learnui.design/img/4-rules/lol-spotify-new-bottom.png)

For bonus points, (1) use the inline button UNTIL it's about to go offscreen, and at that point switch to the anchored solution and (2) make it more visible than Spotify's button, which took me months to notice while I haplessly right-clicked individual songs to add them to my playlists!

Brilliant! And this is what Spotify has done.

Another possibility is to say “Hey, we can’t _reliably and consistently_ show the button at the bottom of the list. Where’s the _nearest logical place_ to put it?”

And the answer is, (I think pretty obviously) _the top of the list_.

![image](https://www.learnui.design/img/4-rules/lol-spotify-new-top.png)

I wish.

Sacrebleu! This is actually just what Spotify-competitor Rdio did, before they were acqui-shut-down by Pandora.

![image](https://www.learnui.design/img/4-rules/lol-rdio.png)

Reconstructed from memory (like all reality, if you think about it)

The lesson here is clear. Never sell your company, and always always obey the Law of Locality.

(There are actually 3 laws of locality, and “Put UI elements where they effect change” is only the first. If you’re interested, read more [here](https://www.learnui.design/blog/the-3-laws-of-locality.html))

Next!

## 2. ABD: Anything but Dropdowns

_Any time you feel tempted to use a dropdown, ask yourself if one of these 12 controls is better instead._

One **non-obvious lesson of UX design** is that **dropdowns** are pretty much the **worst control**.

![image](https://www.learnui.design/img/4-rules/dropdown-hell.png)

Welcome to hell!

They’re not _always_ bad, but you’re working against the following:

- Dropdowns take **too many clicks/taps**. One to open, a few more to scroll around to the right option (on mobile), another to select the right option, and (on mobile) another to close. (Compare to the _single click use-cases_ of many of the options listed below)
- Dropdowns **don’t show you the options**! You have to click into them to see the possible values, and on mobile, you can often only see a couple at a time.
- Long dropdowns are **ridiculous to navigate**. A country dropdown for an app used worldwide could have 195+ countries. At some point, almost _any other method_ of asking a user their country would be quicker than having them scroll through a dropdown (“Smoke signals?” AGCKKHKGH).

This is pretty straightforward, so let’s just cover some examples for the various major cases of dropdown replacement.

### If you’re choosing between 2 options…

We already have some fantastic options for allowing users to choose 1 of 2 things, all of which (A) show the options right away and (B) require fewer taps/clicks.

For questions to which there is no “default” answer, and either might be picked with roughly equal frequency, try a **segmented button**.

![image](https://www.learnui.design/img/4-rules/dropdown-two-segment.png)

segmented button instead of dropdown control

If there is a “default state” that corresponds to “Off”, try a **checkbox**. A checkbox is also good for settings that don’t effect change _until the user presses Save or Submit_.

![image](https://www.learnui.design/img/4-rules/dropdown-two-checkbox.png)

checkbox instead of dropdown control

Similar to the checkbox is the **switch**, which is good for changes that should apply _immediately_.

![image](https://www.learnui.design/img/4-rules/dropdown-two-switch.png)

switch instead of dropdown control

Checkboxes and switches _only_ make sense when there are two options. However, the following controls make sense for 2 to roughly 5 options, so you might try some of the following instead.

### If you’re choosing between 2–5 options…

We covered **segmented buttons** above (and they apply here too) but it’s worth mentioning that when there are more options, **vertical segmented buttons** allow even more flexibility of answer length.

![image](https://www.learnui.design/img/4-rules/dropdown-few-segment.png)

vertical segmented button instead of dropdown control

**Radio buttons** are similar, but particularly useful if you need to display a couple sub-elements for each choice.

![image](https://www.learnui.design/img/4-rules/dropdown-few-radio.png)

radio button instead of dropdown control

For detailed displays of just a few choices, **cards** are where it’s at.

![image](https://www.learnui.design/img/4-rules/dropdown-few-cards.png)

cards instead of dropdown control

One trick I like is displaying **visual options** literally.

![image](https://www.learnui.design/img/4-rules/dropdown-tesla-visual.png)

Tesla likes it too, apparently.

### If you’re choosing between many options…

When there are enough options that scrolling through them is annoying, consider a **typeahead** control. It’s like a search bar that shows top matching results as you type.

![image](https://www.learnui.design/img/4-rules/dropdown-many-typeahead.png)

typeahead control instead of dropdown control

### If you’re choosing a date…

Picking a date from dropdowns is the _worst_. If I ever do this, then I’ve _really_ failed as a UX designer.

![image](https://www.learnui.design/img/4-rules/dropdown-dates-dont.png)

don't use dropdown controls for choosing dates

But what do you use instead? Well, it depends. First question: what _type of date are you picking_?

1. **Poisson dates**. Dates _most likely to be in the near future_, tapering off as you go farther into the future (or nearer to the present), e.g. date of an appointment you’re scheduling, date of a flight you’re purchasing
2. **High-variability dates**. Dates that have a _similar probability of being anywhere in a wide range of time_, e.g. date of birth, day-and-month of your birthday

(Yes, I named “Poisson dates” after the [mathematical distribution](https://en.wikipedia.org/wiki/Poisson_distribution) 🤓)

For different types of date-picking, you should use different controls.

![image](https://www.learnui.design/img/4-rules/dropdown-dates-chart.png)

chart of poisson vs. wide-range dates

For Poisson dates, you want to make it DEAD SIMPLE to pick dates in the most common range (e.g. for scheduling an appointment, it might be the next, say, 14 days). It’s perfectly OK if picking dates outside of that range is a little tougher.

A **calendar control** fits the bill rather well for Poisson dates. If you know the date to-be-picked is most likely in the next 2–4 weeks, you’re golden.

![image](https://www.learnui.design/img/4-rules/dropdown-date-calendar.png)

calendar control instead of dropdown control

Rather creatively, [Google Flights](https://www.google.com/flights) _defaults_ to you selecting a flight roughly 2 weeks in the future, which is perhaps an opportunity for confusion (“I didn’t choose this!”), but probably a better date to default to, and closer to the hump in the Poisson curve.

![image](https://www.learnui.design/img/4-rules/dropdown-google-flights.png)

Google Flights defaults to hump in Poisson distribution fo flight dates

**Date text inputs** are probably the best option for high-variability dates, where (A) there’s no reason to favor any date over another, meaning (B) all options will be equally difficult to select.

![image](https://www.learnui.design/img/4-rules/dropdown-date-text.png)

Remember, input[type=date] is your friend… on desktop, at least

### If you’re choosing a number…

Numbers come in all kinds of flavors, but you’re most likely to be tempted to use dropdowns when you’re dealing with _**counts **_- e.g. the number of _tickets_, the number of _people_, the number of _rooms_, etc.

How often do you need 1 ticket? _Plenty_.

How often do you need 10 tickets? _Not so much_.

How often do you need 10,000 tickets? _Is this some kind of cruel joke?_

For counts of things, you’re also dealing with Poisson distributions, and should use a control that biases towards lower numbers - like a **stepper**.

![image](https://www.learnui.design/img/4-rules/dropdown-number-stepper.png)

stepper control instead of dropdown control

For wide-range numbers (like, say, SSNs), you weren’t going to use a dropdown anyways… _I hope_.

### So can I ever use a dropdown?

Sure.

Remember, they work OK when…

- Users **rarely need to change the default** value
- There are **very few options** - e.g. only 3 will be visible on the default iOS control
- The user is **not on mobile** (whereby many of these problems are mitigated)

The particularly observant among you may have noticed that the Google Flights interface I lauded above actually has **three prominent dropdowns**!

![image](https://www.learnui.design/img/4-rules/dropdown-google-flights.gif)

Brilliant detail: on mobile, the 'Economy' dropdown is removed.

They actually do a great job with this. The potential usability issues are swiftly mitigated with:

- **Custom controls** that show all options on tap (including on mobile) – and replace 4 dropdowns (for Adults, Children, and Seated Infants and Lap Infants) with **4 steppers in a single dropdown**.
- **Removing** the “Economy” dropdown on mobile
- **Few options** and **smart defaults** for each control

If you want to print this section out and stick it on your wall, I’ve created a [printable cheatsheet](https://www.learnui.design/extras/selection-controls-ux-checklist.html) of dropdown replacements.

Anyhow. Let’s move on.

## 3. Pass the Squint Test

_If you squint your eyes, the Most Important Thing should catch your eye first – and the least important elements should catch your eye last._

Pop quiz: what does a user need to do to _use this page_?

(NB: I’ve blurred it out so you have to go by gut instinct, but it’s a data entry form, to give you a hint)

![image](https://www.learnui.design/img/4-rules/squint-ak1-blur.png)

blurred out version of a train ticketing ui

My best guess is **two things**:

1. Check any **applicable checkboxes** (??) in the yellow area
2. Press the **blue “Submit” button**

Did you guess the same?

Wrong and wrong.

![image](https://www.learnui.design/img/4-rules/squint-ak1.png)

train ticketing ui violates squint test

1. The “checkboxes” are actually very small numerical text inputs. (If you already read [Anything But Dropdowns](https://www.learnui.design/blog/4-rules-intuitive-ux.html#2-abd-anything-but-dropdowns), you know Poisson numbers should be steppers)
2. The **Most Important Thing** (“Find Options” – which is a very confusing way to say “Submit”, by the way) is **gray and unnoticeable**. A much _less_ important thing (“Help”) is immediately next it, but bigger and _more_ visible.

The Squint Test says the _Most Important Thing_ must be the _most visible thing_. What’s the MIT? The ticket textbox (or stepper 😉) controls and “Submit” button.

If you make it past this page, the next page is even worse.

![image](https://www.learnui.design/img/4-rules/squint-ak2-blur.png)

blurred out version of a train ticketing ui

What will you click: gray button the _left_, or identical gray button on the _right_?

Hope you chose left!

![image](https://www.learnui.design/img/4-rules/squint-ak2.png)

In rushing through this form, I actually clicked 'help' first. Oops. My second time on this page, I clicked 'Go Back', having processed there was an 'Add' and 'Go Back' button, and in the other 99.999% of (left-to-right language) websites, 'Go Back' is always on the left.

Again, _when I squint my eyes and look at the design, I can’t tell what’s important_.

Like the Law of Locality and Anything But Dropdowns, the Squint Test is a fairly simple law to enforce. Here’s like a 30-second wireframey redesign.

![image](https://www.learnui.design/img/4-rules/squint-ak2-redesign.png)

wireframe redesign of a train ticketing ui to pass the squint test

Does it work?

![image](https://www.learnui.design/img/4-rules/squint-ak2-redesign-blur.png)

blurred out version of a train ticketing ui wireframe reddesign to pass the squint test

You tell me. Four radios and a button. And a tiny little link below it.

I’m not trying to pick on AlaskaTrain.com. You see this kind of stuff all over.

Here’s the signup screen for the city guide app Foursquare (blurred, of course).

![image](https://www.learnui.design/img/4-rules/foursquare-blur.png)

blurred out foursquare ui

How do you actually submit the required data? (i.e. the Most Important Thing)

Hint: it’s hidden in plain text in the upper-right corner 😉

![image](https://www.learnui.design/img/4-rules/squint-foursquare.png)

foursquare ui redesigned to pass squint test

But Foursquare is just following Apple’s design standards here. Unfortunately, violating the Squint Test is a tradition even among industry leaders.

![image](https://www.learnui.design/img/4-rules/squint-calendar.png)

ios calendar app failing the squint test

One way to find the Most Important Thing is to consider what percentage of pageviews will involve a certain action. Here’s flashcard/memorization software Anki analyzed in this way.

![image](https://www.learnui.design/img/4-rules/squint-anki.png)

action frequency analysis of Anki UI

For every 100 flashcards I view, I will _then go on to…_

- Show the answer (approx. 95 times)
- Navigate back to the list of decks (twice)
- Start adding cards (twice)
- Use some other feature (very rarely)

This sort of analysis really hints at what kind of interface would work better here.

- **Emphasize** the most-commonly used functionality (at first approximation, “most used” equals “most important”)
- **Deemphasize, hide, or remove** the less commonly used functionality

![image](https://www.learnui.design/img/4-rules/squint-anki-redesign.png)

wireframe redesign of Anki UI to pass the squint test

Now this is just a start (I’d want to see if users understood that the unlabelled plus button _added cards_, for instance). But with just a couple simple heuristics, we’ve reduced a cluttered, confusing interface of 10 UI elements down to just 5. A reduction of… check my math here… 50%.

For more on the Squint Test, check out my YouTube [video redesign](https://youtu.be/B4XHaboNMOI) of the Timezon.es web app. Or, if you don’t have 10 minutes, here’s a [scannable, illustrated blog post](https://www.learnui.design/blog/squint-test-ui-design-case-study.html) with the same step-by-step redesign.

## 4. Teach by example

_If you’re introducing users to new concepts, a few examples can be worth 1000 words - which your users wouldn’t read, anyways._

We have a weird tendency to try and explain things _in words_ when _examples_ would be much clearer.

Consider the startup Teeming.ai, who reached out to me to ask about their homepage design. Subheads on the page read:

- “Teeming takes the **isolation out of remote work**”
- “Teeming helps with **remote team building**” as well as “**learning, problem solving, having fun, and motivating each other**”
- “Teeming and video for **synchronous [communication]**”
- “Works with all your **favorite video platforms**”

But here’s my question for you. _**What does Teeming actually do?**_

![image](https://www.learnui.design/img/4-rules/tbe-teeming-1.png)

teeming.ai UI doesn't teach by example

It’s tough to tell. I know it has something to do with… _good vibes for remote workers_? But I have no concrete idea how it would help me, so I wouldn’t otherwise try it, recommend it, etc.

(Sorry Teeming, you know I ❤️ you)

Next, let’s look at [IFTTT](https://ifttt.com/). Maybe you already know what they do - in which case, _pretend you don’t_, and try to figure it out from these headlines on their homepage:

- Automatically light the way for the pizza delivery guy (Dominoes+Hue)
- Post your photo anywhere and see it everywhere (Instagram+twitter)
- Make your voice assistant more personal (Google Assistant+iOS Calendar)

![image](https://www.learnui.design/img/4-rules/tbe-ifttt-1.png)

IFTTT UI teaches by example

![image](https://www.learnui.design/img/4-rules/tbe-ifttt-2.png)

IFTTT UI teaches by example

![image](https://www.learnui.design/img/4-rules/tbe-ifttt-3.png)

IFTTT UI teaches by example

You don’t have to list too many examples to paint a decently clear picture: _IFTTT hooks apps together to do things they couldn’t do alone_.

The crazy part is, if you visit their homepage, they first explain it in text:

_IFTTT helps your apps and devices work together in new ways. IFTTT is the free way to get all your apps and devices talking to each other. Not everything on the internet plays nice, so we’re on a mission to build a more connected world._

YAAAAWN.

My question: **which gives you a better idea of the app?** The examples, or the description? 🤔

I think it’s the _examples_. The description only resonates once I see a few examples of how it can help me.

> 

The description of your complex new app/feature only resonates once I see a few examples of how it can help me.

But examples aren’t just for landing pages. Here’s what you see when you first sign into project management tool [Basecamp](https://basecamp.com/).

![image](https://www.learnui.design/img/4-rules/tbe-basecamp-1.png)

Basecamp UI teaches by example

Rather than seeing a totally blank page, you see _two obviously pre-fabricated example projects_ that **teach you, by example, how the whole app works** (and also gives you an idea of what the tool will look and feel like when you’ve been using it a while).

Seriously, I can browse through fake _chat logs_ by fake _users_ discussing fake _file uploads_ and fake _to-do items_.

![image](https://www.learnui.design/img/4-rules/tbe-basecamp-2.png)

Basecamp UI teaches by example

There’s even a friendly… _mountain?_… telling me I can watch a 2-minute explanatory video about this sample project.

And thank you, Mr. Mountain, for the lead-in: _providing videos showing usage is another way of teaching by example_! Not only does the sample project model teach by example what my projects will look/feel like, but the video teaches by example what it looks like to _use_ the software.

Brilliant.

If your app allows users to _create_ something, a **showcase** is a great way to teach by example just what’s possible.

The beloved painting app [Procreate](https://procreate.art/) won an Apple Design Award, the App Store Editor’s Choice, the App Store Essential awards, and John Gruber called it “groundbreaking”, etc. – and yet none of this is as _viscerally_ exciting as seeing what you can create with it.

![image](https://www.learnui.design/img/4-rules/tbe-procreate.png)

This ain't no ordinary painting app.

Whoa.

That’s no MS Paint.

The showcase is a **powerful tool** for making it clear just what’s possible with your app.

So: if your app does something new and unfamiliar – or relies on new and unfamiliar concepts – you should get acquainted with **the ways of teaching by example**. The _moment you realize_ that you’re introducing users won’t have seen before, you should start thinking: _how can I give an example to make this clearer_?

> 

The moment you realize that you’re introducing users won’t have seen before, you should start thinking: how can I give an example to make this clearer?

In review, my favorite ways of doing this:

1. On any page that tries to get the user to use a feature/app/etc., show examples of what they can do with your tool
2. Use the “first load” experience to provide sample data, showing by example what the properly-working app will look like
3. Strategically inject help content (like articles, videos, or tooltips) inline with the feature that show how to use it
4. Does your app allow users to create something? Include a user-submitted gallery of examples to spur imaginations

Make sense? Let’s call it a day.

Alright, that wraps things up.

There are plenty more rules for “speaking interface” that I cover in my video course [Learn UX Design](https://www.learnui.design/courses/learn-ux-design.html), but these are some of the ones that I’ve used the most over the years. If you like these, check out more of my design writing on Design Hacks, where I send occasional, original design writing – as well as updates when [Learn UX Design](https://www.learnui.design/courses/learn-ux-design.html) is open for enrollment.


