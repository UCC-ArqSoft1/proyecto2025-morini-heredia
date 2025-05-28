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
            <div className="header-container"> 
                <nav className="header-content">
                    <h1 className="header-title">GymPro</h1> 
                    <div className="header-links"> 
                        <a href="/">Inicio</a>
                        <a href="/actividades">Actividades</a>
                        {isLoggedIn ? (
                            <button onClick={logout}>Cerrar sesión</button>
                        ) : (
                            <a href="/login">Iniciar Sesión</a>
                        )}
                    </div>
                </nav>
            </div>
        </header>
    );
}

export default Header;