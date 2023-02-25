create table account_profiles (
    id serial NOT NULL,
    name varchar(100),
    account_id int,
    PRIMARY KEY (id),
    foreign key (account_id) references accounts(id)
)