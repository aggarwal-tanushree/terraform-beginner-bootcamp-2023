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
```

#### auto.tfvars
- Can have any name, but must end with an `.auto.tfvars` extension.
- The variables inside these files will be auto loaded during terraform plan or apply


### Order of terraform input variables
<img src="https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/04391377d8501643fece4e04ef8b97a9fd51dcf0/journal/assets/week-1/TF-var-precedence.gif"  width="40%" height="30%">



## Dealing With Configuration Drift

## What is Configuration drift?
- One challenge when managing infrastructure as code is `drift`.
- `Drift` is the term for when the _real-world state_ of your infrastructure differs from the _state defined in your configuration_. 
- The Terraform state file `terraform.tfstate`  is the record of all resources Terraform manages. Since in our project, we are not commiting the TF state file to our repo, it is lost when we close our environment. The next time our environemnt is launched, TF is unable to determine which resources exist and what is their defined state.

## What happens if we lose our `terraform.tfstate` state file?

If you lose your statefile, you most likley have to tear down all your cloud infrastructure manually.

You can use terraform port but it won't for all Cloud resources. You need check the terraform providers documentation for which resources support import.

### Fix Missing Resources with Terraform Import

`terraform import aws_s3_bucket.bucket bucket-name`

[Terraform Import](https://developer.hashicorp.com/terraform/cli/import)
[AWS S3 Bucket Import](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket#import)
Command to import S3 bucket: `terraform import aws_s3_bucket.<bucket-variable-name> <AWS-bucket_name>`
example: ``tf import aws_s3_bucket.example <bucket_name>``

### Terraform import block TFv1.5.0
https://www.hashicorp.com/blog/terraform-1-5-brings-config-driven-import-and-checks?source=post_page-----a6ca245a7daf--------------------------------
The import block takes two required arguments one optional argument.:
`to` - The instance address this resource will have in your state file. for example, you can provide a resource id for your Azure resource.
`id` - A string with the import ID of the resource in your terraform configuration file.
`provider` (optional) - An optional custom resource provider, This is useful when you’re using multiple providers.

example:
```tf
import {
  # ID of the cloud resource
  # Check provider documentation for importable resources and format
  id = "i-abcd1234"
 
  # Resource address
  to = aws_instance.example
}
```


### Fix Manual Configuration

- when someone goes and delete or modifies cloud resource manually through ClickOps. 

- If we run `terraform plan` it with attempt to put our infrstraucture back into the expected state fixing Configuration Drift

Further reading: https://developer.hashicorp.com/terraform/tutorials/state/resource-drift


## Fix using Terraform Refresh

```sh
terraform apply -refresh-only -auto-approve
```

## Terraform Modules

### Terraform Module Structure

It is recommend to place modules in a `modules` directory when locally developing modules but you can name it whatever you like.

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


### Passing Input Variables

- We can pass input variables to our module.
- The module has to declare the terraform variables in its own `variables.tf`
```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  bucket_name = var.bucket_name
}
```
- However, a definition is also required in the `variables.tf` at the root level. This however, does not need to match the whole definition from the `modules/terrahouse_aws/variables.tf`. Adding a var name, with a type and description should suffice.
Definition in the root level `variables.tf` :
```tf
variable "user_uuid" {
 type = string
}

variable "bucket_name" {
 type = string
}
```

### Modules Sources

Using the source we can import the module from various places eg:
- locally
- Github
- Terraform Registry

```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
}
```

[Modules Sources](https://developer.hashicorp.com/terraform/language/modules/sources)


### Module Outputs
- modules can define their own o/ps
`modules/terrahouse_aws/outputs.tf`
```tf
output "bucket_name" {
  value = aws_s3_bucket.website_bucket.bucket
}
```

- similar to _module input vars_ , the module o/ps need to be referenced in the root level `outputs.tf`
```tf
output "bucket_name" {
  description = "Bucket name for our static website hosting"
  value = module.terrahouse_aws.bucket_name
}
```



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
```

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

3. ## Terraform Import and Configuration Drift
3.1 Create an `Issue`
```txt
Configuration Drift

[ ] use terraform import
[ ] purposely cause configuration drift via clickops, and correct state.
Label: bug
```

3.2 Create a branch and launch in Gitpod

3.3 Run `tf init` to intialize our TF env

3.4 Since we do not have the state file `terraform.tfstate`, we will need to `import` the lost resources manually.

