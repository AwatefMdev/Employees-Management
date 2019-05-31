create table user (
    email varchar(150) not null,
    password char(64) not null,
    id_employee integer (50),
    constraint id_employee
    foreign key (id) 
     REFERENCES employee (id),
     id_role integer REFERENCES role (id),
    created_at timestamp default current_timestamp
);

create table Employee (
    id serial primary key,
    title varchar(150) not null,
    description text not null,
    user_id int not null,
    created_at timestamp default current_timestamp
);
