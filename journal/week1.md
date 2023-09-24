# Terraform Beginner Bootcamp 2023  Week 1  — Building Week Journal

[Week-1 Architecture](#week-1-architecture) :cloud: :potted_plant:

[Notes for revision](#notes-for-revision) :syringe: :medal_sports:

[Personal Documentation](#personal-documentation) :memo: :pencil:

[Task Status](#task-status) :star: :face_in_clouds: :partying_face: :white_check_mark:

# Task Status
| Topic | Status |
| --- | --- | 
| [Journal TOC and Major Version](#journal-toc-and-major-version)  | <ul><li> [x] </li></ul> |
| [Restructure Root Module](#restructure-root-module) | <ul><li> [ ] </li></ul> |
| [Terraform Import and Configuration Drift](#terraform-import-and-configuration-drift) | <ul><li> [ ] </li></ul> |
| [Create Terrahouse Module](#create-terrahouse-module) | <ul><li> [ ] </li></ul> |
| [Static Website Hosting](#static-website-hosting) | <ul><li> [ ] </li></ul> |
| [Content Delivery Network](#content-delivery-network) | <ul><li> [ ] </li></ul> |
| [Terraform Data and Content Version](#terraform-data-and-content-version) | <ul><li> [ ] </li></ul> |
| [Invalidate Cache and Local Exec](#invalidate-cache-and-local-exec) | <ul><li> [ ] </li></ul> |
| [Assets Upload and For Each](#assets-upload-and-for-each) | <ul><li> [ ] </li></ul> |
| [Working Through Git Graph Issues](#working-through-git-graph-issues) | <ul><li> [ ] </li></ul> |
| [Project Validation](#project-validation) | <ul><li> [ ] </li></ul> |

# Week 1 Architecture
![week1-architecture](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/a7c8bb05e343f4daa90e3df29ee00f53993461a2/journal/assets/week-1/Week-1-ArchitecturalDiagram.png)

The complete end-to-end architecture of what we will achieve at the end of the bootcamp is available [here](https://lucid.app/lucidchart/e3f15b1a-2211-4ddb-8c95-f144c2504db4/edit?invitationId=inv_0873b3c6-c652-463f-9f2b-fa0f1b420823&page=0_0#) 

_Diagram copyrights: Andrew Brown from ExamPro.co_

-----------------------------------------------------------------------------------------------------
# Notes For Revision


## Root Module Structure

Our root module structure is as follows:

```
PROJECT_ROOT
│
├── main.tf                 # everything else.
├── variables.tf            # stores the structure of input variables
├── terraform.tfvars        # the data of variables we want to load into our terraform project
├── providers.tf            # defined required providers and their configuration
├── outputs.tf              # stores our outputs
└── README.md               # required for root modules
```

[Standard Module Structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure)

Standard Terraform module structure:
#### Root module
- mandatory element
- Terraform files must exist in the root directory of the repository
- primary entrypoint for the module 

#### **README**
- The root module and any nested modules should have README files
- This file should be named `README` or `README.md`. 
- should contain the description of the module and what it should be used for.
- **examples** can be included in a `examples` directory

####  **LICENSE**
-  The license under which this module is available

`main.tf`, `variables.tf`, `outputs.tf` : are the recommended filenames for a minimal module, even if they are empty.


Example of a minimal recommended module following the standard structure

```txt
$ tree minimal-module/
.
├── README.md
├── main.tf
├── variables.tf
├── outputs.tf
```


Complete structure example:

```txt
$ tree complete-module/
.
├── README.md
├── main.tf
├── variables.tf
├── outputs.tf
├── ...
├── modules/
│   ├── nestedA/
│   │   ├── README.md
│   │   ├── variables.tf
│   │   ├── main.tf
│   │   ├── outputs.tf
│   ├── nestedB/
│   ├── .../
├── examples/
│   ├── exampleA/
│   │   ├── main.tf
│   ├── exampleB/
│   ├── .../
```

## Terraform and Input Variables

## Terraform Cloud Variables

In terraform we can set two kind of variables:
1) **Enviroment Variables** - those you would set in your bash terminal eg. AWS credentials
2) **Terraform Variables** - those that you would normally set in your tfvars file

> We can set Terraform Cloud variables to be sensitive so they are not shown visibliy in the UI.

## Terraform Variables
[Terraform Variables](https://developer.hashicorp.com/terraform/language/values/variables)
- similar to shell input variables/function arguments.
- are 
- let you customize aspects of Terraform modules without altering the module's own source code. 
- allows reusing modules across different Terraform configurations (example: same code used for different envs like dev, prod, qa etc.)
- good practice to briefly describe the purpose of each variable using the optional `description` argument:
> For brevity, input variables are often referred to as just "variables" or "Terraform variables"

### Declaring an Input Variable - example
```tf
variable "availability_zone_name" {
  type    = string
  default = "eu-central-1a"
  description = "The availablity zone in which our AWS resources will be created."
}
```

### Loading Terraform Input Variables

[Terraform Input Variables](https://developer.hashicorp.com/terraform/language/values/variables)

#### var flag
We can use the `-var` flag to set an input variable or override a variable in the tfvars file eg. `terraform -var user_ud="my-user_id"`


#### terraform.tvfars
This is the default file to load in terraform variables in blunk



#### var-file flag
When use different file in place of `.terraform.tfvars` to define the input variables, then we need to explicitly provide the argument `-var-file` during the terraform plan or apply
One use case for having var files other than `.terraform.tfvars` can be: our organization wishes to maintain dedicated var files for each of our environments. Like a `dev.tfvars` for `dev` env, `qa.tfvars` for `qa` env. etc.
Example:


dev.tfvars
```tf
environment = "dev"
resoure_group_location = "us-east-1"
```

```tf
terraform plan -var-file="dev.tfvars"
terraform apply -var-file="dev.tfvars"
``

#### auto.tfvars
- Can have any name, but must end with an `.auto.tfvars` extension.
- The variables inside these files will be auto loaded during terraform plan or apply


### Order of terraform input variables
![input-precedence](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/04391377d8501643fece4e04ef8b97a9fd51dcf0/journal/assets/week-1/TF-var-precedence.gif)



-----------------------------------------------------------------------------------------------------

# Personal Documentation 
1. ## Journal TOC and Major Version
  - Under issue `#19Create TOC Readme ` and PR `#19 create toc readme` we did the following:
    - We restructured the documentation into `journal/<week.md>`
    - We moved our week-0 notes from `./README.md` to `./journal/week-0.md`
    - Created `./journal/week-1.md` for week-1 documentation

2. ## Restructure Root Module
2.1 Create an `Issue`
```txt
Restructure Root Module
[ ] variables.tf
[ ] outputs.tf
[ ] main.tf
[ ] providers.tf
https://developer.hashicorp.com/terraform/language/modules/develop/structure
[ ] Migrate back to TF local state

Label: enhancement
```

2.1 Create a branch and launch in Gitpod

2.3 Create the following files at the project root level: `outputs.tf` , `providers.tf`, `variables.tf`  and `terraform.tfvars`

2.4 Open `main.tf`. We will be reviewing this file and moving code to the relevant newly created files: `outputs.tf` , `providers.tf`, `variables.tf`  and `terraform.tfvars`

2.5 Update `providers.tf` by pasting the provider block from the `main.tf` here

```tf
terraform {
   cloud {
    organization = "aggarwaltanushree"
    workspaces {
      name = "terra-house-1"
    }
   }
  required_providers {
    random = {
      source = "hashicorp/random"
      version = "3.5.1"
    }
    aws = {
      source = "hashicorp/aws"
      version = "5.16.2"
    }
  }
}

provider "aws" {
}
provider "random" {
  # Configuration options
}
```

2.6 Update  `outputs.tf` file
```tf
output "random_bucket_name" {
  value = random_string.bucket_name.result
}
```

2.7 Update `variables.tf` with the definintion of a new var `user_uuid`
```tf
variable "user_uuid" {
  description = "The UUID of the user"
  type        = string
  validation {
    condition        = can(regex("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$", var.user_uuid))
    error_message    = "The user_uuid value is not a valid UUID."
  }
}
``

2.8 Next, we would like to add `tags`  to our `S3 bucket` (it will be the name of our S3 bucket). We need to verify in the [Terraform Registry S3 bucket documentation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket) if S3 buckets support tagging and the associated syntax.
Add `tags` to S3 bucket resource in `main.tf`
```tf
  tags = {
    UserUuid = var.user_uuid
  }
```

2.9.1 Update `terraform.tfvars` by creating a new var named `user_uuid` and assign it the value of your `ExamPro a/c UUID`
```tf
user_uuid="<add value here>"
```
2.9.2 Also create a TF cloud `terraform variable` for the same, while we are at it. This will be referenced when we migrate to TF Cloud.

2.10 Revert back from TF Cloud to TF Local
2.10.1 Run `tf init` followed by `tf destroy` to destory your TF cloud resources.

![tf-destory](
tf-cloud-detroy.png

2.10.2 Comment out the `cloud provider block` in `providers.tf` 

```tf
terraform {
 #  cloud {
 #   organization = "aggarwaltanushree"
 #   workspaces {
 #     name = "terra-house-1"
 #   }
 #  }
  required_providers {
    random = {
      source = "hashicorp/random"
      version = "3.5.1"
    }
    aws = {
      source = "hashicorp/aws"
      version = "5.16.2"
    }
  }
}	


2.10.3 Delete the `terrafor.lock.hcl` file and `.terraform` folder so we can migrate back to TF local.

2.11 Run `tf init` to initialize TF local 

2.12 Run `terraform plan -var user_uuid="<insert value here>"`. 
Now that we know that our var works correctly, lets try `terraform plan` (this will pick the var value from our `terraform.tfvars` file followed by `tf apply`.

> Since our TF state file: `terraform.tfstate` will not be checked into our code repository and we are not running `tf destroy`, we will face issues the next time we try to work with our resources. We will tackle this in our next branch!

2.13 Note: we will not be checking-in `terraform.tfvars` file. So we can add a step in our `.gitpod.yml` to copy `terraform.tfvars.example` as `terraform.tfvars` in our project

Let's update the `terraform` section of our `.gitpod.yml`

```yml
cp $PROJECT_ROOT/terraform.tfvars.example $PROJECT_ROOT/terraform.tfvars
```

2.14 Update the documentation

2.15 Stage, commit and sync the changed to Github

2.16 Create a PR and Merge this branch `21-restructure-root-module` to the `main` branch.

2.17 Issue tags to the `main branch` as `1.1.0`