# Write your MySQL query statement below
select
    customer_id,count(Visits.visit_id)  as count_no_trans
from 
    Visits
left join
    Transactions
on
    (Visits.visit_id = Transactions.visit_id) 
  Where  
    Transactions.transaction_id  is NULL
group by Visits.customer_id
order by count(Visits.visit_id) desc