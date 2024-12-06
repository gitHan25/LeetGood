WITH EmployeeDetail AS (
    SELECT p.project_id, p.employee_id, e.name, e.experience_years 
    FROM Project p 
    JOIN Employee e ON p.employee_id = e.employee_id
)
SELECT 
    project_id, 
    
    ROUND(AVG(experience_years),2) as average_years
FROM EmployeeDetail
GROUP BY project_id;