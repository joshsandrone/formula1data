import argparse
import json
from scraper import scrapeRaceWeekend, scrapeDriverTeamProfile, scrapeSeason
from classes import RaceWeekend
from db import MongoDB
from tools import logger

def parseArgs():
    parser = argparse.ArgumentParser()
    parser.add_argument('--race', type=str, help='Race location as specified on formula1.com')
    parser.add_argument('--season', type=str, help='Season, e.g. 2024')
    parser.add_argument('--id', type=int, help='The id from the URL at formula1.com')
    parser.add_argument('--raceFormat', type=str, choices=['standard', 'sprint2020', 'sprint2021', 'sprint2022', 'sprint2023', 'sprint2024'], help='Specify the race weekend format: standard, sprint2024, sprint2023, sprint2022, sprint2021, sprint2020')
    parser.add_argument('--profileSeason', type=str, help='Profile the drivers and teams from a season. This arg needs to be used alone.')
    parser.add_argument('--parseSeasonFromJson', type=str, help='Profile the season and parse each race. This arg needs to be used alone.')


    # Parse the arguments
    args = parser.parse_args()

    # profileSeason can only be used alone
    if args.profileSeason and (args.race or args.season or args.id or args.raceFormat or args.parseSeasonFromJson):
        parser.error("--profileSeason must be passed alone")

    # parseSeasonFromJson can only be used alone
    if args.parseSeasonFromJson and (args.race or args.season or args.id or args.raceFormat or args.profileSeason):
        parser.error("--parseSeasonFromJson must be passed alone")

    # Check if required arguments are provided unless using --profileSeason
    if not args.profileSeason and not args.parseSeasonFromJson and not (args.race and args.season and args.id and args.raceFormat):
        parser.error('--race, --season, --id, and --raceFormat are required unless using --profileSeason or --parseSeasonFromJson')

    return args


def addRaceToDB(mongoDB, season : str, raceLocation : str, id : str, raceFormat : str) -> None:
    logger.info(f"Scraping race: {raceLocation}, {season} with format {raceFormat}")

    raceWeekend = scrapeRaceWeekend(season, raceLocation, id, raceFormat)
    if not raceWeekend:
        logger.error("Error scraping the race weekend!")
        return

    mongoDB.insert_race_weekend(raceWeekend, season)


def addSeasonProfileToDB(mongoDB, season : str) -> None:
    logger.info(f"Profiling season: {season}")

    seasonProfile = scrapeDriverTeamProfile(season)
    if not seasonProfile:
        logger.error("Error scraping the season profile!")
        return
    
    mongoDB.insert_season_profile(seasonProfile)


def parseSeasonJsonandAddToDB(mongoDB, jsonFile):
    logger.info(f"Scraping season from: {jsonFile}")
    try:
        with open(jsonFile, 'r') as file:
            seasonData = json.load(file)
    except Exception as e:
        logger.error(f"Error parsing the json: {e}")
        return None

    season = scrapeSeason(seasonData)
    # Add a profile for the season
    addSeasonProfileToDB(mongoDB,seasonData["year"])
    # Add each race 
    for raceWeekend in season.getRaces():
        mongoDB.insert_race_weekend(raceWeekend, seasonData["year"])


if __name__ == "__main__":
    args = parseArgs()

    mongoDB = MongoDB() 

    if args.profileSeason:
        addSeasonProfileToDB(mongoDB, args.profileSeason)
    elif args.parseSeasonFromJson:
        parseSeasonJsonandAddToDB(mongoDB, args.parseSeasonFromJson)
    else:
        addRaceToDB(mongoDB, args.season,args.race,args.id, args.raceFormat)

    mongoDB.close_connection()
