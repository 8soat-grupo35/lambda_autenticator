# Lambda (substitua pelo ARN da sua lambda se já existir)
resource "aws_lambda_function" "auth_lambda" {
  filename         = data.archive_file.function_archive.output_path # Zip do código da Lambda
  function_name    = "lambda-authenticator"
  role             = data.aws_iam_role.labrole.arn
  handler          = "main"
  runtime          = "provided.al2023"
  source_code_hash = data.archive_file.function_archive.output_base64sha256
}
