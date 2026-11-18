---
title: "What if the browser was the server? - Arthur Cornil"
notion_id: 35e54f1c-7d23-8127-8dc2-d9406c717061
notion_url: https://app.notion.com/p/What-if-the-browser-was-the-server-Arthur-Cornil-35e54f1c7d2381278dc2d9406c717061
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://arthurcornil.com/blog/ship-it-to-the-user/
tags: ["English", "Web Development", "Javascript", "WebAssembly", "Software Development", "Performance", "Article", "Arthur Cornil"]
---
## The SaaS paradigm

SaaS won because a URL beats an installer. And honestly, I love it. You can access software through your web browser in seconds. Great distribution.

However, modern SaaS products are often not much more than a nice interface you use in order to interact with a database. Your computer has become a machine used to render web pages and make HTTP requests to a remote server that will process your data and store it. These companies then justify their 15$/month subscription by the fact that they store your data.

Clearly, this established paradigm is justified for some tools that require real-time collaboration or large computing resources. However, when it comes to personal data and a single user, this architecture becomes more of a constraint than anything else. Think about it: the vast majority of the time performing an action is lost through network communication between your browser and the server.

Additionally you don't benefit in any way from having your data stored and processed on a remote computer. This is like hiring someone to live in another city just to hold your grocery list. Every time you want to add "milk" or check what you need, you have to call them, wait for them to pick up, tell them the update, and wait for them to read it back to you.

## Ship the backend to the user

The average computer has become increasingly capable. Why should we offload work to a remote computer? What if the browser **was** the server?

If a backend is simply code that processes data and a database, we can actually ship the backend to the user. No authentication required. No server holding personal data and no HTTP request for each CRUD action.

Here are the two main things that make this possible:

- **WASM (WebAssembly)**: You can compile any language to a binary the browser can run
- **OPFS (Origin private file system)**: The browser now has a real file system API with persistent storage

`sqlite.org/sqlite-wasm` is a version of SQLite that is built to run in the browser using WASM, making it easier to interact with using JS. However, thanks to WASM once again, **you don't have to stick to JavaScript** for the "_backend_". You could litterally write it in C, link against SQLite and ship the compiled `.wasm` file to the browser.

## The architecture

Because we are running the software in a browser, we will have to deal with JavaScript, at least to some extent.

Historically JS has been single-threaded. However, browsers gave us a way to run scripts in the background, separate from the main execution thread. How perfect would that be for someone looking to run a backend and a database from within the user's browser tab?

You could have your frontend and backend run independently, so that big queries and resource demanding processes don't freeze the UI.

The main execution thread and the worker can communicate using `postMessage()`. Here is a communication pattern:

```plain text
Main thread                    Worker
     │                            │
     │  postMessage({             │
     │    type: 'performAction',  │
     │    payload: { ... }        │
     │  })                        │
     │ ─────────────────────────► │
     │                            │ onmessage fires
     │                            │ calls to the WASM-compiled backend.
     │                            │ It performs the SQL work
     │                            │
     │          postMessage({     │
     │            type: 'result'  │
     │            data: { ... }   │
     │          })                │
     │ ◄───────────────────────── │
     │                            │
onmessage fires                   │
UI updates                        │
```

## The Proof of Concept

I made a very simple Time Tracker in order to demonstrate this paradigm. I used Anthropic's _Sonnet 4.6_ to generate a UI (CSS and Vanilla JS). I then wrote a backend in Go to demonstrate that you can ship a backend written in _any_ language.

I used `ncruces/go-sqlite3` module, a wrapper for a WASM built SQLite instance. This way, the Go backend owns SQLite.

The first challenge was to store a `.db` file on the user's machine using _OPFS_. `ncruces/go-sqlite3` does not yet provide an easy way to use OPFS as a file system. Fortunately, I found `danmestas/go-sqlite3-opfs`, an OPFS VFS for `ncruces/go-sqlite3`.

