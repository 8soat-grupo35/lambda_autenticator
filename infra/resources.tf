resource "null_resource" "function_binary" {
    provisioner "local-exec" {
        command = "GOOS=linux GOARCH=amd64 go build -o bootstrap ../main.go"
    }
}