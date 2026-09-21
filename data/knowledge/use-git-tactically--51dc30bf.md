---
title: "Use Git tactically"
notion_id: 51dc30bf-1504-47ca-a37d-d72932287a76
notion_url: https://app.notion.com/p/Use-Git-tactically-51dc30bf150447caa37dd72932287a76
last_edited: 2023-04-25T13:19:00.000Z
source_url: https://stackoverflow.blog/2022/04/06/use-git-tactically/
tags: ["English", "Programming", "Article", "Stack Overflow Blog"]
---
[https://stackoverflow.blog/2022/04/06/use-git-tactically/](https://stackoverflow.blog/2022/04/06/use-git-tactically/)

[Ed. note: While we take some time to rest up over the holidays and prepare for next year, we are re-publishing our top ten posts for the year. Please enjoy our favorite work this year and we’ll see you in 2023.]

In the movie Free Solo the rock climber Alex Honnold trains to perform a free solo climb of El Capitan, a mountain in Yosemite.

(El Capitan. Photo by Mike Murphy, 2005.)

It's a good movie, but if you haven't seen it, free solo climbing is when you scale a rock face without ropes, harness, or safety equipment. If you lose your grip and fall, you'll die. El Capitan, just to rub it in, is 914 meters of vertical rock. Free-climbing it is an incredible endeavor, but Honnold gets it done by committing to one move at a time (this article is about using Git, after all).

Honnold didn't just free-climb El Capitan. He trained deliberately towards the goal of free climbing El Capitan.

The documentary shows how he repeatedly climbs El Capitan with safety equipment. He plans a route and climbs it several times. On each of the training ascensions, he's using ropes, a harness, and various fasteners for the ropes. When he falls during training, he doesn't fall far, because the rope, harness, and fasteners stop the fall at the last point of fixation.

It's almost like a video game save point.

In one memorable scene, Honnold considers a jump from one position to another. Hundreds of meters in the air, parallel to a vertical rock face. It's a truly precarious maneuver. If he fails, he'll die.

Or, that's true for the free climb. At first, he rehearses the move using rope and harness. This enables him to perform a potentially fatal jump in relative safety. When it goes wrong, he's back at where he fixed his rope, and he may try again.

When you’re making large code changes, even migrating to a new implementation, you can create save points to prevent catastrophes. Like Alex Honold, you can fix your code in place to give you a better chance to get to the next successful build.

When you edit code, you go from one working state to another, but during the process, the code doesn't always run or compile.

Consider an interface like this:

```plain text
public interface IReservationsRepository { Task Create(Reservation reservation); Task<IReadOnlyCollection<Reservation>> ReadReservations( DateTime dateTime); Task<Reservation?> ReadReservation(Guid id); Task Update(Reservation reservation); Task Delete(Guid id); }
```

This, as most of the code in this article, is from my book Code That Fits in Your Head. As I describe in the section on the Strangler Fig pattern, at one point I had to add a new method to the interface. The new method should be an overload of the ReadReservations method with this signature:

Task<IReadOnlyCollection<Reservation>> ReadReservations(DateTime min, DateTime max);

Once you start typing that method definition, however, your code no longer works:

```plain text
Task<IReadOnlyCollection<Reservation>> ReadReservations( DateTime dateTime); T Task<Reservation?> ReadReservation(Guid id);
```

If you're editing in Visual Studio, it'll immediately light up with red squiggly underlines, indicating that the code doesn't parse.

You have to type the entire method declaration before the red squiggly lines disappear, but even then, the code doesn't compile. While the interface definition may be syntactically valid, adding the new method broke some other code. The code base contains classes that implement the IReservationsRepository interface, but none of them define the method you just added. The compiler knows this, and complains:

> Error CS0535 'SqlReservationsRepository' does not implement interface member 'IReservationsRepository.ReadReservations(DateTime, DateTime)'

There's nothing wrong with that. I'm just trying to highlight how editing code involves a transition between two working states:

In Free Solo the entire climb is dangerous, but there's a particularly perilous maneuver that Alex Honnold has to make because he can't find a safer route. For most of the climb, he climbs using safer techniques, moving from position to position in small increments, never losing grip or footing as he shifts his center of gravity.

There's a reason he favors climbing like that. It's safer.

You can't edit code without temporarily breaking it. What you can do, however, is move in small, deliberate steps. Every time you reach a point where the code compiles and all tests pass: commit the changes to Git.

Tim Ottinger calls this a micro-commit. Not only should you commit every time you have a green bar—you should deliberately move in such a way that the distance between two commits is as short as possible. If you can think of alternative ways to change the code, choose the pathway that promises the smallest steps.

Why make dangerous leaps when you can advance in small, controlled moves?

Git is a wonderful tool for maneuverability. Most people don't think of it like that. They start programming, and hours later, they may commit to Git in order to push a branch out.

Tim Ottinger doesn't do that, and neither do I. I use Git tactically.

I'll walk you through an example.

As described above, I wanted to add a ReadReservations overload to the IReservationsRepository interface. The motivation for that is described in Code That Fits in Your Head, but that's not the point here. The point is to use Git to move in small increments.

When you add a new method to an existing interface, the code base fails to compile when you have existing classes that implement that interface. How do you deal with that situation? Do you just forge ahead and implement the new method? Or are there alternatives?

Here's an alternative path that moves in smaller increments.

First, lean on the compiler (as Working Effectively with Legacy Code puts it). The compiler errors tell you which classes lack the new method. In the example code base, it's SqlReservationsRepository and FakeDatabase. Open one of those code files, but don't do anything yet. Instead, copy the new ReadReservations method declaration to the clipboard. Then stash the changes:

$ git stash

Saved working directory and index state WIP on tactical-git: [...]

The code is now back in a working state. Now find a good place to add the new method to one of the classes that implement the interface.

I’ll start with the SqlReservationsRepository class. Once I've navigated to the line in the file where I want to add the new method, I paste in the method declaration:

```plain text
Task<IReadOnlyCollection<Reservation>> ReadReservations(DateTime min, DateTime max);
```

That doesn't compile because the method ends with a semicolon and has no body.

So I make the method public, delete the semicolon, and add curly brackets:

```plain text
public Task<IReadOnlyCollection<Reservation>> ReadReservations(DateTime min, DateTime max) { }
```

This still doesn't compile, because the method declaration promises to return a value, but the body is empty.

What's the shortest way to a working system?

```plain text
public Task<IReadOnlyCollection<Reservation>> ReadReservations(DateTime min, DateTime max) { throw new NotImplementedException(); }
```

You may not want to commit code that throws NotImplementedException, but this is in a brand-new method that has no callers. The code compiles and all tests pass—of course they do: no existing code changed.

Commit the changes:

```plain text
$ git add . && git commit [tactical-git 085e3ea] Add ReadReservations overload to SQL repo 1 file changed, 5 insertions(+)
```

This is a save point. Saving your progress enables you to back out of this work if something else comes up. You don't have to push that commit anywhere. If you feel icky about that NotImplementedException, take comfort that it exists exclusively on your hard drive.

Moving from the old working state to the new working state took less than a minute.

The natural next step is to implement the new method. You may consider doing this incrementally as well, using TDD as you go, and committing after each green and refactor step (assuming you follow the red-green-refactor checklist).

I'm not going to do that here because I try to keep SqlReservationsRepository a Humble Object. The implementation will turn out to have a cyclomatic complexity of 2. Weighed against how much trouble it is to write and maintain a database integration test, I consider that sufficiently low to forgo adding a test (but if you disagree, nothing prevents you from adding tests in this step).

```plain text
public async Task<IReadOnlyCollection<Reservation>> ReadReservations(DateTime min, DateTime max) { const string readByRangeSql = @" SELECT [PublicId], [Date], [Name], [Email], [Quantity] FROM [dbo].[Reservations] WHERE @Min <= [Date] AND [Date] <= @Max"; var result = new List<Reservation>(); using var conn = new SqlConnection(ConnectionString); using var cmd = new SqlCommand(readByRangeSql, conn); cmd.Parameters.AddWithValue("@Min", min); cmd.Parameters.AddWithValue("@Max", max); await conn.OpenAsync().ConfigureAwait(false); using var rdr = await cmd.ExecuteReaderAsync().ConfigureAwait(false); while (await rdr.ReadAsync().ConfigureAwait(false)) result.Add( new Reservation( (Guid)rdr["PublicId"], (DateTime)rdr["Date"], new Email((string)rdr["Email"]), new Name((string)rdr["Name"]), (int)rdr["Quantity"])); return result.AsReadOnly(); }
```

Granted, this takes more than a minute to write, but if you've done this kind of thing before, it probably takes less than ten—particularly if you've already figured the SELECT statement out on beforehand, perhaps by experimenting with a query editor.

Once again, the code compiles and all tests pass. Commit:

```plain text
$ git add . && git commit [tactical-git 6f1e07e] Implement ReadReservations overload in SQL repo 1 file changed, 25 insertions(+), 2 deletions(-)
```

Status so far: We're two commits in, and all code works. The time spent coding between each commit has been short.

The other class that implements IReservationsRepository is called FakeDatabase. It's a Fake Object (a kind of Test Double) that exists only to support automated testing.

The process for implementing the new method is exactly the same as for SqlReservationsRepository. First, add the method:

```plain text
public Task<IReadOnlyCollection<Reservation>> ReadReservations(DateTime min, DateTime max) { throw new NotImplementedException(); }
```

The code compiles and all tests pass. Commit:

```plain text
$ git add . && git commit [tactical-git c5d3fba] Add ReadReservations overload to FakeDatabase 1 file changed, 5 insertions(+)
```

Then add the implementation:

```plain text
public Task<IReadOnlyCollection<Reservation>> ReadReservations(DateTime min, DateTime max) { return Task.FromResult<IReadOnlyCollection<Reservation>>( this.Where(r => min <= r.At && r.At <= max).ToList()); }
```

The code compiles and all tests pass. Commit:

```plain text
$ git add . && git commit [tactical-git e258575] Implement FakeDatabase.ReadReservations overload 1 file changed, 2 insertions(+), 1 deletion(-)
```

Each of these commits represent only a few minutes of programming time; that's the whole point. By committing often, you have granular save points you can retreat to if things start to go wrong.

Keep in mind that we've been adding the methods in anticipation that the IReservationsRepository interface will change. It hasn't changed yet, remember. I stashed that edit.

The new method is now in place everywhere it needs to be in place: both on SqlReservationsRepository and FakeDatabase.

Now pop the stash:

```plain text
$ git stash pop On branch tactical-git Changes not staged for commit: (use "git add <file>..." to update what will be committed) (use "git restore <file>..." to discard changes in working directory) modified: Restaurant.RestApi/IReservationsRepository.cs no changes added to commit (use "git add" and/or "git commit -a") Dropped refs/stash@{0} (4703ba9e2bca72aeafa11f859577b478ff406ff9)
```

This re-adds the ReadReservations method overload to the interface. When I first tried to do this, the code didn't compile because the classes that implement the interface didn't have that method.

Now, on the other hand, the code immediately compiles and all tests pass. Commit.

```plain text
$ git add . && git commit [tactical-git de440df] Add ReadReservations overload to repo interface 1 file changed, 2 insertions(+)
```

We're done. By a tactical application of git stash, it was possible to partition what looked like one long, unsafe maneuver into five smaller, safer steps.

Someone once, in passing, mentioned that one should never be more than five minutes away from a commit. That's the same kind of idea. When you begin editing code, do yourself the favor of moving in such a way that you can get to a new working state in five minutes.

This doesn't mean that you have to commit every five minutes. It's okay to take time to think. Sometimes, I go for a run, or go grocery shopping, to allow my brain to chew on a problem. Sometimes, I just sit and look at the code without typing anything. And sometimes, I start editing the code without a good plan, and that's okay, too... Often, by dawdling with the code, inspiration comes to me.

When that happens, the code may be in some inconsistent state. Perhaps it compiles; perhaps it doesn't. It's okay. I can always reset to my latest save point. Often, I reset by stashing the results of my half-baked experimentation. That way, I don't throw anything away that may turn out to be valuable, but I still get to start with a clean slate.

git stash is probably the command I use the most for increased maneuverability. After that, being able to move between branches locally is also useful. Sometimes, I do a quick-and-dirty prototype in one branch. Once I feel that I understand the direction in which I must go, I commit to that branch, reset my work to a more proper commit, make a new branch and do the work again, but now with tests or other things that I skipped during the prototype.

Being able to stash changes is also great when you discover that the code you're writing right now needs something else to be in place (e.g. a helper method that doesn't yet exist). Stash the changes, add the thing you just learned about, commit that, and the pop the stash. Subsection 11.1.3 Separate Refactoring of Test and Production Code in Code That Fits in Your Head contains an example of that.

