# Week 0 — Prep Week Journal

## Task Status
| Topic | Status |
| --- | --- | 
| [Branching Tagging PR](#branching-tagging-pr) | <ul><li> [x] </li></ul> |
| [Terraform CLI Refactor](#terraform-cli-refactor) | <ul><li> [x] </li></ul> |
| [Project Root Env Var](#project-root-env-var) | <ul><li> [x] </li></ul> |
| [AWS CLI Refactor](#aws-cli-refactor) | <ul><li> [x] </li></ul> |
| [Random Terraform Provider Init Plan Apply](#random-terraform-provider-init-plan-apply) | <ul><li> [x] </li></ul> |
| [Terraform Provider S3 bucket](#terraform-provider-s3-bucket) | <ul><li> [x] </li></ul> |
| [Terraform Cloud and Terraform Login](#terraform-cloud-and-terraform-login) | <ul><li> [x] </li></ul> |
| [Terraform Login Workaround](#terraform-login-workaround) | <ul><li> [ ] </li></ul> |
| [TF Alias](#tf-alias) | <ul><li> [ ] </li></ul> |
| [Project Validation](#project-validation) | <ul><li> [ ] </li></ul> |

## Personal Documentation
1. ### Branching Tagging PR
1.1 Open your bootcamp repository in [Github](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023)

1.2 Navigate to the `Issues` tab and click on `New Issues`
![issues](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/0-issues.png)

1.3 Give the issue a name and description:
```txt
Add semantic documentation to project

We want to add semantic versioning in our Terraform bootcamp project.
```
then click `Submit new issue`
![issue_create](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/1-issue-create.png)

1.4 Click on the newly created issue to view its details. 
> Notice that an _issue number_ has been auto generated.
![view_issue](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/2-issue-num.png)

1.5 On the right you will notice an option to create a `branch` for this `issue`. Click the option.
![create_branch](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/3-create-pr.png)

1.6 Give the branch a meaningful `name` and select the radio button: `Checkout locally`
![create_branch_1](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/4-create-branch-name.png)

Click `Create branch`

1.7 Copy the command returned. This will be used to checkout the newly created branch in Gitpod.
![cpy_cmd](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/5-checkout-cmd.png)

`git checkout 1-add-semantic-documentation-to-project`

1.8 Launch the repo in `Gitpod`
![launch_gitpod](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/6-launch-gitpod.png)

1.9 At the Gitpod terminal, type the command copied in step-1.7 to checkout the issue branch
`git checkout 1-add-semantic-documentation-to-project`

![chkout_branch](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/7-checkout-branch.png)

1.10 Open the `README.md` file and update it with the following information:

```md
## Semantic Versioning

This project is going to use Semantic versioning for its tagging.
[semver.org](https://semver.org/)

The general format is: 

**MAJOR.MINOR.PATCH**, example: `1.0.1` 

>Note: some projects like to name the versions with a `v` prefix. example: `v1.0.1` . This however is not scematic versioning_

- **MAJOR** version when you make incompatible API changes
- **MINOR** version when you add functionality in a backward compatible manner
- **PATCH** version when you make backward compatible bug fixes
Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.
```

1.11 Stage the file by clicking `+` sign, then `Commit` the file to your Gihub repo.
**Note: make sure you enter a descriptive comment message**
example: `"#1 Add semantic versioning to project"`

![commit-branch](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/8-commit-branch.png)

![sync-branch](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/9-sync-branch.png)

1.12 Add and push `Tags` to the Github repo using the below commands:

```sh
git tag 0.1.0
git push --tags
```

![tags](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/10-tagging.png)

1.13 Open the bootcamp repo in Github and verify that the commit was synced.

1.14 Next, we need to bring this feature back to the `main` branch. For this we will be creating a `pull request`.

In Github, navigate to the `Pull requests` tab, and click on `New pull request`.

![new_pull](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/11-new-pull.png)

1.15 Select `base: main` and `compare: add-semantic-documentation` as the branches.


> Notice the message:  ✔️`Able to merge` . This indicates that we are Git will be able to merge our feature branch into the `main` branch.

Click `Create pull request` 

![pull_create](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/12-pull-create.png)

1.16 Next, we `merge` the `pull request`

>Note: In real-world projects, someone else from the team review's the code before merging it. Since bootcamp is for individual submissions, we will be skipping this step and merging the branch ourself without an reviewer_

![merge](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/13-merge.png)

Confirm the `Merge` to complete the merging with `main` branch and `close the issue`.

![merge_close](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/3ec7459e623b566ea6c4c983f165e5cd18d822a8/journal/assets/week-0/14-merge-close-issue.png)

>Note: in production scenarios, ideally branches are deleted after they have been inspected and merged. For the bootcamp we will be retaining the individual branches for grading verification ._

1.17 Stop the `Gitpod workspace`.

![stop_gitpod](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/008ed87ac39ddcbf21eaa6c7f03a9650d1c605b8/journal/assets/week-0/15-stop-gitpod-workspace.png)


2. ### Terraform CLI Refactor 

2.1 Create a new feature branch in your Github repositiory.

```txt
Issue name: Refactor Terraform CLI
Issue description: Fixing this issue with installing the Terraform CLI. We need to make sure it automatically installs to completion without user input.
Label: bug
```

![create-issue](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/31bb5c4d8cc3f8b36ceafdca71cad2cd9e095447/journal/assets/week-0/16-new-issue.png)

2.2 Create a branch for this issue and launch it in Gitpod.

2.3 Open the `.gitpod.yml` file. We need to figure out the issue with this file, i.e. why Terraform CLI is not getting installed without a user input.
In order to troubleshoot, start executing each line of code from the `.gitpod.yml` in the `Terminal`. Keep an eye out for the line of code that prompts for user input.

![troubleshoot](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/31bb5c4d8cc3f8b36ceafdca71cad2cd9e095447/journal/assets/week-0/17-terrafomr-cli-installation.png)

Now that we have figured out the issue, let us begin refactoring our code.
Open the official [Terraform CLI installation documentation](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli) on the Hashicorp website in your web browser.
We will be using this documentation as a reference to configure our installation script.

2.4 Create a new folder at the root level of the project and name it `bin`

2.5 Create a bash shell script inside `bin` named `install_terraform_cli.sh`
```sh
#!/usr/bin/env bash

sudo apt-get update && sudo apt-get install -y gnupg software-properties-common curl

wget -O- https://apt.releases.hashicorp.com/gpg | \
gpg --dearmor | \
sudo tee /usr/share/keyrings/hashicorp-archive-keyring.gpg

gpg --no-default-keyring \
--keyring /usr/share/keyrings/hashicorp-archive-keyring.gpg \
--fingerprint

echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] \
https://apt.releases.hashicorp.com $(lsb_release -cs) main" | \
sudo tee /etc/apt/sources.list.d/hashicorp.list

sudo apt update

sudo apt-get install terraform -y
```

2.6 Try executing this shell script in the terminal:
`source ./bin/install_terraform_cli.sh`

![exec-script](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/31bb5c4d8cc3f8b36ceafdca71cad2cd9e095447/journal/assets/week-0/18-check-script-exec.png)

Notice the output, the script was able to execute properly.
> Note: a more common way to execute a shell script is simply executing it from the relative path, without using `source`. 
Try executing this script this way.
`./bin/install_terraform_cli`

>You will notice an error of `permission denined`. The reason for this error is, that our script lacks `executable permissions`.
Grant the script executable permissions using `chmod u+x ./bin/install_terraform_cli`

Try executing the script again.
`./bin/install_terraform_cli`
This time it works! Good job!

2.7 Now that we have a working terraform-cli installation script, let's configure it in our `.gitpod.yml` file.
Edit the `.gitpod.yml` file as below:

```yml
tasks:
  - name: terraform
    before: |
      source ./bin/install_terraform_cli
  - name: aws-cli
    env:
      AWS_CLI_AUTO_PROMPT: on-partial
    before: |
      cd /workspace
      curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
      unzip awscliv2.zip
      sudo ./aws/install
      cd $THEIA_WORKSPACE_ROOT
vscode:
  extensions:
    - amazonwebservices.aws-toolkit-vscode
    - hashicorp.terraform
```

Essentially we are updating:

~~init~~ -> before
and calling our terraform-cli installation shell script.

2.8 Update the documentation `README.md` with all the learnings from this activity.

2.9 Stage and commit the updates. Sync the changes.
> Make sure to include the issue number in the commit message.
Example: `#3 refactor terraform cli into a bash script for .gitpod.yml`

2.10 Back in Github repo, open a pull request for the #3 branch and merge it with the main branch.
![2-pr](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/31bb5c4d8cc3f8b36ceafdca71cad2cd9e095447/journal/assets/week-0/19-pull-request.png)

![2-merge](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/31bb5c4d8cc3f8b36ceafdca71cad2cd9e095447/journal/assets/week-0/20-merge.png)

2.11 In Gitpod, tag and push the changes to `main`

```sh
git fetch
git checkout main
git pull
git tag 0.3.0
git push --tags
```
![add-tags](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/31bb5c4d8cc3f8b36ceafdca71cad2cd9e095447/journal/assets/week-0/21-tags.png)

2.12 Stop the Gitpod workspace.

3. ### Project Root Env Var

3.1 Create a new issue 

```txt
Issue name: Project Root Env Var
Issue description: Setup an environment variable for PROJECT__ROOT that can be referenced in our bash scripts.
Label: documentation
```
3.2 Create a branch for this issue and launch it in Gitpod.

3.3 Type `env` at the terminal to check all existing environment variables (Env Var).

3.4 Create a `.env.example` file containing examples of the env vars that will be used in this project.
```txt
PROJECT_ROOT='/workspace/terraform-beginner-bootcamp-2023'

```
3.4 Create a env var in Gitpod for our `PROJECT_ROOT` :

`gp env PROJECT_ROOT='/workspace/terraform-beginner-bootcamp-2023'`

![env-var](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/31bb5c4d8cc3f8b36ceafdca71cad2cd9e095447/journal/assets/week-0/22-proj-root-env-var.png)

3.5 Update the ./bin/install_terraform_cli shell script to reference the newly created env var

```sh
#!/usr/bin/env bash

cd /workspace

sudo apt-get update && sudo apt-get install -y gnupg software-properties-common curl

wget -O- https://apt.releases.hashicorp.com/gpg | \
gpg --dearmor | \
sudo tee /usr/share/keyrings/hashicorp-archive-keyring.gpg

gpg --no-default-keyring \
--keyring /usr/share/keyrings/hashicorp-archive-keyring.gpg \
--fingerprint

echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] \
https://apt.releases.hashicorp.com $(lsb_release -cs) main" | \
sudo tee /etc/apt/sources.list.d/hashicorp.list

sudo apt update

sudo apt-get install terraform -y

cd $PROJECT_ROOT
```

3.6 Update the documentation for this task in `README.MD` .

3.7 Stage, commit and sync the changes.
> Make sure to add the issue nummber #5 in the commit message

3.8 Stop the Gitpod workspace and re-launch it. Verify if our PROJECT_ROOT env var persists.

3.9 Back in github, create and merge the branch `5-project-root-environment-variable` with the `main` branch.
![pr](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/31bb5c4d8cc3f8b36ceafdca71cad2cd9e095447/journal/assets/week-0/23-pr.png)

3.10 In Gitpod, tag and push the changes to `main`.

```sh
git fetch
git checkout main
git pull
git tag 0.3.0
git push --tags
```

![add-tags](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/31bb5c4d8cc3f8b36ceafdca71cad2cd9e095447/journal/assets/week-0/24-add-tags.png)

3.11  Stop the `Gitpod workspace`.

4. ### AWS CLI Refactor
4.1 Create an `issue`
```txt
Name: refactor aws cli script
description: 
- [] Refactor AWS CLI into bash script
- [] Provide env var examples for AWS CLI requirements
- [] set our env vars for AWS using gp env
label: enhancement
```

![create-issue](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/28-create-issue.png)

4.2 Create a branch for the issue and launch in Gitpod

4.3 Under `bin` directory, create a new file named `install_aws_cli`

```sh
#!/usr/bin/env bash

cd /workspace

rm -f '/workspace/awscliv2.zip'
rm -rf '/workspace/aws'

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

aws sts get-caller-identity

cd $PROJECT_ROOT
```

4.4 Update the `.gitpod.yml` file with the `./bin/install_aws_cli` script name as below:
Since we have created a new script `./bin/install_aws_cli` containing the steps for AWS CLI installation, we do not need to maintain the installation steps in our `.gitpod.yml`, we can simplify this file by calling our bash script from here.

Replace 
```yml
  - name: aws-cli
    env:
      AWS_CLI_AUTO_PROMPT: on-partial
    before: |
      cd /workspace
      curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
      unzip awscliv2.zip
      sudo ./aws/install
      cd $THEIA_WORKSPACE_ROOT
```
with -> 

```yml
	  - name: aws-cli
		env:
		  AWS_CLI_AUTO_PROMPT: on-partial
		before: |
		  source ./bin/install_aws_cli
```

4.5 Grant the script executable permisssions `chmod u+x ./bin/install_aws_cli`

4.6 Execute the script to verify that it executes and is able to install the AWS CLI.

4.7 Check if any AWS credentials are configured, by executing the following command at the terminal:
`aws sts get-caller-identity`

A blank line is returned by the command. This indicates that there aren't any credentails set. 
Next, we will be creating these credentails.

4.8 Login to the `AWS Management Console` and navigate to [IAM](https://us-east-1.console.aws.amazon.com/iamv2/home?region=eu-central-1#/home)

Create a new [IAM user](https://us-east-1.console.aws.amazon.com/iamv2/home?region=eu-central-1#/users) for the bootcamp.

User: `terraform-beginner-bootcamp`
Permissions: `AdministratorAccess (or add it to your Admin group)`

![create-iam-user](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/25-create-iam-user.png)

![add-perms-to-iam-user](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/26-create-iam-user.png)

4.9 Create `access keys` for the newly created IAM user.

![create-access-key](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/27-create-access-keys.png)

4.10 Back in Gitpod workspace, update the `.env.example` file with the examples of required AWS credentails.

```txt
AWS_ACCESS_KEY_ID='AKIAIOSFODNN7EXAMPLE'
AWS_SECRET_ACCESS_KEY='wAbcdEFGhiJkl/M7NOPQR/sTuvWXYzEXAMPLEKEY'
AWS_DEFAULT_REGION='eu-central-1'
```

4.11 Set these credentials in Gitpod execting the commands at the terminal:
```sh
gp env AWS_ACCESS_KEY_ID='AKIAIOSFODNN7EXAMPLE'
gp env AWS_SECRET_ACCESS_KEY='wAbcdEFGhiJkl/M7NOPQR/sTuvWXYzEXAMPLEKEY'
gp env AWS_DEFAULT_REGION='eu-central-1'
```
Verify your [Gitpod variables](https://gitpod.io/user/variables) list to check if they were successfully set.

4.12 Update `README.md` with the learnings.

4.13 Stage, Commit and sync the changes. Stop the Gitpod environment and relaunch it.

4.14 Check if the AWS CLI is installed and if our credentails persisted.
`aws sts get-caller-identity`

![check-identity](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/29-check-identity.png)

4.15 Create pull request and merge the change to main branch.

![issue-pr](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/31-pr-and-merge.png)

4.16 Add the tags.

```sh
git fetch
git checkout main
git pull
git tag 0.4.0
git push --tags
```

![add-tags](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/30-add-tags.png)

4.17 Stop the Gitpod workspace.

5. ### Random Terraform Provider Init Plan Apply
5.1 Create an `issue`
```txt
Name: Terraform Random Bucket Name
Description:
- [] explore the terraform registry
- [] install the terraform random provider
- [] run terraform init
- [] generate out a random bucket name
- [] output the random bucket name to outputs
- [] run terraform plan
- [] run terraform apply

label: good-first-issue
```

5.2 Create a branch for the issue and launch in Gitpod
5.3 Open **main.tf** file and paste the Terraform **provider code** from the [Terraform Registry website](https://registry.terraform.io/providers/hashicorp/random/latest) -> Use Provider
We will be using this provider to generate random things in our Project.
> Note: these steps are being performed just to gain inderstanding of Terraform providers. Random provider is not directly related to our project implementation.


```tf
terraform {
  required_providers {
    random = {
      source = "hashicorp/random"
      version = "3.5.1"
    }
  }
}

provider "random" {
  # Configuration options
}
```

5.4 Next, we create a **resource** for this **provider**. Copy the code from the [documentation](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string)

```tf
resource "random_string" "random" {
  length           = 16
  special          = false
}
```
We will modify it to generate our S3 bucket name name.

5.5 We wish to output the S3 bucket name, for this we refer the [output documentation](https://developer.hashicorp.com/terraform/language/values/outputs)

5.6 Update the `main.tf` file as below:
```tf
terraform {
  required_providers {
    random = {
      source = "hashicorp/random"
      version = "3.5.1"
    }
  }
}

provider "random" {
  # Configuration options
}

resource "random_string" "bucket_name" {
  length   = 16
  special  = false
}

output "random_bucket_name" {
  value = random_string.bucket_name.result
}

```

5.7 Next, open the **terraform terminal** in Gitpod and execute `terraform`
This will confirm that our Terraform CLI has been properly installed. 
> In case of issues, troubleshoot and fix. DO NOT proceed to the next step if the terraform CLI is not working!

5.8 Run `terraform init`
> Notice that a new file **.terraform.lock.hcl** has been created, this file is maintained automatically by "terraform init". Any manual edits may be lost in future updates.
A folder named `terraform` has also been geenrated. This stores the terraform `provider` binaries, that have been downloaded from the Terraform registry by `terraform init`.

![terraform-init](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/32-terraform-init.png)

5.9 Next, run `terraform plan` to generate the change set, i.e set of resources that will be created when we run `terraform apply`.

5.10 Finally, run `terraform apply --auto-approve` which will create our resources defined in our **main.tf**.

5.11 To view the output, we can execute the command:
`terraform output`

Or to view a specific output (in this case the S3 bucket name) :
`terraform output random_bucket_name`

![terraform-output](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/33-terraform-output.png)

5.12 Update the documentation with the learnings `README.md`

5.13 Stage, Commit and Sync the changes.
> Don't forget to reference the issue number #9

5.14 Create a pull request and merge `9-terraform-random-bucket-name` into `main` branch.

5.15 Add tags
```sh
git checkout main
git pull
git tag 0.5.0
git push --tags
```
![add-tags](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/34-add-tags.png)

5.16 Stop Gitpod


6. ### Terraform Provider S3 bucket
6.1 Create an `issue`
```txt
Name: Simple S3 Bucket
Description:
- [] Define an S3 Bucket in Terraform
- [] Use the random resource string for the name
- [] Install the AWS Terraform Provider
- [] Configure AWS Provider
- [] Terraform Apply and Terraform Destroy

label: enhancement
```

6.2 Create a branch for the issue and launch in Gitpod
6.3 At the terminal in Gitpod, execute `aws s3 ls`.
> If our environment is setup correctly, this should be able to connect to our AWS acoount and fetch the names of S3 buckets available there.

6.4 Run `terraform init` and `terraform apply` to download the terraform provider binaries and create the resources defined in our `main.tf`

6.5 Open the [Terraform Registry](https://registry.terraform.io/providers/hashicorp/aws/latest) in a web browser and copy the **provider** code for **AWS**.  
Then navigate to **aws_s3_bucket** resource  [Terraform Registry](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket).

We will be incorporating these in our `main.tf`

>Important pointer: A Terraform project can contain only **one required_providers block**

6.6 Update the `main.tf` file to incude the AWS provider.

```tf
terraform {
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

# https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string
resource "random_string" "bucket_name" {
  length   = 32
  special  = false
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket
resource "aws_s3_bucket" "example" {
  bucket = random_string.bucket_name.result
}

output "random_bucket_name" {
  value = random_string.bucket_name.result
}

```

6.7 Run `terraform init` again, so the AWS provider binaries are downloaded to our environment.
> Whenever a new provider is added to a project, `terraform init` needs to be executed, so the required binaries are downloaded in our environment.

![provider-init](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/35-aws-provider.png)

>Important:
Terraform needs our AWS credentials in order to create the defined AWS resource (S3 bucket in this case).
These credentails can be defined in the `main.tf` file, but this is **not recommended**, since the main.tf will be commited to our code repository. The better way to safely use credentials is to set them as env vars.
Terraform flow for reading credentials:
(1) check config file _if not present, then_ -> (2) read from env vars


6.8 Execute `terraform plan`, followed by `terraform apply --auto-approve`
Notice the error while creating the `S3 bucket`
This is because **S3 naming convention does not allow _uppercase_ letters in the bucket names.**

![s3-error](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/36-s3-error.png)

Refer the [documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html) for all rules relating to the S3 bucket naming.

6.9 We need to exclude _uppercase letters_ from our random S3 bucket name geenrator.
Let's update the code in `main.tf`:

```tf
terraform {
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

# https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string
resource "random_string" "bucket_name" {
  length   = 32
  lower = true
  upper = false
  special  = false
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket
resource "aws_s3_bucket" "example" {
  # Bucket Naming Rules
  #https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html?icmpid=docs_amazons3_console
  bucket = random_string.bucket_name.result
}

output "random_bucket_name" {
  value = random_string.bucket_name.result
}
```

6.10 Run the plan and apply commands again. 
```hcl
terraform plan
terraform apply --auto-approve
```

Did the bucket get created?

Yes!

![s3-bucket-created](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/37-bucket-created.png)

Verify the same by checking in the AWS Management Console
![s3-bucket-available](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/38-bucket-verified.png)

6.11 After the verification, we are good to proceed with the resource deletion.
``terraform destroy --auto-aprove`

![resources-destroyed](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/39-s3-destroyed.png)

6.12 Stage, Commit and Sync the changes.

6.13 Create a pull request and merge `11-simple-s3-bucket` with the `main` branch.

6.14 Add tags
```sh
git checkout main
git pull
git tag 0.6.0
git push --tags
```

6.15 Stop the Gitpod workspace.
![tagging](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/dc4ebb49d2bcd9fbdc912835b3fb3754cd270e21/journal/assets/week-0/40-tags.png)

7. ### Terraform Cloud and Terraform Login
7.1 Create an `issue`
```txt
Issue name: Terraform Cloud Backend
Desc:
- [] Configure Terraform Cloud Backend
- [] Workaround for Terraform Login
- [] Migrate our local state to remote state
- [] Create a new Project and Workspace in Terraform Cloud
 
Label: Enhancement
```

7.2 Create a branch for the issue and launch in Gitpod

7.3 Login to Gitpod and initialize your Terraform providers and create the resources like we have been doing in the previous steps.

```tf
terraform init
terraform plan
terraform apply --auto-approve
```

7.4 Login to your [Terraform Cloud Account](https://app.terraform.io/app)

7.5 Create a new `Project` with the name: `terraform-beginner-bootcamp-2023`

![tfc-project](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/43-tfc-project.png)

7.6 Inside the `terraform-beginner-bootcamp-2023` project, create a new workspace 
	- Select workspacr type: CLI-driven workflow
	- workspace name: terra-house-1
	- workspace description: Terrahouse infrastructure that will connect to TerraTowns.

![tfc-workspace](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/46-save-token-2.png)

7.7 Back in Gitpod workspace, update the `main.tf` code to include the `remote cloud provider`

```tf
terraform {
  cloud {
    organization = "aggarwaltanushree"
    workspaces {
      name = "terra-house-1"
    }
```


7.8 Then execute `terraform init` at the terminal.
Notice the error.

![init-error](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/45-init-error.png)

> Since we are trying to connect to Terraform cloud (app.terraform.io), we first need to configure our credentails, which can be done by `terraform login`

When promted, enter `yes`
Shift+p to view the available options.
![save-token](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/45-save-token.png)

Copy the document path `https://app.terraform.io/app/settings/tokens?source=terraform-login` and create the token there.
Generate the `token` for one day. Copy and path where the token needs to be stored. 

![save-tokens](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/46-save-token-2.png)

Create the credentials file and paste the token into it.

```sh
touch /home/gitpod/.terraform.d/credentails.trfc.json
open /home/gitpod/.terraform.d/credentails.trfc.json
```

![credentials](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/47-save-token-in-credential.png)

7.9 Once the credentials have been saved, we can proceed with `terraform init`.
At this stage we have the option to save our workspace terraform state (i.e our `terraform.tfstate` json file) into the terraform cloud workspace. Enter `yes` when prompted to do so.

![terraform-init-cloud](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/48-terraform-init-cloud.png)

7.10 Back in your Terraform Workspace in TFC, verify if our resources and outputs are visible.

![terraform-init-cloud-resources](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/49-terraform-init-cloud-resources.png)

7.11 Execute `terraform plan` and `terraform apply`

![run-plan-creds-error](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/51-run-plan-creds-error.png)

> We run into an error! Open the detailed log in your TF cloud a/c to see the details.
The error indicates that we do not have AWS credentials configured!
You may now wonder, that we did this already while configuring the AWS CLI. You are correct.
However, this time around we are running in the _remote state_. We need to configure the AWS credentials as `env vars` in our Terraform cloud a/c as well.

Let's get on with it!

Navigate to [variable sets](https://app.terraform.io/app/aggarwaltanushree/settings/varsets) and create a new `variable set` for your AWS credentials.

> Tips:
> - make sure you set these as env vars and not TF vars
> - since these are sensitive credentials, enable the `sensitive` checkbox.
> - set scope only as necessay, i.e our `terraform-beginner-bootcamp-2023` project instead of `global` scope.

![set-tfc-aws-env-vars](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/52-tfc-aws-env-vars.png)

7.12  Execute `terraform plan` again, this time it should work. Next, execute `terraform apply`

![run-plan](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6e9373c9ed6a2b1be62aee1bfb5595bc50bdac3e/journal/assets/week-0/50-run-plan.png)

7.13 Stage, commit and sync the changes

7.14 PR and Merge `13-terraform-cloud-backend` to `main`

7.15 Add Tags

```sh
git checkout main
git pull
git tag 0.7.0
git push --tags
```
7.16. Stop the Gitpod Workspace

8. #### Terraform Login Workaround

9. #### TF Alias

10. #### Project Validation
    
