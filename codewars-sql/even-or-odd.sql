-- column is_even containing "Even" or "Odd" depending on number column values.

SELECT
    CASE
        WHEN MOD(number,2) = 0 THEN 'Even'
        ELSE 'Odd'
        END as "is_even"
from numbers
