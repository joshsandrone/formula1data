import json
from tools import logger
from .driverClass import Driver

class DriverSession:

    def __init__(self, driverFirstName : str, driverSurname : str, driverTeam : str, position : int, fastestLap : str) -> None:
        self.driver = Driver(driverFirstName, driverSurname, driverTeam)
        self.position = position
        self.fastestLap = fastestLap

    def getDriver(self) -> Driver:
        return self.driver   
    
    def getPosition(self) -> int:
        return self.position
    
    def getFastestLap(self) -> str:
        return self.fastestLap
    

class DriverQualfyingSession(DriverSession):

    def __init__(self, driverFirstName : str, driverSurname : str, driverTeam : str, laptimes : list, position : int) -> None:
        super().__init__(driverFirstName, driverSurname, driverTeam, position, None)
        self.q1Lap = laptimes[0]
        self.q2Lap = laptimes[1]
        self.q3Lap = laptimes[2]
    
    def getQ1Time(self) -> str:
        return self.q1Lap
    
    def getQ2Time(self) -> str:
        return self.q2Lap
    
    def getQ3Time(self) -> str:
        return self.q3Lap
    

class DriverPracticeSession(DriverSession):

    def __init__(self, driverFirstName : str, driverSurname : str, driverTeam : str, fastestLap : str, position : int) -> None:
        super().__init__(driverFirstName, driverSurname, driverTeam, position, fastestLap)


class DriverRaceSession(DriverSession):

    def __init__(self, driverFirstName : str, driverSurname : str, driverTeam : str, fastestLap : str, position : int, points : int, dnf : bool) -> None:
        super().__init__(driverFirstName, driverSurname, driverTeam, position, fastestLap)
        self.points = points
        self.dnf = dnf

    def setFastestLapData(self, fastestLap : dict) -> None:
        self.fastestLap = fastestLap

    def getPoints(self) -> int:
        return self.points
    
    def getDnf(self) -> bool:
        return self.dnf


class Session:

    def __init__(self, trackName) -> None:
        self.driverResults = []
        self.trackName = trackName

    def registerDriverResult(self, driverSession : DriverSession) -> None:
        self.driverResults.append(driverSession)

    def getDriverSessionObject(self, driverFirstName : str, driverSurname : str):
        for driverSession in self.driverResults:
            if driverSession.getDriver().getName() == f"{driverFirstName} {driverSurname}":
                return driverSession
            
        logger.warning(f"No driver session object for '{driverFirstName} {driverSurname}' in session")
        return None

    def __repr__(self):
        return str(self.getJsonRepr())


class QualifyingSession(Session):

    def getJsonRepr(self) -> str:
        reprJson = []
        for driverResult in self.driverResults:

            driver = driverResult.getDriver()
            driverResultJson = {
                "Position" : driverResult.getPosition(),
                "Driver" : driver.getName(),
                "Team" : driver.getTeam(),
                "Q1 Time" : driverResult.getQ1Time(),
                "Q2 Time" : driverResult.getQ2Time(),
                "Q3 Time" : driverResult.getQ3Time()
            }

            reprJson.append(driverResultJson)

        return reprJson

    

class PracticeSession(Session):

    def __init__(self, trackName, praticeNumber) -> None:
        super().__init__(trackName)
        self.practiceNumber = praticeNumber

    def getJsonRepr(self):
        reprJson = []
        for driverResult in self.driverResults:

            driver = driverResult.getDriver()
            driverResultJson = {
                "Position" : driverResult.getPosition(),
                "Driver" : driver.getName(),
                "Team" : driver.getTeam(),
                "FastestLap" : driverResult.getFastestLap()
            }

            reprJson.append(driverResultJson)

        return reprJson

    

class RaceSession(Session):

    def getJsonRepr(self):
        reprJson = []
        for driverResult in self.driverResults:

            driver = driverResult.getDriver()
            driverResultJson = {
                "Position" : driverResult.getPosition(),
                "Driver" : driver.getName(),
                "Team" : driver.getTeam(),
                "FastestLap": driverResult.getFastestLap(),
                "Points": driverResult.getPoints(),
                "Dnf": driverResult.getDnf()
            }

            reprJson.append(driverResultJson)    

        return reprJson   

    

class Weekend:

    def __init__(self,location : str, practice1 : PracticeSession, qualifying : QualifyingSession, race : RaceSession) -> None:
        self.location = location
        self.practice1 = practice1
        self.qualifying = qualifying
        self.race = race 


class RaceWeekend(Weekend):

    def __init__(self, location: str, practice1 : PracticeSession, practice2 : PracticeSession, practice3 : PracticeSession, qualifying : QualifyingSession, race : RaceSession) -> None:
        super().__init__(location, practice1,qualifying, race)
        self.practice2 = practice2
        self.practice3 = practice3

    def getJsonRepr(self):
        reprJson = {
            "location" : self.location,
            "Race" : self.race.getJsonRepr(),
            "Qualifying" : self.qualifying.getJsonRepr(),
            "Practice1" : self.practice1.getJsonRepr(),
            "Practice2" : self.practice2.getJsonRepr(),
            "Practice3" : self.practice3.getJsonRepr()
        }

        return reprJson
    

class RaceWeekendSprint2024(Weekend):

    def __init__(self, location: str, practice1 : PracticeSession, sprintQualifying : QualifyingSession, sprintRace : RaceSession, qualifying : QualifyingSession, race : RaceSession) -> None:
        super().__init__(location, practice1,qualifying, race)
        self.sprintQualifying = sprintQualifying
        self.sprintRace = sprintRace

    def getJsonRepr(self):
        reprJson = {
            "location" : self.location,
            "Race" : self.race.getJsonRepr(),
            "Qualifying" : self.qualifying.getJsonRepr(),
            "SprintRace" : self.sprintRace.getJsonRepr(),
            "SprintQualifying" : self.sprintQualifying.getJsonRepr(),
            "Practice1" : self.practice1.getJsonRepr()
        }

        return reprJson




class Season:

    def __init__(self, year : int):
        self.year = year
        self.races = []

    def addRace(self, race : Weekend):
        self.races.append(race)

    
    def getJsonRepr(self):
        reprJson = {f"{self.year}" : {
            "raceWeekends" : []
        }}

        for race in self.races:
            reprJson[f"{self.year}"]["raceWeekends"].append(race.getJsonRepr())

        return reprJson
    
    def getRaces(self):
        return self.races
    
    def exportJson(self):
        jsonOut = self.getJsonRepr()  # Get the dictionary representation
        json_string = json.dumps(jsonOut, indent=4)  # Convert dictionary to JSON string with indentation
        with open("exportedJson.json", "w") as json_file:
            json_file.write(json_string)


    def __repr__(self):
        return str(self.getJsonRepr())