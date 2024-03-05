
CREATE DATABASE golang;
GO;

CREATE table "ohlc_data"
(
    "symbol" text primary key,
    "timestamp" bigint,
    "high"   float NOT NULL,
    "low"    float NOT NULL,
    "open"   float NOT NULL ,
    "close"  float NOT NULL,
    "change" float NOT NULl

);
DROP TABLE ohlc_data ;

select * from ohlc_data;

CREATE TABLE ohlc_data (
                           symbol TEXT NOT NULL,
                           timestamp BIGINT NOT NULL,
                           open float NOT NULL,
                           high float NOT NULL,
                           low float NOT NULL,
                           close float NOT NULL,
                           change float NOT NULL,
                           PRIMARY KEY (symbol, timestamp)
);

truncate table ohlc_data
