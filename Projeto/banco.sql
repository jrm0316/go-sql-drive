create database bank;
use bank;
create table tabela(nome varchar(30) NOT NULL, idade int(5));
insert into tabela(nome, idade) values ("Juliano", 42);
select * from tabela;
truncate tabela;
