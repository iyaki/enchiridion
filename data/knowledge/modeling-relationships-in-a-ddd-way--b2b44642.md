---
title: "Modeling Relationships in a DDD Way"
notion_id: b2b44642-e72a-41a4-b904-037fa2031481
notion_url: https://app.notion.com/p/Modeling-Relationships-in-a-DDD-Way-b2b44642e72a41a4b904037fa2031481
last_edited: 2026-09-21T17:40:00.000Z
source_url: https://enterprisecraftsmanship.com/posts/modeling-relationships-in-ddd-way/
tags: ["English", "Domain Driven Design", "System Design / Software Architecture", "Databases", "Article", "Enterprise Craftsmanship"]
---
[https://enterprisecraftsmanship.com/posts/modeling-relationships-in-ddd-way/](https://enterprisecraftsmanship.com/posts/modeling-relationships-in-ddd-way/)

April 19, 2022

Let’s talk about modeling of relationships, including the dreaded many-to-many relationships, in a DDD way.

## 1. Types of relationships

There are 4 types of relationships between tables in a relational database:

- One-to-many,
- Many-to-one,
- Many-to-many,
- One-to-one.

Let’s discuss each of these types separately.

### 1.1. One-to-many relationships

This is the most common type of relationship. To give an example, let’s say we have a Student and Enrollment tables in our database. The relationship between them would be one-to-many:

One-to-many relationship

The notation of "one-to-many" reads as follows:

- For each student row, there can be multiple enrollment rows — This is what the second word in the X-to-Y notation means.
- For each enrollment, there can be only one student — Which is what the first word in X-to-Y refers to.

At the database level, you can introduce a one-to-many relationship by creating a foreign key constraint from the Enrollment table to Student's Id.

### 1.2. Many-to-one relationships

The inverse of one-to-many is a many-to-one relationship. For each one-to-many relationship you automatically get a many-to-one relation that goes in the opposite direction. In the above example, it’s the relationship between enrollments and students.

Another example is a relationship between students and their favorite courses:

Many-to-one relationship

The only difference between one-to-many and many-to-one relationships is which side of the relation you are reading it from.

### 1.3. Many-to-many

A many-to-many relationship would be between students and instructors, where each student can have many instructors and each instructor can also have many students:

Many-to-many relationship

In the database, you can create a many-to-many relationship by stacking two one-to-many relationships. You can do that using an intermediate table, such as Student2Instructor. Both students and instructors would have a one-to-many relationship with that table.

### 1.4. One-to-one relationships

Finally, a one-to-one relationship is when each row in one table can only have one row in the related table, in both directions.

A typical example is a student and his or her student details:

One-to-one relationship

At the database level, a one-to-one relationship can be introduced by creating a foreign key constraint from StudentDetailsID to StudentID.

## 2. Object–relational impedance mismatch

The above section shows how types of relationships look in a relational database: as associations between tables (represented by foreign key constraints).

What about the domain model? How can different types of relationships be represented there?

Let’s take the most common type of relationship as an example: a one-to-many relationship. At the domain model level, one-to-many relationships are represented with collections on one side and with a singular reference on the other:

One-to-many relationship in the domain model

Here’s how it may look in code:

```plain text
public class Student { public string Name { get; } public string Email { get; } public ICollection<Enrollment> Enrollments { get; } // one-to-many } public class Enrollment { public string Grade { get; } public string Course { get; } public Student Student { get; } // many-to-one }
```

Notice the difference between the database and domain models:

- To create an association between Student and Enrollment in the database, you only need to set up the enrollment’s StudentID,
- But in the domain model, you have to set up both, the student’s Enrollments and the enrollment’s Student.

And that’s where we step into the territory of the object–relational impedance mismatch. This term describes a set of conceptual and technical difficulties that we encounter when we try to combine an object-oriented programming language with a relational database.

This difference between how relationships are represented in the domain model and in the database is the most common and annoying of such difficulties.

In the database, all relationships are represented with a foreign key constraint. There’s no such notion as a one-to-many or many-to-many relationship at the database level. All your database knows about is the foreign key from one table to another. It’s us, programmers, who ascribe meaning to these foreign keys depending on the situation.

There’s also no such thing as unidirectional vs bidirectional relationships in the database. At the database level, relationships are always bidirectional: each foreign key constraint always involves two tables.

In the domain mode, however, we may decide to keep only one end of the relationship:

```plain text
public class Student { public string Name { get; } public string Email { get; } public Course FavoriteCourse { get; } // many-to-one } public class Course { // no one-to-many public string Name { get; } }
```

And therefore, make the relationship unidirectional instead of bidirectional.

This difference between how relationships look in the database and in the domain model leads to complications when you try to persist your domain objects.

For example, what if we set the two ends of the one-to-many relationship such that they are inconsistent with each other? How would the following domain objects be persisted?

```plain text
Student student1 = GetStudent(1); Student student2 = GetStudent(2); var enrollment = new Enrollment { Student = student2 /* Setting student2 here... */ }; student1.Enrollments.Add(enrollment); /* ...but then using it for student1 */
```

This inconsistency is only possible because the relationship has a representation in both the Student and Enrollment classes. In the database, no such inconsistencies may take place because the relationship exists in the form of just one column: Enrollment.StudentID.

ORMs use quite sophisticated heuristics to figure out which domain class is responsible for each particular relationship. For the Student and Enrollment classes, for instance, an ORM may default to treating the Enrollment.Student property as the owner of the relationship. But if this property isn’t declared, then the ORM may inverse the ownership and make Student.Enrollments responsible for that relationship instead.

As a general guideline, try to always make relationships unidirectional when possible. This would simplify your domain model and also alleviate the object–relational impedance mismatch.

## 3. One-to-one relationships

Let’s now talk about one-to-one relationships. In the domain model, such relationships are represented with a singular reference on both ends:

One-to-one relationship in the domain model

Here’s how it looks in code:

```plain text
public class Student { public string Name { get; } public string Email { get; } public StudentDetails Details { get; } // one-to-one } public class StudentDetails { public string Address { get; } public string Preferences { get; } public string SportsParticipation { get; } public Student Student { get; } // one-to-one }
```

In the above code, each Student may only have one details object and each StudentDetails may have only one student.

This type of relationship is completely unnecessary. You don’t get any benefits out of it. The only thing one-to-one relationships do is they clutter your database.

If two tables relate to each other as one-to-one, merge them into one table.

You don’t really need to split the Student table like this. All the student details data can (and should) be put into the Student itself. This would simplify your database because, all else equal, it’s easier to deal with one table than two.

The only exception is when you need to split a table for performance reasons. For example, when the student details contains blobs, such as files or large chunks of text, that you don’t want to retrieve alongside the student itself.

Alright, but how the domain model would look after removing the StudentDetails table? Do you have to move all its properties to the Student class?

Not necessarily. If the grouping of properties in StudentDetails makes sense from a business perspective, you can keep that grouping, but instead of modeling it as an entity, represent it as a value object:

```plain text
public class Student : Entity { public string Name { get; } public string Email { get; } public StudentDetails Details { get; } } public class StudentDetails : ValueObject { public string Address { get; } public string Preferences { get; } public string SportsParticipation { get; } }
```

For all practical purposes, value objects behave exactly the same way as one-to-one relationships but have two significant advantages:

- They are stored alongside the host entity, so no additional tables in the database are needed.
- They are easier to work with because of their immutability.

## 4. Many-to-many relationships

Finally, let’s discuss many-to-many relationships. In the domain model, such relationships are represented by collections on both sides of the relation:

Many-to-many relationship

This is how it may look in code:

```plain text
public class Student : Entity { public string Name { get; } public string Email { get; } public ICollection<Instructor> Instructors { get; } // many-to-many } public class Instructor : Entity { public string Name { get; } public ICollection<Student> Students { get; } // many-to-many }
```

This type of relationship is usually the most tricky one. The problem is that it’s hard to decide which side of a many-to-many relation is the owner of that relation.

In other words, which class should be the entry point for setting up links between students and instructors? You have to pick one to avoid inconsistencies, where the relation is set up on one side but not on the other.

Note that there’s no such issue with one-to-many relationships. One-to-many relationships are almost always set up from the one-to-many end, which is the aggregate root. In our previous example, for instance, the Student is the aggregate root and the Enrollment is an internal entity in the same aggregate.

In many-to-many relationships, though, the two sides of the relation are equally illegible to become the owner because both sides are usually aggregate roots.

### 4.1. Reducing parity of many-to-many relationships

So, what to do here? The first thing you need to consider is reducing parity of the many-to-many relationship by making it unidirectional instead of bidirectional.

In a lot of cases, you don’t really need both sides of the relation in your domain model. For example, the Instructor might need the collection of students, but does the Student also need the collection of instructors? If you can manage without it, then remove that collection, and solve the problem this way.

But what if you can’t reduce parity? How should you act then?

In this case, you just have to make a choice and decide which side of the relation is going to be the owner of it. In the example with Instructor and Student, it should probably be the Instructor:

```plain text
public class Student : Entity { public string Name { get; } public string Email { get; } private readonly IList<Instructor> _instructors; public IReadOnlyList<Instructor> Instructors => _instructors; internal void AddInstructor(Instructor instructor) { _instructors.Add(instructor); } } public class Instructor : Entity { public string Name { get; } private readonly IList<Student> _students; public IReadOnlyList<Student> Students => _students; public void AddStudent(Student student) { _students.Add(student); student.AddInstructor(this); } }
```

Notice that both Instructor and Student expose an AddXXX method, but only Instructor has this method defined as public and it’s responsible for setting up both ends of the relationship. The corresponding method in the Student is marked as internal, to avoid any ambiguities regarding which class is the owner here.

### 4.2. Working with intermediate tables in many-to-many relationships

Mapping of many-to-many relationships onto the domain model also raises issues. As I mentioned previously, there’s no notion of many-to-many relationship at the database level; that’s just two one-to-many relationships stacked upon each other with the help of an intermediate table.

How do you work with that intermediate table in the domain model?

Follow this guideline:

- If the intermediate table only contains references to the related tables, then don’t introduce a class for that table.
- If the intermediate table contains other information, then do introduce a class for it.

In the above example, Student2Instructor only contains references to Student and Instructor, and so we shouldn’t introduce a separate class for that table.

But if we were to add, let’s say, a DateAdded field, then we should promote that table into a separate domain class, like this:

```plain text
public class Student : Entity { public string Name { get; } public string Email { get; } private readonly IList<StudentInstructor> _studentInstructors; public IReadOnlyList<Instructor> Instructors => _studentInstructors .Select(x => x.Instructor) .OrderBy(x => x.DateAdded) .ToList(); internal void AddInstructor(StudentInstructor instructor) { _studentInstructors.Add(instructor); } } public class Instructor : Entity { public string Name { get; } private readonly IList<StudentInstructor> _studentInstructors; public IReadOnlyList<Student> Students => _studentInstructors .Select(x => x.Student) .OrderBy(x => x.DateAdded) .ToList(); public void AddStudent(Student student) { var studentInstructor = new StudentInstructor(student, this, DateTime.Now); _studentInstructors.Add(studentInstructor); student.AddInstructor(studentInstructor); } } public class StudentInstructor : Entity { public Student Student { get; } public Instructor Instructor { get; } public DateTime DateAdded { get; } public StudentInstructor(Student student, Instructor instructor, DateTime dateAdded) { Student = student; Instructor = instructor; DateAdded = dateAdded; } }
```

Notice that the Student and Instructor still expose a collection of instructors and students; they don’t expose objects of the new StudentInstructor class. It’s always a good idea to do so if the client code doesn’t need the additional information from the StudentInstructor. If the client does need that info, however, then we’d need to expose the StudentInstructor directly.

This refactoring has essentially converted our many-to-many relationship into two one-to-many relationships.

Note that I’ve also renamed Student2Instructor into StudentInstructor. I recommend to reserve the X2Y naming pattern only for tables that aren’t present in the domain model. This will help you distinguish many-to-many relationships from one-to-many relationships at the database level.

Here are other changes to the intermediate table:

Many-to-many intermediate table conversion

Try to avoid composite primary keys in the database. Only use them for many-to-many intermediate tables that don’t contain any other data. For all other tables, use a single-column primary key. And if you need uniqueness guarantees, just introduce a separate unique index.

## 5. Summary

- There are 4 types of relationships between tables in a relational database: One-to-many, Many-to-one, Many-to-many, One-to-one.
- Try to reduce parity of the relationship by making it unidirectional instead of bidirectional.
- Replace one-to-one relationships with value objects.
- When working with intermediate tables in many-to-many relationships: Don’t introduce a class for that table if it only contains references to related tables. Do introduce a class for that table if it contains any other information. In this case, convert the many-to-many relationship into two one-to-many relationships.

## Subscribe

## Comments

comments powered by
