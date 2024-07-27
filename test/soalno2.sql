
SELECT u.name, 
       o.amount, 
       o.created_at
FROM users u
JOIN orders o ON u.id = o.user_id
ORDER BY u.name, o.created_at;
