

drop table  if exists user_table;
drop table  if exists category;
drop table  if exists product_table;
drop table  if exists shopping_cart_table;


--остальное:
create table user_table(login varchar(100) PRIMARY KEY NOT NULL,
                         password varchar(100) NOT NULL,
                         role varchar(100) DEFAULT 'customer');

create table category(id_category serial,
                      name_category text NOT NULL PRIMARY KEY,
                      number_of_product_ads int default 0);

--для продавца:
create table product_table(id_product serial PRIMARY KEY,
                           id_seller varchar(100) REFERENCES user_table(login) NOT NULL,
                           product_name text NOT NULL,
                           product_category text REFERENCES category(name_category) NOT NULL,
                           product_description text not null default '-'::text,
                           product_image bytea not null default pg_read_binary_file('/imgs/default_image.jpg'),
                           product_cost int not null default 0,
                           product_availability int NOT NULL);

--для покупателя:
create table shopping_cart_table(id_cart int PRIMARY KEY,
                                 id_user varchar(100) REFERENCES user_table(login),
                                 id_product int REFERENCES product_table(id_product));

-- заполнение категорий товара:
insert into category(name_category) values ('Одежда и обувь');
insert into category(name_category) values ('Аксессуары к одежде');
insert into category(name_category) values ('Бытовая техника');
insert into category(name_category) values ('Электроника');
insert into category(name_category) values ('Детские товары');
insert into category(name_category) values ('Товары для хобби');
insert into category(name_category) values ('Товары для дома и сада');
insert into category(name_category) values ('Бытовая химия');
insert into category(name_category) values ('Косметика');
insert into category(name_category) values ('Остальные категории');


-- триггеры для таблицы category
create or REPLACE function categoryUp()
    returns trigger AS
    $$
begin
update category
set  number_of_product_ads = number_of_product_ads + 1
where name_category = new.product_category;
return new;
end;
$$
language 'plpgsql';

create or REPLACE function categoryDown()
    returns trigger AS
    $$
begin
update category
set  number_of_product_ads = number_of_product_ads - 1
where name_category = new.product_category;
return new;
end;
$$
language 'plpgsql';


create trigger category_after_insert
    after insert on product_table
    for each row
    execute  procedure categoryUp();

create trigger category_after_delete
    after delete on product_table
    for each row
    execute  procedure categoryUp();


