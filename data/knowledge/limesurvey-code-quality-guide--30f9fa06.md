---
title: "LimeSurvey Code Quality Guide"
notion_id: 30f9fa06-52ef-4c87-af0b-fce0a69167e3
notion_url: https://app.notion.com/p/LimeSurvey-Code-Quality-Guide-30f9fa0652ef4c87af0bfce0a69167e3
last_edited: 2022-12-21T13:39:00.000Z
source_url: https://manual.limesurvey.org/Code_quality_guide
tags: ["English", "PHP", "Guide"]
---
## LimeSurvey Code Quality Guide

## Principles

1. Design to isolate change and to prepare for change by making parts composable and interchangeable.
2. It's harder to read code than to write code. Choose patterns that are easy to read, not easy to write.
3. Make your code communicate intent.
4. Always leave code better than you found it.
5. Remember cross-cutting concerns: translation and security.
6. Fail hard, fail early. The earlier the program crashes when something is wrong, the less likely you'll end up with corrupt data.

## Classes

### Rules

- Avoid empty OOP ceremony — don't define getters and setters for all properties. Make them public unless they have invariants to uphold.
- Avoid static methods — they cannot be mocked. Exception: factory methods (Foo::make(...)).
- Inject all dependencies via the constructor to enable unit-testing.
- Keep inheritance hierarchies shallow (max 2-3 levels deep).
- Prefer composition over inheritance.

### Types of classes in LimeSurvey

- ActiveRecord — communicates with the database only; no business logic
- Controllers — glue code only; no business logic; includes CLI controllers
- Helper functions — short, obey command-query separation; no dependencies
- Helper classes — like helper functions but have injectable dependencies; no state
- Service classes — clear, separate task logic (e.g. CreateSurvey, ImportSurvey)
- Question render classes — logic for how questions are rendered
- Plugin classes — using the plugin event system
- Widgets — HTML logic for custom HTML widgets

### Naming

- Suffixes: Exception, Trait, Controller, Command
- Prefixes: Abstract
- ActiveRecord classes: singular of the table name (table "lime_users" → model "User")
- Command classes: imperative verb (e.g. CreateSurveyCommand)
- Helper classes: as specific as possible to prevent growth

### Interfaces

- Use interfaces to define contracts between modules and enable mocking in tests.

## Functions

### Size

- Maximum ~60 lines (one screen)
- Maximum 4 arguments; if more, make it a class
- If a function grows too big, split it or create a class

### Naming

- Name tells what it does; docs tell why
- Pattern: verb [+ adjective] + noun — e.g. createFieldmap(), quoteText(), getRelativePath()
- Prefix boolean functions with "is" — e.g. isCaptchaEnabled()
- Never use "and" in function names — split into two functions

### Contract

- First docblock line: one sentence describing the relation between arguments and output
- Document all parameters with correct types
- Don't use array when it's a list of objects — use Survey[]

Example:

```plain text
/** * Creates a random password of length $length (defaults to 12) * @param int $length * @return string */ function createPassword($length = 12)
```

### Side-effects and testability

A side-effect is anything that makes a function return different results for the same input (database, file system, randomization, echo).

- Functions should either have side-effects or return a result — not both (command-query separation).
- Move side-effects higher up in the stack trace to make functions testable.

Example — lift the dependency:

```plain text
// Bad: queries database internally, needs fixture to test function getSurveyInfo(int $surveyid, string $languagecode = ''): array; // Good: caller provides the data, function is now testable function getSurveyInfo(Survey $survey, string $languagecode = ''): array;
```

### Assertions

- Use assert() for internal invariants — never seen by end users.
- Use exceptions for external inputs (browser, database, file system).
- Assertions have zero runtime cost when disabled in production.

Example:

```plain text
function gT($sToTranslate, $sEscapeMode = 'html', $sLanguage = null) { assert(is_string($sToTranslate)); assert(strlen($sToTranslate) > 0); assert($sEscapeMode === 'html' || $sEscapeMode === 'js' || $sEscapeMode === 'unescaped'); return quoteText(Yii::t('', $sToTranslate, array(), null, $sLanguage), $sEscapeMode); }
```

### Exceptions

- Used for inputs from the outside world (browser, database, file system)
- Don't use for control flow — only for fatal situations
- Provide actionable error messages
- Never do an empty catch-block
- Throw InvalidArgumentException for violated pre-conditions

### If-statements

Factor long conditions into named methods:

