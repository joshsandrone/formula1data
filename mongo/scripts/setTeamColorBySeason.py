import argparse
from connectToMongo import MongoDB


def parseArgs():
    parser = argparse.ArgumentParser()
    parser.add_argument('--season', type=str, help='Season, e.g. 2024', required=True)
    # Parse the arguments
    args = parser.parse_args()

    return args

def getTeamColors(db, season : str):
    season = db.find_document("seasons", {"season" : season} ,{})
    return season["teams"]

def updateTeamColors(db, teams : dict):
    for team in teams:
        print(team)
        updatedColor = "empty"
        updatedColor = input(f"{team['name']} : current color is: {team['primaryColor']} . Enter updated color in hex (N to not update): ")

        if updatedColor[0] == '#':
            updateTeamColor(db, team['name'], updatedColor)

def updateTeamColor(db, teamname, color):
    db.update_document("seasons", {"teams.name" : teamname}, {"$set" : {"teams.$.primaryColor" : color}})


if __name__ == "__main__":
    args = parseArgs()
    db = MongoDB()
    teams = getTeamColors(db, args.season)
    updateTeamColors(db, teams)

