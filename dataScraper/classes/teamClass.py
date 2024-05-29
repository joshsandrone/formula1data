from .driverClass import Driver

class Team:

    def __init__(self, teamName : str, season : str) -> None:
        self.teamName = teamName
        self.season = season
        self.drivers = []
        self.points = 0

    def getTeamName(self) -> str:
        return self.teamName
    
    def addDriver(self, driver : Driver) -> None:
        self.drivers.append(driver)

    def addPoints(self, points) -> None:
        self.points += points
    
    def getJsonRepr(self):
        return {
            "name" : self.teamName,
            "drivers" : [driver.getJsonRepr() for driver in self.drivers],
            "points" : self.points,
            "primaryColor" : "#674ea7"  #Ideally scrape the teams primary color - default set to purple atm
        }