---
title: "UI-First Development"
notion_id: 4825cf65-adae-4e8f-bf58-31ff6b743099
notion_url: https://app.notion.com/p/UI-First-Development-4825cf65adae4e8fbf5831ff6b743099
last_edited: 2026-09-21T17:22:00.000Z
source_url: https://medium.com/newsonthebloc/ui-first-development-94092907ebd6
tags: ["Article", "Medium", "English", "Programming"]
---
[https://medium.com/newsonthebloc/ui-first-development-94092907ebd6](https://medium.com/newsonthebloc/ui-first-development-94092907ebd6)

7 min readNov 17, 2014

--

## Not sure where to start? Just put a button on the page.

When building a website, it’s hard to know where to start. It’s unclear from the get-go what code will be necessary. And designing the model, the view, or the controller in isolation is difficult.

To solve that blank-page problem, try UI-first development. The methodology of UI-first development is simple:

1. Add something to the UI so the user can take action.
2. Fix the resulting error chain.
3. Repeat.

Want users to contact you? Start by adding a page with a textarea on it. Then fix the resulting errors.

Want users to be able to log out? Start by adding a “log out” link to your header. Then fix the resulting errors.

UI-first development has lots of benefits. It:

- Makes you think about the ultimate goal: Your starting point has to be related to the user’s experience.
- Gives you a starting point. The first thing to do is to add something to the UI.
- Tells you what to do next. The next step is always to fix the subsequent error.
- Builds only the things you need. If you limit yourself to fixing the immediate error, you won’t waste time over-engineering the solution.
- Gives you more feedback. Since the UI is always working (or at least broken in a way that makes sense) you can get good feedback from the browser.
- Tells you when you’re done. You’re done when there are no more errors to fix.

A lot of these benefits are familiar from TDD.

## A Complete Example

Let’s go through a more complete example. We’ll use Rails for the example but a similar technique can be used with any web framework.

Check out the github repository if you want to follow along.

We want to build a feature where the user can upvote posts. Our starting point is a simple Rails application showing a list of posts:

Press enter or click to view image in full size

The initial situation

Let’s start from the beginning. What does the user need in order to upvote? A button.

Let’s add a button tag to the markup:

```plain text
<button>↑</button>
```

Refresh the browser. The page now looks like this:

Press enter or click to view image in full size

The upvote button

Let’s click one of the three upvote buttons. It won’t work magically but we still want to know how it fails. In doing this, we are:

- Getting satisfaction from seeing the result of our work.
- Confirming our expectations up to that point.
- Getting information from the subsequent error.
- Asking what we should be doing next.

After clicking the button, we should ask ourselves a few questions. Did anything happen? What was the error? Does the error make sense?

In our case, nothing happened after the click. This is wrong: let’s fix it. We will add an event handler that will submit the upvote request to the server. If we were not using Rails, we could add an onclick to the button and use jQuery’s $.ajax to post the request to the server. In Rails, this can be done by using button_to:

```plain text
<%= button_to ‘↑’, upvote_path(post) %>
```

We think we fixed the previous error so we click the button again. Or at least we try. Upon refreshing the browser, we are greeted with:

Press enter or click to view image in full size

NoMethodError in Posts#index

This error page is the equivalent of a failing test. At this point we want to do the least amount of work that will fix this error.

The error message is useful: it says there is no method named upvote_path. This method would be available if there were a route for our upvote request to go through. Let’s add the missing route to routes.rb:

```plain text
post ‘posts/:id/upvote’ => ‘posts#upvote’, as: ‘upvote’
```

Once the route is in place, we refresh the browser and we gladly see our list of posts again:

Press enter or click to view image in full size

With apprehension, we click the upvote button and are greeted with:

Press enter or click to view image in full size

Unknown action

Fair enough. We created a route to a controller action that does not exist yet. Let’s create it in posts_controller.rb:

```plain text
def upvote @post = Post.find params[:id] @post.upvote @post.saveend
```

This should be enough to carry us through the next step. When we click the button one more time, we are told:

Press enter or click to view image in full size

NoMethodError in PostsController#upvote

There is no method upvote on the Post class: we need to update our model to support voting. The first step would be to create the upvote method on the Post model and increment the number of votes:

```plain text
def upvote self.votes += 1end
```

With this method in place, we think we fixed the previous error. It’s time to click the upvote button again (or just refresh the browser) and we get this new message:

NoMethodError in PostsController#upvote

Every time we get rid of the previous error message and receive a new one, we know we made progress. And we know we only built the necessary components. Let’s keep going.

The new error message is complaining that votes does not exist in the Post class. Let’s create it by running this generator in the terminal:

```plain text
rails generate migration AddVotesToPosts votes:integer
```

Once the migration is generated, let’s open the migration file to confirm that everything is as expected. Let’s also add a default value to the votes attribute:

```plain text
def change add_column :posts, :votes, :integer, default:0end
```

Then we can run rake db:migrate and refresh the page:

Press enter or click to view image in full size

Template is missing

Excellent: we have a new error message. It’s a bonus of UI-first development that it makes us less anxious about errors. We don’t have to see them as a bad thing. We learn to like error messages and expect them every step of the way.

So what is happening now? It looks like Rails is trying to render a template back to the browser. Before we fix this error though, let’s take some time to confirm our expectations: the post should have been upvoted. Let’s open up the Rails console and verify:

```plain text
$ rails c Loading development environment (Rails 4.1.7) >> Post.first.votes Post Load (0.2ms) SELECT “posts”.* FROM “posts” ORDER BY “posts”.”id” ASC LIMIT 1 => 1
```

All is good: our post received the vote. We can click the upvote button again and the 1 will become 2. Expectations confirmed.

Let’s go back to the UI and let’s fix the error about the missing template. The underlying problem is that we are not rendering a proper response back to the browser. The easiest way to fix this is to add a render call to the action method:

```plain text
def upvote @post = Post.find params[:id] @post.upvote @post.save render inline: ‘OK’end
```

This way, when we click the upvote button, we get a simple response in the browser:

Press enter or click to view image in full size

The server responds: OK

Of course, this is not ideal for the user. We want the upvote to happen through AJAX. Let’s add an html option to our button:

```plain text
<%= button_to ‘↑’, upvote_path(post), remote:true %>
```

We’re almost there.

Let’s go back to our list of posts and click the upvote button again. Chrome’s network tab tells us that the request was triggered and the response from the server received. Our post was upvoted. Nothing happened in the UI, not because something is broken, but because things are working: we have come full circle.

Press enter or click to view image in full size

Chrome’s network tab shows the response from the server

Now would be a good time to add some feedback for the user. Perhaps we should show the current number of votes and update that number after the user’s upvote. We could get started on that feature by adding the current number of upvotes to the UI. We would realize that the count is wrong after an upvote from the user. And we would fix that error. And the next error after that.

We could have built these two features at the same time, but it’s a good idea to restrict our efforts to one small feature. It makes it easier to follow the error chain. And it allows for changing our mind about that second feature. The second feature might not be needed after all or it might have to be re-thought.

A long time ago, we added a button to the UI, and then we never had to ask ourselves what we should do next: we were always fixing the next error.

We built only the required functionality. And we could always rely on clicking the upvote button in order to test each new bit of functionality. The UI became our flashlight into the codebase.

Not sure how to get started with a new feature? Just start with the UI.
