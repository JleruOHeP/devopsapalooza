---
title: "Diary - Week 1"
date: 2022-10-01T18:44:53+02:00
tags: ["sre", "work diary"]
summary: "Week 1 - after vacation"
---
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
