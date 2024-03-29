create table if not exists blog_entry
(
    id INTEGER,
    title VARCHAR2(255),
    comment_id INTEGER,
    PRIMARY KEY(id),
    FOREIGN KEY(comment_id)
    REFERENCES(comment,id)
)