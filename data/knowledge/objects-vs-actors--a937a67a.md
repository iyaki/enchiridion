---
title: "Objects vs. Actors"
notion_id: a937a67a-ea55-4561-9999-5ffbabf7637d
notion_url: https://app.notion.com/p/Objects-vs-Actors-a937a67aea55456199995ffbabf7637d
last_edited: 2026-09-21T17:19:00.000Z
source_url: https://hackernoon.com/objects-vs-actors-6l1g3ex4
tags: ["Programming", "System Design / Software Architecture", "Article", "Hackernoon", "English"]
---
[https://hackernoon.com/objects-vs-actors-6l1g3ex4](https://hackernoon.com/objects-vs-actors-6l1g3ex4)

Readers are assumed to have at least a passing familiarity with an object-oriented language and an actor-oriented language (e.g. Erlang,Pony). Note that objects and actors are orthogonal abstractions andtherefore a language may implement both – Pony is an example.Language abstractions orthogonal to both objects and actors will notbe considered here, e.g. static vs. dynamic typing. Because Ponymixes objects and actors, this paper will use Erlang as the prototypefor actors in order to draw a sharper contrast between objects andactors. Further exposition of Pony will require its own paper.

Fundamentally, software is logic interacting with state. Logic is encapsulated in functions (aka methods) and state is visible to one or morefunctions. When state is visible to multiple simultaneously activefunctions, and one or more of those functions can write that state(aka shared mutable state), accesses must be coordinated in order toguarantee consistent behavior.

Both objects and actors are encapsulation abstractions which define aboundary around logic and state. Figure 1 depicts the structure of atypical object and a typical actor. An object encapsulates one ormore public methods, zero or more private methods, and internalstate. Each method encapsulates logic and local state. All methodsmay call each other and the public methods of other objects to whichthey have a reference. The internal state is visible to all methods,e.g. instance variables. An actor encapsulates one or more functionsand a message queue. Each function encapsulates logic and local stateand all functions may call each other. All functions may receivemessages from the queue and send messages to other actors.

For a given problem domain, the logic and state in an objectimplementation and an actor implementation will be substantially thesame – modulo the abstractions made available by the language. Themost salient differences between objects and actors are the entrypoints across their encapsulation boundary and the sharing of statewithin this boundary.

As depicted in Figure 1, every public method of an object is an entrypoint; actors have only a single entry point in the form of a messagequeue. Within an object there is mutable state visible to multiplefunctions – thus each individual object is potentially an instanceof the shared mutable state problem in microcosm. Within an actorthere is no mutable state visible to multiple functions – themessage queue is visible to multiple functions but it is read only.

Given that public methods of an object can be called at any time in anyorder by any external method with a reference to that object, these public methods can be simultaneously active even within a single thread. Therefore the object’s internal state can change while a method is executing even if that method did not change said state. A simple example of this scenario is depicted in Figure 2. In general, the shared mutablestate problem may occur when two or more instances of methods withinthe same object are simultaneously live, i.e. have frames on thestack. Note that these could be instances of the same method. Inorder to guarantee this behavior does not occur, access coordinationlogic will need to be implemented around the internal state of theobject even for single threaded implementations, e.g. locking. Thisdefinition encompasses recursion but in that situation statemodifications are contained within the same function and areimmediately apparent to the developer.

Messages can be sent to an actor at any time in any order [1] by any actorwith a reference to that actor. Messages are processed serially byany function within the actor; logic within the function determinesthe point at which messages are processed and which specific messagein the queue to process. Because messages are processed serially, thecorresponding logic is executed serially. Thus the actor model doesnot exhibit the shared mutable state problem inherent in the objectmodel.

Even if an object were to be designed to have only a single public methodto mimic the single entry point of actors, that public method canstill be called at any time in any order by any external method witha reference to that object and therefore the shared mutable stateproblem is still present.

For objects, methods external to the object determine when, and in whatorder, logic within that object is activated. For actors, logicinternal to the actor itself determines when it will be activated (byextracting a message from the queue). Thus, fully comprehending thebehavior of an object requires developers to look beyond the objectitself – potentially far beyond. Fully comprehending the behaviorof an actor requires inspection of only the functions within thatactor.

Given that logic within the actor decides when to process the next messageand which message to process, it would be simple to implement logicto guarantee that invariants are maintained before each message isprocessed. Implementing a similar guarantee for objects is morecomplicated.

Because the actor model does not exhibit the shared mutable state problem, an actor implementation that executes correctly in a serial scenario will execute correctly in a concurrent scenario without modification. The same cannot be said for an object implementation.

To gain experience with the actor model, I suggest starting with alanguage for which the actor model is intrinsic. The actor languagewith the largest ecosystem and market footprint is Erlang. Erlangprograms are compiled to a virtual machine that schedules actors.Pony is a promising new actor language that blends actors, objects,and capabilities. Pony is a compiled language and an actor scheduleris linked into the binaries.

When transitioning from an object model with threaded concurrency to anactor model with message passing concurrency, developers will need tomake a mental realignment. Actor creation is almost as fast as afunction call and this changes the optimal approach to scaling.Imagine a system that processes messages requiring a sequence ofsteps. Rather than having permanently running actors for each stepwhere each actor serially executes a step in the message processingsequence, a better approach is to spawn a set of transient actors foreach message when it arrives.

The bottom line: When compared to the actor model, the object model ismore brittle, its encapsulation more permeable, and its cognitiveload on developers higher. To summarize, the advantages of actorsrelative to objects are:

No shared mutable state problem in serial or concurrent scenarios

Zero modification needed to go from serial to concurrent scenarios

Easier implementation of invariant enforcement between logic activations

Lower cognitive load on the developer

[1] Erlang and Pony guarantee that messages sent by actor A to actorB appear in B’s message queue in the order A sent them.
