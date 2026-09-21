---
title: "btn.social - Implement social authentication in seconds"
notion_id: 7ac41f14-8bc9-4a7e-b263-348c514d570a
notion_url: https://app.notion.com/p/btn-social-Implement-social-authentication-in-seconds-7ac41f148bc94a7eb263348c514d570a
last_edited: 2023-04-22T20:02:00.000Z
source_url: https://btn.social/
tags: ["Web Development", "Information Security", "Service", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Lots of providers, with more to come!

Every product speaks to different audiences on different platforms.
You and _your users_ deserve the freedom to choose the social network that works best for them.

That's why we support a variety of social OAuth providers.
From Facebook to GitHub to Instagram... our list is always expanding to keep up with your needs!

```plain text
{
  "network": "facebook",
  "user": {
    "id": "1026983XXXXXX51213",
    "email": "alex.smith@me.com",
    "first_name": "Alex",
    "last_name": "Smith",
    "name": "Alex Smith",
    "name_format": "{first} {last}",
    "short_name": "Alex",
    "picture": {
      "data": {
        "height": 50,
        "is_silhouette": false,
        "url": "https://platform-lookaside.fbsbx.com/platform/profilepic/?asid=1026983XXXXXX51213&height=50&width=50",
        "width": 50
      }
    }
  }
}
```

## Affordable – at any scale!

Unlike other authentication services, we're _not interested_ in tracking monthly active users and we don't believe in significant price jumps as a reward for growing your audience. Pricing should never be a guessing game. We charge for every login request with volume-based discounts. We let _you decide_ how often a user must login. It's your application and your bill... you're in control!

That's **$0.01** every **100** logins!

**$5/mo**

Includes **50,000** logins Extra **10,000** logins per **$1.00**

Business **$79/mo**

Includes **1,000,000** logins Extra **12,500** logins per **$1.00**

Enterprise **$169/mo**

Includes **2,500,000** logins Extra **22,500** logins per **$1.00**

### Pricing Calculator

It will cost$25.00for 250,000 logins in a month.

All plans allow for _unlimited_ applications.
All applications are billed together under one subscription per account.

## Examples – Show, don't tell.

Don't miss the [Usage](https://docs.btn.social/usage/) documentation!

### Attach in Seconds

Adding a Social SSO link with `btn.social` is just like adding any other HTML link:

```plain text
<a href="https://login.btn.social/{APPID}/facebook">Login with Facebook</a>

```

You just need to include the JavaScript SDK once per page:

```plain text
<a href="https://login.btn.social/{APPID}/facebook">Login with Facebook</a>
<a href="https://login.btn.social/{APPID}/google">Login with Google</a>

<script async src="https://btn.social/v1/web.js"></script>

```

This will allow the OAuth flow to work.

However, you must define an `onlogin` callback to receive the user's data upon successful logins. With the web SDK, this can be done by passing a `data-jsonp` attribute with the name of the callback function to run once the SDK has been initialized:

```plain text
<a href="https://login.btn.social/{APPID}/facebook">Login with Facebook</a>
<a href="https://login.btn.social/{APPID}/google">Login with Google</a>

<!-- Call the "myCallback" global function when SDK is loaded -->
<script async data-jsonp="myCallback" src="https://btn.social/v1/web.js"></script>
<script defer>
function myCallback() {
console.log('the btn.social SDK is ready!');
btnsocial.onlogin(function (payload) {
let { network, user, token } = payload;
console.log(`user logged in via "${network}" provider`, { user, token });
    });
  }
</script>

```

### Fully Customizable

Because these are _your links_, you may customize the CSS and HTML to your heart's content.

The web SDK only requires that `<a>` tags be used with the `"https://login.btn.social/"` address.

```plain text
<a class="btn social github" aria-label="login with github" href="https://login.btn.social/{APPID}/github">
<svg>...</svg>
<span>Login with GitHub</span>
</a>

<a class="btn social twitter" aria-label="login with twitter" href="https://login.btn.social/{APPID}/twitter">
<svg>...</svg>
<span>Login with Twitter</span>
</a>

<style>
.btn.social {
cursor: pointer;
position: relative;
display: inline-flex;
justify-content: center;
align-items: center;
text-decoration: none;
border-radius: 0.25rem;
transition: all ease-in 300ms;
border: 1px solid var(--c);
background-color: white;
padding: 0.6rem 1.2rem;
text-align: center;
margin: 0 0.25rem;
color: var(--c);
  }
.btn.social:hover {
background-color: var(--c);
color: white;
  }
.btn.social.github {
--c: #181717;
  }
.btn.social.twitter {
--c: #1DA1F2;
  }
.btn.social span {
margin-left: 0.5rem;
  }
.btn.social svg {
fill: currentColor;
height: 20px;
width: 20px;
  }
@media screen and (max-width: 600px) {
.btn.social span {
display: none;
    }
  }
</style>

```

### Configure `"popup"` Mode

Redirects are used by default. You may _opt into_ the `"popup"` mode, which will house the OAuth workflow within a dialog window (if possible).

The dialog's dimensions are **500×600** pixels by default, but may be customized with the `data-width` and `data-height` script attributes for the Web SDK.

```plain text
<a class="btn social twitter" href="https://login.btn.social/{APPID}/twitter">Login with Twitter</a>
<a class="btn social github" href="https://login.btn.social/{APPID}/github">Login with GitHub</a>

<!-- Enable "popup" mode via the <script> tag -->
<script async data-mode="popup" src="https://btn.social/v1/web.js"></script>

```

When using the npm package, you may programmatically call the `btn.popup()` method and customize the `width` and `height` options:

```plain text
// Use "popup" mode via the "btn.social" package
import * as btn from 'btn.social';

/*
  HTML:
    <button class="btn social" data-network="twitter">Login with Twitter</button>
    <button class="btn social" data-network="github">Login with GitHub</button>
*/

const APPID = 'btn-social-appid';
$('.btn.social').on('click', async evt => {
evt.preventDefault();

let width = 500, height = 700;
let provider = evt.target.getAttribute('data-network');

if (provider === 'twitter') {
height = 600;
width = 400;
  }

let { network, user, token } = await btn.popup(APPID, provider, { width, height });
console.log(`user logged in via "${network}" provider`, { user, token });
});

```

## Frequently Asked Questions
