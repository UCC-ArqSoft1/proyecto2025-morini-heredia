-- Eliminar la restricción de clave foránea existente
ALTER TABLE inscripcions
DROP FOREIGN KEY fk_actividads_inscripciones;

-- Recrear la restricción con ON DELETE CASCADE
ALTER TABLE inscripcions
ADD CONSTRAINT fk_actividads_inscripciones
FOREIGN KEY (id_actividad) 
REFERENCES actividads(id_actividad)
ON DELETE CASCADE;

DROP TRIGGER IF EXISTS revisar_cupo_insert;
DROP TRIGGER IF EXISTS revisar_cupo_update_ins;
DROP TRIGGER IF EXISTS revisar_cupo_update_act;

DELIMITER //

-- revisamos si es posible insertar un registro en inscripcions basandonos en el cupo maximo de la actividad
CREATE TRIGGER revisar_cupo_insert
BEFORE INSERT ON inscripcions
FOR EACH ROW
BEGIN
    DECLARE actividad_cupo INT;

    -- Obtener el cupo de la actividad
    SELECT cupo INTO actividad_cupo
    FROM actividads
    WHERE id_actividad = NEW.id_actividad;

    -- Contar las inscripciones actuales para la actividad
    IF (SELECT COUNT(*) FROM inscripcions WHERE id_actividad = NEW.id_actividad AND is_activa) >= actividad_cupo THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'No se puede inscribir, el cupo de la actividad ha sido alcanzado.';
    END IF;
END;

//

-- en caso de activar la inscripcion mediante UPDATE, revisamos el cupo disponible
CREATE TRIGGER revisar_cupo_update_ins
BEFORE UPDATE ON inscripcions
FOR EACH ROW
BEGIN
    DECLARE actividad_cupo INT;

    -- Verificar si se está activando la inscripción
    IF NEW.is_activa = 1 AND OLD.is_activa = 0 THEN
        -- Obtener el cupo de la actividad
        SELECT cupo INTO actividad_cupo
        FROM actividads
        WHERE id_actividad = NEW.id_actividad;

        -- Contar las inscripciones activas para la actividad
        IF (SELECT COUNT(*) FROM inscripcions WHERE id_actividad = NEW.id_actividad AND is_activa) >= actividad_cupo THEN
            SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'No se puede activar la inscripción, el cupo de la actividad ha sido alcanzado.';
        END IF;
    END IF;
END;

//

-- revisamos si es posible cambiar el cupo de una actividad
CREATE TRIGGER revisar_cupo_update_act
BEFORE UPDATE ON actividads
FOR EACH ROW
BEGIN
    DECLARE inscripciones_activas INT;

    -- Contar las inscripciones activas para la actividad
    SELECT COUNT(*) INTO inscripciones_activas
    FROM inscripcions
    WHERE id_actividad = NEW.id_actividad AND is_activa = 1;

    -- Verificar si el nuevo cupo es menor que las inscripciones activas
    IF NEW.cupo < inscripciones_activas THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'No se puede cambiar el cupo, hay inscripciones activas que superan el nuevo límite.';
    END IF;
END;

//

DELIMITER ;

DROP VIEW IF EXISTS actividads_lugares;

CREATE VIEW actividads_lugares AS
SELECT *, cupo - (SELECT COUNT(*) FROM inscripcions ins WHERE ins.id_actividad = ac.id_actividad AND ins.is_activa) AS lugares
FROM actividads ac;
