-- Usuarios
INSERT INTO usuarios (id_usuario, nombre, apellido, username, password, is_admin) VALUES
  (2, 'Ana',    'Gómez',    'ana.gomez',    'pass123', 0),
  (3, 'Bruno',  'Pérez',    'bruno.pere',   'abc321',  0),
  (4, 'Carla',  'López',    'carla.lopez',  'xyz789',  1),
  (5, 'David',  'Santos',   'david.santos', 'qwe456',  0),
  (6, 'Elena',  'Martín',   'elena.mart',   'zxc852',  0)
;

UPDATE usuarios SET password = sha2(password, 256) WHERE id_usuario BETWEEN 2 AND 6;

-- Actividades
INSERT INTO actividads (id_actividad, foto_url, titulo, descripcion, cupo, dia, horario_inicio, horario_final, instructor, categoria) VALUES
  (1, 'SAMPLE_URL', 'Karate Miyagi-do',        'Clase de karate con el estilo del señor Miyagi', 10, 'Sabado',    '07:00:00', '09:00:00', 'Sr. Miyagi',      'karate'),
  (2, 'SAMPLE_URL', 'Yoga Suave',              'Clase de yoga para principiantes',               15, 'Lunes',     '10:00:00', '11:00:00', 'Laura Ruiz',      'yoga'),
  (3, 'SAMPLE_URL', 'Crossfit Básico',         'Entrenamiento funcional de fuerza',              10, 'Miercoles', '18:30:00', '19:30:00', 'Martín Díaz',     'fitness'),
  (4, 'SAMPLE_URL', 'Natación Adultos',        'Lecciones de natación nivel intermedio',         8,  'Viernes',   '09:00:00', '10:00:00', 'Sandra Pérez',    'natación'),
  (5, 'SAMPLE_URL', 'Pilates',                 'Pilates para tonificar el core',                 12, 'Martes',    '17:00:00', '18:00:00', 'Carlos Méndez',   'pilates'),
  (6, 'SAMPLE_URL', 'Spinning',                'Clase de ciclismo indoor de alta intensidad',    20, 'Jueves',    '19:00:00', '20:00:00', 'Lucía Herrera',   'ciclismo'),
  (7, 'SAMPLE_URL', 'Boxeo profesional',       'Clase de boxeo',                                 4,  'Martes',    '19:00:00', '20:00:00', 'Mike Tyson',      'boxeo'),
  (8, 'SAMPLE_URL', 'Entrenamiento de fuerza', 'Entrenamiento con pesas y maquinas',             12, 'Miercoles', '19:00:00', '20:00:00', 'Anatoli Cleaner', 'fuerza'),
  (9, 'SAMPLE_URL', 'Entrenamiento de fuerza', 'Entrenamiento con pesas y maquinas',             12, 'Miercoles', '20:00:00', '21:00:00', 'Anatoli Cleaner', 'fuerza')
;

-- Inscripciones
INSERT INTO inscripcions (id_usuario, id_actividad) VALUES
  (2, 1),
  (2, 7),
  (3, 1),
  (3, 7),
  (6, 1),
  (6, 2),
  (6, 4),
  (6, 5),
  (6, 7)
;

