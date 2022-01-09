cp .env_example .env && docker network create -d bridge lib-net && docker-compose \
-f docker/kong.yaml \
-f docker/redis.yaml \
-f docker/kafka.yaml \
-f docker/account.yaml \
-f docker/management.yaml \
-f docker/platform.yaml \
up --build