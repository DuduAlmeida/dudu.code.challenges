# Write your MySQL query statement below
with
employee_department_ranked as (
    select 
        d.name as department,
        e.name as employee,
        e.salary as salary,
        DENSE_RANK() OVER(PARTITION BY d.name ORDER BY e.salary DESC) AS rn_salary
    from Employee as e
    left join Department as d on e.departmentId = d.id
)

select 
    department as Department,
    employee as Employee,
    salary as Salary
    #, rn_salary
from employee_department_ranked
where  rn_salary <= 3