insert into peach 
(id, size, juice)
values ($1, $2, $3)
returning id;
