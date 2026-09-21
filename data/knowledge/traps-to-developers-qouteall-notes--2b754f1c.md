---
title: "Traps to Developers | qouteall notes"
notion_id: 2b754f1c-7d23-816e-b785-d2a72a379a71
notion_url: https://app.notion.com/p/Traps-to-Developers-qouteall-notes-2b754f1c7d23816eb785d2a72a379a71
last_edited: 2025-11-26T19:05:00.000Z
source_url: https://qouteall.fun/qouteall-blog/2025/Traps%20to%20Developers
tags: ["English", "Web Development", "HTML", "CSS", "Programming", "Testing", "Article", "qouteall notes"]
---
A summarization of some traps to developers. There traps are unintuitive things that are easily misunderstood and cause bugs.

## Summarization of traps

### HTML and CSS

- 
-  
-  
-    
- 
- 
- 
- 
- 
- 
-  
- 
- 
- 
- 
- 

### Unicode and text encoding

- Two concepts: code point (rune), grapheme cluster: 
- Different in-memory string behaviors in different languages: 
- Some text files have byte order mark (BOM) at the beginning. For example, EF BB BF is a BOM that denotes the file is in UTF-8 encoding. It's mainly used in Windows. Some non-Windows software does not handle BOM.
- When converting binary data to string, often the invalid places are replaced by � (U+FFFD)
- Normalization. é can be U+00E9 or U+0065 U+0301
- Unicode unification. Different characters in different language use the same code point. Different languages' font variants render the same code point differently. [HTML code](https://github.com/qouteall/qouteall-blog/blob/main/blog/2025/unicode-unification-example.html)

### Floating point

- 
- 
- 
- 
-  
- 
- 
-  
-  
- Leap second. Unix timestamp is "transparent" to leap second, which means that converting between unix timestamp and UTC time ignores leap second. The time measured in unix timestamp will stretch or squeeze near a leap second (leap smear) to hide existence of leap second.
- Time zone. UTC and Unix timestamp is globally uniform and doesn't care about time zone. But human-readable time is time-zone-dependent. It's recommended to store timestamp in database and convert to human-readable time in UI instead of storing human-readable time in database.
- Daylight Saving Time (DST): In some region people adjust clock forward by one hour in warm seasons.
- Time may "go backward" due to NTP sync.
- It's recommended to configure the server's time zone as UTC. Different nodes having different time zones will cause trouble in distributed system. After changing system time zone, the database may need to be reconfigured or restarted.
- There are two clocks: hardware clock and system clock. The hardware clock itself doesn't care about time zone. Linux treats it as UTC by default. Windows treats it as local time by default.

### Java

- `==` compares object reference. Should use `.equals` to compare object content.
- Forget to override `equals` and `hashcode`. It will use object identity equality by default in map key and set.
- Mutate the content of map key object (or set element object) makes the container malfunciton.
- A method that returns `List<T>` may sometimes return mutable `ArrayList`, but sometimes return `Collections.emptyList()` which is immutable. Trying to modify on `Collections.emptyList()` throws `UnsupportedOperationException`.
- A method that returns `Optional<T>` may return `null`.
- Return in `finally` block swallows any exception thrown in the `try` or `catch` block. The method will return the value from `finally`.
- Interrupt. Some libraries ignore interrupt. Interrupt may break class initialization that contains IO.
- Thread pool does not log exception of tasks sent by `.submit()` by default. You can only get exception from the future returned by `.submit()`. Don't discard the future. And `scheduleAtFixedRate` task silently stop if exception is thrown.
- Literal number starting with 0 will be treated as octal number. (`0123` is 83)
- When debugging, debugger will call `.toString()` to local variables. Some class' `.toString()` has side effect, which cause the code to run differently under debugger. This can be disabled in IDE.

### Golang

- `append()` reuses memory region if capacity allows. Appending to a subslice can overwrite parent if they share memory region.
- `defer` executes when the function returns, not when the lexical scope exits.
- `defer` capture mutable variable.
- About `nil`: 

### C/C++

