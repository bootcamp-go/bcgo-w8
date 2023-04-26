USE movies_db;

/* ---------------------- */
/* JOIN - GROUPBY- HAVING */
/* ---------------------- */

/* Mostrar aquellas peliculas favoritas de los actores, y cuantos actores la consideran como tal */
SELECT COUNT(*) as liked_by_actors, mo.title, mo.rating, mo.awards FROM movies as mo INNER JOIN actors ac ON mo.id = ac.favorite_movie_id
GROUP BY mo.title HAVING liked_by_actors > 1
ORDER BY liked_by_actors DESC, mo.title ASC;

/* ---------------------- */
/* SUB-QUERIES         	  */
/* ---------------------- */

/* Mostrar info de las peliculas de los actores, cuyo rating fue mayor a 9 */
SELECT am.* FROM actor_movie as am
WHERE movie_id IN (SELECT id FROM movies WHERE rating=9.0);

SELECT * FROM actor_movie as am
WHERE EXISTS (
  SELECT * FROM movies m
  WHERE m.id = am.movie_id AND m.rating = 9.0
);

SELECT am.* FROM actor_movie as am INNER JOIN movies as mo ON am.movie_id = mo.id AND mo.rating = 9.0;