---
title: "Production Go-Live Checklist"
date: 2022-01-16T18:20:49+01:00
tags: []
---
Often before you go to production with a new application, you need to make sure that every base is covered. It can be a task for a Change Manager or fall into the SRE team’s caring hands. Here is a quick checklist of what might be required:

## SECURITY
* Application endpoints must be protected with authentication and authorization.

* Penetration testing has been conducted against all publicly exposed endpoints;

  All findings must be addressed or signed off (accepted as a security risk) by the business.

* Access to Production resources is restricted to support team only.

* Web applications are binded to organization production domain.

## QUALITY ASSURANCE
* Application has been tested by a QA resource;

  Any failed findings must be addressed and retested by a QA resource.

  Final test results must be signed off by the business.

* Requirements were signed off by business.

* Performance testing has been conducted with the expected production workload;

  Test results must be signed off by the business.

* All cloud resources are created and configured using infrastructure-as-code (e.g. Terraform).

* Application code is in source control.

* Code build/deployments are automated using CI/CD tool (e.g. Azure DevOps).

## OPERATIONAL READINESS
* Support model has been documented and signed off by the business. This should include details such as:
  - An incident prioritization matrix (signed off by support teams and product owners)
  - Incident escalation
  - After-hours support (if any)
  - Application tiering

* Application is storing its logs in a central location (e.g. Sumologic).

* Application is monitored by Application Performance Monitoring tool (e.g. New Relic).

* Operational procedures have been documented (e.g. a page in Confluence on how to make a change; how to grant a user access; etc.).

* Handover to the BAU support team has been done and signed off by the owner. (For supporting triage and assign inbound issues).

* All operational costs for supporting the application have been clearly defined and signed off by the business. This may include items such as:
  - Cost for Azure resources
  - Cost for support team

* User access to systems is planned and organized (e.g. with an AD/OKTA/Auth0 setup).

* ServiceNow setup (if applicable):
  - All application resources are defined as configuration items in ServiceNow.
  - Resolver group has been created in ServiceNow (if new team);
  - This refers to the team which will be supporting the application.
