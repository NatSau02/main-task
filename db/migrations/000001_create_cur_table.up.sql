CREATE TABLE IF NOT EXISTS curren(
   id serial PRIMARY KEY,
   currency1 VARCHAR (100) NOT NULL,
   currency2 VARCHAR (100) NOT NULL,
   rate VARCHAR (100),
   date VARCHAR (100) 
);