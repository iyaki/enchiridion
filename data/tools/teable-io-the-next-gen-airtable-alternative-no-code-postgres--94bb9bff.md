---
title: "teable.io - The Next Gen Airtable Alternative: No-Code Postgres"
notion_id: 94bb9bff-3684-402d-a1f3-ad0f4d21a7ed
notion_url: https://app.notion.com/p/teable-io-The-Next-Gen-Airtable-Alternative-No-Code-Postgres-94bb9bff3684402da1f3ad0f4d21a7ed
last_edited: 2024-05-10T20:41:00.000Z
source_url: https://github.com/teableio/teable
tags: ["Office", "Project Management", "Tool", "Service", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

#    

![image](https://github.com/teableio/teable/raw/develop/static/assets/images/teable-vertical-light.png)

### **Postgres-Airtable Fusion**

Teable is a Super fast, Real-time, Professional, Developer friendly, No-code database built on Postgres. It uses a simple, spreadsheet-like interface to create complex enterprise-level database applications. Unlock efficient app development with no-code, free from the hurdles of data security and scalability.

[Home](https://teable.io/) | [Help](https://help.teable.io/) | [Blog](https://blog.teable.io/) | [Template](https://template.teable.io/) | [Roadmap](https://app.teable.io/share/shr04TEw1u9EOQojPmG/view) | [Discord](https://discord.gg/uZwp7tDE5W) | [Twitter](https://twitter.com/teableio)

#    

![image](https://github.com/teableio/teable/raw/develop/static/assets/images/teable-interface-light.png)

## Quick Guide

1. Looking for a quick experience? Select a scenario from the [template center](https://template.teable.io/) and click "Use this template".
2. Seeking high performance? Try the [1 million rows demo](https://app.teable.io/share/shrVgdLiOvNQABtW0yX/view) to feel the speed of Teable.
3. Want to learn to use it quickly? Click on this [tutorial](https://help.teable.io/quick-start/build-a-simple-base)
4. Interested in deploying it yourself? Click [Deploy on Railway](https://railway.app/template/wada5e?referralCode=rE4BjB)

## ✨Features

### 📊 Spreadsheet-like Interface

All you want is here

- Cell Editing: Directly click and edit content within cells.
- Formula Support: Input mathematical and logical formulas to auto-calculate values.
- Data Sorting and Filtering: Sort data based on a column or multiple columns; use filters to view specific rows of data.
- Aggregation Function: Automatically summarize statistics for each column, providing instant calculations like sum, average, count, max, and min for streamlined data analysis.
- Data Formatting: formatting numbers, dates, etc.
- Grouping: Organize rows into collapsible groups based on column values for easier data analysis and navigation.
- Freeze Columns: Freeze the left column of the table so they remain visible while scrolling.
- Import/Export Capabilities: Import and export data from other formats, e.g., .csv, .xlsx.
- Row Styling & Conditional Formatting: Change row styles automatically based on specific conditions. (coming soon)
- Charts & Visualization Tools: Create charts from table data such as bar charts, pie charts, line graphs, etc. (coming soon)
- Data Validation: Limit or validate data that are entered into cells. (coming soon)
- Undo/Redo: Undo or redo recent changes. (coming soon)
- Comments & Annotations: Attach comments to rows, providing explanations or feedback for other users. (coming soon)

### 🗂️ Multiple Views

Visualize and interact with data in various ways best suited for their specific tasks.

- Grid View: The default view of the table, which displays data in a spreadsheet-like format.
- Form View: Input data in a form format, which is useful for collecting data.
- Kanban View: Displays data in a Kanban board, which is a visual representation of data in columns and cards.
- Calendar View: Displays data in a calendar format, which is useful for tracking dates and events. (coming soon)
- Gallery View: Displays data in a gallery format, which is useful for displaying images and other media. (coming soon)
- Gantt View: Displays data in a Gantt chart, which is useful for tracking project schedules. (coming soon)
- Timeline View: Displays data in a timeline format, which is useful for tracking events over time. (coming soon)

### 🚀 Super Fast

Amazing response speed and data capacity

- Millions of data are easily processed, and there is no pressure to filter and sort
- Automatic database indexing for maximum speed
- Supports batch data operations at one time

### 👨‍💻 Full-featured SQL Support

Seamless integration with the software you are familiar with

- BI tools like Metabase PowerBi...
- No-code tools like Appsmith...
- Direct retrieve data with native SQL

### 🔒 Privacy-First

You own your data, in spite of the cloud

- Bring your own database (coming soon)

### ⚡️ Real-time Collaboration

Designed for teams

- No need to refresh the page, data is updated in real-time
- Seamlessly integrate collaboration member invitation and management
- Perfect permission management mechanism, from table to column level

### 🧩 Extensions (Coming Soon)

Expand infinite possibilities

- Backend-less programming capability based on React
- Customize your own application with extremely low cost
- Extremely easy-to-use script extensions mode

### 🤖 Automation (Coming Soon)

Empower data-driven workflows effortlessly and seamlessly

- Design your workflow with AI or Visual programming
- Super easy to retrieve data from the table

### 🧠 Copilot (Coming Soon)

Native Integrated AI ability

- Chat 2 App. "Create a project management app for me"
- Chat 2 Chart. "Analyze the data in the order table using a bar chart"
- Chat 2 View. "I want to see the schedule for the past week and only display participants"
- Chat 2 Action. "After the order is paid and completed, an email notification will be sent to the customer"
- More actions...

### 🗄️ Support for Multiple Databases (Coming Soon)

Choose the SQL database you like

- Sqlite, PostgreSQL, MySQL, MariaDB, TiDB...

# Structure

![image](https://camo.githubusercontent.com/1470bc5ed68e2671a712665be23b2c0612229ac893c1eee461399476f43bc28e/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4f70656e253230496e2d476974706f642e696f2d2532333139363644323f7374796c653d666f722d7468652d6261646765266c6f676f3d676974706f64)

```plain text
.
├── apps
│   ├── nextjs-app          (front-end, include a nextjs app)
│   └── nestjs-backend      (backend, include a nestjs app)
└── packages
    ├── common-i18n         (locales)
    ├── core                (share code and interface)
    ├── sdk                 (sdk for extensions)
    ├── db-main-prisma      (schema, migrations, prisma client)
    ├── eslint-config-bases (to shared eslint configs)
    └── ui-lib              (ui component)

```

## Deploy

### Deploy With Docker

```plain text
cd dockers/examples/standalone/
docker-compose up -d
```

for more details, see [dockers/examples](https://github.com/teableio/teable/blob/develop/dockers/examples)

### One Click Deployment

These platforms are easy to deploy with one click and come with free credits.

![image](https://camo.githubusercontent.com/d07713342bc583232f8752c33a6a24e5f367d73725183a63f2f5fdd7c00606a3/68747470733a2f2f7261696c7761792e6170702f627574746f6e2e737667)

![image](https://camo.githubusercontent.com/7bbd819d76b4e67c04fe4d8a0d9cc7220e691f3d2177f6e54f84689ba380184e/68747470733a2f2f7a65616275722e636f6d2f627574746f6e2e737667)

![image](https://raw.githubusercontent.com/labring-actions/templates/main/Deploy-on-Sealos.svg)

## Development

### 1. Initialize

```plain text
# Enabling the Help Management Package Manager
corepack enable

# Install project dependencies
pnpm install
```

### 2. Select Database

we currently support `sqlite` and `postgres`, you can switch between them by running the following command

```plain text
make switch-db-mode
```

### 3. Custom Environment Variables（Optional）

```plain text
cd apps/nextjs-app
copy .env.development .env.development.local
```

### 4. Run Dev Server

you just need to start backend, it will start next server for frontend automatically, file change will be auto reload

```plain text
cd apps/nestjs-backend
pnpm dev
```

## Why Teable?

No-code tools have significantly speed up how we get things done, allowing non-tech users to build amazing apps and changing the way many work and live. People like using spreadsheet-like UI to handle their data because it's easy, flexible, and great for team collaboration. They also prefer designing their app screens without being stuck with clunky templates.

Giving non-techy people the ability to create their software sounds exciting. But that's just the start:

- As businesses expand, their data needs intensify. No one wishes to hear that once their orders reach 100k, they'll outgrow their current interface. Yet, many no-code platforms falter at such scales.
- Most no-code platforms are cloud-based. This means your important data sits with the provider, and switching to another platform can be a headache.
- Sometimes, no-code tools can't do what you want because of their limitations, leaving users stuck.
- If a tool becomes essential, you'll eventually need some tech expertise. But developers often find these platforms tricky.
- Maintaining systems with complex setups can be hard for developers, especially if these aren't built using common software standards.
- Systems that don't use these standards might need revamping or replacing, costing more in the long run. It might even mean ditching the no-code route and going back to traditional coding.

### What We Think the Future Of No-code Products Look Like

- An interface that anyone can use to build applications easily.
- Easy access to data, letting users grab, move, and reuse their information as they wish.
- Data privacy and choice, whether that's in the cloud, on-premise, or even just on your local.
- It needs to work for developers too, not just non-tech users.
- It should handle lots of data, so it can grow with your business.
- Flexibility to integrate with other software, combining strengths to get the job done.
- Last, native AI integration to takes usability to the next level.

In essence, Teable isn't just another no-code solution, it's a comprehensive answer to the evolving demands of modern software development, ensuring that everyone, regardless of their technical proficiency, has a platform tailored to their needs.

## Sponsors ❤️

If you are enjoying some this project in your company, I'd really appreciate a [sponsorship](https://github.com/sponsors/teableio), a [coffee](https://ko-fi.com/teable) or a dropped star. That gives me some more time to improve it to the next level.

# License

AGPL-3.0
