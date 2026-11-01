---
title: "How to Use AI to Do Stuff: An Opinionated Guide"
notion_id: e45c5ce4-87bc-43b9-abd7-d0ff578359d3
notion_url: https://app.notion.com/p/How-to-Use-AI-to-Do-Stuff-An-Opinionated-Guide-e45c5ce487bc43b9abd7d0ff578359d3
last_edited: 2023-07-21T11:59:00.000Z
source_url: https://www.oneusefulthing.org/p/how-to-use-ai-to-do-stuff-an-opinionated
tags: ["English", "Productivity", "Producer (Individual Contributor)", "Artificial Intelligence (AI)", "Article", "One Useful Thing Substack"]
---
Increasingly powerful AI systems are being released at an increasingly rapid pace. This week saw the debut of Claude 2, likely the second most capable AI system available to the public. The week before, Open AI released Code Interpreter, the most sophisticated mode of AI yet available. The week before that, some AIs [got the ability to see images](https://www.oneusefulthing.org/p/on-giving-ai-eyes-and-ears).

And yet not a single AI lab seems to have provided any user documentation. Instead, the only user guides out there appear to be Twitter influencer threads. Documentation-by-rumor is a weird choice for organizations claiming to be concerned about proper use of their technologies, but here we are.

I can’t claim that this is going to be a complete user guide, but it will serve as a bit of orientation to the current state of AI. I have been putting together a Getting Started Guide to AI for my students (and interested readers) every few months, and each time, it requires major modifications. The last couple of months have been particularly insane.

This guide is opinionated, based on my experience, and focused on how to pick the right tool to do things. I have written separately about [the kinds of tasks you may want AI to do](https://www.oneusefulthing.org/p/on-boarding-your-ai-intern), which might be useful to read first.

# The Major Large Language Models

When we talk about AI right now, we are usually talking about Large Language Models, or LLMs. Most AI applications are powered by LLMs, of which there are just a few Foundation Models, created by a handful of organizations. Each company gives direct access to their models via a Chatbot: OpenAI makes GPT-3.5 and GPT-4, which power [ChatGPT ](https://chat.openai.com/)and Microsoft’s [Bing ](https://www.bing.com/search?q=Bing%20AI&showconv=1&FORM=hpcodx&sydconv=1)(access it on an Edge browser). Google has a variety of models under the label of [Bard](https://bard.google.com/). And Anthropic makes Claude and [Claude 2](https://claude.ai/).

There are other LLMs I won’t be discussing. The first is [Pi](https://pi.ai/talk), a chatbot built by Inflection. Pi is optimized for conversation, and really, really wants to be your friend (seriously, try it to see what I mean). It does not like to do much besides chat, and trying to get it to do work for you is an exercise in frustration. We also won’t cover the variety of open source models that anyone can use and modify. They are generally not accessible or useful for the casual user today, but have real promise. Future guides may include them.

So here is your quick reference chart, summarizing the state of LLMs:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The first four (including Bing) are all OpenAI systems. There are basically two major OpenAI AIs today: 3.5 and 4. The 3.5 model kicked off the current AI craze in November, the 4 model premiered in the Spring and is much more powerful. A new variation uses plugins to connect to the internet and other apps. There are a lot of plugins, most of which are not very useful, but you should feel free to explore them as needed. Code Interpreter as is an extremely powerful version of ChatGPT that can run Python programs. If you have never paid for OpenAI, you have only used 3.5. Aside from the plugins variation, and a temporarily suspended version of GPT-4 with browsing, none of these models are connected to the internet. Microsoft’s Bing uses a mix of 4 and 3.5, and is usually the first model in the GPT-4 family to roll out new features. For example, it can both create and view images, and it can read documents in the web browser. It is connected to the internet. [Bing is a bit weird to use, but powerful.](https://oneusefulthing.substack.com/p/power-and-weirdness-how-to-use-bing)

Google has been testing its own AI for consumer use, which they call Bard, but which is powered by a variety of Foundation Models, most recently one called PaLM 2. For the company that developed LLM technology, they have been pretty disappointing, although improvements announced yesterday show they are still working on the underlying technology, so I have hope. It has already gained the capability to run limited code and interpret images, but I would generally avoid it for now.

The final company, Anthropic has released Claude 2. Claude is most notable for having a very large context window - essentially the memory of the LLM. Claude can hold almost an entire book, or many PDFs, in memory. It has been built to be less likely to act maliciously than other Large Language Models, which means, practically, that it tends to scold you a bit about stuff.

Now, on to some uses:

# Write stuff

**Best free options:** [Bing ](https://www.bing.com/search?q=Bing%20AI&showconv=1&FORM=hpcodx)and [Claude 2](https://claude.ai/)

**Paid option: **[ChatGPT](https://chat.openai.com/chat) 4.0/ChatGPT with plugins

For right now, GPT-4 is still the most capable AI tool for writing, which you can access at Bing (select“creative mode”) for free or by purchasing a $20/month subscription to ChatGPT. Claude, however, is a close second, and has a limited free option available.

These tools are also being integrated directly into common office applications. Microsoft Office will include a copilot powered by GPT and Google Docs will integrate suggestions from Bard. [The implications of what these new innovations mean for writing are pretty profound.](https://www.oneusefulthing.org/p/setting-time-on-fire-and-the-temptation)

Here are some ways to use AI to help you write.

- Writing drafts of anything. Blog posts, essays, promotional material, speeches, lectures, chose-you-own adventures, scripts, short stories - you name it, AI does it, and pretty well. All you have to do is prompt it. Prompt crafting is not magic, but basic prompts result in boring writing, [but getting better at prompting is not that hard, just work interactively with the system.](https://www.oneusefulthing.org/p/on-boarding-your-ai-intern) You will find AI systems to be much more capable as writers with a little practice.
- Make your writing better. Paste your text into an AI. Ask it to improve the content, or for suggestions about how to make it better for a particular audience. Ask it to create 10 drafts in radically different styles. Ask it to make things more vivid, or add examples. Use it to inspire you to do better work.
- Help you with tasks. AI can do things you don’t have the time to do. Use it like an intern to write emails, create sales templates, give you next steps in a business plan, and a lot more. [Here is what I could accomplish with it in 30 minutes in supporting a product launch.](https://oneusefulthing.substack.com/p/superhuman-what-can-ai-do-in-30-minutes)
- [Unblock yourself.](https://oneusefulthing.substack.com/p/how-to-use-ai-to-unstick-yourself) It is very easy to get distracted from a task by one difficult challenge. AI provides a way of giving yourself momentum.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

_Some things to worry about: _In a bid to respond to your answers, it is very easy for the AI to “hallucinate” and generate plausible facts. It can generate entirely false content that is utterly convincing. Let me emphasize that: _**AI lies continuously and well**_. Every fact or piece of information it tells you may be incorrect. You will need to check it all. Particularly dangerous is asking it for references, quotes, citations, and information for the internet (for the models that are not connected to the internet). Bing will usually hallucinate less than other models, because GPT-4 is generally more grounded and because Bing’s internet connection means it can actually pull in relevant facts. [Here is a guide to avoiding hallucinations](https://oneusefulthing.substack.com/p/how-to-get-an-ai-to-lie-to-you-in), but they are impossible to completely eliminate.

And also note that AI doesn’t explain itself, it only makes you think it does. If you ask it to explain why it wrote something, it will give you a plausible answer that is completely made up. When you ask it for its thought process, is not interrogating its own actions, it is just generating text that sounds like it is doing so. This makes understanding biases in the system very challenging, even though those biases almost certainly exist.

It also can be used unethically to manipulate or cheat. You are responsible for the output of these tools.

# Make images

**Most transparent option: **[Adobe Firefly](https://www.adobe.com/sensei/generative-ai/firefly.html)**
Open Source Option:** [Stable Diffusion](https://stable-diffusion-ui.github.io/)**Best free option:** Bing or [Bing Image Creator](https://www.bing.com/images/create) (which uses DALL-E), [Playgound ](https://playgroundai.com/)(which lets you use multiple models)**Best quality images:** [Midjourney](https://midjourney.com/)

There are four big image generators available for most people:

1. Stable Diffusion, which is open source and you can run from any high-end computer. It takes effort to get started, since you have to learn to craft prompts properly, but once you do it can produce great results. It is especially good for combining AI with images from other sources. [Here is a nice guide to Stable Diffusion if you go that route (be sure to read both parts 1 and part 2).](https://www.jonstokes.com/p/stable-diffusion-20-and-21-an-overview)
2. DALL-E, from OpenAI, which is incorporated into Bing (you have to use creative mode) and Bing image creator. This system is solid, but worse than Midjourney.
3. Midjourney, which is the best system in mid-2023. It has the lowest learning-curve of any system: just type in "thing-you-want-to-see --v 5.2" (the --v 5.2 at the end is important, it uses the latest model) and you get a great result. Midjourney requires Discord.[ Here is a guide to ](https://www.pcworld.com/article/540080/how-to-use-discord-a-beginners-guide.html)using Discord.
4. Adobe Firefly, built into a variety of Adobe products, but it lags DALL-E and Midjourney in terms of quality. However, while the other two models have been unclear about the source images that they used to train their AIs, Adobe has declared that it is only using images it has the right to use.

Here is how they compare (each image is labelled with the model):

Prompt: “Fashion photoshoot of sneakers inspired by Van Gogh” - the first images that were created by each model

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

_Some things to worry about:_ These systems are built around models that have built-in biases due to their training on Internet data (if you ask it to create a picture of an entrepreneur, for example, you will likely see more pictures featuring men than women, unless you specify “female entrepreneur”), you can use [this explorer](https://huggingface.co/spaces/society-ethics/DiffusionBiasExplorer) to see these biases at work.

These systems are also trained on existing art on the internet in ways that are not transparent and [potentially legally and ethically questionable](https://www.nytimes.com/2022/09/02/technology/ai-artificial-intelligence-artists.html). Though technically you own copyright of the images created, legal rules are still hazy.

Also, right now, they don’t create text, just a bunch of stuff that looks like text. But Midjourney has nailed hands.

# Come up with ideas

**Best free option:** [Bing](https://www.bing.com/search?q=Bing%20AI&showconv=1&FORM=hpcodx)

**Paid option: **[ChatGPT](https://chat.openai.com/chat) 4.0, but Bing is likely better because of its internet connections

Despite of (or in fact, because of) all its constraints and weirdness, AI is perfect for idea generation. You often need to have a lot of ideas to have good ideas, and AI is good at volume. With the right prompting, you can also force it to be very creative. Ask Bing in creative mode to look up your favorite unusual idea generation techniques, like Brian Eno's oblique strategies or Mashall McLuhan's tetrads, and apply them. Or ask for something weird, like ideas inspired by a random patent, or your favorite superhero…

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# Make videos

**Best animation tool:** [D-i](https://www.d-id.com/)D for animating faces in videos. [Runway v2](https://app.runwayml.com/) for creating videos from text

**Best voice cloning:** [ElevenLabs](https://beta.elevenlabs.io/speech-synthesis)

It is now trivial to generate a video with a completely AI generated character, reading a completely AI-written script, talking in an AI-made voice, animated by AI. [It can also deepfake people, as you can see in this link where I deepfaked myself. Instructions and more information here.](https://oneusefulthing.substack.com/p/a-quick-and-sobering-guide-to-cloning) Use with caution, but this can be great for explainer videos and introductions.

The first commercially available text-to-video tool was also recently released, Runway v2. It creates short 4-second clips, and is more of a demonstration of what is to come, but is worth taking a look at if you want a sense of the future development in this space.

_Some things to worry about: _Deep fakes are a huge concern, and these systems need to be used ethically.

# Work with documents and data

**For data (And also any weird ideas you have with code):** Code Interpreter

**For documents: **Claude 2 for large documents or many documents at once, Bing Sidebar for smaller documents and webpages (the sidebar, part of the Edge browsers can “see” what is in your browser, letting Bing work with that information, though the size of the context window is limited)

[I wrote about Code Interpreter last week](https://www.oneusefulthing.org/p/what-ai-can-do-with-a-toolbox-getting). It is a mode of GPT-4 that lets you upload files to the AI, allows the AI to write and run code, and lets you download the results provided by the AI. It can be used to execute programs, run data analysis (though you will need to know enough about statistics and data to check its work), and create all sorts of files, [web pages](https://twitter.com/prkeshari/status/1678155933606637568?s=20), and even [games](https://twitter.com/icreatelife/status/1678184683702566922?s=20). Though there has been a lot of debate since its release about the risks associated with untrained people using it for analysis, many experts testing Code Interpreter are pretty impressed, [to the degree that one paper suggests it will require changing the way we train data scientists.](https://twitter.com/emollick/status/1678615507128164354?s=20) Go to my previous post if you want more details on how to use it. I also made an initial prompt to set up Code Interpreter to create useful data visualizations. It gives it some basic principles of good chart design & also reminds it that it can output many kinds of files. You can find that [here](https://t.co/m4yAdKROiJ).

For working with text, and especially PDFs, Claude 2 is excellent so far. I have pasted in entire books into the previous version of Claude, with impressive results, and the new model is much stronger. You can see my previous experience, and some prompts that might be interesting to use, [here](https://www.oneusefulthing.org/p/what-happens-when-ai-reads-a-book). I also gave it numerous complex academic articles and asked it to summarize the results, and it does a good job! Even better, you can then interrogate the material by asking follow-up questions: what is the evidence for that approach? What do the authors conclude? And so on…

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

_Some things to worry about: _These systems still hallucinate, though in more limited ways. You need to check over their results if you want to ensure accuracy.

# Get information and learn stuff

**Best free option:** [Bing](https://www.bing.com/search?q=Bing%20AI&showconv=1&FORM=hpcodx)

**Paid option: **Usually Bing is best. For children, [Khanmigo ](https://www.khanacademy.org/khan-labs)from Khan Academy offers good AI-driven tutoring powered by GPT-4.

If you are going to use AI as a search engine, probably don’t do that. The risk of hallucination is high and most AIs are not connected to the Internet, anyway (which is why I suggest you use Bing. Bard, Google’s AI, hallucinates much more). However, there is some evidence that AI can often provide more useful answers than search when used carefully, [according to a recent pilot study](https://arxiv.org/abs/2307.01135). Especially in cases where search engines aren’t very good, [like tech support, deciding where to eat, or getting advice](https://twitter.com/emollick/status/1643718474668097538?s=20), Bing is often better than Google as a starting point. This is an area that is evolving rapidly, but you should be careful about these uses for now. [You don’t want to get in trouble.](https://www.nytimes.com/2023/06/08/nyregion/lawyer-chatgpt-sanctions.html)

But more exciting is the possibility of using AIs to help education, including helping us learn. [I have written about how AI can be used for teaching](https://www.oneusefulthing.org/p/assigning-ai-seven-ways-of-using) and to [help make teachers’ lives easier and their lessons more effective](https://oneusefulthing.substack.com/p/using-ai-to-make-teaching-easier), but it can also work for self-guided learning as well. You can ask the AI to explain concepts and get ver good results. This [prompt is a good automated tutor](https://twitter.com/emollick/status/1669434927761313807?s=20), and use can find a [direct link to activate the tutor in ChatGPT here](https://chat.openai.com/share/ec1018ec-1d86-4160-b587-354253c7d5cb). Because we know the AI could be hallucinating, you would be wise to (carefully!) double-check any critical data against another source.

# And more?

Thanks to rapid advances in technology, these are likely the worst AI tools you will ever use, as the past few months of development have shown. I have no doubt I will need to make a new guide soon. But remember two key points that remain true about AI:

- AI is a tool. It is not always the right tool. Consider carefully whether, given its weaknesses, it is right for the purpose to which you are planning to apply it.
- There are many ethical concerns you need to be aware of. AI can be used to infringe on copyright, or to cheat, or to steal the work of others, or to manipulate. And how a particular AI model is built and who benefits from its use are often complex issues, and not particularly clear at this stage. Ultimately, you are responsible for using these tools in an ethical manner.

We are in the early days of a very rapidly advancing revolution. Are there other uses you want to share? Let me know in the comments.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This post is licensed under a [Creative Commons Attribution-NonCommercial 4.0 International License](http://creativecommons.org/licenses/by-nc/4.0/).