```plain text
// Bad if ((empty($instance->oOptions->{$attribute})) || (!empty($instance->oOptions->{$attribute}) && ($instance->oOptions->{$attribute} === 'inherit' || $instance->oOptions->{$attribute} === 'I' || $instance->oOptions->{$attribute} == '-1'))) { // Good $value = $instance->options->{$attribute}; if (empty($value) || $this->isInherit($value)) { }
```

### Looping

- Factor loop bodies into separate functions.
- Use plural→singular naming: foreach ($questions as $question)
- Loop indices: $i, $j, $k

## Variables

### Naming rules

- Follow PSR-12: camelCase, no underscores
- ActiveRecord variables: match class name — $survey, $question
- Plural for lists: $surveys, $questions
- Never include "array" or "list" in names — plural is enough
- $qid, $gid, $sid are acceptable established abbreviations
- Use $result only when baking a return value; prefer a more meaningful name
- Use /** @var */ annotations instead of Hungarian notation

Examples:

```plain text
// Bad $result = LabelSet::model()->findAll(); // Good $labelSets = LabelSet::model()->findAll(); // Type annotation /** @var array<string, string> */ $data = [];
```

## Code duplication

### When to refactor duplicates

- Ask: will the clones always change together, or will they diverge?
- The longer the duplicate, the stronger the argument to refactor it.
- Single-file duplicates almost always change together — refactor them.

### Common bugs in duplicated code

- Missing null check / isset()
- Inconsistent default values
- Inconsistent validation behaviour

### Example

```plain text
// Duplicated $this->connection->createCommand("SQL COMMAND")->execute(); $this->connection->createCommand("OTHER SQL COMMAND")->execute(); // Refactored public function executeCommand($command) { $this->connection->createCommand($command)->execute(); }
```

### Tools

- phpcpd — PHP copy-paste detector
- jscpd — JS/multi-format copy-paste detector
- SonarQube, Scrutinizer — full-system analysis

## Security

- XSS: Use ActiveRecord filters; always encode/decode output
- CSRF: Always use POST with CSRF token for state-changing actions (delete, add, update)
- SQL Injection: Use prepared statements and bind params — never concatenate user input into SQL
- Permissions: Always check permission in controller actions before executing logic

## Performance

- Never load all rows into PHP models when you only need a subset — use LIMIT/pagination.
- Test with large datasets to ensure scalability.

## Localization

### Basics

Currently we are using the gT(),eT(),ngT() and neT() functions for translations.

Since LimeSurvey is available in 60 languages it is very important that your original English string (which will be picked up automatically for translation by our translators) is done right from the start. Imagine if we have to correct only one string at a later time - then 60 strings become invalid across the project and need to be re-translated by the poor translator souls.

