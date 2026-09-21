---
title: "DDD and bulk operations"
notion_id: 32b714e0-8428-4085-a6c1-7afcb6ea283f
notion_url: https://app.notion.com/p/DDD-and-bulk-operations-32b714e084284085a6c17afcb6ea283f
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/ddd-bulk-operations/
tags: ["Programming", "System Design / Software Architecture", "Object Oriented Programming", "Domain Driven Design", "Article", "Enterprise Craftsmanship"]
---
[https://enterprisecraftsmanship.com/posts/ddd-bulk-operations/](https://enterprisecraftsmanship.com/posts/ddd-bulk-operations/)

October 28, 2019

Combining bulk operations with Domain-Driven Design is a tough problem. In this article, we’ll look at why that is so and discuss ways to marry the two.

This article is also a response to a reader question. The question contained an interesting example, which I’ll use throughout this post:

> Hi Vladimir! Do you have any article about bulk actions in a DDD context? I didn’t find anything useful concerning this. Please, consider the following example: Given a list of tasks, I want to set an execution date to all tasks that match a selected month and category, Also, I can’t set an execution date for an already completed task, For a given month and category, there can be up to 30,000 tasks. Currently I created a SetExecutionDateDomainService that: Queries tasksRepository.GetBy(month, category), For each task checks if task.CanSetExecutionDate(), Calls taskRepository.Update(task). Any comments/tips on how to handle this?

There are three ways to handle this problem:

- Process objects one by one (the way the question author does it),
- Rely on SQL bulk updates,
- Use a combination of the Specification and Command patterns.

The first two options have trade-offs. I particularly like the third one.

## Sequential processing

The most straightforward way to handle this problem is to retrieve all suitable objects, and update them sequentially, one by one:

```plain text
public class Task { public int Month { get; private set; } public string Category { get; private set; } public bool IsCompleted { get; private set; } public DateTime? ExecutionDate { get; private set; } public bool CanSetExecutionDate() { return IsCompleted == false; } public void SetExecutionDate(DateTime executionDate) { Guard.Require(CanSetExecutionDate(), "CanSetExecutionDate()"); ExecutionDate = executionDate; } } public class SetExecutionDateService { public void SetExecutionDate(int month, string category, DateTime executionDate) { IReadOnlyList<Task> tasks = _repository.GetBy(month, category); foreach (Task task in tasks) { if (task.CanSetExecutionDate()) { task.SetExecutionDate(executionDate); _repository.Update(task); } } } }
```

The main benefit of this solution is that all domain knowledge is contained within the domain model. Specifically, the knowledge of when the execution date can be set (the CanSetExecutionDate method).

The drawback here is the lack of performance: processing and updating tasks separately requires a lot of database roundtrips — one per each such update.

DDD in general doesn’t work particularly well outside of OLTP-type operations (transactional processing of smalls amounts of data). This is also true for this use case of bulk updating a large number of tasks — it falls outside DDD’s "area of comfort".

> A bulk operation (or a bulk update) is updating large amounts of data in one database roundtrip.

## Using raw SQL

If DDD doesn’t work well with bulk updates, what does? That’s where raw SQL shines. SQL is specifically designed to work with large sets of related data, and we can use it for our scenario too:

```plain text
UPDATE dbo.Task SET ExecutionDate = @ExecutionDate WHERE Category = @Category AND Month = @Month AND IsCompleted = 0 -- domain knowledge duplication
```

This approach is fast and simple but violates the DRY principle: you have to put the knowledge of what tasks are eligible for setting the execution date to both the SQL (the IsCompleted = 0 line) and application code (the CanSetExecutionDate method).

The use of raw SQL might not be that bad of an option, especially in simple projects, but there’s a better way.

## Using the Specification pattern

I’ve been writing about the Specification pattern for a while now:

- Specification pattern: C# implementation,
- Pluralsight course,
- CQRS vs Specification pattern.

In short, the Specification pattern is about encapsulating a piece of domain knowledge into a single unit (called specification), which you then reuse in three scenarios:

- Data retrieval
- In-memory validation
- Creation of new objects ("Construction-to-order" on the figure below).

Specification pattern

I’ve also written that, although the idea of this pattern looks interesting, it goes against the CQRS pattern and thus should be discarded. The reason is that CQRS provides another benefit — loose coupling, which is more important than DRY in the vast majority of cases.

CQRS vs Specification pattern

CQRS achieves loose coupling by splitting a single unified model into two: one for reads (data retrieval, the sphere of raw SQL queries) and the other one for writes (in-memory validation, the sphere of DDD). This separation is where the contradiction lies: the Specification pattern advocates for keeping the unified model.

So how the Specification pattern helps in the scenario with bulk updates?

It turns out that you can use specifications not only for querying the database, but updating it too. Let me first show a typical use for this pattern. I’ll then demonstrate how it can be extended for the bulk update use case.

In the above use case with setting tasks' execution date, we’ll need the following three specifications:

```plain text
public sealed class TaskIsCompletedSpecification : Specification<Task> { public override Expression<Func<Task, bool>> ToExpression() { return task => task.IsCompleted; } } public sealed class TaskMonthSpecification : Specification<Task> { private readonly int _month; public TaskMonthSpecification(int month) { _month = month; } public override Expression<Func<Task, bool>> ToExpression() { return task => task.Month == _month; } } // + TaskCategorySpecification, which is the same as TaskMonthSpecification
```

With these specifications, Task now looks as follows:

```plain text
public class Task { public int Month { get; private set; } public string Category { get; private set; } public bool IsCompleted { get; private set; } public DateTime? ExecutionDate { get; private set; } public bool CanSetExecutionDate() { var spec = new TaskIsCompletedSpecification(); '1 return spec.IsSatisfiedBy(this) == false; '1 } public void SetExecutionDate(DateTime executionDate) { Guard.Require(CanSetExecutionDate(), "CanSetExecutionDate()"); ExecutionDate = executionDate; } }
```

Note the use of TaskIsCompletedSpecification in '1. It may look redundant (after all, this specification checks the IsCompleted field of the same Task instance), but it’s important to be consistent in allocating domain knowledge across your application. Once you introduce a specification to hold a piece of that knowledge, all other classes should start using it too to comply with the DRY principle.

Here’s the domain service:

```plain text
public class SetExecutionDateService { public void SetExecutionDate(int month, string category, DateTime executionDate) { var monthSpec = new TaskMonthSpecification(month); var categorySpec = new TaskCategorySpecification(category); var isNotCompletedSpec = new TaskIsCompletedSpecification().Not(); Specification<Task> spec = monthSpec.And(categorySpec).And(isNotCompletedSpec); '1 IReadOnlyList<Task> tasks = _repository.GetList(spec); '2 foreach (Task task in tasks) { if (task.CanSetExecutionDate()) { task.SetExecutionDate(executionDate); _repository.Update(task); } } } }
```

The domain service combines the three specifications (line '1) and passes them to the repository ('2). The repository looks as follows (I’m using NHibernate but the code is the same for Entity Framework):

```plain text
public IReadOnlyList<Task> GetList(Specification<Task> specification) { return _session.Query<Task>() .Where(specification.ToExpression()) .ToList(); }
```

This code relies on a sophisticated ORM functionality that traverses the specification’s expression tree and converts it into SQL. For example, this combined specification

```plain text
var monthSpec = new TaskMonthSpecification(month); var categorySpec = new TaskCategorySpecification(category); var isNotCompletedSpec = new TaskIsCompletedSpecification().Not(); Specification<Task> spec = monthSpec.And(categorySpec).And(isNotCompletedSpec);
```

gets translated into

```plain text
Month = @Month AND Category = @Category AND NOT(IsCompleted = 1)
```

C# expressions in combination with an ORM are a powerful duo. But even they only take you so far. ORMs allow you to use expressions to query the database but not update it. Updating the database is what we need in order to implement the bulk update functionality (set the execution date to all tasks in one database roundtrip).

So, what to do, then?

The good news is that you don’t have to rely neither on ORMs nor C# expressions to work with the database using the Specification pattern. Expression trees are a handy tool that simplifies the implementation of your specifications, but they are just that — a tool, of which you have many.

Another tool is the plain SQL itself. In fact, you can combine the two approaches: use expression trees for in-memory validation and querying the database, and raw SQL for bulk updates. The idea is that, in addition to the ToExpression method, each specification has to also implement ToSql() where it would generate an appropriate filter for the update SQL query.

Here’s how the base Specification class looks (again, check out this GitHub repository for the full source code):

```plain text
public abstract class Specification<T> { public bool IsSatisfiedBy(T entity) { Func<T, bool> predicate = ToExpression().Compile(); return predicate(entity); } public abstract Expression<Func<T, bool>> ToExpression(); /* And(), Or(), Not() methods */ }
```

You need to add two new abstract methods:

```plain text
public abstract string ToSql(); public abstract IEnumerable<SqlParameter> ToSqlParameters();
```

ToSql converts the specification into SQL, and ToSqlParameters provides required parameters for that SQL.

Now you need to implement these two methods in all specification subclasses. Here’s an example:

```plain text
public sealed class TaskMonthSpecification : Specification<Task> { private readonly int _month; public TaskMonthSpecification(int month) { _month = month; } public override Expression<Func<Task, bool>> ToExpression() { return task => task.Month == _month; } public override string ToSql() { return "[Month] = @Month"; } public override IEnumerable<SqlParameter> ToSqlParameters() { yield return new SqlParameter("Month", _month); } }
```

And, finally, this is how the bulk update looks:

```plain text
// Domain service public void SetExecutionDate(int month, string category, DateTime executionDate) { var monthSpec = new TaskMonthSpecification(month); var categorySpec = new TaskCategorySpecification(category); var isNotCompletedSpec = new TaskIsCompletedSpecification().Not(); Specification<Task> spec = monthSpec.And(categorySpec).And(isNotCompletedSpec); _repository.UpdateExecutionDate(executionDate, spec); } // TaskRepository public void UpdateExecutionDate(DateTime executionDate, Specification<Task> specification) { string sql = @" UPDATE dbo.Task SET ExecutionDate = @ExecutionDate WHERE " + specification.ToSql(); using (DbCommand command = _session.Connection.CreateCommand()) { command.CommandText = sql; command.Parameters.AddRange(specification.ToSqlParameters().ToArray()); command.Parameters.Add(new SqlParameter("ExecutionDate", executionDate)); command.ExecuteNonQuery(); } }
```

Such a use of the Specification pattern brings up the forth scenario, bulk updates:

Specification pattern, revisited

Note that this use case doesn’t contradict CQRS: the reuse of domain knowledge for in-memory validation and bulk updates takes place within the write part of your application. Hence, I’d like to retract my earlier statement that specifications are only useful in simple scenarios (in which loose coupling isn’t that important). Bulk updates are a perfectly valid use case for this pattern, and that use case can occur in applications of any complexity.

With the above implementation, the business requirements regarding how to select tasks for a bulk update are all located in the domain layer. Those requirements are a combination of three preconditions, all of which are encapsulated in specifications:

1. Tasks of a specific month,
2. Tasks with a specific category,
3. Tasks that are not completed.

So, problem solved? Not really (yet). While we’ve encapsulated the knowledge of what tasks are suitable for an update, the update itself is still scattered between the Task domain class and TaskRepository ('1 and '2):

```plain text
public class Task { /* Month, Category, IsCompleted, ExecutionDate properties */ public bool CanSetExecutionDate() { var spec = new TaskIsCompletedSpecification(); return spec.IsSatisfiedBy(this) == false; } public void SetExecutionDate(DateTime executionDate) { Guard.Require(CanSetExecutionDate(), "CanSetExecutionDate()"); ExecutionDate = executionDate; '1 } } // TaskRepository public void UpdateExecutionDate(DateTime executionDate, Specification<Task> specification) { string sql = @" UPDATE dbo.Task SET ExecutionDate = @ExecutionDate '2 WHERE " + specification.ToSql(); using (DbCommand command = _session.Connection.CreateCommand()) { command.CommandText = sql; command.Parameters.AddRange(specification.ToSqlParameters().ToArray()); command.Parameters.Add(new SqlParameter("ExecutionDate", executionDate)); command.ExecuteNonQuery(); } }
```

This is another instance of domain logic duplication. To deal with this issue, we’ll need the other piece of the puzzle: the Command pattern.

## Meet the Command pattern

The duplication in the listing above might not seem like a big deal because it’s the assignment of just one field. But in fact, it is a big deal — there’s also a precondition that requires the task not to be completed for it to have an execution date:

```plain text
public void SetExecutionDate(DateTime executionDate) { /* This precondition is an intrinsic part of the execution date assignment */ Guard.Require(CanSetExecutionDate(), "CanSetExecutionDate()"); ExecutionDate = executionDate; }
```

The act of setting an execution date is the whole SetExecutionDate method, not just the assignment operation (=) in it. The precondition from that method is also present in the SQL query TaskRepository generates:

```plain text
UPDATE dbo.Task SET ExecutionDate = @ExecutionDate WHERE [Month] = @Month AND Category = @Category AND NOT(IsCompleted = 1) -- the precondition
```

The problem is that there’s nothing preventing TaskRepository from setting an execution date without consulting with this precondition. The connection between the IsCompleted and ExecutionDate fields is an important piece of domain knowledge, which you have to keep in mind and duplicate in both Task and TaskRepository.

And imagine having to assign not a primitive type like DateTime but a value object containing multiple fields. It becomes really easy to get the logic in Task and TaskRepository out of sync.

So, how to overcome this issue and avoid the duplication of the assignment logic? This is where the Command pattern comes into play.

The Command pattern essentially serves the same purpose as specifications but instead of examining the properties of domain objects, a command changes those properties. You can think of the differences between the two patterns as following:

- The Specification pattern encapsulates the knowledge of what data to update.
- The Command pattern encapsulates the knowledge of how to update that data.

Also, while you can use specifications in 4 scenarios, commands are only useful in two: in-memory updates and bulk updates.

Command pattern

This is how the Command base class looks:

```plain text
public abstract class Command<T> { /* Restrictions in addition to preconditions */ protected readonly IReadOnlyList<Specification<T>> _restrictions; '1 protected Command(IReadOnlyList<Specification<T>> restrictions) { _restrictions = restrictions; } /* Command's preconditions */ protected abstract IEnumerable<Specification<T>> GetPreconditions(); '2 private Specification<T> CombinedSpecification => GetPreconditions() .Concat(_restrictions) .Aggregate(Specification<T>.All, (x, y) => x.And(y)); protected abstract void ExecuteCore(T entity); protected abstract string GetTableName(); protected abstract string ToSqlCore(); protected abstract IEnumerable<SqlParameter> ToSqlParametersCore(); /* In-memory update */ public bool CanExecute(T entity) { return CombinedSpecification.IsSatisfiedBy(entity); } public void Execute(T entity) { if (CanExecute(entity) == false) throw new InvalidOperationException(); ExecuteCore(entity); } /* SQL for bulk update */ public string ToSql() { return @" UPDATE " + GetTableName() + @" SET " + ToSqlCore() + @" WHERE " + CombinedSpecification.ToSql(); } /* SQL parameters for bulk update */ public IReadOnlyList<SqlParameter> ToSqlParameters() { return CombinedSpecification.ToSqlParameters() .Concat(ToSqlParametersCore()) .ToArray(); } }
```

The class looks a bit large but the idea behind it is pretty simple — bake the preconditions into the command so that there isn’t even an option to omit those preconditions. In addition to the preconditions (line '2), there are additional restrictions ('1) that could be imposed on the command.

Here’s our bulk update command:

```plain text
public class SetExecutionDateCommand : Command<Task> { private readonly DateTime _executionDate; public SetExecutionDateCommand( DateTime executionDate, params Specification<Task>[] restrictions) : base(restrictions) { _executionDate = executionDate; } protected override IEnumerable<Specification<Task>> GetPreconditions() { yield return new TaskIsCompletedSpecification().Not(); } protected override void ExecuteCore(Task entity) { entity.ExecutionDate = _executionDate; } protected override string GetTableName() { return "dbo.Task"; } protected override string ToSqlCore() { return "ExecutionDate = @ExecutionDate"; } protected override IEnumerable<SqlParameter> ToSqlParametersCore() { yield return new SqlParameter("ExecutionDate", _executionDate); } }
```

And here’s the usage:

```plain text
// SetExecutionDateService public void SetExecutionDate(int month, string category, DateTime executionDate) { var monthSpec = new TaskMonthSpecification(month); '1 var categorySpec = new TaskCategorySpecification(category); '1 var command = new SetExecutionDateCommand(executionDate, monthSpec, categorySpec); _repository.BulkUpdate(command); } // TaskRepository public void BulkUpdate(SetExecutionDateCommand command) { using (DbCommand dbCommand = _session.Connection.CreateCommand()) { dbCommand.CommandText = command.ToSql(); dbCommand.Parameters.AddRange(command.ToSqlParameters().ToArray()); dbCommand.ExecuteNonQuery(); } }
```

Note that the specifications-restrictions ('1) are optional (you may or may not apply them to the command), but the specification-precondition is mandatory. In fact, you don’t even have an option to specify that precondition — it’s baked into the command itself. This is the essence of encapsulation: you can’t trust yourself to do the right thing all the time; you must eliminate the very possibility of doing the wrong thing.

Also note that I’m not familiar with the application’s specifics and assume the month and category restrictions are optional. If they are not, you should move them to the GetPreconditions method too, in which case the command and the domain service will become even simpler:

```plain text
public class SetExecutionDateCommand : Command<Task> { private readonly DateTime _executionDate; private readonly int _month; private readonly string _category; public SetExecutionDateCommand(DateTime executionDate, int month, string category) : base(new Specification<Task>[0]) { _category = category; _month = month; _executionDate = executionDate; } protected override IEnumerable<Specification<Task>> GetPreconditions() { yield return new TaskIsCompletedSpecification().Not(); yield return new TaskMonthSpecification(_month); yield return new TaskCategorySpecification(_category); } /* The rest is the same */ } // SetExecutionDateService public void SetExecutionDate(int month, string category, DateTime executionDate) { var command = new SetExecutionDateCommand(executionDate, month, category); _repository.BulkUpdate(command); }
```

Again, the raw SQL is probably still a better option for most projects due to its simplicity, even though it doesn’t adhere to the DRY principle. But the combination of the Specification and Command patterns may be useful for projects with sophisticated domain logic that you want to reuse between in-memory and bulk updates.

## Summary

- DDD is good for transactional processing of small amounts of data (OLTP) and doesn’t work well with bulk operations.
- A bulk operation (or a bulk update) is updating large amounts of data in one database roundtrip.
- There are three ways to handle bulk updates: Sequential processing (adheres to the DRY principle, bad for performance), Using raw SQL (good for performance, violates the DRY principle) Using a combination of the Specification and Command patterns (adheres to DRY and good for performance).
- Bulk operations is the forth use case for the Specification pattern, in addition to in-memory validation, querying the database, and creation of new objects.
- The Specification pattern encapsulates the knowledge of what data to update. The Command pattern encapsulates the knowledge of how to update that data. Both patterns allow you to reuse that knowledge between the domain model and bulk operations.
- Commands use specifications as preconditions for In-memory updates, Bulk updates.

## Subscribe

## Comments

comments powered by