`tf import aws_s3_bucket.example <bucket_name>`
`tf import random_string.bucket_name <bucket_name>`

3.5 Run `tf plan` to see if it is able to determine our resource state now.
Notice, that it works fine for the S3 bucket, but not for the `random`.

> So, this is where the use of `random` ends in our project. It worked great for our initial hands-on practice, but is no longer needed going forward.

3.6 Delete the `random` provider from `providers.tf` 

```tf
random = {
      source = "hashicorp/random"
      version = "3.5.1"
    }
```

3.7 Define a bucket name variable name, condition and validation in `variables.tf`

```tf
variable "bucket_name" {
  description = "The name of the S3 bucket"
  type        = string

  validation {
    condition     = (
      length(var.bucket_name) >= 3 && length(var.bucket_name) <= 63 && 
      can(regex("^[a-z0-9][a-z0-9-.]*[a-z0-9]$", var.bucket_name))
    )
    error_message = "The bucket name must be between 3 and 63 characters, start and end with a lowercase letter or number, and can contain only lowercase letters, numbers, hyphens, and dots."
  }
}
```

3.8.1 Create a var example for the bucket name in `terraform.tfvars.example`

```tf
bucket_name="<enter-bucket-name>"
```


3.8.2 Add this var to `terraform.tfvars` file with the real bucket name.


3.9 Update the S3 bucket name to `website_bucket` in `main.tf`, and delete the `random_string` resource block from the file.

```tf
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket
resource "aws_s3_bucket" "website_bucket" {
  # Bucket Naming Rules
  #https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html?icmpid=docs_amazons3_console
  bucket = var.bucket_name

  tags = {
    UserUuid = var.user_uuid
  }
}
```

3.10 Update the `outputs.tf` with the updated bucket name.

```tf
output "bucket_name" {
  value = aws_s3_bucket.website_bucket.bucket
}
```

3.11 Run `tf plan` if it works, follow it with `tf apply`

3.12 Update the documentation

3.13 Stage, commit and sync the changed to Github

3.14  Create a PR and Merge this branch `23-configuration-drift` to the `main` branch.

3.15 Issue tags to the `main branch` as `1.2.0`

4. ## Create Terrahouse Module
4.1 Create an `Issue`
```txt
AWS Terrahouse Module
 Setup directory structure for our module
 port our s3 bucket into the module.
label: enhancement
```

4.2 Create a branch and launch in Gitpod

4.3 Add a new folder named `modules/terrahouse_aws` and create the following filed under it `main.tf`, `outputs.tf`, `variables.tf`, `README.md`, `LICENSE` and `resource-storage.tf`

4.4.1 Update `modules/terrahouse_aws/main.tf` with the contents of `./main.tf`. 
4.4.2 Move the `aws` provider block from `./provider.tf` to `modules/terrahouse_aws/main.tf`. Delete the `./provider.tf` file.

```tf
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "5.16.2"
    }
  }
}
provider "aws"{
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket
resource "aws_s3_bucket" "website_bucket" {
  # Bucket Naming Rules
  #https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html?icmpid=docs_amazons3_console
  bucket = var.bucket_name

  tags = {
    UserUuid = var.user_uuid
  }
}
```


4.5 Move the contents of `./variables.tf` to `modules/terrahouse_aws/variables.tf`
```tf
variable "user_uuid" {
  description = "The UUID of the user"
  type        = string
  validation {
    condition        = can(regex("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$", var.user_uuid))
    error_message    = "The user_uuid value is not a valid UUID."
  }
}

variable "bucket_name" {
  description = "The name of the S3 bucket"
  type        = string

  validation {
    condition     = (
      length(var.bucket_name) >= 3 && length(var.bucket_name) <= 63 && 
      can(regex("^[a-z0-9][a-z0-9-.]*[a-z0-9]$", var.bucket_name))
    )
    error_message = "The bucket name must be between 3 and 63 characters, start and end with a lowercase letter or number, and can contain only lowercase letters, numbers, hyphens, and dots."
  }
}
```

4.6 Move the contents of `./outputs.tf` to `modules/terrahouse_aws/outputs.tf`

```tf
output "bucket_name" {
  value = aws_s3_bucket.website_bucket.bucket
}
```

4.7 Import our new module `terrahouse_aws` to `./main.tf` and its associated variables.

```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  bucket_name = var.bucket_name
}
```

4.8 Let's try to execute to see if there are any errors.
`tf init`
> Notice the **warning** 

