create or replace function downgrade_value(value int)
	RETURNS integer
	LANGUAGE plpgsql
	STRICT
AS $function$
declare
output int := 0;
begin
output := value - 1;
RETURN (output);
END;
$function$;

drop table  if exists user_table;
drop table  if exists category;
drop table  if exists product_table;
drop table  if exists shopping_cart_table;

--остальное:
create table user_table( login varchar(100) PRIMARY KEY NOT NULL,
                         password varchar(100) NOT NULL,
                         role varchar(100) DEFAULT 'customer');

create table category(id_category int PRIMARY KEY,
                      name_category text NOT NULL);

--для продавца:
create table product_table(id_product int PRIMARY KEY,
                           id_seller varchar(100) REFERENCES user_table(login) NOT NULL,
                           product_name text NOT NULL,
                           product_category_id int REFERENCES category(id_category) NOT NULL,
                           product_description text not null default '-'::text,
                           product_image bytea not null default 'https://github.com/io-qar/Market-Bot/blob/master/default_image.jpg?raw=true',
                           product_cost float4 not null default 0.0,
                           product_availability int NOT NULL);


--для покупателя:
create table shopping_cart_table(id_cart int PRIMARY KEY,
                                 id_user varchar(100) REFERENCES user_table(login),
                                 id_product int REFERENCES product_table(id_product));




--insert into scientist (id, firstname, lastname) values (1, 'albert', 'einstein');
--insert into scientist (id, firstname, lastname) values (2, 'isaac', 'newton');
--insert into scientist (id, firstname, lastname) values (3, 'marie', 'curie');
--select * from scientist;