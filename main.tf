terraform {
    required_providers {
    terratowns = {
      source = "local.providers/local/terratowns"
      version = "1.0.0"
    }
  }
  #cloud {
  #  organization = "aggarwaltanushree"
  #  workspaces {
  #    name = "terra-house-1"
  #  }
  #}
}

provider "terratowns" {
  endpoint = var.terratowns_endpoint
  user_uuid = var.teacherseat_user_uuid
  token = var.terratowns_access_token
}

module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.teacherseat_user_uuid
  index_html_filepath = var.index_html_filepath
  error_html_filepath = var.error_html_filepath
  content_version = var.content_version
  assets_path = var.assets_path
}

resource "terratowns_home" "home" {
  name = "How to play Pokemon in 2023!"
  description = <<DESCRIPTION
Pokémon GO is an Adventure game developed by Niantic, Inc.
BlueStacks app player is the best platform to play Android games on your PC or 
Mac for an immersive gaming experience.
Explore the world alongside Trainers from all over as they uncover Pokémon locations.
DESCRIPTION
  domain_name = module.terrahouse_aws.cloudfront_url
  # Mock CDN below:
  #domain_name = "3fdq3gz.cloudfront.net"
  #town = "gamers-grotto"
  town = "missingo"
  content_version = 1
}