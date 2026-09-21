---
title: "GPT best practices - OpenAI"
notion_id: 266188cb-9026-41b9-858f-54542d668fa1
notion_url: https://app.notion.com/p/GPT-best-practices-OpenAI-266188cb902641b9858f54542d668fa1
last_edited: 2026-09-21T17:00:00.000Z
source_url: https://developers.openai.com/api/docs/guides/prompt-engineering
tags: ["English", "Artificial Intelligence (AI)", "Guide"]
---
This guide shares strategies and tactics for getting better results from GPTs. The methods described here can sometimes be deployed in combination for greater effect. We encourage experimentation to find the methods that work best for you.

Some of the examples demonstrated here currently work only with our most capable model, `gpt-4`. If you don't yet have access to `gpt-4` consider joining the [waitlist](https://openai.com/waitlist/gpt-4-api). In general, if you find that a GPT model fails at a task and a more capable model is available, it's often worth trying again with the more capable model.

## [Six strategies for getting better results](https://platform.openai.com/docs/guides/gpt-best-practices/six-strategies-for-getting-better-results)

### [Write clear instructions](https://platform.openai.com/docs/guides/gpt-best-practices/write-clear-instructions)

GPTs can’t read your mind. If their outputs are too long, ask for brief replies. If their outputs are too simple, ask for expert-level writing. If you dislike the format, demonstrate the format you’d like to see. The less GPTs have to guess at what you want, the more likely you’ll be to get it.

Tactics:

- [Include details in your query to get more relevant answers](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-include-details-in-your-query-to-get-more-relevant-answers)
- [Ask the model to adopt a persona](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-ask-the-model-to-adopt-a-persona)
- [Use delimiters to clearly indicate distinct parts of the input](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-use-delimiters-to-clearly-indicate-distinct-parts-of-the-input)
- [Specify the steps required to complete a task](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-specify-the-steps-required-to-complete-a-task)
- [Provide examples](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-provide-examples)
- [Specify the desired length of the output](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-specify-the-desired-length-of-the-output)

### [Provide reference text](https://platform.openai.com/docs/guides/gpt-best-practices/provide-reference-text)

GPTs can confidently invent fake answers, especially when asked about esoteric topics or for citations and URLs. In the same way that a sheet of notes can help a student do better on a test, providing reference text to GPTs can help in answering with fewer fabrications.

Tactics:

- [Instruct the model to answer using a reference text](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-instruct-the-model-to-answer-using-a-reference-text)
- [Instruct the model to answer with citations from a reference text](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-instruct-the-model-to-answer-with-citations-from-a-reference-text)

Just as it is good practice in software engineering to decompose a complex system into a set of modular components, the same is true of tasks submitted to GPTs. Complex tasks tend to have higher error rates than simpler tasks. Furthermore, complex tasks can often be re-defined as a workflow of simpler tasks in which the outputs of earlier tasks are used to construct the inputs to later tasks.

Tactics:

- [Use intent classification to identify the most relevant instructions for a user query](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-use-intent-classification-to-identify-the-most-relevant-instructions-for-a-user-query)
- [For dialogue applications that require very long conversations, summarize or filter previous dialogue](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-for-dialogue-applications-that-require-very-long-conversations-summarize-or-filter-previous-dialogue)
- [Summarize long documents piecewise and construct a full summary recursively](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-summarize-long-documents-piecewise-and-construct-a-full-summary-recursively)

If asked to multiply 17 by 28, you might not know it instantly, but can still work it out with time. Similarly, GPTs make more reasoning errors when trying to answer right away, rather than taking time to work out an answer. Asking for a chain of reasoning before an answer can help GPTs reason their way toward correct answers more reliably.

Tactics:

