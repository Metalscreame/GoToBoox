CREATE SCHEMA gotoboox;

CREATE TABLE gotoboox.tags(
id SERIAL primary key,
title character varying (150) not null
);
CREATE TABLE gotoboox.books(
id SERIAL PRIMARY KEY,
title CHARACTER VARYING (250) NOT NULL,
description CHARACTER VARYING (999) NOT NULL,
popularity REAL NOT NULL DEFAULT 0,
evaluation_number INT DEFAULT 0,
state TEXT DEFAULT 'FREE',
image BYTEA
);

CREATE TABLE gotoboox.books_tags(
book_id int references gotoboox.books (id) not null,
tag_id int references gotoboox.tags (id) not null
);


CREATE TABLE gotoboox.users (
id SERIAL PRIMARY KEY,
nickname CHARACTER VARYING (150) NOT NULL UNIQUE,
email CHARACTER VARYING (250) NOT NULL UNIQUE ,
password CHARACTER VARYING (250) NOT NULL,
exchanges_number INT DEFAULT 0,
book_id  INT REFERENCES  gotoboox.books  (id),
notification_get_new_books BOOLEAN DEFAULT FALSE,
notification_get_when_book_reserved BOOLEAN DEFAULT FALSE,
notification_daily BOOLEAN DEFAULT FALSE,
has_book_for_exchange BOOLEAN DEFAULT FALSE,
returning_book_id INT REFERENCES gotoboox.books (id) DEFAULT NULL,
taken_book_id INT REFERENCES gotoboox.books (id) DEFAULT NULL,
register_date DATE
);

CREATE TABLE gotoboox.categories(
id SERIAL PRIMARY KEY,
title CHARACTER VARYING (250) NOT NULL UNIQUE
);

CREATE TABLE gotoboox.comments (
  id SERIAL PRIMARY KEY,
  book_id INT REFERENCES gotoboox.books (id),
  user_nickname CHARACTER VARYING (150) REFERENCES gotoboox.users (nickname),
  user_email CHARACTER VARYING (250) REFERENCES gotoboox.users (email) ,
  commentary TEXT NOT NULL,
  commentary_date DATE
);


CREATE TABLE gotoboox.roles(
role_id SERIAL primary key,
title character varying (150) not null
);

CREATE TABLE gotoboox.users_roles(
id int references gotoboox.users (id) not null,
role_id int references gotoboox.roles (role_id) not null
);

INSERT INTO gotoboox.users  (nickname,email,password,register_date)
VALUES
('admin', 'admin@admin', 'asd', '2018-01-01'),
('nick2', 'nick2@gmail.com', 'pass2', '2018-01-21'),
('nick3', 'nick3@gmail.com', 'pass3', '2018-01-11'),
('nick4', 'nick4@gmail.com', 'pass4', '2018-01-19'),
('nick5', 'nick5@gmail.com', 'pass5', '2018-03-30'),
('nick6', 'nick6@gmail.com', 'pass6', '2018-02-08');

INSERT INTO gotoboox.categories (title)  VALUES('Romance'),('Fantasy'),('Science');

