import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './AdminPanel.css';

const AdminPanel = () => {
    const [actividades, setActividades] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        const isAdmin = localStorage.getItem("isAdmin") === "true";
        if (!isAdmin) {
            navigate('/');
            return;
        }
        fetchActividades();
    }, [navigate]);

    const fetchActividades = async () => {
        try {
            const response = await fetch('http://localhost:8080/actividades');
            if (response.ok) {
                const data = await response.json();
                setActividades(data);
            }
        } catch (error) {
            console.error("Error al cargar actividades:", error);
        }
    };

    const handleEditar = (actividad) => {
        // TODO: Implementar edición
        console.log("Editar actividad:", actividad);
    };

    const handleEliminar = async (actividad) => {
        if (!actividad.id) {
            console.error("Error: La actividad no tiene ID", actividad);
            alert('Error: No se puede eliminar la actividad porque no tiene ID');
            return;
        }

        if (window.confirm('¿Estás seguro de que deseas eliminar esta actividad?')) {
            try {
                console.log("Intentando eliminar actividad con ID:", actividad.id); // Para debug
                const response = await fetch(`http://localhost:8080/actividades/${actividad.id}`, {
                    method: 'DELETE',
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('access_token')}`,
                        'Content-Type': 'application/json'
                    }
                });

                if (response.ok) {
                    // Actualizar la lista de actividades
                    fetchActividades();
                    alert('Actividad eliminada con éxito');
                } else {
                    const errorData = await response.json().catch(() => ({}));
                    alert(errorData.message || 'Error al eliminar la actividad');
                }
            } catch (error) {
                console.error("Error al eliminar:", error);
                alert('Error al eliminar la actividad');
            }
        }
    };

    return (
        <div className="admin-container">
            <h2>Panel de Administración</h2>
            <div className="admin-table-container">
                <table className="admin-table">
                    <thead>
                        <tr>
                            <th>Título</th>
                            <th>Descripción</th>
                            <th>Instructor</th>
                            <th>Categoría</th>
                            <th>Día</th>
                            <th>Horario</th>
                            <th>Cupo</th>
                            <th>Acciones</th>
                        </tr>
                    </thead>
                    <tbody>
                        {actividades.map((actividad) => (
                            <tr key={actividad.id}>
                                <td>{actividad.titulo}</td>
                                <td>{actividad.descripcion}</td>
                                <td>{actividad.instructor}</td>
                                <td>{actividad.categoria}</td>
                                <td>{actividad.dia}</td>
                                <td>{actividad.hora_inicio} - {actividad.hora_fin}</td>
                                <td>{actividad.cupo}</td>
                                <td className="acciones-column">
                                    <button 
                                        className="action-button edit-button"
                                        onClick={() => handleEditar(actividad)}
                                        title="Editar"
                                    >
                                        ✏️
                                    </button>
                                    <button 
                                        className="action-button delete-button"
                                        onClick={() => handleEliminar(actividad)}
                                        title="Eliminar"
                                    >
                                        🗑️
                                    </button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default AdminPanel; 