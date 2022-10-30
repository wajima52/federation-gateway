CREATE TABLE products
(
    id         SERIAL       NOT NULL,
    name       VARCHAR(255) NOT NULL,
    price      integer NOT NULL ,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
)