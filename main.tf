terraform {
    required_providers {
    terratowns = {
      source = "local.providers/local/terratowns"
      version = "1.0.0"
    }
  }
  cloud {
    organization = "aggarwaltanushree"
    workspaces {
      name = "terra-house-1"
    }
  }
}

provider "terratowns" {
  endpoint = var.terratowns_endpoint
  user_uuid = var.teacherseat_user_uuid
  token = var.terratowns_access_token
}

module "home_krakow_hosting" {
  source = "./modules/terrahome_aws"
  user_uuid = var.teacherseat_user_uuid
  public_path = var.krakow.public_path
  content_version = var.krakow.content_version
}

resource "terratowns_home" "home" {
  name = "A weekend getway to Krakow!"
  description = <<DESCRIPTION
Last weekend I took a short trip to Krakow,
a city in Poland. This page showcases some of the interesting spots I 
visited in Krakow.
DESCRIPTION
  domain_name = module.home_krakow_hosting.domain_name
  town = "the-nomad-pad"
  content_version = var.krakow.content_version
}

module "home_recipe_hosting" {
  source = "./modules/terrahome_aws"
  user_uuid = var.teacherseat_user_uuid
  public_path = var.recipe.public_path
  content_version = var.recipe.content_version
}

resource "terratowns_home" "home_recipe" {
  name = "How to annoy your husband - A Cookbook"
  description = <<DESCRIPTION
This is an original cookbook on how to annoy your husband.
DESCRIPTION
  domain_name = module.home_recipe_hosting.domain_name
  town = "cooker-cove"
  content_version = var.recipe.content_version
}