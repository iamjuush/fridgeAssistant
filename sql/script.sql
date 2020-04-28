CREATE TABLE IF NOT EXISTS groceries
(
    Name     varchar(100) NOT NULL PRIMARY KEY ON CONFLICT REPLACE,
    Quantity float,
    Unit     varchar(20),
    Category varchar(50),
    Expiry   date
);
