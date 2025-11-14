# rest_service_go
Link-shortener REST service on golang.

# Start service:
$env:CONFIG_PATH="C:\deals\rest_service_go\config\local.yaml" 
go run .\cmd\url-shortener\main.go

# TEST-CASES
curl -X POST http://localhost:8081/url -H "Content-Type: application/json" -u "artem:123" -d "{\"url\": \"https://google2.com\", \"alias\": \"test_alias2\"}"
{"status":"OK","alias":"test_alias2"}

http://localhost:8081/test_alias
<a href="https://google.com">Found</a>.

# 
ssh -i rest_service_go  -o StrictHostKeyChecking=no root@178.72.149.152