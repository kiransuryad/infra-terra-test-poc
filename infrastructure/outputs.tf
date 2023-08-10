output "vpc_id" {
  description = "The ID of the VPC"
  value       = aws_vpc.my_vpc.id
}

output "vpc_tags" {
  description = "The tags of the VPC"
  value       = aws_vpc.my_vpc.tags
}
