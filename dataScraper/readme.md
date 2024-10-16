# Command to run the docker container (use host network)

The below instructions give showcase how to use the datascraper using docker.

### To scrape an individual race result.
1. Find the dedicated link on F1.com (e.g. For Miami 2024: https://www.formula1.com/en/results/2024/races/1234/miami/race-result)
2. the format of the link above is: https://www.formula1.com/en/results/{year}/races/{id}/{race}/race-result
3. Unfortunealty, not all ids are sequential race to race, I have prepared Json files for the 2023 season (please contact).
4. Please provide the race format (sprint2024, sprint2023, standard)

### Season Profiles
1. Scraping a season profile creates adds to the season collection in the MongoDB
2. It scrapes Drivers, Teams
3. e.g. docker run --rm  --network host scraper --profileSeason 2024

### Example commands
docker build -t scraper .
docker run --rm  --name scraper --network host scraper
docker run --rm  --network host scraper --profileSeason 2024
docker run --rm  --network host scraper --race italy --season 2024 --id 1235 --raceFormat standard
docker run --rm  --network host scraper --race miami --season 2024 --id 1234 --raceFormat sprint2024
docker run --rm  --network host scraper --race bahrain --season 2023 --id 1141 --raceFormat standard
