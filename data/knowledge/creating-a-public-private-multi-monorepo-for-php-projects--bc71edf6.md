---
title: "Creating A Public/Private Multi-Monorepo For PHP Projects"
notion_id: bc71edf6-715b-439b-9485-d8deab078bde
notion_url: https://app.notion.com/p/Creating-A-Public-Private-Multi-Monorepo-For-PHP-Projects-bc71edf6715b439b9485d8deab078bde
last_edited: 2023-01-25T18:22:00.000Z
source_url: https://www.smashingmagazine.com/2021/08/public-private-multi-monorepo-php-projects/
tags: ["English", "Monorepositories", "PHP", "Article"]
---
Let’s see how to use a “multi-monorepo” approach for making the development experience faster, yet keeping your PHP packages private. This solution can be especially beneficial for PRO plugin creators.

To make my development experience faster, I moved all of the PHP packages required by my projects to a monorepo. When each package is hosted in its own repository (the multirepo approach), I would need to develop and test it on its own and then publish it to Packagist before I can install it in other packages via Composer. With the monorepo, because all packages are hosted together, they can be developed, tested, versioned, and released at the same time.

The monorepo that hosts my PHP packages is public, accessible to anyone on GitHub. Git repositories cannot grant different access to different assets; it’s all either public or private. Because I plan to release a pro WordPress plugin, I want its packages to be kept private, meaning they can’t be added to the public monorepo.

The solution I found is to use a “multi-monorepo” approach, comprising two monorepos: one public and one private, with the private monorepo embedding the public one as a Git submodule, allowing it to access its files. The public monorepo can be considered the upstream one, and the private monorepo the downstream one.

Architecture of a multi-monorepo. (Large preview)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

