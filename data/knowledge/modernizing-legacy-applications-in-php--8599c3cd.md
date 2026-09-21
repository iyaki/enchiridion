---
title: "Modernizing Legacy Applications in PHP"
notion_id: 8599c3cd-5cc9-44bd-bc0c-9ff1ed460168
notion_url: https://app.notion.com/p/Modernizing-Legacy-Applications-in-PHP-8599c3cd5cc944bdbc0c9ff1ed460168
last_edited: 2026-09-21T17:10:00.000Z
source_url: https://leanpub.com/mlaphp
tags: ["English", "System Design / Software Architecture", "PHP", "Book"]
---
Get your code under control in a series of small, specific steps.

Minimum price

Free!

$34.99

PDF

EPUB

4,139

Readers

230

Pages

58,569Words

## About

## "You will breeze through your code like the wind. It will be autoloaded, dependency-injected, unit-tested, layer-separated, and front-controlled."

Is your legacy PHP application composed of page scripts placed directly in the document root of the web server? Do your page scripts, along with any other classes and functions, combine the concerns of model, view, and controller into the same scope? Is the majority of the logical flow incorporated as include files and global functions rather than class methods? If so, you already know that the wide use of global variables means that making a change in one place leads to unexpected consequences somewhere else. These and other factors make it overly difficult and expensive for you to add features and fix bugs. Working with this legacy application feels like dragging your feet through mud.

But it doesn't have to be that way! This book will show you how to modernize your application by extracting and replacing its legacy artifacts. We will use a step-by-step approach, moving slowly and methodically, to improve your application from the ground up. Moreover, we will keep your application running the whole time. Each completed step in the process will keep your codebase fully operational with higher quality. When we are done, you will be able to breeze through your code like the wind. Your code will be autoloaded, dependency-injected, unit-tested, layer-separated, and front-controlled.

Please note that this book is about modernizing in terms of practice and technique, and not in terms of tools. We are not going to discuss the latest, hottest frameworks or libraries. With the exception of testing systems like PHPUnit, and one or two standalone third-party libraries, the book does not advocate adding third-party code to your existing legacy application. Most of the very limited code we do add to your application is specific to this book. We will be improving ourselves as programmers, as well as improving the quality of our legacy application.

If you feel overwhelmed by a legacy application, "Modernizing Legacy Applications in PHP" is the book for you. If you prefer a paper copy, you can purchase one at Amazon.com.

If you're still on the fence, the video "It Was Like That When I Got Here" (embedded above on this page) outlines the first few chapters for free, and you can read some reviews of the book on the feedback page. You can also read the 4.5-star review from SitePoint.

Be sure to follow the book on Twitter @mlaphp, and tell all your friends!

## Bundle

- Pricing$49.95Minimum priceBought separately$64.98Suggested price$49.95

## Author

Paul M. Jones

Paul M. Jones is an internationally recognized PHP expert who has worked as everything from junior developer to VP of Engineering in all kinds of organizations (corporate, military, non-profit, educational, medical, and others). Paul's latest open-source project is the Atlas Persistence Framework for PHP. Among his other accomplishments, Paul is the lead developer on Aura for PHP and Solar Framework, and the creator of the Savant template system. He has authored a series of authoritative benchmarks on dynamic framework performance, and was a founding contributor to the Zend Framework (the DB, DB_Table, and View components). Paul was a founding member of the PHP Framework Interoperability Group, where he shepherded the PSR-1 and PSR-2 recommendations, and was the primary author on the PSR-4 autoloader recommendation. He was also a member of the Zend PHP 5.3 Certification education advisory board. He blogs at paul-m-jones.com. In a previous career, Paul was an operations intelligence specialist for the US Air Force, and enjoys putting .308 holes in targets at 400 yards.

The Leanpub Podcast

## Podcast

## Translations

## Testimonials

- [This book] gives developers an easy-to-follow, practical and powerful process to bring their applications up to a modern baseline. As I followed the exercises in the book, my questions almost seemed to be anticipated and answered before the chapter was over. If you are in the midst of a legacy refactor or you find yourself in a state of despair caused by the code you have to work on, I implore you: read Modernizing Legacy Applications in PHP. Change is really possible!jblotusEasy-to-follow, practical, and powerful
- This book helped me slay a 300k line of code giant and has allowed me to break out my shell. It has been my best PHP book purchase yet and I continue to improve the code base. The code base is now 80% unit tested and I am now implementing continuous integration.Chris SmithKiss your technical debt goodbye
- Reading through the book, it feels like you're pair programming with the author. I'm at the keyboard, driving, and the author is navigating, telling me where to go and what to do next. Each step is practical, self-contained and moves you closer to the end goal you seek: maintainable code. I *highly* recommend this book. Even if you're a seasoned developer (I've been writing code professionally more than 20 years), you will benefit from Paul's approach and detailed documentation of the process.jclermontA 12-step program for PHP devs
- This is one of those books that PHP developers from all skill levels will be able to glean value from, and I know after just a single read-through that it will be an oft-referenced resource when I need to convert my old legacy-based procedural code into something cleaner, object-oriented, and testable. This is a very thorough guide to ... getting developers stuck with legacy codebases up to speed with the tools that are available to them. I will recommend this to anyone who will listen.J. Michael WardSuperb

