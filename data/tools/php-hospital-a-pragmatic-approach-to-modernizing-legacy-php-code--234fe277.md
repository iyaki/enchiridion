---
title: "php.hospital - A pragmatic approach to modernizing legacy PHP code"
notion_id: 234fe277-3df9-4e5b-a8b9-2774c6983c42
notion_url: https://app.notion.com/p/php-hospital-A-pragmatic-approach-to-modernizing-legacy-PHP-code-234fe2773df94e5ba8b92774c6983c42
last_edited: 2024-03-25T17:31:00.000Z
source_url: https://php.hospital/
tags: ["English", "PHP", "Website", "Guide"]
---
# php.hospital

I started this website to share some of my learnings and some of the pitfalls I've encountered while modernizing legacy PHP code (my clients' code as well as my own side projects).

I ordered it in a way that I think makes sense, but feel free to jump around as you see fit. I hope you find it useful.

Let's get from here

```php
// No consistent styling, not using modern PHP features to reduce errors
class SettingService
{
    protected $all_settings = array();
    private   $doctrin;

    public function __construct(EntityManagerInterface $em)
    {
        $this->doctrine = $em;
    }

    /**
     * @param $name Can be one of 'site_name', 'site_keywords', 'author'
     */
    public function getSettingByName($name)
    {

        if ($name === null || $name === '') {
            return null;
        }
        $setting = $this->doctrine->getRepository('App:Setting')->findOneBy(['name' => $name]);
        return $setting;
    }
}
```

to here

```php
// No redundancy, using typehints to understand what the code is doing
readonly class SettingService
{
    public function __construct(
        private SettingRepository $settingRepository,
    ) {
    }

    public function getSettingByName(WebsiteSettingName $name): ?WebsiteSetting
    {
        return $this->settingRepository->findOneBy(['me' => $name]);
    }
}
```

and from here

```php
// All code crammed into one function
public function findJobsByCriteria(
    ?string $jobName = null,
    ?string $startDateEarliest = null,
    ?string $startDateLatest = null,
    ?string $page = null,
    ?string $pageSize = null,
): mixed {
    $queryBuilder = $this->createQueryBuilder('j');

    if ($jobName !== null) {
        $queryBuilder
            ->andWhere('j.name LIKE :jobName')
            ->setParameter('jobName', '%'.$jobName.'%')
        ;
    }

    if ($startDateEarliest !== null) {
        // ...
    }

    // ...
    $queryBuilder
        ->setFirstResult(($page - 1) * $pageSize)
        ->setMaxResults($pageSize)
    ;

    return $queryBuilder->getQuery()->getResult();
}
```

to here

```php
// Applying testable, extendible and modular query filters using the Filter Pattern
public function findJobsByCriteria(array $queryParameters): array {
    $queryBuilder = $this->createQueryBuilder('j');
    $queryFilter->apply($queryBuilder, Job::class, $queryParameters);

    return $queryBuilder->getQuery()->getResult();
}
```
