# Terraform Beginner Bootcamp 2023  Week 0  — Prep Week Journal

[Week-0 Architecture](#week-0-architecture) :cloud: :seedling:

[Notes for revision](#notes-for-revision) :syringe: :medal_sports:

[Personal Documentation](#personal-documentation) :memo: :pencil:

[Task Status](#task-status) :star: :face_in_clouds: :partying_face: :white_check_mark:


# Task Status
| Topic | Status |
| --- | --- | 
| [Branching Tagging PR](#branching-tagging-pr)  | <ul><li> [x] </li></ul> |
| [Terraform CLI Refactor](#terraform-cli-refactor) | <ul><li> [x] </li></ul> |
| [Project Root Env Var](#project-root-env-var) | <ul><li> [x] </li></ul> |
| [AWS CLI Refactor](#aws-cli-refactor) | <ul><li> [x] </li></ul> |
| [Random Terraform Provider Init Plan Apply](#random-terraform-provider-init-plan-apply) | <ul><li> [x] </li></ul> |
| [Terraform Provider S3 bucket](#terraform-provider-s3-bucket) | <ul><li> [x] </li></ul> |
| [Terraform Cloud and Terraform Login](#terraform-cloud-and-terraform-login) | <ul><li> [x] </li></ul> |
| [Terraform Login Workaround](#terraform-login-workaround) | <ul><li> [x] </li></ul> |
| [TF Alias](#tf-alias) | <ul><li> [x] </li></ul> |
| [Project Validation](#project-validation) | <ul><li> [x] </li></ul> |


# Week 0 Architecture
![week0-architecture](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/7f19b23b71dbc5815dfeec76897bbc8ba19b8996/journal/assets/week-0/Week0-ArchitecturalDiagram.png)

The complete end-to-end architecture of what we will achieve at the end of the bootcamp is available [here](https://lucid.app/lucidchart/e3f15b1a-2211-4ddb-8c95-f144c2504db4/edit?invitationId=inv_0873b3c6-c652-463f-9f2b-fa0f1b420823&page=0_0#) 

_Diagram copyrights: Andrew Brown from ExamPro.co_

-----------------------------------------------------------------------------------------------------
# Notes For Revision

## Semantic Versioning

This project is going to use Semantic versioning for its tagging.
[semver.org](https://semver.org/)

The general format is: 

**MAJOR.MINOR.PATCH**, example: `1.0.1` 

_Note: some projects like to name the versions with a `v` prefix. example: `v1.0.1` . This however is not scematic versioning_

- **MAJOR** version when you make incompatible API changes
- **MINOR** version when you add functionality in a backward compatible manner
- **PATCH** version when you make backward compatible bug fixes

Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.


## Install the Terraform CLI

### Considerations with the Terraform CLI changes
The Terraform CLI installation instructions have changed due to gpg keyring changes. So we needed refer to the latest install CLI instructions via Terraform Documentation and change the scripting for install.

[Install Terraform CLI](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)


### Considerations for Linux Distribution

This project is built against Ubunutu.
Please consider checking your Linux Distrubtion and change accordingly to distrubtion needs. 

[How To Check OS Version in Linux](
https://www.cyberciti.biz/faq/how-to-check-os-version-in-linux-command-line/)

Example of checking OS flavor and version:

```
$ cat /etc/os-release

PRETTY_NAME="Ubuntu 22.04.3 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04.3 LTS (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy
```

### Refactoring into Bash Scripts

While fixing the Terraform CLI gpg depreciation issues we notice that bash scripts steps were a considerable amount more code. So we decided to create a bash script to install the Terraform CLI.

This bash script is located here: [./bin/install_terraform_cli](./bin/install_terraform_cli)

- This will keep the Gitpod Task File ([.gitpod.yml](.gitpod.yml)) tidy.
- This allow us an easier to debug and execute manually Terraform CLI install
- This will allow better portablity for other projects that need to install Terraform CLI.

#### Shebang 

A Shebang (prounced Sha-bang) tells the bash script what program that will interpet the script. eg. `#!/bin/bash` or `#!/usr/bin/env bash`

Purpose of Shebang:

- for portability for different OS distributions 
-  will search the user's PATH for the bash executable

[More on Shebang_(Unix)](https://en.wikipedia.org/wiki/Shebang_(Unix))

#### Execution Considerations

When executing the bash script we can use the `./` shorthand notiation to execute the bash script.

eg. `./bin/install_terraform_cli`

If we are using a script in .gitpod.yml  we need to point the script to a program to interpert it.

eg. `source ./bin/install_terraform_cli`

#### Linux Permissions Considerations

In order to make  bash scripts executable we need to change Linux permission for the script to be exetuable at the user/owner mode.

```sh
chmod u+x ./bin/install_terraform_cli
```

alternatively:

```sh
chmod 744 ./bin/install_terraform_cli
```

https://en.wikipedia.org/wiki/Chmod

#### Bash aliases
Bash allows setting shortcuts for long commands, these are referred to as `alias`. 
[Aliases](https://tldp.org/LDP/abs/html/aliases.html) can be defined in the bash profile `~/.bash_profile` as below:

`alias tf="terraform"`

> The `~/.bash_profile` file is read automatically when a new shell is loaded. So if we wish to execute it in the current shell, we can do so by executing `./bash_profile`



### Github Lifecycle (Before, Init, Command)

We need to be careful when using the Init because it will not rerun if we restart an existing workspace. Check out the lifecycle details at the below url:

https://www.gitpod.io/docs/configure/workspaces/tasks

### Working with Environment Variables a.k.a env vars

#### env command

- In Shell, we can list out all Enviroment Variables (Env Vars) using the `env` command

- Using `grep` along with `env` allows filtering. eg. `env | grep AWS_ will return all env vars defined in our environment matching the pattern `AWS_`
- One important use case for setting env vars in projects is defining paths for various files. Setting env vars is a good practice and a reusable way of setting paths, and avoiding hardcoding path in code.

#### Setting and Unsetting Env Vars

In the terminal we can set using `export AWS_DEFAULT_REGION="eu-central-1"

In the terrminal we unset using `unset AWS_DEFAULT_REGION`

We can set an environemnt variables temporarily when just running a command, as below:

```sh
HELLO='world' ./bin/print_message
```
Within a bash script we can set an environment variable without writing export eg.

```sh
#!/usr/bin/env bash

HELLO='world'

echo $HELLO
```

#### Printing Vars

We can print an env var using echo eg. `echo $HELLO`

#### Scoping of Env Vars

When you open up new bash terminals in VSCode it will not be aware of env vars that you have set in another window.

To make Env Vars to persist across all future bash terminals that are open, you need to set env vars in your [bash profile](https://linuxhint.com/simple-guide-to-create-open-edit-bash-profile/). eg. `.bash_profile`

#### Persisting Env Vars in Gitpod

We can persist env vars into gitpod by storing them in you account's Gitpod Environement [Variable list](https://gitpod.io/user/variables)
https://www.gitpod.io/docs/configure/projects/environment-variables

```
gp env HELLO='world'
```

All future workspaces launched will set the env vars for all bash terminals opened in those workspaces.

You can also set env vars in the `.gitpod.yml` but this can only contain non-senstive env vars.

### AWS CLI Installation

[AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html) is installed for the project via the bash script [`./bin/install_aws_cli`](./bin/install_aws_cli)


[Getting Started Install (AWS CLI)](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
[AWS CLI Env Vars](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)

We can check if our AWS credentials is configured correctly by running the following AWS CLI command:

```sh
aws sts get-caller-identity
```

If it is succesful you should see a json payload return that looks like this:

```json
{
    "UserId": "ABCDEF12ZPVHJ5WIJ5KR",
    "Account": "123456789012",
    "Arn": "arn:aws:iam::123456789012:user/terraform-beginner-bootcamp"
}
```

We will need to generate AWS CLI credentials for the [IAM User](https://us-east-1.console.aws.amazon.com/iamv2/home?region=eu-central-1#/users) in order to enable the user's access to AWS CLI.

## Terraform Basics

### Terraform Registry

Terraform sources their providers and modules from the Terraform registry which is located at [registry.terraform.io](https://registry.terraform.io/)

- **Providers** is an interface to APIs that will allow to create resources in terraform.
- **Modules** are a way to make large amount of terraform code modular, portable and sharable.

[Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)

> `main.tf` is a Terraform root level module.

### Terraform Console

We can see a list of all the Terrform commands by simply typing `terraform`


#### Terraform Init

At the start of a new terraform project we will run `terraform init` to download the binaries for the Terraform providers that we'll use in this project.

#### Terraform Plan

`terraform plan`

This will generate out a changeset, about the state of our infrastructure and what will be changed.

We can output this changeset ie. "plan" to be passed to an apply, but often you can just ignore outputting.

#### Terraform Apply

`terraform apply`

This will run a plan and pass the changeset to be execute by terraform. Apply should prompt yes or no.

If we want to automatically approve an apply we can provide the auto approve flag eg. `terraform apply --auto-approve`

#### Terraform Destroy

`teraform destroy`
This will destroy resources.

You can also use the auto approve flag to skip the approve prompt. eg. `terraform destroy --auto-approve`

### Terraform Lock Files

`.terraform.lock.hcl` contains the locked versioning for the providers or modulues that should be used with this project.

The Terraform Lock File **should be committed** to your Version Control System (VSC) eg. Github

### Terraform State Files

Running terraform apply generates a file named `.terraform.tfstate`. This file contains information about the current state of your infrastructure.

This file **should not be commited** to your VCS, as the file can contain sensentive data.

If you lose this file, you lose knowning the state of your infrastructure.

`.terraform.tfstate.backup` is the previous state file state.

### Terraform Directory

`.terraform` directory contains binaries of Terraform providers.

### AWS Credentials
Terraform needs our AWS credentials in order to create the defined AWS resource (S3 bucket in this case).
These credentails can be defined in the `main.tf` file, but this is **not recommended**, since the main.tf will be commited to our code repository. The better way to safely use credentials is to set them as `env vars`.

#### Terraform flow for reading credentials:
(1) check config file _if not present, then_ -> (2) read from env vars

#### Error while implementing Random Provider
S3 buckets have specific naming conventions. During our implementation we ran into an error, since our random bucket name generator allowed `UPPERCASE alphabets`. To tackle this, we explicitely denied the presence of uppercase letters in the code.
```tf
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
Refer the [documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html) for all rules relating to the S3 bucket naming.

### Terraform Cloud Basics

> In this bootcamp we will be utilizing the Terraform cloud free tier, which allows upto 500 resources free per month.
[Terraform Cloud Pricing](https://www.hashicorp.com/products/terraform/pricing)

![pricing](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/7f19b23b71dbc5815dfeec76897bbc8ba19b8996/journal/assets/week-0/41-TFC-pricing.png)

#### Terraform Workspace v/s Terraform Project
- **Workspace** : A container in the Terraform Cloud for infrastructure state, configuration and settings.
- **Project** : A conceptual way to group Terraform workspaces together, which server a particular purpose/goal.

![workspace-vs-proj](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/7f19b23b71dbc5815dfeec76897bbc8ba19b8996/journal/assets/week-0/42-TFC-workspace-vs-project.png)

#### Types of Terraform Workspace Workflows
Version control workflow: Stores Terraform configs in a Git repo and triggers runs  based on PRs. (This related to the Gitpods flow, and not in scope of the bootcamp.)
CLI-driven workflow: Triggers remote Terraform runs from local command line. (example: Gitpod). This is the type we will be using in our bootcamp.
API-driven workflow: It is an advanced option for programmatic coding, it integrates Terraform with a larger pipeline using Terraform API. (also not in scope of the bootcamp.)

#### terraform login and credentails.trfc.json file
`terraform login`
The terraform [login](https://developer.hashicorp.com/terraform/cli/commands/login) command can be used to automatically obtain and save an API token for Terraform Cloud, Terraform Enterprise, or any other host that offers Terraform services

`credentails.trfc.json`
- Terraform will obtain an API token and save it in plain text in a local CLI configuration file called credentials.tfrc.json. 
- When you run terraform login, it will explain specifically where it intends to save the API token and give you a chance to cancel if the current configuration is not as desired.
- The path in Gitpod is `/home/gitpod/.terraform.d/credentails.trfc.json`

### Issues with Terraform Cloud Login and Gitpod Workspace

When attempting to run `terraform login` it will launch bash a wiswig view to generate a token. However it does not work expected in Gitpod VsCode in the browser.

The workaround is manually generate a token in Terraform Cloud

```
https://app.terraform.io/app/settings/tokens?source=terraform-login
```

Then create open the file manually here:

```sh
touch /home/gitpod/.terraform.d/credentials.tfrc.json
open /home/gitpod/.terraform.d/credentials.tfrc.json
```

Provide the following code (replace your token in the file):

```json
{
  "credentials": {
    "app.terraform.io": {
      "token": "YOUR-TERRAFORM-CLOUD-TOKEN"
    }
  }
}
``````

We have automated this workaround with the following bash script  [bin/generate_tfrc_credentials](bin/generate_tfrc_credentials)

### TF alias
We wish to set an alias for terraform, so instead of typing `terraform` to execute each command, we will be able to use `tf`. example: `tf init`, `tf apply` etc.
We will be setting this alias in our bash profile `~/.bash_profile` as below:

`alias tf="terraform"`

>The `~/.bash_profile` file is read automatically when a new shell is loaded. So if we wish to execute it in the current shell, we can do so by executing `./bash_profile``

-----------------------------------------------------------------------------------------------------
# Personal Documentation 
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

==============================================================================

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

==============================================================================

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

==============================================================================

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

==============================================================================

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

==============================================================================

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

==============================================================================

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

==============================================================================

8. #### Terraform Login Workaround
8.1 Create the issue.
```txt
Issue name: Generate TFRC
Issue Description:
- [] Create a bash script using ChatGPT to create trfc file.
- [] Create new token for 30 days in Terraform Cloud

Issue label: enhancement
```

8.2 Create a branch for the issue and launch in Gitpod.

8.3 Add an example for `TERRAFORM_CLOUD_TOKEN` in `.env.example1
TERRAFORM_CLOUD_TOKEN='YOUR SECRET TERRAFORM CLOUD TOKEN'

8.4 Login to terraform cloud settings and create a token for the duration of this bootcamp.
![terratowns-token](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/09c04945dc73718a185c7c16a50781f447ea05be/journal/assets/week-0/53-terratowns-token.png)

8.5  Set this is Gitpod env variable:

`gp env TERRAFORM_CLOUD_TOKEN='YOUR SECRET TERRAFORM CLOUD TOKEN'`

`export TERRAFORM_CLOUD_TOKEN='YOUR SECRET TERRAFORM CLOUD TOKEN'`

8.6 Next, we will create a bash script that will refer our variable.
- Create script `generate_tfrc_credentials` under the `bin` folder
```sh
#!/usr/bin/env bash

# Define target directory and file
TARGET_DIR="/home/gitpod/.terraform.d"
TARGET_FILE="${TARGET_DIR}/credentials.tfrc.json"

# Check if TERRAFORM_CLOUD_TOKEN is set
if [ -z "$TERRAFORM_CLOUD_TOKEN" ]; then
    echo "Error: TERRAFORM_CLOUD_TOKEN environment variable is not set."
    exit 1
fi

# Check if directory exists, if not, create it
if [ ! -d "$TARGET_DIR" ]; then
    mkdir -p "$TARGET_DIR"
fi

# Generate credentials.tfrc.json with the token
cat > "$TARGET_FILE" << EOF
{
  "credentials": {
    "app.terraform.io": {
      "token": "$TERRAFORM_CLOUD_TOKEN"
    }
  }
}
EOF

echo "${TARGET_FILE} has been generated."

```

- Grant the script executable permissions `chmod u+x ./bin/generate_tfrc_credentials`

- Execute the script `./bin/generate_tfrc_credentials`

![tfrc-creds](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/09c04945dc73718a185c7c16a50781f447ea05be/journal/assets/week-0/54-tfrc-creds.png)

8.7 Update the `.gitpod.yml` file to call out script.
`source ./bin/generate_tfrc_credentials`

The final file should look like:
```yml
tasks:
  - name: terraform
    before: |
      source ./bin/install_terraform_cli
      source ./bin/generate_tfrc_credentials
  - name: aws-cli
    env:
      AWS_CLI_AUTO_PROMPT: on-partial
    before: |
      source ./bin/install_aws_cli

vscode:
  extensions:
    - amazonwebservices.aws-toolkit-vscode
    - hashicorp.terraform
```


8.8 Stage, commit and sync the changes

8.9 PR and Merge `13-terraform-cloud-backend` to `main`

8.10 Add Tags

```sh
git checkout main
git pull
git tag 0.8.0
git push --tags
```
8.11 Stop the Gitpod Workspace

==============================================================================

9. #### TF Alias

9.1 Create an `issue`
```txt
Issue name: TF alias for TerraformTerraform Cloud Backend
Desc:
[ ] Set an alias for terraform to be tf in our bash profile.
 
Label: Enhancement
```

9.2 Create a branch for the issue and launch in Gitpod

9.3 We wish to set an alias for terraform, so instead of typing `terraform` to execute each command, we will be able to use `tf`. example: `tf init`, `tf apply` etc.
We will be setting this alias in our bash profile `~/.bash_profile` as below:

`alias tf="terraform"`

> The `~/.bash_profile` file is read automatically when a new shell is loaded. So if we wish to execute it in the current shell, we can do so by executing `./bash_profile`

9.4 Now test it by executing `tf`

![tf-alias](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/09c04945dc73718a185c7c16a50781f447ea05be/journal/assets/week-0/55-tf-alias.png)

> We have created the alias in our bash profile file. `./bash_profile` is a file read by the `bash terminal`. We want the file to be read even for programmatic access, this can be achieved by creating a shell script.

 9.5 Create script `set_tf_alias` inside `bin`

```sh
#!/usr/bin/env bash

# Check if the alias already exists in the .bash_profile
grep -q 'alias tf="terraform"' ~/.bash_profile

# $? is a special variable in bash that holds the exit status of the last command executed
if [ $? -ne 0 ]; then
    # If the alias does not exist, append it
    echo 'alias tf="terraform"' >> ~/.bash_profile
    echo "Alias added successfully."
else
    # Inform the user if the alias already exists
    echo "Alias already exists in .bash_profile."
fi

# Optional: source the .bash_profile to make the alias available immediately
source ~/.bash_profile
```

9.6 Grant the script executable permissions `chmod u+x ./bin/set_tf_alias`. 

9.7 Execute the script `./bin/set_tf_alias`

![tfrc-alias-script](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/09c04945dc73718a185c7c16a50781f447ea05be/journal/assets/week-0/56-tfrc-alias-script.png)

9.8 Incorporate the script into `.gitpod.yml`
```yml
tasks:
  - name: terraform
    before: |
      source ./bin/set_tf_alias
      source ./bin/install_terraform_cli
      source ./bin/generate_tfrc_credentials
  - name: aws-cli
    env:
      AWS_CLI_AUTO_PROMPT: on-partial
    before: |
      source ./bin/set_tf_alias
      source ./bin/install_aws_cli

vscode:
  extensions:
    - amazonwebservices.aws-toolkit-vscode
    - hashicorp.terraform
```

9.9 Stage, commit and sync the changes

9.10 PR and Merge `17-tf-alias-for-terraform` to `main`

9.11 Add Tags

```sh
git checkout main
git pull
git tag 0.9.0
git push --tags
```
9.12. Stop the Gitpod Workspace

==============================================================================

10. #### Project Validation
week-0 code passed the validation.
![validation-CF-stack](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/7f19b23b71dbc5815dfeec76897bbc8ba19b8996/journal/assets/week-0/57-validation-CF-stack.png)

![week0-validation-pass](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/7f19b23b71dbc5815dfeec76897bbc8ba19b8996/journal/assets/week-0/58-week0-validation-pass.png)
    
-----------------------------------------------------------------------------------------------------
