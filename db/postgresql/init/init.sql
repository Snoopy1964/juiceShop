CREATE SCHEMA juice_shop;
SET search_path TO juice_shop;

CREATE TABLE accounts (
    ID varchar(256) NOT NULL,
    email varchar(256) NOT NULL,
    passwd varchar(256) NOT NULL,
    PRIMARY KEY(ID)
);
