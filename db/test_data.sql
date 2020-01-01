INSERT INTO watchlist(id, owner_id, name, created_on, items)
VALUES (1, 1, 'My List', '2019-11-20 3:00:00', '{1, 2}');

INSERT INTO item_types(id, name)
VALUES (1, 'Movie');

INSERT INTO item_types(id, name)
VALUES (2, 'TV Show');

INSERT INTO items(watchlist_id, item_type, title,  description, release_date, rating, genre, watched)
VALUES (1, 1, 'John Wick: Chapter 3 - Parabellum', 'John Wick is on the run after killing a member of the international assassins guild, and with a $14 million price tag on his head, he is the target of hit men and women everywhere.', '2019-05-17 12:00:00', 'R', 'Action, Crime, Thriller', false);

INSERT INTO items(watchlist_id, item_type, title,  description, release_date, rating, genre, watched)
VALUES (1, 1, 'Venom', 'A failed reporter is bonded to an alien entity, one of many symbiotes who have invaded Earth. But the being takes a liking to Earth and decides to protect it.', '2018-10-05 12:00:00', 'PG-13', 'Action, Sci-Fi, Thriller', false);