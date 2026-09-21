---
title: "Step-by-Step Guide to Building and Launching your Chrome Extension"
notion_id: 52b0a938-5a85-4e1d-b412-cf01412dd006
notion_url: https://app.notion.com/p/Step-by-Step-Guide-to-Building-and-Launching-your-Chrome-Extension-52b0a9385a854e1db412cf01412dd006
last_edited: 2023-04-25T14:20:00.000Z
source_url: https://hackernoon.com/step-by-step-guide-to-building-and-launching-your-chrome-extension-th3h3tx1
tags: ["Article", "Hackernoon", "English", "Programming"]
---
[https://hackernoon.com/step-by-step-guide-to-building-and-launching-your-chrome-extension-th3h3tx1](https://hackernoon.com/step-by-step-guide-to-building-and-launching-your-chrome-extension-th3h3tx1)

While building my first chrome extension, Foragear- Quick Search Tool, I struggled to find an article that covered the entire ideating, building, and launching process of chrome extensions. To make the lives of future chrome extension builders easier, here is an all-in-one guide to help you through the process.

Hopefully, this saves you the time of looking through and synthesizing information from multiple articles.

The article is broken down into the following sections:

1. Ideating: why Chrome extensions and what to build
2. Building: viewing the code of similar extensions
3. Building: Chrome extension architecture
4. Building: Chrome APIsBuilding: testing and troubleshooting
5. Launching: navigating the chrome web store submission process

Without further ado, let’s dive right into it.

## Ideating: Why Chrome Extensions?

Ease of development- A key advantage of building browser extensions is its low barrier to entry. For beginner coders, extensions are generally easier to develop as compared to mobile applications or websites that deliver similar functionality and value.

Personally, a chrome extension was my first ever ‘live project’ and I completed the ideation, building, and launching process within 2 weeks as a beginner.

Extent of reach- From the graph below, it is clear that Chrome beats other browsers by a large margin in market share. To maximize the number of users that can download and use your browser extension, it therefore makes the most sense to build for Chrome. Once launched, your extension would be available for download by all Chrome users on the Chrome Web Store.

Why not just launch it on all browsers? That is possible, but you ultimately have to start with one browser, and chrome looks like the best bet. It is also important to note that while extensions may be adapted to work in different browsers, the codebase cannot be copied wholesale and may require significant editing.

If your application makes extensive use of browser-specific APIs, it would be a hassle to adapt it for another browser. Not all browsers may have corresponding analogous API with similar functionalities that you can tap into, and even if such APIs exist, the API functions may look slightly different.

## Ideating: What to Build?

While extensions are great, they are not the best platform for all types of product ideas. If you are looking to build complex, multi-functional software, browser extensions would not be appropriate. Rather, extensions should be simple to use and should focus on tackling one specific pain point or deliver a single function effectively. In fact, in the Chrome Web Store Submission process, Google states that “an extension must have a single purpose that is narrow and easy-to-understand”.

For example, my extension (Foragear) alleviates the pain of navigating through long webpages. All it does is to auto-scroll to and highlight keywords (based on user search query) and phrases (based on google search description) so that users can quickly identify relevant sections of a webpage.

To aid in your brainstorming process, here are some key categories of extensions from a developer’s point of view:

1. New tab extensions: customize new tab pages for aesthetic purposes, provide useful information on each new tab, or integrate some functionality such as note-taking into the new tab. Example: Momentum

2. Website-specific extensions: their functionality is restricted to specific websites. Example: RateFlix, which overlays Rotten Tomatoes ratings on your Netflix catalog.

3. General extensions: they work on most/ all websites. Example: Dark Reader, which allows you to switch to dark mode and control the display of websites you browse.

## Building: viewing the code of similar extensions

As a new developer unfamiliar with chrome extension development, you can look through the code of other similar extensions for inspiration. Whenever you add an extension to your browser, the code gets stored locally on your desktop, allowing you to view it.

Step-by-step process:

1. Enter “chrome://extensions” into your search bar
2. Turn on developer mode, scroll to the relevant extension and copy the ID
3. Go to your file explorer and search for a folder named with the ID. You can also search manually using the following directory (may differ)- C:\Users\<user>\AppData\Local\Google\Chrome\User Data\Default\Extensions\<ID name>.

All of the extension’s associated files should be in the folder and the files can be viewed using a code editor like Visual Studio Code. However, you should refrain from copying chunks of code (after all it is not your intellectual property). If you want to copy, do check Github if they have made their code open source. Besides that, many extensions minify their files for storage reasons. For such files, you can ‘unminify’ them online if you wish to read the code. Personally, I viewed other extensions’ code to better understand how similar chrome extensions work.

## Building: Chrome extension architecture

Like most web-based software, Chrome extensions are built using Javascript, HTML, and CSS. To get an in-depth understanding of how extensions work, it would be useful to read through Google Chrome’s Developer Guide. As Google’s guide provides little examples, I have included small snippets from Foragear’s code to illustrate how extensions work.

Manifest.json

Extensions are required to have a short manifest.json file that gives the browser basic details of the extension and maps out the other files used in it. Here is an adapted version of Foragear’s manifest.json:

```plain text
{ "name": "Foragear- Quick Search Tool", "version": "1.0.0", "description": "Autoscroll to relevant sections of a page and have relevant keywords/ phrases highlighted", "icons": { "128": "icons/128.png", "48": "icons/48.png" }, "background": { "persistent": false, "scripts": [ "background.js" ] }, "content_scripts": [ { "js": [ "contentmain.js" ], "matches": [ "https://www.google.com/search?*", "http://www.google.com/search?*" ] } ], "page_action": { "default_title": "Foragear- Quick Search Tool", "default_popup": "popup/popup.html" }, "permissions": [ "activeTab", "<all_urls>" ] }
```

The file can be broken down to the following sections, in order:

1. Basic details of the extension and directories to the extension icons (2 different resolutions provided)
2. We tell the browser that the background script’s name is background.js
3. We tell the browser that the content script’s name is contentmain.js and to run the content script only when the browser’s URL starts with “www.google.com/search?”
4. We provide the directory of our popup’s HTML file, which would be triggered when clicking the icon on the chrome toolbar
5. We request for user permissions

Background Script

The background script stays dormant most of the time and contains listeners that trigger the script only when certain browser events occur.

```plain text
chrome.runtime.onMessage.addListener( function(request, sender, sendResponse) { if (request.message === "custom_message") { //do something } } )
```

In the example above, the chrome.runtime.onMessage.addListener API has been used to wait and listen for a message called custom_message to be sent from another script. Once received, the background script is instructed to do something (eg. send a message or execute a content script).

## User Interface (UI) Elements

This would be where HTML and CSS come into play. The main UI element for your extension is likely to be the popup that appears when users click your extension icon on the Chrome toolbar. You can also use Javascript to make the UI elements interactive.

Content Scripts

Content scripts read and modify webpages. They are written in Javascript and are executed when the webpage loads. For example, Foragear’s content script is used to read through a page and highlight keywords and phrases, directly modifying the content displayed on the user’s screen, as shown here:

The figure below, taken from Google’s guide, illustrates the interactions between the various files. Sending and receiving messages is the key method of communication between the files, and is used to coordinate the entire extension’s functionality.

## Building: Chrome APIs

You have an idea and you understand the extensions architecture. Now what? In this section, we’ll dive into the most useful tool for your development process: Chrome APIs. These APIs allow you to request for user permissions, listen and react to events, store user data, interact with the browser, and access special features such as Bluetooth, audio, and notifications. In short, they act as the building blocks for your extension.

You can look through Chrome’s API documentation and check out APIs that may be relevant to your extension. To get you started, we would cover some of the most common APIs:

Chrome.runtime

1. chrome.runtime.sendMessage: it allows you to send a single message to event listeners, allowing interaction between your different scripts (unable to send to content scripts)
2. chrome.runtime.onMessage.addListener: listen to and gets triggered when a message is received from an extension process/ another script

Chrome.tabs

If your extension has anything to do with browser tabs, you need this API.

1. chrome.tabs.get: get details about any specified tab (eg. URL, title, id, whether it is active). This is useful if you want to trigger actions only when the user is on certain websites (eg. if your extension is website-specific).
2. chrome.tabs.getCurrent: get details of the current tab
3. chrome.tabs.sendMessage: send a message to a specified tab’s content script(s), with an optional callback to run when a response is sent back
4. chrome.tabs.create: create a new tab (you can specify a URL)

## Building: testing and troubleshooting

Using chrome://extensions

While building your extension, it is a good practice to constantly test and debug it. Testing it after each ‘stage of development’ would allow you to quickly identify and fix errors in your code. If you write a significant amount of code without testing and errors pop up, it may be hard to understand the source of the error.

To test your extension:

1. Type “chrome://extensions” into your chrome search bar
2. Turn on developer mode and click “load unpacked” to upload a folder containing your draft code and other materials (eg. extension icon image)
3. Turn on your extension and test it out! Note that you can edit the code locally, then toggle the extension off/on again to update its functionality and test it again. You don’t have to click “Load unpacked” every time. This works as long as you do not change the directory of your files.

The “chrome://extensions” page is also useful in that if the script cannot be executed fully, chrome would identify the specific files and lines of code where an error occurs.

## Using the Console Panel

Another useful tool for debugging is the console panel. You can access it by clicking the three dots > More Tools > Developer Tools.

While browsing in chrome, the console panel highlights any scripts that have failed to run and the specific lines that failed. Moreover, you could use console.log() at different parts of your scripts. The content of these logs would appear in the console panel, allowing you to know which parts of the code were successfully executed and where it failed.

Frustration

Just wanted to end this section with an ominous note. Testing and debugging would probably take the lion’s share of your time, and you would often feel like ripping your hair out (especially when you don’t understand why your code doesn’t work). In any case, Google and Stackoverflow are your best friends.

## Launching: navigating the chrome web store submission process

Refer to Google’s article for a detailed explanation of the process of publishing an extension to the chrome web store. In short, here is the process:

1. Register for a chrome web store developer account. It costs 5 USD to register, but this one-time fee allows you to publish up to 20 extensions.
2. Upload your extension files as a single zipped folder to the developer dashboard
3. Fill in details such as description and summary, upload promotional materials (promo tiles, logo icon, demo video etc.), and justify user permissions you request for
4. Submit the extension for review.

One notable aspect of the submissions process is the permission justification section. For any permission that you request for in the manifest.json file of your extension, you have to justify why it is needed for your extension to function.

When developing the extension, be careful to ensure that you are requesting as little permissions as possible. For example, if your extension can work with the “activeTab” permission that gives temporary access to the currently active tab, don’t request for the “<all_urls>” permission that allows you to fetch information from non-active tabs too. You may find a list of permissions here.

## Conclusion

This article is long enough as it is, so I’ll end it off here. If you are looking to launch a cool extension or any other product, feel free to contact me and I’ll be happy to help you launch it on Product Hunt. Happy coding!

Also published on: https://codeburst.io/ideating-building-and-launching-your-chrome-extension-step-by-step-d713ed71cfd8
