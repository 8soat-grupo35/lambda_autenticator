provider "aws" {
  region = "us-east-1" # Altere para a região desejada
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda_execution_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "go_lambda" {
  function_name = "go_lambda_function"
  role          = aws_iam_role.lambda_role.arn
  handler       = "main"
  runtime       = "go1.x"

  filename         = "./function.zip" # Caminho do arquivo zipado com o código
  source_code_hash = filebase64sha256("./function.zip")

  environment {
    variables = {
      ENV_VAR = "example_value" # Substitua com suas variáveis de ambiente
    }
  }

  tags = {
    Environment = "dev"
    Project     = "example"
  }
}

resource "aws_lambda_permission" "allow_api_gateway" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.go_lambda.function_name
  principal     = "apigateway.amazonaws.com"
}

resource "aws_api_gateway_rest_api" "lambda_api" {
  name        = "go-lambda-api"
  description = "API Gateway for the Go Lambda Function"
}

resource "aws_api_gateway_resource" "lambda_resource" {
  rest_api_id = aws_api_gateway_rest_api.lambda_api.id
  parent_id   = aws_api_gateway_rest_api.lambda_api.root_resource_id
  path_part   = "example"
}

resource "aws_api_gateway_method" "lambda_method" {
  rest_api_id   = aws_api_gateway_rest_api.lambda_api.id
  resource_id   = aws_api_gateway_resource.lambda_resource.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "lambda_integration" {
  rest_api_id = aws_api_gateway_rest_api.lambda_api.id
  resource_id = aws_api_gateway_resource.lambda_resource.id
  http_method = aws_api_gateway_method.lambda_method.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.go_lambda.invoke_arn
}

resource "aws_api_gateway_deployment" "lambda_deployment" {
  rest_api_id = aws_api_gateway_rest_api.lambda_api.id
  depends_on = [
    aws_api_gateway_integration.lambda_integration
  ]
  stage_name = "dev"
}

output "api_url" {
  value = aws_api_gateway_deployment.lambda_deployment.invoke_url
}
