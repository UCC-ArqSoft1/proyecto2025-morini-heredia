import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import './EditarActividad.css';

const EditarActividad = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [actividad, setActividad] = useState({
        titulo: '',
        descripcion: '',
        cupo: 0,
        dia: '',
        hora_inicio: '',
        hora_fin: '',
        instructor: '',
        categoria: ''
    });
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchActividad = async () => {
            try {
                const response = await fetch(`http://localhost:8080/actividades/${id}`);
                if (response.ok) {
                    const data = await response.json();
                    setActividad(data);
                } else {
                    setError('Error al cargar la actividad');
                }
            } catch (error) {
                setError('Error al conectar con el servidor');
            }
        };

        fetchActividad();
    }, [id]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await fetch(`http://localhost:8080/actividades/${id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('access_token')}`
                },
                body: JSON.stringify(actividad)
            });

            if (response.ok) {
                alert('Actividad actualizada correctamente');
                navigate('/admin');
            } else {
                const errorData = await response.json();
                setError(errorData.error || 'Error al actualizar la actividad');
            }
        } catch (error) {
            setError('Error al conectar con el servidor');
        }
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setActividad(prev => ({
            ...prev,
            [name]: value
        }));
    };

    return (
        <div className="editar-actividad-container">
            <h2>Editar Actividad</h2>
            {error && <div className="error-message">{error}</div>}
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="titulo">Título:</label>
                    <input
                        type="text"
                        id="titulo"
                        name="titulo"
                        value={actividad.titulo}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="descripcion">Descripción:</label>
                    <textarea
                        id="descripcion"
                        name="descripcion"
                        value={actividad.descripcion}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="cupo">Cupo:</label>
                    <input
                        type="number"
                        id="cupo"
                        name="cupo"
                        value={actividad.cupo}
                        onChange={handleChange}
                        required
                        min="1"
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="dia">Día:</label>
                    <select
                        id="dia"
                        name="dia"
                        value={actividad.dia}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Seleccione un día</option>
                        <option value="Lunes">Lunes</option>
                        <option value="Martes">Martes</option>
                        <option value="Miércoles">Miércoles</option>
                        <option value="Jueves">Jueves</option>
                        <option value="Viernes">Viernes</option>
                        <option value="Sábado">Sábado</option>
                    </select>
                </div>

                <div className="form-group">
                    <label htmlFor="hora_inicio">Hora de inicio:</label>
                    <input
                        type="time"
                        id="hora_inicio"
                        name="hora_inicio"
                        value={actividad.hora_inicio}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="hora_fin">Hora de fin:</label>
                    <input
                        type="time"
                        id="hora_fin"
                        name="hora_fin"
                        value={actividad.hora_fin}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="instructor">Instructor:</label>
                    <input
                        type="text"
                        id="instructor"
                        name="instructor"
                        value={actividad.instructor}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="categoria">Categoría:</label>
                    <select
                        id="categoria"
                        name="categoria"
                        value={actividad.categoria}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Seleccione una categoría</option>
                        <option value="musculacion">Musculación</option>
                        <option value="cardio">Cardio</option>
                        <option value="yoga">Yoga</option>
                        <option value="baile">Baile</option>
                        <option value="funcional">Funcional</option>
                    </select>
                </div>

                <div className="form-buttons">
                    <button type="submit" className="btn-guardar">Guardar Cambios</button>
                    <button type="button" className="btn-cancelar" onClick={() => navigate('/admin')}>
                        Cancelar
                    </button>
                </div>
            </form>
        </div>
    );
};

export default EditarActividad; 