# infra-terra-test-poc
terratest POC

This POC does the following tasks,

1. Creates a VPC.
2. Sets up an Internet Gateway for public subnets.
3. Defines a public route table that routes traffic to the Internet Gateway.
4. Creates two public subnets and associates them with the public route table.
5. Creates two private subnets.
6. Defines a security group to allow SSH access.


