# Week 0 — Prep Week Journal

## Task Status

#### 1. Branching Tagging PR
1.1 Open your bootcamp repository in [Github](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023)

1.2 Navigate to the `Issues` tab and click on `New Issues`
![issues](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/0-issues.png)

1.3 Give the issue a name and description:
```txt
Add semantic documentation to project

We want to add semantic versioning in our Terraform bootcamp project.
```
then click `Submit new issue`
![issue_create](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/1-issue-create.png)

1.4 Click on the newly created issue to view its details. Notice that an _issue number_ has been auto generated.
![view_issue](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/2-issue-num.png)

1.5 On the right you will notice an option to create a `branch` for this `issue`. Click the option.
![create_branch](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/3-create-pr.png)

1.6 Give the branch a meaningful `name` and select the option to `Checkout locally`
![create_branch_1](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/4-create-branch-name.png)

Click `Create branch`

1.7 Copy the command returned. This will be used to checkout the newly created branch in Gitpod.
![cpy_cmd](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/5-checkout-cmd.png)

`git checkout 1-add-semantic-documentation-to-project`

1.8 Launch the repo in `Gitpod`
![launch_gitpod](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/6-launch-gitpod.png)

1.9 At the Gitpod terminal, type the command copied in step-1.7 to checkout the issue branch
`git checkout 1-add-semantic-documentation-to-project`

![chkout_branch](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/7-checkout-branch.png)

1.10 Open the `README.md` file and update it with the following information:

```md
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
```

1.11 Stage and Commit the file to your Gihub repo.
**Note: make sure you enter a descriptive comment message**
example: `"#1 Add semantic versioning to project"`

![commit-branch](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/8-commit-branch.png)

![sync-branch](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/9-sync-branch.png)

1.12 Add and push `Tags` to the Github repo using the below commands:

```sh
git tag 0.1.0
git push --tags
```

![tags](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/10-tagging.png)

1.13 Open the bootcamp repo in Github and verify that the commit was synced.

1.14 Next, we need to bring this feature back to the `main` branch. For this we will be creating a `pull request`.

In Github, navigate to the `Pull requests` tab, and click on `New pull request`.

![new_pull](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/11-new-pull.png)

1.15 Select `base: main` and `compare: add-semantic-documentation` as the branches.
Notice the message:  ✔️`Able to merge` . This indicates that we are Git will be able to merge our feature branch into the `main` branch.

Click `Create pull request` 

![pull_create](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/12-pull-create.png)

1.16 Next, we `merge` the `pull request`

_Note: In real-world projects, someone else from the team review's the code before merging it. Since bootcamp is for individual submissions, we will be skipping this step and merging the branch ourself without an reviewer_

![merge](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/13-merge.png)

Confirm the `Merge` to complete the merging with `main` branch and `close the issue`.

![merge_close](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/9a4995651b8bd537f3bb9d3aff5cfe46a700c51d/journal/assets/week0/14-merge-close-issue.png)

_Note: in production scenarios, ideally branches are deleted after they have been inspected and merged. For the bootcamp we will be retaining the individual branches for grading verification . _


#### Terraform CLI Refactor 