INSERT INTO gotoboox.books (title, description, popularity,evaluation_number)
VALUES
('Fool for Love', 'Joe Cantrell, owner of the Gansett Island Ferry Company, has been in love with Janey McCarthy for as long as he can remember. At the same time, Janey has been dating or engaged to doctor-in-training David Lawrence. When things go horribly wrong between David and Janey, she calls her fifth brother', '9', '1'),
('Fatal Invasion', 'This compelling romantic suspense novel has all the right elements to keep the reader turning pages, whether engaged in the seamy details of the case or the steamy elements of Sam''s relationship with her hot green-eyed husband. Marie Force, a New York Times bestselling author, excels at creating living, breathing characters and tangling them up in a believable, compelling plot. Fatal Threat is the eleventh book in Force''s Fatal series of romantic thrillers, and Force shows no signs of slowing down.', '4', '1'),
('The Fellowship of the Ring: Being the First Part of The Lord of the Rings', 'One Ring to rule them all, One Ring to find them, One Ring to bring them all and in the darkness bind them. In ancient times the Rings of Power were crafted by the Elven-smiths, and Sauron, the Dark Lord, forged the One Ring, filling it with his own power so that he could rule all others. But the One Ring was taken from him, and though he sought it throughout Middle-earth, it remained lost to him. After many ages it fell into the hands of Bilbo Baggins, as told in The Hobbit. In a sleepy village in the Shire, young Frodo Baggins finds himself faced with an immense task, as his elderly cousin Bilbo entrusts the Ring to his care. Frodo must leave his home and make a perilous journey across Middle-earth to the Cracks of Doom, there to destroy the Ring and foil the Dark Lord in his evil purpose.', '9.5', '2'),
('A Storm of Swords', 'An immersive entertainment experience unlike any other, A Song of Ice and Fire has earned George R. R. Martindubbed the American Tolkien by Time magazine international acclaim and millions of loyal readers. Now here is the entire monumental cycle:','7','1'),
('Big History: Examines Our Past, Explains Our Present, Imagines Our Future', 'Featuring a foreword by the father of Big History, David Christian, and produced in association with the Big History Institute, Big History provides a comprehensive understanding of the major events that have changed the nature and course of life on the planet we call home. This first fully integrated visual reference on Big History for general readers places humans in the context of our universe, from the Big Bang to virtual reality.', '6', '3'),
('Big History: Between Nothing and Everything', 'Big History: Between Nothing and Everything surveys the past not just of humanity, or even of planet Earth, but of the entire universe. In reading this book instructors and students will retrace a voyage that began 13.7 billion years ago with the Big Bang and the appearance of the universe. Big history incorporates findings from cosmology, earth and life sciences, and human history, and assembles them into a single, universal historical narrative of our universe and of our place within it.', '8', '3');

INSERT INTO gotoboox.books (title, description, popularity,evaluation_number,state)
VALUES
  ('Feel for Love', 'Joe Cantrell, owner of the Gansett Island Ferry Company, has been in love with Janey McCarthy for as long as he can remember. At the same time, Janey has been dating or engaged to doctor-in-training David Lawrence. When things go horribly wrong between David and Janey, she calls her fifth brother', '9', '1','TAKEN'),
  ('Tekken taken game', 'Joe Cantrell, owner of the Gansett Island Ferry Company, has been in love with Janey McCarthy for as long as he can remember. At the same time, Janey has been dating or engaged to doctor-in-training David Lawrence. When things go horribly wrong between David and Janey, she calls her fifth brother', '9', '1','TAKEN'),
  ('What a pretty bookkie', 'Joe Cantrell, owner of the Gansett Island Ferry Company, has been in love with Janey McCarthy for as long as he can remember. At the same time, Janey has been dating or engaged to doctor-in-training David Lawrence. When things go horribly wrong between David and Janey, she calls her fifth brother', '9', '1','RESERVED');


INSERT INTO gotoboox.tags(
  title)
VALUES
  ('Crime'), ('Comedy'),('DarkFantasy'),
  ('HistFic'),('Historical'),('LitFic'),('Mystery'),('NonFiction'),('Poetry'),('Romantic'),
  ('RomanticSuspense'), ('TrueStories'),('ScienceFiction'),('ShortStory'),('Suspense'),('UrbanFantasy'),
  ('YA'),('YALit'),('YABook'),('BestBookEver'),('Dystopian'),
  ('EpicReads'),('MyFav'),('NeededToBeRead');

INSERT INTO gotoboox.books_tags(
  id, tag_id)
VALUES ('1', '1'), ('1', '2'), ('1', '3'), ('2', '4'), ('3', '5'), ('4', '1'), ('5', '2'), ('6', '3'),
  ('7', '15'), ('8', '1'), ('5', '11'), ('5', '16'),('5', '5') ;

INSERT INTO gotoboox.roles VALUES (1,'User');