data "archive_file" "lambda_zip_file_int" {
  type        = "zip"
  output_path = "${var.lambda_src_path}/hello-world.zip"
  source {
    content  = file("${var.lambda_src_path}/hello-world.js")
    filename = "hello-world.js"
  }
}

resource "aws_lambda_function" "test_lambda" {
  # If the file is not in the current working directory you will need to include a 
  # path.module in the filename.
  filename      = "${var.lambda_src_path}/hello-world.zip"
  function_name = "${var.project}-test-function"
  role          = aws_iam_role.iam_role_lambda.arn
  handler       = "hello-world.app"

  runtime = "nodejs12.x"

  vpc_config {
    subnet_ids         = [aws_subnet.subnet1.id]
    security_group_ids = [aws_security_group.lambda-sg.id]
  }

  environment {
    variables = {
      foo = "bar"
    }
  }
}

# SECURITY GROUPS #
resource "aws_security_group" "lambda-sg" {
  name   = "${var.project}-lambda_sg"
  vpc_id = aws_vpc.du_vpc.id

  # HTTP access from anywhere
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = [var.vpc_cidr_block] # allow traffic from in vpc
  }

  # outbound internet access
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.project}-lambda-security-group"
  }
}