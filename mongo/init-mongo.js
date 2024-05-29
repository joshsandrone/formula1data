/* 

*/

// create the database
db = db.getSiblingDB("f1data")

// collection 1 - Driver Profiles - Name, Team, Birthdate, Nationality
db.createCollection("driverProfiles")

// collection 2 = Team Profiles - Name, Primary Color, Entry per Year, Drivers
db.createCollection("teamProfiles")

// collection 3 - Season's
db.createCollection("seasons")

// collection 4 - Race's
db.createCollection("races")


