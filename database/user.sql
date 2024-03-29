create table if not exists user
(
    id INTEGER,
    name VARCHAR2(255),
    email VARCHAR2(255),
    blog_id INTEGER,
    PRIMARY KEY(id),
    FOREIGN KEY(blog_id),
    REFERENCES blog_entry (id)
             ON DELETE CASCADE
             ON UPDATE NO ACTION,
)