```plain text
Web Worker
│
│  pre-opens OPFS file handles
│  passes them to the backend
│
▼
Go WASM binary
│
│  sql.Open("file:track.db?vfs=opfs")
│  thinks it's a normal file
│
▼
VFS shim (danmestas/go-sqlite3-opfs)
│
│  createSyncAccessHandle()
│
▼
OPFS — track.db lives here, on the user's actual disk
```

Here is how I handled it from the Web Worker:

```plain text
const root = await navigator.storage.getDirectory();

// Getting handles for each SQLite related file
const handles = {};
for (const suffix of ["", "-journal", "-wal"]) {
    const name = "track.db" + suffix;
    const fh = await root.getFileHandle(name, { create: true });
    handles[name] = await fh.createSyncAccessHandle();
}

await fetch('wasm_exec.js')
    .then(r => r.text())
    .then(text => eval(text));
const go = new Go();
const result = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject);
go.run(result.instance);

// Hand the pre-opened file handles to the VFS shim
_opfs_init(handles);

// Exposed Go function that initiates a SQLite instance
initConnection();
```

A very basic and trivial way to initiate SQLite would be to use a `CREATE IF NOT EXISTS` statement.

```plain text
func initConnection(this js.Value, args []js.Value) any {
	var err error
	db, err = sql.Open("sqlite3", "file:track.db?vfs=opfs")
	if err != nil {
		log.Fatal("failed to open db:", err)
	}

	db.SetMaxOpenConns(1)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS entries (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		name		TEXT,
		started_at  INTEGER NOT NULL,
		ended_at    INTEGER NOT NULL,
		created_at  INTEGER NOT NULL
	)`)
	if err != nil {
		log.Fatal("failed to create table:", err)
	}
	return nil
}
```

Simply create database schema if the `.db` file doesn't have it yet. This is fine for a simple proof of concept, but the best option for a real project would be to implement a migration system.

What's left now is for the frontend to trigger specific actions in the backend. Like adding a new entry in the time tracker.

The main execution thread informs the worker that the user wants to add an entry.

```plain text
worker.postMessage({
    action: 'addEntry',
    payload: {
        name: timerDesc.value.trim() || 'Untitled',
        started_at: startedAt,
        ended_at: endedAt,
    }
});
```

The worker receives the message and simply calls the appropriate Go function. Here is what the `addEntry` Go function looks like:

```plain text
func addEntry(this js.Value, args []js.Value) any {
	payload := args[0].String()

	var entry Entry
	if err := json.Unmarshal([]byte(payload), &entry); err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	entry.CreatedAt = time.Now().UnixMilli()
	_, err := db.Exec(`INSERT INTO entries (name, started_at, ended_at, created_at) VALUES (?, ?, ?, ?)`,
		entry.Name,
		entry.StartedAt,
		entry.EndedAt,
		entry.CreatedAt,
	)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return "ok"
}
```

Once this function executed, the worker will `postMessage()` to the main execution thread, to tell it the action executed and that it can now be reflected accordingly in the UI.

Repo: [https://github.com/arthurcornil/localtrack](https://github.com/arthurcornil/localtrack)

## Limitations

As stated in the introduction, this architecture is simply not applicable to software that require constant real-time sync. For softwares that are based on deep real-time collaboration like Figma, Google Docs, Slack, ... the SaaS architecture simply makes perfect sense.

With the explored architecture, the user's data lives entirely in OPFS. Even though that makes the data survive browser relaunches and machine reboots, the data is intrinsically bound to the browser. Switching browser would result in an entirely new database. That's why an "Export" and "Import" data feature might be absolutely necessary for any real project.

## Conclusion

I noticed, after building the time tracker POC, that [Notion actually uses this paradigm in production](https://www.notion.com/blog/how-we-sped-up-notion-in-the-browser-with-wasm-sqlite). This shows that you can ship software this way.

The established SaaS paradigm is the default choice, even when it isn't the _right_ choice, because it's the path of least resistance. But let's question what's established. A finance tracker, a journaling app, a habit tracker; none of these need a server. Ship the software to the user. Not just the UI.
