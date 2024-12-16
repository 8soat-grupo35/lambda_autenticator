resource "aws_apigatewayv2_api" "fastfood-api-gateway" {
    name = "Api gateway fastfood"
    protocol_type = "HTTP"
}

output "apigateway_endpoint" {
  value = aws_apigatewayv2_api.fastfood-api-gateway.api_endpoint
}