---
title: "Bookmarkable by Design: URL-Driven State in HTMX | Loren Stewart"
notion_id: 2b754f1c-7d23-818c-9d6d-ca36c7370e6f
notion_url: https://app.notion.com/p/Bookmarkable-by-Design-URL-Driven-State-in-HTMX-Loren-Stewart-2b754f1c7d23818c9d6dca36c7370e6f
last_edited: 2025-11-26T18:54:00.000Z
source_url: https://www.lorenstew.art/blog/bookmarkable-by-design-url-state-htmx
tags: ["Article", "Tutorial", "Loren Stewart", "English", "Web Development", "Frontend", "HTMX"]
---
When you move from React to HTMX, you trade complex state management for server-side simplicity. But you still need to handle filters, sorting, pagination, and search. Where does that state live now?

The answer is surprisingly elegant: in the URL itself. By treating URL parameters as your single source of truth, you get bookmarkable, shareable application state without needing to install another dependency.

## The Pattern in Action

A URL like `/?status=active&sortField=price&sortDir=desc&page=2` tells you everything about the current view. It’s not just an address: it’s a complete state representation that users can bookmark, share, or refresh without losing context.

## Quick Start: Three Essential Steps

The entire pattern revolves around three synchronized steps:

1. **Server reads URL parameters** and renders the appropriate view
2. **Forms and hx-include preserve all state** when making HTMX requests
3. **Browser URL updates automatically** with hx-push-url

Let’s build this step by step.

## Step 1: Server Reads URL State

Your server endpoint reads query parameters and uses them to render the initial view:

```plain text
@Get("/")
@Render("data-table.eta")
async homepage(
  @Query("sortField") sortField: string,
  @Query("sortDir") sortDir: "asc" | "desc",
  @Query("status") status: string,
  @Query("page") page: string,
) {
  // Parse with defaults
  const pageNum = parseInt(page) || 1;

  // Apply state to data query
  const result = await this.dataService.getItems({
    sortField: sortField || "name",
    sortDir: sortDir || "asc",
    status: status || "",
    page: pageNum,
  });

  // Return both data and state for template
  return {
    items: result.items,
    totalItems: result.total,
    sortField: sortField || "name",
    sortDir: sortDir || "asc",
    status: status || "",
    page: pageNum,
  };
}
```

