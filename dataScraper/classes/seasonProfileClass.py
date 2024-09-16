from typing import List
from .teamClass import Team
from .driverClass import Driver

class SeasonProfile:

    def __init__(self, season : str, teams : List[Team], drivers: List[Driver]):
        self.teams = teams
        self.drivers = drivers
        self.season = season
        self.races = []

    def addRace(self, race):
        self.races.append(race)

    def getJsonRepr(self):
        return {
            "season" : self.season,
            "teams" : [team.getJsonRepr() for team in self.teams],
            "races" : [race.getJsonRepr() for race in self.races]
        }

class SeasonRace:

    def __init__(self, round : int, location : str, date : str, laps : int):
        self.round = round
        self.location = location
        self.date = date
        self.laps = laps

    def getJsonRepr(self):
        return {
            "round" : self.round,
            "location" : self.location,
            "date" : self.date,
            "laps" : self.laps
        }