# Command to run the docker container (use host network)
docker build -t scraper .
docker run --rm  --name scraper --network host scraper


docker run --rm  --network host scraper --profileSeason 2024

docker run --rm  --network host scraper --race italy --season 2024 --id 1235 --raceFormat standard

docker run --rm  --network host scraper --race miami --season 2024 --id 1234 --raceFormat sprint2024

docker run --rm  --network host scraper --race bahrain --season 2023 --id 1141 --raceFormat standard