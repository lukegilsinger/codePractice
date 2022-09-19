resource "aws_glue_job" "glue-job" {
  name = "${var.project}-glue-job"
  role_arn = aws_iam_role.iam_role_glue.arn
  max_retries = "0"
  command {
    script_location = "s3://${aws_s3_bucket.test-bucket.id}/scripts/glue-script.py"
    python_version = "3"
  }
}
