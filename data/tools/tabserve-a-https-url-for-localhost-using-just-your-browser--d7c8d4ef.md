---
title: "Tabserve - A HTTPS URL for localhost using just your browser"
notion_id: d7c8d4ef-f362-42ff-b06b-0f18150c86b4
notion_url: https://app.notion.com/p/Tabserve-A-HTTPS-URL-for-localhost-using-just-your-browser-d7c8d4eff36242ffb06b0f18150c86b4
last_edited: 2023-07-05T19:17:00.000Z
source_url: https://tabserve.dev/
tags: ["Web Development", "Network", "REST API", "Untried", "Service", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Intro

Tabserve is a web app that uses a single [Cloudflare worker](https://workers.cloudflare.com/) combined with browser based [web workers](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API/Using_web_workers) to create a reverse proxy from the internet to your localhost.

This gives you an HTTPS URL on `your-domain.com` (such as `https://anything.your-domain.com`) that forwards traffic to any `localhost` address, such as `http://localhost:1234`.

You can instantly create any number of subdomains as Cloudflare provisions a TLS wildcard certificate - no waiting for a DNS propagation to take effect.

This lets you share your local web server with the world, receive webhooks, test from different devices, and use browser JS API’s that are HTTPS-only. You can also run small to medium production services from any machine with a web browser (although Tabserve is currently alpha - [please submit any issues you encounter](https://github.com/emadda/tabserve/issues/new)).

Your local web service is protected and optimized by Cloudflares Web Application Firewall (WAP), and uses the latest HTTP protocols and optimizations (HTTP-3). Your CF worker will be created in the CF location nearest to you in one of [300 global locations](https://www.cloudflare.com/en-gb/network/) which reduces latency and speeds up instant reloading during development.

Each subdomain currently maps to a single [Durable Object](https://developers.cloudflare.com/workers/learning/using-durable-objects/). This gives you a throughput of between 100-500 RPS (informally tested), and isolates each subdomain from each other for security and fault/load isolation. The [WebSocket Hibernation API](https://developers.cloudflare.com/workers/runtime-apis/durable-objects/#websockets-hibernation-api) is used, so you only get charged for the time you are forwarding requests, and should generally be less than $5/month ([pricing](https://developers.cloudflare.com/workers/platform/pricing/#durable-objects)).

Current reverse proxy tools require you to install a CLI on your machine, which gives it access to your entire machine and can be difficult to observe and configure.

Tabserve “installs” by just loading the web page (less than 300ms), and is [isolated using the Chrome browser](https://www.chromium.org/Home/chromium-security/brag-sheet/). V8 and other JS engines have some of the most robust and well tested sandboxing. Cloudflare workers also use V8, which means both the server and client of the reverse proxy have strong sandboxes. V8 is also a very fast runtime.

Note: You must upload a single small [Cloudflare worker](https://github.com/emadda/worker-tabserve-reverse-proxy) to your Cloudflare account and have a domain added to Cloudflare DNS to use Tabserve.

## Interesting things you can do

- A public HTTPS URL for localhost 
- Share your local web server with your any team members around the world.
- Receive web hooks from API’s to your local web server.
- Use web API’s that require HTTPS.
- Test on different devices.
- Unlimited subdomains such as `a.your-domain.com`, `b.your-domain.com`. 
- Each subdomain has a valid HTTPS cert instantly - no DNS or certificate set up required.
- Use any device with a browser as a reverse proxy. 
- iPad, iPhone etc (alpha)
- This allows you to route web traffic from the internet to the LAN without setting up firewall rules or installing third party proxy clients that have access to your whole computer.
- Forward HTTP requests to any domain, not just localhost.
- Configure reverse proxies remotely. 
- Each browser you leave open on a machine shows up in your Tabserve UI.
- Servers can be added and removed remotely.
- Debug your server using Chrome dev tools. 
- The Tabserve UI runs completely locally and uses your browsers `fetch` API - these show up in the network tab.
- Chrome dev tools is familiar and feature rich.
- Chain reverse proxies (Tabserve -> GUI -> local web server). 
- Use your preferred HTTP GUI to observe requests.
- Use Charles or another GUI to set up a reverse proxy with rule like: `localhost:1234 (Tabserve forward_to) -> localhost:5678 (your actual dev web server)`
- Use as a stand-in for a production server. 
- You can use `xyz.your-domain.com` in the Tabserve UI, and then switch your Cloudflare DNS to a real production web server when ready.
- You can use Cloudflare features for the public server URL. 
- Latest HTTP protocol support.
- Web Application Firewall (WAP)
- On-the-fly optimizations for JS and images.
- Use multiple root domains. 
- You can add any number of domains to Tabserve.
- And probably billions of other things.
- [Let me know how you’re using Tabserve in the show and tell](https://github.com/emadda/tabserve/discussions/new?category=show-and-tell).
- Contact [support@tabserve.dev](mailto:support@tabserve.dev) for direct help.

## Sequence diagram

[eyJ2ZXJzaW9uIjoiMSIsImVuY29kaW5nIjoiYnN0cmluZyIsImNvbXByZXNzZWQiOnRydWUsImVuY29kZWQiOiJ4nO1d61biyFx1MDAxNv7fT8Fy/lx1MDAwZZm6X+afgrZ4V1BbT89yXHUwMDA1iFx1MDAxMFxyXHQmQcBZvdY8y3m08ySnXHUwMDAySC5cdTAwMTBcZjZonJZe7YJKKtmp2t/+9t61XHUwMDBi/v5SKGz4w66x8Wdhw1x1MDAxODR0y2y6en/j96D90XA907HVITT67Dk9tzE6s+37Xe/PP/5cYntoXHKnM+5lWEbHsH1PnfdcdTAwMWb1uVD4e/RXXHUwMDFkMZtB353KYf2OXHUwMDE2z/fJsH5+s3OLYdE9XHUwMDE4dVx1MDAxZJ30LIxl2kbYOlBNUFx1MDAwMFxyRF9senioXHUwMDBlMzD92Debfls1hS1tw2y1fdVEXCKn6XbLMmKneb7r3Fx1MDAxYiXHctxAht+gXHUwMDEx/Fx1MDAwYsWo6437luv07Ob0XHUwMDFj39Vtr6u76pHD825Ny6r6Q2s8WHqj3XMjV1x1MDAxOd/lclwiJEy0T/s1da9tNMNu6r6ttm14Xkxkp6s3TD9cdTAwMThcdTAwMDBcYsLWQMhupTmahL9CsVxcvWNUglmwe5ZcdTAwMTW9sN2cXFz4ebLCmUCTllx1MDAxZqGchlx1MDAxMVxcg2NcdTAwMDGZQEROXHUwMDBmhFxuXHUwMDAzpUi2XHUwMDFlOfZIeSBiXGZgQimenmF6ZaU1/uiqt7rlXHUwMDE54YBcdTAwMDeibYdcdTAwMWFcdTAwMTWTu9dt6uNOSlx1MDAwZc4hkpJcdTAwMTBcdTAwMWGKo1ToPtnHclx1MDAxYfdz7tN1zKjKXHUwMDA2r/BdIVx1MDAxY+7Rh+n7v35/+WylcOH5X1x1MDAxMv02LN3zS06nY/rqQU5cdTAwMDIhklx1MDAwMnu+7vpbpt007VbymGE3U46Mem26rtNvXHUwMDFi+syEq37JY19cIlM8XHUwMDA17HwsRsxcdTAwMDJGqbMsIOSYccQzzPLYMFxcM94olSsnjbPN/lx1MDAxZLI9XHUwMDFmWLXNXHUwMDFjYWrWXFxcZkZmJ/ycMEOrsiaLzVnMJkZcdTAwMWV8XGZSSCFSXHUwMDEzIWF4ZK22IVx02MhcdTAwMWRWXGLX14FcIiPYloHN21x1MDAxOJD4XHUwMDE0z5qT14GX0jTwXCIoXHTjgoUnvITd9v711t1Wi1x1MDAxOVx1MDAxNZNVXHUwMDExZ2feUcXMPXYhQO9cZt5cdTAwMTlOXHUwMDA1klx1MDAwM0Ws4ch/ovWjoXVV+Fx1MDAxNKnkiiUjnFxuXHUwMDFjQuVFfDKK2lx1MDAxY29J3uQn23flTWI8XFzlXHUwMDFmn4TkXGafhFxiXHUwMDA0XHUwMDA1g5/4/Nfjc1xmnMtcdTAwMDfr6/XmvXFcblx1MDAxYlx1MDAwZbvG3r23eX0zXHUwMDFirbpGw1x1MDAxZutwTIFFXFx9i1x1MDAxMbp51jtcdTAwMDTmaJ5YhIzPIHVcdTAwMTY6eFx1MDAwNjpcdTAwMTPASshcdTAwMTnGXFzM+DrqXHUwMDE4TTWwjFxigVx1MDAwMVx1MDAxNVlcdTAwMWOgXHUwMDE5jE6F+zuid1NN8Y1BOFx1MDAwN1x1MDAxMUV7ZNbWI39cdTAwMWFcZthh5/Zwu1l13SrdmJ7341lXV2hcdTAwMDBcdTAwMTao/XxpZtQ+9jBjk42FxjhcdTAwMDCYYUjV4EtcdTAwMWVDXHUwMDAw1+hcZlx1MDAwMlx1MDAwNNJcdTAwMTiVXHUwMDEyqXBRQEZcImmdKVx1MDAxZVx1MDAxMP3Ew1x1MDAwMjzEre6EqSBQto5LOEfxI6yaUHwulWvBIMhTbmakZGpoK7ZvuLZcdTAwMTGdMcf2q+bTiDtBrHVH75jWMDZRwWU2LbNcdTAwMTVcZsBGw1xiLlx1MDAxNnO7fLOhW9NcdTAwMTM6ZrNpxZTIM0aeWpDamLY21K101epWspCF45ot09atWupcdTAwMDNcdTAwMDV32H1WeqihqP3ZXHT04Ej9X1xi31x1MDAwNYRcdTAwMTTOP081fELNXHUwMDExXHUwMDExnGX3LJtcdTAwMTd71bNDT3fbpauD1p1VqrtmN0egmu9ZUpDOzKv2LF9i+Fx0YDFX7q4keC2Jmll+XFxcdTAwMTFtXXblcWnnpH5Z263ZZVYzKsdn9XegrfmMXHUwMDE0UfnUfDSWXHUwMDEwcEphdpWf/8y5V3lcbjQgXHUwMDAxIJJcdTAwMTOswkFcdTAwMWNcdTAwMDdcdTAwMDBdm/5LqVx0KNWdKURM0oiCT9FAZ9BcdTAwMDBcdFwiXGJxKFdcdTAwMDCH2IG3XHSplmamYFx1MDAxNEuW02veWmpcdTAwMTi/25eOe1x1MDAxZiWon2OvXHUwMDA0T71gs5M8tUiudMaKMCbBr+eqWVx1MDAwNyZcXEhChCNKlyCrev8rXHUwMDE57lx1MDAxZdbOirjt2mxIpbtbyT1yJcpcdTAwMWRZQaQ8RYipXGatyEdgq9JBp0OvcKfW3CudXHUwMDFmXHUwMDEyU24/XHUwMDE5X3PIVlx1MDAxMKaHplx1MDAxNCAgXHUwMDExlpmVfv5D517pXHTVuFxuJDmSmFx1MDAwM1xi43HkXHUwMDFh6VxuXHUwMDEyoakpXHUwMDE2lEBEJWd0Tjw6h69cdTAwMDRBgmKOfiG+qul1z3BcdTAwMWZcdTAwMTUrlNqu0zFcbnXX6XtrY61cdTAwMTeMd5K1XpZu7dxcdTAwMDVh6lx1MDAxMlx1MDAxYueMMVx1MDAxMC2eeFx0xqTfXHUwMDE0x+6F0brqbjfc+7t+XHKdw9zDXHUwMDE4YpI78lK+KEWSRFJTXHUwMDFmgbusonEzrPfPvZuW5XXq1ZYo3V3mkbswS1N6SFx0XHUwMDExmGGSJb206Knzr/VcdTAwMTJpNFxitlx1MDAxMFSRXGaNRJ/JLOjKoy0yirYgVTNOXHUwMDA0wWJcdTAwMTZcdTAwMTFcdTAwMTFDN3XnXGJU0lwiuFx1MDAwMkh8XHUwMDE09jpwXHUwMDE097RcdTAwMWTPX1x1MDAxM129YK+TdDVHnEz8XHUwMDA0xVxclE4441xubG3vXoPTU1Gqe0VUvzWPz2bT+HqwXHUwMDBlN1NsXHUwMDE5U1lcdTAwMTgx48+qRiDXgkVcdTAwMTJGXHUwMDA1XHUwMDE2iHOMZ3XtwyxkeY5lptiAsM/7XHUwMDE1W0Kg7IhkmM9byEIg1c9Qdlx1MDAwMEkoXHUwMDA1yFx1MDAxMiQvn9GPrJC/V7Vlulx1MDAxNia6z1x1MDAxMOTSxZeR+ZpUOY+hbrrGPTrzjk6LrXte3qm3W4PecYTDlSlq9ILnRVx1MDAxYYq0tvTuaIZm5nxFdZ1cdTAwMTNcXC9k8Vx1MDAwNPQjNFx1MDAxZSHpmfIwpEZZgiXqT0r44pqePtm7Q4hcdTAwMDZnvc2Le/KwlVx1MDAxZlx1MDAwNKdkTKXUVECHXHUwMDEwZoJcblx1MDAxOC83h1x1MDAwMmpcdTAwMDCokE+oeJxzXHUwMDEyWVNZNaUvZ2efXHKGZMrRiiVcbnJVqvI6y5GLUpUkaN7XnL1cdTAwMTLeLJUyXGKQQSpcdTAwMDNnr9x+2O2dfKtcdTAwMTiMyk1nt2ua5zq29/OObqjcXY1cdTAwMTJMiVS2jIJEpVx1MDAwMkJAXHUwMDFibVQgRDnTiFx1MDAxMZ6Q7L3xrZx1XHUwMDA0kFx1MDAwMGuJYj/xnW98j3HnfcNnJ1x1MDAxN0ZbXHUwMDE3O1x1MDAwN0/s7LFre1x1MDAxNGZz75Xea5JTglx1MDAwNERcdTAwMTiROLdhXHUwMDAyNUwxU4BcdTAwMDaUXHUwMDEyXHUwMDFhKVx1MDAwN1x0dVbFmYBhXHUwMDEw1IrzaKXUp+u/etdcdTAwMWZcdTAwMTBcdTAwMGUlmrvRiqW6aVCqfohEQ4Z/metfTNXCRPcvicvkY9/Vz/nnOLWiQVA1XHUwMDFjXHUwMDE4gey5ZVx1MDAwNjvFo4ejzevq7s3TYK96UH44dfKDsjRcdTAwMDJXs89cdTAwMTGB6mEh5ijigY+MmHLQXHUwMDExQpJyoLhcdTAwMWQjuU5cdTAwMDJfwlx1MDAxOE5QLTBcdTAwMDSCI5TTSvJP+l6RyXkluln6zkrIJIRcdTAwMDTK7PA+bFxu5Sd2TshcdTAwMTOs7FdcdTAwMDfHnjj/JvNcdTAwMGVvKqnGlOshMFx1MDAwZVx1MDAxY1x1MDAxNEBi6CbKeYeCIa7IXHUwMDExclx1MDAxNbKQhGDvi27IXHUwMDE5XHUwMDE0Sur1rDF9wjvf8Fx1MDAxZaOu8uhWT7tcdTAwMTfG9iku61x1MDAwM703fHw6t7LU0JNEvU9cdTAwMTQr0ypcdTAwMDWKNcKlUFGEXHUwMDE0MVx1MDAxZu9cdTAwMDNcdTAwMTbNv4dcdTAwMDM+r2aeqWhfxkpBs2xcdTAwMTahQnBljzPlS97KzX5e5Nqt1U6qXHUwMDA113joXHUwMDE53kpcbuct4za+6lx1MDAxYl9cdTAwMTPznW5Uh16umY89THJNLEX61VXJp61hR1x1MDAxNr+S7MuEVO7mXHUwMDEyK9jdw21vf9Oz0J3N9If9Wqv1RFe99XLl5ItcdTAwMTImXGKL1ZuOqSljUFx1MDAwM4FcdTAwMTVcdTAwMDNMiqDqd6Epm6a3R1ZcdTAwMWGAXHUwMDE1pLd/asX6dfh81Yr1M1x1MDAxY7yuY3uRXHUwMDAxfzVAX1x1MDAwMcXkvX9mjfolXHUwMDE0yvTUXHUwMDA2RspKXHUwMDAztkT5lNiT5dNcdTAwMGJGLvd3StbTfXnwXHLJ47zDkC9YhV85XGYl0JRtXHUwMDEzSLnaXG5cXJF891x1MDAwMlx1MDAxOErGZbBcbvbedSNviMJLo15cci7nXHUwMDE3OkpwvfXGQFxccPt1YjFagDtDiZAygZZcdEj3XHUwMDFl+c7XbpVcdTAwMWU0dk6Or1x1MDAxZoA0nM1Vb1x1MDAxYVs5XHUwMDE4WaKSXHUwMDEx51xyjEHST8VcdTAwMDRyXHUwMDE1RfmfaMwzXHUwMDFh0/1TgiFcdTAwMTeYwuz+6fVBu7jfwNVhj1x1MDAxNv2rri/EQ7mZdyxCxFx1MDAxMsy4Tlx1MDAwNzVYXHUwMDBlU1x1MDAwZSdT0FJuXHUwMDA3z+agcoSIcptXsT77UcBcdTAwMTg4ibPh2lu5p8tcdTAwMDSKKyBEkF7mjFSESFx1MDAxMVx1MDAxNNlBWCHl3rC9V/vm75efXHUwMDFhW1x1MDAwN9jc2nnMPVxiQVx1MDAwMoRcdTAwMTisXHUwMDEzhEJcdTAwMTNBwlx1MDAwYlx1MDAwNaX7anAzXHUwMDE1NlNGg7296Fx1MDAxN4tcdTAwMTLfMUhcXHGMOIbHwaOFao/9rYFpblx1MDAxNW+2YXv/XHUwMDFhbmdKpSY0lM4pY8ZcdTAwMTRrkktcdTAwMGWogIBcdTAwMTIm5yjWQizkKpfadII0e1x1MDAwZZKpMNiFgFx0hfPKXHUwMDE2ePo3kDBcdTAwMTF8bYnMUlv61tnU//3z38I4XHUwMDEx0lCQKdw6bkG3rILXqzedjsKI993+blx1MDAwNyeNd0lcdTAwMTdMrzBa2jCahZbhtFxcvdtcdTAwMGVwZlxyv9tccsvxXHUwMDE0V1x1MDAxNXyncGu66s2EvJ6vcGxcdTAwMWKFcs/V65ZROK7fXHUwMDE5XHK/0FVXnN5qeqepXHUwMDAzumvWXHLX1n01loXNk0qh50XV4O0zvpEoZFx0XHUwMDEz8oGHeP1paUjTN8PT0VdW8SVcdTAwMWP/rn9S2ePSvi5f7qGd0kntxnXyZI1S9lx1MDAwNYN0i75qn1x1MDAwM1x1MDAwM6kxKHlQXHUwMDFjyihDc1x1MDAxNtnm7Fx1MDAwNEZcdTAwMDRhjpTl+3V8jmc49Y16oT9cdTAwMDZmXHUwMDAySm/riCwjzyp2XHUwMDAxp1x1MDAwMXbRxlx1MDAxY/VSXHUwMDFhxWlcdTAwMTaiXHUwMDFiXHUwMDAz9qhqnveLT6y8RcTO8eFhvd9cdTAwMWRcdTAwMWPlXHUwMDFlsDM7gNeJWITYaPM+XHUwMDEwmEhcdTAwMDbkvO2PcyArpfI6IFx1MDAwNr9Q4iyAyJXTc1x1MDAwYlaw8XCElNF2efc9sPqiIGtcdTAwMDUpSv+tXHUwMDAyiFx1MDAxNKNKkqWC1sywXHUwMDE1LK9cdTAwMTgtwlx1MDAwNEbhXHUwMDFhI3nlpWhcdTAwMDJLjiDjXG53L5SuTL9cdTAwMGZKUlx1MDAwNlx1MDAwNKBrZ9XJgcTu/Cz7aWPaXHUwMDE2qe1cdTAwMWQp5eRdflxmQPhcdTAwMWIueqdrXHUwMDE52lBhsDimx+mPuUw7rNtcYmRcdTAwMTZmnSk9jNJcdTAwMTecIeBIkcpcdTAwMTLudZteXHUwMDAz+2xIadHePTPBt67z9bSRe0tcdTAwMDCTXHUwMDE5k3Um1lx1MDAxMWBcdTAwMWFcZla5JFx1MDAwNHC+fz3HXHUwMDEysOB7ZWl0R/ObkzVcdTAwMDaKXHUwMDE4XCKm6C2wqtBhPX9HwJ9QxVx1MDAxN2+P0Fx1MDAwNVwiLIXLL1x1MDAxMzu4oXe7VV+N7MZzlevGo2n0t2bV6Lfb0SuwpyNUXHUwMDA38DFGxbE/vvz4P6pcdTAwMGXrtyJ9
  
    
      @font-face {
        font-family: "Virgil";
        src: url("https://excalidraw.com/Virgil.woff2");
      }
      @font-face {
        font-family: "Cascadia";
        src: url("https://excalidraw.com/Cascadia.woff2");
      }
    
    
  
  InternetCloudflareWorkerTabserveChrome browserLocalhostHTTPS requestHTTPS responseWebSocket messageWebSocket messageHTTP requestHTTP response• HTTPS cert for all subdomains• Worker is started geographicallyclosest to first request• One Durable Object per subdomain• WebSocket Hibernation API used• One web worker per subdomain• Your local web serverhttps://example.your-domain.comhttp://localhost:1234](https://tabserve.dev/_astro/sequence-diagram.c049e797.svg)

## Setting up the Cloudflare Worker and DNS

## Limitations

- Limitations that will be fixed in future versions: 
- Responses over 5MB will temporarily block the worker event loop and slow all requests for that particular domain (cached responses fix this).
- No WebSocket requests.
- Some browsers try to sleep/restore browser Web Worker threads - this requires more testing. 
- Chrome and Firefox desktop browsers work fine when left open for many days.
- Only 100-500 RPS.
- No TCP/UDP connections. Tabserve is a HTTP only proxy.
- You must [enable CORS](https://enable-cors.org/server.html) on the localhost web server.
