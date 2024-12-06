# Write your MySQL query statement below

SELECT name
FROM Employee e1
JOIN(
    SELECT managerID
    FROM Employee
    GROUP BY managerId
    HAVING COUNT(managerId)>=5
) e2

on e1.id = e2.managerId

