printf "Creating Services...\n\n"

curl --location --request POST 'http://localhost:8001/services' \
--form 'name="management-service"' \
--form 'protocol="http"' \
--form 'host="management-service"' \
--form 'port="8081"' \
--form 'path="/"'

curl --location --request POST 'http://localhost:8001/services' \
--form 'name="account-service"' \
--form 'protocol="http"' \
--form 'host="account-service"' \
--form 'port="8082"' \
--form 'path="/"'

curl --location --request POST 'http://localhost:8001/services' \
--form 'name="platform-service"' \
--form 'protocol="http"' \
--form 'host="platform-service"' \
--form 'port="8083"' \
--form 'path="/"'

printf "Creating Routes...\n\n"

curl --location --request POST 'http://localhost:8001/services/management-service/routes' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=manager' \
--data-urlencode 'paths=/mgr' \
--data-urlencode 'methods=GET' \
--data-urlencode 'methods=POST' \
--data-urlencode 'methods=PUT' \
--data-urlencode 'methods=OPTIONS' \
--data-urlencode 'methods=DELETE'

curl --location --request POST 'http://localhost:8001/services/account-service/routes' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=account' \
--data-urlencode 'paths=/auth' \
--data-urlencode 'methods=GET' \
--data-urlencode 'methods=POST' \
--data-urlencode 'methods=PUT' \
--data-urlencode 'methods=OPTIONS' \
--data-urlencode 'methods=DELETE'

curl --location --request POST 'http://localhost:8001/services/platform-service/routes' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=platform' \
--data-urlencode 'paths=/' \
--data-urlencode 'methods=GET' \
--data-urlencode 'methods=POST' \
--data-urlencode 'methods=PUT' \
--data-urlencode 'methods=OPTIONS' \
--data-urlencode 'methods=DELETE'

printf "Setting up authentication...\n\n"

curl --location --request POST 'http://localhost:8001/services/platform-service/plugins' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=jwt' \
--data-urlencode 'config.secret_is_base64=false' \
--data-urlencode 'config.run_on_preflight=true'

curl --location --request POST 'http://localhost:8001/services/management-service/plugins' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=jwt' \
--data-urlencode 'config.secret_is_base64=false' \
--data-urlencode 'config.run_on_preflight=true'

curl --location --request POST 'http://localhost:8001/plugins' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=cors' \
--data-urlencode 'config.origins=*' \
--data-urlencode 'config.methods=GET' \
--data-urlencode 'config.methods=POST' \
--data-urlencode 'config.methods=PUT' \
--data-urlencode 'config.methods=DELETE' \
--data-urlencode 'config.methods=OPTIONS' \
--data-urlencode 'config.headers=Content-Type' \
--data-urlencode 'config.headers=Access-Control-Allow-Origin' \
--data-urlencode 'config.exposed_headers=Access-Control-Allow-Origin' \
--data-urlencode 'config.credentials=true' \
--data-urlencode 'config.preflight_continue=true' \
--data-urlencode 'config.max_age=3600'