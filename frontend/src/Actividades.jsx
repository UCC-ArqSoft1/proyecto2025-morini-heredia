

const Actividades = () => {
    const actividades = [
        {
          nombre: "taekwondo",
          descripcion: "Arte marcial coreana",
          horarios: [
            { dia: 2, "hora-inicio": "18:30", "hora-fin": "20:00" },
            { dia: 4, "hora-inicio": "18:30", "hora-fin": "20:00" }
          ]
        },
        {
          nombre: "zumba",
          descripcion: "ritmos latinos",
          horarios: [
            { dia: 1, "hora-inicio": "19:30", "hora-fin": "20:30" },
            { dia: 3, "hora-inicio": "19:30", "hora-fin": "20:30" }
          ]
        }
      ];

      const diasSemana = ["Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"];

      const handleInscription = (nombreActividad) => {
        alert(`Inscripto en ${nombreActividad}`)
      }

      const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
      
      return (
        <div className="actividades-container">
            <p>Hola mundo</p>
            {actividades.map((actividad, index) => (
                <div className="actividad-card" key={index}>
                    <h3>{ actividad.nombre }</h3>
                    <p> { actividad.descripcion }</p>
                    <ul>
                        { actividad.horarios.map((horario, i) => (
                            <li key={i}>
                                Dia: { diasSemana[horario.dia] } -
                                Hora Inicio: { horario["hora-inicio"]} -
                                Hora Fin: { horario["hora-fin"]}
                            </li>
                        ))}
                    </ul>
                    {
                      isLoggedIn && (
                        <button onClick={() => handleInscription(actividad.nombre)}>Inscribir</button>
                      )
                    }
                </div>
            )
            
        )}
        </div>
      )

}
export default Actividades;
