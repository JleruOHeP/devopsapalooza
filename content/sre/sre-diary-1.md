---
title: "SRE Diary - Weeks 1 to 8"
date: 2025-01-25T22:31:00+01:00
tags: ["sre", "work diary"]
summary: "Combination of old diary posts"
---
I started to put the diary entries together in October 2022, and managed to post 8 of the weeks. Unfortunately, I somehow lost the source code for the last 4 of them somehow. So to rebuild it, I decided to wrap them all into 1 post now, for historical reasons.

## Week 1 - after vacation
After my last vacation, I felt very energetic and decided to revive this website (as you probably noticed). And one of the ideas I wanted to try is to do some sort of a diary. I want to log my work and exciting topics throughout the week and post it here. I believe it can do some good :)

### Monday
	* So good to have email filters - I sorted >300 emails in 30 minutes!
	* Sync meeting about work - quick overview of plans.
	* Story refinement meeting.
	* Work - completed PR for security improvement and implemented the same stuff for other microservices.

### Tuesday
	* When technology is tested and implemented in one service - it is easy to replicate it in other services.
	* Good to stricken security rules correctly. I found that we had some endpoints not protected 
	as initially designed by rewriting the security configuration and implementing it in 
	a standard way for all microservices.
	* Meeting to define a plan for the following PI (product iteration) - the main goal is 
	to keep every interested party in a loop and manage expectations.
	* An issue with Microsoft and uBlock Origin extension. I helped another colleague create 
	his first support ticket and then posted the investigation results to the Teams channel, 
	so everyone could have it implemented.
	* Presented how to use cost control in azure.

### Wednesday
	* Sync with a colleague about one of the stories - to correctly split it 
	by frontend/backend tasks.
	* Worked on the backend part of the story - first implemented a basic endpoint 
	to share it (first git commit to a branch), then added actual logic, then tests, 
	and finally refactored the code to be cleaner.
	* Systems group work to update three applications I am responsible for 
	in the evening, after working hours.

### Thursday
	* Workshop on story splitting for another team to plan work for the upcoming PI.
	* Discussed data model for endpoints and how to combine similar DTOs 
	but different use cases.
	* Started doing a story about the custom domain for FrontDoor - unfortunately, 
	it is a manual process because terraform is not there yet.

### Friday
	* Terraform and DNS records - to make a custom domain for FrontDoor.
	* Play with Terragrunt to deploy shared infrastructure from the pipeline.

## Week 2 - end of sprint
This week was the last week of the sprint, so a bit of extra pressure to finish all the stories and check that everything is stable in our QA environment.

### Monday
	* Story preparation meeting that continued into 1-1 design session of pattern 
    to efficiently process large amount of data on a daily schedule.
	* Discussion of efficiency of terragrunt run-all apply.
    * Monthly service reporting.
    
### Tuesday
	* PR review of a big story.
	* Onboarded new system to identity provider.
	* Small change in Kotlin microservice - working with Kotlin for the first time.

### Wednesday
	* Rewrote one infrastructure repo to use vanilla terraform and pipelines instead of terrugrunt.
	* Story preparation meeting for the next sprint (of another team).
	* Integration between Quarkus microservice and Azure ServiceBus.
    
### Thursday
	* Facilitated sprint results presentation for stakeholders (we have it as a rotation duty).
	* Sprint planning.
	* QA deployment.
	* Worked on front-end story to calculate and display coverage area of a tilted sensor. 
    
### Friday
	* Finished working on Quarkus-Service Bus integration.

## Week 3 - dev work
This week felt nice. Sometime it is great to mostly focus on development and not architecture or plannings. Development has the fastest feedback loop. You write your code and run unit tests locally and see the results right away!  

### Monday 
	* Refactoring of the Kotlin microservice
	* Story preparation for the next Sprint
	* Architectural meeting to discuss what security practices we want to implement to harden K8S cluster

