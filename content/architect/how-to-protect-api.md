---
title: "How to protect APIs"
date: 2022-11-08T20:12:39+02:00
tags: ["architect", "security"]
summary: "API security best practices"
---
It feels like nowadays, almost every product is developed as a web app. With that many APIs in the wild, average quality is expected to go down. But why? And even more importantly - how to make sure that your application is good?

There are a lot of points to consider in regard to quality. For example, you should ensure that your product solves your customers' real problems or that its performance is adequate. Right now, I want to list some of the security best practices. Your customers will appreciate it if the application is secure.

### Authentication
The number one thing you need to implement is authentication. Authentication is the process of determining whether someone or something is who or what it says it is. You need to always use secure authentication techniques like JWT or OAuth to confirm user identity. Here are the most common options:

1. Simple HTTP authentication should never be used as it sends fields without encryption.
2. Basic HTTP Auth. Is not the best as well. The credentials are sent in an HTTP header, making them easily interceptable.
3. Bearer HTTP Auth. With this method, the user must first authenticate using a bearer authentication to access your resources. A username or other identity is often embedded within the token that is later passed to the server: ``Authorization: Bearer <token>``
4. JWT. JSON Web Tokens offer user-level access to API, and it is possible to include different user-specific details and claims in the token. Many Identity providers generate JWT tokens that other services can validate and use.

### Authorization
After we have confirmed the user’s identity, it would be good to ascertain his level of access, or in other words, Authorize access. It is not required to have both, but it would make your application more secure. To authorize requests, you can implement the following:

1. API Keys. It is basically a unique string/code used to access APIs. First, you generate those keys and send them to your users somehow. And later validate if the request is made with a valid key. It will also allow tracking of users' identities and use patterns.
2. OAuth. OAuth is an open authorization standard that allows users to grant third-party access to their web resources without sharing their passwords.

### Inputs and Outputs
The next piece of advice is to validate all inputs. You could specify the correct input format for your functions in your API documentation (openapi specs!). And when you receive any request - first validate the input and only then process it. The best illustration of why will be the Bobby Tables xkcd. On the other end of the process, return only the minimal required data back for each request. It will improve the network throughput, but much more importantly, it could hide possible implementation details from malicious parties. For example, removing any default X-Powered-By and Server headers from your HTTP response would make perfect sense.

### API Gateways
And as the last thing, leys discuss API gateways. I would recommend to always hide your services behind a gateway. API gateways combine both security-related activities and useful business-related operations. Rate limiting, throttling, blocking malicious clients, and the simple possibility of changing routing between services without affecting clients are all characteristics of API gateways.

Rate limiting limits the number of requests a user can make to an API within a given period. It is a protective measure to prevent overload of the API. API rate limiting is always implemented on the user level. At the same time, API throttling is implemented on the gateway level and artificially slows down the response time of API if too many calls are being made in a given time. Those techniques can help in the prevention of denial-of-service (DoS) attacks. Unfortunately, they are much less effective in preventing distributed denial-of-service (DDoS) attacks.

Very often, API gateways have WAF (Web Application Firewall) available as well. It automatically protects your applications from common web exploits. If you know specific IP ranges you do not want to allow to your resources, it is possible to add policies that will blacklist those too. Stop your competitors from abusing your servers!

An API gateway can also route requests as you want. It is possible to have the same API exposed through multiple instances of the app. Or, if you’re going to update the version of the app - you can do a rolling deployment with new functionality available only to a percent of the whole user base.