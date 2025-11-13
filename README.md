# rest_service_go
Link-shortener REST service on golang.

# TEST-CASES
curl -X POST http://localhost:8081/url -H "Content-Type: application/json" -u "artem:123" -d "{\"url\": \"https://google2.com\", \"alias\": \"test_alias2\"}"
{"status":"OK","alias":"test_alias2"}

http://localhost:8081/test_alias
<a href="https://google.com">Found</a>.