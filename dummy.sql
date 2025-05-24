-- Usuarios
INSERT INTO usuarios (id, nombre, apellido, username, password, rol, fecha_registro) VALUES
  (5, 'Ana',    'Gómez',    'ana.gomez',    'pass123', 'socio', '2025-01-15 09:30:00'),
  (6, 'Bruno',  'Pérez',    'bruno.pere',   'abc321',   'socio', '2025-02-01 14:20:00'),
  (7, 'Carla',  'López',    'carla.lopez',  'xyz789',   'admin', '2025-03-10 08:00:00'),
  (8, 'David',  'Santos',   'david.santos', 'qwe456',   'socio', '2025-04-05 17:45:00'),
  (9, 'Elena',  'Martín',   'elena.mart',   'zxc852',   'socio', '2025-05-01 12:10:00')
;

UPDATE usuarios SET password = sha2(password, 256) WHERE id BETWEEN 5 AND 9;

-- Actividades
INSERT INTO actividads (id, titulo, descripcion, cupo, dia, horario_inicio, horario_final, instructor, categoria) VALUES
  (1, 'Yoga Suave',      'Clase de yoga para principiantes', 15, 'Lunes',    '2025-06-02 10:00:00', '2025-06-02 11:00:00', 'Laura Ruiz',     'yoga'),
  (2, 'Crossfit Básico', 'Entrenamiento funcional de fuerza', 10, 'Miércoles','2025-06-04 18:30:00', '2025-06-04 19:30:00', 'Martín Díaz',    'fitness'),
  (3, 'Natación Adultos','Lecciones de natación nivel intermedio', 8, 'Viernes', '2025-06-06 09:00:00', '2025-06-06 10:00:00', 'Sandra Pérez',   'natación'),
  (4, 'Pilates',        'Pilates para tonificar el core',   12, 'Martes',   '2025-06-03 17:00:00', '2025-06-03 18:00:00', 'Carlos Méndez',  'pilates'),
  (5, 'Spinning',       'Clase de ciclismo indoor de alta intensidad', 20, 'Jueves', '2025-06-05 19:00:00', '2025-06-05 20:00:00', 'Lucía Herrera',  'ciclismo')
;

-- Inscripciones
INSERT INTO inscripciones (usuario_id, actividad_id) VALUES
  (9, 1),
  (9, 4),
  (8, 2),
  (8, 5),
  (5, 1),
  (5, 2),
  (5, 3),
  (6, 3),
  (6, 5)
;

insert into inscripcions (usuario_id, actividad_id) select * from inscripciones;
update inscripcions set estado_inscripcion = 'inactiva' where id in (5,6,9);
