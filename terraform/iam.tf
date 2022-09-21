data "aws_iam_policy_document" "AWSLambdaTrustPolicy" {
  version = "2012-10-17"
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "iam_role_lambda" {
  assume_role_policy = data.aws_iam_policy_document.AWSLambdaTrustPolicy.json
  name               = "${var.project}-iam-role-lambda"
}

resource "aws_iam_role_policy_attachment" "iam_role_policy_attachment_lambda_basic_execution" {
  role       = aws_iam_role.iam_role_lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "iam_role_policy_attachment_lambda_vpc_access_execution" {
  role       = aws_iam_role.iam_role_lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

data "aws_iam_policy_document" "s3" {
  statement {
    effect = "Allow"
    resources = ["*"]

    actions = [
      "s3:*",
    ]
  }
}

resource "aws_iam_policy" "s3" {
  policy = data.aws_iam_policy_document.s3.json
}

resource "aws_iam_role_policy_attachment" "s3" {
  role = aws_iam_role.iam_role_lambda.name
  policy_arn = aws_iam_policy.s3.arn
}

data "aws_iam_policy_document" "AWSGlueTrustPolicy" {
  version = "2012-10-17"
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["glue.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "iam_role_glue" {
  assume_role_policy = data.aws_iam_policy_document.AWSGlueTrustPolicy.json
  name               = "${var.project}-iam-role-glue"
}