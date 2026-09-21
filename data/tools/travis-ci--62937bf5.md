---
title: "Travis-CI"
notion_id: 62937bf5-422d-4d81-a2b6-96c2e2462a54
notion_url: https://app.notion.com/p/Travis-CI-62937bf5422d4d81a2b696c2e2462a54
last_edited: 2022-12-20T04:33:00.000Z
source_url: https://www.travis-ci.com/
tags: ["English", "Continuous Integration/Continuous Delivery", "DevOps", "Untried", "Service"]
---
## How developers build simple, trustworthy CI/CD pipelines

Join hundreds of thousands who define tests and deployments in minutes, then scale up simply with parallel or multi-environment builds using Travis CI’s precision syntax—all designed with the developer experience in mind.

Where developers trust Travis CI with testing and automation

1.

## Build your new CI/CD pipeline in 20 minutes or less

Whether starting from scratch or rebuilding a rich ecosystem, get pipelines flowing with preconfigured environments for your language and a minimal syntax with up to 50% less YAML or JSON.

```plain text
language: python python: - "3.7" - "3.8" - "3.9" # Command to install dependencies install: - pip install -r requirements.txt - pip install pytest pytest-cov # Command to run tests script: - pytest --cov=./ tests/ # Specify branches to build branches: only: - main - develop # Cache pip dependencies cache: pip # Run jobs in parallel jobs: include: - name: "Lint" python: "3.9" before_script: - pip install flake8 script: - flake8 . - name: "Docs" python: "3.9" before_script: - pip install sphinx script: - sphinx-build -b html docs docs/_build after_success: - pip install coveralls - coveralls notifications: email: on_success: never on_failure: always
```

```plain text
language: node_js node_js: - "12" - "14" - "16" # Specify the operating system os: - linux - osx # Cache dependencies cache: directories: - node_modules # Install dependencies install: - npm ci # Run tests script: - npm run test - npm run build # Run jobs in parallel jobs: include: - stage: lint node_js: "16" script: npm run lint - stage: coverage node_js: "16" script: npm run test:coverage # Deploy to npm deploy: provider: npm email: "[email protected]" api_key: $NPM_TOKEN on: tags: true branch: main # Notifications notifications: email: on_success: never on_failure: always slack: rooms: - secure: "encrypted-slack-webhook-url"
```

```plain text
language: java jdk: - openjdk11 - openjdk14 # Cache dependencies cache: directories: - $HOME/.m2 # Build stages stages: - test - name: deploy if: branch = main jobs: include: - stage: test name: "Unit Tests" script: mvn test - stage: test name: "Integration Tests" script: mvn verify -Pintegration-tests - stage: deploy name: "Deploy to Staging" script: mvn deploy -Pstaging # Notifications notifications: email: recipients: - [email protected] on_success: change on_failure: always # Before install steps before_install: - chmod +x mvnw # Main build script script: - ./mvnw clean install # After success steps after_success: - bash <(curl -s https://codecov.io/bash) # Only build these branches branches: only: - main - develop
```

```plain text
language: cpp compiler: - gcc - clang before_install: - sudo apt-get update - sudo apt-get install -y cmake script: - mkdir build - cd build - cmake .. - make - make test notifications: email: false
```

```plain text
language: php php: - 7.4 - 8.0 - 8.1 before_script: - composer self-update - composer install --no-interaction script: - vendor/bin/phpunit - vendor/bin/phpcs cache: directories: - $HOME/.composer/cache notifications: email: false
```

```plain text
language: rust rust: - stable - beta - nightly matrix: allow_failures: - rust: nightly fast_finish: true cache: cargo before_script: - rustup component add clippy script: - cargo build --verbose - cargo test --verbose - cargo clippy -- -D warnings notifications: email: false
```

```plain text
language: go go: - 1.16.x - 1.17.x - 1.18.x - tip matrix: allow_failures: - go: tip fast_finish: true before_install: - go get -t -v ./... script: - go test -race -coverprofile= coverage.txt -covermode=atomic - go vet ./... - diff -u <(echo -n) <(gofmt -d .) after_success: - bash <(curl -s https://codecov.io/bash) notifications: email: false
```

