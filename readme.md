# Devopsapalooza website

Content for my blog available at https://devopsapalooza.com/

# Infrastructure

Done manually so far - it is:
* S3 bucket 
* Route 53 DNS zone
* CloudFront
* Lambda and Dynamo DB for subscription function

# Deployment

Only site is built and deployed automatically, infrastructure and lambda were developed long time ago and deployed manually. 

# Local development

## Hugo

Site is built with hugo (https://gohugo.io/), so first - install it (current version for the site is 0.142.0)
```
choco install hugo-extended
```

Then you copy config.yaml.tmp into the real config.yaml, replacing the google analytics id and then run

```
hugo server
```
to serve the site locally.

## Subscribe function

Located under lambdas folder. It is a C# code.
