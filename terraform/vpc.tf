# Declare the az data source
data "aws_availability_zones" "available" {
  state = "available"
}

resource "aws_vpc" "du_vpc" {
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
  cidr_block              = var.subnet_public_cidr_block
  vpc_id                  = aws_vpc.du_vpc.id
  map_public_ip_on_launch = true
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
    Name = "${var.project}-subnet-public"
  }
}

resource "aws_subnet" "subnet2" {
  cidr_block              = var.subnet_private_cidr_block
  vpc_id                  = aws_vpc.du_vpc.id
  map_public_ip_on_launch = false
  availability_zone = data.aws_availability_zones.available.names[1]

  tags = {
    Name = "${var.project}-subnet-private"
  }
}

# ROUTING #
resource "aws_route_table" "rt_public" {
  vpc_id = aws_vpc.du_vpc.id

  route { # outward traffic
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "${var.project}-route-table-public"
  }
}

resource "aws_route_table" "rt_private" {
  vpc_id = aws_vpc.du_vpc.id

  route { # outward traffic
    cidr_block = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat_gateway.id
  }

  tags = {
    Name = "${var.project}-route-table-private"
  }
}

resource "aws_route_table_association" "rta-subnet1" {
  subnet_id      = aws_subnet.subnet1.id
  route_table_id = aws_route_table.rt_public.id
}

resource "aws_route_table_association" "rta-subnet2" {
  subnet_id      = aws_subnet.subnet2.id
  route_table_id = aws_route_table.rt_private.id
}

resource "aws_eip" "eip" {
  vpc        = true
  depends_on = [aws_internet_gateway.igw]
  tags = {
    Name = "${var.project}-eip"
  }
}

resource "aws_nat_gateway" "nat_gateway" {
  allocation_id = aws_eip.eip.id
  subnet_id     = aws_subnet.subnet1.id
  tags = {
    Name = "${var.project}-nat-gateway"
  }
}

resource "aws_default_network_acl" "default_network_acl" {
  default_network_acl_id = aws_vpc.du_vpc.default_network_acl_id
  subnet_ids             = [aws_subnet.subnet1.id, aws_subnet.subnet2.id]

  ingress {
    protocol   = -1
    rule_no    = 100
    action     = "allow"
    cidr_block = "0.0.0.0/0"
    from_port  = 0
    to_port    = 0
  }

  egress {
    protocol   = -1
    rule_no    = 100
    action     = "allow"
    cidr_block = "0.0.0.0/0"
    from_port  = 0
    to_port    = 0
  }

  tags = {
    Name = "${var.project}-default-network-acl"
  }
}

resource "aws_default_security_group" "default_security_group" {
  vpc_id = aws_vpc.du_vpc.id

  ingress {
    protocol  = -1
    self      = true
    from_port = 0
    to_port   = 0
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    # cidr_blocks = ["127.0.0.1/32"]
  }

  tags = {
    Name = "${var.project}-default-security-group"
  }
}