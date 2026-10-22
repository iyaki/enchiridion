---
title: "Sandi Metz' Rules For Developers"
notion_id: 9dae2b54-2962-4f33-a209-20f46d0c4289
notion_url: https://app.notion.com/p/Sandi-Metz-Rules-For-Developers-9dae2b5429624f33a20920f46d0c4289
last_edited: 2026-09-21T17:35:00.000Z
source_url: https://thoughtbot.com/blog/sandi-metz-rules-for-developers
tags: ["Programming", "Article", "thoughtbot Blog", "English"]
---
[https://thoughtbot.com/blog/sandi-metz-rules-for-developers](https://thoughtbot.com/blog/sandi-metz-rules-for-developers)

Back in January, Sandi Metz introduced her rules for developers in a Ruby Rogues podcast episode episode. Around the time Sandi’s rules were published, the team I am on was starting a new project. This post details the experience of that team applying Sandi’s rules to the new application.

## The rules

There are four rules.

Here are the rules:

1. Classes can be no longer than one hundred lines of code.
2. Methods can be no longer than five lines of code.
3. Pass no more than four parameters into a method. Hash options are parameters.
4. Controllers can instantiate only one object. Therefore, views can only know about one instance variable and views should only send messages to that object (@object.collaborator.value is not allowed).

## When to break these rules

Paraphrasing Sandi, “You should break these rules only if you have a good reason or your pair lets you.” Your pair or the person reviewing your code are the people who you should ask.

Think of this as rule zero. It is immutable.

## 100-line classes

Despite the large number of private methods we wrote, keeping classes short proved easy. It forced us to consider what the single responsibility of our class was, and what should be extracted.

This applied to specs as well. In one case, we found a spec file ran over the limit which helped us realize we were testing too many features. We split the file into a few, more focused, feature specs.

That made us realize that git diffs wouldn’t necessarily show us when we exceed 100 lines.

## Five lines per method

Limiting methods to five lines per method is the most interesting rule.

We agreed if, else, and end are all lines. In an if block with two branches, each branch could only be one line.

For example:

```plain text
def validate_actor if actor_type == 'Group' user_must_belong_to_group elsif actor_type == 'User' user_must_be_the_same_as_actor end end
```

Five lines ensured that we never use else with elsif.

Having only one line per branch urged us to use well-named private methods to get work done. Private methods are great documentation. They need very clear names, which forced us to think about the content of the code we were extracting.

## Four method arguments

The four method arguments rule was particularly challenging in Rails, and particularly in the views.

View helpers such as link_to or form_for can end up requiring many parameters to work correctly. While we put some effort into not passing too many arguments, we fell back to Rule 0 and left the parameters if we couldn’t find a better way to do it.

## Only instantiate one object in the controller

This rule raised the most eyebrows before we started the experiment. Often, we needed more than one type of thing on a page. For example, a homepage needed both an activity feed and a notification counter.

We solved this using the Facade Pattern. It looked like this:

app/facades/dashboard.rb:

```plain text
class Dashboard def initialize(user) @user = user end def new_status @new_status ||= Status.new end def statuses Status.for(user) end def notifications @notifications ||= user.notifications end private attr_reader :user end
```

app/controllers/dashboards_controller.rb:

```plain text
class DashboardsController < ApplicationController before_filter :authorize def show @dashboard = Dashboard.new(current_user) end end
```

app/views/dashboards/show.html.erb:

```plain text
<%= render 'profile' %> <%= render 'groups', groups: @dashboard.group %> <%= render 'statuses/form', status: @dashboard.new_status %> <%= render 'statuses', statuses: @dashboard.statuses %>
```

The Dashboard class provided a common interface for locating the user’s collaborator objects and we passed the dashboard’s state to view partials.

We didn’t count instance variables in controller memoizations toward the limit. We used a convention of prefixing unused variables with an underscore to make it clear what is meant to be used in a view:

```plain text
def calculate @_result_of_expensive_calculation ||= SuperCalculator.get_started(thing) end
```

## Great success

We recently concluded our experiment as a success, published results in our research newsletter, and have incorporated the rules into our best practices guide.
