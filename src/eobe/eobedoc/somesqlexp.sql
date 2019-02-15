
CREATE TABLE APPTASKSTEST2_TBL (
	id			INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	username 	VARCHAR(16) NOT NULL, 
    title		CHAR(64) NOT NULL, 
    details 	MEDIUMTEXT,
    crtdt		TIMESTAMP
)CHARACTER SET 'utf8' COLLATE 'utf8_general_ci';



delimiter $$
	create procedure FixedRandomChnWorkTEST2(num int)
	begin
			declare i int default 1;

			while i <=num do
				insert into APPTASKSTEST2_TBL  (title, username, details, crtdt)  
						values(substring('产生随机数小abcdefghijklmnopqrstuvwxyz1234567890. _结一下这里有俄神吐槽：华裔夫宅区街道被否媒：俄圣彼神吐槽：华裔夫妇神吐槽：华裔夫妇买豪宅区街道被否买豪abcdefghijklmnopqrstuvwxyz1234567890宅区街道被否得abcdefghijklmnopqrstuvwxyz1234567890堡民潘神吐槽：华裔夫妇买豪宅区街道被否石屹：帮家里有俄媒：俄圣彼得堡里有俄媒：俄圣彼得堡民民彼得堡民abcdefghijklmnopeqrstuvwxyz潘石屹乡彼得堡民abcdABCEDFGHIJKLMNOPQRSTUWVWXYZ妇买豪宅区街道被否媒：俄圣彼神吐槽：华裔夫妇神吐槽：华裔夫妇买豪宅区街道被否买豪abcdefghijklmnopqrstuvwxyz1234567890宅区街道被否得abcdefghijklmnopqrstuvwxyz1234567890堡民潘神吐槽：华裔夫妇买豪宅区街道被否石屹：帮家里有俄媒：俄圣彼得堡里有俄媒：俄圣彼得堡民民彼得堡民abcdefghijklmnopeqrstuvwxyz潘石屹乡彼得堡民abcdABCEDFGHIJKLMNOPQRSTUWVWXYZ潘石屹卖了一个彼得堡民abcdefghijklmnopqrstuvwxyz1234567890潘石屹亿苹果 祖坟却被扒了房神吐槽：华裔夫妇买豪宅区街道被否内发现一中神吐槽：华裔夫妇买豪宅区街道被否国女学生全身一使用该随机数发生器种子值个表专门产生互联网打造中国物流“大脑” | “乌镇时间” | 十九大 随机的字符串', floor(1+128*rand()), 8),  
							GendRandString(16), 
							GendRandString(128), 
							DATE_SUB(NOW(), INTERVAL (465*RAND()) DAY)); 				
				set i = i + 1;
			end while;
	end
    $$
delimiter ;


call FixedRandomChnWorkTEST2(1901); 



insert into APPTASKSTEST2_TBL  (title, username, details, crtdt)  
		values(substring('产生随机数小结一下这里有俄神吐槽：华裔夫妇买豪宅区街道被否媒：俄圣彼神吐槽：华裔夫妇神吐槽：华裔夫妇买豪宅区街道被否买豪宅区街道被否得堡民潘神吐槽：华裔夫妇买豪宅区街道被否石屹：帮家里有俄媒：俄圣彼得堡里有俄媒：俄圣彼得堡民民彼得堡民潘石屹乡彼得堡民潘石屹卖了一个彼得堡民潘石屹亿苹果 祖坟却被扒了房神吐槽：华裔夫妇买豪宅区街道被否内发现一中神吐槽：华裔夫妇买豪宅区街道被否国女学生全身一使用该随机数发生器种子值个表专门产生互联网打造中国物流“大脑” | “乌镇时间” | 十九大 随机的字符串', floor(1+128*rand()), 8),  
					'testbenchmark',
                    GendRandString(128), 
                    DATE_SUB(NOW(), INTERVAL (465*RAND()) DAY)); 

commit;
     
select randChr(); 

select substring('产生随机数小结一下这里有一使用该随机数发生器种子值个表专门产生随机的字符串', floor(1+20*rand()), 1);
     
     
SELECT COUNT(*) FROM APPTASKSTEST2_TBL;


SELECT COUNT(*)  FROM APPTASKSTEST2_TBL WHERE username='testbenchmark';
SELECT COUNT(*)  FROM APPTASKS_TBL WHERE username='testbenchmark';

SELECT * FROM APPTASKSTEST2_TBL WHERE username='testbenchmark';
#SELECT * FROM APPTASKS_TBL WHERE WHERE (title!=:1 AND username=:2) OR crdt > :3;



SELECT * FROM APPTASKS_TBL where TRUE;

DESC APPTASKS_TBL;



