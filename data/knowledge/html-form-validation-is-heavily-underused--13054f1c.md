---
title: "HTML Form Validation is heavily underused"
notion_id: 13054f1c-7d23-8160-a0aa-e8ff39973914
notion_url: https://app.notion.com/p/HTML-Form-Validation-is-heavily-underused-13054f1c7d238160a0aae8ff39973914
last_edited: 2024-11-15T20:09:00.000Z
source_url: https://expressionstatement.com/html-form-validation-is-heavily-underused
tags: ["English", "Web Development", "HTML", "Article", "Expression Statement"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

HTML Forms have powerful validation mechanisms, but they are heavily underused. In fact, not many people even know much about them. Is this because of some flaw in their design? Let’s explore.

## Attributes, methods, and properties

It’s easy to disallow empty inputs by adding a [`required`](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/required) attribute:



```html
<input required={true} />
```

Beyond that, there is a bunch of other ways that you can add constraints to your input. Precisely, there are three ways to do it:

- Using specific `type` attribute values, such as `"email"`, `"number"`, or `"url"`
- Using other input attributes that create constraints, such as `"pattern"` or `"maxlength"`
- **Using the ****`setCustomValidity`**** DOM method of the input**

The last one is the most powerful as it allows to create arbitrary validation logic and handle complex cases. Do you notice how it differs from the first two techniques? The first two are defined with _attributes_, but `setCustomValidity` is a _method_.

Here’s a great write-up that explains the differences between DOM attributes and properties: [https://jakearchibald.com/2024/attributes-vs-properties/](https://jakearchibald.com/2024/attributes-vs-properties/)

## The nuance of an imperative API

The fact that `setCustomValidity` API is exposed only as a method and doesn’t have an attribute equivalent leads to some terrible ergonomics. I’ll show you with an example.

But first, a very quick intro to how this API works:



```javascript
// Make input invalid
input.setCustomValidity("Any text message");
```

This would make input _invalid_ and the browser will show the reason as “Any text message”.



```javascript
// Remove custom constraints and make input valid
input.setCustomValidity("");
```

Passing an empty string makes the input _valid_ (unless other constraints are applied).

That’s pretty much it! Now let’s apply this knowledge.

Let’s say we want to implement an equivalent of the `required` attribute. That means that an empty input must be prevent the form from being submitted.



```html
<input
  name="example"
  placeholder="..."
  onChange={(event) => {
    const input = event.currentTarget;
    if (input.value === "") {
      input.setCustomValidity("Custom message: input is empty");
    } else {
      input.setCustomValidity("");
    }
  }}
/>
```

This kind of looks like we’re done and this code should be enough to accomplish the task. But try to see it in action:

It may seem to work, but there’s just one important edge case: the input is in a valid state _initially_. If you reset the component and press the “submit” button, the form submission will go through. But surely, before we ever touch the input, it is empty, and therefore must be invalid. But we only ever do something when the input value _changes_.

How can we fix this?

Let’s execute some code when the component mounts:



```javascript
import { useRef, useLayoutEffect } from "react";

function Form() {
  const ref = useRef();
  useLayoutEffect(() => {
    // Make input invalid on initial render if it's empty
    const input = ref.current;
    const empty = input.value !== "";
    input.setCustomValidity(empty ? "Initial message: input is empty" : "");
  }, []);

  return (
    <form>
      <input
        ref={ref}
        name="example"
        onChange={(event) => {
          const input = event.currentTarget;
          if (input.value === "") {
            input.setCustomValidity("Custom message: input is empty");
          } else {
            input.setCustomValidity("");
          }
        }}
      />
      <button>Submit</button>
    </form>
  );
}
```

Great! Now everything works as expected. But at what cost?

## The boilerplate problem

Let’s look at our clumsy way to validate the initial value:



```javascript
const ref = useRef();
useLayoutEffect(() => {
  // Make input invalid on initial render if it's empty
  const input = ref.current;
  const empty = input.value !== "";
  input.setCustomValidity(empty ? "Initial message: input is empty" : "");
}, []);
```

Ugh! Wouldn’t want to write that one each time. Let’s think about what’s wrong with this.

- The validation logic is duplicated between the onChange handler and the initial render phase
- The initial validation is not co-located with the input, so we’re losing code cohesion. It’s fragile: if you update validation logic, you might forget to update code in both places.
- The `useRef` + `useLayouEffect` + `onChange` combo is just too much ceremony, especially when a form has a lot of inputs. And it gets even more confusing if only some of those inputs use `customValidity`

This is what happens when you deal with a purely imperative API in a declarative component.

> Unlike validation attributes, `CustomValidity` is a purely imperative API. In other words, there’s no input attribute that we can use to set custom validity.

In fact, I would argue that this is **the main reason for poor adoption of native form validation**. If the API is cumbersome, sometimes it just does not matter how powerful it is.

## The missing part

In essence, this is the attribute we need:



```html
<input custom-validity="error message" />
```

In a declarative framework, this would allow to define input validations in a very powerful way:



```javascript
function Form() {
  const [value, setValue] = useState();
  const handleChange = (event) => setValue(event.target.value);
  return (
    <form>
      <input
        name="example"
        value={value}
        onChange={handleChange}
        custom-validity={value.length ? "Fill out this field" : ""}
      />
      <button>Submit</button>
    </form>
  );
}
```

Pretty cool! In my opinion, at least. Though you can rightfully argue that this accomplishes only what the existing `required` attribute is already capable of. Where’s the “power”?

Let me show you, but first, since there’s no actual `custom-validity` currently in the [HTML Spec](https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#the-constraint-validation-api), let’s implement it in userland.



```javascript
function Input({ customValidity, ...props }) {
  const ref = useRef();
  useLayoutEffect(() => {
    if (customValidity != null) {
      const input = ref.current;
      input.setCustomValidity(customValidity);
    }
  }, [customValidity]);

  return <input ref={ref} {...props} />;
}
```

This will work well for our demo purposes.

For a production-ready component check out a more [complete implementation](https://gist.github.com/everdimension/a5c1e991a8a6b6aab060ce349b37b825).

## The power

Now we’ll explore which non-trivial cases this design can help solve.

In real-world apps, validation often gets more complex than local checks. Imagine a username input that should be **valid only if the username is not taken**. This would require async calls to your server and an intermediary state: the form should not be valid while the check is in progress. Let’s see how our abstraction can handle this.

Username

Helpers:

Play around with this example. It uses the `required` to prevent empty inputs. But then it relies on `customValidity` to mark input as invalid during the loading state and based on the response.

### Implementation

First, we create an async function to check whether the username is unique that imitates a server request with a delay.



```javascript
export async function verifyUsername(userValue) {
  // imitate network delay
  await new Promise((r) => setTimeout(r, 3000));
  const value = userValue.trim().toLowerCase();
  if (value === "bad input") {
    throw new Error("Bad Input");
  }
  const validationMessage = value === "taken" ? "Username is taken" : "";
  return { validationMessage };
}
```

Next, we’ll create a controlled form component and use [react-query](https://tanstack.com/query/latest) to manage to server request when the input value changes:



```javascript
import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { verifyUsername } from "./verifyUsername";
import { Input } from "./Input";

function Form() {
  const [value, setValue] = useState("");
  const { data, isLoading, isError } = useQuery({
    queryKey: ["verifyUsername", value],
    queryFn: () => verifyUsername(value),
    enabled: Boolean(value),
  });

  return (
    <form>
      <Input
        name="username"
        required={true}
        value={value}
        onChange={(event) => {
          setValue(event.currentTarget.value);
        }}
      />
      <button>Submit</button>
    </form>
  );
}
```

Great! We have the setup in place. It consists of two crucial parts:

- Verification request state managed by `useQuery`
- Our custom `<Input />` component that is capable of taking the `customValidity` prop

Let’s put those pieces together:



```javascript
import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { verifyUsername } from "./verifyUsername";
import { Input } from "./Input";

function Form() {
  const [value, setValue] = useState("");
  const { data, isLoading, isError } = useQuery({
    queryKey: ["verifyUsername", value],
    queryFn: () => verifyUsername(value),
    enabled: Boolean(value),
  });

  const validationMessage = data?.validationMessage;

  return (
    <form>
      <Input
        name="username"
        required={true}
        customValidity={
          isLoading
            ? "Verifying username..."
            : isError
            ? "Could not verify"
            : validationMessage
        }
        value={value}
        onChange={(event) => {
          setValue(event.currentTarget.value);
        }}
      />
      <button>Submit</button>
    </form>
  );
}
```

That’s it! We’re describing the whole async validation flow, including loading, error and success states, _in one attribute_. You can go back to see [the result](https://expressionstatement.com/html-form-validation-is-heavily-underused#example-async-username) again if you wish

### One more

This one will be shorter, but also interesting, because it covers dependent input fields. Let’s implement a form that requires to repeat the entered password:



```javascript
import { useState } from "react";
import { Input } from "./Input";

function ConfirmPasswordForm() {
  const [password, setPassword] = useState("");
  const [confirmedPass, setConfirmedPass] = useState("");

  const matches = confirmedPass === password;
  return (
    <form>
      <Input
        type="password"
        name="password"
        required={true}
        value={password}
        onChange={(event) => {
          setPassword(event.currentTarget.value);
        }}
      />
      <Input
        type="password"
        name="confirmedPassword"
        required={true}
        value={confirmedPass}
        customValidity={matches ? "" : "Password must match"}
        onChange={(event) => {
          setConfirmedPass(event.currentTarget.value);
        }}
      />
      <button>Submit</button>
    </form>
  );
}
```

You can try it out:

Enter Password Repeat Password

## Conclusion

I hope I’ve been able to show you how `setCustomValidity` can cover validation needs of all kinds.

But the real power comes from great APIs.

And hopefully, you are now equipped with one of those.

And even more hopefully, we will see it natively in the HTML Spec one day.
