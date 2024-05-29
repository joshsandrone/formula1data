# F1 Data Project

## Components:

### dataScraper
* Python with BeautifulSoup
* Scrapes data from formula1.com
### webapp1 (to rename) (API)
* Golang with Gin Gonic
* Process the F1 Data stored in the database
* Provides an API for the F1 Data
### mongo
* MongoDB database
* Stores the collected data from the dataScraper


## Running:
* The dataScraper is designed as standalone
* It has a DockerFile that will connect the scrapper to the exposed mongo port via localhost
* The other components run in tandem using docker-compose
    * `docker-compose up -d --build`
    * `docker-compose down`
