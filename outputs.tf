output "bucket_name_1" {
  description = "Bucket name for our static website hosting"
  value = module.home_krakow_hosting.bucket_name
}

output "s3_website_endpoint_1" {
  description = "S3 Static Website hosting endpoint"
  value = module.home_krakow_hosting.website_endpoint
}

output "cloudfront_url_1" {
  description = "The CloudFront Distribution Domain Name"
  value = module.home_krakow_hosting.domain_name
}

output "bucket_name_2" {
  description = "Bucket name for our static website hosting"
  value = module.home_recipe_hosting.bucket_name
}

output "s3_website_endpoint_2" {
  description = "S3 Static Website hosting endpoint"
  value = module.home_recipe_hosting.website_endpoint
}

output "cloudfront_url_2" {
  description = "The CloudFront Distribution Domain Name"
  value = module.home_recipe_hosting.domain_name
}