create table if not exists dbo.fridges
(
    id   int auto_increment
        primary key,
    name varchar(100) null,
    constraint uix_fridges_name
        unique (name)
);

create table if not exists dbo.groceries
(
    id        int auto_increment
        primary key,
    name      varchar(100) null,
    quantity  double       null,
    unit      varchar(50)  null,
    category  varchar(50)  null,
    expiry    date         null,
    fridge_id int          null,
    constraint uix_groceries_name
        unique (name),
    constraint groceries_fridge_id_fridges_ID_foreign
        foreign key (fridge_id) references dbo.fridges (id)
            on delete cascade
);

