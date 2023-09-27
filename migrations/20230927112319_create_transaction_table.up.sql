CREATE TABLE "transaction" (
      id serial not null primary key ,
      user_id int not null ,
      saving_account_id int not null ,
      transaction_code varchar(255) not null ,
      amount float not null,
      created_at timestamp
)