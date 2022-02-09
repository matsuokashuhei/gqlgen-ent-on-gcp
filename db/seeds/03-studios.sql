DELETE FROM studios;
SELECT id INTO @school_id FROM schools WHERE name = "Studio Landin'";

INSERT INTO studios (name, location, school_studios, create_time, update_time) VALUE ("立川校", "東京都立川市", @school_id, now(), now()),
                                                                                     ("国分寺校", "東京都国分寺市", @school_id, now(), now());