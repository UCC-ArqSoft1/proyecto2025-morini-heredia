import React, { useState, useEffect } from "react";
import EditarActividadModal from './EditarActividadModal';
import "./Actividades.css";
import { useNavigate } from "react-router-dom";

const Actividades = () => {
    const [actividades, setActividades] = useState([]);
    const [actividadesFiltradas, setActividadesFiltradas] = useState([]);
    const [inscripciones, setInscripciones] = useState([]);
    const [actividadEditar, setActividadEditar] = useState(null);
    const [filtros, setFiltros] = useState({
        busqueda: "",
        categoria: "",
        dia: ""
    });
    const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
    const isAdmin = localStorage.getItem("isAdmin") === "true";
    const navigate = useNavigate();
    // const idUsuario = parseInt(localStorage.getItem("idUsuario"))

    useEffect(() => {
        fetchActividades();
        fetchInscripciones();
    }, []);

    useEffect(() => {
        filtrarActividades();
    }, [filtros, actividades]);

    const fetchActividades = async () => {
        try {
            const response = await fetch("http://localhost:8080/actividades");
            if (response.ok) {
                const data = await response.json();
                console.log("Actividades cargadas:", data);
                setActividades(data);
                setActividadesFiltradas(data);
            }
        } catch (error) {
            console.error("Error al cargar actividades:", error);
        }
    };
    
    const fetchInscripciones = async () => {
        try {
            const response = await fetch("http://localhost:8080/inscripciones", {
                headers: {'Authorization': `Bearer ${localStorage.getItem('access_token')}`
            },
            });
            if (response.ok) {
                const resp = await response.json();
                const data = resp.filter(insc => insc.is_activa)
                
                console.log("Inscripciones cargadas:", data);
                setInscripciones(data);
            }
        } catch (error) {
            console.error("Error al cargar inscripciones:", error);
        }
    };

    const handleFiltroChange = (e) => {
        const { name, value } = e.target;
        setFiltros(prev => ({
            ...prev,
            [name]: value
        }));
    };

    const filtrarActividades = () => {
        let actividadesFiltradas = [...actividades];

        // Filtrar por búsqueda (título o descripción)
        if (filtros.busqueda) {
            const busquedaLower = filtros.busqueda.toLowerCase();
            actividadesFiltradas = actividadesFiltradas.filter(actividad =>
                actividad.titulo.toLowerCase().includes(busquedaLower) ||
                actividad.descripcion.toLowerCase().includes(busquedaLower)
            );
        }

        // Filtrar por categoría (ahora como búsqueda de texto)
        if (filtros.categoria) {
            const categoriaLower = filtros.categoria.toLowerCase();
            actividadesFiltradas = actividadesFiltradas.filter(actividad =>
                actividad.categoria.toLowerCase().includes(categoriaLower)
            );
        }

        // Filtrar por día
        if (filtros.dia) {
            actividadesFiltradas = actividadesFiltradas.filter(actividad =>
                actividad.dia.toLowerCase() === filtros.dia.toLowerCase()
            );
        }

        setActividadesFiltradas(actividadesFiltradas);
    };

    const handleEnroling = async (actividadId) => {
        if (!isLoggedIn) {
            navigate("/login");
            return;
        }

        try {
            const response = await fetch("http://localhost:8080/inscripciones", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${localStorage.getItem("access_token")}`,
                },
                body: JSON.stringify({
                    id: actividadId,
                }),
            });

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.error || "Error al inscribirse en la actividad");
            }

            // Actualizar la lista de inscripciones
            fetchInscripciones();
            // Actualizar la lista de actividades para reflejar el cambio en los cupos
            fetchActividades();
            alert("¡Inscripción exitosa!");
        } catch (error) {
            console.error("Error al inscribirse:", error);
            alert(error.message);
        }
    };
    
    const handleUnenrolling = async (id_actividad) => {
        try {
            const response = await fetch("http://localhost:8080/inscripciones", {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('access_token')}`,
                },
                body: JSON.stringify({
                    id: parseInt(id_actividad)
                })
            });

            if (response.status == 204) {
                alert(`Desinscripto exitosamente`);
                fetchInscripciones();
            } else {
                alert(`Ups! algo salio mal, vuelve a intentarlo mas tarde`);
            }

            fetchActividades();
        } catch (error) {
            alert(`Ups! algo salio mal, vuelve a intentarlo mas tarde`);
            console.error("Error al desinscribir el usuario:", error);
        }
    };

    const handleEditar = (actividad) => {
        setActividadEditar(actividad);
    };

    const handleCloseModal = () => {
        setActividadEditar(null);
    };

    const handleSaveEdit = () => {
        fetchActividades();
    };

    const handleEliminar = async (actividad) => {
        if (!actividad.id_actividad) {
            console.error("Error: La actividad no tiene ID", actividad);
            alert('Error: No se puede eliminar la actividad porque no tiene ID');
            return;
        }

        if (window.confirm('¿Estás seguro de que deseas eliminar esta actividad?')) {
            try {
                console.log("Intentando eliminar actividad con ID:", actividad.id_actividad);
                const response = await fetch(`http://localhost:8080/actividades/${actividad.id_actividad}`, {
                    method: 'DELETE',
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('access_token')}`,
                        'Content-Type': 'application/json'
                    }
                });

                if (response.ok) {
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

    const estaInscripto = (id_actividad) => {
        return inscripciones.some(insc => 
            insc.id_actividad === id_actividad &&
            insc.is_activa
        )
    };

    return (
        <div className="actividades-container">
            <div className="filtros-container">
                <div className="search-wrapper">
                    <span className="search-icon">🔍</span>
                    <input
                        type="text"
                        name="busqueda"
                        placeholder="Buscar actividad..."
                        value={filtros.busqueda}
                        onChange={handleFiltroChange}
                        className="filtro-input"
                    />
                </div>
                <input
                    type="text"
                    name="categoria"
                    placeholder="Categoría..."
                    value={filtros.categoria}
                    onChange={handleFiltroChange}
                    className="filtro-input"
                />
                <select
                    name="dia"
                    value={filtros.dia}
                    onChange={handleFiltroChange}
                    className="filtro-select"
                >
                    <option value="">Día</option>
                    <option value="Lunes">Lunes</option>
                    <option value="Martes">Martes</option>
                    <option value="Miercoles">Miercoles</option>
                    <option value="Jueves">Jueves</option>
                    <option value="Viernes">Viernes</option>
                    <option value="Sabado">Sabado</option>
                </select>
            </div>

            <div className="actividades-grid">
                {actividadesFiltradas.map((actividad) => (
                    <div className="actividad-card" key={actividad.id_actividad}>
                        <h3>{actividad.titulo}</h3>
                        <div className="actividad-info">
                            <p>{actividad.descripcion}</p>
                            <p>Instructor: {actividad.instructor || "No especificado"}</p>
                            <p>Categoría: {actividad.categoria || "No especificada"}</p>
                            <p>
                                Día: {actividad.dia || "No especificado"}
                                <span className="actividad-horario">
                                    Horario: {actividad.hora_inicio} a {actividad.hora_fin}
                                </span>
                            </p>
                            <p>Cupo total: {actividad.cupo} | Lugares disponibles: {actividad.lugares}</p>
                        </div>

                        {isLoggedIn && (
                            <div className="card-actions">
                                {isAdmin ? (
                                    <>
                                        <button
                                            className="action-button edit-button"
                                            onClick={() => handleEditar(actividad)}
                                            title="Editar"
                                        >
                                            <span>✏️</span>
                                            Editar
                                        </button>
                                        <button
                                            className="action-button delete-button"
                                            onClick={() => handleEliminar(actividad)}
                                            title="Eliminar"
                                        >
                                            <span>🗑️</span>
                                            Eliminar
                                        </button>
                                    </>
                                ) : (
                                    <button
                                        className="inscripcion-button"
                                        onClick={() => 
                                            estaInscripto(actividad.id_actividad) ? 
                                                handleUnenrolling(actividad.id_actividad) :
                                                handleEnroling(actividad.id_actividad)
                                        }
                                    >
                                        {estaInscripto(actividad.id_actividad) ? "Desinscribir ❌" : "Inscribir ✔️"}
                                    </button>
                                )}
                            </div>
                        )}
                    </div>
                ))}
            </div>

            {actividadEditar && (
                <EditarActividadModal
                    actividad={actividadEditar}
                    onClose={handleCloseModal}
                    onSave={handleSaveEdit}
                />
            )}
        </div>
    );
};

export default Actividades;