- Storing a pointer of an element inside `std::vector` and then grow the vector, `vector` may re-allocate content, making element pointer no longer valid.
- `std::string` created from literal string may be temporary. Taking `c_str()` from a temporary string is wrong.
- Iterator invalidation. Modifying a container when looping on it.
- `std::remove` doesn't remove but just rearrange elements. `erase` actually removes.
- Literal number starting with 0 will be treated as octal number. (`0123` is 83)
- Undefined behaviors. The compiler optimization aim to keep defined behavior the same, but can freely change undefined behavior. Relying on undefined behavior can make program break under optimization. [See also](https://russellw.github.io/undefined-behavior) 
- Alignment. 

### Python

- Default argument is a stored value that will not be re-created on every call.
- There are subtle differences between numpy and pytorch.

### SQL Databases

- Null is special. 
- Date implicit conversion can be timezone-dependent.
- Complex join with disctinct may be slower than nested query. [See also](https://www.red-gate.com/simple-talk/databases/sql-server/t-sql-programming-sql-server/dont-use-distinct-as-a-join-fixer/)
- In MySQL (InnoDB), if string field doesn't have `character set utf8mb4` then it will error if you try to insert a text containing 4-byte UTF-8 code point.
- MySQL (InnoDB) default to case-insensitive.
- MySQL (InnoDB) can do implicit conversion by default. `select '123abc' + 1;` gives 124.
- MySQL (InnoDB) gap lock may cause deadlock.
- In MySQL (InnoDB) you can select a field and group by another field. It gives nondeterministic result.
- In SQLite the field type doesn't matter unless the table is `strict`.
- Foreign key may cause implicit locking, which may cause deadlock.
- Locking may break repeatable read isolation (it's database-specific).
- Distributed SQL database may doesn't support locking or have weird locking behaviors. It's database-specific.
- If the backend has N+1 query issue, the slowness may won't be shown in slow query log, because the backend does many small queries serially and each individual query is fast.
- Long-running transaction can cause problems (e.g. locking). It's recommended to make all transactions finish quickly.
- Whole-table locks that can make the service temporarily unusable: 
- About ranges: 
- `volatile`: 
- Time-of-check to time-of-use (TOCTOU).
- In SQL database, for special uniqueness constraint that doesn't fit simple unique index (e.g. unique across two tables, conditional unique, unique within time range), and constraint is enforced by application, then: 
- Atomic reference counting (`Arc`, `shared_ptr`) is slow when many threads frequently change the same counter. [See also](https://pkolaczk.github.io/server-slower-than-a-laptop/)
- About read-write lock: some read-write locks doesn't support upgrading from read lock to write lock. Try to write lock when holding read lock may deadlock.

### Common in many languages

- Forget to check for null/None/nil.
- Modifying a container when for looping on it. Single-thread "data race".
- Unintended sharing of mutable data. For example in Python `[[0] * 10] * 10` does not create a proper 2D array.
- For integer `(low + high) / 2` may overflow. A safer way is `low + (high - low) / 2`
- Short circuit. `a() || b()` will not run `b()` if `a()` returns true. `a() && b()` will not run `b()` when `a()` returns false.
- When using profiler: the profiler may by default only include CPU time which excludes waiting time. If your app spends 90% time waiting on database, the flamegraph may not include that 90% which is misleading.

### Linux and bash

- If the current directory is moved, `pwd` still shows the original path. `pwd -P` shows the real path.
- `cmd > file 2>&1` make both stdout and stderr go to file. But `cmd 2>&1 > file` only make stdout go to file but don't redirect stderr.
- File name is case sensitive (unlike Windows).
- There is a capability system for executables, apart from file permission sytem. Use `getcap` to see capability.
- Unset variables. If `DIR` is unset, `rm -rf $DIR/` becomes `rm -rf /`. Using `set -u` can make bash error when encountering unset variable.
- If you want a script to add variables and aliases to current shell, it should be executed by using `source script.sh`, instead of directly executing. But the effect of `source` is not permanent and doesn't apply after re-login. It can be made permanent by putting into `~/.bashrc`.
- Bash has caching between command name and file path of command. If you move one file in `$PATH` then using that command gives ENOENT. Refresh cache using `hash -r`
- Using a variable unquoted will make its line breaks treated as space.
- K8s `livenessProbe` used with debugger. Breakpoint debugger usually block the whole application, making it unable to respond health check request, so it can be killed by K8s `livenessProbe`.

### React

- Modifying state in rendering code.
- Hook used inside if or loop.
- Value not included in `useEffect` dependency array.
- Forget clean up in `useEffect`.
- Closure trap (capturing outdated state).
- Accidentally change data in wrong places (impure component).
- Forgetting `useCallback` cause unnecessary re-render.
- Passing non-memoed things into a memoed component makes memo useless.

### Git

- Rebase can rewrite history. After rebasing local branch, normal push will give weird result (because history is rewritten). Rebase should be used with force push. If remote branch's history is rewritten, pulling should use `-rebase`. 
- Reverting a merge doesn't fully cancel the side effect of the merge. If you merge B to A and then revert, merging B to A again has no effect. One solution is to revert the revert of merge. (A cleaner way to cancel a merge, instead of reverting merge, is to backup the branch, then hard reset to commit before merge, then cherry pick commits after merge, then force push.)
- In GitHub, if you accidentally commited secret (e.g. API key) and pushed to public, even if you override it using force push, GitHub will still record that secret. [Guest Post: How I Scanned all of GitHub’s “Oops Commits” for Leaked Secrets](https://trufflesecurity.com/blog/guest-post-how-i-scanned-all-of-github-s-oops-commits-for-leaked-secrets) [Example activity tab](https://github.com/SharonBrizinov/test-oops-commit/compare/e6533c7bd729957b2eb31e88065c5158d1317c5e...9eedfa00983b7269a75d76ec5e008565c2eff2ef)
- In GitHub, if there is a private repo A and you forked it as B (also private), then when A become public, the private repo B's content is also publicly accessible, even after deleting B. [See also](https://trufflesecurity.com/blog/anyone-can-access-deleted-and-private-repo-data-github).
- `git stash pop` does not drop the stash if there is a conflict.
- It's recommended to add `*/.DS_Store` into `.gitignore` because MacOS auto adds `.DS_Store` files into every folder.

### Networking

- 
- 
- 
- 
- 
- 
- 
- 
-  

### Other

- YAML is space-sensitive, unlike JSON. `key:value` is wrong. `key: value` is correct.
- When using Microsoft Excel to open a CSV file, Excel will do a lot of conversions, such as date conversion (e.g. turn `1/2` and `1-2` into `2-Jan`) and Excel won't show you the original string. [The gene SEPT1 was renamed due to this Excel issue](https://en.wikipedia.org/wiki/SEPTIN1). Excel will also make large numbers inaccurate (e.g. turn `12345678901234567890` into `12345678901234500000`) and won't show you the original accurate number, because Excel internally use floating point for number.
