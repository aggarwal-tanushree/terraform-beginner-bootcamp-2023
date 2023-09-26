output "bucket_name" {
  value = aws_s3_bucket.website_bucket.bucket
}
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_website_configuration#website_endpoint
output "website_endpoint" {
  value = aws_s3_bucket_website_configuration.website_configuration.website_endpoint
}