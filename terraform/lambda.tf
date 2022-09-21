data "archive_file" "lambda_zip_file_int" {
  type        = "zip"
  output_path = "${var.lambda_src_path}/test2.zip"
  source_dir = "${var.lambda_src_path}/lambda-src"
}

resource "aws_lambda_function" "test_lambda" {
  # If the file is not in the current working directory you will need to include a 
  # path.module in the filename.
  filename      = "${var.lambda_src_path}/test2.zip"
  function_name = "${var.project}-test-function"
  role          = aws_iam_role.iam_role_lambda.arn
  handler       = "test.app"

  runtime = "nodejs12.x"
  timeout = 10

  vpc_config {
    subnet_ids         = [aws_subnet.subnet2.id]
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
    protocol  = -1
    self      = true
    from_port = 0
    to_port   = 0
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
