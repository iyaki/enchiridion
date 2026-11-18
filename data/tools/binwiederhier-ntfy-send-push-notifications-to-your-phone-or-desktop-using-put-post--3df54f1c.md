---
title: "binwiederhier/ntfy: Send push notifications to your phone or desktop using PUT/POST"
notion_id: 3df54f1c-7d23-81d4-8145-da1a43618c3b
notion_url: https://app.notion.com/p/binwiederhier-ntfy-Send-push-notifications-to-your-phone-or-desktop-using-PUT-POST-3df54f1c7d2381d48145da1a43618c3b
last_edited: 2026-09-18T01:19:00.000Z
source_url: https://github.com/binwiederhier/ntfy
tags: ["English", "Web Development", "Notification", "Open Source", "DevOps", "Automation", "Tool", "Service", "GitHub"]
---
---

![image](https://github.com/binwiederhier/ntfy/raw/main/web/public/static/images/ntfy.png)

![image](https://camo.githubusercontent.com/c9b062ec22f04cbe5db8f9380c36479cb16c09f928f1384c0926d091751694df/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f72656c656173652f62696e776965646572686965722f6e7466792e7376673f636f6c6f723d73756363657373267374796c653d666c61742d737175617265)

![image](https://camo.githubusercontent.com/0cd449b85bb265b2fd2ff969a79b28287b819bd3d5c7f262b36b801f9f134bb5/68747470733a2f2f706b672e676f2e6465762f62616467652f6865636b656c2e696f2f6e7466792e737667)

![image](https://github.com/binwiederhier/ntfy/workflows/test/badge.svg)

![image](https://camo.githubusercontent.com/e44b67759b5df5e10fae37e771aa2f85361060c8465aee44268a34be71db07c4/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f62696e776965646572686965722f6e746679)

![image](https://camo.githubusercontent.com/5e84814cdcfd41a9ade478ee4bbec831f6e84c0251778847fca5bc616d415767/68747470733a2f2f636f6465636f762e696f2f67682f62696e776965646572686965722f6e7466792f6272616e63682f6d61696e2f67726170682f62616467652e7376673f746f6b656e3d413539374b5134363347)

![image](https://camo.githubusercontent.com/8a4449c246f7879e97147acedc49b11ccb86a21182712af9fa3d2a8e7ce48eec/68747470733a2f2f696d672e736869656c64732e696f2f646973636f72642f3837343339383636313730393239353632363f6c6162656c3d446973636f7264)

![image](https://camo.githubusercontent.com/fe617100675e43cc23023bf36650a059b22ec1881f9f6dff5eb498869993145a/68747470733a2f2f696d672e736869656c64732e696f2f6d61747269782f6e7466793a6d61747269782e6f72673f6c6162656c3d4d6174726978)

![image](https://camo.githubusercontent.com/ddd9e9cfa3cbe4854acee1b11845f5ac8bb71983eb0bfbfd1cdb7709dd397a21/68747470733a2f2f696d672e736869656c64732e696f2f6d61747269782f6e7466792d73706163653a6d61747269782e6f72673f6c6162656c3d4d61747269782b7370616365)

![image](https://camo.githubusercontent.com/3290bd448b9226d36766be89d1d991b6cbb62defb1e47a55ddfcfe8e08fe603e/68747470733a2f2f6865616c7468636865636b732e696f2f62616467652f36386236353937362d623362302d343130322d616563392d3938303932312f6b636f4567724c592e737667)

![image](https://camo.githubusercontent.com/d262538135adaf89726a6c0813405f23df5c85990e758c22037048cb4d4b01fe/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f436f6e74726962757465253230776974682d476974706f642d3930386138353f6c6f676f3d676974706f64)

**ntfy** (pronounced "_notify_") is a simple HTTP-based [pub-sub](https://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern)
notification service. With ntfy, you can **send notifications to your phone or desktop via scripts** from any computer,
**without having to sign up or pay any fees**. If you'd like to run your own instance of the service, you can easily do
so since ntfy is open source.

You can access the free version of ntfy at [**ntfy.sh**](https://ntfy.sh/). There is also an [open-source Android app](https://github.com/binwiederhier/ntfy-android)
available on [Google Play](https://play.google.com/store/apps/details?id=io.heckel.ntfy) or [F-Droid](https://f-droid.org/en/packages/io.heckel.ntfy/),
as well as an [open source iOS app](https://github.com/binwiederhier/ntfy-ios) available on the [App Store](https://apps.apple.com/us/app/ntfy/id1625396347).

![image](https://github.com/binwiederhier/ntfy/raw/main/docs/static/img/badge-googleplay.png)

![image](https://github.com/binwiederhier/ntfy/raw/main/docs/static/img/badge-fdroid.svg)

![image](https://github.com/binwiederhier/ntfy/raw/main/docs/static/img/badge-appstore.png)

![image](https://github.com/binwiederhier/ntfy/raw/main/.github/images/screenshot-curl.png)

![image](https://github.com/binwiederhier/ntfy/raw/main/.github/images/screenshot-web-detail.png)

![image](https://github.com/binwiederhier/ntfy/raw/main/.github/images/screenshot-phone-main.jpg)

![image](https://github.com/binwiederhier/ntfy/raw/main/.github/images/screenshot-phone-detail.jpg)

![image](https://github.com/binwiederhier/ntfy/raw/main/.github/images/screenshot-phone-notification.jpg)

## [ntfy Pro](https://ntfy.sh/app) 💸 🎉

I now offer paid plans for [ntfy.sh](https://ntfy.sh/) if you don't want to self-host, or you want to support the development of
ntfy (→ [Purchase via web app](https://ntfy.sh/app)). You can **buy a plan for as low as $5/month**.
You can also donate via [GitHub Sponsors](https://github.com/sponsors/binwiederhier), and [Liberapay](https://liberapay.com/ntfy).
I would be very humbled by your sponsorship. ❤️

## [**Documentation**](https://ntfy.sh/docs/)

[Getting started](https://ntfy.sh/docs/) |
[Android/iOS](https://ntfy.sh/docs/subscribe/phone/) |
[API](https://ntfy.sh/docs/publish/) |
[Install / Self-hosting](https://ntfy.sh/docs/install/) |
[Building](https://ntfy.sh/docs/develop/)

## Chat/forum

There are a few ways to get in touch with me and/or the rest of the community. Feel free to use any of these methods. Whatever
works best for you:

- [Discord server](https://discord.gg/cT7ECsZj9w) - direct chat with the community
- [Matrix room #ntfy](https://matrix.to/#/#ntfy:matrix.org) (+ [Matrix space](https://matrix.to/#/#ntfy-space:matrix.org)) - same chat, bridged from Discord
- [GitHub issues](https://github.com/binwiederhier/ntfy/issues) - questions, features, bugs

## Announcements/beta testers

For announcements of new releases and cutting-edge beta versions, please subscribe to the [ntfy.sh/announcements](https://ntfy.sh/announcements)
topic. If you'd like to test the iOS app, join [TestFlight](https://testflight.apple.com/join/P1fFnAm9). For Android betas,
join Discord/Matrix (I'll eventually make a testing channel in Google Play).

## Sponsors

If you'd like to support the ntfy maintainers, please consider donating to [GitHub Sponsors](https://github.com/sponsors/binwiederhier) or
and [Liberapay](https://liberapay.com/ntfy). We would be humbled if you helped carry the server and developer
account costs. Even small donations are very much appreciated.

Thank you to our commercial sponsors, who help keep the service running and the development going:

![image](https://camo.githubusercontent.com/bbed225b7b0ffa5d4c6c991ce9c155447132f44a168b0bc4571931a915bf3ab4/68747470733a2f2f6f70656e736f757263652e6e7963332e63646e2e6469676974616c6f6365616e7370616365732e636f6d2f6174747269627574696f6e2f6173736574732f5356472f444f5f4c6f676f5f686f72697a6f6e74616c5f626c75652e737667)

![image](https://raw.githubusercontent.com/warpdotdev/brand-assets/refs/heads/main/Logos/Warp-Wordmark-Black.png)

And a big fat **Thank You** to the individuals who have sponsored ntfy in the past, or are still sponsoring ntfy:

![image](https://github.com/neutralinsomniac.png)

![image](https://github.com/aspyct.png)

![image](https://github.com/nickexyz.png)

![image](https://github.com/qcasey.png)

![image](https://github.com/mckay115.png)

![image](https://github.com/Salamafet.png)

![image](https://github.com/codinghipster.png)

![image](https://github.com/HinFort.png)

![image](https://github.com/Lexevolution.png)

![image](https://github.com/johnnyip.png)

![image](https://github.com/JonDerThan.png)

![image](https://github.com/12nick12.png)

![image](https://github.com/eanplatter.png)

![image](https://github.com/fnoelscher.png)

![image](https://github.com/bnorick.png)

![image](https://github.com/snh.png)

![image](https://github.com/hen-x.png)

![image](https://github.com/JamieGoodson.png)

![image](https://github.com/cremesk.png)

![image](https://github.com/dangowans.png)

![image](https://github.com/mnault.png)

![image](https://github.com/nwithan8.png)

![image](https://github.com/peterleiser.png)

![image](https://github.com/portothree.png)

![image](https://github.com/finngreig.png)

![image](https://github.com/skrollme.png)

![image](https://github.com/gergepalfi.png)

![image](https://github.com/tonyakwei.png)

![image](https://github.com/crosbyh.png)

![image](https://github.com/mdlnr.png)

![image](https://github.com/p-samuel.png)

![image](https://github.com/zugaldia.png)

![image](https://github.com/NathanSweet.png)

![image](https://github.com/msdeibel.png)

![image](https://github.com/ksurl.png)

![image](https://github.com/CodingTimeDEV.png)

![image](https://github.com/Terrormixer3000.png)

![image](https://github.com/voroskoi.png)

![image](https://github.com/Nickwasused.png)

![image](https://github.com/bahur142.png)

![image](https://github.com/vinhdizzo.png)

![image](https://github.com/Ge0rg3.png)

![image](https://github.com/biopsin.png)

![image](https://github.com/thebino.png)

![image](https://github.com/sky4055.png)

![image](https://github.com/julianlam.png)

![image](https://github.com/andreapx.png)

![image](https://github.com/billycao.png)

![image](https://github.com/zoic21.png)

![image](https://github.com/IanKulin.png)

![image](https://github.com/Joachim256.png)

![image](https://github.com/overtone1000.png)

![image](https://github.com/oakd.png)

![image](https://github.com/KucharczykL.png)

![image](https://github.com/hansbickhofe.png)

![image](https://github.com/caseodilla.png)

![image](https://github.com/0xAF.png)

![image](https://github.com/soonoo.png)

![image](https://github.com/nichu42.png)

![image](https://github.com/samliebow.png)

![image](https://github.com/johman10.png)

![image](https://github.com/R-Gld.png)

![image](https://github.com/FingerlessGlov3s.png)

![image](https://github.com/Twisterado.png)

![image](https://github.com/ScrumpyJack.png)

![image](https://github.com/andrejarrell.png)

![image](https://github.com/oaustegard.png)

![image](https://github.com/CreativeWarlock.png)

![image](https://github.com/darkdragon-001.png)

![image](https://github.com/jonathan-kosgei.png)

![image](https://github.com/KevinWang15.png)

![image](https://github.com/darkmattercoder.png)

![image](https://github.com/bmcgonag.png)

![image](https://github.com/skorokithakis.png)

![image](https://github.com/eenturk.png)

![image](https://github.com/spirossi.png)

![image](https://github.com/teomarcdhio.png)

![image](https://github.com/MarcMichalsky.png)

![image](https://github.com/LuckVintage.png)

![image](https://github.com/spartan.png)

![image](https://github.com/alexandzors.png)

![image](https://github.com/dkramer95.png)

![image](https://github.com/YezGotIt.png)

![image](https://github.com/thomasskou.png)

![image](https://github.com/surfernv.png)

![image](https://github.com/richardleach.png)

![image](https://github.com/bear.png)

![image](https://github.com/cminter.png)

![image](https://github.com/pgwiebes.png)

![image](https://github.com/ralhei.png)

![image](https://github.com/TechMDW.png)

![image](https://github.com/ubipo.png)

![image](https://github.com/tka85.png)

![image](https://github.com/beekeeb.png)

![image](https://github.com/Emiliaaah.png)

![image](https://github.com/zark0s.png)

![image](https://github.com/tomershvueli.png)

![image](https://github.com/CataIana.png)

![image](https://github.com/ajay-actuary.png)

![image](https://github.com/mursec.png)

![image](https://github.com/FrameXX.png)

![image](https://github.com/vovayartsev.png)

![image](https://github.com/dwain-lab.png)

![image](https://github.com/brookmg.png)

![image](https://github.com/siebej.png)

![image](https://github.com/rxsantos.png)

![image](https://github.com/hermannx5.png)

![image](https://github.com/rwxd.png)

![image](https://github.com/Integral-Tech.png)

![image](https://github.com/TheTomik1.png)

![image](https://github.com/dav23r.png)

![image](https://github.com/stannynuytkens.png)

![image](https://github.com/danbartram.png)

![image](https://github.com/arthurgleckler.png)

![image](https://github.com/tomroth04.png)

![image](https://github.com/Circenn5130.png)

![image](https://github.com/jceloria.png)

![image](https://github.com/afunworm.png)

![image](https://github.com/PTR-inc.png)

![image](https://github.com/spudooli.png)

![image](https://github.com/IMarkoMC.png)

![image](https://github.com/rubund.png)

![image](https://github.com/Riolku.png)

![image](https://github.com/arnbrhm.png)

![image](https://github.com/herzkerl.png)

![image](https://github.com/0x45796164.png)

![image](https://github.com/madchr1st.png)

![image](https://github.com/avalentic.png)

![image](https://github.com/TheCraiggers.png)

![image](https://github.com/sheetd.png)

![image](https://github.com/dlt-green.png)

![image](https://github.com/suhlig.png)

![image](https://github.com/Proximus888.png)

![image](https://github.com/wielandp.png)

![image](https://github.com/chxseh.png)

![image](https://github.com/user8446.png)

![image](https://github.com/cdf-eagles.png)

## Contributing

I welcome any contributions. Just create a PR or an issue. For larger features/ideas, please reach out
on Discord/Matrix first to see if I'd accept them. To contribute code, check out the [build instructions](https://ntfy.sh/docs/develop/)
for the server and the Android app. Or, if you'd like to help translate 🇩🇪 🇺🇸 🇧🇬, you can start immediately in
[Hosted Weblate](https://hosted.weblate.org/projects/ntfy/).

![image](https://camo.githubusercontent.com/c698d279e959946fd8cdd69e94fd684902f4d9db8d58d8788feb1d71ce703dd4/68747470733a2f2f686f737465642e7765626c6174652e6f72672f776964676574732f6e7466792f2d2f6d756c74692d626c75652e737667)

## Code of Conduct

We as members, contributors, and leaders pledge to make participation in our community a harassment-free experience for
everyone, regardless of age, body size, visible or invisible disability, ethnicity, sex characteristics, gender identity
and expression, level of experience, education, socio-economic status, nationality, personal appearance, race, caste,
color, religion, or sexual identity and orientation.

**We pledge to act and interact in ways that contribute to an open, welcoming, diverse, inclusive, and healthy community.**

_Please be sure to read the complete _[_Code of Conduct_](https://github.com/binwiederhier/ntfy/blob/main/CODE_OF_CONDUCT.md)_._

## License

Made with ❤️ by [Philipp C. Heckel](https://heckel.io/).

The project is dual licensed under the [Apache License 2.0](https://github.com/binwiederhier/ntfy/blob/main/LICENSE) and the [GPLv2 License](https://github.com/binwiederhier/ntfy/blob/main/LICENSE.GPLv2).

Third-party libraries and resources:

- [github.com/urfave/cli](https://github.com/urfave/cli) (MIT) is used to drive the CLI
- [Mixkit sounds](https://mixkit.co/free-sound-effects/notification/) (Mixkit Free License) are used as notification sounds
- [Sounds from notificationsounds.com](https://notificationsounds.com/) (Creative Commons Attribution) are used as notification sounds
- [Roboto Font](https://fonts.google.com/specimen/Roboto) (Apache 2.0) is used as a font in everything web
- [React](https://reactjs.org/) (MIT) is used for the web app
- [Material UI components](https://mui.com/) (MIT) are used in the web app
- [MUI dashboard template](https://github.com/mui/material-ui/tree/master/docs/data/material/getting-started/templates/dashboard) (MIT) was used as a basis for the web app
- [Dexie.js](https://github.com/dexie/Dexie.js) (Apache 2.0) is used for web app persistence in IndexedDB
- [GoReleaser](https://goreleaser.com/) (MIT) is used to create releases
- [go-smtp](https://github.com/emersion/go-smtp) (MIT) is used to receive e-mails
- [stretchr/testify](https://github.com/stretchr/testify) (MIT) is used for unit and integration tests
- [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) (MIT) is used to provide the persistent message cache
- [Firebase Admin SDK](https://github.com/firebase/firebase-admin-go) (Apache 2.0) is used to send FCM messages
- [github/gemoji](https://github.com/github/gemoji) (MIT) is used for emoji support (specifically the [emoji.json](https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json) file)
- Go's [text/template](https://pkg.go.dev/text/template) (BSD-3-Clause) is vendored under [template/gotext/](https://github.com/binwiederhier/ntfy/blob/main/template/gotext) with a small patch adding an execution deadline (see [template/gotext/README.md](https://github.com/binwiederhier/ntfy/blob/main/template/gotext/README.md))
- [Lightbox with vanilla JS](https://yossiabramov.com/blog/vanilla-js-lightbox) as a lightbox on the landing page
- [HTTP middleware for gzip compression](https://gist.github.com/CJEnright/bc2d8b8dc0c1389a9feeddb110f822d7) (MIT) is used for serving static files
- [Regex for auto-linking](https://github.com/bryanwoods/autolink-js) (MIT) is used to highlight links (the library is not used)
- [Statically linking go-sqlite3](https://www.arp242.net/static-go.html)
- [Linked tabs in mkdocs](https://facelessuser.github.io/pymdown-extensions/extensions/tabbed/#linked-tabs)
- [webpush-go](https://github.com/SherClockHolmes/webpush-go) (MIT) is used to send web push notifications
- [Sprig](https://github.com/Masterminds/sprig) (MIT) is used to add template parsing functions
