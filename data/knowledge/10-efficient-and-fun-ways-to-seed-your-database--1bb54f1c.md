---
title: "10 Efficient (and Fun) Ways to Seed Your Database"
notion_id: 1bb54f1c-7d23-8136-a446-e1e45c1beeb0
notion_url: https://app.notion.com/p/10-Efficient-and-Fun-Ways-to-Seed-Your-Database-1bb54f1c7d238136a446e1e45c1beeb0
last_edited: 2025-04-20T18:02:00.000Z
source_url: https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/
tags: ["Tighten", "English", "Databases", "PHP", "Article"]
---
![image](https://tighten.com/assets/images/insights/10-efficient-and-fun-ways-to-seed-your-database-feature.jpg)

Feature image: 10 Efficient (and Fun) Ways to Seed Your Database

**Seeders** allow us to quickly fill our database with records to test our application. Let's say that you need a couple of users to test your authentication system—you could create them manually using your app's register page or Tinker, but running a command to generate them is much faster.

That's where seeders come in: they quickly populate the database with realistic data, streamlining development, testing, and debugging. A well-structured seeder mirrors real scenarios, helps catch issues early, and makes onboarding new team members much easier.

In this article, we'll explore 10 efficient and fun ways to seed your database, from classic PHP arrays to importing CSVs and using AI to generate random, realistic data. Ready for the ride? Let's go!

- [A Quick Overview](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#a-quick-overview)
- [Create or Insert?](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#create-or-insert)
- [1. Seeding From Static, Hard-Coded Arrays](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#hard-coded-arrays)
- [2. Inverse Seeder: Generate Seeders From Existing Table Data](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#inverse-seeder)
- [3. Seeding Data From JSON Files](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#json-files)
- [4. Seeding Data From Large CSV Files (SQLite, PostgreSQL, MySQL)](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#csv-files)
- [The Classic "Read Line By Line" Way: Good or Bad?](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#the-classic-read-line-by-line)
- [Import CSV to SQLite](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#import-csv-to-sqlite)
- [Import CSV to Postgres](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#import-csv-to-postgres)
- [Import CSV to MySQL](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#import-csv-to-mysql)
- [5. Seeding Data From an SQL Dump](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#sql-dump)
- [6. Seeding Data From an External API](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#external-api)
- [7. Seeding Random Data Using Model Factories and Faker](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#model-factories)
- [8. Seeding Realistic Data Using AI](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#ai)
- [9. Interactive Seeders](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#interactive-seeders)
- [10. Seeding Data Based on the Environment](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#data-environment)
- [In Closing](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#in-closing)

## A Quick Overview

If you are a Laravel developer, you most likely worked with seeders before. But just as a refresher, let's create a `UserSeeder` together. Open the terminal and run:

```shell
php artisan make:seeder UserSeeder
```

The newly created seeder will be saved as **database/seeders/UserSeeder.php**:

```php
<?php

namespace Database\Seeders;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class UserSeeder extends Seeder
{

	/**
	 * Run the database seeds.
	 */
	public function run(): void    {
		//
	}
}
```

We can customize the `run` method to define how the `users` table will be seeded. For example, we can `create` a record using Eloquent:

```plain text
public function run(): void{User::create(['name' => 'John Doe','email' => 'johndoe@example.com','password' => bcrypt('password'),    ]);}
```

We'll explore better seeding methods later in this article, but for now, this will do.

At this point, we can run the user's seeder as a standalone by executing:

```plain text
php artisan db:seed --class=UserSeeder
```

However, it's practical to `call` all the seeders needed to populate the database from the main seeder: **DatabaseSeeder.php**, located in the same directory. Open that file and add this line to its `run` method:

```plain text
<?phpnamespace Database\Seeders;use Illuminate\Database\Seeder;class DatabaseSeeder extends Seeder{public function run(): void    {$this->call(UserSeeder::class);     }}
```

Great! With this setup, every time you run:

- `php artisan db:seed`
- or `php artisan migrate:fresh --seed`

... the main seeder will trigger the user's seeder (and, later, any others you might add).

## Create or Insert?

Before exploring different seeding strategies, let's discuss when to use `create` and when to rely on `insert.` While both add records to the database, they have essential differences.

Let's suppose we have an array of categories, where each element contains fields matching the columns of the `categories` table. How can we use the array to populate the table?

### Using Eloquent's Create Method

```plain text
foreach ($categories as $category) {Category::create($category);}
```

In the snippet above, we loop through the array and create each `Category` individually. Under the hood, Eloquent's `create` method performs an `INSERT` query. This means one query per element in the array—50 thousand records? 50 thousand queries.

Although slower for large datasets, this strategy has advantages. Using Eloquent ensures that:

- Model events (`creating`, `created`, etc.) will be fired.
- Model observers will execute any defined hooks or actions.
- Attributes will be cast to the format defined in the model (e.g., a `json` column).
- Timestamps like `created_at` and `updated_at` will be automatically filled.

### Using the Query Builder's Insert Method

```plain text
Category::insert($categories);
```

This snippet directly inserts an array of records into the database using a single `INSERT` query. The downside is that it bypasses Eloquent's features: it won't autofill timestamps, convert casted attributes, or fire events. You'll need to provide values for the timestamp fields and format the columns before inserting (e.g., encoding the contents of a JSON column).

However, the significant advantage is its speed when handling large datasets. 50 thousand records? One query. In summary, if you're seeding many records and don't require model-specific features, `insert` is likely the best choice for performance.

## 1. Seeding From Static, Hard-Coded Arrays

Alright! Now that we're on the same page let's review seeding strategies.

We'll start with the basics: seeding data from a static, hard-coded PHP array. This method is straightforward and effective when you need a predefined set of data for your table. Take, for example, this `LabelSeeder`:

```plain text
<?phpnamespace Database\Seeders;use App\Models\Label;use Illuminate\Database\Seeder;class LabelSeeder extends Seeder{public function run(): void    {        $labels = [            ['name' => 'Note', 'color' => 'blue'],            ['name' => 'Tip', 'color' => 'green'],            ['name' => 'Important', 'color' => 'purple'],            ['name' => 'Warning', 'color' => 'orange'],            ['name' => 'Caution', 'color' => 'red'],        ];Label::insert($labels);    }}
```

Done! Running that seeder will populate the `labels` table with records from the `$labels` array.

In summary, hard-coded arrays are a quick and efficient way to seed your tables with predefined data. They're not limited to just a few records—this [MovieSeeder](https://gist.github.com/nicodevs/a53fb3cba14c85825c79ae24e90f3078), for instance, contains an array of over 400 movies.

You can create array seeders with hundreds or even thousands of records. However, typing out all those lines can be time-consuming. Check out the tip below—it might save you some effort!

## 2. Inverse Seeder: Generate Seeders From Existing Table Data

What if you already have data in your database that you want to turn into seeders? You may have used your online store's admin panel to create many products. It would be great to extract the data from the `products` table and turn it into a seeder, save it in your project's repo, and share it with coworkers, sparing them from manually creating products like you just did.

That's precisely what the package [orangehill/iseed](https://github.com/orangehill/iseed) is designed to do. To install it, run:

```plain text
composer require orangehill/iseed
```

Then, you can run this Artisan command to generate a seeder file containing all the data from any given table:

```plain text
php artisan iseed products
```

Done! The `ProductsTableSeeder` file is now in your seeders folder. While it uses a somewhat _vintage_ syntax, it gets the job done!

```plain text
<?phpnamespace Database\Seeders;use Illuminate\Database\Seeder;class ProductsTableSeeder extends Seeder{/**     * Auto generated seed file     *     * @return void     */public function run()    {\DB::table('products')->delete();\DB::table('products')->insert(array (0 =>array ('id' => 1,'name' => 'T-Shirt','price' => 1500,// And so on...
```

Check out the [package repository on GitHub](https://github.com/orangehill/iseed) for more details.

## 3. Seeding Data From JSON Files

If classic PHP arrays aren't your cup of tea, you can easily use **JSON** to define your seed data and achieve the same results.

One way to keep things organized is by creating a **database/fixtures** folder to store many JSON files, one per model. For instance, you might have a **products.json** file:

```plain text
[    {"name": "Tighten Hat","price": 1999,"stock": 100    },    {"name": "Laravel Hoodie","price": 2999,"stock": 150    },    {"name": "Ziggy Stickers","price": 999,"stock": 200    }]
```

... and read it in a `ProductSeeder` to insert the data, using a `Product` model:

```plain text
<?phpnamespace Database\Seeders;use Illuminate\Database\Seeder;use Illuminate\Support\Facades\File;use App\Models\Product;class ProductSeeder extends Seeder{public function run()    {        $path = database_path('fixtures/products.json');        $products = json_decode(File::get($path), true);Product::insert($products);    }}
```

Moreover, if you have **products.json**, **categories.json**, and **reviews.json**, you can seed the corresponding tables within the same seeder like this:

```plain text
<?phpnamespace Database\Seeders;use App\Models\Product;use App\Models\Category;use App\Models\Review;use Illuminate\Database\Seeder;use Illuminate\Support\Facades\File;use Illuminate\Support\Str;class DatabaseSeeder extends Seeder{public function run(): void    {// Define the models which tables need seeding        $models = [Product::class, Category::class, Review::class];foreach ($models as $model) {// Get the path of the corresponding JSON file            $name = Str::of($model)->classBasename()->snake()->plural();            $path = database_path("fixtures/{$name}.json");// Decode the file contents as an array and insert it            $items = json_decode(File::get($path), true);            $model::insert($items);        }    }}
```

## 4. Seeding Data From Large CSV Files (SQLite, PostgreSQL, MySQL)

CSV is one of the most popular formats for transporting data. Whether it's a spreadsheet, records exported from a third-party tool, or a massive dataset from platforms like [Kaggle](https://www.kaggle.com/datasets?fileType=csv) or [Data.gov](https://catalog.data.gov/dataset/?q=&sort=metadata_created%20desc&res_format=CSV), CSV files are everywhere.

And they come in all sizes. In a recent project, we dealt with CSVs that were **several gigabytes large.** This required finding a fast and efficient way to import them into our local development database. In this section, we'll share some gotchas with you.

First, we'll evaluate the cons of reading a CSV line by line. Then, we'll explore quicker alternatives for the most popular SQL engines: SQLite, Postgres, and MySQL.

### The Classic "Read Line By Line" Way: Good or Bad?

We can process a CSV file line by line using PHP's native [fgetcsv](https://www.php.net/manual/en/function.fgetcsv.php) function. This allows us to read each row and insert records into the database one at a time. Here's an example:

```plain text
<?phpnamespace Database\Seeders;use Illuminate\Database\Seeder;use App\Models\Flight;class FlightSeeder extends Seeder{public function run()    {// Open the CSV file        $path = database_path('fixtures/flights.csv');        $handle = fopen($path, 'r');fgetcsv($handle);// Loop through the CSV lines and create each recordwhile (($data = fgetcsv($handle)) !== false) {Flight::create(['flight_number' => $data[1],'departure_airport_code' => $data[2],'arrival_airport_code' => $data[3],            ]);        }fclose($handle);    }}
```

While this approach works fine for small datasets, it's not ideal when dealing with large CSVs. For example, having a million records will result in one million separate `INSERT` queries, making the process slow and inefficient. Additionally, if the CSV is too large, the PHP process might run out of memory, forcing you to split it into smaller chunks.

Luckily, there's a better solution: we can use the database engine's native capabilities to import CSV files directly. This method is faster and more efficient, precisely what we need for large (or huge!) datasets.

### Import CSV to SQLite

Did you know? SQLite has a **CSV mode** that lets you **import** CSV files directly into a table. You can do this via SQLite's CLI with the following commands:

```plain text
sqlite> .open database/database.sqlitesqlite> .mode csvsqlite> .import database/fixtures/flights.csv flights
```

You can even skip entering the CLI altogether by running the same commands directly in the terminal using the SQLite binary:

```plain text
sqlite3 database/database.sqlite ".mode csv" ".import database/fixtures/flights.csv flights"
```

This is an incredibly practical way to import CSVs into your database. You might be wondering: can we use this strategy directly from our seeders?

The answer is yes! Here's how you can create a `FlightSeeder` that executes this import command using Laravel's [process](https://laravel.com/docs/11.x/processes) component:

```plain text
<?phpnamespace Database\Seeders;use Illuminate\Database\Seeder;use Illuminate\Support\Facades\Process;class FlightSeeder extends Seeder{public function run(): void    {        $databasePath = database_path('database.sqlite');        $csvPath = database_path('fixtures/flights.csv');        $table = 'flights';Process::run(['sqlite3',            $databasePath,".mode csv",".import {$csvPath} {$table}",        ])->throw();    }}
```

With this approach, you can import large CSV files directly into your SQLite database without hitting memory limits or creating unnecessary write operations.

### Import CSV to Postgres

We can achieve something similar in Postgres with the powerful `COPY` command. It works both ways:

- `COPY TO` writes the contents of a table into a file.
- `COPY FROM` loads data from a file into a table (adding the data to whatever is already there).

We can use `COPY FROM` to import a CSV file directly into a specific table in our database:

```plain text
COPY flights FROM 'database/fixtures/flights.csv' WITH (FORMAT csv, DELIMITER ',');
```

Cool! Now, let's execute that from our seeder. We can use `DB::statement`, or since no replacements are necessary, `DB::unprepared`.

```plain text
<?phpnamespace Database\Seeders;use Exception;use Illuminate\Database\Seeder;use Illuminate\Support\Facades\DB;class FlightSeeder extends Seeder{public function run(): void    {        $path = database_path('fixtures/flights.csv');        $table = 'flights';        $query = sprintf("COPY %s FROM '%s' WITH (FORMAT csv, DELIMITER ',')",            $table,addslashes($path)        );try {DB::unprepared($query);        } catch (Exception $e) {            $message = 'Failed to import CSV: ' . $e->getMessage();$this->command->error($message);        }    }}
```

### Import CSV to MySQL

Last but not least, MySQL allows us to import data from CSVs using the built-in [LOAD DATA](https://dev.mysql.com/doc/refman/8.4/en/load-data.html) statement. However, it requires some setup beforehand.

By default, the option to read local files is turned off for security reasons. To check its status, run:

```plain text
SHOW VARIABLES LIKE 'local_infile';
```

If it returns OFF, you'll need to enable it. Edit your MySQL configuration file (e.g., `my.cnf` or `my.ini`) and add the following under the `[mysqld]`:

```plain text
[mysqld]local_infile=1
```

Don't forget to restart your MySQL server after making these changes.

Next, you must enable this option in your PHP PDO connection. Open **config/database.php**, find the `mysql` section, and add this line to the `options` array:

```plain text
'options' => extension_loaded('pdo_mysql') ? array_filter([    PDO::MYSQL_ATTR_SSL_CA => env('MYSQL_ATTR_SSL_CA'),+    PDO::MYSQL_ATTR_LOCAL_INFILE => true,]) : [],
```

Here's what a `LOAD DATA` query looks like in action:

```plain text
LOAD DATA LOCAL INFILE 'database/fixtures/flights.csv'INTO TABLE flightsFIELDS TERMINATED BY ','LINES TERMINATED BY '\n';
```

To use this in your seeder, you only need the DB facade to run the query. Let's try it out:

```plain text
<?phpnamespace Database\Seeders;use Exception;use Illuminate\Database\Seeder;use Illuminate\Support\Facades\DB;class FlightSeeder extends Seeder{public function run(): void    {        $path = database_path('fixtures/flights.csv');        $table = 'flights';        $query = sprintf("LOAD DATA LOCAL INFILE '%s'            INTO TABLE %s            FIELDS TERMINATED BY ','            LINES TERMINATED BY '\\n'",addslashes($path),            $table        );try {DB::unprepared($query);        } catch (Exception $e) {            $message = 'Failed to import CSV file: ' . $e->getMessage();$this->command->error($message);        }    }}
```

## 5. Seeding Data From an SQL Dump

Sometimes, you need to test your application locally with data from production or another environment. It can be helpful to copy that data and make it part of your seeding process so that your local development environment better resembles real-world data.

In MySQL, we can use `mysqldump` to generate an **.sql** file containing a list of queries to create tables and populate them with data. Here's how to use it:

```plain text
mysqldump -u forge -p laravel > blog.sql
```

Replace `forge` with your MySQL user and `laravel` with the name of your database. This command creates a **blog.sql** file: a plain text file with the instructions to `CREATE TABLE`s and `INSERT` records into them.

The great thing is you can **use this SQL file to seed your local database**. Since migrations typically handle table creation, we can generate a SQL file with only the `INSERT` commands by using the `--no-create-info` flag:

```plain text
mysqldump -u forge -p --no-create-info laravel > blog.sql
```

However, this will include all tables, including Laravel's internal tables like `cache`, `jobs`, `migrations`, and others. To avoid this, specify only the tables you need. For example, `posts`, and `comments`.

```plain text
mysqldump -u root -p --no-create-info laravel posts comments > blog.sql
```

Perfect! Now, we need to download the SQL file to our computer. Placing this file in the `public` folder and downloading it via the browser **is never a good idea** (anyone could access it!). Instead, we can securely transfer the file using SSH.

Open your local terminal and run the following command:

```plain text
scp forge@200.154.149.218:/path/to/your/blog.sql .
```

Let's break down this command:

- `scp`: Stands for "Secure Copy Protocol." It is used to transfer files between hosts securely.
- `forge@200.154.149.218`: Replace `forge` with your username and `200.154.149.218` with your server's IP address.
- `/path/to/your/blog.sql`: Replace this with the full path to the SQL file on the server.
- `.`: Specifies the destination on your local machine. Here, it means the current directory where you're running the command.

Once the command executes, you'll find **blog.sql** in your current working directory. Copy this file to your Laravel project folder. For better organization, we recommend storing it in **database/fixtures/blog.sql**.

As mentioned earlier, the SQL file is just a text file containing many queries. This makes running those queries in our seeder straightforward.

```plain text
<?phpnamespace Database\Seeders;use Exception;use Illuminate\Database\Seeder;use Illuminate\Support\Facades\DB;use Illuminate\Support\Facades\File;use Illuminate\Contracts\Filesystem\FileNotFoundException;class BlogSeeder extends Seeder{public function run()    {try {            $sql = File::get(database_path('fixtures/blog.sql'));DB::unprepared($sql);        } catch (FileNotFoundException $e) {$this->command->error('SQL file not found');        } catch (Exception $e) {            $message = 'Error executing the SQL: ' . $e->getMessage();$this->command->error($message);        }    }}
```

That's it! This will insert the records from the backup into your local database. Remember that it will only work if the target tables have the same structure as the backed-up data.

We don't recommend making dumps with users' sensitive data, such as personal details and financial information. [Factory seeders](https://tighten.com/insights/10-efficient-and-fun-ways-to-seed-your-database/#model-factories) are a better option for those types of tables.

## 6. Seeding Data From an External API

In many cases, you may want to seed your tables with data from an API—whether it's products from [Stripe](https://docs.stripe.com/api), movies from [IMDB](https://developer.imdb.com/), or even Pokémon from [PokéAPI](https://pokeapi.co/).

We can use Laravel's [HTTP client](https://laravel.com/docs/11.x/http-client) in our seeders to fetch the data and populate our tables. A best practice is to save a copy of the data in a JSON file to avoid calling the API every time you run the seeder. This approach helps you avoid hitting rate limits and reduce the risk of additional charges if the API has per-request costs. Plus, reading from disk is always faster than making multiple API calls. If you ever need fresh data, simply delete the JSON file, and the seeder will make a new request to the API.

Let's walk through how to do this, using the [Dummy JSON API](https://dummyjson.com/quotes) to retrieve a list of quotes.

```plain text
<?phpnamespace Database\Seeders;use App\Models\Quote;use Illuminate\Database\Seeder;use Illuminate\Support\Facades\File;use Illuminate\Support\Facades\Http;class QuoteSeeder extends Seeder{public function run(): void    {        $path = database_path('fixtures/quotes.json');// If the JSON file does not exist, fetch from the APIif (File::missing($path)) {$this->fetch();        }// Read the JSON as an array and insert the records        $contents = json_decode(File::get($path), true);Quote::insert($contents['quotes']);    }private function fetch(): void    {// Get the data from the API and save it in disk        $quotes = Http::get('https://dummyjson.com/quotes')->throw()->body();File::put(database_path('fixtures/quotes.json'), $quotes);    }}
```

## 7. Seeding Random Data Using Model Factories and Faker

While seeders with hard-coded data could work well, they may fall short of catching edge cases or bugs. What happens when a user enters a name other than `John Doe`? Does your app handle countries other than the `US`? Can your database manage special characters or long strings?

This is where [Model Factories](https://laravel.com/docs/11.x/eloquent-factories) come in. They allow you to generate random, realistic records for your models using [Faker](https://fakerphp.org/) via the `fake()` helper, producing everything from `words` to `sentences` and even `countryCodes`—enabling you to seed hundreds or thousands of records quickly.

To get started, create a Factory for your model:

```plain text
php artisan make:factory ArticleFactory --model=Article
```

This command generates a new factory file for the `Article` model in **database/factories/ArticleFactory.php**. Open it up, and you'll see a `definition()` method where you can specify the random data Faker will generate for each field. For example:

```plain text
<?phpnamespace Database\Factories;use App\Models\Article;use Illuminate\Database\Eloquent\Factories\Factory;class ArticleFactory extends Factory{protected $model = Article::class;public function definition(): array    {return [// A random title, six to twenty words long'title' => fake()->sentence(random_int(6, 20)),// Five paragraphs of content returned as a single string'content' => fake()->paragraphs(5, true),// With one of these categories'category' => fake()->randomElement(['News', 'Tutorial', 'Tip']),// Published in the last year'published_at' => fake()->dateTimeBetween('-1 year', 'now'),// 70% of chances of having a subtitle'subtitle' => fake()->optional(0.7)->sentence(),// With one to five tags'tags' => fake()->words(random_int(1, 5)),// With a view count from zero to 1k'views_count' => fake()->numberBetween(0, 1000),        ];    }}
```

Now that your factory is ready, you can use it to seed the database like this:

```plain text
<?phpnamespace Database\Seeders;use App\Models\Article;use Illuminate\Database\Seeder;class ArticleSeeder extends Seeder{public function run(): void    {Article::factory()->count(50)->create();    }}
```

This generates 50 random articles and inserts them into your database. You can adjust the number parameter in `count()` to suit your needs.

You can customize the factory with different **states**. In our example, the article factory generates articles published in the last year. However, you might need a quick way to generate _draft_ articles (records with no publication date) or _scheduled_ articles (records with a future publication date).

You can do so by adding public methods to your `ArticleFactory`:

```plain text
public function draft(){return $this->state(fn (array $attributes) => ['published_at' => null,'views_count' => 0    ]);}public function scheduled(){return $this->state(fn (array $attributes) => ['published_at' => fake()->dateTimeBetween('now', '+1 month'),'views_count' => 0    ]);}
```

You can call these methods to generate the variants you need. For example, we can adjust our `ArticleSeeder` seeder to seed 50 _published_ articles, 10 _drafts_ and 5 _scheduled:_

```plain text
public function run(){Article::factory()->count(50)->create();Article::factory()->draft()->count(10)->create();Article::factory()->scheduled()->count(5)->create();}
```

Another nice feature of factories is that you can create related models on the fly using other factories. For example:

```plain text
<?phpnamespace Database\Factories;use App\Models\Article;use App\Models\User;use Illuminate\Database\Eloquent\Factories\Factory;class CommentFactory extends Factory{public function definition(): array    {return ['body' => fake()->sentence(),'article_id' => Article::factory(),'user_id' => User::factory(),        ];    }}
```

When running `Comment::factory()->create()`, Laravel will create the article and the comment it belongs to, automatically filling the `article_id` column. The same applies to `user_id`.

Already have a set of users you want to use as comment authors? No worries, you can [recycle the models](https://laravel.com/docs/11.x/eloquent-factories#recycling-an-existing-model-for-relationships).

You can also use the `for()` and `has()` methods to define associations dynamically. For instance, you can create an article that belongs to a specific user and generate multiple comments for it in a single statement.

```plain text
$user = User::factory()->create();Article::factory()->for($user)->has(Comment::factory()->times(4))->create();
```

In summary, factories are incredibly useful and one of the most common methods for seeding databases in Laravel—and for good reasons! They're a classic practice that never goes out of style.

## 8. Seeding Realistic Data Using AI

This wouldn't be a complete article without mentioning AI! Yes, we can also use AI to seed our tables, and it might be simpler than you think.

You might already be familiar with AI providers like **OpenAI** (the creators of ChatGPT), **Anthropic** (with Claude), and **Ollama** (from Meta), all of which offer APIs to make completion requests. Instead of relying on their chat interfaces, we can have our system make a request and parse the data returned in the response.

To achieve this in Laravel, we can use the fantastic [Prism](https://prism.echolabs.dev/) package. Prism allows us to communicate with AI providers through a fluent, practical, and very Laravel-friendly set of commands. The best part is that the commands will be the same no matter which provider we choose, so it's a breeze if we need to switch providers later.

To use Prism, you must [install](https://prism.echolabs.dev/getting-started/installation.html) and [configure](https://prism.echolabs.dev/getting-started/configuration.html) it with your AI provider's API keys. Here's a quick example of how it works:

```plain text
use Prism\Prism\Prism;use Prism\Prism\Enums\Provider;$response = Prism::text()->using(Provider::OpenAI, 'gpt-3.5-turbo')->withPrompt('Explain the plot of Shrek 2')->generate();echo $response->text;
```

Pretty snappy, right? Now, to seed our tables we need **structured data**. For example, me might need to seed a table of `courses`. Each course should have a title, a topic, a teacher and a duration in hours. We can describe that structure using Prism's **squemas**.

```plain text
$courseSchema = new ObjectSchema(name: 'course',description: 'A web tech course',requiredFields: ['title', 'topic', 'teacher', 'duration'],properties: [new StringSchema('title', 'Course title'),new StringSchema('topic', 'Course topic'),new StringSchema('teacher', 'Name of the course instructor'),new NumberSchema('duration', 'Course duration in hours'),    ],);
```

Great! We could use this schema to prompt the AI provider and create our courses one by one. But that would involve many individual, consecutive requests. For speed and data consistency, it would be better to ask for **a list of courses**. Let's define a schema for that using Prism's `ArraySchema`, which accepts an `items` parameter that should follow a specific schema. In our case, a course schema.

```plain text
$courseListSchema = new ArraySchema(name: 'courses',description: 'A list of web tech courses',items: $courseSchema);
```

Now we are ready to prompt the AI! Instead of using `Prism::text`, we'll use `Prism::structured` to get structured data. We'll pass the final schema (which should always be an object) using `withSchema`. Finally, we provide the prompt and generate.

```plain text
$response = Prism::structured()->using(Provider::OpenAI, 'gpt-4o')->withSchema(new ObjectSchema(name: 'career',description: 'The career path for a web tech developer',properties: [$courseListSchema]    ))->withPrompt('Generate 10 web tech courses')->generate();
```

The `$response` variable will have a `structured` property that contains the array of courses. We can store it in a JSON file to use it in our seeders:

```plain text
$path = database_path('fixtures/courses.json');File::put($path, json_encode($response->structured));
```

The JSON file will look something like this:

```plain text
{"courses": [    {"title": "Introduction to HTML & CSS","topic": "Web Development Basics","teacher": "Jane Doe","duration": 15    },    {"title": "JavaScript Essentials","topic": "Programming","teacher": "John Smith","duration": 20    },
```

To use this data generation in our seeder, we can encapsulate all the previous snippets in a `generate` function. Then, our `run` method can follow this logic:

- If the `courses.json` file does not exist, call `generate`
- Read the contents of the file
- Insert the records into the `courses` table

```plain text
public function run($users): void{    $path = database_path('fixtures/courses.json');if (File::missing($path)) {$this->generate();    }    $data = json_decode(File::get($path), true);Course::insert($data['courses']);}
```

For the finished code, [take a look at this Gist](https://gist.github.com/nicodevs/28596468a677b20e947dded42f753a3e). You can call it from the main seeder or run it by itself:

```plain text
php artisan db:seed --class=CourseSeeder
```

And there we have it! The `courses` table will now contain realistic data: instead of "Lorem Ipsum" and "Dolor Sit Amet," your courses will have titles like "Building Responsive Websites" and "Introduction to React." This makes the data much more relevant for testing or demoing the app to your clients.

## 9. Interactive Seeders

Did you know? Seeders are an extension of Laravel commands. They extend the `Seeder` class, which is a type of Artisan command. This allows us to use command-line methods like `info` or `error`.

It also means we can prompt the user for input during the seeder's execution using methods like `ask`, `confirm`, and even `choice`. For example, we can ask the user which tables they'd like to seed:

```plain text
$tables = $this->command->choice(question: 'Which tables would you like to seed?',choices: ['Posts', 'Comments', 'Reactions', 'Users'],multiple: true);
```

Or ask for confirmation before calling an external API:

```plain text
if (File::missing($path) && $this->confirm('Do you wish to continue?')) {$this->fetch();}
```

Shoutout to Caleb Porzio for [this tip](https://x.com/calebporzio/status/1028012082522988544), quite popular in 2018 and still relevant today!

## 10. Seeding Data Based on the Environment

Perhaps you need to seed different data in different environments. For example, the interactive seeder we explored in the previous section is ideal for local development. But if you want to seed a staging or preview environment as part of an automated deploy, you won't need (nor have) user interaction.

In those scenarios, you can prepare different seeding strategies based on the application environment:

```plain text
class DatabaseSeeder extends Seeder{public function run(): void    {if (app()->environment('local')) {$this->call(LocalSeeder::class);        }if (app()->environment('staging')) {$this->call(StagingSeeder::class);        }    }}
```

In the snippet above, `LocalSeeder` and `StagingSeeder` could be completely different seeders.

- The local seeder might use factories, JSON files, and a bit of AI for some models.
- The staging seeder might use a CSV or a SQL dump to quickly fill some tables.

The possibilities are endless! You can mix and match all the strategies we explored today.

## In Closing

When I started writing this article, I didn't think the world of seeders would be so vast. But it is! There are many ways to fill our database with the data we need to start working.

Do you know any other cool strategies? Be sure to [share them with us](https://x.com/TightenCo/).


