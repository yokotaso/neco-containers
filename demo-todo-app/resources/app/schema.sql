create database if not exists todo;

create table if not exists todo.todo (
    todo_id bigint(20) not null AUTO_INCREMENT,
    todo_title varchar(30),
    finished boolean,
    created_at timestamp,
    primary key (todo_id)
);
