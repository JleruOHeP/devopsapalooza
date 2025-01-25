---
title: "Azure Governance"
date: 2022-10-27T20:12:39+02:00
tags: ["architect", "documentation"]
summary: "Temaplte with explanations for a “Azure Governance” document"
---
It does not matter if you work in a small startup or in a huge corporation. Managing cloud resources requires a certain level of discipline, the more the better. To structure it nicely, you can introduce governance documentation. It will help with cowboy-style access to production resources, unexplainable bills, and the overall security of your systems. I’ll take Azure as an example and outline a template that could be used. I will also assume that this document is somewhat live - you do not write it and then forget; it is a handy tool. To make it shareable, hosting it somewhere like confluence would make sense.

### Azure tenants and subscriptions
The first part of this document can be treated as a reference book for your setup. Describe what tenants you have. Most likely, there will be one tenant in your organization, but for bigger ones, there could be more than that. It will also make sense to mention how many billing accounts there are.

Next, dive a bit deeper - list all subscriptions (existing and planned). Sometimes there is a single subscription to host all resources. Still, to make it cleaner, it is possible to have different subscriptions for different environments and/or even for different products. I prefer one subscription for one environment, for example, PROD, QA, and DEV. To make this list of subscriptions a bit more useful, l have it as a table - with a subscription name and a short summary of its purpose. If there are any plans to change the existing composition, add that too - at least as a todo item or a comment.

Also, let me take a short detour here and mention that I like to have as much interactivity as possible for this kind of document. If there is a subscription mentioned in a list - I will make it as a hyperlink to this subscription overview in the Azure portal; if it is a comment that this subscription is used for “Project X”, then I’ll add a link to confluence space of this project. And even more - if there is an explanation of some used service, I would like to have a link to its documentation (for example, to the Microsoft Docs site).

If you use resource groups as a separation between projects/environments and host everything in a single subscription - add links to those RGs instead.

The final part of this overview of existing resources will be a list of Management Groups - if you have and use them. I sometimes find them a bit overkill; if there are only three subscriptions, it would be too much to set those up. But as usual - you do you 🙂

With those, the reference part of the document is finished, and you can focus on the management rules you want to follow.

### Resource management
One of the sections in those rules could be a “Resource management” section - where you can briefly explain what resource groups are and what their purpose is. Like “In azure, all resources are created inside Resource groups. New resource groups should be created to represent the same life-cycle of the resource and to help group related services together.” It would make sense to add rules around the naming convention for everything. Another thing to add here would be tagging principles. Tags are a super powerful tool that will come again and again with this document.

### Policies
Azure policies could be a subsection of the whole Resource management. You might not need custom policies from the very beginning, but at some stage, you would need a tool that will help you validate that this governance document is actually followed. Policies could be introduced as an audit tool first, but later they could be changed to be enforced. In this subsection list all the policies, their status, and actions that need to be taken.

### Deployment process
Two more essential parts of resource management are resource locks and if you have them. It could be beneficial to lock critical resources from being changed or deleted (the first thing that comes to mind would be a storage account with your terraform state). And if we started on infrastructure as code - you need to define in this document how resources are expected to be created. A good sentence on this matter could be “All of the resources in Azure must be created using automated processes”.

### Scalability, Monitoring and Backups
“Scalability, Monitoring and Backups” is the next section I would add. Here I will outline what is expected from teams. It could be some sort of a go-live checklist that teams need to complete before their product is accepted into the world. It would be good to have a generic approach on how to tackle disaster recovery, status, and performance problems described here as well.

### Security
The perfect continuation of this section is “Security”. Describe what tools are used for security checks. What level of security is implemented for different environments, and what is the expected level for production? Mention if regular external security audits are being done by some third-party company, when was the last one and how you process the findings. A link to the risk registry will be helpful.

### Cost management
My next section would be “Cost management”. After external audits, it will be a breather. I suggest you have budgets defined in Azure based on tags. And also have alerts enabled. For cost reporting, it would make sense to use a cost analysis tool and split the overall cost by tags and resource groups.

### Access management
The last thing that is required to be in the governance document is “Access management”. Describe how access is controlled - what roles and groups are used. It is a good practice to never assign roles to individual users and instead always use groups. You want to move towards the least privileged principle, so users will have as low-level access as possible, with as small scope as possible (you want to avoid granting Owner permission over the whole subscription). The development team needs to be able to freely work in a DEV environment but only monitor PROD. To fix incidents, it is possible to either use Privileged Identity Management (Azure AD P2 feature) or have a separate “support” group with elevated access. To ensure the least privileged access principle, you will perform regular access reviews based on groups and remove roles/people that are no longer needed. Describe this process.

Ok, I hope this could be helpful to you. You can copy this post, keep the headers only and replace my explanations with real “meat”.