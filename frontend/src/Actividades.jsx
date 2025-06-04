import React, { useState, useEffect } from "react";
import "./Actividades.css";

const Actividades = () => {
  const [actividades, setActividades] = useState([]);
  const [inscripciones, setInscripciones] = useState([]);
  const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
  const isAdmin = localStorage.getItem("isAdmin") === "true";

  useEffect(() => {
    fetchActividades();
    const savedInscripciones = localStorage.getItem("inscripciones");
    if (savedInscripciones) {
      setInscripciones(JSON.parse(savedInscripciones));
    }
  }, []);

  const fetchActividades = async () => {
    try {
      const response = await fetch("http://localhost:8080/actividades");
      if (response.ok) {
        const data = await response.json();
        console.log("Actividades cargadas:", data); // Para debug
        setActividades(data);
      }
    } catch (error) {
      console.error("Error al cargar actividades:", error);
    }
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
      {actividades.map((actividad) => (
        <div className="actividad-card" key={actividad.id}>
          <h3>{actividad.titulo}</h3>
          <p>{actividad.descripcion}</p>
          <p>Instructor: {actividad.instructor || "No especificado"}</p>
          <p>Categoría: {actividad.categoria || "No especificada"}</p>
          <p>
            Día: {actividad.dia || "No especificado"} - Horario:{" "}
            {actividad.hora_inicio} a {actividad.hora_fin}
          </p>
          <p>Cupo disponible: {actividad.cupo || "No especificado"}</p>
          
          {isLoggedIn && (
            <div className="card-actions">
              {isAdmin ? (
                <>
                  <button 
                    className="action-button edit-button"
                    onClick={() => handleEditar(actividad)}
                    title="Editar"
                  >
                    ✏️ Editar
                  </button>
                  <button 
                    className="action-button delete-button"
                    onClick={() => handleEliminar(actividad)}
                    title="Eliminar"
                  >
                    🗑️ Eliminar
                  </button>
                </>
              ) : (
                <button onClick={() => handleInscription(actividad.titulo)}>
                  {estaInscripto(actividad.titulo) ? "Desinscribir" : "Inscribir"}
                </button>
              )}
            </div>
          )}
        </div>
      ))}
    </div>
  );
};

export default Actividades;