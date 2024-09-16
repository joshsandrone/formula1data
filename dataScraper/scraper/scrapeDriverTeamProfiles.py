from .scrape import scrapeWebpage
from tools import logger
from classes import (
    Driver,
    Team,
    SeasonProfile,
    SeasonRace
)



def scrapeDriverTeamProfile(season : str) -> SeasonProfile:
    logger.info(f"-> SCRAPING: Driver and Team Profiles from Season, {season}")

    # First scrape the drivers championship from the season
    # This will tell us the drivers, teams and points of the team
    # Using the drivers championship for team points prevents any
    # team point deductions harming the data.
    seasonProfile = extractDriversAndTeam(season)

    if not seasonProfile:
        return None

    # Next add the races in the season to the seasonProfile
    # These will be scraped from the season results page
    extractRaces(seasonProfile, season)
    return seasonProfile


def extractDriversAndTeam(season : str) -> SeasonProfile:
    logger.info(f"---> Extracting Drivers and Teams from season: {season}")
    url = f"https://www.formula1.com/en/results.html/{season}/drivers.html"
    soup = scrapeWebpage(url)

    teams = []
    drivers = []

    if soup:

        # Scraping the relationship between the teams and drivers
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
                    logger.debug(f"Identified new team: {driverTeam}")
                    teamObj = Team(driverTeam, season)
                    teams.append(teamObj)

                # Add the driver to the team object
                teamObj.addDriver(driver)
                teamObj.addPoints(driverPoints)
                position += 1

        
        return SeasonProfile(season, teams, drivers)
    
    logger.error("Error: Beautiful Soup failed!")
    return None

def extractRaces(seasonProfile: SeasonProfile, season : str):
    logger.info(f"---> Extracting Races from season: {season}")
    url = f"https://www.formula1.com/en/results.html/{season}/races.html"
    soup = scrapeWebpage(url)

    if soup:
        table = soup.select(".resultsarchive-table")[0]
        races = table.select("tr")

        round = 1
        for race in races:
            span_elements = race.select("span")
            if span_elements:
                raceLocation = race.select("td")[1].text.strip()
                raceDate = race.select("td")[2].text
                raceLaps = race.select("td")[5].text
                logger.debug(f"{raceLocation = } {raceDate = } {raceLaps = } {round = }")
                
                seasonRace = SeasonRace(round, raceLocation, raceDate, int(raceLaps))
                seasonProfile.addRace(seasonRace)

                round += 1

    else:
        logger.error("Error: Beautiful Soup failed!")
