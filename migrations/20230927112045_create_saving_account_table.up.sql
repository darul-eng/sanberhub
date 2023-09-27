CREATE TABLE "saving_account" (
    id serial not null primary key ,
    user_id int not null ,
    account_number int not null ,
    balance float not null
)