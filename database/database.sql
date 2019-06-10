

create table employee (
    id serial primary key,
    firstname varchar(150) not null,
    lastname varchar(150) not null,
    email varchar(150) not null,
    adress varchar(150) not null,
    gender varchar(150) not null,
    mobilenumber INT (50), 
    description text not null,
    idtools integer REFERENCES tools (id),
    idattendance integer REFERENCES attendance (id),
    idleaves integer REEFERENCES leaves (id),
    created_at timestamp default current_timestamp
);

create table tools(
    email varchar(150) not null,
    password char(64) not null,
    id_employee integer (50),
    constraint id_employee
    foreign key (id) 
     REFERENCES employee (id),
     id_role integer REFERENCES role (id),
    created_at timestamp default current_timestamp
);
create table attendance (
    email varchar(150) not null,
    password char(64) not null,
    id_employee integer (50),
    constraint id_employee
    foreign key (id) 
     REFERENCES employee (id),
     id_role integer REFERENCES role (id),
    created_at timestamp default current_timestamp
);

create table leaves(
    email varchar(150) not null,
    password char(64) not null,
    id_employee integer (50),
    constraint id_employee
    foreign key (id) 
     REFERENCES employee (id),
     id_role integer REFERENCES role (id),
    created_at timestamp default current_timestamp
);
create table meetingroom (
    email varchar(150) not null,
    password char(64) not null,
    id_employee integer (50),
    constraint id_employee
    foreign key (id) 
     REFERENCES employee (id),
     id_role integer REFERENCES role (id),
    created_at timestamp default current_timestamp
);

create table training (
    email varchar(150) not null,
    password char(64) not null,
    id_employee integer (50),
    constraint id_employee
    foreign key (id) 
     REFERENCES employee (id),
     id_role integer REFERENCES role (id),
    created_at timestamp default current_timestamp
);

create table employeemeetingroom(
    email varchar(150) not null,
    password char(64) not null,
    id_employee integer (50),
    constraint id_employee
    foreign key (id) 
     REFERENCES employee (id),
     id_role integer REFERENCES role (id),
    created_at timestamp default current_timestamp
);

create table employeemeetingroom(
    email varchar(150) not null,
    password char(64) not null,
    id_employee integer (50),
    constraint id_employee
    foreign key (id) 
     REFERENCES employee (id),
     id_role integer REFERENCES role (id),
    created_at timestamp default current_timestamp
);





