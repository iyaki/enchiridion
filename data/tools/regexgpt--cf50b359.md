---
title: "RegExGPT"
notion_id: cf50b359-4457-42f7-889e-71d9a8f84020
notion_url: https://app.notion.com/p/RegExGPT-cf50b359445742f7889e71d9a8f84020
last_edited: 2023-09-08T17:30:00.000Z
source_url: https://regexgpt.app/
tags: ["English", "Programming", "Untried", "Service"]
---
RegexGPT is a tool that lets people generate regex patterns by inputting an example of the text they would like to transform and another input for the expected result. There is also a natural language input that allows the user to explain the pattern they would like to match. The final input lets users select the programming language for the output.

## How to use RegExGPT?

1. Enter an example of the text you would like to transform
2. Enter the expected result
3. Optionally, enter a natural language explanation
4. Select the programming language for the output
5. Click the "Generate" button to generate the regex pattern

## How to use the RegEx Pattern?

Once you have generated the regex pattern, you can use it in your code to transform the input text into the expected result. Below is an example usage of the generated regex pattern in JavaScript:

const regex = /{generated-regex-pattern}/;

const inputText = '{example-text}';

const result = inputText.replace(regex, '{expected-result}');

console.log(result); // Outputs the expected result
