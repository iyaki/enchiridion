---
title: "basaran: Basaran is an open-source alternative to the OpenAI text completion API. It provides a compatible streaming API for your Hugging Face Transformers-based text generation models"
notion_id: f7b44fb5-5a80-4884-9d15-b7bbcc530d00
notion_url: https://app.notion.com/p/basaran-Basaran-is-an-open-source-alternative-to-the-OpenAI-text-completion-API-It-provides-a-comp-f7b44fb55a8048849d15b7bbcc530d00
last_edited: 2023-04-22T01:16:00.000Z
source_url: https://github.com/hyperonym/basaran
tags: ["English", "Artificial Intelligence (AI)", "Untried", "Tool"]
---
# Basaran

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Basaran is an open-source alternative to the [OpenAI text completion API](https://platform.openai.com/docs/api-reference/completions/create). It provides a compatible streaming API for your [Hugging Face Transformers](https://huggingface.co/docs/transformers/index)-based [text generation models](https://huggingface.co/models?pipeline_tag=text-generation).

The open source community will eventually witness the [Stable Diffusion](https://stability.ai/blog/stable-diffusion-public-release) moment for large language models (LLMs), and Basaran allows you to replace OpenAI's service with the latest open-source model to power your application [without modifying a single line of code](https://github.com/hyperonym/basaran/blob/master/README.md#openai-client-library).

The key features of Basaran are:

- Streaming generation using various decoding strategies.
- Support for both decoder-only and encoder-decoder models.
- Detokenizer that handles surrogates and whitespace.
- Multi-GPU support with optional quantization.
- Real-time partial progress using [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format).
- Compatibility with OpenAI API and client libraries.
- Comes with a fancy web-based playground!

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Quick Start

### TL;DR

Replace `user/repo` with your [selected model](https://huggingface.co/models?pipeline_tag=text-generation) and `X.Y.Z` with the [latest version](https://hub.docker.com/r/hyperonym/basaran/tags), then run:

```plain text
docker run -p 80:80 -e MODEL=user/repo hyperonym/basaran:X.Y.Z
```

And you're good to go!

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

```plain text
Playground: http://127.0.0.1/
API: http://127.0.0.1/v1/completions

```

### Installation

### Using Docker (Recommended)

Docker images are available on [Docker Hub](https://hub.docker.com/r/hyperonym/basaran/tags) and [GitHub Packages](https://github.com/orgs/hyperonym/packages?repo_name=basaran).

For GPU acceleration, you also need to install the [NVIDIA Driver](https://docs.nvidia.com/datacenter/tesla/tesla-installation-notes/index.html) and [NVIDIA Container Runtime](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/install-guide.html). Basaran's image already comes with related libraries such as CUDA and cuDNN, so there is no need to install them manually.

Basaran's image can be used in three ways:

- **Run directly**: By specifying the `MODEL="user/repo"` environment variable, the corresponding model can be downloaded from Hugging Face Hub during the first startup.
- **Bundling**: Create a new Dockerfile to [preload a public model](https://github.com/hyperonym/basaran/blob/master/deployments/bundle/bloomz-560m.Dockerfile) or [bundle a private model](https://github.com/hyperonym/basaran/blob/master/deployments/bundle/private.Dockerfile).
- **Bind mount**: Mount a model from the local file system into the container and point the `MODEL` environment variable to the corresponding path.

For the above use cases, you can find sample [Dockerfiles](https://github.com/hyperonym/basaran/tree/master/deployments/bundle) and [docker-compose files](https://github.com/hyperonym/basaran/tree/master/deployments/compose) in the [deployments directory](https://github.com/hyperonym/basaran/tree/master/deployments).

### Using pip

Basaran is tested on Python 3.8+ and PyTorch 1.13+. You should create a [virtual environment](https://docs.python.org/3/library/venv.html) with the version of Python you want to use, and activate it before proceeding.

1. Install with `pip`:

```plain text
pip install basaran
```

1. Install dependencies required for GPU acceleration (optional):

```plain text
pip install accelerate bitsandbytes
```

1. Replace `user/repo` with the selected model and run Basaran:

```plain text
MODEL=user/repo PORT=80 python -m basaran
```

For a complete list of environment variables, see [`__init__.py`](https://github.com/hyperonym/basaran/blob/master/basaran/__init__.py).

### Running From Source

If you want to access the latest features or hack it yourself, you can choose to run from source using `git`.

1. Clone the repository:

```plain text
git clone https://github.com/hyperonym/basaran.git && cd basaran
```

1. Install dependencies:

```plain text
pip install -r requirements.txt
```

1. Replace `user/repo` with the selected model and run Basaran:

```plain text
MODEL=user/repo PORT=80 python -m basaran
```

### Basic Usage

### cURL

Basaran's HTTP request and response formats are consistent with the [OpenAI API](https://platform.openai.com/docs/api-reference).

Taking [text completion](https://platform.openai.com/docs/api-reference/completions/create) as an example:

```plain text
curl http://127.0.0.1/v1/completions \
 -H 'Content-Type: application/json' \
 -d '{ "prompt": "once upon a time,", "echo": true }'
```

Details

### OpenAI Client Library

If your application uses [client libraries](https://github.com/openai/openai-python) provided by OpenAI, you only need to modify the `OPENAI_API_BASE` environment variable to match Basaran's endpoint:

```plain text
OPENAI_API_BASE="http://127.0.0.1/v1" python your_app.py
```

The [examples](https://github.com/hyperonym/basaran/tree/master/examples) directory contains examples of [using the OpenAI Python library](https://github.com/hyperonym/basaran/blob/master/examples/openai-python-library/main.py).

### Using as a Python Library

Basaran is also available as a library on [PyPI](https://pypi.org/project/basaran/). It can be used directly in Python without the need to start a separate API server.

1. Install with `pip`:

```plain text
pip install basaran
```

1. Use the `load_model` function to load a model:

```plain text
from basaran.model import load_model

model = load_model("user/repo")
```

1. Generate streaming output by calling the model:

```plain text
for choice in model("once upon a time"):
 print(choice)
```

The [examples](https://github.com/hyperonym/basaran/tree/master/examples) directory contains examples of [using Basaran as a library](https://github.com/hyperonym/basaran/blob/master/examples/basaran-python-library/main.py).

## Compatibility

Basaran's API format is consistent with OpenAI's, with differences in compatibility mainly in terms of parameter support and response fields. The following sections provide detailed information on the compatibility of each endpoint.

### Models

Each Basaran process serves only one model, so the result will only contain that model.

### Completions

Although Basaran does not support the `model` parameter, the OpenAI client library requires it to be present. Therefore, you can enter any random model name.

| Parameter | Basaran | OpenAI | Default Value | Maximum Value |
| --- | --- | --- | --- | --- |
| `model` | ○ | ● | - | - |
| `prompt` | ● | ● | `""` | `COMPLETION_MAX_PROMPT` |
| `suffix` | ○ | ● | - | - |
| `min_tokens` | ● | ○ | `0` | `COMPLETION_MAX_TOKENS` |
| `max_tokens` | ● | ● | `16` | `COMPLETION_MAX_TOKENS` |
| `temperature` | ● | ● | `1.0` | - |
| `top_p` | ● | ● | `1.0` | - |
| `n` | ● | ● | `1` | `COMPLETION_MAX_N` |
| `stream` | ● | ● | `false` | - |
| `logprobs` | ● | ● | `0` | `COMPLETION_MAX_LOGPROBS` |
| `echo` | ● | ● | `false` | - |
| `stop` | ○ | ● | - | - |
| `presence_penalty` | ○ | ● | - | - |
| `frequency_penalty` | ○ | ● | - | - |
| `best_of` | ○ | ● | - | - |
| `logit_bias` | ○ | ● | - | - |
| `user` | ○ | ● | - | - |

### Chat

Providing a unified chat API is currently difficult because each model has a different format for chat history.

Therefore, it is recommended to pre-format the chat history based on the requirements of the specific model and use it as the prompt for the completion API.

### [GPT-NeoXT-Chat-Base-20B](https://huggingface.co/togethercomputer/GPT-NeoXT-Chat-Base-20B)

```plain text
**Summarize a long document into a single sentence and ...**

<human>: Last year, the travel industry saw a big ...

<bot>: If you're traveling this spring break, ...

<human>: But ...

<bot>:

```

### [chatglm-6b](https://huggingface.co/THUDM/chatglm-6b)

```plain text
[Round 0]
问：你好
答：你好!有什么我可以帮助你的吗?
[Round 1]
问：你是谁？
答：

```

## Roadmap

- API 
- Models 
- List models
- Retrieve model
- Completions 
- Create completion
- Chat 
- Create chat completion
- Model 
- Architectures 
- Encoder-decoder
- Decoder-only
- Decoding strategies 
- Random sampling with temperature
- Nucleus-sampling (top-p)
- Stop sequences
- Presence and frequency penalties

See the [open issues](https://github.com/hyperonym/basaran/issues) for a full list of proposed features.

## Contributing

This project is open-source. If you have any ideas or questions, please feel free to reach out by creating an issue!

Contributions are greatly appreciated, please refer to [CONTRIBUTING.md](https://github.com/hyperonym/basaran/blob/master/CONTRIBUTING.md) for more information.

## License

Basaran is available under the [MIT License](https://github.com/hyperonym/basaran/blob/master/LICENSE).

© 2023 [Hyperonym](https://hyperonym.org/)