- gT() is the original translation function. It will return a translated version of the string in the currently selected language. ie: echo gT("Hello");
- ngT() returns multiple translations of a sentence or phrase for which alternate plural forms may apply. ie: echo sprintf(ngT('Please select at least %s answer','Please select at least %s answers',iMinimumAnswers),$iMinimumAnswers);
- eT() echos the translation directly. (ie: instead of echo gT("Hello"); you can use eT("Hello");
- neT() echos the multiple translations of the sentence or phrase for which alternate plural forms may apply directly.

Please follow these important rules:

### Spacing

- Do not embed margin spaces in your translation. Instead use proper CSS formatting

Wrong:

```plain text
eT('Visible? ');
```

Right:

```plain text
eT('Visible?');
```

### Punctuation

- Punctuation is always part of the translation

Wrong:

```plain text
echo gT('Summary').':<br>'.gT('This is a great summary.');
```

Right:

```plain text
echo gT('Summary:').'<br>'.gT('This is a great summary.');
```

### Linebreaks

- Do not embed system linebreaks in your translation.

Wrong:

```plain text
eT('This is a very long text. We hope that you will really read it.');
```

Right:

```plain text
eT('This is a very long text. We hope that you will really read it.');
```

### Capitalization

- Do not capitalize words except where grammatically correct (like at the beginning of a sentence or for a brand name). LimeSurvey is not a newspaper.

Wrong:

```plain text
eT('Create A New Label Set'); eT('Google Maps API Key');
```

Right:

```plain text
eT('Create a new label set'); eT('Google Maps API key');
```

### Full sentences

- Do not split sentences across several translation strings to form a sentence. If it is one sentence it should be in one translation unit.

Wrong:

```plain text
echo gT('This is a very long sentence and'); echo '<br>'; echo gT('this is the rest of it. This will be confusing for the translator.');
```

Right:

```plain text
echo gT('This is a very long sentence and this is the rest of it.'); echo '<br>'; echo gT('This will not be confusing for the translator.');
```

### Concatenation

- Do not concatenate several translations to form a sentence or concatenate to an additional information (like a number or string). Instead use the sprint() or sprintf() function with placeholders.

Wrong:

```plain text
echo gT('The sum must not be bigger than ') . $maxSum;
```

Right:

```plain text
echo sprintf(gT('The sum must not be bigger than %d'), $maxSum);
```

This will avoid the problem that in other languages the word positioning of a sentence can be very different and the information $maxSum from the example above might be needed at a different position in the sentence, not just at the end.

Wrong:

```plain text
echo gT('The user ') . $username . gT(' was deleted.');
```

Right:

```plain text
echo sprintf(gT('The user %s was deleted.'), $username);
```

Also don't concatenate words or sentences using place holders.

Wrong:

```plain text
echo sprintf(gT('The user was %s.')), gT( $action=='deleted'?'deleted':'created' );
```

Right:

```plain text
if ($action == 'deleted') { eT('The user was deleted.') } else { eT('The user was created.'); }
```

### Plurals

- Use the n*T functions where applicable. These functions show a different translation depending the number of items. Currently LimeSurvey only supports the numbers/situations 1 and >1.

Wrong:

```plain text
echo sprintf(gT('Please select at least %d answer(s).'), $minimumAnswers);
```

Right:

```plain text
echo sprintf( ngT('Please select at least %s answer', 'Please select at least %d answers', $minimumAnswers), $minimumAnswers );
```

### Embedded HTML

- Text should not contain HTML. If you desperately need HTML in there (for a link) use printf() and %s tags, but even this is a bad solution because if the replacement tags are not properly reflected in the translation, the application could break. So, better don't use this at all.

Wrong:

```plain text
eT('This is an <a href="importance.html">important</a> message.');
```

Right (still avoid this at all):

```plain text
printf(gT('This is an %simportant%s message.'),'<b>','</b>');
```

### Variables

- Don't embed a variable or constant instead of the real string. The text will not be picked up by our translation system.

Wrong:

```plain text
$strMessage = 'Very important message, that needs to be translated'; eT($strMessage);
```

Right:

```plain text
eT('Very important message, that needs to be translated');
```

### JavaScript

The script that analyzes the PHP files for translation strings do not work on JavScript. That means that all translatable strings must be written in PHP, using gT() or eT() functions.

The solution to this is to register a global JavaScript variable via PHP, to register translations. Example code in a controller, using heredoc and regsiterScript:

```plain text
$undo = gT("Undo (ctrl + Z)", "js"); $redo = gT("Redo (ctrl + Y)", "js"); $find = gT("Find (ctrl + F)", "js"); $replace = gT("Replace (ctrl + H)", "js"); App()->getClientScript()->registerScript( "SurveyThemeEditorLanguageData", <<<JAVASCRIPT surveyThemeEditorLanguageData = { undo: "$undo", redo: "$redo", find: "$find", replace: "$replace" }; JAVASCRIPT, CClientScript::POS_BEGIN );
```

## Testing

### Types

- Unit tests — no fixture, uses mocking/stubs, runs in parallel, fast
- Functional tests — with fixture/database setup
- Integrity tests — full install, scripted browser; slow but requires no code adaptation

### Rules

- Inject dependencies via constructor to enable mocking.
- Functions outside classes cannot be mocked — only use them when mocking is not needed.
- Don't test the framework; don't test PHP itself.
- Include negative tests (verify correct failure behaviour).

Example — dependency injection for testability:

```plain text
// Bad: implicit dependency, cannot mock class CopyQuestion { private function copySubquestions($parentId) { $subquestions = \Question::model()->findAllByAttributes(['parent_qid' => $parentId]); } } // Good: injectable, mockable class CopyQuestion { private $questionModel; public function __construct(Question $questionModel) { $this->questionModel = $questionModel; } private function copySubquestions($parentId) { $subquestions = $this->questionModel->findAllByAttributes(['parent_qid' => $parentId]); } }
```

## Static analysis

- Run php -l for syntax checking
- Use Psalm for type checking (strictness is configurable)
- Use CodeSniffer for style enforcement (PSR-12)
- Use phpcpd for duplicate detection

## The PHP of yesterday

Abandon these patterns:

- Don't use associative arrays for structured data — use data-transfer objects (no performance overhead)
- Use [] instead of array()
- Don't use Hungarian notation — use @param and @var annotations checked by static analysis

## The PHP of tomorrow

Adopt when the minimum PHP version allows:

- Union and intersection types
- Short closures (arrow functions)
- Enums
- Match expressions
- Attributes (instead of docblock annotations)
