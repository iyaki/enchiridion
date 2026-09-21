---
title: "Contract Tests"
notion_id: 825ee350-6e20-4fe4-8c97-5a134129866a
notion_url: https://app.notion.com/p/Contract-Tests-825ee3506e204fe48c975a134129866a
last_edited: 2023-04-25T13:24:00.000Z
source_url: https://www.kai-sassnowski.com/post/contract-tests/
tags: ["English", "Testing", "Object Oriented Programming", "Article", "on error resume next (Kai Sassnowski)"]
---
[https://www.kai-sassnowski.com/post/contract-tests/](https://www.kai-sassnowski.com/post/contract-tests/)

In this post, I want to take a look at how you can write tests to ensure consistent behavior across all implementations of an interface.

## Interfaces, what gives?

Interfaces are used to decouple code from any specific implementation. Instead, we depend on a more general contract. This contract can then be implemented by many different classes to provide different implementations.

A good example of this is a file system abstraction like Flysystem. Flysystem provides many different implementations of the same generic file system interface. This allows us to work with files in a more general way rather than having to worry about if we’re saving our files to S3, an FTP server or to the local hard drive directly.

The idea of decoupling our code from any specific file system is that we can swap in and out different implementations without changing our code. We could, for example, save to our local hard drive during development, use an in-memory file system in our tests and Amazon S3 in production. All without making a single change to our code!

This all hinges on the fact, however, that not only do all these different implementation have to implement the same interface, they need to behave exactly the same. And this is where things get interesting.

## Where things get interesting

Consider a very basic interface like this.

```plain text
interface Filesystem { public function saveFile(File $file, string $filename): void; }
```

This looks fairly straight-forward. We can save a file by calling saveFile with a File object and a desired $filename. The function doesn’t return any result, just to keep the example simple.

Now let’s look at three different implementations of this interface.

```plain text
class ImplementationA implements Filesystem { public function saveFile(File $file, string $filename): void { $path = $this->pathPrefix . $filename; file_put_contents($path, $file->getContent()); } } class ImplementationB implements Filesystem { public function saveFile(File $file, string $filename): void { $path = $this->pathPrefix . $filename; if (file_exists($path)) { throw new FileAlreadyExistsException($path); } file_put_contents($path, $file->getContent()); } } class ImplementationC implements Filesystem { public function saveFile(File $file, string $filename): void { // Gotcha, sucka! } }
```

See the problem?

All three classes are valid implementations of the Filesystem interface. They behave very differently, however:

- ImplementationA doesn’t care if the file already exists or not, it will simply override the contents
- ImplementationB throws an exception if a file already exists with the same name
- ImplementationC is just the worst. It doesn’t even save the file. And if you decide to check out the implementation to see what’s going wrong it even taunts you!

This is a problem because even though all three classes implement the same interface they aren’t interchangeable. They behave differently in the same situation so our code has to account for that. Which is not the point of using an interface.

When a class implements an interface, all that means is that it provides the correct methods with the correct signatures. It doesn’t provide any guarantees that these methods behave consistently, or even do anything at all.

## What’s in a contract

I have been using the words “contract” and “interface” more or less interchangeably up until now. From now on, I want be a bit more specific about these terms.

- When I say interface, I mean the Interface keyword in your programming language of choice. It’s a language construct.
- A contract is a more general concept. It’s what the interface promises to the consumer beyond just the method signatures. This is about how the object behaves.

Looking back at our Filesystem interface from above, which behaviors do we want to be part of its contract?

First of all, we want to actually save the file. This might sound like a joke but as you saw above, an interface cannot even guarantee this simple fact. This is a valid implementation of the interface.

```plain text
class ImplementationC implements Filesystem { public function saveFile(File $file, string $filename): void { // Gotcha, sucka! } }
```

Secondly, we need to define what should happen when a file with the same name already exists. Should we overwrite it, silently fail, or throw an exception? The point is that there is no “right” answer to this question. But whatever you decide on, every class that implements this interface needs to behave the same way when the target file already exists.

So how can we enforce the contract and not just the interface? This is where contract tests come into play.

## Contract tests

What we want to do is write a generic test for the Filesysteminterface. This test should describe the expected behavior in terms of only the methods that are available on the interface. This way, we can’t accidentally create a test that only works for a specific implementation.

For this reason, we need to extend our interface slightly to add a fileExists method as well.

```plain text
interface Filesystem { public function storeFile(File $file, string $name): void; public function fileExists(string $filename): bool; }
```

With these two methods in place, we can write a generic test to ensure that storeFile actually stores the file.

But... wait a second. What object are we actually calling these methods on? Because remember, we’re not writing a test for any particular implementation. You cannot create an instance of an interface, so new Filesystem() is out.

The solution is to define an abstract method that needs to be implemented inside the test case of each implementation. When called, this method will return an instance of whatever class we’re testing. Like so:

```plain text
protected abstract function getInstance(): Filesystem; /** @test */ public function storingAFileForTheFirstTime(): void { $filesystem = $this->getInstance(); $this->assertFalse( $filesystem->fileExists('::filename::'); ) $filesystem->storeFile( new MockFile('::contents::'), '::filename::' ); $this->assertTrue( $filesystem->fileExists('::filename::') ); }
```

This test defines the expected behavior when trying to save a file for the first time (i.e. the file doesn’t exist). There are a few interesting things to note:

- This test doesn’t know what implementation of a Filesystem it’s dealing with.
- So in order to get an instance of the class we’re supposed to be testing, it defines an abstract method getInstance. This method needs to be implemented in the test file of the concrete implementation.
- The test describes the expected behavior of the storeFile method in very general terms. After we store a file, we expect it to exist. We make no assumptions about what it means to store a file or for a file to exist. Instead, we leverage the fileExists method that we added to our interface.
- To ensure that the fileExists method doesn’t simply return true, we check that file does in fact not exist before we save it.

So where do we put these two methods? There are two options: an abstract class or a trait. To not muck around with inheritance too much, I usually opt for a trait.

```plain text
trait FilesystemContractTests { protected abstract function getInstance(): Filesystem; /** @test */ public function storingAFileForTheFirstTime(): void {…} }
```

Any class that wants to implement the Filesystem interface should use this trait and implement any abstract methods.

```plain text
class InMemoryFilesystemTest extends TestCase { use FilesystemContractTest; protected abstract function getInstance(): Filesystem { return new InMemoryFilesystem(); } }
```

Here is an example implementation of the InMemoryFilesystem class that would make this contract test pass.

```plain text
class InMemoryFilesystem implements Filesystem { private array $files = []; public function storeFile(File $file, string $name): void { $this->files[$name] = $file; } public function fileExists(string $filename): bool { return isset($this->files[$filename]); } }
```

The FilesystemContractTest can now be reused to test any class that wants to be a Filesystem. This way we can ensure that not only does the class implement the interface, but also behaves the way a Filesystem should. Because that is what the consumer of an interface depends on.

Let’s say that we want our Filesystem to throw an exception if a file with the same name already exists. Let’s add a test for this to our FilesystemContractTest.

```plain text
/** @test */ public function throwExceptionIfFileAlreadyExists(): void { $this->expectException(FileAlreadyExistsException::class); $this->expectExceptionMessage("The file '::filename::' already exists"); $filesystem = $this->getInstance(); $filesystem->storeFile( new MockFile('::contents::'), '::filename::' ); // Trying to store another file with the same name should blow up. $filesystem->storeFile( new MockFile('::contents::'), '::filename::' ); }
```

To test this scenario, we try and save two files with the same name. In this case, the contract states that a FileAreadyExistsException with an appropriate message should be thrown.

Let’s implement this for our InMemoryFilesystem.

```plain text
class InMemoryFilesystem implements Filesystem { private array $files = []; public function storeFile(File $file, string $name): void { if ($this->fileExists($name)) { throw new FileAlreadyExistsException( "The file '{$name}' already exists" ); } $this->files[$name] = $file; } public function fileExists(string $filename): bool { return isset($this->files[$filename]); } }
```

Sweet.

## Conclusion

There is more to the contract of an interface than just the method signatures. Contract test are incredibly useful to ensure that the behavior of each implementation of an interface behaves consistently with the rest.

That’s it folks, that’s the post. Please reach out to me on Twitter if you have any comments. I’d love to hear about them.