The template embeds this state directly in the DOM (using [ETA templates](https://eta.js.org/) in this case):

```plain text
<div id="data-table">
  <!-- Filter form with all state -->
  <form hx-get="/api/data"
        hx-target="#data-table"
        hx-push-url="true"
        hx-params="*"
        hx-trigger="change">

    <!-- Status filter -->
    <select name="status">
      <option value="" <%= it.status === '' ? 'selected' : '' %>>All</option>
      <option value="active" <%= it.status === 'active' ? 'selected' : '' %>>Active</option>
    </select>

    <!-- Hidden state fields -->
    <input type="hidden" name="sortField" value="<%= it.sortField %>">
    <input type="hidden" name="sortDir" value="<%= it.sortDir %>">
    <input type="hidden" name="page" value="<%= it.page %>">
  </form>

  <!-- Sortable columns -->
  <th class="sortable <%= it.sortField === 'name' ? 'sorted' : '' %>"
      hx-get="/api/data"
      hx-target="#data-table"
      hx-push-url="true"
      hx-params="*"
      hx-include="[name='status']"
      hx-vals="{sortField: 'name', sortDir: '<%= it.sortField === 'name' && it.sortDir === 'asc' ? 'desc' : 'asc' %>', page: 1}">
    Name
    <% if (it.sortField === 'name') { %>
      <%= it.sortDir === 'asc' ? '↑' : '↓' %>
    <% } %>
  </th>

  <th class="sortable <%= it.sortField === 'price' ? 'sorted' : '' %>"
      hx-get="/api/data"
      hx-target="#data-table"
      hx-push-url="true"
      hx-params="*"
      hx-include="[name='status']"
      hx-vals="{sortField: 'price', sortDir: '<%= it.sortField === 'price' && it.sortDir === 'asc' ? 'desc' : 'asc' %>', page: 1}">
    Price
    <% if (it.sortField === 'price') { %>
      <%= it.sortDir === 'asc' ? '↑' : '↓' %>
    <% } %>
  </th>
</div>
```

Key insights:

- State flows from URL → Server → DOM
- Hidden form fields preserve state across different interactions
- CSS classes reflect current state (e.g., `sorted` class on the active sort column)
- Server-side rendering means the page works before any JavaScript loads

## Step 2: State Coordination with Forms and Hidden Fields

Instead of complex JavaScript, we use HTML forms to coordinate state:

```plain text
<!-- Form captures all filter state -->
<form hx-get="/api/data"
      hx-target="#data-table"
      hx-push-url="true"
      hx-params="*"
      hx-trigger="change">

  <select name="status">
    <option value="">All</option>
    <option value="active">Active</option>
  </select>

  <!-- Preserve other state as hidden fields -->
  <input type="hidden" name="sortField" value="<%= it.sortField %>">
  <input type="hidden" name="sortDir" value="<%= it.sortDir %>">
  <input type="hidden" name="page" value="<%= it.page %>">
</form>

<!-- Each sortable column uses templating for dynamic values -->
<th class="sortable <%= it.sortField === 'date' ? 'sorted' : '' %>"
    hx-get="/api/data"
    hx-target="#data-table"
    hx-push-url="true"
    hx-params="*"
    hx-include="[name='status']"
    hx-vals="{sortField: 'date', sortDir: '<%= it.sortField === 'date' && it.sortDir === 'asc' ? 'desc' : 'asc' %>', page: 1}">
  Date
  <% if (it.sortField === 'date') { %>
    <%= it.sortDir === 'asc' ? '↑' : '↓' %>
  <% } %>
</th>
```

Key benefits of this approach:

- HTMX automatically preserves all form state in the URL
- No JavaScript needed for state coordination
- Sort direction toggling handled in the template logic using conditional expressions
- Filters are preserved when sorting, sorting resets pagination
- Each sortable column uses templating to dynamically set sortField and toggle sortDir

## Step 3: Automatic URL Syncing with hx-push-url

With `hx-push-url="true"` and `hx-params="*"`, HTMX automatically handles URL updates:

```plain text
<!-- All form data becomes URL parameters -->
<form hx-get="/api/data"
      hx-target="#data-table"
      hx-push-url="true"
      hx-params="*">
```

HTMX automatically:

1. Sends all form data as query parameters to the server
2. Updates the browser URL with those same parameters
3. Creates proper history entries for back/forward navigation

No JavaScript required for URL management because HTMX handles everything declaratively.

## Production Considerations

**URL Length Limits**: Browsers support URLs up to ~2000 characters. For complex filters, consider using abbreviated parameter names or moving some state server-side.

**Parameter Validation**: Always validate and sanitize URL parameters on the server. Treat them as untrusted user input.

**Testing**: The pattern is highly testable since state is explicit in the URL and form values; no complex JavaScript functions to mock.

## The Architecture Payoff

This URL-first approach delivers multiple benefits without the complexity of client-side state management libraries. Every view is inherently shareable, so you can send a colleague a link and they see will exactly what you see. The browser’s back button works as expected by returning to previous filter states. SEO is built in since search engines can crawl every state combination. And the debugging experience is transparent because the current state is always visible in the address bar.

By embracing the URL as your state store, you’re not working around the web platform; you’re working with it. This pattern scales from simple sorting to complex multi-filter interfaces while maintaining the simplicity that drew you to HTMX in the first place.

The next time you reach for a state management library, consider whether the humble URL might be all you need. In many cases, it’s not just sufficient; it’s superior.