```plain text
language: csharp mono: none dotnet: 6.0.100 install: - dotnet restore script: - dotnet build - dotnet test notifications: email: false branches: only: - main - develop cache: directories: - $HOME/.nuget/packages
```

```plain text
language: ruby rvm: - 2.6 - 2.7 - 3.0 install: - bundle install script: - bundle exec rake test after_success: - bundle exec rake coveralls:push notifications: email: recipients: - [email protected] on_success: always on_failure: always
```

1.

## Build your new CI/CD pipeline in no time

Whether starting from scratch or rebuilding a rich ecosystem, get pipelines flowing with preconfigured environments for your language and a minimal syntax with up to 50% less YAML or JSON.

```plain text
language: python python: - "3.7" - "3.8" - "3.9" # Command to install dependencies install: - pip install -r requirements.txt - pip install pytest pytest-cov # Command to run tests script: - pytest --cov=./ tests/ # Specify branches to build branches: only: - main - develop # Cache pip dependencies cache: pip # Run jobs in parallel jobs: include: - name: "Lint" python: "3.9" before_script: - pip install flake8 script: - flake8 . - name: "Docs" python: "3.9" before_script: - pip install sphinx script: - sphinx-build -b html docs docs/_build after_success: - pip install coveralls - coveralls notifications: email: on_success: never on_failure: always
```

```plain text
language: node_js node_js: - "12" - "14" - "16" # Specify the operating system os: - linux - osx # Cache dependencies cache: directories: - node_modules # Install dependencies install: - npm ci # Run tests script: - npm run test - npm run build # Run jobs in parallel jobs: include: - stage: lint node_js: "16" script: npm run lint - stage: coverage node_js: "16" script: npm run test:coverage # Deploy to npm deploy: provider: npm email: "[email protected]" api_key: $NPM_TOKEN on: tags: true branch: main # Notifications notifications: email: on_success: never on_failure: always slack: rooms: - secure: "encrypted-slack-webhook-url"
```

```plain text
language: java jdk: - openjdk11 - openjdk14 # Cache dependencies cache: directories: - $HOME/.m2 # Build stages stages: - test - name: deploy if: branch = main jobs: include: - stage: test name: "Unit Tests" script: mvn test - stage: test name: "Integration Tests" script: mvn verify -Pintegration-tests - stage: deploy name: "Deploy to Staging" script: mvn deploy -Pstaging # Notifications notifications: email: recipients: - [email protected] on_success: change on_failure: always # Before install steps before_install: - chmod +x mvnw # Main build script script: - ./mvnw clean install # After success steps after_success: - bash <(curl -s https://codecov.io/bash) # Only build these branches branches: only: - main - develop
```

```plain text
language: cpp compiler: - gcc - clang before_install: - sudo apt-get update - sudo apt-get install -y cmake script: - mkdir build - cd build - cmake .. - make - make test notifications: email: false
```

```plain text
language: php php: - 7.4 - 8.0 - 8.1 before_script: - composer self-update - composer install --no-interaction script: - vendor/bin/phpunit - vendor/bin/phpcs cache: directories: - $HOME/.composer/cache notifications: email: false
```

```plain text
language: rust rust: - stable - beta - nightly matrix: allow_failures: - rust: nightly fast_finish: true cache: cargo before_script: - rustup component add clippy script: - cargo build --verbose - cargo test --verbose - cargo clippy -- -D warnings notifications: email: false
```

```plain text
language: go go: - 1.16.x - 1.17.x - 1.18.x - tip matrix: allow_failures: - go: tip fast_finish: true before_install: - go get -t -v ./... script: - go test -race -coverprofile= coverage.txt -covermode=atomic - go vet ./... - diff -u <(echo -n) <(gofmt -d .) after_success: - bash <(curl -s https://codecov.io/bash) notifications: email: false
```

```plain text
language: csharp mono: none dotnet: 6.0.100 install: - dotnet restore script: - dotnet build - dotnet test notifications: email: false branches: only: - main - develop cache: directories: - $HOME/.nuget/packages
```

