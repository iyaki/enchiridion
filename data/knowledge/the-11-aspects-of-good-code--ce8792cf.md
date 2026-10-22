---
title: "The 11 Aspects of Good Code"
notion_id: ce8792cf-d301-47f3-99a7-8559f90a9547
notion_url: https://app.notion.com/p/The-11-Aspects-of-Good-Code-ce8792cfd30147f399a78559f90a9547
last_edited: 2023-07-28T15:08:00.000Z
source_url: https://www.pathsensitive.com/2023/07/the-11-aspects-of-good-code.html
tags: ["Article", "Path-Sensitive", "English", "Programming"]
---
Lessons on code quality start in the first few weeks of learning to program, when a newcomer to the field is taught the basics of variable naming and told why programming languages have comments. They continue in countless blog posts and in every debate on a pull request.

Avoid it or embrace it, code quality training permeates one's entire career.

But it is so easy to lose sight of why.

Says the skeptic: pretty code is a distraction, like the gargoyles on a cathedral or the curlicues of Baroque architecture.

Says the maximalist: Have you ever had a debugging session that demanded you turn over so many stones it felt as if the world was crumbling under you? You don't achieve 10x velocity by becoming 10x faster at debugging, but by writing code that doesn't need debugging at all.

And says the guru: that which is merely a pretty distraction is not code quality.

I've made a professional quest to clarify all the fuzzy terms of software engineering. Students of our [course](https://mirdin.com/the-advanced-software-design-course/) learn definitions of code knowledge and coupling in the first lessons. One of my research papers gives [a rigorous definition](https://www.jameskoppel.com/files/papers/demystifying_dependence.pdf) of dependence.

Today, I carry this quest to its apotheosis: what is quality code?

To this question, there can be no short exhaustive answer. Asking “what is good code” is a lot like asking “how do chemicals work?” It's the subject of an entire field.

But we can more easily ask what is the purpose of chasing code quality, even if achieving it is a craft worthy of a lifetime of study. To recognize quality code, we begin by asking: what are the external and internal properties quality code should have?

## **External Properties**

### **Good code is done code**

