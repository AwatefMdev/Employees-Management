ALTER TABLE employee 
ADD CONSTRAINT leavesfkey FOREIGN KEY (idleaves) REFERENCES leaves (id);
ALTER TABLE employee 
ADD CONSTRAINT toolsfkey FOREIGN KEY (idtools) REFERENCES tools(id);
ALTER TABLE employee 
ADD CONSTRAINT attendancefkey FOREIGN KEY (idattendance) REFERENCES  attendance(id);
ALTER TABLE employee_training 
ADD CONSTRAINT employeefkey FOREIGN KEY (idemployee) REFERENCES employee(id);
ALTER TABLE parking
ADD CONSTRAINT employeeparkingfkey FOREIGN KEY (idemployee) REFERENCES employee(id);
ALTER TABLE role
ADD CONSTRAINT employeefkey FOREIGN KEY (idemployee) REFERENCES employee(id);
ALTER TABLE employee_training 
ADD CONSTRAINT trainingfkey FOREIGN KEY (idtraining) REFERENCES training(id);




create table employee (
    id serial primary key,
    firstname varchar(150) not null,
    lastname varchar(150) not null,
    email varchar(150) not null,
    adress varchar(150) ,
    gender varchar(150) ,
    mobilenumber integer , 
    description text,
    idtools integer,
    idattendance integer,
     idleaves integer,
    created_at timestamp default current_timestamp
);

create table tools(
   id serial primary key,
    name varchar(45),
    quantity int,
    date Date,  
    created_at timestamp default current_timestamp
);
create table attendance (
    id serial primary key,
   timecheckin time,
   timecheckout time,
    created_at timestamp default current_timestamp
);

create table leaves(
    id serial primary key,
  type varchar(255),
    starttime time,
   endtime time ,
   reason varchar (255),
    created_at timestamp default current_timestamp
);
create table meetingroom (
    id serial primary key,
   capacity integer,
    description varchar(255),
    created_at timestamp default current_timestamp
);

create table training (
    id serial primary key,
    type varchar(45),
    date Date,
    duration int,
    link varchar (255),
    created_at timestamp default current_timestamp
);

create table employeemeetingroom(
    id serial primary key,
   date Date,
    employee_invited varchar (255),
    timeend time,
    timestart time,
    created_at timestamp default current_timestamp
);

create table parking(
   id serial primary key,
   capacity int ,
    date date,
    newcar varchar (255)
    created_at timestamp default current_timestamp
);

create table role(
   id serial primary key, 
    label varchar (255),
    created_at timestamp default current_timestamp
);


create table employee_training(
   id serial primary key, 
    created_at timestamp default current_timestamp
);



