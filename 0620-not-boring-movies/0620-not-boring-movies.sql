SELECT id, movie, description, rating 
FROM (
    SELECT *, ROW_NUMBER() OVER (ORDER BY id) as row_num 
    FROM Cinema
) as numbered_cinema
WHERE 
    row_num % 2 = 1 AND 
    description != 'boring'
ORDER BY rating DESC;