cp .env_example .env \
&& docker network create -d bridge lib-net || true \
&& docker-compose \
-f docker/platform.yaml up --build