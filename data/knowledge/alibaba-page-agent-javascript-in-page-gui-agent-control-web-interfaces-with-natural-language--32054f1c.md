---
title: "alibaba/page-agent: JavaScript in-page GUI agent. Control web interfaces with natural language."
notion_id: 32054f1c-7d23-81ef-aed0-f90fe6d3d629
notion_url: https://app.notion.com/p/alibaba-page-agent-JavaScript-in-page-GUI-agent-Control-web-interfaces-with-natural-language-32054f1c7d2381efaed0f90fe6d3d629
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://github.com/alibaba/page-agent
tags: ["Tool", "Article", "GitHub", "English", "Web Development", "Javascript", "AI", "SaaS", "Frontend", "Productivity"]
---
# Page Agent

![image](https://camo.githubusercontent.com/79058d74ca2396ef43f2bdb2fd2c90024a3dc22ed53c202c975e84d63fb6ce8b/68747470733a2f2f696d672e616c6963646e2e636f6d2f696d6765787472612f69312f4f31434e30314e434d4b586a31476e34746b46547378665f2121363030303030303030303636362d322d7470732d313238302d3235362e706e67)

![image](https://camo.githubusercontent.com/fdf2982b9f5d7489dcf44570e714e3a15fce6253e0cc6b5aa61a075aac2ff71b/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4c6963656e73652d4d49542d79656c6c6f772e737667)

![image](https://camo.githubusercontent.com/a6ed1655e865a26dfcec492b68d0fa0ed6d8b7c1aaf3b974bfbe28f2b71e6232/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2533432532462533452d547970655363726970742d2532333030373463312e737667)

![image](https://camo.githubusercontent.com/b6b7528b856d0deefabe139e7038a120b1b84f6664cd4d95ec35aac0c64b1907/68747470733a2f2f696d672e736869656c64732e696f2f6e706d2f64742f706167652d6167656e742e737667)

![image](https://camo.githubusercontent.com/89a58997d17d9b044a28417f593de585b640f055545ec00673a8503f7c3abdba/68747470733a2f2f696d672e736869656c64732e696f2f62756e646c6570686f6269612f6d696e7a69702f706167652d6167656e74)

![image](https://camo.githubusercontent.com/b94b57233c1c4f16db54d8c0930c36700eb9c72689bbb5d69b4ab3b85977183a/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f73746172732f616c69626162612f706167652d6167656e742e737667)

The GUI Agent Living in Your Webpage. Control web interfaces with natural language.

🌐 **English** | [中文](https://github.com/alibaba/page-agent/blob/main/docs/README-zh.md)

👉 [**🚀 Demo**](https://alibaba.github.io/page-agent/) | [**📖 Documentation**](https://alibaba.github.io/page-agent/docs/introduction/overview) | [📢 Join HN Discussion](https://news.ycombinator.com/item?id=47264138)

**page-agent-demo-0227.mp4**

## ✨ Features

- **🎯 Easy integration** 
- **📖 Text-based DOM manipulation** 
- **🧠 Bring your own LLMs**
- **🎨 Pretty UI with human-in-the-loop**
- **🐙 Optional **[**chrome extension**](https://alibaba.github.io/page-agent/docs/features/chrome-extension)** for multi-page tasks.**

## 💡 Use Cases

- **SaaS AI Copilot** — Ship an AI copilot in your product in lines of code. No backend rewrite needed.
- **Smart Form Filling** — Turn 20-click workflows into one sentence. Perfect for ERP, CRM, and admin systems.
- **Accessibility** — Make any web app accessible through natural language. Voice commands, screen readers, zero barrier.
- **Multi-page Agent** — Extend your agent's reach across browser tabs with the optional [chrome extension](https://alibaba.github.io/page-agent/docs/features/chrome-extension).

## 🚀 Quick Start

### One-line integration

Fastest way to try PageAgent with our free Demo LLM:

```plain text
<script src="{URL}" crossorigin="true"></script>
```

| Mirrors | URL |
| --- | --- |
| Global | [https://cdn.jsdelivr.net/npm/page-agent@1.5.5/dist/iife/page-agent.demo.js](https://cdn.jsdelivr.net/npm/page-agent@1.5.5/dist/iife/page-agent.demo.js) |
| China | [https://registry.npmmirror.com/page-agent/1.5.5/files/dist/iife/page-agent.demo.js](https://registry.npmmirror.com/page-agent/1.5.5/files/dist/iife/page-agent.demo.js) |

> 

### NPM Installation

```plain text
npm install page-agent
```

```plain text
import { PageAgent } from 'page-agent'

const agent = new PageAgent({
    model: 'qwen3.5-plus',
    baseURL: 'https://dashscope.aliyuncs.com/compatible-mode/v1',
    apiKey: 'YOUR_API_KEY',
    language: 'en-US',
})

await agent.execute('Click the login button')
```

For more programmatic usage, see [📖 Documentations](https://alibaba.github.io/page-agent/docs/introduction/overview).

## 🤝 Contributing

We welcome contributions from the community! Follow our instructions in [CONTRIBUTING.md](https://github.com/alibaba/page-agent/blob/main/CONTRIBUTING.md) for environment setup and local development.

Please read [Code of Conduct](https://github.com/alibaba/page-agent/blob/main/docs/CODE_OF_CONDUCT.md) before contributing.

## 👏 Acknowledgments

This project builds upon the excellent work of [**`browser-use`**](https://github.com/browser-use/browser-use).

`PageAgent` is designed for **client-side web enhancement**, not server-side automation.

```plain text
DOM processing components and prompt are derived from browser-use:

Browser Use
Copyright (c) 2024 Gregor Zunic
Licensed under the MIT License

Original browser-use project: <https://github.com/browser-use/browser-use>

We gratefully acknowledge the browser-use project and its contributors for their
excellent work on web automation and DOM interaction patterns that helped make
this project possible.

Third-party dependencies and their licenses can be found in the package.json
file and in the node_modules directory after installation.
```

## 📄 License

[MIT License](https://github.com/alibaba/page-agent/blob/main/LICENSE)

**⭐ Star this repo if you find PageAgent helpful!**

![image](https://camo.githubusercontent.com/a925269ab55e399b7b9ff5d640ee2958c70cba31982cab5ff2ca077eb25c5dd4/68747470733a2f2f6170692e737461722d686973746f72792e636f6d2f696d6167653f7265706f733d616c69626162612f706167652d6167656e7426747970653d64617465266c6567656e643d746f702d6c65667426763d32)
