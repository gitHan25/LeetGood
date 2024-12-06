SELECT 
    employee_id, 
    name, 
    reports_count, 
    ROUND(average_age, 0) AS average_age
FROM (
    SELECT 
        e1.employee_id, 
        e1.name, 
        COUNT(e2.employee_id) AS reports_count,
        AVG(e2.age) AS average_age
    FROM 
        Employees e1
    LEFT JOIN 
        Employees e2 ON e1.employee_id = e2.reports_to
    GROUP BY 
        e1.employee_id, e1.name
    HAVING 
        reports_count > 0
) AS manager_stats
ORDER BY 
    employee_id;