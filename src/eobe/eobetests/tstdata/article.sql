# 1. Clean up
DROP TABLE APILIST_TBL;
DROP TABLE APIPARAM_TBL;
DROP TABLE APIRESLT_TBL;

# 2. Create Table
 CREATE TABLE APILIST_TBL (
	id 			int AUTO_INCREMENT not null primary key,
    apiname		char(32) not null UNIQUE,
    apidesc		TEXT
 );
 
 CREATE TABLE APIPARAM_TBL(
	id 			int AUTO_INCREMENT not null primary key,
	paramname	char(32) not null,
    apiname 	char(32) not null,
    allwsrc		TINYINT not null,	/* allowed data source, 0xE(14) for variables only, 0x1 for literal only, 0xF for both*/
    datatype	TINYINT not null,	/* 0x3(0011): string,  0xC(1100): int*/
    posnum		TINYINT not null,	/* Parameter position, index start from 1. e.g.,: APICall(pos1, pos2, pos3)*/
    paramdesc	TEXT
 );
 
 CREATE TABLE APIRESLT_TBL(
	id 			int AUTO_INCREMENT not null primary key,
	rsltname	char(32) not null,
    apiname 	char(32) not null,
    paramdesc	TEXT
 );
 
 
# 3. Insert inital test data
INSERT INTO APILIST_TBL VALUES(1, 'compareInt',  'Compare two integer values, only get a bool result without stop the procedure.'); 
INSERT INTO APIPARAM_TBL VALUES(1, 'ParamLeft', 'compareInt', 0x0E, 0x0C, 1, 'left source for compare, only accepty variables from: ^/?/$');
INSERT INTO APIPARAM_TBL VALUES(default, 'ParamOprt', 'compareInt', 0x01, 0x03, 2, 'operator for compare, only accepty literal string, like: eq, gt, lt, ge....'); 
INSERT INTO APIPARAM_TBL VALUES(default, 'ParamRght', 'compareInt', 0x0F, 0x0C, 3, 'Right source for compare, accepty both literal and variables from: ^/?/$'); 
INSERT INTO APIRESLT_TBL VALUES(1, 'retCmpIntResult', 'compareInt',  'Return lower case \"true\" or \"false\" as a result.');


# 4. Commit
commit;
 
 

