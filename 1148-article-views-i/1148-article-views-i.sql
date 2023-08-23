# Write your MySQL query statement below
SELECT  author_id as id
FROM Views
WHERE (viewer_id = author_id)>=1
group by author_id
order by author_id asc
