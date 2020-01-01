CREATE SEQUENCE users_id_seq;

CREATE TABLE users(
    id integer default nextval('users_id_seq') NOT NULL,
    username text NOT NULL,
    password text  NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL
);

CREATE SEQUENCE watchlist_id_seq;

CREATE TABLE watchlist(
    id integer default nextval('watchlist_id_seq') NOT NULL,
    owner_id integer NOT NULL,
    name text NOT NULL,
    created_on timestamp without time zone NOT NULL,
    items integer[]
);

CREATE SEQUENCE items_id_seq;

CREATE TABLE items(
    id integer default nextval('items_id_seq') NOT NULL,
    watchlist_id integer NOT NULL,
    item_type integer NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    release_date timestamp without time zone NOT NULL,
    rating text NOT NULL,
    genre text NOT NULL,
    watched boolean NOT NULL
);

CREATE SEQUENCE itemtypes_id_seq;

CREATE TABLE item_types(
    id integer,
    name text NOT NULL
);