DELETE FROM schedules;
INSERT INTO schools VALUE (NULL, "Studio Landin'", now(), now());

INSERT INTO rooms VALUE (NULL, "A", 50, @studio_id, now(), now());
SELECT id INTO @room_id FROM rooms WHERE studio_id = @studio_id AND name = "A";