```plain text
language: ruby rvm: - 2.6 - 2.7 - 3.0 install: - bundle install script: - bundle exec rake test after_success: - bundle exec rake coveralls:push notifications: email: recipients: - [email protected] on_success: always on_failure: always
```

2.

## A precision tool for every CI/CD demand

Travis CI elegantly compacts all your testing automation into one minimal Configuration as Code file for a true developer experience. Simple syntax and effortless extensibility.

```plain text
language: python python: - "3.7" - "3.8" - "3.9" env: - DJANGO_VERSION=2.2 - DJANGO_VERSION=3.0 - DJANGO_VERSION=3.1 matrix: include: - python: "3.9" env: DJANGO_VERSION=3.2 exclude: - python: "3.7" env: DJANGO_VERSION=3.1 allow_failures: - python: "3.9" env: DJANGO_VERSION=3.2 before_install: - pip install -U pip - pip install -U setuptools - pip install -q Django==$DJANGO_VERSION install: - pip install -r requirements.txt script: - python manage.py test after_success: - coveralls notifications: email: false
```

Run tests against different versions of your runtimes or dependencies—or even multiple languages—for comprehensive automation and absolute quality guarantees on your way to production.

```plain text
language: python python: - "3.8" services: - mongodb - redis - mysql env: - DB=mongodb - DB=redis - DB=mysql before_script: - if [[ "$DB" == "mongodb" ]]; then mongo mydb_test --eval 'db.createUser({user:"travis",pwd:"test",roles: ["readWrite"]});'; fi - if [[ "$DB" == "mysql" ]]; then mysql -e 'CREATE DATABASE IF NOT EXISTS mydb_test; '; fi script: - if [[ "$DB" == "mongodb" ]]; then python test_mongodb.py; fi - if [[ "$DB" == "redis" ]]; then python test_redis.py; fi - if [[ "$DB" == "mysql" ]]; then python test_mysql.py; fi
```

Quickly split complex operations, like unit and integration tests, into multiple jobs that give you faster visual feedback on code quality and results—stop waiting on your pipeline to finish and get back to your IDE.

```plain text
language: python python: - "3.8" # Install dependencies install: - pip install -r requirements.txt # Run tests script: - pytest # Integration with Slack for notifications notifications: slack: your-slack-workspace:your-slack-token # Integration with Codecov for code coverage reports after_success: - pip install codecov - codecov # Integration with GitHub Pages for documentation deployment deploy: provider: pages skip_cleanup: true github_token: $GITHUB_TOKEN on: branch: main
```

Enable the must-haves, like build status images, static code analysis, and key management without wasting time on needlessly complex glue. For everything else, the Travis CI API has your back.

```plain text
language: python python: - "3.8" script: - python -m pytest tests/ notifications: email: recipients: - [email protected] on_success: change on_failure: always slack: rooms: - secure: "encrypted-slack-webhook-url" on_success: always on_failure: always template: - "Build <%{build_url}|#%{build_number}> (<%{compare_url}|%{commit}>) of %{repository_slug} @%{branch} by %{author} %{result} in %{duration}" webhooks: urls: - https://webhook.site/your-unique-id on_success: change on_failure: always irc: channels: - "irc.freenode.org#myproject" template: - "%{repository} (%{commit}) : %{message}" - "Build details: %{build_url}" pushover: api_key: "pushover-api-key" users: - "pushover-user-key" template: "%{repository} (%{commit}) : %{message} - %{duration}" after_failure: - cat /home/travis/build/your-github-username/ your-repo-name/tests/test-suite.log
```

Switch on highly customizable notifications to email, Slack, Opsgenie, any webhook destination, and many more. All you need to bring is your API token and as little as two fresh lines of YAML.