```sh
Warning: Redundant empty provider block
│ 
│   on modules/terrahouse_aws/main.tf line 9:
│    9: provider "aws"{
│ 
│ Earlier versions of Terraform used empty provider blocks ("proxy provider configurations") for child modules to declare their need to be
│ passed a provider configuration by their callers. That approach was ambiguous and is now deprecated.
│ 
│ If you control this module, you can migrate to the new declaration syntax by removing all of the empty provider "aws" blocks and then
│ adding or updating an entry like the following to the required_providers block of module.terrahouse_aws:
│     aws = {
│       source = "hashicorp/aws"
│     }
│ 
│ (and one more similar warning elsewhere)
╵

Terraform has been successfully initialized!
```

> This indicates that we can delete the `aws` provider block from our `modules/terrahouse_aws/main.tf, since it is blank.
```tf
provider "aws"{
}
```

4.9 Lets' do a `tf plan` now.
> Error : complains about undeclared variables. 
```sh
Error: Reference to undeclared input variable
│ 
│   on main.tf line 3, in module "terrahouse_aws":
│    3:   user_uuid = var.user_uuid
│ 
│ An input variable with the name "user_uuid" has not been declared. This variable can be declared with a variable "user_uuid" {} block.
╵
╷
│ Error: Reference to undeclared input variable
│ 
│   on main.tf line 4, in module "terrahouse_aws":
│    4:   bucket_name = var.bucket_name
│ 
│ An input variable with the name "bucket_name" has not been declared. This variable can be declared with a variable "bucket_name" {}
│ block.
╵
```
> We know that we defined these variables in our `modules/terrahouse_aws/variables.tf`. However, a definition is also required in the `variables.tf` at the root level. This however, does not need to match the whole definition from the `modules/terrahouse_aws/variables.tf`. Adding a var name, with a type and description should suffice.

`variables.tf`
```tf
variable "user_uuid" {
 type = string
}

variable "bucket_name" {
 type = string
}
```

Try `tf plan` again. This time it should be able to read the input vars.

4.10 Let's proceed with `tf apply`
> Notice that it is able to create the resources, but does not list any outputs. This was working fine before we moved our o/p definitions to the  module. This is because, even though we have our o/p def in the module `outputs.tf`, they need to be referenced in the `./outputs.tf` for them to be displayed.
```sh
Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
gitpod /workspace/terraform-beginner-bootcamp-2023 (25-aws-terrahouse-module) $ tf output
╷
│ Warning: No outputs found
│ 
│ The state file either has no outputs defined, or all the defined outputs are empty. Please define an output in your configuration with
│ the `output` keyword and run `terraform refresh` for it to become available. If you are using interpolation, please verify the
│ interpolated value is not empty. You can use the `terraform console` command to assist.
╵
gitpod /workspace/terraform-beginner-bootcamp-2023 (25-aws-terrahouse-module) $ 
```

4.11 Reference the `bucket_name` o/p defined in `modules/terrahouse_aws/outputs.tf` in our root level `outputs.tf`
```tf
output "bucket_name" {
  description = "Bucket name for our static website hosting"
  value = module.terrahouse_aws.bucket_name
}
```

4.12 `tf plan` and `tf apply --auto-approve`.  Check the outputs: `tf outputs`
```sh
gitpod /workspace/terraform-beginner-bootcamp-2023 (25-aws-terrahouse-module) $ tf apply --auto-approve
module.terrahouse_aws.aws_s3_bucket.website_bucket: Refreshing state... [id=jvf0qijub046z6nj13vhm9463gjgf9g7]

Changes to Outputs:
  + bucket_name = "jvf0qijub046z6nj13vhm9463gjgf9g7"

You can apply this plan to save these new output values to the Terraform state, without changing any real infrastructure.

Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

Outputs:

bucket_name = "jvf0qijub046z6nj13vhm9463gjgf9g7"
gitpod /workspace/terraform-beginner-bootcamp-2023 (25-aws-terrahouse-module) $ tf output
bucket_name = "jvf0qijub046z6nj13vhm9463gjgf9g7"
gitpod /workspace/terraform-beginner-bootcamp-2023 (25-aws-terrahouse-module) $ 
```

4.13 Run `tf destroy --auto-approve`

4.14 Update the documentation

4.15 Stage, commit and sync the changed to Github

4.16  Create a PR and Merge this branch `25-aws-terrahouse-module` to the `main` branch.

4.17 Issue tags to the `main branch` as `1.3.0`
