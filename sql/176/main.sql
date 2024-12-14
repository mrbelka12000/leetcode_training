select (
           select salary from(
                                 select distinct salary ,  dense_rank() over (ORDER BY salary DESC NULLS LAST) as rank from Employee
                             ) where rank = 2

       ) as secondhighestsalary;
