import { Fragment, useState } from "react";
import './Login.css';
import { useNavigate } from "react-router-dom";

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handlerLogin = async (e) => {
        e.preventDefault();

        if (username === "admin" && password === "admin") {
            console.log("Login OK");
            localStorage.setItem("isLoggedIn", "true");
            navigate("/actividades");
        } else {
            alert("Usuario o contraseña incorrectos."); 
        }
    };

    return (
        <div className="login-container">
            <form className="login-form" onSubmit={handlerLogin}>
                <h2>Iniciar Sesión</h2>
                <div className="input-group"> 
                    <input
                        type="text"
                        placeholder="Usuario"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        required
                    />
                </div>
                <div className="input-group"> 
                    <input
                        type="password"
                        placeholder="Contraseña"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                <button type="submit">Ingresar</button>
            </form>
        </div>
    );
};

export default Login;