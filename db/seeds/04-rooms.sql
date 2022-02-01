DELETE FROM rooms;

SELECT id INTO @studio_id FROM studios WHERE name = "立川校";
INSERT INTO rooms VALUE (NULL, "A", 30, @studio_id, now(), now());

SELECT id INTO @studio_id FROM studios WHERE name = "国分寺校";
INSERT INTO rooms VALUE (NULL, "A", 30, @studio_id, now(), now()),
                        (NULL, "B", 10, @studio_id, now(), now());