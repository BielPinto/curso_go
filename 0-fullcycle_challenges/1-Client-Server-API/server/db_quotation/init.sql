

USE goexpert ;
CREATE TABLE IF NOT EXISTS  quotation (
  id varchar(255),
  code varchar(80),
  codein varchar(80),
  name varchar(80),
  high  decimal(10,4),
  low decimal(10,4),
  varBid decimal(10,4),
  PctChange decimal(10,2),
  Bid decimal(10,4),
  Ask decimal(10,4),
  CreateDate varchar(80),
  primary key (id)
);
-- ALTER DATABASE device_commander OWNER TO device_commander;



-- CREATE TABLE IF NOT EXISTS products ( id INTEGER PRIMARY KEY AUTOINCREMENT, code varchar(80), codein varchar(80), name varchar(80), high  decimal(10,4), low decimal(10,4), varBid decimal(10,4), PctChange decimal(10,2), Bid decimal(10,4), Ask decimal(10,4), CreateDat varchar(80));
   
