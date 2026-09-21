---
title: "EasyMock"
notion_id: 70305d30-d0c7-4b8b-a879-722f15458c0c
notion_url: https://app.notion.com/p/EasyMock-70305d30d0c74b8ba879722f15458c0c
last_edited: 2022-12-21T14:44:00.000Z
source_url: https://easymock.org/
tags: ["Programming", "Testing", "Untried", "Tool", "English"]
---
## Why

Great testing includes isolation

Most parts of a software system do not work in isolation, but collaborate with other parts to get their job done.

In a lot of cases, we do not care about using real collaborators implementation in unit testing, as we trust these collaborators.

Mock Objects replace collaborators of the unit under test.

## How

Isolation involves Mock Objects

To test a unit in isolation or mount a sufficient environment, we have to simulate the collaborators in the test.

A Mock Object is a test-oriented replacement for a collaborator. It is configured to simulate the object that it replaces in a simple way.

In contrast to a stub, a Mock Object also verifies whether it is used as expected.

## What

EasyMock makes mocking easier

EasyMock has been the first dynamic Mock Object generator, relieving users of hand-writing Mock Objects, or generating code for them.

EasyMock provides Mock Objects by generating them on the fly using Java proxy mechanism.
