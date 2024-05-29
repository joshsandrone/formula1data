from .scrape import scrapeWebpage
from tools import logger
from classes import (
    DriverQualfyingSession, 
    QualifyingSession, 
    DriverPracticeSession, 
    PracticeSession, 
    DriverRaceSession, 
    RaceSession, 
    RaceWeekend, 
    RaceWeekendSprint2024, 
    Season
)


def scrapeSeason(seasonData : dict) -> Season:
    season = Season(seasonData["year"])
    for race in seasonData["races"]:
        raceWeekend = scrapeRaceWeekend(seasonData["year"], race['location'], race['f1id'], race['raceFormat'])
        season.addRace(raceWeekend)
    return season


def scrapeRaceWeekend(year : str, location : str, formula1ID : str, format : str) -> RaceWeekend:
    baseURL = f"https://www.formula1.com/en/results.html/{year}/races/{formula1ID}/{location}"

    if format == "standard":
        raceWeekend = scrapeStandardRaceWeekend(baseURL, location, year)
    elif format == "sprint2024":
        raceWeekend = scrapeSprintRaceWeekend_format2024(baseURL, location, year)
    elif format == "sprint2023":
        # Given the same number of practice, races and qualifyings, can use the 2024 func
        raceWeekend = scrapeSprintRaceWeekend_format2024(baseURL, location, year)
    else:
        logger.warning(f"Unable to scrape race weekend of type {format}")
        return None
    
    return raceWeekend


def scrapeStandardRaceWeekend(baseURL, raceLocation, raceYear) -> RaceWeekend:
    logger.info(f"SCRAPING: {raceLocation}, {raceYear}")

    race = scrapeRace(baseURL,raceLocation,raceYear,sprint=False)
    qualifying = scrapeQualfyingSession(baseURL,raceLocation,raceYear)
    practice1 = scrapePracticeSessopn(baseURL,raceLocation,raceYear,1)
    practice2 = scrapePracticeSessopn(baseURL,raceLocation,raceYear,2)
    practice3 = scrapePracticeSessopn(baseURL,raceLocation,raceYear,3)
    raceWeekend = RaceWeekend(raceLocation, practice1, practice2, practice3, qualifying, race)
    return raceWeekend


def scrapeSprintRaceWeekend_format2024(baseURL, raceLocation, raceYear):

    race = scrapeRace(baseURL, raceLocation, raceYear)
    qualifying = scrapeQualfyingSession(baseURL, raceLocation, raceYear)
    sprintRace = scrapeRace(baseURL, raceLocation, raceYear, sprint=True)
    sprintQualifying = scrapeQualfyingSession(baseURL, raceLocation, raceYear, sprint=True)
    practice1 = scrapePracticeSessopn(baseURL,raceLocation,raceYear,1)
    raceWeekend = RaceWeekendSprint2024(raceLocation, practice1, sprintQualifying, sprintRace, qualifying, race)
    return raceWeekend


def scrapeQualfyingSession(baseURL,raceLocation,raceYear,sprint=False) -> QualifyingSession:

    if sprint:
        logger.info(f"---> SCRAPING: Qualifying, {raceLocation}, {raceYear}")
        soup = scrapeWebpage(f"{baseURL}/sprint-qualifying.html")
    else:
        logger.info(f"---> SCRAPING: Sprint Qualifying, {raceLocation}, {raceYear}")
        soup = scrapeWebpage(f"{baseURL}/qualifying.html")

    if soup:
        qualfyingSession = QualifyingSession(raceLocation)
        resultsTable = soup.select(".resultsarchive-table")[0]
        driverPositions = resultsTable.select("tr")

        position = 1
        for pos in driverPositions:
            span_elements = pos.select("span")
            if span_elements:
                driverFirstName = span_elements[0].text
                driverSurname = span_elements[1].text
                driverTeam = pos.select("td")[4].text
                driverLapTimes = [lap_time.text for lap_time in pos.select("td")[5:8]]
                driverQualfyingSession = DriverQualfyingSession(driverFirstName, driverSurname, driverTeam, driverLapTimes, position)
                qualfyingSession.registerDriverResult(driverQualfyingSession)
                position += 1


        logger.debug(qualfyingSession.__repr__())
        return qualfyingSession
    return None
    