```plain text
language: python python: - "3.8" stages: - test - name: deploy if: branch = main jobs: include: - stage: test name: "Unit Tests" script: python -m unittest discover tests/unit - stage: test name: "Integration Tests" script: python -m unittest discover tests/integration - stage: deploy name: "Deploy to Production" script: - pip install awscli - aws s3 sync . s3://my-bucket/ --delete env: global: - PYTHONPATH=$PYTHONPATH:$TRAVIS_BUILD_DIR before_install: - pip install -r requirements.txt branches: only: - main - develop notifications: email: on_success: never on_failure: always
```

Infinitely organize complex CI/CD pipelines into groups that only run if other parallel jobs have been completed successfully. Catch failures faster or layer in smart conditionals to catch edge cases other CI/CD tools lose track of.

```plain text
language: generic jobs: include: # Linux builds - os: linux arch: amd64 - os: linux arch: arm64 # macOS builds - os: osx osx_image: xcode12.5 # Windows builds - os: windows script: - echo "Running tests on $TRAVIS_OS_NAME $TRAVIS_CPU_ARCH" - # Add your build and test commands here notifications: email: false
```

Test your code in parallel on more flexible combinations than any other CI/CD tool. Mix-and-match amd64, ppc64le, s390x, arm64, and arm64-graviton2 CPUs with Linux, macOS, and Windows environments to make your builds edge case-free.

2.

## A precision tool for every CI/CD demand

Travis CI elegantly compacts all your testing automation into one minimal Configuration as Code file for a true developer experience. Simple syntax and effortless extensibility.

```plain text
language: python python: - "3.7" - "3.8" - "3.9" env: - DJANGO_VERSION=2.2 - DJANGO_VERSION=3.0 - DJANGO_VERSION=3.1 matrix: include: - python: "3.9" env: DJANGO_VERSION=3.2 exclude: - python: "3.7" env: DJANGO_VERSION=3.1 allow_failures: - python: "3.9" env: DJANGO_VERSION=3.2 before_install: - pip install -U pip - pip install -U setuptools - pip install -q Django==$DJANGO_VERSION install: - pip install -r requirements.txt script: - python manage.py test after_success: - coveralls notifications: email: false
```

Run tests against different versions of your runtimes or dependencies—or even multiple languages—for comprehensive automation and absolute quality guarantees on your way to production.

```plain text
language: python python: - "3.8" services: - mongodb - redis - mysql env: - DB=mongodb - DB=redis - DB=mysql before_script: - if [[ "$DB" == "mongodb" ]]; then mongo mydb_test --eval 'db.createUser({user:"travis",pwd:"test",roles: ["readWrite"]});'; fi - if [[ "$DB" == "mysql" ]]; then mysql -e 'CREATE DATABASE IF NOT EXISTS mydb_test; '; fi script: - if [[ "$DB" == "mongodb" ]]; then python test_mongodb.py; fi - if [[ "$DB" == "redis" ]]; then python test_redis.py; fi - if [[ "$DB" == "mysql" ]]; then python test_mysql.py; fi
```

Quickly split complex operations, like unit and integration tests, into multiple jobs that give you faster visual feedback on code quality and results—stop waiting on your pipeline to finish and get back to your IDE.

```plain text
language: python python: - "3.8" # Install dependencies install: - pip install -r requirements.txt # Run tests script: - pytest # Integration with Slack for notifications notifications: slack: your-slack-workspace:your-slack-token # Integration with Codecov for code coverage reports after_success: - pip install codecov - codecov # Integration with GitHub Pages for documentation deployment deploy: provider: pages skip_cleanup: true github_token: $GITHUB_TOKEN on: branch: main
```

Enable the must-haves, like build status images, static code analysis, and key management without wasting time on needlessly complex glue. For everything else, the Travis CI API has your back.

```plain text
language: python python: - "3.8" script: - python -m pytest tests/ notifications: email: recipients: - [email protected] on_success: change on_failure: always slack: rooms: - secure: "encrypted-slack-webhook-url" on_success: always on_failure: always template: - "Build <%{build_url}|#%{build_number}> (<%{compare_url}|%{commit}>) of %{repository_slug} @%{branch} by %{author} %{result} in %{duration}" webhooks: urls: - https://webhook.site/your-unique-id on_success: change on_failure: always irc: channels: - "irc.freenode.org#myproject" template: - "%{repository} (%{commit}) : %{message}" - "Build details: %{build_url}" pushover: api_key: "pushover-api-key" users: - "pushover-user-key" template: "%{repository} (%{commit}) : %{message} - %{duration}" after_failure: - cat /home/travis/build/your-github-username/ your-repo-name/tests/test-suite.log
```

