CREATE TABLE IF NOT EXISTS Users(
    ID serial NOT NULL,
    Username VARCHAR (127) NOT NULL UNIQUE,
    Password VARCHAR (127) NOT NULL,
    PRIMARY KEY (ID)
)