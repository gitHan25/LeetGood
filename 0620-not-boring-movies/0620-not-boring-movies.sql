# Write your MySQL query statement below
SELECT id,movie,description,rating FROM Cinema  WHERE id mod 2 != 0 and description !='boring' order by rating DESC