Switch on highly customizable notifications to email, Slack, Opsgenie, any webhook destination, and many more. All you need to bring is your API token and as little as two fresh lines of YAML.

```plain text
language: python python: - "3.8" stages: - test - name: deploy if: branch = main jobs: include: - stage: test name: "Unit Tests" script: python -m unittest discover tests/unit - stage: test name: "Integration Tests" script: python -m unittest discover tests/integration - stage: deploy name: "Deploy to Production" script: - pip install awscli - aws s3 sync . s3://my-bucket/ --delete env: global: - PYTHONPATH=$PYTHONPATH:$TRAVIS_BUILD_DIR before_install: - pip install -r requirements.txt branches: only: - main - develop notifications: email: on_success: never on_failure: always
```

Infinitely organize complex CI/CD pipelines into groups that only run if other parallel jobs have been completed successfully. Catch failures faster or layer in smart conditionals to catch edge cases other CI/CD tools lose track of.

```plain text
language: generic jobs: include: # Linux builds - os: linux arch: amd64 - os: linux arch: arm64 # macOS builds - os: osx osx_image: xcode12.5 # Windows builds - os: windows script: - echo "Running tests on $TRAVIS_OS_NAME $TRAVIS_CPU_ARCH" - # Add your build and test commands here notifications: email: false
```

Test your code in parallel on more flexible combinations than any other CI/CD tool. Mix-and-match amd64, ppc64le, s390x, arm64, and arm64-graviton2 CPUs with Linux, macOS, and Windows environments to make your builds edge case-free.

language: gogo:- "1.15"- "1.16"- "1.17"

install:- go mod download

script:- go test ./...

branches:only:- main- develop

cache:directories:- $GOPATH/pkg/mod

jobs:include:- name: "Lint"go: "1.17"script:- go install golang.org/x/lint/golint@latest && golint ./...

- name: "Docs"go: "1.17"script:- go install github.com/swaggo/swag/cmd/swag@latest && swag init

language: rubyrvm:- 2.6- 2.7- 3.0

install:- bundle install

script:- bundle exec rspec

branches:only:- main- develop

cache:directories:- vendor/bundle

jobs:include:- name: "Lint"rvm: 3.0script:- bundle exec rubocop

- name: "Docs"rvm: 3.0script:- bundle exec yard

3.

## The most resilient name in CI/CD

Others commodify CI/CD, building overly complex DevSecOps Platforms. We remain steadfast in our original mission to help developers conquer their CI/CD pipelines faster, more fluently, and, dare we say… with more fun?

### Why pay for CI/CD

That’s a tough but fair question. We’re still delivering the simplest and most extensible CI/CD for developers who want precision tools, not bloated platforms. Travis CI uniquely helps you take ownership of code quality, collaborate better with your peers, and take ownership of the results you create together.

### Engineering-driven customer support

Our CI/CD experts are engineers, not just support agents. They understand your challenges and speak your language. Whether you're exploring Travis CI Enterprise or have a technical question, we’re here to guide you. Get fast, knowledgeable help when you need it most.

### A genuine community

You’re not alone. From build optimization tips to custom pipeline setups, our CI/CD community is always ready to help. Share knowledge, swap ideas, or get feedback from real developers who’ve been there. At Travis CI, community means collaboration.

### Our security-forward promises

Security is built into everything we do. From GDPR compliance to build isolation, we help protect your code and team. With features like HashiCorp Vault integration, secure artifact signing, and scoped credentials, Travis CI lets you ship with confidence and peace of mind.

### Simple to start.Intuitive to extend.

Developed upon by hundreds of thousands.

Newsletter

Get tips, promotions and exclusive offers

© Copyright 2026, All Rights Reserved