### Tuesday
	* Prod deployment
	* Reorganize repositories between different projects and update their pipeline triggers
	* Worked on identity provider automation 

### Wednesday
	* Worked on JPA classes, Adapters and Services for a new microservice
	* Added FrontDoor routes for a new microservice

### Thursday
	* Worked on Java Mappers
	* Planned work for Systems group for the next Product Increment (5 sprints)
	* Presented Authentication mechanisms for Azure (Access Policies vs Connection Strings vs SAS tokens)

### Friday
	* Worked on Azure Active Directory authentication with OIDC for Hashicorp Vault
    * Prepared for a feedback round 

## Week 4 - more dev work
Continuation of work on a new microservice. A lot of development and syncronization between team. Nice effort from everyone to make it finished by the end of the sprint!

And we started the next one - this will be IP sprint, focused on improvements, tech debt and Non Functional requirements. A lot of fun coming!

### Monday
    * Meetings day
    * Feedback circle
    * Story refinement
    
### Tuesday
	* Worked on Java app to manage dynamic Azure resources
	
### Wednesday
	* Continue working on Java microservice to manage Azure
  
### Thursday
	* Sprint day
	* QA Deployment
	* Configured Hashicorp Vault JWT authentication using Azure tokens
  
### Friday
	* Mini-hackaton to improve one of the microservices

## Week 5 - terraform and pipelines refactorings

As I said last time, this sprint is a IP sprint. And on top of that it is the last sprint of PI. That means at least one extra day to plan the next one.

### Monday
	* Story refinement
	* Meeting to discuss team velocity and approach to planning
	* Refactor Azure DevOps Release pipelines to yaml-based multistaged pipelines for some of the older repos
### Tuesday
	* Prod deployment
	* Terraform code refactoring - deprecated azurerm resources, introduction of shared variables
	* Interview with a candidate for a chapter lead position
### Wednesday
	* PI planning day
### Thursday
	* Knowledge sharing day between different product teams
### Friday
	* Improve CORS setup of our microservices to allow access to openapi specs

## Week 6 - low morale
This week I had to do the deployment out of the schedule, because the person who was responsible for it left the company. It also lowered the team morale quite a bit. But Hopefully everything will be fine!

### Monday
	* Story refinement
	* Small front end story for Angular - to conditionally hide mat tab headers 
### Tuesday
	* ADR review - provision AKS resources using terraform or helm directly
	* Presentation on how to use Azure App Service Log Stream
	* Monthly service reporting	
### Wednesday
	* Setup of a new project: repo, pipeline and basic infrastructure  
### Thursday
	* Sprint planning day
	* QA deployment  
### Friday
	* Development work on Java microservice - implementation of the hexagonal architecture
	* Team building event

## Week 7 - Salesforce integration
This week was all about Salesforce integration. We are developing a new licensing system and want to bill customers correctly.

### Monday
	* Story refinement
	* Worked on Salesforce integration
### Tuesday
	* Prod deployment
	* Worked on Azure storage container blob listing defect in code.
	* Previous PI results presentation	
### Wednesday
	* Worked on Salesforce integration 
### Thursday
	* Worked on Salesforce integration
### Friday
	* Worked on Salesforce integration

## Week 8 - Murphy’s law
Finishing touches for the Salesforce integration. It is the last week of the sprint, so the demo is planned with a lot of interested stakeholders participating.

### Monday
	* Story refinement
	* Worked on Salesforce integration
### Tuesday
	* Alerts and monitoring for the Salesworce service
	* Scheduler function to trigger license reporting on a monthly basis
### Wednesday
	* Finishing touches and testing of the SF integration - works as expected!
### Thursday
	* Sprint day
	* SF integration did not work during the demo. Murphys law, I guess.
	* Fixing up issues with integration
### Friday
	* Infrastructure work on AKS cluster for a new set of services

And here I decided to slow down with posts because of no responses. Will see how it will work out this time :)