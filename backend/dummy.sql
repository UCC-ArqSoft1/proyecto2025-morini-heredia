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
INSERT INTO actividads (id_actividad, titulo, descripcion, cupo, dia, horario_inicio, horario_final, instructor, categoria) VALUES
  (1, 'Karate Miyagi-do',       'Clase de karate con el estilo del señor Miyagi', 10, 'Sabado', '2025-06-05 07:00:00', '2025-06-05 09:00:00', 'Sr. Miyagi',  'karate'),
  (2, 'Yoga Suave',      'Clase de yoga para principiantes', 15, 'Lunes',    '2025-06-02 10:00:00', '2025-06-02 11:00:00', 'Laura Ruiz',     'yoga'),
  (3, 'Crossfit Básico', 'Entrenamiento funcional de fuerza', 10, 'Miercoles','2025-06-04 18:30:00', '2025-06-04 19:30:00', 'Martín Díaz',    'fitness'),
  (4, 'Natación Adultos','Lecciones de natación nivel intermedio', 8, 'Viernes', '2025-06-06 09:00:00', '2025-06-06 10:00:00', 'Sandra Pérez',   'natación'),
  (5, 'Pilates',        'Pilates para tonificar el core',   12, 'Martes',   '2025-06-03 17:00:00', '2025-06-03 18:00:00', 'Carlos Méndez',  'pilates'),
  (6, 'Spinning',       'Clase de ciclismo indoor de alta intensidad', 20, 'Jueves', '2025-06-05 19:00:00', '2025-06-05 20:00:00', 'Lucía Herrera',  'ciclismo'),
  (7, 'Boxeo profesional',       'Clase de boxeo', 4, 'Martes', '2025-06-05 19:00:00', '2025-06-05 20:00:00', 'Mike Tyson',  'boxeo')
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

