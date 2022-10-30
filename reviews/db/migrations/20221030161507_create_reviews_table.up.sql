CREATE TABLE reviews
(
    id         SERIAL  NOT NULL,
    user_id    integer NOT NULL,
    product_id integer NOT NULL,
    point integer NOT NULL,
    PRIMARY KEY (id)
)