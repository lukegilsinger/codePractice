resource "aws_db_instance" "postgres_instance" {
  identifier = lower("${var.project}-postgres-db")
  db_name              = "postgres1"
  engine               = "postgres"
  engine_version       = "13.7"
  instance_class       = "db.t3.micro"
  allocated_storage    = 10
  username             = "postgres"
  password             = "postgres"
  # parameter_group_name = "default.mysql5.7"
  skip_final_snapshot  = true

  vpc_security_group_ids   = ["${aws_security_group.rds_security_group.id}"]
  # db_subnet_group_name     = "${var.rds_public_subnet_group}"
  port                     = 5432
  publicly_accessible      = true
  multi_az                 = false
  storage_encrypted        = true # you should always do this

  # allocated_storage        = 256 # gigabytes
  # backup_retention_period  = 7   # in days
  # engine                   = "postgres"
  # engine_version           = "9.5.4"
  # identifier               = "mydb1"
  # instance_class           = "db.r3.large"
  # name                     = "mydb1"
  # parameter_group_name     = "mydbparamgroup1" # if you have tuned it
  # password                 = "${trimspace(file("${path.module}/secrets/mydb1-password.txt"))}"
  # storage_type             = "gp2"
  # username                 = "mydb1"
}

resource "aws_security_group" "rds_security_group" {
  name = "${var.project}-rds-sg"

  description = "RDS (terraform-managed)"
  vpc_id      = aws_vpc.vpc.id

  # Only Postgres in
  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = [var.vpc_cidr_block]
  }

  # Allow all outbound traffic.
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

}