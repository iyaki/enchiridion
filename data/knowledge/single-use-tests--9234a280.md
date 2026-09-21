---
title: "Single-use Tests"
notion_id: 9234a280-bd6f-4676-a366-71c66eed0d19
notion_url: https://app.notion.com/p/Single-use-Tests-9234a280bd6f4676a36671c66eed0d19
last_edited: 2023-10-17T12:29:00.000Z
source_url: https://timacdonald.me/single-use-tests/
tags: ["Article", "Tim MacDonald", "English", "Programming", "Testing"]
---
The tests I write are usually written in Pest, PHPUnit, or Jest. I also usually commit them to a code repository - but not all automated tests should be committed. I wrote what I consider to be an automated test that was a "single-use test". It served its purpose and it will forever be in my heart, but not in my repository.

A tiny script I created allowed me to verify, with certainty, that I had not modified any route definitions after a pretty decent-sized refactor that touched a lot of different parts of the application. I thought it was pretty neat and gave me a lot of confidence to ship the changes, so I thought I would share.

In a project I’m working on there are duplicated middleware added to a lot of routes. At some point, in this projects long life, the way middleware was added to routes was moved from controller constructor middleware to the route file middleware (where they belong 😎).

If you aren’t familiar with these two options for adding middleware to a route, Laravel allows you to add middleware in a routes file...

```plain text
// routes/web.php

use App\Http\Middleware\Authenticate;
use App\Http\Middleware\MonitorUsage;
use Illuminate\Support\Facades\Route;

// Apply middleware to a group of routes...
Route::middleware(Authenticate::class)->group(function () {

    // Apply middleware to a single route...
    Route::get(/* ... */)->middleware(MonitorUsage::class);
});

```

Laravel also allows you to add middleware to routes via the controller’s constructor. These middleware will apply to all methods on the controller unless otherwise configured.

```plain text
namespace App\Http\Controllers;

use App\Http\Middleware\Authenticate;
use App\Http\Middleware\MonitorUsage;

class UserController extends Controller
{
    public function __construct()
    {
        $this->middleware(Authenticate::class);
        $this->middleware(MonitorUsage::class, ['only' => 'show']);
    }

    // ...
}

```

As I mentioned, this project had moved from using controller middleware to route file middleware, however the middleware in the controllers had not yet been tidied up and removed.

This wasn’t really a problem.

It wasn’t like these middleware were being run twice. Laravel deduplicates middleware applied to a route, i.e., if we add the same middleware to a route more than once it is only executed once.

We can see this with a quick example. I'll create an empty middleware that dumps a message when it is run.

```plain text
namespace App\Http\Middleware;

class MyMiddleware
{
    public function handle($request, $next)
    {
        dump('Running my middleware.');

        return $next($request);
    }
}

```

I can then create a route and apply the middleware to it twice; once via a wrapping group and once on the actual route itself.

```plain text
// routes/web.php

use App\Middleware\MyMiddleware;
use Illuminate\Support\Facades\Route;

Route::middleware(MyMiddleware::class)->group(function () {
    Route::get('/', fn () => 'Welcome!')->middleware(MyMiddleware::class);
});

```

When we access this page we see the following output rendered in the browser. You will notice that the dump is only seen once, indicating that Laravel has removed the duplicate middleware from the stack.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

So cleaning this up was about aesthetics and consistency - or so I thought!

I decided to finally dig in and refactor it away. I came up with a plan and started doing a, rather manual, systematic route-by-route check.

1. Look at the route.
2. Work out the applied middleware in the routes file.
3. Move to the controller constructor.
4. Verify the middleware applied matched or was at least a subset of the route file middleware.
5. Remove the middleware from the constructor.
6. Make any required adjustments in the route file.

I did this for about 2 minutes before it dawned on me that whoever reviews this is gonna have a bad time. We would either YOLO it into production and hope the testsuite caught any issues or some poor soul would have to do all this manual checking again while reviewing. Not a fun time.

It wasn’t worth the time / improvement trade off.

I stopped coding; I moved away from the keyboard; I thought…

Do I scrap the changes? It’s not hurting anyone…but it did hurt _me_. Won’t somebody think of the children!

Then I realised I could _potentially_ automate the review process and MAKE THE ROBOTS DO IT 🤖

Laravel ships with an artisan command that displays all routes defined within the application. There are a few flags on the command to change the commands output.

The bare command, `php artisan route:list`, will show you the routes.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

But this doesn’t show the route middleware. Using the verbose flag, `php artisan route:list -v`, will show middleware for each route.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Great! We are getting somewhere useful! I then remembered that a `--json` flag was available on the command to output a JSON representation of the routes and their middleware instead of the visual representation we see in the terminal.

With the combination of all of this I could get all the routes and their respective middleware output before I made the change and then compare the file to the routes after I made the change.

So no more route-by-route review; I went straight for the jugular and stripped out all the middleware from all the controllers in a matter of minutes.

I committed my changes to my `working` branch.

I crafted the following helper script.

```plain text
git checkout main
php artisan route:list --json -v | jq > before.json

git checkout working
php artisan route:list --json -v | jq > after.json

```

The script does the following steps:

1. Check out the repository’s `main` branch
2. Outputs all the routes as JSON.
3. The output is “piped” into [`jq`](https://github.com/jqlang/jq) which will “pretty print”, or format, the JSON into something human readable.
4. The formatted output is then written to a `before.json` file.

It then repeats the steps for my working branch but the output is written to the `after.json` file.

The content of the `before.json` and `after.json` files now looks something like the following, which has been truncated for brevity...

```plain text
[
  {
    "domain": null,
    "method": "GET|HEAD",
    "uri": "/",
    "name": null,
    "action": "Closure",
    "middleware": [
      "web"
    ]
  },
  {
    "domain": null,
    "method": "GET|HEAD",
    "uri": "api/user",
    "name": null,
    "action": "Closure",
    "middleware": [
      "api",
      "App\\Http\\Middleware\\Authenticate:sanctum"
    ]
  },
  {
    "domain": null,
    "method": "GET|HEAD",
    "uri": "confirm-password",
    "name": "password.confirm",
    "action": "App\\Http\\Controllers\\Auth\\ConfirmablePasswordController@show",
    "middleware": [
      "web",
      "App\\Http\\Middleware\\Authenticate"
    ]
  },
  // ...
]

```

The only thing needed now was to compare the two files and see what the damage is. Of course I don't want this to be a manual process either! This step is why I wanted the JSON formatted nicely. If it was all a single line it would be impossible to see the difference. With formatted JSON I can use some tooling to see the exact differences before and after the refactor.

My tool of choice here is the diffing tool built into Kitty, the terminal I use, but you could use any diffing tool you have access to.

```plain text
kitty +kitten diff before.json after.json

```

The result was something that look liked the following...

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This gave me a way to visualise, with precision, the differences between each route's middleware before and after the change and detect inconsistecies.

I found this super valuable.

I ended up finding a few differences that were inconsistencies between the contructor and route file middleware - so it turns out it wasn't _just_ aesthetics and consistency but actually help identify some places where our constructor middleware was lying to us.

So a PR with **548** deletions and **359** additions across **111** files was easily merge without manual testing.

Why do you care? Because now you can migrate your middleware to the routes file and also completely refactor your routes file to remove the use of `Route::prefix` and have a flat routes file, the way it should be, without worrying about breaking everything 😎

But hopefully it also gives you some ideas of how you could create single-use tests for your own PRs.
