CREATE TABLE reviews
(
    id         SERIAL  NOT NULL,
    user_id    integer NOT NULL,
    product_id integer NOT NULL,
    point integer NOT NULL,
    text VARCHAR(255),
    PRIMARY KEY (id)
)