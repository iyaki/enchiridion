---
title: "The Original Sin of Free Software"
notion_id: da6bfbc7-fa57-4627-ac7d-058e22de40b4
notion_url: https://app.notion.com/p/The-Original-Sin-of-Free-Software-da6bfbc7fa574627ac7d058e22de40b4
last_edited: 2023-08-01T00:17:00.000Z
source_url: https://lipu.dgold.eu/original-sin
tags: ["English", "Reflection", "Programming", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Someone asked me about this very recently so I felt it was worth writing out my thoughts in more detail. Note that this isn’t a manifesto, or a declaration, its just an effort to start putting some structure on my thoughts.

## The History

Since Richard Stallman wrote his GNU Manifesto[[1]](https://lipu.dgold.eu/original-sin#fn:1) in the mid–1980s, it has come to be regarded as the ur-text of Free and Open Software. The principles enunciated were later expanded upon and codified in the Free Software Definition[[2]](https://lipu.dgold.eu/original-sin#fn:2), and have been adopted, to varying degrees, by different organisations, like Debian’s Free Software Guidelines[[3]](https://lipu.dgold.eu/original-sin#fn:3) and the OSI’s Open Source Definition[[4]](https://lipu.dgold.eu/original-sin#fn:4). Before I continue, its probably a good idea to restate the “Four Freedoms” of the FSD:-

1. The freedom to run the program as you wish, for any purpose
2. The freedom to study how the program works, and change it so it does your computing as you wish. Access to the source code is a precondition for this.
3. The freedom to redistribute copies so you can help your neighbor.
4. The freedom to distribute copies of your modified versions to others. By doing this you can give the whole community a chance to benefit from your changes. Access to the source code is a precondition for this.

While these are each expressed as _positive_ freedoms, when the Four Freedoms (I’ll refer to these as 4F from here, and call the individual Freedoms by number, e.g. 4F0) were adopted by Debian and OSI, there was a distinct shift from positive freedoms to _restrictive_ freedoms - most particularly the requirements that the software license:-

_No Discrimination Against Fields of Endeavor_ The license must not restrict anyone from making use of the program in a specific field of endeavor. For example, it may not restrict the program from being used in a business, or from being used for genetic research. _Distribution of License_ The rights attached to the program must apply to all to whom the program is redistributed without the need for execution of an additional license by those parties.

## The Original Sins

There are two separate, albeit related, problems which underlie these apparently positive definitions: Individualism & Capitalism.

In each case, be it the 4F, DSG or the OSI, all positive rights are expressed purely in _individual_ terms. There are no collective rights in these documents, just individual rights. The ideal of the developer as an individual is hard-coded into the DNA of the Free Software Movement and its various children.

There are various reasons for this; Stallman’s original tracts, for example, idealise his struggles with the collapse of the “social club” of MIT’s AI Lab and emphasise his private struggles against the loss of his community, these experiences are the cornerstone of his personal identity. Bruce Perens, father of the Debian Project, holds to the peculiar American cult of the Individualist Libertarian, as does Eric Raymond, the programmer/chronicler who devised the Open Source concept with Perens in the late 1990s - free software without the freedom.

It is important to remember that the American version of Libertarianism, as espoused by Ayn Rand and her ideologues, is nothing to do with the Franco-European tradition of libertarianism, a spectrum of leftist anarchism running from Babeuf through Déjacque to Faure. American Libertarianism should more accurately be described as “anarchist-capitalism”, a strain of pseudo-political thought which idolises the popular concept of the Old West as a high-point of western civilisation, when men were men and justice as dispensed from the barrel of a gun.

It is notable that the sole non-individual right enunciated in the DSG and the OSD is that of the **corporation**. Perens and Raymond were certainly motivated by the tenets of capitalism, the primary purpose of their “Open Source Definition” was the packaging up of “Free Software” to make it acceptable to capitalists and rent-seekers of every stripe. Their collaborators in this, most notably Brian Behlensdorf and Larry Augustin, later seamlessly transitioned to the burgeoning Venture Capitalist industry.

Be it as a result of accident, as with Stallman, or of design, as with Perens/Raymond _et al_, these designers of the “Holy Texts” of Free and Open Software applied no value to the concepts of Community and Commons.

To look again at 4F3 & 4F4, the Freedoms described are explicitly the Freedom for the Individual Programmer to decide whether to share his value-added programs with his community. More than that, in his detailed expansion on the Freedoms, Stallman is at pains to explain that his Freedoms do not enforce any _obligation_ upon the Individual Programmer to release _his_ works. In fact, in his expanded notes, Stallman entitles 4F3 _“The freedom to redistribute, __**if you wish**__”_:-

> 

You should also have the freedom to make modifications and use them privately in your own work or play, without even mentioning that they exist. If you do publish your changes, you should not be required to notify anyone in particular, or in any particular way.

Again, the entire emphasis of Stallman’s concept of Freedom rests on the specific rights of the Individual. Stallman, moreover, is the “hardcore” of this movement, the supposedly non-commercial member of the group. The emphasis placed here on the right to not distribute if you wish may seem trivial, however, it is central to the overall failure of Free and Open Source Software.

## The Resulting Problems

Today, some 35 years after Stallman’s 4Fs, we can see that Free and Open Source Software is, in almost all conceivable ways, triumphant. It has won its struggles against Microsoft, it is the core of almost all smartphones, it lies at the heart of the networked world we all inhabit.

Amidst all this glory, however, there is something utterly broken about the FOSS world. In short, it is one-sided, devoid of economic rights, and an active agent in the oppression of us all.

As a result of the underlying assumptions and conceptions of the FOSS movement, there is no methodology to _require_ the sharing of changes and improvments with the commons. The entire focus of the FOSS movement is specifically geared towards individual rights, the focus of the Open-Source movement is on ensuring that coporations can freely take the product of individual labour and profit from it.

Because of Stallman’s original sin, there is no accepted structure within the FOSS world wherein such obligation to share changes can be enforced by the originator of a piece of software. The DSG, in fact, specifically requires this as a _precondition_ for inclusion in its repositories.

This means that improvements in software, bug-fixes, security-patches and the like, remain the absolute property of the person who has made them, despite the fact that they are, in Newton’s words, _“standing on the shoulders”_ of those who came before. The originator of the software, in other words, is someone who can count herself lucky to be afforded the “exposure” which may come from the future sale of her works by a more powerful entity. The entire function of FOSS is for developers to provide the fruits of their labour, either without recompense or for fractional recompense, to others.

The individual programmer has the rights afforded them by the 4Fs and the OSD, but those rights are inherently solitary. What possibility can there be to enforce those rights, limited as they may be, against the powerhouses of our gilded age? What economic rights accrue to our developers, the people who, in Stallman’s conception, are the origin point of the FOSS movement?

Two recent and pertinent examples can be examined to provide answers: Heartbleed, the security vulnerability in the OpenSSL Library; and GitHub.

OpenSSL, a software library used to secure communications against eavesdropping, was (until Heartbleed) primarily written by one person, Stephen Henson. The OpenSSL Foundation, who employed him, had never had more than $1m in annual funding, and this funding was obtained almost exclusively through third-party contract work which focused on specific deliverable product, **not** on the _“fundamental maintenance and development activities like releases management, code review and refactoring, performance and security, etc.”_[[5]](https://lipu.dgold.eu/original-sin#fn:5) The beneficiaries of OpenSSL’s public nature had donated, on average, some **$2,000** a year for this critical piece of software.[[6]](https://lipu.dgold.eu/original-sin#fn:6)

The Heartbleed, vulnerability, when it became known, was a disaster of almost unimaginable proportions. Entire sectors of the internet were vulnerable to this exploit thanks to their reliance on this crucial underfunded piece of software. Companies like Amazon Web Services (annual sales, $5.11bn), Stripe (market valuation, $245bn) and Cisco Systems (annual income, $50bn) were exposed.

That’s not just standing on the shoulders, that’s standing on the shoulders while the owner of the shoulders is five feet under water. To a great extent, the people who write FOSS software don’t even see this as a problem, they’re happy to have someone shoving them under the water.

Take, as a different example, GitHub - the repository of much of the FOSS world’s source code. Its entire business model is predicated on the willingness of creators to share the product of their labour with the wider world. These creators[[7]](https://lipu.dgold.eu/original-sin#fn:7) offer the product of their labour to the site, solely in exchange for the chance of exposure and discovery. There is no contract with GitHub which can provide payment for this delivery, save the ephemeral back-slap of the “Star” and the prospect of collaboration with others.

Microsoft, when they acquired GitHub in 2018 for $7.5bn[[8]](https://lipu.dgold.eu/original-sin#fn:8), barely mentioned the creators in their carefully-crafted lawyer-approved statement. They placed a value of seven thousand five hundred million dollars on this company, absolutely none of which accrued to the people who offered their labour to that company. Apart from a few cranks[[9]](https://lipu.dgold.eu/original-sin#fn:9), almost no-one protested this. This has become such a feature of the FOSS world that it barely merited attention.

It is simply not sustainable for the commons to be a one-way supplier of intellectual property to massive corporations. In addition, the rights of the programmer must include the right to be _paid_ for the work of programming and for value derived from that work.

As things are, FOSS not only doesn’t provide value to the writers (by recompense for their labours), it only delivers value to those who are best placed to _exploit_ this labour. These exploiters, thanks to the dogmatism of the FOSS world, are under no positive obligation to provide any recompense.

If you are a lone programmer, working in a field which you understand implicitly, then your works will be taken by corporate entities and monetised. As long as they do not sell that _software_ (narrowly defined) then they are under no obligation to note your contribution. If they utilise it internally as an organisation, then they don’t even have to share their changes (security patches, general improvements) with anyone.

In the case of permissive licenses (like the MIT License) the few protections afforded the software source evaporate completely. Even with strong copyleft licensing in place, the enforcement of that copyleft license is of course left entirely to the programmer. While there are some examples of assistance being provided by the Software Freedom Conservancy and others, these are not ordinarily available to the average GitHub user.

The SFC, for example, will only act in the case of software “member projects”. Just to become a member project, a developer/organisation is required to license their software under a FSD/OSD License _and_ to license all documentation under either that license _or_ CC-BY-SA.[[10]](https://lipu.dgold.eu/original-sin#fn:10) Again, we see the institutions of FOSS acting primarily in the interests of free labour and expropriation.

The imbalance of resources between that user and, say, Amazon or Google, renders any chance of enforcing a copyleft restriction so small as to be non-existent. Even if the evidentiary standard is met, the costs of litigation are prohibitive.

Finally, FOSS is readily available for use by the very worst elements of our techno-dystopian world. Again, deriving from the Original Sin of the Free Software movement, in particular 4F0:-

> 

The freedom to run the program as you wish, **for any purpose**.

This is a gaping sore at the heart of the FOSS world. “For any purpose” encompasses a multitude - be that the use of FOSS by military organisations, such as the USAF’s “Global Hawk” drone program; or by surveillance operations, like GCHQ or the NSA; or by face-recognition programs in public places.

It is the stated position of the FOSS movement that all of these uses are _fine_, and all are to be welcomed. Usually, this declaration is accompanied with some expert deployment of the consequentialist fallacy. Like the tech-giants and hate speech, the reasoning proceeds that if we do anything to stop it, then we have to stop everything. Better then to do **nothing**.

Any attempt to question this fallacious reasoning usually results in a cavalcade of FOSS-advocates attacking the querist. It is clear that this issue triggers such behaviour at a deep, fundamental level. This is hardly surprising, after all, anyone who is in deep denial will instinctively lash out against any impingement of reality into their delusions, and so it is with FOSS.

## Towards an Alternative

In these circumstances, what then are the features of a license which I feel would offer support to the values I hold? I’d prefer to have a different set of Freedoms, which emphasise the importance of community and collaboration, which require the sharing of improvements, which are deliberately structured so as to prevent the expropriation of labour, and which have certain ethical standards at their core.

These then are my proposals for a new set of freedoms:-

1. 

The freedom for any person to decide to limit the use of their software for any purpose or category of purposes shall not be restricted, subject only to the requirements of the principles of equality and non-discrimination.

1. 

The freedom for any natural person to run the program as they wish, subject to the first freedom.

1. 

The freedom to study how the program works, and change it so it does your computing as you wish. Access to the source code is a requirement of this freedom.

1. 

The right of the commons to be provided with changes to the software is essential to ensuring the sustainability and security of the commons. Provision of all changes to the commons, in so far as can be permitted by circumstances, is required.

1. 

The freedom for any writer to decide that availability of the program to persons other than provided in the first four freedoms is a matter which is outside the scope of the license, and is properly subject to a contract of the writer’s choice being entered into by those persons.

Those are just ideas, ones that I’ve been kicking the tyres of in my own time. I believe that these five freedoms provide a conceptual framework which can

- Assure the viability of the commons
- Provide for the economic rights of the programmer
- Permit ethical design

## Request for Comments

As of this writing[[11]](https://lipu.dgold.eu/original-sin#fn:11), I’ve identified five licenses which meet some, though by no means all, of my criteria:-

- 

Dmitri Kleyner’s Peer Production License [[link](http://wiki.p2pfoundation.net/Peer_Production_License)] describes itself as “copyfarleft”, forbids use by commercial entities, and is heavily based on the Creative Commons “No Commercial” License variants.

- 

The Cooperative Software License [[link](https://coinsh.red/c/csl.txt)] is derived from the PPL, and attempts to shift the focus yet further into software.

- 

The Fully Open Public License [[link](https://github.com/pjakma/fopl)] is not a “copyfarleft” license, but it is an attempt to place source-provision in a license context, and is deserving of attention.

- 

Kyle Mitchell’s LicenseZero project has the Prosperity License [[link](https://licensezero.com/licenses/prosperity)], which is a permissive license which preserves the right of the programmer to be paid.

- 

The same project’s Parity License [[link](https://licensezero.com/licenses/parity)] seeks to preserve the rights of the commons, but the two are not combined. These licenses have the singular virtue of being drafted in a far more modern legal syntax.

Obviously, I’d be delighted to learn of other licenses which start to meet my Five Freedoms.

Please feel free to [contact me with any comments or questions](mailto:comments@lipu.dgold.eu?subject=original-sin) about anything in this essay. This essay isn't performative, so I'm not displaying comments here, but I value interaction and will be delighted to engage positively with you. If I've gotten something completely wrong, tell me!

1. 

The GNU Operating System _“The Gnu Manifesto”_ [https://www.gnu.org/gnu/manifesto.html](https://www.gnu.org/gnu/manifesto.html) Version: 2015/06/02 12:55:15 [accessed 2019–03–28] [ ↩](https://lipu.dgold.eu/original-sin#fnref:1)

1. 

The GNU Operating System _“What is Free Software”_ [https://www.gnu.org/philosophy/free-sw.en.html](https://www.gnu.org/philosophy/free-sw.en.html) Version: 2019/03/20 10:56:16 [accessed 2019–03–28] [ ↩](https://lipu.dgold.eu/original-sin#fnref:2)

1. 

The Debian Project _“The Debian Social Contract”_ [https://www.debian.org/social_contract#guidelines](https://www.debian.org/social_contract#guidelines) Version 1.1 2004/04/25 [accessed 2019–03–28] [ ↩](https://lipu.dgold.eu/original-sin#fnref:3)

1. 

The Open Source Initiative _“The Open Source Definition”_ [https://opensource.org/docs/osd](https://opensource.org/docs/osd) [accessed 2019–03–28] [ ↩](https://lipu.dgold.eu/original-sin#fnref:4)

1. 

Steve Marquess _“Of Money, Responsibility, and Pride”_ [ https://veridicalsystems.com/blog/of-money-responsibility-and-pride/index.html](https://veridicalsystems.com/blog/of-money-responsibility-and-pride/index.html) 2017 [accessed 2019–03–28] [ ↩](https://lipu.dgold.eu/original-sin#fnref:5)

1. 

Marquess, _ibid._ [ ↩](https://lipu.dgold.eu/original-sin#fnref:6)

1. 

disclosure: I once counted myself among them [ ↩](https://lipu.dgold.eu/original-sin#fnref:7)

1. 

Microsoft News Center _“Microsoft to acquire GitHub for $7.5 billion”_ [ https://news.microsoft.com/2018/06/04/microsoft-to-acquire-github-for-7-5-billion/](https://news.microsoft.com/2018/06/04/microsoft-to-acquire-github-for-7-5-billion/) 2018/06/04 [accessed 2019–03–28] [ ↩](https://lipu.dgold.eu/original-sin#fnref:8)

1. 

again, I was one of them [ ↩](https://lipu.dgold.eu/original-sin#fnref:9)

1. 

The Software Freedom Conservancy _“Applying to Join Conservancy as a Member Project”_ [https://sfconservancy.org/projects/apply/](https://sfconservancy.org/projects/apply/) 2018–05–25 [accessed 2019–03–28] [ ↩](https://lipu.dgold.eu/original-sin#fnref:10)