I also use git rebase a lot. While I'm no fan of squashing commits, I've no compunction about reordering commits on my local Git branches. As long as I haven't shared the commits with the world, rewriting history can be beneficial.

Git enables you to experiment, to try out one direction, and to back out if the direction begins to look like a dead end. Just stash or commit your changes, move back to a previous save point and try an alternative direction. Keep in mind that you can leave as many incomplete branches on your hard drive as you like. You don't have to push them anywhere.

That's what I consider tactical use of Git. It's maneuvers you perform to be productive in the small. The artifacts of these moves remain on your local hard drive, unless you explicitly choose to share them with others.

Git is a tool with more potential than most people realize. Usually, programmers use it to synchronize their work with others. Thus, they use it only when they feel the need to do that. That's git push and git pull.

While that's a useful and essential feature of Git, if that's all you do, you might as well use a centralized source control system.

The value of Git is the tactical advantage it also provides. You can use it to experiment, make mistakes, flail, and struggle on your local machine, and at any time, you can just reset if things get too hard.

In this article, you saw an example of adding an interface method, only to realize that this involves more work than you may have initially thought. Instead of just pushing through on an ill-planned unsafe maneuver that has no clear end, just back out by stashing the changes so far. Then move deliberately in smaller steps and finally pop the stash.

Just like a rock climber like Alex Honnold trains with ropes and harness, Git enables you to proceed in small steps with fallback options. Use it to your advantage.
