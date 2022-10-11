CREATE TABLE todos
(
    id SERIAL NOT NULL,
    content VARCHAR NOT NULL ,
    user_id INT not null,
    PRIMARY KEY (id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
)