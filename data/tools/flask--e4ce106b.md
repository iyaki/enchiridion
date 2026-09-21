---
title: "Flask"
notion_id: e4ce106b-cf3d-4a2e-af30-c8c2a4fd5f9c
notion_url: https://app.notion.com/p/Flask-e4ce106bcf3d4a2eaf30c8c2a4fd5f9c
last_edited: 2022-12-19T19:14:00.000Z
source_url: https://palletsprojects.com/p/flask/
tags: ["Web Development", "Python", "Untried", "Framework/Library", "English"]
---
Flask is a lightweight WSGI web application framework. It is designed to make getting started quick and easy, with the ability to scale up to complex applications. It began as a simple wrapper around Werkzeug and Jinja and has become one of the most popular Python web application frameworks.

Flask offers suggestions, but doesn't enforce any dependencies or project layout. It is up to the developer to choose the tools and libraries they want to use. There are many extensions provided by the community that make adding new functionality easy.

```plain text
# app.py from flask import Flask app = Flask(__name__) @app.route("/") def greet(): return "Hello, World!"
```

```plain text
$ flask run * Running on http://127.0.0.1:5000/ (Press CTRL+C to quit)
```
