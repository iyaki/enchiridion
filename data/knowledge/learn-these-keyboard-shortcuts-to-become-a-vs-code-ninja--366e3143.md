---
title: "Learn these keyboard shortcuts to become a VS Code ninja"
notion_id: 366e3143-0bdb-46ff-90ad-7e35c40cb401
notion_url: https://app.notion.com/p/Learn-these-keyboard-shortcuts-to-become-a-VS-Code-ninja-366e31430bdb46ff90ad7e35c40cb401
last_edited: 2022-12-21T17:19:00.000Z
source_url: https://blog.logrocket.com/learn-these-keyboard-shortcuts-to-become-a-vs-code-ninja/
tags: ["Programming", "Productivity", "Article", "LogRocket Blog", "English"]
---
![image](https://blog.logrocket.com/wp-content/uploads/2019/08/keyboard-shortcuts-vs-code-ninja-nocdn.jpg)

Recently, I wanted to limit mouse usage when programming in Visual Studio Code since I found interacting with the IDE through a cursor distracting and a major flow-breaker — so, I tried navigating VSC with keyboard alone.

Here, I would like to present some of the shortcuts that I have found to best increase productivity. Go ahead, open Visual Studio Code and let’s get started.

## **Splitting and focusing**

Unless you are working on a very small screen, chances are you split your IDE into two or three views to switch more smoothly between files.

### **Splitting**

To split the editor, you can use `ctrl + \` (`⌘ + \`).

![image](https://blog.logrocket.com/wp-content/uploads/2019/07/splitting-code-editor.png)

There’s no limit to how many times you can split the editor, but I doubt you will ever want to have more than three views open; it is just not practical beyond that. You can switch between views using `ctrl + 1` (`⌘ + 1`), `ctrl + 2` (`⌘ + 2`), and so on. Alternatively, you can switch between tabs (and, by extension, between views) using `ctrl + page up` / `page down` (`⌘ + page up` / `page down`).

### **Focusing**

Having split the editor, we ended up with the same file open in multiple views. Now we would like to focus the explorer panel. We would like to change this without touching the mouse. Also, it would be nice to focus different views without touching the mouse, too.

To focus the explorer panel we use `ctrl + 0` (`⌘ + 0`). We navigate the panel using the up and down arrows. Using the `enter` key, we open a selected folder or file.

By default, there are two views: the explorer panel and the code view. The code view can be focused by using `ctrl + 1` (`⌘ + 1`). If we create more views by splitting the editor, we use `ctrl + 2` (`⌘ + 2`), `ctrl + 3` (`⌘ + 3`), and so on for the respective views. We can also switch between consecutive tabs with `ctrl + page up` / `page down` (by default, this command isn’t defined on macOS).

Note that this will only work when VSC has access to the whole folder, and only when you are working with an open folder — not individual files.

![image](https://blog.logrocket.com/wp-content/uploads/2019/07/focus-explorer-panel.png)

## **Alternative approach**

There’s also a slightly different approach to selecting files that are farther in the list from the one currently open. We can use `ctrl + p` (`⌘ + p`), which opens up a search bar where we type in either a filename (`http.service.ts`) or a full path (`src/services/http.service.ts`).

![image](https://blog.logrocket.com/wp-content/uploads/2019/07/selecting-files-alt.png)

## **Using the file history**

We often don’t work with all the files in the project at once; we simultaneously work with two, maybe three at most. To shorten the time it takes to switch between them (if we don’t have enough screen space to split the editor), we can use the file history.

---

[**Over 200k developers use LogRocket to create better digital experiences**](https://lp.logrocket.com/blg/learn-more)[**Learn more →**](https://lp.logrocket.com/blg/learn-more)

![image](https://blog.logrocket.com/wp-content/uploads/2022/11/Screen-Shot-2022-09-22-at-12.50.47-PM.png)

![image](https://blog.logrocket.com/wp-content/uploads/2022/08/rocket-button-icon.png)

---

File history, as the name implies, saves the files we last used and provides a quick way to restore them. In VSC we use `ctrl + tab` to switch between the last opened files.

While this is indeed very efficient, there is yet another way which, one might argue, is even faster. By using `alt + left` / `right` arrows (`ctrl + shift + -` / `ctrl + -`) we can switch directly to the previous/next file in the file history.

## **Traversing the code**

Now that we know how to navigate the files, let’s focus on the way we move around the code.

### **Making use of an outline**

In the explorer panel, you can click the **Outline** section to have a code’s outline displayed. While the feature itself is amazing in that it lets us see a more general view of the code, we can also use it to quickly move around.

By using `ctrl + shift + o` (`⌘ + shift + o`), we can bring up the command palette, where we can choose the part of the outline we would like to jump to. After choosing the definition with an up/down arrow, the appropriate piece of code is highlighted for us, making it easier to get where we want to go.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/explorer-panel-outline-section.png)

The same feature could also be used to search the whole project for a given piece of code. By using `ctrl + t` (`⌘ + t`), we again bring up the command palette, where we can now type the name of a variable/function/etc. to search for.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/searching-function.png)

### **Straight to a given line**

Imagine we want to jump straight to a specific line — for example, when there’s an error pointing to it. To jump to a line with a specified index, we can use `ctrl + g`.

---

### **More great articles from LogRocket:**

- Don't miss a moment with [The Replay](https://lp.logrocket.com/subscribe-thereplay), a curated newsletter from LogRocket
- [Learn](https://blog.logrocket.com/rethinking-error-tracking-product-analytics/) how LogRocket's Galileo cuts through the noise to proactively resolve issues in your app
- Use React's useEffect [to optimize your application's performance](https://blog.logrocket.com/understanding-react-useeffect-cleanup-function/)
- Switch between [multiple versions of Node](https://blog.logrocket.com/switching-between-node-versions-during-development/)
- [Discover how to animate](https://blog.logrocket.com/animate-react-app-animxyz/) your React app with AnimXYZ
- [Explore Tauri](https://blog.logrocket.com/rust-solid-js-tauri-desktop-app/), a new framework for building binaries
- Compare [NestJS vs. Express.js](https://blog.logrocket.com/nestjs-vs-express-js/)

---

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/jumping-specific-line.png)

### **Jumping back**

Often we want to fix something quickly in one place of the code and jump right back to where we were before. This we do using `ctrl + u` (`⌘ + u`), which takes the cursor back to where it was prior to the jump.

### **Beginning and end of a file**

To jump to the beginning or the end of a file we can use `ctrl + home` (`⌘ + up`) and `ctrl + end` (`⌘ + down`) respectively.

## **Definitions and references**

Have you ever searched for a definition by hand or by `ctrl + shift + f` (`⌘ + shift + f`)? If you have, then you know how annoying that can be. VSC has a great shortcut for that!

### **Jumping to the definition**

We can jump to the definition of a function or a variable currently highlighted using `F12`.

### **Peeking at the implementation**

Often we only want to have a quick peak at the implementation of, say, a function. Ideally, we don’t want to open another file just to check a few lines. By using `alt + F12` (`option + F12`), we can peek at the implementation of a highlighted function right there next to the cursor. Once we’re done, we just press `esc`.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/peeking-at-implementation.png)

### **Peeking at references**

There’s also a shortcut for peeking at references of a symbol in a similar manner — right next to the cursor. We do this with `shift + F12` (`⌘ + k` and `F12`).

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/peeking-at-references.png)

In both cases, we can use the up and down arrows to choose the definition we would like to see or jump to.

### **Changing the name of a symbol**

Changing the name of a given symbol (e.g., the name of a function) throughout the whole project can be tedious. It is usually done with `ctrl + shift + f` (`⌘ + shift + f`) — we manually replace each usage of the symbol.

This can be done faster with the `F2` shortcut. It prompts a window where we type the new name of a currently highlighted symbol, and that’s it — every occurrence has now been replaced with the new name.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/changing-name-symbol.png)

### **Taking a closer look at errors**

When there’s something wrong with our code, VSC underlines it with a red line. Usually, we could just hover over the code with the mouse cursor and see what’s wrong. We can, however, do it much faster by using `F8`.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/examining-errors.png)

We can leave the “error mode” by clicking the `esc` key.

## **Intellisense**

### **Hover**

As was the case with the errors, when we hover over a symbol, VSC shows us its simplified definition. To achieve the same result with the keyboard, we have to set up our own shortcut.

We can set our own shortcuts by using `ctrl + k` (`⌘ + k`) and then `ctrl + s` (`⌘ + s`), which will open the shortcuts settings view.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/setting-show-hover-shortcut.png)

Then search for the **Show hover** action:

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/setting-custom-shortcut.png)

And set it to your preferred shortcut. I have chosen `alt + shift + s`.

The shortcut in action:

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/show-hover-shortcut-in-action.png)

### **Showing recommended actions**

Sometimes VSC is able to fix our problems by, for example, importing the code we forgot to import ourselves or removing unused code. To see the available actions for the currently highlighted code, we can use `ctrl + .` (`⌘ + .`).

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/showing-recommended-actions.png)

## **Selecting code**

Code is made of blocks, be it a body of a function or an `if` block. Sometimes we want to select the whole thing and, say, remove it without worrying about where the block begins and ends.

The `alt + shift + left` / `right` (`⌘ + ctrl + shift + left` / `right`) shortcut makes it a breeze to select pieces of code based on scope. Repeated use of this shortcut makes the selection appropriately bigger or smaller.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/selecting-code.png)

## **Integrated terminal**

With Visual Studio Code opened in the full-screen mode, it is often convenient to have a terminal right there with us. Switching between the terminal and code calls for a few shortcuts of its own.

### **Opening a terminal**

To open a terminal, we use `ctrl + ``.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/opening-terminal.png)

