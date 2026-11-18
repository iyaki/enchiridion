---
title: "Detect Caps Lock with JavaScript"
notion_id: b5ced514-459f-4bea-836a-6ec568b95d1b
notion_url: https://app.notion.com/p/Detect-Caps-Lock-with-JavaScript-b5ced514459f4bea836a6ec568b95d1b
last_edited: 2024-03-25T14:20:00.000Z
source_url: https://davidwalsh.name/detect-caps-lock
tags: ["English", "Javascript", "Article", "Tutorial", "David Walsh Blog"]
---
Anyone is capable of having their caps lock key on at any given time without realizing so. Users can easily spot unwanted caps lock when typing in most inputs, but when using a `password` `input`, the problem isn't so obvious. That leads to the user's password being incorrect, which is an annoyance. Ideally developers could let the user know their caps lock key is activated.

To detect if a user has their keyboard's caps lock turn on, we'll employ `KeyboardEvent`'s `getModifierState` method:

```plain text
document.querySelector('input[type=password]').addEventListener('keyup', function (keyboardEvent) {
    const capsLockOn = keyboardEvent.getModifierState('CapsLock');
    if (capsLockOn) {
        // Warn the user that their caps lock is on?
    }
});

```

I'd never seen `getModifierState` used before, so I explored the [W3C documentation](https://w3c.github.io/uievents/#event-modifier-initializers) to discover other useful values:

```plain text
dictionary EventModifierInit : UIEventInit {
  boolean ctrlKey = false;
  boolean shiftKey = false;
  boolean altKey = false;
  boolean metaKey = false;

  boolean modifierAltGraph = false;
  boolean modifierCapsLock = false;
  boolean modifierFn = false;
  boolean modifierFnLock = false;
  boolean modifierHyper = false;
  boolean modifierNumLock = false;
  boolean modifierScrollLock = false;
  boolean modifierSuper = false;
  boolean modifierSymbol = false;
  boolean modifierSymbolLock = false;
};

```

`getModifierState` provides a wealth of insight as to the user's keyboard during key-centric events. I wish I had known about `getModifier` earlier in my career!
