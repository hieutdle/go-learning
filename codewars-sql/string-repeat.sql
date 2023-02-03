/*
--# write your SQL statement here: you are given a table 'repeatstr' with columns 'n' and 's',
return a table with all columns and your result in a column named 'res'.

6, "I"     -> "IIIIII"
5, "Hello" -> "HelloHelloHelloHelloHello"

 */

SELECT s, n, REPEAT(s,n) as res
FROM repeatstr