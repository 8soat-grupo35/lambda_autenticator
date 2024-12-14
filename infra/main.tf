provider "aws" {
  region = "us-east-1" # Altere para sua região
}

/*
# API Gateway HTTP API
resource "aws_apigatewayv2_api" "http_api" {
  name          = "lambda-auth-api"
  protocol_type = "HTTP"
}

# Rota do API Gateway para a Lambda
resource "aws_apigatewayv2_integration" "lambda_integration" {
  api_id             = aws_apigatewayv2_api.http_api.id
  integration_type   = "AWS_PROXY"
  integration_uri    = aws_lambda_function.auth_lambda.invoke_arn
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "default_route" {
  api_id    = aws_apigatewayv2_api.http_api.id
  route_key = "ANY /auth" # Substitua por seu endpoint

  target = "integrations/${aws_apigatewayv2_integration.lambda_integration.id}"
}

# Deploy da API
resource "aws_apigatewayv2_stage" "default_stage" {
  api_id      = aws_apigatewayv2_api.http_api.id
  name        = "$default"
  auto_deploy = true
}

# Permitir que o API Gateway invoque a Lambda
resource "aws_lambda_permission" "allow_apigateway" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.auth_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.http_api.execution_arn}/*"
}

# Output do endpoint da API
output "api_endpoint" {
  value = aws_apigatewayv2_api.http_api.api_endpoint
}
*/