/* return to which quarter of the year it belongs as an integer number.
   For example: month 2 (February), is part of the first quarter,
   and month 11 (November), is part of the fourth quarter.*/

SELECT month, CAST(CEILING(month / 3.0) AS INT) AS res
FROM quarterof