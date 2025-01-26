update peach
set 
    id = $1, 
    size = $2, 
    juice = $3
where id = $1;
