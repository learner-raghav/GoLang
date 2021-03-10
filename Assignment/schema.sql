use testDB;
create table brand(brand_id int auto_increment, brand_name varchar(20), PRIMARY KEY(brand_id));
create table model(model_id int auto_increment, model_name varchar(20), brand_id int, PRIMARY KEY(model_id));
create table variant(variant_id int auto_increment, model_id int, name varchar(20), disp float, peak_power float, peak_torque float, PRIMARY KEY(variant_id));
alter table model add Foreign key(brand_id) references brand(brand_id);
alter table variant add Foreign key(model_id) references model(model_id);
