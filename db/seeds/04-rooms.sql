DELETE FROM rooms;

SELECT id INTO @studio_id FROM studios WHERE name = "立川校";
INSERT INTO rooms (name, capacity, studio_rooms, create_time, update_time) VALUE ("A", 30, @studio_id, now(), now());

SELECT id INTO @studio_id FROM studios WHERE name = "国分寺校";
INSERT INTO rooms (name, capacity, studio_rooms, create_time, update_time) VALUE ("A", 30, @studio_id, now(), now()),
                                                                                 ("B", 10, @studio_id, now(), now());