def scrapeRace(baseURL,raceLocation,raceYear,sprint =False) -> RaceSession:
    
    raceSession = RaceSession(raceLocation)

    if sprint:
        logger.info(f"---> SCRAPING: Sprint Race, {raceLocation}, {raceYear}")
        soup = scrapeWebpage(f"{baseURL}/sprint-results.html")
    else:
        logger.info(f"---> SCRAPING: Race, {raceLocation}, {raceYear}")
        soup = scrapeWebpage(f"{baseURL}/race-result.html")

    if soup:
        resultsTable = soup.select(".resultsarchive-table")[0]
        driverPositions = resultsTable.select("tr")

        position = 1
        for pos in driverPositions:
            span_elements = pos.select("span")
            if span_elements:
                driverFirstName = span_elements[0].text
                driverSurname = span_elements[1].text
                driverTeam = pos.select("td")[4].text

                driverDNF = False
                driverRaceDelta = pos.select("td")[6].text
                if (driverRaceDelta == "DNF"):
                    #print(f"DNF for {driverFirstName} {driverSurname} at {raceLocation}, {raceYear}")
                    driverDNF = True

                driverPoints = int(pos.select("td")[7].text)

                driverraceSession = DriverRaceSession(driverFirstName, driverSurname, driverTeam, None, position, driverPoints, driverDNF)
                raceSession.registerDriverResult(driverraceSession)
                position += 1
    else:
        return None
    
    if sprint:
        return raceSession
    
    #Lets know scrape the fastest lap data
    soup = scrapeWebpage(f"{baseURL}/fastest-laps.html")

    if soup:
        resultsTable = soup.select(".resultsarchive-table")[0]
        driverPositions = resultsTable.select("tr")

        position = 1
        for pos in driverPositions:
            span_elements = pos.select("span")
            if span_elements:
                driverFirstName = span_elements[0].text
                driverSurname = span_elements[1].text
                driverFastestLap = pos.select("td")[7].text
                driverFastestLap = {
                    "timeOfDay" : pos.select("td")[6].text,
                    "lapTime" : pos.select("td")[7].text,
                    "lapNumber" : pos.select("td")[5].text,
                    "avgSpeed": pos.select("td")[8].text
                }
                driverRaceSession = raceSession.getDriverSessionObject(driverFirstName, driverSurname)
                if driverRaceSession:
                    driverRaceSession.setFastestLapData(driverFastestLap)
                position += 1
    
        logger.debug(raceSession.__repr__())
        return raceSession
    return None

def scrapePracticeSessopn(baseURL,raceLocation,raceYear, practiceNumber) -> PracticeSession:
    logger.info(f"---> SCRAPING: Practice {practiceNumber}, {raceLocation}, {raceYear}")
    soup = scrapeWebpage(f"{baseURL}/practice-{practiceNumber}.html")

    if soup:
        practiceSession = PracticeSession(raceLocation, practiceNumber)
        resultsTable = soup.select(".resultsarchive-table")[0]
        driverPositions = resultsTable.select("tr")

        position = 1
        for pos in driverPositions:
            span_elements = pos.select("span")
            if span_elements:
                driverFirstName = span_elements[0].text
                driverSurname = span_elements[1].text
                driverTeam = pos.select("td")[4].text
                driverFastestLap = pos.select("td")[5].text
                driverpracticeSession = DriverPracticeSession(driverFirstName, driverSurname, driverTeam, driverFastestLap, position)
                practiceSession.registerDriverResult(driverpracticeSession)
                position += 1

        logger.debug(practiceSession.__repr__())
        return practiceSession
    
    return None


# scrapeSeason(year="2023",
#              raceLocations= [
#                  {"location" : "bahrain", "formula1.com_id" : 1141, "sprintWeekend" : False },
#                  {"location" : "saudi-arabia", "formula1.com_id" : 1142, "sprintWeekend" : False },
#                  {"location" : "australia", "formula1.com_id" : 1143, "sprintWeekend" : False },
#                  {"location" : "azerbaijan", "formula1.com_id" : 1207, "sprintWeekend" : True },
#                  {"location" : "miami", "formula1.com_id" : 1208, "sprintWeekend" : False },
#                  {"location" : "monaco", "formula1.com_id" : 1210, "sprintWeekend" : False },
#                  {"location" : "spain", "formula1.com_id" : 1211, "sprintWeekend" : False },
#                  {"location" : "canada", "formula1.com_id" : 1212, "sprintWeekend" : False },
#                  {"location" : "austria", "formula1.com_id" : 1213, "sprintWeekend" : True },
#                  {"location" : "great-britain", "formula1.com_id" : 1214, "sprintWeekend" : False },
#                  {"location" : "hungary", "formula1.com_id" : 1215, "sprintWeekend" : False },
#                  {"location" : "belgium", "formula1.com_id" : 1216, "sprintWeekend" : True },
#                  {"location" : "netherlands", "formula1.com_id" : 1217, "sprintWeekend" : False },
#                  {"location" : "italy", "formula1.com_id" : 1218, "sprintWeekend" : False},
#                  {"location" : "singapore", "formula1.com_id" : 1219, "sprintWeekend" : False},
#                  {"location" : "japan", "formula1.com_id" : 1220, "sprintWeekend" : False},
#                  {"location" : "qatar", "formula1.com_id" : 1221, "sprintWeekend" : True},
#                  {"location" : "united-states", "formula1.com_id" : 1222, "sprintWeekend" : True},
#                  {"location" : "mexico", "formula1.com_id" : 1223, "sprintWeekend" : False},
#                  {"location" : "brazil", "formula1.com_id" : 1224, "sprintWeekend" : True},
#                  {"location" : "las-vegas", "formula1.com_id" : 1225, "sprintWeekend" : False},
#                  {"location" : "abu-dhabi", "formula1.com_id" : 1226, "sprintWeekend" : False}
#              ])

