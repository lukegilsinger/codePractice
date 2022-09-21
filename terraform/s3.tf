resource "aws_s3_bucket" "test-bucket" {
  bucket = lower("${var.project}-test-bucket")

  tags = local.common_tags
}

resource "aws_s3_object" "glue-script" {
  bucket = aws_s3_bucket.test-bucket.id
  key = "scripts/glue-script.py"
  source = "${var.glue_src_path}/hello-world.py"
}