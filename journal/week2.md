# Terraform Beginner Bootcamp 2023  Week 2  — Building Week Journal

[Week-2 Architecture](#week-2-architecture) :cloud: :evergreen_tree:

[Notes for revision](#notes-for-revision) :syringe: :medal_sports:

[Personal Documentation](#personal-documentation) :memo: :pencil:

[Task Status](#task-status) :star: :face_in_clouds: :partying_face: :white_check_mark:

# Task Status
| Topic | Status |
| --- | --- | 
| [Week 2 Diagramming](#week-2-diagramming)  | <ul><li> [x] </li></ul> |
| [Setting Up Terraform Mock Server](#setting-up-terraform-mock-server) | <ul><li> [x] </li></ul> |
| [Setup Skeleton For Custom Terraform Provider](#setup-skeleton-for-custom-terraform-provider) | <ul><li> [ ] </li></ul> |
| [Provider Block For Custom Terraform Provider](#provider-block-for-custom-terraform-provider) | <ul><li> [ ] </li></ul> |
| [Resource Skeleton](#resource-skeleton) | <ul><li> [ ] </li></ul> |
| [Implementing CRUD](#implementing-crud) | <ul><li> [ ] </li></ul> |
| [Terraform Cloud And Multi Home Refactor](#terraform-cloud-and-multi-home-refactor) | <ul><li> [ ] </li></ul> |
| [Project Validation](#project-validation) | <ul><li> [ ] </li></ul> |



# Week 2 Architecture
![week2-architecture](https://github.com/aggarwal-tanushree/terraform-beginner-bootcamp-2023/blob/6cc6fc3bb3ca792cc74254780928e9d10a4ef099/journal/assets/week-2/Final-ArchitecturalDiagram.png)

The complete end-to-end architecture of what we will achieve at the end of the bootcamp is available [here](https://lucid.app/lucidchart/e3f15b1a-2211-4ddb-8c95-f144c2504db4/edit?invitationId=inv_0873b3c6-c652-463f-9f2b-fa0f1b420823&page=0_0#) 

_Diagram copyrights: Andrew Brown from ExamPro.co_

-----------------------------------------------------------------------------------------------------
# Notes For Revision

## Our Custom Provider - explained
**Custom provider** named `terraform-provider-terratowns` (written in **golang**). 

![physical-diagram](assets/week-2/Week2-Physical-Architectural-Diagram.png)

The custom provider has a **resource** called `Home` which has four actions associated with it, namely:
- Read
- Update
- Create
- Delete

We have two servers:
i)  `TerraTowns` : `Development (Mock) Server` : `sinatra` server :  `localhost:4567`
ii) `TerraTowns.cloud` : `Production Server` : `rails`

Sinatra: light-weight web server
Rails : heavy-duty production server

We will use `bash scripts` under the path `/bin/terratowns/` to mock each of the four HTTP requests : 
i)   `create` 
ii)  `read`
iii) `update`
iv)  `delete`

|   	| TF Resource Action 	| HTTP Request Type 	|          API endpoint         	|
|---	|:------------------:	|:-----------------:	|:-----------------------------:	|
| 1 	| Create             	| POST              	| /api/u/:user_uuid/homes       	|
| 2 	| Read               	| GET               	| /api/u/:user_uuid/:uuid       	|
| 3 	| Update             	| PUT               	| /api/u/:user_uuid/homes/:uuid 	|
| 4 	| Delete             	| DELETE            	| /api/u/:user_uuid/homes/:uuid 	|



## Working with Ruby

### Bundler

- Bundler is a package manager for Ruby, i.e is the tool used to manage gem dependencies in Ruby applications
- It is the primary way to install ruby packages (known as **gems**) for Ruby.
- Bundler helps ensure that the necessary gems are available in development, staging, and production environments, preventing what's commonly known as "dependency hell"

#### Install Gems

You need to create a `Gemfile` and define your gems in that file.

```rb
source "https://rubygems.org"

gem 'sinatra'
gem 'rake'
gem 'pry'
gem 'puma'
gem 'activerecord'
```

Then you need to run the `bundle install` command

This will install the gems on the system globally (unlike nodejs which install packages in place in a folder called node_modules)

A `Gemfile.lock` will be created to lock down the gem versions used in this project.

#### Executing ruby scripts in the context of bundler

We have to use `bundle exec` to tell future ruby scripts to use the gems we installed. This is the way we set context.

### Sinatra

- Sinatra is a DSL (Domain-Specific Language) for rapidly creating web applications in the Ruby programming language. 

- Sinatra is a micro web-framework for ruby to build web-apps.

- It is designed to be lightweight and easy to use, making it an ideal choice for small to medium-sized projects.

- It is great for mock or development servers or for very simple projects.

- You can create a web-server in a single file. 

https://sinatrarb.com/


## Terratowns Mock Server

### Running the web server

We can run the web server by executing the following commands:

```rb
bundle install
bundle exec ruby server.rb
```

All of the code for our server is stored in a single file `terratowns_mock_server/server.rb`.


### HTTP Requests
#### Anatomy of a HTTP request

![http-request-anatomy](assets/week-2/anatomy-of-a-request.png)

#### Bearer Authentication
Bearer authentication (also called token authentication) is an HTTP authentication scheme that involves security tokens called bearer tokens. 
The name “Bearer authentication” can be understood as "give access to the bearer of this token."
The bearer token is a cryptic string, usually generated by the server in response to a login request. 
The client must send this token in the Authorization header when making requests to protected resources: `Authorization: Bearer <token>`
https://swagger.io/docs/specification/authentication/bearer-authentication/



#### HTTP error codes
https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses

Our error codes:
a1001 - token/code/access-code does not match
a1002 - user_uuid not passed in HTTP header
a1003 - token/code/acess-code and user_uuid

-----------------------------------------------------------------------------------------------------

# Personal Documentation 
1. ## Week 2 Diagramming

2. ## Terratowns Mock Server
2.1 Create a new issue in your Github repositiory.

```txt
Issue name: Terratowns Mock Server
Issue description: Download Terratowns mock server into our repo
https://github.com/ExamProCo/terratowns_mock_server

Label: enhancement
```

2.2 Create a branch for this issue and launch it in Gitpod.

2.3 We will now copy the **mock server** code from its Git repository.
`git clone https://github.com/ExamProCo/terratowns_mock_server.git`

```sh
gitpod /workspace/terraform-beginner-bootcamp-2023 (39-terratowns-mock-server) $ git clone https://github.com/ExamProCo/terratowns_mock_server.git
Cloning into 'terratowns_mock_server'...
remote: Enumerating objects: 20, done.
remote: Counting objects: 100% (20/20), done.
remote: Compressing objects: 100% (14/14), done.
remote: Total 20 (delta 7), reused 15 (delta 6), pack-reused 0
Receiving objects: 100% (20/20), 9.34 KiB | 9.34 MiB/s, done.
Resolving deltas: 100% (7/7), done.
```

2.4 We need to delete the `.git` directory, else the imported code will be treated as a sub-module.
```sh
cd terratowns_mock_server
ls -al
rm -rf .git
```
```sh
gitpod /workspace/terraform-beginner-bootcamp-2023 (39-terratowns-mock-server) $ cd terratowns_mock_server
gitpod /workspace/terraform-beginner-bootcamp-2023/terratowns_mock_server (main) $ ls -la
total 44
drwxr-xr-x 4 gitpod gitpod   150 Oct  3 10:46 .
drwxr-xr-x 8 gitpod gitpod  4096 Oct  3 10:46 ..
drwxr-xr-x 2 gitpod gitpod    60 Oct  3 10:46 bin
-rw-r--r-- 1 gitpod gitpod   126 Oct  3 10:46 Gemfile
-rw-r--r-- 1 gitpod gitpod  1069 Oct  3 10:46 Gemfile.lock
drwxr-xr-x 8 gitpod gitpod   163 Oct  3 10:46 .git
-rw-r--r-- 1 gitpod gitpod  1205 Oct  3 10:46 .gitignore
-rw-r--r-- 1 gitpod gitpod    92 Oct  3 10:46 .gitpod.yml
-rw-r--r-- 1 gitpod gitpod 11357 Oct  3 10:46 LICENSE
-rw-r--r-- 1 gitpod gitpod    76 Oct  3 10:46 README.md
-rw-r--r-- 1 gitpod gitpod  4926 Oct  3 10:46 server.rb
gitpod /workspace/terraform-beginner-bootcamp-2023/terratowns_mock_server (main) $ rm -rf .git
gitpod /workspace/terraform-beginner-bootcamp-2023/terratowns_mock_server (39-terratowns-mock-server) $ ls -la
total 44
drwxr-xr-x 3 gitpod gitpod   138 Oct  3 10:47 .
drwxr-xr-x 8 gitpod gitpod  4096 Oct  3 10:46 ..
drwxr-xr-x 2 gitpod gitpod    60 Oct  3 10:46 bin
-rw-r--r-- 1 gitpod gitpod   126 Oct  3 10:46 Gemfile
-rw-r--r-- 1 gitpod gitpod  1069 Oct  3 10:46 Gemfile.lock
-rw-r--r-- 1 gitpod gitpod  1205 Oct  3 10:46 .gitignore
-rw-r--r-- 1 gitpod gitpod    92 Oct  3 10:46 .gitpod.yml
-rw-r--r-- 1 gitpod gitpod 11357 Oct  3 10:46 LICENSE
-rw-r--r-- 1 gitpod gitpod    76 Oct  3 10:46 README.md
-rw-r--r-- 1 gitpod gitpod  4926 Oct  3 10:46 server.rb
gitpod /workspace/terraform-beginner-bootcamp-2023/terratowns_mock_server (39-terratowns-mock-server) $ 
```

2.5 Navigate to `terratowns_mock_server/.gitpod.yml` and cut all the contents of this file. Paste this in `./gitpod.yml` at the project root level.
_place the code as the second task. Add a `cd` command to change dir to terratowns-mock-server_

```yml
  - name: sinatra
    before: | 
      cd $PROJECT_ROOT
      cd terratowns_mock_server
      bundle install
      bundle exec ruby server.rb 
```

Add a `cd $PROJECT_ROOT` for all the tasks defined in the file.

```yml
tasks:
  - name: terraform
    before: |
      cd $PROJECT_ROOT
      source ./bin/set_tf_alias
      source ./bin/install_terraform_cli
      source ./bin/generate_tfrc_credentials
      cp $PROJECT_ROOT/terraform.tfvars.example $PROJECT_ROOT/terraform.tfvars
  - name: sinatra
    before: | 
      cd $PROJECT_ROOT
      cd terratowns_mock_server
      bundle install
      bundle exec ruby server.rb 
  - name: aws-cli
    env:
      AWS_CLI_AUTO_PROMPT: on-partial
    before: |
      cd $PROJECT_ROOT
      source ./bin/set_tf_alias
      source ./bin/install_aws_cli
  - name: http-server
    before: |
      cd $PROJECT_ROOT
      npm install --global http-server
    command: 
      http-server

vscode:
  extensions:
    - amazonwebservices.aws-toolkit-vscode
    - hashicorp.terraform
    - phil294.git-log--graph
    - mhutchie.git-graph
```

2.6 Delete `terratowns_mock_server/.gitpod.yml` since its no longer required

2.7 Rename `terratowns-mock-server/bin` to `terratowns-mock-server/terratowns`. Drag and move this directory  under`./bin`

2.8 Grant executable permissions to these files.
```sh
cd ~
chmod u+x bin/terratowns/*
```

```sh
gitpod /workspace/terraform-beginner-bootcamp-2023 (39-terratowns-mock-server) $ pwd
/workspace/terraform-beginner-bootcamp-2023
gitpod /workspace/terraform-beginner-bootcamp-2023 (39-terratowns-mock-server) $ chmod u+x bin/terratowns/*
gitpod /workspace/terraform-beginner-bootcamp-2023 (39-terratowns-mock-server) $ ls -l bin/terratowns/*
-rwxr-xr-x 1 gitpod gitpod 1051 Oct  3 10:46 bin/terratowns/create
-rwxr-xr-x 1 gitpod gitpod  790 Oct  3 10:46 bin/terratowns/delete
-rwxr-xr-x 1 gitpod gitpod  888 Oct  3 10:46 bin/terratowns/read
-rwxr-xr-x 1 gitpod gitpod 1102 Oct  3 10:46 bin/terratowns/update
```

2.9 Start the Sinatra server (since we have configured this in our `.gitpod.yml` file, the server will automatically load when we restart our environment. 
Since we do not wish to restart the env right now, we can execute this manually, by:
At the terminal:
```sh
cd terratowns_mock_server/
bundle install
bundle exec ruby server.rb
```

cd terratowns_mock_server/
gitpod /workspace/terraform-beginner-bootcamp-2023/terratowns_mock_server (39-terratowns-mock-server) $ bundle install
Bundler 2.4.20 is running, but your lockfile was generated with 2.4.19. Installing Bundler 2.4.19 and restarting using that version.
Fetching gem metadata from https://rubygems.org/.
Fetching bundler 2.4.19
Installing bundler 2.4.19
Fetching gem metadata from https://rubygems.org/.......
Fetching method_source 1.0.0
Fetching rack 2.2.8
Fetching mustermann 3.0.0
Fetching concurrent-ruby 1.2.2
Fetching minitest 5.20.0
Fetching coderay 1.1.3
Fetching nio4r 2.5.9
Fetching tilt 2.2.0
Installing method_source 1.0.0
Installing tilt 2.2.0
Installing mustermann 3.0.0
Installing minitest 5.20.0
Installing coderay 1.1.3
Installing nio4r 2.5.9 with native extensions
Installing rack 2.2.8
Installing concurrent-ruby 1.2.2
Fetching pry 0.14.2
Fetching rack-protection 3.1.0
Installing rack-protection 3.1.0
Fetching tzinfo 2.0.6
Fetching i18n 1.14.1
Installing pry 0.14.2
Installing i18n 1.14.1
Installing tzinfo 2.0.6
Fetching sinatra 3.1.0
Installing sinatra 3.1.0
Fetching activesupport 7.0.8
Installing activesupport 7.0.8
Fetching activemodel 7.0.8
Installing activemodel 7.0.8
Fetching activerecord 7.0.8
Installing activerecord 7.0.8
Fetching puma 6.3.1
Installing puma 6.3.1 with native extensions
Bundle complete! 5 Gemfile dependencies, 20 gems now installed.
Use `bundle info [gemname]` to see where a bundled gem is installed.
gitpod /workspace/terraform-beginner-bootcamp-2023/terratowns_mock_server (39-terratowns-mock-server) $ bundle exec ruby server.rb
== Sinatra (v3.1.0) has taken the stage on 4567 for development with backup from Puma
Puma starting in single mode...
* Puma version: 6.3.1 (ruby 3.2.2-p53) ("Mugi No Toki Itaru")
*  Min threads: 0
*  Max threads: 5
*  Environment: development
*          PID: 4893
* Listening on http://127.0.0.1:4567
* Listening on http://[::1]:4567
Use Ctrl-C to stop
```

2.10 Let's execute our bash scripts manually to test if they are working.
At the terminal, execute: 
`./bin/terratowns/create`
_it should o/p a uuid. Also check in the `sinatra` tab for a HTTP request of type POST_

```sh
gitpod /workspace/terraform-beginner-bootcamp-2023 (39-terratowns-mock-server) $ ./bin/terratowns/create
{"uuid":"535de2f3-3549-4501-8b5e-4046f2f64b4b"}
```

Let's emulate `read`:
`./bin/terratowns/read <house uuid returned by create>`

```sh
gitpod /workspace/terraform-beginner-bootcamp-2023 (39-terratowns-mock-server) $ ./bin/terratowns/read 535de2f3-3549-4501-8b5e-4046f2f64b4b
{
  "uuid": "535de2f3-3549-4501-8b5e-4046f2f64b4b",
  "name": "New House",
  "town": "cooker-cove",
  "description": "A new house description",
  "domain_name": "3xf332sdfs.cloudfront.net",
  "content_version": 1
}
```

`update`:
`./bin/terratowns/update <house uuid returned by create>`

```sh
gitpod /workspace/terraform-beginner-bootcamp-2023 (39-terratowns-mock-server) $ ./bin/terratowns/update 535de2f3-3549-4501-8b5e-40
46f2f64b4b
{"uuid":"535de2f3-3549-4501-8b5e-4046f2f64b4b"}
```


`destroy`:
`./bin/terratowns/delete <house uuid returned by create>`

```sh
gitpod /workspace/terraform-beginner-bootcamp-2023 (39-terratowns-mock-server) $ ./bin/terratowns/delete 535de2f3-3549-4501-8b5e-40
46f2f64b4b
{
  "uuid": "535de2f3-3549-4501-8b5e-4046f2f64b4b"
}
```

all our HTTP actions also being tracked in our Sinatra server :
```sh
== Sinatra (v3.1.0) has taken the stage on 4567 for development with backup from Puma
Puma starting in single mode...
* Puma version: 6.3.1 (ruby 3.2.2-p53) ("Mugi No Toki Itaru")
*  Min threads: 0
*  Max threads: 5
*  Environment: development
*          PID: 4893
* Listening on http://127.0.0.1:4567
* Listening on http://[::1]:4567
Use Ctrl-C to stop
# create - POST /api/homes
name New House
description A new house description
domain_name 3xf332sdfs.cloudfront.net
content_version 1
town cooker-cove
uuid 535de2f3-3549-4501-8b5e-4046f2f64b4b
# read - GET /api/homes/:uuid
# update - PUT /api/homes/:uuid
# delete - DELETE /api/homes/:uuid
```

2.11 Update the documentation

2.12 Stage, commit and sync the changed to Github

2.13  Create a PR and Merge this branch `39-terratowns-mock-server` to the `main` branch.

2.14 Issue tags to the `main branch` as `2.0.0`

