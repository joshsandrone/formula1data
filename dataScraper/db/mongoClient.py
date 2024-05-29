import os
from dotenv import load_dotenv, find_dotenv
import pymongo
from pymongo import MongoClient
from bson import ObjectId
from classes import RaceWeekend, SeasonProfile
from tools import logger

# Load the project env vars from the ../.env
dotenv_path = find_dotenv('../.env')
load_dotenv(dotenv_path)

class MongoDB:

    def __init__(self) -> None:
        self.mongoUsername = os.getenv('MONGO_INITDB_ROOT_USERNAME')
        self.mongoPassword = os.getenv('MONGO_INITDB_ROOT_PASSWORD')
        self.mongoURI = f"mongodb://{self.mongoUsername}:{self.mongoPassword}@localhost:27017/"
        self.database_name = os.getenv('MONGO_DB_NAME')
        self.client = None
        self.establish_connection()

    def establish_connection(self):
        """Establish a connection to the MongoDB server."""
        try:
            self.client = MongoClient(self.mongoURI)
            self.db = self.client[self.database_name]
            logger.info(f"Connected to MongoDB database: {self.database_name}")
        except Exception as e:
            logger.error(f"Error connecting to database: {e}")

    def insert_race_weekend(self, raceWeekend : RaceWeekend, season : str):
        raceWeekendJson = raceWeekend.getJsonRepr()
        raceWeekendJson["season"] = season
        # May be worth adding a Season Object
        insertRaceResult = self.insert_document("races", raceWeekendJson)
        logger.info(f"Race added to DB")

    def insert_season_profile(self, seasonProfile : SeasonProfile):
        seasonProfileJson = seasonProfile.getJsonRepr()
        insertSeasonProfileResult = self.insert_document("seasons", seasonProfileJson)
        logger.info(f"Season Profile added to DB")

    def insert_document(self, collection_name : str, document : dict):
        """Insert a document into a specified collection."""
        collection = self.db[collection_name]
        result = collection.insert_one(document)
        return result

    def find_document(self, collection_name : str, query : dict):
        """Find a document in a specified collection."""
        collection = self.db[collection_name]
        result = collection.find_one(query)
        return result
    
    def update_document(self, collection_name : str, query : dict, updateOperation : dict):
        collection = self.db[collection_name]
        result = collection.update_one(query, updateOperation)
        return result

    def close_connection(self):
        """Close the connection to the MongoDB server."""
        if self.client:
            self.client.close()
            logger.info("Closed connection to MongoDB")

