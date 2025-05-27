import { useNavigate } from "react-router-dom";
import "./Header.css";


const Header = ( ) => {
    const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
    const navigate = useNavigate();
    const logout = () => {
        localStorage.removeItem("isLoggedIn");
        navigate("/");
    }
    
    return (
        <header>
            <div className="header-container"> {/* Clase para el header principal */}
                <nav className="header-content"> {/* Clase para el contenido de la navegación */}
                    <h1 className="header-title">GymPro</h1> {/* Clase para el título */}
                    <div className="header-links"> {/* Clase para el contenedor de enlaces */}
                        <a href="/">Inicio</a>
                        <a href="/actividades">Mis Actividades</a>
                        {isLoggedIn ? (
                            <button onClick={logout}>Cerrar sesión</button>
                        ) : (
                            <a href="/login">Iniciar Sesión</a>
                        )}
                    </div>
                </nav>
                {/* Aquí iría tu "mas codigo" si pertenece al header */}
            </div>
        </header>
    );
}

export default Header;