---
title: "Just a simple website"
date: 2022-01-16T18:19:44+01:00
tags: []
---
We are software developers, and building a new website is not a challenging task. My back-end preference is dotnet; the front-end might be React or Angular; webpack to build it; all the source code in GitHub, and use GitHub actions to automate all deployment activities to AWS. Oh, and terraform for infrastructure.

But then comes a question - is there a better way? Of course, there is. I hope at least… There are heaps of various tools in the wild, and it might be super beneficial to explore them. Not only will it expand your horizons, but it potentially simplifies the task right in front of you.

I was very used to complicated systems, CMS platforms, heavy backends, and fancy frontends in my day-to-day job. Do you really need it for a simple website or your new hobby project? Probably not.

Take this website as an example. Obviously, I did not want to use Sitecore or Umbraco to manage content, and initially, I fell back to my default Single-Page-Application template with Vue.JS for personal stuff. I even updated my previous version of it to use version 3 (and immensely enjoyed it). And then I started to ask myself - how am I going to manage content? The first solution was to create multiple components for each separate page. I did not like all the boilerplate code around it. Then there was an option to load the html content into a page dynamically. It was fun to implement, sure, but it felt a bit overkill.

The right question to ask was, “what am I getting from Vue?”. And I think - nothing. It is excellent for real front-end applications, but for a simple list of articles - nah. And after 5 minutes of googling, I decided to use Hugo instead. It was no less entertaining to understand how to use it, speed up site load time, and simplified content creation.