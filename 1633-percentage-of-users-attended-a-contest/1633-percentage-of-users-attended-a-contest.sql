WITH ContestParticipation AS (
   SELECT 
       contest_id, 
       COUNT(DISTINCT user_id) AS registered_users,
       (SELECT COUNT(*) FROM Users) AS total_users
   FROM Register
   GROUP BY contest_id
)
SELECT 
   contest_id, 
   ROUND(registered_users * 100.0 / total_users, 2) AS percentage 
FROM ContestParticipation
ORDER BY 
   percentage DESC, 
   contest_id ASC