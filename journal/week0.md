# Week 0 — Prep Week Journal

## Task Status


#### 1.  Branching Tagging PR
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


#### 2.  Terraform CLI Refactor 

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

#### 3. Project Root Env Var

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

3.1  Stop the `Gitpod workspace`.
