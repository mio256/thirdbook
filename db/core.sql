create table users (
    id bigserial primary key,
    name varchar(255) not null,
    email varchar(255) unique not null,
    password varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

create type booking_type as enum ('pending', 'approved', 'canceled', 'rejected');

create table bookings (
    id bigserial primary key,
    name varchar(255) not null,
    start_time timestamp not null,
    end_time timestamp not null,
    user_id bigint not null,
    status booking_type not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    constraint bookings_start_end_check check (start_time < end_time)
);

alter table bookings add constraint fk_bookings_user_id foreign key (user_id) references users(id);
