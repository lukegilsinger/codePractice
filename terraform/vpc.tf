# Declare the az data source
data "aws_availability_zones" "available" {
  state = "available"
}

resource "aws_vpc" "du-vpc" {
  cidr_block           = var.vpc_cidr_block
  enable_dns_hostnames = var.enable_dns_hostnames

  tags = {
    Name = "${var.project}-vpc"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.du_vpc.id #referencing the above vpc

  tags = {
    Name = "${var.project}-internet-gateway"
  }
}

resource "aws_subnet" "subnet1" {
  cidr_block              = var.vpc_subnets_cidr_block[0]
  vpc_id                  = aws_vpc.du_vpc.id
  map_public_ip_on_launch = var.map_public_ip_on_launch
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
    Name = "${var.project}-subnet-public-1"
  }
}

resource "aws_subnet" "subnet2" {
  cidr_block              = var.vpc_subnets_cidr_block[1]
  vpc_id                  = aws_vpc.du_vpc.id
  map_public_ip_on_launch = var.map_public_ip_on_launch
  availability_zone = data.aws_availability_zones.available.names[1]

  tags = {
    Name = "${var.project}-subnet-public-2"
  }
}

# ROUTING #
resource "aws_route_table" "rtb" {
  vpc_id = aws_vpc.du_vpc.id

  route { # outward traffic
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "${var.project}-route-table-public"
  }
}

resource "aws_route_table_association" "rta-subnet1" {
  subnet_id      = aws_subnet.subnet1.id
  route_table_id = aws_route_table.rtb.id
}

resource "aws_route_table_association" "rta-subnet2" {
  subnet_id      = aws_subnet.subnet2.id
  route_table_id = aws_route_table.rtb.id
}