As I kept iterating on my code, the repository set-up that I needed to use at each stage of my project also needed to be upgraded. Hence, I didn’t arrive at the multi-monorepo approach on day one; it was a [process that spanned several years](https://css-tricks.com/from-a-single-repo-to-multi-repos-to-monorepo-to-multi-monorepo/) and took a fair amount of effort, going from a single repository, to multiple repositories, to the monorepo, to, finally, the multi-monorepo.

In this article, I will describe how I set up my multi-monorepo using [Monorepo Builder](https://github.com/symplify/monorepo-builder), which works for PHP projects and is based on Composer.

More after jump! Continue reading below ↓

Meet [**Image Optimization**](https://www.smashingmagazine.com/printed-books/image-optimization/), Addy Osmani’s brand new practical guide to optimizing and delivering **high-quality images** on the web. From formats and compression to delivery and maintenance: everything in one single **528-pages** book.[Jump to table of contents ↬](https://www.smashingmagazine.com/printed-books/image-optimization/)

## Reusing Code in the Multi-Monorepo

The public monorepo at [`leoloso/PoP`](https://github.com/leoloso/PoP) is where I keep all of my PHP projects.

This monorepo contains the workflow file [`generate_plugins.yml`](https://github.com/leoloso/PoP/blob/3cd9eee/.github/workflows/generate_plugins.yml), which generates multiple WordPress plugins for distribution when I [create a release on GitHub](https://github.com/leoloso/PoP/actions/runs/1058751996):

Generating plugins when creating a release. (Large preview)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The workflow configuration is not hardcoded in the YAML file, but rather [injected via PHP code](https://github.com/leoloso/PoP/blob/3cd9eee/.github/workflows/generate_plugins.yml#L55-L57):

```plain text
  - id: output_data
    run: |
      echo "::set-output name=plugin_config_entries::$(vendor/bin/monorepo-builder plugin-config-entries-json)"
```

And the configuration is provided via a [custom PHP class](https://github.com/leoloso/PoP/blob/1965e044e371d62ed98f88f05559e48c26183fbd/src/Config/Symplify/MonorepoBuilder/DataSources/PluginDataSource.php#L16):

```plain text
class PluginDataSource
{
  public function getPluginConfigEntries(): array
  {
    return [
      // GraphQL API for WordPress
      [
        'path' => 'layers/GraphQLAPIForWP/plugins/graphql-api-for-wp',
        'zip_file' => 'graphql-api.zip',
        'main_file' => 'graphql-api.php',
        'dist_repo_organization' => 'GraphQLAPI',
        'dist_repo_name' => 'graphql-api-for-wp-dist',
      ],
      // GraphQL API - Extension Demo
      [
        'path' => 'layers/GraphQLAPIForWP/plugins/extension-demo',
        'zip_file' => 'graphql-api-extension-demo.zip',
        'main_file' => 'graphql-api-extension-demo.php',
        'dist_repo_organization' => 'GraphQLAPI',
        'dist_repo_name' => 'extension-demo-dist',
      ],
    ];
  }
}
```

Generating multiple WordPress plugins all together, and configuring the workflow via PHP, has reduced the amount of time I need to manage the project. The workflow currently handles two plugins (the [GraphQL API](https://github.com/leoloso/PoP/blob/master/layers/GraphQLAPIForWP/plugins/graphql-api-for-wp) and its extension demo), but it could handle 200 without additional effort from me.

It is this set-up that I want to reuse for my private monorepo at `leoloso/GraphQLAPI-PRO`, so that the pro plugins can also be generated without effort.

The code to be reused will comprise:

- the GitHub Actions workflows to generate the WordPress plugins (including [scoping](https://graphql-api.com/blog/graphql-api-for-wp-is-now-scoped-thanks-to-php-scoper/), [downgrading from PHP 8.0 to 7.1](https://graphql-api.com/blog/the-plugin-is-now-transpiled-from-php-80-to-71/), and [uploading to the releases page](https://leoloso.com/posts/github-action-to-release-wp-plugin/)).
- the custom PHP services to configure the workflows.

The private monorepo can then generate the pro WordPress plugins simply by triggering the workflows from the public monorepo and overriding their configuration in PHP.

## Linking Monorepos Via Git Submodules

To embed the public repository in the private one, we use [Git submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules):

```plain text
git submodule add <public repo URL>
```

I embedded the public repository in the subfolder `submodules` of the private monorepo, allowing me to add more upstream monorepos in the future if needed. In GitHub, the folder [displays the submodule’s specific commit](https://github.blog/2016-02-01-working-with-submodules/), and clicking on it will take me to that commit at `leoloso/PoP`:

Embedding the public monorepo in the private monorepo. ([Large preview](https://cloud.netlifyusercontent.com/assets/344dbf88-fdf9-42bb-adb4-46f01eedd629/6a567d4e-b656-4e05-9c96-291487f27a8e/03-embedding-public-monorepo.png))

Because the private repository contains submodules, to clone it, we must provide the `--recursive` option:

```plain text
git clone --recursive <private repo URL>
```

## Reusing the GitHub Actions Workflows

GitHub Actions only loads workflows from under `.github/workflows`. Because the public workflows in the downstream monorepo are under `submodules/PoP/.github/workflows`, these must be duplicated in the expected location.

In order to keep the upstream workflows as the single source of truth, we can limit ourselves to copying the files downstream under `.github/workflows`, but never editing them there. If any change is to be done, it must be done in the upstream monorepo and then copied over.

As a side note, notice how this means that the multi-monorepo leaks: The upstream monorepo is not fully autonomous, and it will need to be adapted to suit the downstream monorepo.

In my first iteration to copy the workflows, I created a simple [Composer script](https://getcomposer.org/doc/articles/scripts.md#defining-scripts):

```plain text

{
  "scripts": {
    "copy-workflows": [
      "php -r \"copy('submodules/PoP/.github/workflows/generate_plugins.yml', '.github/workflows/generate_plugins.yml');\"",
      "php -r \"copy('submodules/PoP/.github/workflows/split_monorepo.yaml', '.github/workflows/split_monorepo.yaml');\""
    ]
  }
}
```

Then, after editing the workflows in the upstream monorepo, I would copy them downstream by executing the following:

```plain text
composer copy-workflows
```

But then I realized that just copying the workflows is not enough: They must also be modified in the process. This is because checking out the downstream monorepo requires the option `--recurse-submodules` in order to also check out the submodules.

In GitHub Actions, the checkout downstream is done like this:

```plain text
  - uses: actions/checkout@v2
    with:
        submodules: recursive
```

So, checking out the downstream repository needs the input `submodules: recursive`, but the upstream one does not, and they both use the same source file.

The solution I found is to provide the value for the input `submodules` via the environment variable `CHECKOUT_SUBMODULES`, which by default is [empty for the upstream repository](https://github.com/leoloso/PoP/blob/aec4615/.github/workflows/coding_standards.yml):

```plain text
env:
  CHECKOUT_SUBMODULES: ""

jobs:
  provide_data:
    steps:
      - uses: actions/checkout@v2
        with:
          submodules: ${{ env.CHECKOUT_SUBMODULES }}
```

Then, when copying the workflows from upstream to downstream, the value of `CHECKOUT_SUBMODULES` is replaced with `recursive`:

```plain text
env:
  CHECKOUT_SUBMODULES: "recursive"
```

When modifying the workflow, it’s a good idea to use a regular expression (regex), so that it works for different formats in the source file (such as `CHECKOUT_SUBMODULES: ""` or `CHECKOUT_SUBMODULES:''` or `CHECKOUT_SUBMODULES:`). This will prevent bugs from being created for these kinds of ostensibly harmless changes.

Thus, the `copy-workflows` Composer script shown above is not good enough to handle this complexity.

In my next iteration, I created a PHP command, `CopyUpstreamMonorepoFilesCommand`, to be executed via Monorepo Builder:

```plain text
vendor/bin/monorepo-builder copy-upstream-monorepo-files
```

This command uses a custom service, `FileCopierSystem`, to copy all files from a source folder to the specified destination, while optionally replacing their contents:

```plain text
namespace PoP\GraphQLAPIPRO\Extensions\Symplify\MonorepoBuilder\SmartFile;

use Nette\Utils\Strings;
use Symplify\SmartFileSystem\Finder\SmartFinder;
use Symplify\SmartFileSystem\SmartFileSystem;

final class FileCopierSystem
{
  public function __construct(
    private SmartFileSystem $smartFileSystem,
    private SmartFinder $smartFinder,
  ) {
  }

  /**
   * @param array $patternReplacements a regex pattern to search, and its replacement
   */
  public function copyFilesFromFolder(
    string $fromFolder,
    string $toFolder,
    array $patternReplacements = []
  ): void {
    $smartFileInfos = $this->smartFinder->find([$fromFolder], '*');

    foreach ($smartFileInfos as $smartFileInfo) {
      $fromFile = $smartFileInfo->getRealPath();
      $fileContent = $this->smartFileSystem->readFile($fromFile);

      foreach ($patternReplacements as $pattern => $replacement) {
        $fileContent = Strings::replace($fileContent, $pattern, $replacement);
      }

      $toFile = $toFolder . substr($fromFile, strlen($fromFolder));
      $this->smartFileSystem->dumpFile($toFile, $fileContent);
    }
  }
}
```

When invoking this method to copy all workflows downstream, I also replace the value of `CHECKOUT_SUBMODULES`:

```plain text
/**
 * Copy all workflows to `.github/`, and convert:
 *   `CHECKOUT_SUBMODULES: ""`
 * into:
 *   `CHECKOUT_SUBMODULES: "recursive"`
 */
$regexReplacements = [
  '#CHECKOUT_SUBMODULES:(\s+".*")?#' => 'CHECKOUT_SUBMODULES: "recursive"',
];
(new FileCopierSystem())->copyFilesFromFolder(
  'submodules/PoP/.github/workflows',
  '.github/workflows',
  $regexReplacements
);
```

The workflow in `generate_plugins.yml` needs an additional replacement. When the WordPress plugin is generated, its code is downgraded from PHP 8.0 to 7.1 [by invoking](https://github.com/leoloso/PoP/blob/3cd9eee96b603a124c18abd6f69f6e937a41477a/.github/workflows/generate_plugins.yml#L116) the script [`ci/downgrade/downgrade_code.sh`](https://github.com/leoloso/PoP/blob/3cd9eee96b603a124c18abd6f69f6e937a41477a/.github/workflows/generate_plugins.yml#L116):

```plain text
  - name: Downgrade code for production (to PHP 7.1)
    run: ci/downgrade/downgrade_code.sh "${{ matrix.pluginConfig.rector_downgrade_config }}" "" "${{ matrix.pluginConfig.path }}" "${{ matrix.pluginConfig.additional_rector_configs }}"
```

In the downstream monorepo, this file will be located under `submodules/PoP/ci/downgrade/downgrade_code.sh`. Then, we point the downstream workflow to the right path with this replacement:

```plain text
$regexReplacements = [
  // ...
  '#(ci/downgrade/downgrade_code\.sh)#' => 'submodules/PoP/$1',
];
```

## Configuring Packages in Monorepo Builder

The file `monorepo-builder.php` — placed at the root of the monorepo — holds the [configuration for Monorepo Builder](https://github.com/symplify/monorepo-builder#1-merge-local-composerjson-to-the-root-one). In it, we must indicate where the packages (and plugins, clients, and anything else) are located:

```plain text
use Symfony\Component\DependencyInjection\Loader\Configurator\ContainerConfigurator;
use Symplify\MonorepoBuilder\ValueObject\Option;

return static function (ContainerConfigurator $containerConfigurator): void {
  $parameters = $containerConfigurator->parameters();
  $parameters->set(Option::PACKAGE_DIRECTORIES, [
    __DIR__ . '/packages',
    __DIR__ . '/plugins',
  ]);
};
```

The private monorepo must have access to all code: its own packages, plus those from the public monorepo. Then, it must define all packages from both monorepos in the configuration file. The ones from the public monorepo are located under `/submodules/PoP`:

```plain text
return static function (ContainerConfigurator $containerConfigurator): void {
  $parameters = $containerConfigurator->parameters();
  $parameters->set(Option::PACKAGE_DIRECTORIES, [
    // public code
    __DIR__ . '/submodules/PoP/packages',
    __DIR__ . '/submodules/PoP/plugins',
    // private code
    __DIR__ . '/packages',
    __DIR__ . '/plugins',
    __DIR__ . '/clients',
  ]);
};
```

As they are, the configurations for upstream and downstream are pretty much the same, the differences being that the downstream one will:

- change the path to the public packages,
- add the private packages.

So, it makes sense to rewrite the configuration using object-oriented programming (OOP). Let’s follow the DRY principle (“don’t repeat yourself”) by having a PHP class in the public repository be extended in the private repository.

## Recreating the Configuration Via OOP

Let’s refactor the configuration. In the public repository, the file `monorepo-builder.php` will simply [reference a new class](https://github.com/leoloso/PoP/blob/f958c8f/monorepo-builder.php), [`ContainerConfigurationService`](https://github.com/leoloso/PoP/blob/f958c8f/monorepo-builder.php), where all of the action will happen:

```plain text
use PoP\PoP\Config\Symplify\MonorepoBuilder\Configurators\ContainerConfigurationService;
use Symfony\Component\DependencyInjection\Loader\Configurator\ContainerConfigurator;

return static function (ContainerConfigurator $containerConfigurator): void {
  $containerConfigurationService = new ContainerConfigurationService(
    $containerConfigurator,
    __DIR__
  );
  $containerConfigurationService->configureContainer();
};
```

The `__DIR__` parameter points to the root of the monorepo. It will be needed to obtain the full path to the package directories.

The class `ContainerConfigurationService` is now in charge of [producing the configuration](https://github.com/leoloso/PoP/blob/1965e04/src/Config/Symplify/MonorepoBuilder/Configurators/ContainerConfigurationService.php#L20):

```plain text
namespace PoP\PoP\Config\Symplify\MonorepoBuilder\Configurators;

use PoP\PoP\Config\Symplify\MonorepoBuilder\DataSources\PackageOrganizationDataSource;
use Symfony\Component\DependencyInjection\Loader\Configurator\ContainerConfigurator;
use Symplify\MonorepoBuilder\ValueObject\Option;

class ContainerConfigurationService
{
  public function __construct(
    protected ContainerConfigurator $containerConfigurator,
    protected string $rootDirectory,
  ) {
  }

  public function configureContainer(): void
  {
    $parameters = $this->containerConfigurator->parameters();
    if ($packageOrganizationConfig = $this->getPackageOrganizationDataSource($this->rootDirectory)) {
      $parameters->set(
        Option::PACKAGE_DIRECTORIES,
        $packageOrganizationConfig->getPackageDirectories()
      );
    }
  }

  protected function getPackageOrganizationDataSource(): ?PackageOrganizationDataSource
  {
    return new PackageOrganizationDataSource($this->rootDirectory);
  }
}
```

The configuration can be split across several classes. In this case, `ContainerConfigurationService` retrieves the package configuration through the class `PackageOrganizationDataSource`, whose [implementation you can see](https://github.com/leoloso/PoP/blob/6260d11/src/Config/Symplify/MonorepoBuilder/DataSources/PackageOrganizationDataSource.php):

```plain text
namespace PoP\PoP\Config\Symplify\MonorepoBuilder\DataSources;

class PackageOrganizationDataSource
{
  public function __construct(protected string $rootDir)
  {
  }

  public function getPackageDirectories(): array
  {
    return array_map(
      fn (string $packagePath) => $this->rootDir . '/' . $packagePath,
      $this->getRelativePackagePaths()
    );
  }

  public function getRelativePackagePaths(): array
  {
    return [
      'packages',
      'plugins',
    ];
  }
}
```

## Overriding the Configuration in the Downstream Monorepo

Now that the configuration in the public monorepo has been set up using OOP, we can extend it to suit the needs of the private monorepo.

In order to allow the private monorepo to autoload the PHP code from the public monorepo, we must first configure the downstream `composer.json` file to reference the source code from upstream, which is under the path `submodules/PoP/src`:

```plain text
{
  "autoload": {
    "psr-4": {
      "PoP\\GraphQLAPIPRO\\": "src",
      "PoP\\PoP\\": "submodules/PoP/src"
    }
  }
}
```

Below is the `monorepo-builder.php` file for the private monorepo. Notice that the referenced class `ContainerConfigurationService` in the upstream repository belonged to the `PoP\PoP` namespace but now has been switched to the `PoP\GraphQLAPIPRO` namespace. This class must receive the additional input of `$upstreamRelativeRootPath` (with a value of `submodules/PoP`) in order to recreate the full path to the public packages:

```plain text
use PoP\GraphQLAPIPRO\Config\Symplify\MonorepoBuilder\Configurators\ContainerConfigurationService;
use Symfony\Component\DependencyInjection\Loader\Configurator\ContainerConfigurator;

return static function (ContainerConfigurator $containerConfigurator): void {
  $containerConfigurationService = new ContainerConfigurationService(
    $containerConfigurator,
    __DIR__,
    'submodules/PoP'
  );
  $containerConfigurationService->configureContainer();
};
```

The downstream class `ContainerConfigurationService` overrides which `PackageOrganizationDataSource` class is used in the configuration:

```plain text
namespace PoP\GraphQLAPIPRO\Config\Symplify\MonorepoBuilder\Configurators;

use PoP\PoP\Config\Symplify\MonorepoBuilder\Configurators\ContainerConfigurationService as UpstreamContainerConfigurationService;
use PoP\GraphQLAPIPRO\Config\Symplify\MonorepoBuilder\DataSources\PackageOrganizationDataSource;
use Symfony\Component\DependencyInjection\Loader\Configurator\ContainerConfigurator;

class ContainerConfigurationService extends UpstreamContainerConfigurationService
{
  public function __construct(
    ContainerConfigurator $containerConfigurator,
    string $rootDirectory,
    protected string $upstreamRelativeRootPath
  ) {
    parent::__construct(
      $containerConfigurator,
      $rootDirectory
    );
  }

  protected function getPackageOrganizationDataSource(): ?PackageOrganizationDataSource
  {
    return new PackageOrganizationDataSource(
      $this->rootDirectory,
      $this->upstreamRelativeRootPath
    );
  }
}
```

Finally, the downstream class `PackageOrganizationDataSource` contains the full path to both public and private packages:

```plain text
namespace PoP\GraphQLAPIPRO\Config\Symplify\MonorepoBuilder\DataSources;

use PoP\PoP\Config\Symplify\MonorepoBuilder\DataSources\PackageOrganizationDataSource as UpstreamPackageOrganizationDataSource;

class PackageOrganizationDataSource extends UpstreamPackageOrganizationDataSource
{
  public function __construct(
    string $rootDir,
    protected string $upstreamRelativeRootPath
  ) {
    parent::__construct($rootDir);
  }

  public function getRelativePackagePaths(): array
  {
    return array_merge(
      // Public packages - Prepend them with "submodules/PoP/"
      array_map(
        fn ($upstreamPackagePath) => $this->upstreamRelativeRootPath . '/' . $upstreamPackagePath,
        parent::getRelativePackagePaths()
      ),
      // Private packages
      [
        'packages',
        'plugins',
        'clients',
      ]
    );
  }
}
```

## Injecting the Configuration From PHP Into GitHub Actions

Monorepo Builder offers the command `packages-json`, which we can use to inject the package paths into the GitHub Actions workflow:

```plain text
jobs:
  provide_data:
    steps:
      - id: output_data
        name: Calculate matrix for packages
        run: |
          echo "::set-output name=matrix::$(vendor/bin/monorepo-builder packages-json)"

    outputs:
      matrix: ${{ steps.output_data.outputs.matrix }}
```

This command produces a stringified JSON. In the workflow, it must be converted to a JSON object via `fromJson`:

```plain text
jobs:
  split_monorepo:
    needs: provide_data
    strategy:
      matrix:
        package: ${{ fromJson(needs.provide_data.outputs.matrix) }}
```

Unfortunately, the command `packages-json` [outputs the package names](https://github.com/symplify/symplify/blob/3fddff4c47663d1519f7c97769bade425d6a4e5f/packages/monorepo-builder/src/Json/PackageJsonProvider.php#L23) but not their paths. This would work if all packages were under the same folder (such as `packages/`), but it doesn’t work in our case because the public and private packages are located in different folders.

Fortunately, Monorepo Builder can be [extended with custom PHP services](https://graphql-api.com/blog/extending-the-monorepo-builder/#heading-optimizing-the-monorepo). So, I created a custom command, `package-entries-json` (via the class [`PackageEntriesJsonCommand`](https://github.com/leoloso/PoP/blob/57a7a21a378f718e1ecfd378fe34ea99fa62c168/src/Extensions/Symplify/MonorepoBuilder/Command/PackageEntriesJsonCommand.php)), that outputs the path to the package.

The workflow was then [updated with the new command](https://github.com/leoloso/PoP/blob/514c27d03b6e438175e887c8dba550355b735b7d/.github/workflows/split_monorepo_tagged.yaml#L34-L37):

```plain text
    run: |
      echo "::set-output name=matrix::$(vendor/bin/monorepo-builder package-entries-json)"
```

Executed on the public monorepo, this produces the following packages (among [many others](https://github.com/leoloso/PoP/actions/runs/1050692821)):

```plain text
[
  {
    "name": "graphql-api-for-wp",
    "path": "layers/GraphQLAPIForWP/plugins/graphql-api-for-wp"
  },
  {
    "name": "extension-demo",
    "path": "layers/GraphQLAPIForWP/plugins/extension-demo"
  },
  {
    "name": "access-control",
    "path": "layers/Engine/packages/access-control"
  },
  {
    "name": "api",
    "path": "layers/API/packages/api"
  },
  {
    "name": "api-clients",
    "path": "layers/API/packages/api-clients"
  }
]
```

Executed on the private monorepo, it produces the following entries (among many others):

```plain text
[
  {
    "name": "graphql-api-for-wp",
    "path": "submodules/PoP/layers/GraphQLAPIForWP/plugins/graphql-api-for-wp"
  },
  {
    "name": "extension-demo",
    "path": "submodules/PoP/layers/GraphQLAPIForWP/plugins/extension-demo"
  },
  {
    "name": "access-control",
    "path": "submodules/PoP/layers/Engine/packages/access-control"
  },
  {
    "name": "api",
    "path": "submodules/PoP/layers/API/packages/api"
  },
  {
    "name": "api-clients",
    "path": "submodules/PoP/layers/API/packages/api-clients"
  },
  {
    "name": "graphql-api-pro",
    "path": "layers/GraphQLAPIForWP/plugins/graphql-api-pro"
  },
  {
    "name": "convert-case-directives",
    "path": "layers/Schema/packages/convert-case-directives"
  },
  {
    "name": "export-directive",
    "path": "layers/GraphQLByPoP/packages/export-directive"
  }
]
```

It works quite well. The configuration for the downstream monorepo contains both public and private packages, and the paths to the public ones are prepended with `submodules/PoP`.

## Skipping Public Packages in the Downstream Monorepo

So far, the downstream monorepo includes both public and private packages in its configuration. However, not every command needs to be executed on the public packages.

Take static analysis, for instance. The public monorepo already executes [PHPStan](https://phpstan.org/) on all public packages via the workflow file [`phpstan.yml`](https://github.com/leoloso/PoP/blob/aec4615/.github/workflows/phpstan.yml), as shown in [this run](https://github.com/leoloso/PoP/runs/3176485273?check_suite_focus=true#step:6:1). If the downstream monorepo ran PHPStan once again on the public packages, it would be a waste of computing time. The `phpstan.yml` workflow only needs to run on the private packages.

This means that, depending on the command to be executed in the downstream repository, we might want to include either both public and private packages or only private ones.

To determine whether to add public packages in the downstream configuration, we adapt the downstream class `PackageOrganizationDataSource` to check this condition via the input `$includeUpstreamPackages`:

```plain text
namespace PoP\GraphQLAPIPRO\Config\Symplify\MonorepoBuilder\DataSources;

use PoP\PoP\Config\Symplify\MonorepoBuilder\DataSources\PackageOrganizationDataSource as UpstreamPackageOrganizationDataSource;

class PackageOrganizationDataSource extends UpstreamPackageOrganizationDataSource
{
  public function __construct(
    string $rootDir,
    protected string $upstreamRelativeRootPath,
    protected bool $includeUpstreamPackages
  ) {
    parent::__construct($rootDir);
  }

  public function getRelativePackagePaths(): array
  {
    return array_merge(
      // Add the public packages?
      $this->includeUpstreamPackages ?
        // Public packages - Prepend them with "submodules/PoP/"
        array_map(
          fn ($upstreamPackagePath) => $this->upstreamRelativeRootPath . '/' . $upstreamPackagePath,
          parent::getRelativePackagePaths()
        ) : [],
      // Private packages
      [
        'packages',
        'plugins',
        'clients',
      ]
    );
  }
}
```

Next, we need to provide the value `$includeUpstreamPackages` as either `true` or `false`, depending on the command to be executed.

We can do this by replacing the configuration file `monorepo-builder.php` with two other configuration files: `monorepo-builder-with-upstream-packages.php` (which passes `$includeUpstreamPackages` => `true`) and `monorepo-builder-without-upstream-packages.php` (which passes `$includeUpstreamPackages` => `false`):

```plain text
// File monorepo-builder-without-upstream-packages.php
use PoP\GraphQLAPIPRO\Config\Symplify\MonorepoBuilder\Configurators\ContainerConfigurationService;
use Symfony\Component\DependencyInjection\Loader\Configurator\ContainerConfigurator;

return static function (ContainerConfigurator $containerConfigurator): void {
  $containerConfigurationService = new ContainerConfigurationService(
    $containerConfigurator,
    __DIR__,
    'submodules/PoP',
    false, // This is $includeUpstreamPackages
  );
  $containerConfigurationService->configureContainer();
};
```

We then update `ContainerConfigurationService` to receive the parameter `$includeUpstreamPackages` and pass it along to `PackageOrganizationDataSource`:

```plain text
namespace PoP\GraphQLAPIPRO\Config\Symplify\MonorepoBuilder\Configurators;

use PoP\PoP\Config\Symplify\MonorepoBuilder\Configurators\ContainerConfigurationService as UpstreamContainerConfigurationService;
use PoP\GraphQLAPIPRO\Config\Symplify\MonorepoBuilder\DataSources\PackageOrganizationDataSource;
use Symfony\Component\DependencyInjection\Loader\Configurator\ContainerConfigurator;

class ContainerConfigurationService extends UpstreamContainerConfigurationService
{
  public function __construct(
    ContainerConfigurator $containerConfigurator,
    string $rootDirectory,
    protected string $upstreamRelativeRootPath,
    protected bool $includeUpstreamPackages,
  ) {
    parent::__construct(
      $containerConfigurator,
      $rootDirectory,
    );
  }

  protected function getPackageOrganizationDataSource(): ?PackageOrganizationDataSource
  {
    return new PackageOrganizationDataSource(
      $this->rootDirectory,
      $this->upstreamRelativeRootPath,
      $this->includeUpstreamPackages,
    );
  }
}
```

Next, we have to invoke the `monorepo-builder` with either configuration file, by providing the `--config` option:

```plain text
jobs:
  provide_data:
    steps:
      - id: output_data
        name: Calculate matrix for packages
        run: |
          echo "::set-output name=matrix::$(vendor/bin/monorepo-builder package-entries-json --config=monorepo-builder-without-upstream-packages.php)"
```

However, as we saw earlier, we want to keep the GitHub Actions workflows in the upstream monorepo as the single source of truth, and they clearly do not need these changes.

The solution I found to this issue is to always provide a `--config` option in the upstream repository, with each command getting its own configuration file, such as the [`validate`](https://github.com/leoloso/PoP/blob/aec4615e782494d6596d9106cc04a47d0234459b/.github/workflows/monorepo_validation.yml#L32)[ command receiving the ](https://github.com/leoloso/PoP/blob/aec4615e782494d6596d9106cc04a47d0234459b/.github/workflows/monorepo_validation.yml#L32)[`validate.php`](https://github.com/leoloso/PoP/blob/aec4615e782494d6596d9106cc04a47d0234459b/.github/workflows/monorepo_validation.yml#L32)[ configuration file](https://github.com/leoloso/PoP/blob/aec4615e782494d6596d9106cc04a47d0234459b/.github/workflows/monorepo_validation.yml#L32):

```plain text
  - name: Run validation
    run: vendor/bin/monorepo-builder validate --config=config/monorepo-builder/validate.php
```

Now, there are no configuration files in the upstream monorepo, because it doesn’t need them. But it will not break, because Monorepo Builder [checks whether the configuration file exists](https://github.com/symplify/symplify/blob/90714eec76fb7de8aa0088748d006a40ae75a21c/packages/monorepo-builder/bin/monorepo-builder.php#L45-L60), and, if it does not, it loads the default configuration file instead. So, either we will override the configuration or nothing will happen.

The downstream repository provides the configuration files for each command, specifying whether to add the upstream packages:

As a side note, this is another example of how the multi-monorepo leaks.

```plain text
// File config/monorepo-builder/validate.php
return require_once __DIR__ . '/monorepo-builder-with-upstream-packages.php';
```

## Overriding the Configuration

We are almost done. By now, the downstream monorepo can override the configuration from the upstream monorepo. So, all that’s left to do is to provide the new configuration.

In the `PluginDataSource` class, I override the configuration of which WordPress plugins must be generated, providing the pro ones instead:

```plain text
namespace PoP\GraphQLAPIPRO\Config\Symplify\MonorepoBuilder\DataSources;

use PoP\PoP\Config\Symplify\MonorepoBuilder\DataSources\PluginDataSource as UpstreamPluginDataSource;

class PluginDataSource extends UpstreamPluginDataSource
{
  public function getPluginConfigEntries(): array
  {
    return [
      // GraphQL API PRO
      [
        'path' => 'layers/GraphQLAPIForWP/plugins/graphql-api-pro',
        'zip_file' => 'graphql-api-pro.zip',
        'main_file' => 'graphql-api-pro.php',
        'dist_repo_organization' => 'GraphQLAPI-PRO',
        'dist_repo_name' => 'graphql-api-pro-dist',
      ],
      // GraphQL API Extensions
      // Google Translate
      [
        'path' => 'layers/GraphQLAPIForWP/plugins/google-translate',
        'zip_file' => 'graphql-api-google-translate.zip',
        'main_file' => 'graphql-api-google-translate.php',
        'dist_repo_organization' => 'GraphQLAPI-PRO',
        'dist_repo_name' => 'graphql-api-google-translate-dist',
      ],
      // Events Manager
      [
        'path' => 'layers/GraphQLAPIForWP/plugins/events-manager',
        'zip_file' => 'graphql-api-events-manager.zip',
        'main_file' => 'graphql-api-events-manager.php',
        'dist_repo_organization' => 'GraphQLAPI-PRO',
        'dist_repo_name' => 'graphql-api-events-manager-dist',
      ],
    ];
  }
}
```

Creating a new release on GitHub will trigger the `generate_plugins.yml` workflow and will generate the pro plugins in my private monorepo:

Generating pro plugins. ([Large preview](https://cloud.netlifyusercontent.com/assets/344dbf88-fdf9-42bb-adb4-46f01eedd629/33ea9f4a-ed86-4208-9b8a-c7b561ac3491/04-generating-pro-plugins.png))

Ta-da!

## Conclusion

As always, there is no “best” solution, only solutions that might work better depending on the context. The multi-monorepo approach is not suitable to every kind of project or team. I believe the biggest beneficiaries would be plugin creators who release public plugins that will be upgraded to their pro versions, as well as agencies that customizes plugins for their clients.

In my case, I’m quite happy with this approach. Getting it right takes a bit of time and effort, but it’s a one-time investment. Once the set-up is over, I can focus on building my pro plugins, and the time saved with project management could be huge.

Tips on front-end & UX, delivered weekly in your inbox. Just the things you can actually use.

With practical takeaways, live sessions, video recordings and a friendly Q&A.
