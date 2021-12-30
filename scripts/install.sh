cp .env_example .env && docker-compose \
-f build/docker/services/conf/docker-compose.yaml \
-f build/docker/services/account/docker-compose.yaml \
-f build/docker/services/management/docker-compose.yaml \
-f build/docker/services/platform/docker-compose.yaml \
up --build