We begin with the cliché. All discussions of quality are grounded in the ultimate purpose of the object being designed. The purpose of the vast majority of code is to be executed as software which accomplishes some goal, be it entertaining people, helping them with their taxes, shuttling data, or testing other code. There is also a minority of code built for other purposes: experiments to see if something is possible, examples to explain a library or algorithm, and code that does [tricks](https://en.wikipedia.org/wiki/Quine_(computing)) such as printing its own source.

For all of these, code that fails to achieve its purpose cannot have extrinsic quality any more than an abandoned construction site can be a useful building.

But this does not justify single-mindedness in getting a program to work. There are also non-functional requirements such as performance. Software cannot be quality if sluggishness sends users back to pen-and-paper.

And though an entire business may be dedicated to helping a software product fulfill its purpose, that does not subordinate all other functions to “code quality.” Whether code is quality cannot depend on factors that lie entirely outside the realm of engineering. The failure to market a software package does not make it bad, and a top-down directive requiring people to use it does not make it good.

And so good code is done code.

But we cannot stop there. That is not the end of the story. It is just the beginning.

For if you say “We got it done and delivering value to the customer,” that is not an excuse to your boss when you explain why adding a feature to [wish users “Happy birthday”](https://www.youtube.com/watch?v=y8OnoxKotPQ) will take several years. And it is not an excuse to yourself when you spend [4 days debugging](https://archive.ph/ZfRDI) an issue that turned out to be a typo. Done code is not good code.

### **Good code is understandable**

By one definition, an engineer is someone who understands a system at a deep level.

And, it follows that, for code to have good engineering, it must be understandable.

And sometimes, such as for teaching code, this is its entire purpose.

So you want code to be understandable. But understandable to whom?

To yourself and the people who need to read it.

Or more specifically: to those people _at the time_ they need to read it.

The foolish engineer is offered a new skill, one that will shrink a segment of his code by a factor of 10. “No-one else will understand this” he says as he refuses to learn it. He thus reveals a low expectation of himself masquerading as a low expectation of others. He has hidden within his comfort zone of skill.

The arrogant engineer has a skill and knows it will be effective. “Anyone who cares about this code should be able to learn the technique I used.” If there are to be suitably ambitious readers, the choice turns out correct; but if the audience is one whose concerns lie elsewhere, then it does not. But an uninformed decision cannot be a good one. He has hidden within his comfort zone of empathy.

Either can improve by breaking out of their comfort zone, learning which walls to climb and how, and aiding the rest through construction and placement of ladders.

But the arch-engineer of engineers breaks the comfort zone itself. They are not concerned with climbing nor ladders, for those who follow shall suddenly find themselves atop mountains.

### **Good code is evolvable**

Software is not a point in time but a system.

Precious few programs stand like museum pieces encased in glass, existing only for their own sake, or illustrating a piece of the frozen past. The rest are connected to other programs, to platforms, to growing businesses and rotating customers. They are connected to a changing world.

And now, as we spend our lives glued to screens that came from robotic factories and arrived via satellite-controlled ships, software _is_ the changing world.

You cannot change the world without changing its software. Every software engineer carries the professional burden of building software that is easy to change.

And specifically, to change from one desirable state to another.

We must avoid creating rigid code that is difficult to change at all.

And we must also prevent brittle code that can all-too-easily be changed to something broken.

Good code is easy to extend and difficult to break. The power of a design lies not in what it can do, but rather what it can't do.

But why prepare for a future that may never come? It is impossible to predict the exact ways code will change.

Yet it is often easy to predict _that_ code will change. Or even _where_. And that's all that's needed to create evolvable code.

Yes, it is a folly to design assuming certain changes will need to be made as life takes a certain path. But it is a greater folly to design as if no changes will occur at all.

## **Internal Properties**

It is a lofty goal to say that a program must be correct, understandable, and evolvable.

It is an achievable goal to say that a single function should pass its tests, have few branches, and use abstract types.

And yet the summation of the latter yields the former. Extrinsic quality comes from intrinsic quality. These properties are presented below:

### **Good code can be understood modularly**

Programs are composed of files. Files are composed of declarations. Declarations are composed of lines.

That is to say, programs are built out of pieces.

And every single piece has its purpose.

Every time a line is executed, a change occurs, and there are many true statements that can be said about each change. Most such statements are of no consequence, while some are crucial to the ensuing lines achieving their purpose.

But in some programs, there is a third category of statements. There are facts about the program state that become true on some line, and then are of no consequence until some distant line requires them to be true. Whereas most lines are of concern only to their neighbors, these two lines have grasped hands through a wormhole, their fates entangled. A change in one place can cause breakage on the other side of the universe.

In good code, the purpose of every line can be stated simply. Each line can be understood in isolation. For each, one can reason: if some simple fact about the state of the program is true, then, after running this line, some other simple fact will be true. The assumptions and guarantees of each line click together like Legos, forming simple and correct functions, which in turn click together into simple modules and simple programs.

In bad code, you read a function, ask whether it works, and then read a dozen more in order to have an answer. Changes must be made as tenderly as one playing Jenga, lest the tower collapse.

Good code works by design. Bad code by coincidence.

### **Good code makes it easy to recover the intent of the programmer**

A programmer dreams a new entity. Her mind gradually turns dream into mechanism, mechanism into code, and the dreamed entity is given life.

A new programmer walks in and sees only code. But in his mind, as he reads and understands, the patterns emerge. In his mind, code shapes itself into mechanism, and mechanism shapes itself into dream. Only then can he work. For in truth, a modification to the code is a modification to the dream.

Much of a programmer's work is in recovering information that was already present in the mind of the creator. It is thus the creator's job to make this as simple as possible.

But to do so is a constant struggle.

Every naming decision is a quest to find the word that conjures in the reader's mind the true purpose of the named while warding off misconceptions.

Every function, a quest to carve behavior into something meaningful.

Every module, a quest to create new words that give new powers to the wielder.

Through each such step, we climb towards the ideal of making the program written not in the language of the machine, but in the language of the dream.

They say that for those who have reached the peak, they can simply dream changes to the program and it is instantly so. But great powers are had even by those who only make it partway.

### **Good code expresses intent in a single place**

But it is not enough for it to be easy to go from code to design. It must also be easy to change the design to new code.

The shaper of atoms walks into a room under construction, wide open and brightly lit. “No!” he cries. “I want it to be dark and intimate.” Before him a vast itinerary of work is created, as that one directive demands thousands of strokes of the paintbrush and new choices for every object so contained.

The shaper of bits walks into a website, and says she wants a dark mode. In good code, she speaks the new colors that comprise a dark mode, and it is so. In great code, she merely speaks “dark mode” and the colors are inferred.

Yet all too often such a change, though simple, requires tweaks in thousands of locations, like a thousand well-coordinated strokes of the brush.

The bits should be easier to change than the atoms, for they live inside the machine.

Yet they can be harder, for there are so many more of them.

### **Good code is robust**

If every line serves a purpose, then every line must be correct.

That means that every line is a new opportunity for a mistake to slip in unnoticed.

And how easy it is to make a mistake is something under the control of the software designer.

Some codebases are so treacherous that working in them is like a tightrope walk across the Grand Canyon. There are functions which require consulting a tome to invoke correctly. Writing to a data structure can produce nonsense. Reading from a data structure may produce only a partial story.

Other codebases are more like an elevator ride, to the point where not even deliberate effort can produce an accident. In such code, APIs have guardrails, where any misuse is either disallowed or can only be accomplished by spray-painting on a red flag. Try as you might, no write to a data structure can produce nonsense. If it compiles, it probably works.

As Tony Hoare says, one can write code so simple there are obviously no bugs, or so complex that there are no obvious bugs.

If you must think as hard as you can to check that a program works, it probably doesn't.

But in good code, you barely need think at all.

### **Good code hides secrets**

Software is not a point in time but a system.

And it is not one system, but many interacting ones.

And each is constantly morphing.

But if it looks and acts the same on the outside, no-one will ever know.

It does not matter to the driver when a car's engineer changes its wiring. Unless, that is, the manual had told her in great detail what to expect from its electrical system and she had come to depend on it. The one who learns the car's battery can charge 5 cell phones for exactly 433 minutes before dying is the one able to achieve maximum performance. But, in a changing world, the one who uses this forbidden knowledge tiptoes close to ruin.

Subsystems are joined when their creator's minds are joined, in conversations that should not occur sharing details that should not be shared. Or when the single master fails to erect a firewall in his own mind.

The hotshot boasts about knowing everything. She creates software that can only be worked on by her fellow all-knowing.

The master's virtue is knowing nothing. And that's enough to maintain his software.

### **Good code isolates assumptions**

Minimizing use of knowledge is the path to evolvability. Secrets are but the extreme, known only to their owners. Every use of knowledge ties the program to the World That Was, hindering the creation of the World That Could Be.

For every datum, there are the components that create it, the components that use it, and those in between that merely deliver it. Do those components pass the datum along like a sealed package? Or are those couriers prying into its contents?

A value is passed from one end of the program to another. Every function on the way that calls the value an “int” is another barrier to making it a float. And every function on the way that calls it anything is another barrier to turning this value into two numbers.

The physical world is full of irreversible changes. Build a building and the town shapes around it; burn it back down and forever shall the wind be tainted with its ashes. But when it comes to reshaping the world of bits, the only thing in a programmer's way is himself.

### **Good code is open**

Programs deal with a domain, and both program and domain can be sliced countlessly many ways into sets. Sets of options! Sets of fields! Sets of formats! Sets of formats of fields which represent options!

And as programs and domains change, such sets grow and shrink. When good code deals with such a set, it is to the extent possible agnostic to the set's size.

The simpleton sees an entity with two possible values, and builds the program using a boolean. The next day, the possibilities have grown to three. A rewrite is required.

There is a set of three kinds of entities, and a program is written that can work with each of them. Then comes the day where one is deprecated. If the program was open in this set, then the relevant code is already in a box that can be discarded. If the program was closed, then branches all throughout have become skeletons demanding burial.

The open-minded person accepts new things. So does the open program.

### **Good code uses a programmer's full wisdom**

The journeyman programmer finds a list of 10 principles for good code. She studies them one by one, and after years of toil attains mastery. Before her the baffling complexity of programs stood as stalagmites of wax; before her gaze, it has now melted down and separated into buckets.

The path there is one of toil. Every place where intuition says the code could be simpler, she seeks how. Every issue that was hard to debug, she searches for how it could have been prevented.

And then she declares “that is all.” Her apparent mastery has brought her respect, and her skill cleaves problems that foil others. She accepts her place at the top and rests.

But the one with the potential to become a grandmaster does not rest. They notice the dregs of wax that fall outside the buckets and see in them opportunities to find new explanations. As they search ever deeper, the buckets dissolve and reveal the interconnected whole. As with the programmer learning a codebase, they have stepped into the dream behind the concepts, and are now ready to dream themselves.

This list came from years of observation, reflection, study, teaching, and refinement. It is yours now to study, criticize, preach, ridicule, and extend.

## Resources

**On modularity:** [The 3 Levels of Logic](https://www.pathsensitive.com/2018/01/the-three-levels-of-software-why-code.html)

**On intent:** [My Favorite Principle for Code Quality](https://www.pathsensitive.com/2018/02/making-bugs-impossible-illustrating.html)

**On Robustness:** [State of emergency! The four ways your state might be wrong](https://medium.com/@note89/state-of-emergency-the-four-ways-your-state-might-be-wrong-b6cebea38d78)

**On Secrets:** David Parnas, “The Secret History of Information Hiding”

I mostly lack public resources on openness and the sequestering of assumptions, although [refunctionalization](https://www.pathsensitive.com/2019/07/the-best-refactoring-youve-never-heard.html) is one technique for achieving it.

But for all of these, the best way to learn them is through [deliberate practice.](https://www.pathsensitive.com/2018/02/the-practice-is-not-performance-why.html)

And for that, we have the [Advanced Software Design Course.](https://self-service.mirdin.com/)

_Thank you to Nils Eriksson, Jun Hong “Nemo” Ya”, Emmanuel Genard, Paul Weidinger, and Yongming Han for comments on earlier drafts of this essay._
