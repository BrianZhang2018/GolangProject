create table test
(
    id   int,
    name varchar(100)
);

create table poll (
    pollId int auto_increment PRIMARY KEY,
    title varchar(100),
    description varchar(300),
    options varchar(100),
    urlLink varchar(100),
    ts timestamp default current_timestamp
);

create table pollResponse (
  responseId int auto_increment PRIMARY KEY,
  pollId int,
  description varchar(300),
  selectedOptions varchar(100),
  userId varchar(50),
  ts timestamp default current_timestamp
);

create table User(
    userId int auto_increment PRIMARY KEY,
    name varchar(200),
    age int,
    pwd varchar(200)
)