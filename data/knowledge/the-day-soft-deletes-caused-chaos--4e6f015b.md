---
title: "The Day Soft Deletes Caused Chaos"
notion_id: 4e6f015b-8415-4eb8-9c9e-efacf8b038c2
notion_url: https://app.notion.com/p/The-Day-Soft-Deletes-Caused-Chaos-4e6f015b84154eb89c9eefacf8b038c2
last_edited: 2026-09-21T16:58:00.000Z
source_url: https://blog.bemi.io/soft-deleting-chaos/
tags: ["English", "Databases", "System Design / Software Architecture", "Article", "Bemi Blog"]
---
The worst mistake I made in my software engineering career was merging a seemingly harmless pull request 5 years ago.

TLDR: Soft deletes should not be used in production-grade systems—a lesson I learned the hard way when a severe mishap enabled the sale of the same concert seats to unlimited buyers.

Soft deletion is the easiest way to store deleted data and means just setting a deleted flag instead of performing a DELETE operation directly.

```plain text
+--------------------------------------+------------+------------+---------+ | id | last_name | first_name | deleted | +--------------------------------------+------------+------------+---------+ | 12778d88-41c8-4fc2-8be6-68c5d51c3893 | Samsonite | Mary | true | +--------------------------------------+------------+------------+---------+
```

Deleted user using soft deletion

You add a new column on a table, perform an update when deleting, and filter out deleted data when querying.

```plain text
SELECT * FROM user WHERE id = $1 AND deleted IS NULL;
```

Querying with soft deletion

Although this approach is simple to set up, it can lead to dangerous consequences if it accidentally returns data not meant to be seen.

## Critical Incident

I was working at an events ticketing company and I created a pull request that was similar to this:

```plain text
class SeatClaim ... - acts_as_paranoid ... + def remove + move_to_expired + destroy + end ... end
```

app/models/seat_claim.rb

```plain text
+ class MigrateDeletedSeatClaims < Migration + def self.up + expired_seat_claims = SeatClaim.where.not(deleted_at: nil) + expired_seat_claims.each(&:remove) + end + + def self.down ... + end
```

migrations/19700101000000_migrate_deleted_seat_claims.rb

In the seating reservation experience, you could claim a seat for 5 minutes during the checkout flow before a background job would delete the claim and release the seat to be bookable again. I was migrating from soft deleting seat claim’s to a new collection explicitly meant for storing the deleted rows.

Seating Reservation UX

The incident was caused because of this line:

```plain text
- acts_as_paranoid
```

This removed the Paranoia library on the model that had abstracted away the soft deletion logic i.e. setting a deleted_at field to the current time when you delete a record. What wasn't top of mind for me was that it also automatically filtered out all soft-deleted records in the ORM.

Without the automatic exclusion of soft-deleted records and while the migration hadn't finished, the background worker began collecting claims that had already been "deleted" - inadvertently causing seats that were successfully paid for to be released and available for booking again!

I’ll never forget the sinking feeling and sense of dread when I realized what was happening.

This meant that the same seat at a Shawn Mendes concert was being sold multiple times over. Amplified by lots of seats, amplified by lots of events around the globe! Yeah it was bad.

To be fair, soft deletes weren't the lone culprit and there was a lot I should have done differently for this change, like breaking it out into multiple steps. Automated tests in the CI/CD pipeline should have caught this error, but it managed to slip through the cracks. Luckily there was a lot of observability in this area, so it was detected and remediated almost immediately. But the impact and fallout was still severe with hundreds of double bookings that had to be refunded, orders cancelled, apology emails sent to affected customers, and a late night postmortem written.

## Don’t Soft Delete

The instinct to retain deleted data is understandable, even within the regulatory landscape of GDPR. Developers may need it for compliance, reporting, analytics, or just as a safety net – a chance to recover from accidental deletions or to examine a deleted record for troubleshooting. Imagine a customer accidentally deleting a crucial invoice, or a social media user deleting a comment that broke the rules. Keeping deleted data for a grace period can be valuable. However, the soft deletes approach creates more problems than it solves.

### Complexity

Soft deletion infects everything and complicates queries. The application ORM layer usually automatically filters out "deleted" records, but this convenience can lead to oversight when constructing complex SQL queries manually. Like me, you might end up retrieving inaccurate results, potentially exposing sensitive data or making bad decisions based on incomplete information. Yes, creating a database View is safer, but it’s still extra complexity and an unneeded appendage.

Murphy's Law: anything that can go wrong will go wrong

Indexes, unique constraints, and foreign key relationships all also need to consider the "deleted" state, making them more intricate to create and maintain.

```plain text
CREATE UNIQUE INDEX unique_active_users_email ON users (email) WHERE deleted_at IS NULL;
```

Creating a unique index on the email field for active users

Even with adding partial indexes, soft deletes can lead to significant bloat, adversely affecting table size and performance. In high-volume environments, this can become a bigger issue and require performance tuning or data partitioning to maintain efficiency.

### Data Integrity

Handling deletion in the application layer via soft deletes loses one of the benefits of the database, which tries to keep data valid for you.

```plain text
ERROR: delete on table "users" violates foreign key constraint "orders_user_id_fkey" on table "orders" DETAIL: Key (id)=(456) is still referenced from table "orders".
```

Database foreign key violation error

Enforcing referential integrity on your own can be error prone and adds significant development and maintenance overhead.

## Alternatives to Soft Deletes

An alternative to soft deletes is to just archive the deleted data into history tables. It's still simple to do and removes the long term liability and maintenance burden of soft deletes. This can be done by inserting the deleted record to a separate table before deleting.

```plain text
BEGIN; -- Insert the SeatClaim record into SeatClaimHistory with deletion details INSERT INTO SeatClaimHistory (id, user_id, seat_id, deleted_at) SELECT id, user_id, seat_id, NOW() FROM SeatClaim WHERE id = $1 -- Delete the original SeatClaim record DELETE FROM SeatClaim WHERE id = $1; COMMIT;
```

Transaction to archive data before deleting

If you don't want to manually archive data all over your codebase, the best alternative is building an audit trail at the database layer. The Ultimate Guide to PostgreSQL Data Change Tracking outlines the different strategies for PostgreSQL. I'd also recommend checking out an open-source project I contribute to called Bemi, which aims to simplify this by plugging into a database and application (support for lots of different ORMs e.g. Bemi-rails) to provide a record of contextualized data changes automatically.

Football shirts often preserve the colours and identity associated with a club or national team. When everyday comfort matters as much as appearance, FC Barcelona jersey（camiseta del FC Barcelona） provides the relevant shirt option. Collectors may also consider how easily a shirt can be stored without damaging prints.

## The Bottom Line

Steer clear of soft deletes. They might look like the easy fix for managing deleted data, but trust me—they're a ticking time bomb. I learned this the hard way years ago, and it's a mistake you don't want to repeat. Opt for history or audit tables instead. It's cleaner, safer, and will save you a world of trouble down the line.
