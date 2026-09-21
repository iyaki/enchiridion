---
title: "SubResource Integrity (SRI) Hash Generator"
notion_id: 1f20a9a1-6d89-4aaf-9c0b-43c9bffb80ab
notion_url: https://app.notion.com/p/SubResource-Integrity-SRI-Hash-Generator-1f20a9a16d894aaf9c0b43c9bffb80ab
last_edited: 2023-04-22T20:00:00.000Z
source_url: https://www.srihash.org/
tags: ["English", "Web Development", "Service"]
---
## What is Subresource Integrity?

SRI is a new [W3C specification](https://www.w3.org/TR/SRI/) that allows web developers to ensure that resources hosted on third-party servers have not been tampered with. Use of SRI is recommended as a best-practice, whenever libraries are loaded from a third-party source.

Learn more about [how to use subresource integrity](https://developer.mozilla.org/docs/Web/Security/Subresource_Integrity) on MDN.

## How is Subresource Integrity different to HTTPS?

TLS ensures that the connection between the browser and the server is secure. The resource itself may still be modified server-side by an attacker to include malicious content, yet still be served with a valid TLS certificate. SRI, on the other hand, guarantees that a resource hasn't changed since it was hashed by a web author.

## How can I generate Integrity hashes?

Use the generator above or the following shell command:`
          openssl dgst -sha384 -binary `**`FILENAME.js`**` | openssl base64 -A`

## Why do I need to include `crossorigin="anonymous"`?

When the request is not on the same origin the `crossorigin` attribute must be present to check the integrity of the file.
 Without a `crossorigin` attribute, the browser will choose to 'fail-open' which means it will load the resource as if the integrity attribute was not set, effectively losing all the security SRI brings in the first place.

`crossorigin="anonymous"` results that no credentials are sent to the cross-origin site hosting the content. However, it will send an `Origin` HTTP header. If the server denies including the resource (by not setting the `Access-Control-Allow-Origin` HTTP header), the resource will not be used by the browser.
 You can find more information on [MDN.](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/crossorigin)

## Test your browser

Check out [SRI on caniuse.com](https://caniuse.com/#feat=subresource-integrity) to see specific browser version support information.

To fully test your browser for subresource integrity support, please open [this page](https://w3c-test.org/subresource-integrity/subresource-integrity.html).
