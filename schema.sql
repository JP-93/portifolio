create table revenue (
    revenue_id serial,
    descricao varchar(50),
    valor int,
    data_into varchar,
    primary key (user_id)
);

create table expenses (
        ex_user_id serial,
            user_id int references revenue(user_id),
        descricao varchar(50),
        valor int,
        data_into date,
                      primary key (ex_user_id)
);