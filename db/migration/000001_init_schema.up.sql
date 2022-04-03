
create schema if not exists registration;
create schema if not exists finance;
create schema if not exists access;

create table if not exists registration.building_type
(
    id          int         not null generated always as identity
        constraint pk_building_type primary key,
    description varchar(30) not null
        constraint uk_building_type unique
)
;

create table if not exists registration.condo_type
(
    id          int         not null generated always as identity
        constraint pk_condo_type primary key,
    description varchar(30) not null
        constraint uk_condo_type unique
)
;

create table if not exists registration.supplier_type
(
    id          int         not null generated always as identity
        constraint pk_supplier_type primary key,
    description varchar(30) not null
        constraint uk_supplier_type unique
)
;
create table if not exists registration.condo
(
    id            int         not null generated always as identity
        constraint pk_condo primary key,
    id_condo_type int         not null
        constraint fk_condo_01 references registration.condo_type,
    name          varchar(50) not null
        constraint uk_condo_01 unique,
    nickname      varchar(30) not null
        constraint uk_condo_02 unique,
    address1      varchar(30),
    address2      varchar(30),
    phone1        varchar(20),
    phone2        varchar(20),
    email         varchar(30) not null,
    cnpj          bigint      not null
        constraint uk_condo_03 unique
)
;

create table if not exists registration.building
(
    id               int         not null generated always as identity
        constraint pk_building primary key,
    id_condo         int         not null
        constraint fk_building_01 references registration.condo,
    id_building_type int         not null
        constraint fk_building_02 references registration.building_type,
    description      varchar(30) not null,
    constraint uk_building unique (id_building_type, description)
)
;

create table if not exists registration.supplier
(
    id               int         not null generated always as identity
        constraint pk_supplier primary key,
    id_supplier_type int         not null
        constraint fk_supplier_01 references registration.supplier_type,
    description      varchar(30) not null
        constraint uk_supplier unique
)
;


create table if not exists finance.movement_type
(
    id                  int         not null generated always as identity
        constraint pk_movement_type primary key,
    description         varchar(30) not null
        constraint uk_movement_type unique,
    direction           char(2)     not null
        constraint ch_movement_type_01 check (direction = 'RE' OR direction = 'DF' OR direction = 'OD'), -- receitas, despesas fixas, outras despesas
    id_default_supplier int
        constraint fk_movement_type_01 references registration.supplier (id)
)
;

create table if not exists registration.unity_type
(
    id          int         not null generated always as identity
        constraint pk_unity_type primary key,
    description varchar(30) not null
        constraint uk_unity_type unique
)
;

create table if not exists registration.unity
(
    id              int         not null generated always as identity
        constraint pk_unity primary key,
    id_building     int         not null
        constraint fk_unity_02 references registration.building,
    id_unity_type   int         not null
        constraint fk_unity_01 references registration.unity_type,
    unity_number    varchar(30) not null,
    constraint uk_unity unique (id_building, unity_number)
)
;

create table if not exists access.user_type
(
    id          int generated always as identity
        constraint pk_user_type primary key,
    description varchar(30) not null
        constraint uk_user_type unique
)
;

create table if not exists access.user
(
    id     int generated always as identity
        constraint pk_user primary key,
    login  varchar(20)                                 not null
        constraint uk_user unique,
    email  varchar(30)                                 not null,
    name   varchar(30)                                 not null,
    phone  varchar(20),
    hash   varchar(32)                                 not null,
    active boolean                                     not null default True
)
;

create table if not exists access.user__condo
(
    id_user      int not null
        constraint fk_user__condo_01 references access.user,
    id_condo     int not null
        constraint fk_user__condo_02 references registration.condo,
    id_user_type int not null
        constraint fk_user__condo_03 references access.user_type,
    constraint pk_user__condo primary key (id_user, id_condo)
)
;

