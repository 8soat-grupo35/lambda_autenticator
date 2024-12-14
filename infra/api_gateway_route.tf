resource "aws_apigatewayv2_route" "auth_lambda_route" {
    api_id = aws_apigatewayv2_api.fastfood-api-gateway.id
    route_key = "ANY /Lambda-Autenticator"
    target = "integrations/${aws_apigatewayv2_integration.lambda_auth.id}"
}