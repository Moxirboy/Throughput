create database kirim;
use kirim;
create table goods(
                      id int auto_increment primary key,
                      name varchar(45),
                      sort varchar(45)
);
create table client(
                       id int auto_increment primary key,
                       name varchar(45),
                       created date

);
create table purchase(
                         id int auto_increment primary key,
                         name varchar(45),
                         client_id int ,
                         foreign key(client_id) references client(id)
);

create table purchase_goods(
                               id int auto_increment primary key,
                               purchase_id int, foreign key(purchase_id) references purchase(id),
                               goods_id int,foreign key(goods_id) references goods(id),
                               amount decimal,
                               cort_price int
);
create table requirement(
                            id int auto_increment primary key,
                            date date,
                            client_id int,foreign key(client_id) references client(id)
);
create table requirement_goods(
                                  id int auto_increment primary key,
                                  requirement_id int, foreign key(requirement_id) references requirement(id),
                                  goods_id int,foreign key(goods_id) references goods(id),
                                  amount decimal,
                                  cost_cell int

);