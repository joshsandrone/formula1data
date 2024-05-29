from .scrape import scrapeWebpage
from tools import logger
from classes import (
    Driver,
    Team,
    SeasonProfile
)



def scrapeDriverTeamProfile(season : str) -> SeasonProfile:
    logger.info(f"---> SCRAPING: Driver and Team Profiles from Season, {season}")

    url = f"https://www.formula1.com/en/results.html/{season}/drivers.html"
    soup = scrapeWebpage(url)

    teams = []
    drivers = []

    if soup:

        # Scraping the realtionship between the teams and drivers
        # by using the drivers championship for the season
        table = soup.select(".resultsarchive-table")[0]
        driverPositions = table.select("tr")

        position = 1
        for pos in driverPositions:
            span_elements = pos.select("span")
            if span_elements:
                driverFirstName = span_elements[0].text
                driverSurname = span_elements[1].text
                driverTeam = pos.select("td")[4].text.replace("\n", "")
                driverPoints = int(pos.select("td")[5].text)
                logger.debug(f"{driverFirstName = } {driverSurname = } {driverPoints = }")

                # Set team to None - the link will be set in TeamObj instead
                driver = Driver(driverFirstName, driverSurname, None)
                drivers.append(driver)

                # Find the team object in the list or create a new one if it doesn't exist
                teamObj = next((team for team in teams if team.getTeamName() == driverTeam), None)
                if not teamObj:
                    logger.info(f"Identified new team: {driverTeam}")
                    teamObj = Team(driverTeam, season)
                    teams.append(teamObj)

                # Add the driver to the team object
                teamObj.addDriver(driver)
                teamObj.addPoints(driverPoints)
                position += 1

        
        return SeasonProfile(season, teams, drivers)


    logger.error("Error: Beautiful Soup failed!")
    return None