## Contents

- Foreword
- Preface and Acknowledgments
- 1. Legacy Applications The Typical PHP Application Rewrite or Refactor? Legacy Frameworks Review and Next Steps
- 2. Prerequisites Revision Control PHP Version Editor/IDE Style Guide Test Suite Review and Next Steps
- 3. Implement An Autoloader PSR-0 A Single Location For Classes Add Autoloader Code Common Questions Review and Next Steps
- 4. Consolidate Classes and Functions Consolidate Class Files Consolidate Functions Into Class Files Common Questions Review and Next Steps
- 5. Replace global With Dependency Injection Global Dependencies The Replacement Process Common Questions Review and Next Steps
- 6. Replace new With Dependency Injection Embedded Instantiation The Replacement Process Common Questions Review and Next Steps
- 7. Write Tests Fighting Test Resistance Setting Up A Test Suite Common Questions Review and Next Steps
- 8. Extract SQL Statements To Gateways Embedded SQL Statements The Extraction Process Common Questions Review and Next Steps
- 9. Extract Domain Logic To Transactions Embedded Domain Logic Domain Logic Patterns The Extraction Process Common Questions Review and Next Steps
- 10. Extract Presentation Logic To View Files Embedded Presentation Logic The Extraction Process Common Questions Review and Next Steps
- 11. Extract Action Logic To Controllers Embedded Action Logic The Extraction Process Common Questions Review and Next Steps
- 12. Replace Includes In Classes Embedded include Calls The Replacement Process Common Questions Review and Next Steps
- 13. Separate Public And Non-Public Resources Intermingled Resources The Separation Process Common Questions Review and Next Steps
- 14. Decouple URL Paths From File Paths Coupled Paths The Decoupling Process Common Questions Review and Next Steps
- 15. Remove Repeated Logic In Page Scripts Repeated Logic The Removal Provess Common Questions Review and Next Steps
- 16. Add A Dependency Injection Container What Is A Dependency Injection Container? Adding A DI Container Common Questions Review and Next Steps
- 17. Conclusion Opportunities for Improvement Conversion to Framework Review and Next Steps
- Appendix A: Typical Legacy Page Script
- Appendix B: Code Before Gateways
- Appendix C: Code After Gateways
- Appendix D: Code After Transaction Scripts
- Appendix E: Code Before Collecting Presentation Logic
- Appendix F: Code After Collecting Presentation Logic
- Appendix G: Code After Response View File
- Appendix H: Code After Controller Rearrangement
- Appendix I: Code After Controller Extraction
- Appendix J: Code After Controller Dependency Injection
- Colophon
- About the Author

## Also by the Author

- Solving The N+1 Problem In PHP
- Modernização de Aplicações Legadas em PHP
- Aura Framework v2
- Modernizzare Applicazioni Legacy in PHP

## The Leanpub 60 Day 100% Happiness Guarantee

Within 60 days of purchase you can get a 100% refund on any Leanpub purchase, in two clicks.

See full terms...

## Earn $8 on a $10 Purchase, and $16 on a $20 Purchase

We pay 80% royalties on purchases of $7.99 or more, and 80% royalties minus a 50 cent flat fee on purchases between $0.99 and $7.98. You earn $8 on a $10 sale, and $16 on a $20 sale. So, if we sell 5000 non-refunded copies of your book for $20, you'll earn $80,000.

(Yes, some authors have already earned much more than that on Leanpub.)

In fact, authors have earned over $15 million writing, publishing and selling on Leanpub.

Learn more about writing on Leanpub

## Free Updates. DRM Free.

If you buy a Leanpub book, you get free updates for as long as the author updates the book! Many authors use Leanpub to publish their books in-progress, while they are writing them. All readers get free updates, regardless of when they bought the book or how much they paid (including free).

Most Leanpub books are available in PDF (for computers) and EPUB (for phones, tablets and Kindle). The formats that a book includes are shown at the top right corner of this page.

Finally, Leanpub books don't have any DRM copy-protection nonsense, so you can easily read them on any supported device.

Learn more about Leanpub's ebook formats and where to read them

## Write and Publish on Leanpub

You can use Leanpub to easily write, publish and sell in-progress and completed ebooks and online courses!

Leanpub is a powerful platform for serious authors, combining a simple, elegant writing and publishing workflow with a store focused on selling in-progress ebooks.

Leanpub is a magical typewriter for authors: just write in plain text, and to publish your ebook, just click a button. (Or, if you are producing your ebook your own way, you can even upload your own PDF and/or EPUB files and then publish with one click!) It really is that easy.

Learn more about writing on Leanpub
