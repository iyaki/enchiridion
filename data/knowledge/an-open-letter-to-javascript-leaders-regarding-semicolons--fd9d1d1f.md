---
title: "An Open Letter to JavaScript Leaders Regarding Semicolons"
notion_id: fd9d1d1f-6321-4c5e-b0b0-52dbbb54fb5f
notion_url: https://app.notion.com/p/An-Open-Letter-to-JavaScript-Leaders-Regarding-Semicolons-fd9d1d1f63214c5eb0b052dbbb54fb5f
last_edited: 2024-04-17T19:49:00.000Z
source_url: https://blog.izs.me/2010/12/an-open-letter-to-javascript-leaders-regarding/
tags: ["English", "Javascript", "Article", "blog.izs.me (Isaac Z. Schlueter)"]
---
I got this email last night from Sean Silva:

> I was browsing the code for your npm.js project (this file in particular: [https://github.com/isaacs/npm/blob/master/lib/npm.js](https://github.com/isaacs/npm/blob/master/lib/npm.js)), and noticed you using a style where you line up your commas under the ‘r’ of your var statements, and under your [ and { in array/object literals. I’m very fond of this kind of formatting, but have been reluctant to use it since most js resources expound a fear of js’s automatic semicolon insertion wreaking havoc on your code if you don’t end lines with something which implies a continuation.

Is this safe to place commas like this in browser code, or is it just node where this is possible?

I wrote a few paragraphs, and then decided to shorten it to just this response:

> Yes, it’s quite safe, and perfectly valid JS that every browser understands. Closure compiler, yuicompressor, packer, and jsmin all can properly minify it. There is no performance impact anywhere.

I am sorry that, instead of educating you, the leaders in this language community have given you lies and fear. That was shameful. I recommend learning how statements in JS are actually terminated (and in which cases they are _not_ terminated), so that you can write code that you find beautiful.

Inimino posted [this very clear explanation](http://inimino.org/~inimino/blog/javascript_semicolons). In his typical style, it is crisp, clear, authoritative, well researched, and he generally kept his opinions to himself.

I am going to be a bit more opinionated.

### The Rules

In general, `\n` ends a statement unless:

1. The statement has an unclosed paren, array literal, or object literal or ends in some other way that is not a valid way to end a statement. (For instance, ending with `.` or `,`.)
2. The line is `-` or `++` (in which case it will decrement/increment the next token.)
3. It is a `for()`, `while()`, `do`, `if()`, or `else`, and there is no `{`
4. The next line starts with `[`, `(`, `+`, , `/`, , `,`, `.`, or some other binary operator that can only be found between two tokens in a single expression.

The first is pretty obvious. Even JSLint is ok with `\n` chars in JSON and parenthesized constructs, and with `var` statements that span multiple lines ending in `,`.

The second is super weird. I’ve never seen a case (outside of these sorts of conversations) where you’d want to do write `i\n++\nj`, but, point of fact, that’s parsed as `i; ++j`, not `i++; j`.

The third is well understood, if generally despised. `if (x)\ny()` is equivalent to `if (x) { y() }`. The construct doesn’t end until it reaches either a block, or a statement.

`;` is a valid JavaScript statement, so `if(x);` is equivalent to `if(x){}` or, “If x, do nothing.” This is more commonly applied to loops where the loop check also is the update function. Unusual, but not unheard of.

The fourth is generally the fud-inducing “oh noes, you need semicolons!” case. But, as it turns out, it’s quite easy to _prefix_ those lines with semicolons if you don’t mean them to be continuations of the previous line. For example, instead of this:

```javascript
foo();
[1,2,3].forEach(bar);
```

you could do this:

```javascript
foo()
;[1,2,3].forEach(bar)
```

The advantage is that the prefixes are easier to notice, once you are accustomed to never seeing lines starting with `(` or `[` without semis.

### Restricted Productions

The other common argument for semicolons has to do with semicolon insertion and “restricted productions”. That is, if you have a `\n` immediately after a `return`, `throw`, `break`, or `continue` token, or a `++` or `--` as a postfix operator (that is, `x++\n` or `y--\n`), then it will terminate the statement, no exceptions.

```javascript
//ok
return 7

//probably a mistake
return
       7
```

However, again, this is in fact _easier_ to spot and avoid once you get out of the habit of terminating every statement with a semicolon. When I see the second, my brain instinctively associates the `\n` with “ok, this is over now”, because return is _always_ terminated by just a linebreak.

Lining up the most relevant tokens on the left edge of the screen makes them demonstrably easier for humans to quickly scan. Piles of research on the subject of speed reading and eye-tracking suggest that a missing token on the right is far more likely to be overlooked than one on the left. So, I say, make the right-edge irrelevant, and put the important things on the left.

### So which style is better?

To the extent that there is an objectively “better” choice, it appears to me that the minimal-semicolon/comma-first style is **slightly** superior, both because it is fundamentally more scannable and because it encourages programmers to better understand the language they use.

I can pretty well guarantee that, if you care about this even a little, I care less about your JavaScript style than you do about mine. This isn’t an article where I try to convince you to write your code like I write mine. [We should all decide the pants policy in our own homes.](http://groups.google.com/group/nodejs/msg/428220ab8cd199d2)

Just as a show of good faith…

### Good Reasons to Put Semicolons Everywhere

The best reasons for excessive semicolon usage are esthetics and politics.

“I put semicolons in my JavaScript because without semicolons, it’s not valid C/C++/Java/Whatever.” If you have to write a bunch of Java or C code in a project, and want your JavaScript to not look too different, then that is a valid concern. ([Cassis](http://cassisjs.org/) takes this approach to its absurd end.)

“We do it this way because we use this linter, and it says to.” Consistency is important, and linters are one way to help a group of people stay consistent. Writing an npm-style linter is on my todo list, but it’s not very high up on it.

### The Most Terrible Reason to Put Semicolons Everywhere

“They’re required because ASI is unreliable.” Seriously!?

These rules date back to the early days of JavaScript, in the late 90s. They’re not new, and in my opinion there is no excuse for someone calling themselves a professional JavaScripter and not understanding statement termination. **It is blatantly irresponsible of the thought leaders in the JavaScript community to continue to spread uncertainty rather than understanding.**

Furthermore, the typical place where “automatic semicolon insertion” bites unexpectedly is with the restricted productions. **Adding semicolons to every line will not make ****`return\nfoo`**** return anything other than ****`undefined`****.** The problem is that you **do** use line breaks, not that you **don’t** use semicolons.

The only way to prevent restricted productions from _ever_ being an issue is to always use semicolons _and never use linebreaks_. No one is suggesting that. So stop talking about restricted productions as if they matter, or offering semicolon overuse as an alternative to understanding ASI. You have to understand ASI to be a competent JavaScripter, period.

Which leads me to…

### The Part Where I Get All Opinionated and Piss You Off (despite a noble effort to the contrary)

If you don’t understand how statements in JavaScript are terminated, then you just don’t know JavaScript very well, and shouldn’t write JavaScript programs professionally without supervision, and you _definitely_ should not tell anyone else how to write their JavaScript programs.

I’m guessing I just insulted you. That’s unfortunate. I know that you probably know all sorts of things _around_ JavaScript, like the DOM, and CSS, and MSIE’s little quirks, and jQuery. You have maybe also spent some time learning about closures and prototypes and scope and activation objects, and even hacked a few extensions onto V8 or SpiderMonkey. You’re not a dummy, I’m sure. In fact, you’re almost certainly smarter than I am, and probably better looking and nicer, too. I’m sure we have a lot in common, and could maybe even be friends.

But if you don’t understand _what a JavaScript statement __**is**_, then there is a huge hole in your understanding of perhaps the most fundamental aspect of the language.

And that’s ok. I don’t speak Spanish very well and my C is pretty novice; I also don’t call myself an expert in either one, though I know enough to get by in many situations. If I were to get a job that involved a lot of Spanish speaking or C coding, I’d want someone watching to help me avoid making any serious mistakes.

Like most things in JavaScript, the statement termination rules are not very _well_ designed, but they’re also not particularly _hard_ to understand and use. That understanding just takes a bit of time and effort.

Cozy up with some hot chocolate and the ECMAScript spec some Saturday afternoon. Practice a little. Play around with some test programs. It’s a good time.

Or don’t do that, if you don’t feel like it. It’s your life. You almost certainly have better things to do with it.

Just please stop making authoritative claims like “terminate all lines with semicolons to be safe.” It’s not any safer, or more reliable.

### Addenda 1: “leaders”

> So, mr ruffler of feathers, who are these “leaders” you speak of? Why didn’t you name names?

Because there are too many to name, and I don’t know all of them.

If you have been writing JavaScript for a while, and you provide guidance or leadership to another person who has been writing JavaScript for less time than you, then I’m talking to you. Being a leader is a responsibility. Take it seriously. Don’t spread lies. Be an expert, or admit you’re not an expert. But don’t drive that car without a license.

### Addenda 2: Literary Programmer

> but in English, we put punctuation at the end, not the beginning

JavaScript isn’t english. We also don’t denote ownership (or subject-verb connection) using a period in English. We don’t have Object Literals in English, and we only indent the first line of paragraphs, not all the middle sentences.

This is such a silly argument, I have no choice but to fall in love with it. I started out your detractor, but you won my heart, Literary Programmer. From now on, I’m going to only put line-breaks at the end of functions, and never in the middle, and indent the first line of each.

### Addenda 3: Pedantry

> Why you wanna change my codes? What are you some kinda pedant?

Code however you want. I don’t care, even a little.

Please just don’t lie to people. That’s all I’m asking. It’s such a little bit of politeness. It’s not hard. Just say true things, instead of being a liar, that’s all that this is about.

2010-12-18
