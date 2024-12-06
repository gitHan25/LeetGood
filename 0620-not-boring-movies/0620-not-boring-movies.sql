# Write your MySQL query statement below
SELECT id, movie, description, rating
FROM Cinema
WHERE (id & 1) != 0 
  AND description != 'boring'
ORDER BY rating DESC;
