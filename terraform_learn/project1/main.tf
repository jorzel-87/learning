terraform {
  required_version = ">= 0.12"
}

provider "aws" {
  region = "eu-central-1"
}

# resource "provider_resource_type" "name" {
# some config options
# key = "value"
# key2 = "another_value"
# }

resource "aws_instance" "server_1" {
 ami = "ami-00a205cb8e06c3c4e"
 instance_type = "t2.micro"
 tags = {
   Name = "TF_server"
 }
}

resource "aws_vpc" "tf_vpc" {
  cidr_block = "10.10.0.0/16"
  tags = {
   Name = "TF_vpc"
 }
}

# example ofd refferencing resources
# resorce refferenced, does not need to be defined first!
# order does not matter
 
resource "aws_subnet" "tf_subnet_1" {
  vpc_id     = aws_vpc.tf_vpc.id
  cidr_block = "10.10.1.0/24"

  tags = {
    Name = "prod-subnet"
  }
}