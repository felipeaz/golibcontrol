cp .env_example .env && docker-compose \
-f docker/redis.yaml \
-f docker/kong.yaml \
-f docker/kafka.yaml \
-f docker/account.yaml \
-f docker/management.yaml \
-f docker/platform.yaml \
up --build