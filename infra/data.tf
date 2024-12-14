data "aws_vpc" "vpc" {
  cidr_block = var.vpcCidr
}

data "aws_subnets" "subnets" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.vpc.id]
  }
  filter {
    name = "availability-zone"
    values = ["${var.regionDefault}a","${var.regionDefault}b"]
  }
}

data "aws_subnet" "subnet" {
  for_each = toset(data.aws_subnets.subnets.ids)
  id       = each.value
}

data "aws_iam_role" "labrole" {
  name = "LabRole"
}

data "archive_file" "function_archive" {
  depends_on = [null_resource.function_binary]

  type        = "zip"
  source_file = "bootstrap"
  output_path = "bootstrap.zip"
}