# Terrahome AWS Multi Home Refactor

We altered the paths and variables associated to the content to our static Websites, in order to accomodate the creation of more than one home.

### Steps to add a new home

- Add your home to the `terraform.tfvars.example` and the `terraform.tfvars` making sure to change the block for your new content
    ```tf
    YOURHOME = {
      public_path = "/workspace/terraform-beginner-bootcamp-2023/public/YOURHOMEDIRECTORY"
      content_version = 1 
    }
    ```
- Add you new module to `main.tf` making sure to change the block for your new content
    ```
    module "home_YOURHOME" {
      source = "./modules/terrahome_aws"
      user_uuid = var.teacherseat_user_uuid
      public_path = var.YOURHOME.public_path
      content_version = var.YOURHOME.content_version
    }
    ```
- Add your new resource to `main.tf` making sure to change the block for your new content
    ```
    resource "terratowns_home" "YOURHOME" {
      name = "Travle Blog"
      description = <<DESCRIPTION
    add description here
    DESCRIPTION
      town = "TOWN"
      content_version = var.YOURHOME.content_version
      domain_name = module.YOURNEWMODULE.domain_name
    }
    ```
- Create a new directory under `./public` with your home name
    - Must include `index.html`
    - Must include `error.html`
    - Must include `/assets/`
        - **Only the top level of this directory will be uploaded. Sub-directories will be ignored**

- `terraform init`
- `terraform plan`
- `terraform apply --auto-approve`