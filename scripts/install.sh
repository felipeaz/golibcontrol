cp .env_example .env && docker-compose \
-f docker/conf.yaml \
-f docker/account.yaml \
-f docker/management.yaml \
-f docker/platform.yaml \
up --build