-- write your SQL statement here: you are given a table 'codewars-go' with a column 'n',
-- return a table with 'n' and your result stored in a column naed 'res'.

-- 2 -> 3 (1 + 2)
-- 8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)

SELECT n, n*(n+1)/2 as res
FROM kata