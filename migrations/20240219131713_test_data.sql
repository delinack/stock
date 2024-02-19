-- +goose Up
-- +goose StatementBegin
INSERT INTO items (id, name, size, quantity, created_at) VALUES ('f2aaa171-0118-497b-a0f3-bcffb82533f1', 'dress', 'S', 10, now()),
                                                                ('38293088-21ab-4f10-9217-eca1f42aaa5d', 't-shirt', 'S', 26, now()),
                                                                ('91829929-a842-4a69-a5c5-750dd0074849', 'top', 'M', 39, now()),
                                                                ('00ef538f-02c2-4a23-8695-385918026262', 'socks', 'no_size', 43, now()),
                                                                ('36011ed2-2ae7-4ee9-8567-5900b7b230e3', 'hat', 'XS', 12, now()),;

INSERT INTO stocks (id, name, is_available, created_at) VALUES ('5fe06170-4fb3-429a-b950-1ae1a037376e', 'moscow', true, now()),
                                                               ('719ece6f-65fc-4cc2-b542-7dc673c6c6a8', 'vladivostok', true, now()),
                                                               ('0cc27199-a03a-4f22-8c43-51ad6f336baf', 'omsk', false, now());

INSERT INTO items_stocks (stock_id, item_id, quantity, created_at) VALUES ('5fe06170-4fb3-429a-b950-1ae1a037376e', 'f2aaa171-0118-497b-a0f3-bcffb82533f1', 5, now()),
                                                                          ('719ece6f-65fc-4cc2-b542-7dc673c6c6a8', 'f2aaa171-0118-497b-a0f3-bcffb82533f1', 5, now()),
                                                                          ('5fe06170-4fb3-429a-b950-1ae1a037376e', '38293088-21ab-4f10-9217-eca1f42aaa5d', 26, now()),
                                                                          ('5fe06170-4fb3-429a-b950-1ae1a037376e', '91829929-a842-4a69-a5c5-750dd0074849', 30, now()),
                                                                          ('0cc27199-a03a-4f22-8c43-51ad6f336baf', '91829929-a842-4a69-a5c5-750dd0074849', 9, now()),
                                                                          ('719ece6f-65fc-4cc2-b542-7dc673c6c6a8', '00ef538f-02c2-4a23-8695-385918026262', 43, now()),
                                                                          ('719ece6f-65fc-4cc2-b542-7dc673c6c6a8', '36011ed2-2ae7-4ee9-8567-5900b7b230e3', 12, now());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM items_stocks;
DELETE FROM stocks;
DELETE FROM items;
-- +goose StatementEnd
