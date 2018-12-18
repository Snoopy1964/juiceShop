-- CREATE SCHEMA juice_shop;
-- SET search_path TO juice_shop;

CREATE TABLE users (
    ID int NOT NULL,
    email varchar(255) NOT NULL,
    passwd varchar(255) NOT NULL,
    firstname varchar(255),
    lastname varchar(255),
    lastlogin date,
    PRIMARY KEY(ID)
);

INSERT INTO users VALUES(1, 'ralf@ehret-family.com', '-mwjD-E-k_wrJfujB2FIDPf2GU0CFr6KX3Yf4EGFqipadBdDS8ztMWg9TWQPPmOelOVGQDXQ3D9sGW_NSa26Bg==', 'Ralf', 'Ehret', '2018-01-01');
INSERT INTO users VALUES(2, 'conni@ehret-family.com', 'RG-AmFFQViLKcclQzELlvJ0mbk_dAf4rENeJMevA1WslVW5bMVuzBKQdRVL05Y_06VuDR227zudxhyMF3TBzXQ==', 'Conni', 'Ehret', '2018-01-01');

