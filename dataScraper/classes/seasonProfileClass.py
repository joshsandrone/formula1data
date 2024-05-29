from typing import List
from .teamClass import Team
from .driverClass import Driver

class SeasonProfile:

    def __init__(self, season : str, teams : List[Team], drivers: List[Driver]):
        self.teams = teams
        self.drivers = drivers
        self.season = season

    def getJsonRepr(self):
        return {
            "season" : self.season,
            "teams" : [team.getJsonRepr() for team in self.teams]
        }