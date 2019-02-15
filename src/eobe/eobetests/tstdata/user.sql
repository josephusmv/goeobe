# 1. Clean up
DROP TABLE USER_TBL;


# 2. Create Table
 CREATE TABLE USER_TBL (
	id int AUTO_INCREMENT not null primary key,
    username char(64) not null UNIQUE,
    userpwd char(128) not null,
    permission int not null
 );
 
 
# 3. Insert inital test data, permission in this test: 0xJKLH---> J = User mgmt, K=RWDA, L=RW, H=Readonly....
INSERT INTO USER_TBL VALUES(1, 'sadmin', 'superadmin_enc', 0xFFFF); 
INSERT INTO USER_TBL VALUES(2, 'freshboy', 'freshfresh_enc', 0x0001);
INSERT INTO USER_TBL VALUES(3, 'tommy', 'tommy_enc', 0x0FFF);
INSERT INTO USER_TBL VALUES(4, 'Alex',  'alex_enc',  0x00FF);
INSERT INTO USER_TBL VALUES(5, 'Kate',  'kate_enc',  0x00FF);
INSERT INTO USER_TBL VALUES(6, 'Mike',  'mike_enc',  0x00FF);
INSERT INTO USER_TBL VALUES(7, 'Brow',  'brow_enc',  0x00FF);
INSERT INTO USER_TBL VALUES(8, 'Jung',  'jung_enc',  0x00FF);
INSERT INTO USER_TBL VALUES(9, 'John',  'john_enc',  0x00FF);

# 4. Commit
commit;
 
# 5. Verify
SELECT  * FROM USER_TBL;
SELECT count(*) FROM USER_TBL;

# 6. delete some test intermediate results
DELETE FROM USER_TBL WHERE username='Mickey';


################################################################################################################################################################################################
################################################################################################################################################################################################
#1. Drop Table
DROP TABLE USERPRFL_TBL;


# 2. Create Table
 CREATE TABLE USERPRFL_TBL (
	id int AUTO_INCREMENT not null primary key,
    username char(32) not null UNIQUE,
    age int,
    phone1 char(32) not null,
    phone2 char(32),
    Address TEXT,
    crtdttmcrtdt timestamp not null
 );
 
 #2.1. Change table collection
 alter table USERPRFL_TBL convert to character set utf8 collate utf8_unicode_ci;
 
 
# 3. Insert inital test data
INSERT INTO USERPRFL_TBL VALUES(1, 'sadmin',  24, '710-245-2742623-01', '',  '', '1999-12-31 13:29:19'); 
INSERT INTO USERPRFL_TBL VALUES(2, 'freshboy', 15, '710-265-8391823-01', '', 
		' IPアドレス · MACアドレス · メールアドレス · Uniform Resource Locator（URL）の俗称。 ゴルフ等で打球の準備としての構え。 スズキ・アドレス - スズキが製造販売しているスクータータイプのオートバイ。 ADRES（アドレス） - かつて東芝が開発したノイズリダクション方式。 ADDRESS (アルバム) - 山崎まさよしの ...',
		'2016-10-04 13:27:27');
INSERT INTO USERPRFL_TBL VALUES(3, 'tommy',    46, '210-245-1695623-01', '410-245-9395643-01', 
		'Morgan Stanley Japan Holdings Co., Ltd. Morgan Stanley MUFG Securities Co., Ltd. Morgan Stanley Investment Management (Japan) Co., Ltd. Morgan Stanley Capital K.K.. Morgan Stanley Japan Group Co., Ltd. Otemachi Financial City South Tower 1-9-7 Otemachi, Chiyoda-ku, Toky',
		'2017-10-24 15:57:59');
INSERT INTO USERPRFL_TBL VALUES(4, 'Alex',     64, '410-215-9395643-01', '410-215-9395643-01',
		'Los Angeles Joins Smart City Consortium to Develop Model for IoT Connected Cities. 3. Four People Running on a ... LGBTQ Youth. 3. Los Angeles County Fire Air Operations helicopter uses a fire suppressant Saturday to stop the spread .... Mayor\'s Inaugural Los',
		'2016-12-21 23:19:22');
INSERT INTO USERPRFL_TBL VALUES(5, 'Kate',     37, '010-241-8395623-01', '', '', '2014-02-18 09:57:26');
INSERT INTO USERPRFL_TBL VALUES(6, 'Mike',     32, '014-255-1995663-01', '410-215-9395643-01', 'Building-1, some place for test', '2014-09-30 11:29:59');
INSERT INTO USERPRFL_TBL VALUES(7, 'Brow',     14, '017-245-4951823-01', '',
		'依教育部「中文譯音使用原則」規定，我國中文譯音以漢語拼音為準，進入教育部「中文譯音轉換系統」。 2. 本系統地名譯寫結果，僅供交寄郵件英文書寫參考（請勿作為其他用途書',
		'2017-10-11 23:36:36');
INSERT INTO USERPRFL_TBL VALUES(8, 'Jung',     38, '070-245-742823-01' , '410-215-742823-01' ,
		 'King County, which includes Seattle, is a major center for liberal politics and is a bastion for the Democratic Party. No Republican presidential candidate has garnered the majority of the county\'s votes since Ronald Reagan\'s landslide reelection victory in 1984. In the 2008 election, Barack Obama defeated John McCain in the county by 42 percentage points, a larger margin than any previous election. Slightly more than 29% of the population in the State of Washington reside in King County, making it a significant factor for the Democrats in a few recent close statewide elections. In 2000, it was King County that pushed Maria Cantwell\'s total over that of incumbent Republican Slade Gorton, winning her a seat in the United States Sena',
		 '2016-05-25 13:59:59');
INSERT INTO USERPRFL_TBL VALUES(9, 'John',     18, '022-245-4663867-01', '', '', '2018-01-06 23:59:34');

# 4. Commit
commit;
 
# 5. Verify
SELECT  * FROM USERPRFL_TBL;
SELECT count(*) FROM USERPRFL_TBL;

# 6. delete some test intermediate results
DELETE FROM USERPRFL_TBL WHERE username='Lee';
