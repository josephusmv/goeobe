# 1. Clean up
#DROP TABLE DBTEST_TBL;

# 2. Create Table
 CREATE TABLE DBTEST_TBL (
	id int AUTO_INCREMENT not null primary key,
    username char(32) not null,
    testfield1 char(32) not null,
    testfield2 char(128),
    testfield3 char(128)    
 );
 
# 3. Insert inital test data, permission in this test: 0xJKLH---> J = User mgmt, K=RWDA, L=RW, H=Readonly....
INSERT INTO DBTEST_TBL VALUES(1, 'sadmin', 'avaid', 'This Black History Month', '"Answer to the Ultimate Question of Life, the Universe, and Everything');
INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', 'avaid', 'This Black History Month', '"Answer to the Ultimate Question of Life, the Universe, and Everything');
INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', '11avaid', 'This Black Hi234story Month', '"Answersf to the Ultimate Questwion of Life, the Universe, and Everything');
INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', '222avaid', 'This Black History Month', '"Answer to the Ultimate Questionwe of Life, the Universe, and Everything');
INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', '23423avaid', 'This Black H234istory Month', '"Answer to the Ultimate Questwion of Life, the Universe, and Everything');
INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', '234avaid', 'This Black History Month', '"Answer to the Ultimate23 Question of Life, the Universe, and Everything');
INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', '334avaid', 'This Black Histo234ry Month', '"Answer to the Ult3rimate Question of Life, the Universe, and Everything');

INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', 'testchange', 'This Black Histo234ry Month', '"Answer to the Ult3rimate Question of Life, the Universe, and Everything');
INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', 'RANGEKEYA', 'This Black Histo234ry Month', '"Answer to the Ult3rimate Question of Life, the Universe, and Everything');
INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', 'RANGEKEYB', 'This Black Histo234ry Month', '"Answer to the Ult3rimate Question of Life, the Universe, and Everything');
INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', 'RANGEKEYC', 'This Black Histo234ry Month', '"Answer to the Ult3rimate Question of Life, the Universe, and Everything');
INSERT INTO DBTEST_TBL VALUES(default, 'sadmin', 'RANGEKEYD', 'This Black Histo234ry Month', '"Answer to the Ult3rimate Question of Life, the Universe, and Everything');


# 5. Verify
SELECT  * FROM DBTEST_TBL;
SELECT COUNT(*) FROM DBTEST_TBL;
SELECT  * FROM DBTEST_TBL WHERE testfield1 LIKE 'test%';
SELECT  * FROM DBTEST_TBL WHERE testfield1 LIKE 'RANGE%';
SELECT COUNT(*) FROM DBTEST_TBL WHERE testfield1 LIKE 'RANGEKEYC';


# delete test added
DELETE FROM DBTEST_TBL WHERE id > 40;

#commit;
commit;


# 1. Table for range get DBTEST_TBL 
#DROP TABLE DBRANGETEST_TBL;
# 2. Create Table
 CREATE TABLE DBRANGETEST_TBL (
	id int AUTO_INCREMENT not null primary key,
	srchkey int not null,
    testfield1 char(32) not null unique
 );


# 3. Insert
INSERT INTO DBRANGETEST_TBL VALUES(1, '2014', 'RANGEKEYA');
INSERT INTO DBRANGETEST_TBL VALUES(2, '1985', 'RANGEKEYB');
INSERT INTO DBRANGETEST_TBL VALUES(3, '2014', 'RANGEKEYC');
INSERT INTO DBRANGETEST_TBL VALUES(4, '2019', 'RANGEKEYD');
INSERT INTO DBRANGETEST_TBL VALUES(5, '2014', 'NOTEXSITED');



# 5. Verify
SELECT  * FROM DBRANGETEST_TBL;