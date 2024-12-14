resource "aws_apigatewayv2_integration" "lambda_auth" {
    api_id = aws_apigatewayv2_api.fastfood-api-gateway.id
    integration_type = "AWS_PROXY"
    integration_uri = aws_lambda_function.auth_lambda.invoke_arn
}