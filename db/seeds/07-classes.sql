DELETE FROM classes;

SELECT id INTO @studio_id FROM studios WHERE name = "立川校";
SELECT id INTO @room_id FROM rooms WHERE studio_rooms = @studio_id AND name = "A";
SELECT id INTO @instructor_id FROM instructors WHERE name = "KU-GE.";

SELECT id INTO @schedule_id FROM schedules WHERE room_schedules = @room_id AND day_of_week = 1 AND start_time = "16:45";
INSERT INTO classes (name, level, tuition, start_date, instructor_classes, schedule_classes) VALUE ("Jr.HIPHOP", "小学生", 6000, "2022-01-01", @instructor_id, @schedule_id);
SELECT id INTO @schedule_id FROM schedules WHERE room_schedules = @room_id AND day_of_week = 1 AND start_time = "18:00";
INSERT INTO classes (name, level, tuition, start_date, instructor_classes, schedule_classes) VALUE ("GirlsHIPHOP", "初級", 7000, "2022-01-01", @instructor_id, @schedule_id);
SELECT id INTO @schedule_id FROM schedules WHERE room_schedules = @room_id AND day_of_week = 1 AND start_time = "19:15";
INSERT INTO classes (name, level, tuition, start_date, instructor_classes, schedule_classes) VALUE ("GirlsHIPHOP", "中級", 6000, "2022-01-01", @instructor_id, @schedule_id);
