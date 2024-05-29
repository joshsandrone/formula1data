import requests
from bs4 import BeautifulSoup
from tools import logger
    

def scrapeWebpage(url) -> BeautifulSoup:
    response = requests.get(url)
    if response.status_code == 200:
        soup = BeautifulSoup(response.text, 'html.parser')
        return soup
    else:
        logger.warning(f"Failed to retrieve content from {url}. Status code: {response.status_code}")
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

