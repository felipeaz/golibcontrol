cp .env_example .env \
&& docker network create -d bridge lib-net || true \
&& docker-compose \
-f docker/kong.sh \
-f docker/redis-yaml up \
-f docker/account.yaml up --build
