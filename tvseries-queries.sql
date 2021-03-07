SELECT a.name as "Actor", s.name as "Serie" FROM serie_has_actor as sa
    LEFT JOIN actor as a ON a.id_actor = sa.actor_id_actor
    LEFT JOIN serie as s ON s.id_serie = sa.serie_id_serie
    WHERE s.name = 'BIG SISTER'


SELECT COUNT(e.id_director) as "Episodes Directed" FROM episode as e
    LEFT JOIN director as d ON e.director_id_director = d.id_director
    GROUP BY e.director_id_director
    ORDER BY COUNT(e.id_director) DESC
    LIMIT 1;