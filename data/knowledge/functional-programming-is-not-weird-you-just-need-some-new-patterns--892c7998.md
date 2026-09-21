---
title: "Functional Programming is not weird: you just need some new patterns"
notion_id: 892c7998-9b8d-486d-86cf-95e32455106d
notion_url: https://app.notion.com/p/Functional-Programming-is-not-weird-you-just-need-some-new-patterns-892c79989b8d486d86cf95e32455106d
last_edited: 2023-04-25T13:24:00.000Z
source_url: https://medium.com/@cameronp/functional-programming-is-not-weird-you-just-need-some-new-patterns-7a9bf9dc2f77
tags: ["Functional programming", "Programming", "Learning", "Article", "Medium", "English"]
---
[https://medium.com/@cameronp/functional-programming-is-not-weird-you-just-need-some-new-patterns-7a9bf9dc2f77](https://medium.com/@cameronp/functional-programming-is-not-weird-you-just-need-some-new-patterns-7a9bf9dc2f77)

## How I learned to stop worrying and love immutability

9 min readFeb 29, 2016

--

[UPDATE: I gave a talk on this subject at the Empire Elixir Conference in May. You can watch it here.]

I saw something on Hacker News this weekend that I would have agreed with a year ago, but not anymore: Functional Programming Is Not Popular Because It Is Weird.

> Writing functional code is often backwards and can feel more like solving puzzles than like explaining a process to the computer. In functional languages I often know what I want to say, but it feels like I have to solve a puzzle in order to express it to the language.

This line really resonated with me. When learning Elixir, and before that Scala, I often felt that I’d spend an hour ruminating on a problem, turning my brain inside out to try and imagine how I was going to model it without manipulating state. Then I’d eventually hit on the solution, which would turn out to be about three lines of lovely, elegant code, that worked on the first try. It felt more like playing some kind of game than the kind of thing I could imagine using for Serious Work™.

This was me a year ago

But I enjoyed the game, and when I came to understand the benefits of the Erlang backend, in terms of concurrency, I decided to stick with it. And something unexpected happened. I’d been doing little bite-size problems for some time. Project Euler, the problems in Dave Thomas’ Programming Elixir, etc. At some point, I noticed that not only was I no longer “solving puzzles” to achieve simple effects, I was actually writing code to solve problems at a much faster pace, and with much greater precision than ever before. Complex algorithms felt like they were flowing out of me, and they kept amazing me by working correctly right out of the gate. It was as if the all the promises of FP advocates came true at once. Weird.

I wondered why this was. Had I suddenly developed the ability to think in the upside down way that FP was asking of me? There was no question I was faster, and at the risk of sounding like a zealot, the feeling was exhilarating. So I’ve been thinking about what it was that happened, and how I can both keep it up, and help others to do the same.

It’s about patterns. Not the big architecture GoF-style patterns people argue about incessantly, but the modest everyday patterns that we use to write code. They’re so small that they don’t even have names half the time, but whether you’re an FP programmer, or an Imperative programmer, you use them constantly. Some are specific to a particular language, and others are common patterns within language families. If you’ve been writing imperative code for awhile, you have a big toolbox of these that you can instinctively reach for at any moment, and you don’t even have to think about it. Here’s an example:

```plain text
void print_arr(int *list, int count) { for(int i = 0; i <= count; i++) { printf(“%d “, list[i] ); } printf(“\n”);}
```

Do you see the bug? If you’ve worked with C, or Java, or any language like it, you know right away what this code is supposed to do, and, you know right away that it should be i < count, not i <= count. You hammer these kinds of loops out without thinking, but I’ll bet that years ago, when you were first learning, you’d have had to think about this quite a lot. But this is now a pattern you can deploy any time you feel the need to “walk an array”, or “walk a string”, and you can recognize it when you see it without thinking much about how it works.

Back in the 90’s, when I spent a lot of time interviewing C programmers, a very common interview question was “write a function to remove the duplicates from a sorted array of ints”. I loved this question, and I used to ask it over and over again. I’ve probably asked 200 people to solve this problem, because at the time, it was a great way to see how comfortable the candidate was with C. (I know, I know, whiteboard interviews aren’t a good test, we don’t write code on whiteboards, etc. Give me a break, it was the 90’s. We did all sorts of things that would horrify you today.)

The reason I liked sorted-array-with-dupes, as we called it, was that 99% of candidates could describe a basic strategy for how they were going to solve it:

1. Walk the array…
2. keeping track of the most recent value we’ve taken…
3. and if it changes, take the new value, and set that to be our most recent value
4. until we get to the end, and then return the new array.

Everyone made it this far, but only about 10–15% of the candidates could easily translate that into code. This is because it involves two iterators, and keeping two iterators under control is a pattern that requires experience. It was also the kind of pattern that came up again and again when programming in C.

Here’s what this code might look like in C:

```plain text
int rem_dupes(int *array, int count) { if (count == 0) return 0; int current_val = array[0]; int current_output_position = 1; for (int i = 1; i < count; i++) { if (array[i] != current_val) { array[current_output_position++] = array[i]; current_val = array[i]; } } return(current_output_position);}
```

After dealing with the edge case of count==0, it walks the array with the main for-loop iterator, and keeps the second iterator (current_output_position) up to date manually. Then it returns the new count. If you’ve written a bunch of C, or a bunch of Java, you can probably grok this without difficulty.

But I think you’ll agree with me, that it’s not exactly the simplest thing in the world, and that once upon a time you might have struggled to write this without bugs. Also, even now, how confident are you that there aren’t any bugs in that code? I’m not 100% confident, and I’ve compiled and run it just this morning.

Now here is the same function, written recursively in Elixir, using multiple function heads and pattern matching:

```plain text
def rem_dupes([]), do: []def rem_dupes([first | t]), do: [first | rem_dupes(t, first)]def rem_dupes([], _), do: []def rem_dupes([same | t], same), do: rem_dupes(t, same)def rem_dupes([next | t], _last), do: [next | rem_dupes(t, next)]
```

[UPDATE: I’ve made this more idiomatic thanks to advice from pseudothere on /r/elixir]

If you don’t know Elixir, or if you struggle with FP patterns, then this probably looks like exactly the sort of magical incantation that annoys you. You’re thinking “Look at that: it’d take me an hour to solve the puzzle, and when I finished it, I’d have code that would never make sense to me again.” And you’d be wrong. To someone who is experienced with Elixir (or another functional language), this code is quite readable because it uses a few really common patterns that any Elixir dev will have in their back pocket after awhile. Allow me to illustrate:

```plain text
def rem_dupes([]), do: []
```

This is an extremely common maneuver. Handle the edge case in a separate function head, and handle it first. No need to worry about it again. Is it more confusing than the C approach? Probably not.

```plain text
def rem_dupes([first | t]), do: [first | rem_dupes(t, first)]
```

Ok, now that’s a handful, I’ll admit. The notation “[ first | t ]” causes first to be bound to the first element, and t to be bound to the rest of the list. All we’re basically doing here is what we did implicitly in the C code with our initialization of current_val and current_output_position. We’re capturing the first item, and setting it as the “latest” item in our call to the two-parameter version of rem_dupes. We’re setting up to “walk the list”. Using a 1-arity function head to set things up for the real work to get done by a 2-arity or 3-arity function is also an extremely common pattern in Elixir. Maybe it looks weird to you now, but I assure you, it won’t look weird after you’ve done it 100 times. So lets look at the 2-arity function, where the real work happens:

```plain text
def rem_dupes([], _), do: []
```

We’ve gotten to the end of the list, there’s nothing more to return, and we don’t care what the last value was. This is the kind of thing that will drive you nuts if you’re an imperative programmer, because it feels like we’re doing the last part first. But I ask you, is this any different to checking “i < count” at the top of the for-loop instead of the bottom? It’s just a pattern. At one time I had to think about it, now it’s obvious what it’s doing.

```plain text
def rem_dupes([same | t], same), do: rem_dupes(t, same)
```

What dark magic is this!? Here we’re using pattern matching to skip over elements that are equal to the value we’re looking for (the second parameter). This will only match, and therefore this function will only execute, if the head of the list is equal to the second parameter. At first this kind of programming will seem batshit crazy, but eventually you’ll internalize the pattern, and you’ll read and write it without much thought.

```plain text
def rem_dupes([next | t], _last), do: [next | rem_dupes(t, next)]
```

And finally our step to grab the next value. Notice we apply “next” to the second parameter of the recursive call, and we no longer care about the value of last (that’s what the _ means).

Each of these lines represents very typical patterns in Elixir, and you will not master these patterns overnight. You didn’t master the patterns you know overnight either. I would say that it took me at least 80–100 hours of programming Elixir, before it suddenly became a hundred times easier. If someone came running into the office today, panting and perspiring heavily, and told me that someone needed to write some code to parse a bunch of text files, execute a bunch of arcane business rules, and update a database or two, or the world would end, I would absolutely be reaching for Elixir first. There’s no doubt in my mind that I’d solve whatever screwed up mess they’d brought me a hell of a lot faster with Elixir than with any other tool.

If you trust me, and you’d like to try and internalize a bunch of these patterns yourself, here’s my advice:

1. Do not put any pressure on yourself. Assume that you’re going to struggle to get your head around these patterns at first, and don’t worry about how long it’s going to take. Trust your old pal Cameron, who promises you’ll get it eventually, and it’ll be totally worth it.
2. 100 little challenges are going to be much more helpful than one big stonking challenge. Do not try and write an online version of Monopoly or something right out of the gate. Solve the kinds of problems you’d expect to see in a first year programming class. I list several great sources of these kinds of problems below.
3. Solve the same kinds of problems over and over. These micropatterns need to become almost like muscle-memory for you, and the only way to achieve that is to do them many, many times. To give you an idea what I mean, I wrote that Elixir solution to the problem above in less than a minute this morning, and all in one go, because it only uses micropatterns I’ve used many times before.
4. Solve the exact same problems again. Go back to problems you’ve solved before, and solve them again, with your newfound understanding.
5. Don’t feel bad about putting a problem aside. If you’re not enjoying a problem, just skip it. Noone’s paying you to do this, and there are a hell of a lot of problems out there. Drop it and do something fun. Maybe you’ll come back to it later, maybe not.
6. Ask for help! The Elixir Slack team, and the Elixir mailing list are the friendliest places on the whole internet. There’s a #beginners channel in Slack also, and there’s helpful folks in there all the time. There are Elixir meetups in most major cities. If you’re in NYC, the Empire Elixir Conference is coming in May. I’ll be there, and so will a lot of other passionate Elixir and Erlang devs. Early Bird tickets are going on sale in a couple of days.

### Great sources of problems to solve:

- Project Euler: These problems are obviously a little math heavy, but I learned a lot by doing the first 50 or so.
- Exercism: There’s an Elixir track here, and a common thing to do is submit your solution to exercism, and then ping the #general channel in slack to solicit feedback. I’d encourage that behavior.
- Advent of Code: This was tremendously helpful for me, and brilliant fun. Especially the first 14–15 days. Highly recommended.
- Programming Elixir by Dave Thomas. This is a great book, as you might expect from Dave, and offers many little exercises throughout, along with an online forum where solutions are discussed.
- Exercises for Programmers by Brian Hogan. This book includes 57 bite-sized exercises. I haven’t done them all, but I’ve done a few and it’s been fun so far.

### Thank you for your time. I hope you enjoyed this, or found it useful, or both. If so, I’d really appreciate it if you clicked the heart button below so more people will see it. Thanks!
