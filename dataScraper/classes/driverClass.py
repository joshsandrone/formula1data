from typing import List, Union
from tools import logger

class Driver:

    # To keep backwards compatability, team can be passed as object or str.
    def __init__(self, firstName : str, surName : str, team : Union[None, str]) -> None:
        self.name = f"{firstName} {surName}"
        self.team = team
    
    def getName(self) -> str:
        return self.name
    
    def getTeam(self) -> Union[None, str]:
        return self.team
    
    def getJsonRepr(self) -> None:
        return {
            "name" : self.name
        }