- [Instruct the model to work out its own solution before rushing to a conclusion](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-instruct-the-model-to-work-out-its-own-solution-before-rushing-to-a-conclusion)
- [Use inner monologue or a sequence of queries to hide the model's reasoning process](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-use-inner-monologue-or-a-sequence-of-queries-to-hide-the-model-s-reasoning-process)
- [Ask the model if it missed anything on previous passes](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-ask-the-model-if-it-missed-anything-on-previous-passes)

Compensate for the weaknesses of GPTs by feeding them the outputs of other tools. For example, a text retrieval system can tell GPTs about relevant documents. A code execution engine can help GPTs do math and run code. If a task can be done more reliably or efficiently by a tool rather than by a GPT, offload it to get the best of both.

Tactics:

- [Use embeddings-based search to implement efficient knowledge retrieval](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-use-embeddings-based-search-to-implement-efficient-knowledge-retrieval)
- [Use code execution to perform more accurate calculations or call external APIs](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-use-code-execution-to-perform-more-accurate-calculations-or-call-external-apis)

Improving performance is easier if you can measure it. In some cases a modification to a prompt will achieve better performance on a few isolated examples but lead to worse overall performance on a more representative set of examples. Therefore to be sure that a change is net positive to performance it may be necessary to define a comprehensive test suite (also known an as an "eval").

Tactic:

- [Evaluate model outputs with reference to gold-standard answers](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-evaluate-model-outputs-with-reference-to-gold-standard-answers)

Each of the strategies listed above can be instantiated with specific tactics. These tactics are meant to provide ideas for things to try. They are by no means fully comprehensive, and you should feel free to try creative ideas not represented here.

In order to get a highly relevant response, make sure that requests provide any important details or context. Otherwise you are leaving it up to the model to guess what you mean.

| **Worse** | **Better** |
| --- | --- |
| How do I add numbers in Excel? | How do I add up a row of dollar amounts in Excel? I want to do this automatically for a whole sheet of rows with all the totals ending up on the right in a column called "Total". |
| Who’s president? | Who was the president of Mexico in 2021, and how frequently are elections held? |
| Write code to calculate the Fibonacci sequence. | Write a TypeScript function to efficiently calculate the Fibonacci sequence. Comment the code liberally to explain what each piece does and why it's written that way. |
| Summarize the meeting notes. | Summarize the meeting notes in a single paragraph. Then write a markdown list of the speakers and each of their key points. Finally, list the next steps or action items suggested by the speakers, if any. |

The system message can be used to specify the persona used by the model in its replies.

Delimiters like triple quotation marks, XML tags, section titles, etc. can help demarcate sections of text to be treated differently.

For straightforward tasks such as these, using delimiters might not make a difference in the output quality. However, the more complex a task is the more important it is to disambiguate task details. Don’t make GPTs work to understand exactly what you are asking of them.

Some tasks are best specified as a sequence of steps. Writing the steps out explicitly can make it easier for the model to follow them.

Providing general instructions that apply to all examples is generally more efficient than demonstrating all permutations of a task by example, but in some cases providing examples may be easier. For example, if you intend for the model to copy a particular style of responding to user queries which is difficult to describe explicitly. This is known as "few-shot" prompting.

You can ask the model to produce outputs that are of a given target length. The targeted output length can be specified in terms of the count of words, sentences, paragraphs, bullet points, etc. Note however that instructing the model to generate a specific number of words does not work with high precision. The model can more reliably generate outputs with a specific number of paragraphs or bullet points.

If we can provide a model with trusted information that is relevant to the current query, then we can instruct the model to use the provided information to compose its answer.

Given that GPTs have limited context windows, in order to apply this tactic we need some way to dynamically lookup information that is relevant to the question being asked. [Embeddings](https://platform.openai.com/docs/guides/embeddings/what-are-embeddings) can be used to implement efficient knowledge retrieval. See the tactic ["Use embeddings-based search to implement efficient knowledge retrieval"](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-use-embeddings-based-search-to-implement-efficient-knowledge-retrieval) for more details on how to implement this.

If the input has been supplemented with relevant knowledge, it's straightforward to request that the model add citations to its answers by referencing passages from provided documents. Note that citations in the output can then be verified programmatically by string matching within the provided document.

For tasks in which lots of independent sets of instructions are needed to handle different cases, it can be beneficial to first classify the type of query and to use that classification to determine which instructions are needed. This can be achieved by defining fixed categories and hardcoding instructions that are relevant for handling tasks in a given category. This process can also be applied recursively to decompose a task into a sequence of stages. The advantage of this approach is that each query will contain only those instructions that are required to perform the next stage of a task which can result in lower error rates compared to using a single query to perform the whole task. This can also result in lower costs since larger prompts cost more to run ([see pricing information](https://openai.com/pricing)).

Suppose for example that for a customer service application, queries could be usefully classified as follows:

Based on the classification of the customer query, a set of more specific instructions can be provided to a GPT model to handle next steps. For example, suppose the customer requires help with "troubleshooting".

Notice that the model has been instructed to emit special strings to indicate when the state of the conversation changes. This enables us to turn our system into a state machine where the state determines which instructions are injected. By keeping track of state, what instructions are relevant at that state, and also optionally what state transitions are allowed from that state, we can put guardrails around the user experience that would be hard to achieve with a less structured approach.

Since GPTs have a fixed context length, dialogue between a user and an assistant in which the entire conversation is included in the context window cannot continue indefinitely.

There are various workarounds to this problem, one of which is to summarize previous turns in the conversation. Once the size of the input reaches a predetermined threshold length, this could trigger a query that summarizes part of the conversation and the summary of the prior conversation could be included as part of the system message. Alternatively, prior conversation could be summarized asynchronously in the background throughout the entire conversation.

An alternative solution is to dynamically select previous parts of the conversation that are most relevant to the current query. See the tactic ["Use embeddings-based search to implement efficient knowledge retrieval"](https://platform.openai.com/tactic-use-embeddings-based-search-to-implement-efficient-knowledge-retrieval).

Since GPTs have a fixed context length, they cannot be used to summarize a text longer than the context length minus the length of the generated summary in a single query.

To summarize a very long document such as a book we can use a sequence of queries to summarize each section of the document. Section summaries can be concatenated and summarized producing summaries of summaries. This process can proceed recursively until an entire document is summarized. If it’s necessary to use information about earlier sections in order to make sense of later sections, then a further trick that can be useful is to include a running summary of the text that precedes any given point in the book while summarizing content at that point. The effectiveness of this procedure for summarizing books has been studied in previous [research](https://openai.com/research/summarizing-books) by OpenAI using variants of GPT-3.

Sometimes we get better results when we explicitly instruct the model to reason from first principles before coming to a conclusion. Suppose for example we want a model to evaluate a student’s solution to a math problem. The most obvious way to approach this is to simply ask the model if the student's solution is correct or not.

But the student's solution is actually not correct! We can get the model to successfully notice this by prompting the model to generate its own solution first.

The previous tactic demonstrates that it is sometimes important for the model to reason in detail about a problem before answering a specific question. For some applications, the reasoning process that a model uses to arrive at a final answer would be inappropriate to share with the user. For example, in tutoring applications we may want to encourage students to work out their own answers, but a model’s reasoning process about the student’s solution could reveal the answer to the student.

Inner monologue is a tactic that can be used to mitigate this. The idea of inner monologue is to instruct the model to put parts of the output that are meant to be hidden from the user into a structured format that makes parsing them easy. Then before presenting the output to the user, the output is parsed and only part of the output is made visible.

Alternatively, this can be achieved with a sequence of queries in which all except the last have their output hidden from the end user.

First, we can ask the model to solve the problem on its own. Since this initial query doesn't require the student’s solution, it can be omitted. This provides the additional advantage that there is no chance that the model’s solution will be biased by the student’s attempted solution.

Next, we can have the model use all available information to assess the correctness of the student’s solution.

Finally, we can let the model use its own analysis to construct a reply in the persona of a helpful tutor.

Suppose that we are using a model to list excerpts from a source which are relevant to a particular question. After listing each excerpt the model needs to determine if it should start writing another or if it should stop. If the source document is large, it is common for a model to stop too early and fail to list all relevant excerpts. In that case, better performance can often be obtained by prompting the model with followup queries to find any excerpts it missed on previous passes.

A model can leverage external sources of information if provided as part of its input. This can help the model to generate more informed and up-to-date responses. For example, if a user asks a question about a specific movie, it may be useful to add high quality information about the movie (e.g. actors, director, etc…) to the model’s input. Embeddings can be used to implement efficient knowledge retrieval, so that relevant information can be added to the model input dynamically at run-time.

A text embedding is a vector that can measure the relatedness between text strings. Similar or relevant strings will be closer together than unrelated strings. This fact, along with the existence of fast vector search algorithms means that embeddings can be used to implement efficient knowledge retrieval. In particular, a text corpus can be split up into chunks, and each chunk can be embedded and stored. Then a given query can be embedded and vector search can be performed to find the embedded chunks of text from the corpus that are most related to the query (i.e. closest together in the embedding space).

Example implementations can be found in the [OpenAI Cookbook](https://github.com/openai/openai-cookbook/blob/main/examples/vector_databases/Using_vector_databases_for_embeddings_search.ipynb). See the tactic [“Instruct the model to use retrieved knowledge to answer queries”](https://platform.openai.com/docs/guides/gpt-best-practices/tactic-instruct-the-model-to-use-retrieved-knowledge-to-answer-queries) for an example of how to use knowledge retrieval to minimize the likelihood that a model will make up incorrect facts.

GPTs cannot be relied upon to perform arithmetic or long calculations accurately on their own. In cases where this is needed, a model can be instructed to write and run code instead of making its own calculations. In particular, a model can be instructed to put code that is meant to be run into a designated format such as triple backtics. After an output is produced, the code can be extracted and run. Finally, if necessary, the output from the code execution engine (i.e. Python interpreter) can be provided as an input to the model for the next query.

Another good use case for code execution is calling external APIs. If a model is instructed in the proper use of an API, it can write code that makes use of it. A model can be instructed in how to use an API by providing it with documentation and/or code samples showing how to use the API.

**WARNING: Executing code produced by a model is not inherently safe and precautions should be taken in any application that seeks to do this. In particular, a sandboxed code execution environment is needed to limit the harm that untrusted code could cause.**

Sometimes it can be hard to tell whether a change — e.g., a new instruction or a new design — makes your system better or worse. Looking at a few examples may hint at which is better, but with small sample sizes it can be hard to distinguish between a true improvement or random luck. Maybe the change helps performance on some inputs, but hurts performance on others.

Evaluation procedures (or "evals") are useful for optimizing system designs. Good evals are:

- Representative of real-world usage (or at least diverse)
- Contain many test cases for greater statistical power (see table below for guidelines)
- Easy to automate or repeat

| Difference to detect | Sample size needed for 95% confidence |
| --- | --- |
| 30% | ~10 |
| 10% | ~100 |
| 3% | ~1,000 |
| 1% | ~10,000 |

Evaluation of outputs can be done by computers, humans, or a mix. Computers can automate evals with objective criteria (e.g., questions with single correct answers) as well as some subjective or fuzzy criteria, in which model outputs are evaluated by other model queries. [OpenAI Evals](https://github.com/openai/evals) is an open-source software framework that provides tools for creating automated evals.

Model-based evals can be useful when there exists a range of possible outputs that would be considered equally high in quality (e.g. for questions with long answers). The boundary between what can be realistically evaluated with a model-based eval and what requires a human to evaluate is fuzzy and is constantly shifting as models become more capable. We encourage experimentation to figure out how well model-based evals can work for your use case.

Suppose it is known that the correct answer to a question should make reference to a specific set of known facts. Then we can use a model query to count how many of the required facts are included in the answer.

For example, using the following system message:

Here's an example input where both points are satisfied:

Here's an example input where only one point is satisfied:

Here's an example input where none are satisfied:

There are many possible variants on this type of model-based eval. Consider the following variation which tracks the kind of overlap between the candidate answer and the gold-standard answer, and also tracks whether the candidate answer contradicts any part of the gold-standard answer.

Here's an example input with a substandard answer:

Here's an example input with a good answer:
