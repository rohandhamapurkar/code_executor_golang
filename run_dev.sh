docker system prune -f
docker-compose -f docker-compose.dev.yml build
docker system prune -f
docker-compose -f docker-compose.dev.yml up -d