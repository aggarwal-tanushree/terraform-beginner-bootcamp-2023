# Terraform Beginner Bootcamp 2023

# Week-0 Learnings
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

# Detailed step-by-step documentation
The complete step-by-step manual for week-0 is available [here](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/b652a3bdd02574dc026f3fab198aa20d51700f85/journal/week0.md)
-------------------------------------------------------------------------------------------------------------
