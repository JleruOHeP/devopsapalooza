---
title: "My Terraform Deployment Pipeline"
date: 2022-10-05T18:50:53+02:00
tags: ["devops", "terraform"]
summary: "Azure DevOps pipeline template for infrastructure project"
---
Recently I had to work on a small shared infrastructure project, and for the initial implementation of it, we as a team decided to try out terragrunt. 

It was nice and easy (maybe just slightly unfamiliar to have .hcl files) in the beginning. Repository had 2 modules and I applied both of them using `terragrunt run-all apply`. For the first rollout it was great, but then I asked myself how it will look like in the future. It would be nice to have an approval gate for PROD environment where someone would be able to review exact plan that will be applied soon. And then I found this discussion where it is stated that run-all is not the best practice (https://github.com/gruntwork-io/terragrunt/issues/720#issuecomment-497888756), that there might be issues with inter-depencencies and then that it would not be possible to pass plan to multiple modules at the same time. So if I wanted to keep terragrunt, then deployment pipeline will be split into separate stages for each module - quite an overhead for my taste. By the way, if you know what is the better approach, please let me know!

So instead, I rolled back this code back to my preferred terraform template.

For the project structure, it will be a classical one with 2 additions:
```
modules
| - moduleA
    | - main.tf
    | - variables.tf
    | - output.tf
| - moduleB 
variables
| - dev.tfvars.json
| - qa.tfvars.json
| - prod.tfvars.json
configs
| - dev.config
| - qa.config
main.tf
variables.tf
```

There is an obvious drawback of code duplication here - it will be required to copy variables from the root module to submodules. And the state file can become quite big, so the refresh could take longer than preferred. Alas, this approach is not ideal too…

Two additional folders (variables and configs) have backend configs to init terraform locally (using -backend-config=PATH option) and environment specific non-secret variables that could be passed to plan command (-var-file=FILENAME). This way I try to keep everything checked into git.

Now lets move to pipeline template. I will give an example from Azure DevOps deploying to Azure.


```yaml
trigger:
  branches:
    include:
    - '*'
  paths:
    exclude:
    - .gitignore
    - README.md
 
variables:
  releaseBranchName: 'main'
  isReleaseBranch: $[eq(variables['Build.SourceBranchName'], variables['releaseBranchName'])]
  terraformVersion: 1.3.1
  azureSubscription: 'ARM-Connection - Shared'
  stateRG: 'infra-shared-rg'
  stateStorage: 'infrasharedsa'
  stateContainer: 'tfstate'
  stateFile: 'infra-shared.tfstate'
 
resources:
  - repo: self
 
pool:
  vmImage: ubuntu-latest
 
stages:
  - stage: Build
    jobs:
      - job: validateTerraform
        steps:
          - task: TerraformInstaller@0
            displayName: 'Install Terraform'
            inputs:
              terraformVersion: $(terraformVersion)
          - task: TerraformTaskV3@3
            displayName: 'Initialize Terraform'
            inputs:
              provider: 'azurerm'
              command: 'init'
              backendServiceArm: $(azureSubscription)
              backendAzureRmResourceGroupName: $(stateRG)
              backendAzureRmStorageAccountName: $(stateStorage)
              backendAzureRmContainerName: $(stateContainer)
              backendAzureRmKey: $(stateFile)
          - task: TerraformTaskV3@3
            displayName: 'Validate Infrastructure (Terraform)'
            inputs:
              provider: 'azurerm'
              command: 'validate'
              workingDirectory: '$(System.DefaultWorkingDirectory)'
 
  - stage: Deploy    
    condition: and(succeeded(), eq(variables.isReleaseBranch, 'true'))
    jobs:
      - job: plan
        steps:
          - task: TerraformInstaller@0
            displayName: Install Terraform
            inputs:
              terraformVersion: $(terraformVersion)
          - task: TerraformTaskV3@3
            displayName: 'Initialize Terraform'
            inputs:
              provider: 'azurerm'
              command: 'init'
              backendServiceArm: $(azureSubscription)
              backendAzureRmResourceGroupName: $(stateRG)
              backendAzureRmStorageAccountName: $(stateStorage)
              backendAzureRmContainerName: $(stateContainer)
              backendAzureRmKey: $(stateFile)
          - task: TerraformTaskV3@3
            displayName: 'Plan Terraform'
            inputs:
              provider: 'azurerm'
              command: 'plan'
              commandOptions: '-var-file=variables/shared.tfvars.json -out tfplan -no-color'
              environmentServiceNameAzureRM: $(azureSubscription)
          - publish: tfplan
            displayName: 'Publish plan'
            artifact: tfplan-shared
 
      - job: waitForValidation
        dependsOn: plan
        displayName: Wait for external validation
        pool: server
        timeoutInMinutes: 1440
        steps:
          - task: ManualValidation@0
            timeoutInMinutes: 1440 # task times out in 1 day
            inputs:
              instructions: 'Please validate the output from the plan step'
              onTimeout: 'reject'
 
      - deployment: apply
        displayName: Apply Terraform
        environment: SHARED
        dependsOn: waitForValidation
        pool:
          vmImage: 'ubuntu-latest'
        strategy:
          runOnce:
            deploy:
              steps:
                - checkout: self
                - task: TerraformInstaller@0
                  displayName: 'Install Terraform'
                  inputs:
                    terraformVersion: $(terraformVersion)
                - task: TerraformTaskV3@3
                  displayName: 'Initialize Terraform'
                  inputs:
                    provider: 'azurerm'
                    command: 'init'
                    backendServiceArm: $(azureSubscription)
                    backendAzureRmResourceGroupName: $(stateRG)
                    backendAzureRmStorageAccountName: $(stateStorage)
                    backendAzureRmContainerName: $(stateContainer)
                    backendAzureRmKey: $(stateFile)
                - task: DownloadPipelineArtifact@2
                  displayName: 'Download tfplan'
                  inputs:
                    artifact: tfplan-shared
                    path: '$(Build.SourcesDirectory)'
                - task: TerraformTaskV3@3
                  displayName: 'Apply Terraform'
                  inputs:
                    command: apply
                    commandOptions: '$(Pipeline.Workspace)/tfplan-shared/tfplan'
                    environmentServiceNameAzureRM: $(azureSubscription)
```

In this pipeline I first validate terraform on every commit, and then after the PR was merged into main branch, it will also run deployment into SHARED environment (which might have approval gates configured too!). First step of the deployment is to generate the tfplan based on the current state and publish it as pipeline artefact. After this there will be a manual validation step, where someone will need to review the plan and accept it. Next the generated plan will be downloaded for the next job and terraform will apply that pacific plan. So if infrastructure was changed between plan and apply, it would not be changed to something that was not reviewed, but instead it will fail and deployment team will be notified. I like that more than blid applies :)

This pipeline works for all the changes and will try to apply all modules at the same time and save the state to a single file. It has its pros and cons, but for now I like it more than terragrunt.
