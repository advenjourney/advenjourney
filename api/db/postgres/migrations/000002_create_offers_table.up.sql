CREATE TABLE IF NOT EXISTS Offers(
    ID serial NOT NULL UNIQUE,
    Title VARCHAR (255) ,
    Location VARCHAR (255) ,
    Description VARCHAR (10000) ,
    TitleImageURL VARCHAR (1000) ,
    UserID INT ,
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,
    PRIMARY KEY (ID)
)
