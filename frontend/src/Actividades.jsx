import React, { useState, useEffect } from "react";
import EditarActividadModal from './EditarActividadModal';
import "./Actividades.css";

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

  useEffect(() => {
    fetchActividades();
    const savedInscripciones = localStorage.getItem("inscripciones");
    if (savedInscripciones) {
      setInscripciones(JSON.parse(savedInscripciones));
    }
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

    // Filtrar por categoría
    if (filtros.categoria) {
      actividadesFiltradas = actividadesFiltradas.filter(actividad => 
        actividad.categoria.toLowerCase() === filtros.categoria.toLowerCase()
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

  const handleInscription = (nombreActividad) => {
    const isInscripto = inscripciones.includes(nombreActividad);
    
    let nuevasInscripciones;
    if (isInscripto) {
      nuevasInscripciones = inscripciones.filter(nombre => nombre !== nombreActividad);
      alert(`Desinscripto de ${nombreActividad}`);
    } else {
      nuevasInscripciones = [...inscripciones, nombreActividad];
      alert(`Inscripto en ${nombreActividad}`);
    }
    
    setInscripciones(nuevasInscripciones);
    localStorage.setItem("inscripciones", JSON.stringify(nuevasInscripciones));
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
    if (!actividad.id) {
      console.error("Error: La actividad no tiene ID", actividad);
      alert('Error: No se puede eliminar la actividad porque no tiene ID');
      return;
    }

    if (window.confirm('¿Estás seguro de que deseas eliminar esta actividad?')) {
      try {
        console.log("Intentando eliminar actividad con ID:", actividad.id);
        const response = await fetch(`http://localhost:8080/actividades/${actividad.id}`, {
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

  const estaInscripto = (nombreActividad) => {
    return inscripciones.includes(nombreActividad);
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
        <select
          name="categoria"
          value={filtros.categoria}
          onChange={handleFiltroChange}
          className="filtro-select"
        >
          <option value="">Categoría</option>
          <option value="funcional">Funcional</option>
          <option value="spinning">Spinning</option>
          <option value="yoga">Yoga</option>
          <option value="pilates">Pilates</option>
          <option value="mma">MMA</option>
        </select>
        <select
          name="dia"
          value={filtros.dia}
          onChange={handleFiltroChange}
          className="filtro-select"
        >
          <option value="">Día</option>
          <option value="lunes">Lunes</option>
          <option value="martes">Martes</option>
          <option value="miercoles">Miércoles</option>
          <option value="jueves">Jueves</option>
          <option value="viernes">Viernes</option>
          <option value="sabado">Sábado</option>
        </select>
      </div>

      <div className="actividades-grid">
        {actividadesFiltradas.map((actividad) => (
          <div className="actividad-card" key={actividad.id}>
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
              <p>Cupo disponible: {actividad.cupo || "No especificado"}</p>
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
                    onClick={() => handleInscription(actividad.titulo)}
                  >
                    {estaInscripto(actividad.titulo) ? "Desinscribir" : "Inscribir"}
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