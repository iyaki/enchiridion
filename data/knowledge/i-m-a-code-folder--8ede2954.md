---
title: "I'm a code folder"
notion_id: 8ede2954-886d-4376-b768-327ad10ae500
notion_url: https://app.notion.com/p/I-m-a-code-folder-8ede2954886d4376b768327ad10ae500
last_edited: 2023-02-16T19:36:00.000Z
source_url: https://stitcher.io/blog/code-folding
tags: ["English", "Programming", "Productivity", "Article", "stitcher blog"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This is where the ad would go. Instead though, I'd like to point you towards my [GitHub Sponsors](https://aggregate.stitcher.io/links/27b1f2b4-2cde-4e8f-8857-491bd3bccee4) page. If you're a regular reader and my content is helping you, you can consider a one-time or monthly sponsorship. If you're a company looking for dedicated ad placements on this blog or my newsletter, you can email me at [brendt@stitcher.io](mailto:brendt@stitcher.io)

I know it looks strange the first time you see it, but hear me out for a minute: I am a code folder.

```php
class TwitterSyncCommand extends Command
{
    protected $signature = 'twitter:sync {--clean}';

    /** @var \Illuminate\Database\Eloquent\Collection<Mute> */
    private Collection $mutes;

    public function handle(Twitter $twitter) { /* … */ }

    public function syncFromList(Twitter $twitter): void { /* … */ }

    public function syncFromSearch(Twitter $twitter): void { /* … */ }

    private function storeTweets(array $tweets, TweetFeedType $feedType): void { /* … */ }

    private function shouldBeRejected(Tweet $tweet): ?RejectionReason { /* … */ }
}
```

I hide most of my code, most of the time. I have keyboard shortcuts to easily show and hide blocks of code; and when I open a file, all method and function bodies are collapsed by default.

The reason? I’m not a superhuman speed reader that understands dozens of lines of code at one glance. And… I also don’t have a two-metre-high screen.

I just can’t read and understand all of this — you know?

```php
class TwitterSyncCommand extends Command
{
    protected $signature = 'twitter:sync {--clean}';

    /** @var \Illuminate\Database\Eloquent\Collection<Mute> */
    private Collection $mutes;

    public function handle(Twitter $twitter)
    {
        $this->mutes = Mute::query()->select('text')->get();

        if ($this->option('clean')) {
            $this->error('Truncating tweets!');

            Tweet::truncate();
        }

        $this->syncFromSearch($twitter);

        $this->syncFromList($twitter);

        $this->info('Done');
    }

    public function syncFromList(Twitter $twitter): void
    {
        do {
            $lastTweet = Tweet::query()
                ->where('feed_type', TweetFeedType::LIST)
                ->orderByDesc('tweet_id')
                ->first();

            $tweets = $twitter->request('lists/statuses.json', 'GET', [
                'list_id' => config('services.twitter.list_id'),
                'since_id' => $lastTweet?->tweet_id,
                'count' => 200,
                'tweet_mode' => 'extended',
            ]);

            $count = count($tweets);

            if ($count === 0) {
                $this->comment('No more new tweets');
            } else {
                $this->comment("Syncing {$count} tweets from list");

                $this->storeTweets($tweets, TweetFeedType::LIST);
            }
        } while ($tweets !== []);
    }

    public function syncFromSearch(Twitter $twitter): void
    {
        do {
            $lastTweet = Tweet::query()
                ->where('feed_type', TweetFeedType::SEARCH)
                ->orderByDesc('tweet_id')
                ->first();

            $tweets = $twitter->request('/search/tweets.json', 'GET', [
                'q' => 'phpstorm',
                'since_id' => $lastTweet?->tweet_id,
                'count' => 200,
                'tweet_mode' => 'extended',
            ])->statuses;

            $count = count($tweets);

            if ($count === 0) {
                $this->comment('No more new tweets');
            } else {
                $this->comment("Syncing {$count} tweets from search");

                $this->storeTweets($tweets, TweetFeedType::SEARCH);
            }
        } while ($tweets !== []);
    }

    private function storeTweets(array $tweets, TweetFeedType $feedType): void
    {
        foreach ($tweets as $tweet) {
            $subject = $tweet->retweeted_status ?? $tweet;

            $tweet = Tweet::updateOrCreate([
                'tweet_id' => $tweet->id,
            ], [
                'state' => TweetState::PENDING,
                'feed_type' => $feedType,
                'text' => $subject->full_text ,
                'user_name' => $subject->user->screen_name,
                'retweeted_by_user_name' => isset($tweet->retweeted_status)
                    /** @phpstan-ignore-next-line  */
                    ? $tweet->user->screen_name
                    : null,
                'created_at' => Carbon::make($subject->created_at),
                'payload' => json_encode($tweet),
            ]);

            if ($reason = $this->shouldBeRejected($tweet)) {
                $tweet->update([
                    'state' => TweetState::REJECTED,
                    'rejection_reason' => $reason->message,
                ]);
            }

            (new ParseTweetText)($tweet);
        }
    }

    private function shouldBeRejected(Tweet $tweet): ?RejectionReason
    {
        if ($tweet->isRetweet() && $tweet->feed_type === TweetFeedType::SEARCH) {
            return RejectionReason::retweetedFromSearch();
        }

        // Reject tweets containing a specific word
        foreach ($this->mutes as $mute) {
            if ($tweet->containsPhrase($mute->text)) {
                return RejectionReason::mute($mute->text);
            }
        }

        // Reject replies
        if ($tweet->getPayload()->in_reply_to_status_id) {
            return RejectionReason::isReply();
        }

        // Reject mentions
        if (str_starts_with($tweet->text, '@')) {
            return RejectionReason::isMention();
        }

        // Reject non-english tweets
        $language = $tweet->getPayload()->lang;

        if ($language !== 'en') {
            return RejectionReason::otherLanguage($language);
        }

        return null;
    }
}
```

I feel overwhelmed looking at all that code, especially when I’m searching for one specific method.

Now, when I create a new file, it’s all fine, it’s a blank slate, and I’ve got control. But code grows rapidly, and real life projects more often than not include work in existing files instead of blank slates. So I really need a way to reduce cognitive load, I can’t “take it all in at once”. And that’s why I’ve grown to like code folding so much.

Give yourself a week to get used to it. Configure your IDE to automatically fold method bodies and functions, and assign proper key bindings to show and hide blocks. I promise you, you’ll like it.
