resource "aws_apigatewayv2_deployment" "api_gateway_deployment" {
  depends_on = [aws_apigatewayv2_route.auth_lambda_route]
  api_id     = aws_apigatewayv2_api.fastfood-api-gateway.id
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_apigatewayv2_stage" "default_stage" {
  api_id        = aws_apigatewayv2_api.fastfood-api-gateway.id
  name          = "$default"
  deployment_id = aws_apigatewayv2_deployment.api_gateway_deployment.id
}
