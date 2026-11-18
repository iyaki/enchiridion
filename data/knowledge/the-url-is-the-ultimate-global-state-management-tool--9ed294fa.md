---
title: "The URL is the ultimate global state management tool"
notion_id: 9ed294fa-29a6-41e6-9d7a-f67eb8813f42
notion_url: https://app.notion.com/p/The-URL-is-the-ultimate-global-state-management-tool-9ed294fa29a641e69d7af67eb8813f42
last_edited: 2026-09-21T17:00:00.000Z
source_url: https://www.notion.so/9ed294fa29a641e69d7af67eb8813f42
tags: ["English", "Web Development", "Article", "Jacob Paris"]
---
[← Back to guides](https://www.jacobparis.com/guides)

The [Web History API](https://developer.mozilla.org/en-US/docs/Web/API/History) is the most underrated state management tool available to developers.

It's the only state management tool that's built into the browser, and supported by every browser. It doesn't require any libraries or frameworks, and **it outperforms every other state managment tool.**

The more your URL reflects the data that needs to be loaded, the more you can optimize your data fetching strategy.

Client side state management tools do not pass their knowledge along to the server, so the browser needs to do an extra round trip.

1. fetching initial data
2. rendering the page
3. accessing the state management tool
4. fetching the data it needed in the first place
5. rendering again

If you're using a server side state management tool The URL is one of these., you can skip the first 3 steps.

A URL that looks like `?query=swimming&sort=price&page=4` tells the server exactly which data it should fetch up-front and send to the browser, fully rendered.

## [Sharable forms and filters](https://www.jacobparis.com/guides/url-as-state-management#sharable-forms-and-filters)

Let's look at an Airbnb URL

Go to Airbnb, open a listing, and in your browser console, open `new URL(window.location.href)`.

You'll get an object that looks something like this

```plain text
hostname: www.airbnb.comhref: https://www.airbnb.com/rooms/12345678?adults=1&children=0&infants=0&pets=0&check_in=2023-05-05&check_out=2023-05-10pathname: /rooms/12345678searchParams:adults: "1",children: "0",infants: "0",pets: "0",check_in: "2023-05-05",check_out: "2023-05-10"
```

The pathname part is obvious – of course we need to know what page we're on. But the searchParams part is interesting.

Each page contains a small booking form, where you can enter your information and get a price quote.

What should happen when a user refreshes the page? Should the form be cleared, ready for them to enter their information again? If we want to pre-fill the form with the information they've already entered, storing it in the URL isn't the only way. It could be part of local storage, or a cookie, or in a database.

But people also share prospective bookings with other people. If I'm planning a trip with a friend, I can send them a link to the listing I'm looking at, and they'll see the same price quote I'm seeing, with the same dates and number of guests.

No other state management tool achieves the sharability of state that the URL does.

## [Back and forward through paginated data](https://www.jacobparis.com/guides/url-as-state-management#back-and-forward-through-paginated-data)

The history stack is a list of URLs that the user has visited. When you hit the back button, you go back to the previous URL in the stack.

This has important usability implications. If you're on a page, and you click a link to go to another page, and then you hit the back button, you expect to go back to the page you were on before you clicked the link.

One particularly frustrating example of this done poorly was the now-defunct app Spectrum [Spectrum](https://github.com/withspectrum/spectrum) was acquired by GitHub and is now called GitHub Discussions. .

Spectrum was a forum, where discussion topics were split into pages upon pages of threads. They had a search feature, which correctly used the URL to store the search query. If you searched for something, and then refreshed the page, you would still see the search results for that query.

Unfortunately, they didn't use the URL to store the page number.  You could be on page 10 of a search, click something that looked promising, and the back button would take you back to page 1 of the search results.

Users use the back and forward buttons as part of their natural browsing experience. Mobile users will swipe pages left and right to navigate forward and back.

All web applications should be built with the understanding that history navigation is part of the user interface.

## [Opening a modal sometimes feels like a navigation](https://www.jacobparis.com/guides/url-as-state-management#opening-a-modal-sometimes-feels-like-a-navigation)

Modal popup dialogs can be a useful way to show additional information to users, without taking them away from the page they're on.

On smaller screens, there is rarely enough space to show a modal and the page content at the same time, so many websites make the modal take up the whole screen, which looks like a new page to the user.

For this reason, it's important that the modal can be closed by hitting the back button. Mobile users will often swipe left to trigger their device's back navigation. If you take them to the previous page instead of closing the modal, they'll be confused and feel like they just went back 2 pages.

If the user should be able to bookmark the modal, or share it with someone else, then it should have its own URL. It's not important whether this is part of the pathname like `/issues/edit` or a query parameter like `/issues?modal=edit`, as long as it's in the URL.

But people use modals for different things, so the best way to handle this can change depending on the context.

### [Sharing the URL without sharing the modal](https://www.jacobparis.com/guides/url-as-state-management#sharing-the-url-without-sharing-the-modal)

If the modal is relevant to the user, but not to anyone else, then you may want a URL that doesn't share the modal.

Instead of a query param, you can use [Location State](https://developer.mozilla.org/en-US/docs/Web/API/History/pushState) to store whether the modal is open or closed without it being part of the URL.

This state is still stored in the browser history. Forward, back, and "re-open closed tab" will all work as expected.

React Router's and therefore Remix  Link component has a `state` prop that can be used to store arbitrary data in the browser history.

```plain text
export default function Example() {const location = useLocation()const shouldShowModal = location.state?.modal === "edit"  return shouldShowModal ? (    <Modal />  ) : (    <Link to="/issues" state={{ modal: "edit" }}>      Edit Issue    </Link>  )
```

### [Use replaceState when you close the modal](https://www.jacobparis.com/guides/url-as-state-management#use-replacestate-when-you-close-the-modal)

Modals often have a close button, whether it's a little X in the corner or a button that says "Cancel".

It may seem intuitive to use a simple link for this, since navigating to a page with the modal closed is the same as closing the modal, but this is a common mistake.

The back button should always take the user to the previous logical state. Pressing the back button to close a modal makes sense, and pressing it again should take them to the previous page.

Hitting a physical cancel button **also** takes them back to a previous logical state, so we need to reflect that in the history stack. Otherwise, the user will hit the back button and be surprised when the modal re-opens..

The best way to do this is to use the [replaceState](https://developer.mozilla.org/en-US/docs/Web/API/History/replaceState) method to replace the current history entry with a new one that has the modal closed.

In Remix and React Router, you can use the `replace` prop on the Link component This works on forms too!  to replace the current history entry instead of adding a new one.

```plain text
function Modal() {return (    <div>      <h1>Edit Issue</h1>      <Link to="/issues" replace>        Cancel      </Link>    </div>  )
```

## [Caveats](https://www.jacobparis.com/guides/url-as-state-management#caveats)

- There is a limit to how much data you can store in the URL. If you're storing a lot of data, you may need to use a different storage mechanism, such as keeping an ID in the URL and storing the data in a database.
- Don't store anything sensitive in the URL. It's not encrypted, and it's stored in the browser history.
- If you store an access token in the URL, like Google Docs does when you ask for a shareable link to a private document, inform the user of the risks and give them a way to revoke the token.
