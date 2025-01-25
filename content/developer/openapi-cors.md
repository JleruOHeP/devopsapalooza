---
title: "How to configure openapi specs for Confluence"
date: 2022-10-24T20:12:39+02:00
tags: ["developer", "java", "documentation"]
summary: "openapi CORS snippet"
---
When you develop a WebAPI it is a really good idea to start with design and documentation. That way you will write something that is useful and not just a collection of random endpoints. To help you with that there is an excellent tool (as always :) ) - OpenAPI. It is an API description format for REST APIs. You can first define how your API will be used, validate it and only then start the actual development. Even better, there is also a number of open api generators that will take the spec file and generate all the boilerplate code for you (think about all the controllers and dtos you don’t need to write!). I won’t go into implementation details for this now, but still would highly recommend this approach.

What I want to discuss today is the logical next step. After you develop your application, you want to share the current documentation with your consumers and the future you. If you have a Confluence, it is possible to use [Open API (Swagger) Integration for Confluence](https://toshihiro.atlassian.net/wiki/spaces/OASI/overview?homepageId=229378) macros to automatically read the latest version of the published specification and display it in a familiar swagger UI style. Everything can be setup in a couple of clicks and will just work!

With a tiny issue. Security. Confluence macro runs on one server, and your application is on another. Usually, it is not a good idea to allow access to your resources left and right :) There is a special mechanism to enable it - CORS. It is a mechanism that lets you setup access to resources from domains outside your domain. There are a number of options for how you can configure CORS.

First of all, you can say that you don’t really care about security and allow access to your webapp from anywhere by anyone. It is NOT a good idea, so I won’t spend much time on it.

There is a second not-so-good idea. You can configure your app so that it is accessible from a list of trusted domains and add “https://toshihiro.herokuapp.com” to the list. It is better but still allows too much.

Finally, if you publish your specs as “/openapi.json” it is possible to let anyone read those (because you want to share the knowledge! Expand your client base!) but still restrict access to actual APIs only to the trusted list. That would be my preference. Here is a quick Java snippet that will do it:

```java
@Bean
protected CorsConfigurationSource corsConfigurationSource() {
  final CorsConfiguration corsConfiguration = new CorsConfiguration();
  // configure your default settings based on env. variables or constants
  corsConfiguration.setAllowedOrigins(CORS_ALLOWED_ORIGINS);
  corsConfiguration.setAllowCredentials(true);
  corsConfiguration.addAllowedHeader(CORS_ALLOWED_HEADER);
  corsConfiguration.setAllowedMethods(CORS_ALLOWED_METHODS);
  corsConfiguration.setMaxAge(CORS_MAX_AGE);
  corsConfiguration.addExposedHeader("Content-Disposition");
  corsConfiguration.addExposedHeader("Content-Length");

  final CorsConfiguration openApiCorsConfiguration = new CorsConfiguration();
  openApiCorsConfiguration.setAllowedOrigins(List.of("*"));
  openApiCorsConfiguration.setAllowedMethods(List.of("GET"));

  UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
  source.registerCorsConfiguration("/openapi.json", openApiCorsConfiguration);
  source.registerCorsConfiguration("/**", corsConfiguration);

  return source;
}
```
CORS configurations applied in the same order that they are listed, so the most restrictive one should be the top-most.