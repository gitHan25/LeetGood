# Write your MySQL query statement below

select name,bonus
from Employee
left join Bonus on Employee.empId = Bonus.empId
WHERE bonus < 1000 OR Bonus IS NULL
