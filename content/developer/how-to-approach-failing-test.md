---
title: "How to Approach Failing Test"
date: 2022-09-27T20:50:08+02:00
tags: ["to be updated", "developer"]
summary: "Switch from failing integration test to working unit test"
---

Today at work, I had an interesting task. We wanted to update our Spring Boot security configuration to use the newer and standardized version across multiple microservices. But the first step is to integrate it with one microservice and only then propagate it to the rest. The actual implementation was easy; troubles came from testing.

I needed to add MockMvc to test requests, and I wanted it auto-injected into my test class. Huh, apparently, there is an issue with TestContainers in Java when you use both @Autowired and @MockBean or @SpyBean. ApplicationContext is unavailable in @SpringBootTest… sometimes. It could be that one of the tests for the security configuration would fail, where we test that requests to certain endpoints should be rejected without proper authentication. Or it could be one of the database integration tests. Super annoying!

So first thing first - Google and Stackoverflow. A similar issue was asked multiple times on Stackoverflow, but I have not found the exact match. There is never an exact match; otherwise, it would be boring work, right? And many search results were even for spring 1.5; if I tried to implement some of those options blindly, it would be an absolute mess! Still, there were some good pointers. For example, someone suggested creating a MockMvc manually instead of auto-wiring it and injecting required security filter chains into it.

That made me think - do we really need a full integration test with a @SpringBootTest? First of all, those are slow. And not working. A better solution is to use @WebMvcTest with @ContextConfiguration - it is faster and not breaking other integration tests. 

And everything would be perfect if not one small detail. With that setup, the actuator health check endpoint was not available yet. In the test, the request returned 404; but we wanted to ensure that this endpoint was available without authorization. Back to the drawing board, I mean google search? Yes! But keep it time limited. We have a “good enough” working solution, and 100% test coverage of all possible scenarios is impossible anyway.  

The lesson I wanted to highlight with this post is to keep everything realistic. Striving for excellence is good, but the working solution right now is better. I want to come back to the actuator issue and will do it - when other more important topics will be finished. Or maybe in my spare time, just for fun. Another thing is that questioning the architecture is always good. You could prove that it is the correct one or maybe find and implement a better one. 
