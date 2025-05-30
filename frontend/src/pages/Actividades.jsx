import React, { useState, useEffect } from "react";
import "./Actividades.css";

const Actividades = () => {
  const [actividades, setActividades] = useState([]);
  const [inscripciones, setInscripciones] = useState([]);
  const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";

  useEffect(() => {
    fetch("http://localhost:8080/actividades")
      .then((res) => res.json())
      .then((data) => setActividades(data))
      .catch((err) => console.error("Error fetching actividades:", err));
    
    // Cargar inscripciones guardadas en localStorage
    const savedInscripciones = localStorage.getItem("inscripciones");
    if (savedInscripciones) {
      setInscripciones(JSON.parse(savedInscripciones));
    }
  }, []);

  const handleInscription = (nombreActividad) => {
    // Verificar si ya está inscripto
    const isInscripto = inscripciones.includes(nombreActividad);
    
    let nuevasInscripciones;
    if (isInscripto) {
      // Desinscribir
      nuevasInscripciones = inscripciones.filter(nombre => nombre !== nombreActividad);
      alert(`Desinscripto de ${nombreActividad}`);
    } else {
      // Inscribir
      nuevasInscripciones = [...inscripciones, nombreActividad];
      alert(`Inscripto en ${nombreActividad}`);
    }
    
    // Actualizar estado y localStorage
    setInscripciones(nuevasInscripciones);
    localStorage.setItem("inscripciones", JSON.stringify(nuevasInscripciones));
  };

  const handleDetails = (nombreActividad) => {
    console.log(`Detalles para ${nombreActividad} (pendiente de implementar)`);
  };

  const estaInscripto = (nombreActividad) => {
    return inscripciones.includes(nombreActividad);
  };

  return (
    <div className="actividades-container">
      {actividades.map((actividad, index) => (
        <div className="actividad-card" key={index}>
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
              <button onClick={() => handleInscription(actividad.titulo)}>
                {estaInscripto(actividad.titulo) ? "Desinscribir" : "Inscribir"}
              </button>
              <button
                className="detalles-btn"
                onClick={() => handleDetails(actividad.titulo)}
              >
                Detalles
              </button>
            </div>
          )}
        </div>
      ))}
    </div>
  );
};

export default Actividades;