To open more terminals, we use `ctrl + shift + ``.

### **Splitting it up**

Terminal, just like the editor, can be split up into panels. For this we use `ctrl + shift + 5`.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/splitting-terminal.png)

### **Focusing**

To focus on a terminal, while in the editor, we use `ctrl + ``. If we use `ctrl + `` while the terminal is focused, we can toggle its state from shown to hidden.

### **Focusing split panels**

Once we are focused on the terminal, we can use `alt + left`/`right` (`⌘ + option + left` / `right`) to switch focus between split panels.

### **Bonus**

Here are some terminal shortcuts I found to be very helpful.

### **Killing a terminal**

Killing a terminal can be achieved by clicking the trash icon in the top-right corner of the terminal, but in order for it to be a mouse-free experience, we have to set up a shortcut.

While in the shortcuts settings input, type in “workbench.action.terminal.kill” and then click on it to set up the shortcut. I have chosen to use `ctrl + shift + x`, but whatever works for you is fine.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/killing-terminal-shortcut.png)

### **Maximizing a terminal**

Sometimes, when there’s a lot of logs coming in, we would like to temporarily make the terminal larger. Same spiel as before, but in the settings, type `workbench.action.toggleMaximizedPanel`. Here, I have chosen to put it under `ctrl + shift + q`.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/maximizing-terminal-shortcut.png)

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/terminal-maximized.png)

## **Everything else**

In case you don’t know what the shortcut is for something, you can always open up the command palette with `ctrl + shift + p` (`⌘ + shift + p`) and type in what you want the shortcut to do, e.g., “open terminal.” Most of the time, it will show you the correct action with the shortcut next to the name.

![image](https://blog.logrocket.com/wp-content/uploads/2019/08/open-new-terminal-shortcut.png)

## **Summary**

The key to mastering these shortcuts is consistency. Try to implement them gradually, and before you know it, you’ll find yourself relying less and less on the mouse, which in turn will make your coding experience much smoother and more efficient.

Want to learn more? Here’s Visual Studio Code’s [documentation](https://code.visualstudio.com/docs/getstarted/keybindings).
