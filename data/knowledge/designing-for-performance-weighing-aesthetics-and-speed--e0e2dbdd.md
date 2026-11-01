---
title: "Designing for Performance - Weighing Aesthetics and Speed"
notion_id: e0e2dbdd-ae2d-4a65-a32a-d92e00423983
notion_url: https://app.notion.com/p/Designing-for-Performance-Weighing-Aesthetics-and-Speed-e0e2dbddae2d4a65a32ad92e00423983
last_edited: 2023-09-13T14:44:00.000Z
source_url: https://designingforperformance.com/
tags: ["Web Development", "Book", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

As a web designer, you encounter tough choices when it comes to weighing aesthetics and performance. Good content, layout, images, and interactivity are essential for engaging your audience, and each of these elements have an enormous impact on page load time and the end-user experience. In this practical book, [Lara Callender Hogan](https://larahogan.me/) helps you approach projects with page speed in mind, showing you how to test and benchmark which design choices are most critical.

Lara is donating all of the proceeds from the book to various charities focused on supporting marginalized people in tech, including [Girl Develop It](https://www.girldevelopit.com/), [Women Who Code](https://www.womenwhocode.com/), [Black Girls Code](https://www.blackgirlscode.com/), [#YesWeCode](https://www.yeswecode.org/), [Latinas in STEM](http://www.latinasinstem.com/), [Hack the Hood](https://www.hackthehood.org/), and DonorsChoose projects like [Growing "Girls with Gadgets"!](https://www.donorschoose.org/project/growing-girls-with-gadgets/1468056/) and [Girls Can Code, Too!](https://www.donorschoose.org/project/girls-can-code-too/1457602/). So if you enjoy the book, please [consider buying it](http://shop.oreilly.com/product/0636920033578.do)!

## Table of Contents

1. [Praise for ](https://designingforperformance.com/praise/)[_Designing for Performance_](https://designingforperformance.com/praise/)[Dedication](https://designingforperformance.com/dedication/)
2. [Foreword by Randy J. Hunt](https://designingforperformance.com/foreword-hunt/)
3. [Preface](https://designingforperformance.com/preface/)
4. 1. [Performance Is User Experience](https://designingforperformance.com/performance-is-ux/) 
5. [Impact on Your Brand](https://designingforperformance.com/performance-is-ux/#impact-on-your-brand) 
6. [Returning Users](https://designingforperformance.com/performance-is-ux/#returning-users)
7. [Search Engine Rankings](https://designingforperformance.com/performance-is-ux/#search-engine-rankings)
8. [Impact on Mobile Users](https://designingforperformance.com/performance-is-ux/#impact-on-mobile-users) 
9. [Mobile Networks](https://designingforperformance.com/performance-is-ux/#mobile-networks)
10. [Mobile Usage Patterns](https://designingforperformance.com/performance-is-ux/#mobile-usage-patterns)
11. [Mobile Hardware](https://designingforperformance.com/performance-is-ux/#mobile-hardware)
12. [Designers’ Impact on Performance](https://designingforperformance.com/performance-is-ux/#designers-impact-on-performance)
13. 2. [The Basics of Page Speed](https://designingforperformance.com/basics-of-page-speed/) 
14. [How Browsers Render Content](https://designingforperformance.com/basics-of-page-speed/#how-browsers-render-content) 
15. [Requests](https://designingforperformance.com/basics-of-page-speed/#requests)
16. [Connections](https://designingforperformance.com/basics-of-page-speed/#connections)
17. [Page Weight](https://designingforperformance.com/basics-of-page-speed/#page-weight)
18. [Perceived Performance](https://designingforperformance.com/basics-of-page-speed/#perceived-performance) 
19. [Critical Rendering Path](https://designingforperformance.com/basics-of-page-speed/#critical-rendering-path)
20. [Jank](https://designingforperformance.com/basics-of-page-speed/#jank)
21. [Other Impacts on Page Speed](https://designingforperformance.com/basics-of-page-speed/#other-impacts-on-page-speed) 
22. [Geography](https://designingforperformance.com/basics-of-page-speed/#geography)
23. [Network](https://designingforperformance.com/basics-of-page-speed/#network)
24. [Browser](https://designingforperformance.com/basics-of-page-speed/#browser)
25. 3. [Optimizing Images](https://designingforperformance.com/optimizing-images/) 
26. [Choosing an Image Format](https://designingforperformance.com/optimizing-images/#choosing-an-image-format) 
27. [JPEG](https://designingforperformance.com/optimizing-images/#jpeg)
28. [GIF](https://designingforperformance.com/optimizing-images/#gif)
29. [PNG](https://designingforperformance.com/optimizing-images/#png)
30. [Additional Compression](https://designingforperformance.com/optimizing-images/#additional-compression)
31. [Replacing Image Requests](https://designingforperformance.com/optimizing-images/#replacing-image-requests) 
32. [Sprites](https://designingforperformance.com/optimizing-images/#sprites)
33. [CSS3](https://designingforperformance.com/optimizing-images/#css3)
34. [Data URIs and Base64-Encoding Images](https://designingforperformance.com/optimizing-images/#data-uris-and-base64-encoding-images)
35. [SVG](https://designingforperformance.com/optimizing-images/#svg)
36. [Image Planning and Iterating](https://designingforperformance.com/optimizing-images/#image-planning-and-iterating) 
37. [Schedule Routine Checks](https://designingforperformance.com/optimizing-images/#schedule-routine-checks)
38. [Create Style Guides](https://designingforperformance.com/optimizing-images/#create-style-guides)
39. [Mentor Other Image Creators](https://designingforperformance.com/optimizing-images/#mentor-other-image-creators)
40. 4. [Optimizing Markup and Styles](https://designingforperformance.com/optimizing-markup-and-styles/) 
41. [Cleaning Your HTML](https://designingforperformance.com/optimizing-markup-and-styles/#cleaning-your-html) 
42. [Divitis](https://designingforperformance.com/optimizing-markup-and-styles/#divitis)
43. [Semantics](https://designingforperformance.com/optimizing-markup-and-styles/#semantics)
44. [Accessibility](https://designingforperformance.com/optimizing-markup-and-styles/#accessibility)
45. [Frameworks and Grids](https://designingforperformance.com/optimizing-markup-and-styles/#frameworks-and-grids)
46. [Cleaning Your CSS](https://designingforperformance.com/optimizing-markup-and-styles/#cleaning-your-css) 
47. [Unused Styles](https://designingforperformance.com/optimizing-markup-and-styles/#unused-styles)
48. [Combine and Condense Styles](https://designingforperformance.com/optimizing-markup-and-styles/#combine-and-condense-styles)
49. [Clean Stylesheet Images](https://designingforperformance.com/optimizing-markup-and-styles/#clean-stylesheet-images)
50. [Remove Specificity](https://designingforperformance.com/optimizing-markup-and-styles/#remove-specificity)
51. [Optimizing Web Fonts](https://designingforperformance.com/optimizing-markup-and-styles/#optimizing-web-fonts)
52. [Creating Repurposable Markup](https://designingforperformance.com/optimizing-markup-and-styles/#creating-repurposable-markup) 
53. [Style Guides](https://designingforperformance.com/optimizing-markup-and-styles/#style-guides)
54. [Additional Markup Considerations](https://designingforperformance.com/optimizing-markup-and-styles/#additional-markup-concerns) 
55. [CSS and JavaScript Loading](https://designingforperformance.com/optimizing-markup-and-styles/#css-and-javascript-loading)
56. [Minification and gzip](https://designingforperformance.com/optimizing-markup-and-styles/#minification-and-gzip)
57. [Caching Assets](https://designingforperformance.com/optimizing-markup-and-styles/#caching-assets)
58. 5. [Responsive Web Design](https://designingforperformance.com/responsive-web-design/) 
59. [Deliberately Loading Content](https://designingforperformance.com/responsive-web-design/#deliberately-loading-content) 
60. [Images](https://designingforperformance.com/responsive-web-design/#images)
61. [Fonts](https://designingforperformance.com/responsive-web-design/#fonts)
62. [Approaches](https://designingforperformance.com/responsive-web-design/#approaches) 
63. [Project Documentation](https://designingforperformance.com/responsive-web-design/#project-documentation)
64. [Mobile First](https://designingforperformance.com/responsive-web-design/#mobile-first)
65. [Measure Everything](https://designingforperformance.com/responsive-web-design/#measure-everything)
66. 6. [Measuring and Iterating on Performance](https://designingforperformance.com/measuring-and-iterating/) 
67. [Browser Tools](https://designingforperformance.com/measuring-and-iterating/#browser-tools) 
68. [YSlow](https://designingforperformance.com/measuring-and-iterating/#yslow)
69. [Chrome DevTools](https://designingforperformance.com/measuring-and-iterating/#chrome-devtools)
70. [Synthetic Testing](https://designingforperformance.com/measuring-and-iterating/#synthetic-testing)
71. [Real User Monitoring](https://designingforperformance.com/measuring-and-iterating/#real-user-monitoring)
72. [Changes over Time](https://designingforperformance.com/measuring-and-iterating/#changes-over-time)
73. 7. [Weighing Aesthetics and Performance](https://designingforperformance.com/weighing-aesthetics-and-performance/) 
74. [Finding the Balance](https://designingforperformance.com/weighing-aesthetics-and-performance/#finding-the-balance)
75. [Make Performance Part of Your Workflow](https://designingforperformance.com/weighing-aesthetics-and-performance/#make-performance-part-of-your-workflow)
76. [Approach New Designs with a Performance Budget](https://designingforperformance.com/weighing-aesthetics-and-performance/#approach-new-designs-with-a-performance-budget)
77. [Experiment on Designs with Performance in Mind](https://designingforperformance.com/weighing-aesthetics-and-performance/#experiment-on-designs-with-performance-in-mind)
78. 8. [Changing Culture at Your Organization](https://designingforperformance.com/changing-culture/) 
79. [Performance Cops and Janitors](https://designingforperformance.com/changing-culture/#performance-cops-and-janitors)
80. [Upward Management](https://designingforperformance.com/changing-culture/#upward-management) 
81. [Impact on Business Metrics](https://designingforperformance.com/changing-culture/#impact-on-business-metrics)
82. [Experiencing Site Speed](https://designingforperformance.com/changing-culture/#experiencing-site-speed)
83. [Working with Other Designers and Developers](https://designingforperformance.com/changing-culture#working-with-other-designers-and-developers) 
84. [Educating](https://designingforperformance.com/changing-culture/#educating)
85. [Empowering](https://designingforperformance.com/changing-culture/#empowering)

## About the author

[**Lara Callender Hogan**](https://larahogan.me/) is the Senior Engineering Manager of the Performance team at Etsy. Lara previously managed Etsy’s Mobile Web Engineering team. Before joining Etsy, Lara was a User Experience Manager and self-taught frontend developer at a number of startups. She’s been certified as an EMT, owned her own photography business, and co-founded an LGBT wedding website. She also believes it’s important to celebrate career achievements with [donuts](https://larahogan.me/donuts/).

## Colophon

The animal on the cover of _Designing for Performance_ is a tufted coquette (_Lophornis ornatus_), a tiny hummingbird that breeds in eastern Venezuela, Trinidad, Guiana, and northern Brazil.

Also known as the splendid coquette, this hummingbird is so tiny that it can easily be confused with a large bee as it moves from flower to flower. Its red beak has a black tip and is short and straight. The female doesn’t have very flashy plumage, but the male has striking blackspotted and orange-colored feathers that project from the sides of his neck and an orange head crest.

Hummingbirds in general are quite solitary, so the tufted coquette is mostly found alone or in small groups, as it searches for nectar and small insects to feed on.

Many of the animals on O’Reilly covers are endangered; all of them are important to the world. To learn more about how you can help, go to [animals.oreilly.com](http://animals.oreilly.com/).

The cover image is from Wood’s Natural History. The cover fonts are URW Typewriter and Guardian Sans.
