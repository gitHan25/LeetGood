SELECT 
    contest_id, 
    ROUND(COUNT(r.user_id) * 100.0 / (SELECT COUNT(*) FROM Users),2) AS percentage 
FROM Register r 
JOIN Users u ON r.user_id = u.user_id 
GROUP BY contest_id
order by 
percentage DESC,
contest_id ASC