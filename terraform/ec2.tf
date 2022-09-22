data "aws_ami" "amazon-linux-2" {
 most_recent = true

 filter {
   name   = "owner-alias"
   values = ["amazon"]
 }

 filter {
   name   = "name"
   values = ["amzn2-ami-hvm*"]
 }
}

resource "aws_network_interface" "nw_interface" {
  subnet_id   = aws_subnet.subnet1.id
  private_ips = ["10.0.0.100"]

  tags = {
    Name = "primary_network_interface"
  }
}

resource "aws_instance" "vpc_bastion" {
  ami           = data.aws_ami.amazon-linux-2.id
  instance_type = "t2.micro"

  network_interface {
    network_interface_id = aws_network_interface.nw_interface.id
    device_index         = 0
  }

  tags = {
    Name = "${var.project}-ec2-bastion"
  }
}

resource "aws_security_group" "bastion_host" {
  name        = "bastion_host_SG"
  description = "Allow SSH"
  vpc_id      =  aws_vpc.du_vpc.id                

  ingress {
    description = "SSH from VPC"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "${var.project}-bastion-security-group"
  }
}