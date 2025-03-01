CREATE TABLE "GreenObject" (
  "Id" uuid UNIQUE PRIMARY KEY,
  "LocationName" varchar(50) NOT NULL,
  "Location" uuid,
  "Shape" varchar(30),
  "TrashType" varchar(15),
  "GreenScore" uuid,
  "Disabled" boolean
);

CREATE TABLE "GreenScore" (
  "Id" uuid UNIQUE PRIMARY KEY,
  "Verification" integer,
  "Report" integer,
  "TrashRank" varchar(15)
);

CREATE TABLE "Location" (
  "Id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "Latitude" float8,
  "Longitude" float8,
  "Street" varchar(100),
  "City" varchar(100),
  "Country" varchar(50)
);

CREATE TABLE "UserRank" (
  "Id" uuid UNIQUE PRIMARY KEY,
  "UserPoints" integer,
  "UserRank" varchar(15),
  "Premium" varchar(15),
  "PremiumLength" date
);

CREATE TABLE "User" (
  "Id" uuid UNIQUE PRIMARY KEY,
  "Username" varchar(50) UNIQUE NOT NULL,
  "Mail" varchar(50) UNIQUE NOT NULL,
  "Password" varchar NOT NULL,
  "Name" varchar(50) NOT NULL,
  "Surname" varchar(50) NOT NULL,
  "Picture" bytea,
  "City" varchar(50) NOT NULL,
  "Country" varchar(50) NOT NULL,
  "Sex" varchar(15),
  "Birthday" date,
  "Disabled" boolean,
  "UserRank" uuid,
  "Role" varchar(15)
);

CREATE TABLE "UserRelationship" (
  "User1" uuid,
  "User2" uuid,
  "Date" date,
  "Type" varchar(15)
);

CREATE TABLE "TOU" (
  "TOId" uuid UNIQUE,
  "UId" uuid UNIQUE
);

ALTER TABLE "TOU" ADD FOREIGN KEY ("TOId") REFERENCES "GreenObject" ("Id");

ALTER TABLE "GreenObject" ADD FOREIGN KEY ("Location") REFERENCES "Location" ("Id");

ALTER TABLE "GreenObject" ADD FOREIGN KEY ("GreenScore") REFERENCES "GreenScore" ("Id");

ALTER TABLE "TOU" ADD FOREIGN KEY ("UId") REFERENCES "User" ("Id");

ALTER TABLE "User" ADD FOREIGN KEY ("UserRank") REFERENCES "UserRank" ("Id");

ALTER TABLE "UserRelationship" ADD FOREIGN KEY ("User1") REFERENCES "User" ("Id");

ALTER TABLE "UserRelationship" ADD FOREIGN KEY ("User2") REFERENCES "User" ("Id");
