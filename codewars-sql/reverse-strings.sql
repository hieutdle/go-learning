-- write your SQL statement here: you are given a table 'solution' with column 'str'
-- return a table with column 'str' and your result in a column named 'res'.
-- 'world'  =>  'dlrow'
-- 'word'   =>  'drow'

SELECT str, REVERSE(str) AS res
FROM solution