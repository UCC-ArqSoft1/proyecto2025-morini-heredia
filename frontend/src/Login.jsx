import { useState } from "react";
import './Login.css';
import { useNavigate } from "react-router-dom";

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handlerLogin = async (e) => {
        e.preventDefault();
        setIsLoading(true);
        setError("");

        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    username: username.trim(),
                    password: password
                })
            });

            if (response.ok) {
                const data = await response.json();
                
                localStorage.setItem("access_token", data.access_token);
                localStorage.setItem("token_type", data.token_type);
                localStorage.setItem("isLoggedIn", "true");
                
                navigate("/actividades");
            } else {
                const errorData = await response.json();
                setError(errorData.error || "Error de autenticación");
                alert("Error al loguearse");
            }

        } catch (error) {
            setError("Error de conexión");
            console.error("Error de conexión:", error);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="login-container">
            <form className="login-form" onSubmit={handlerLogin}>
                <h2>Iniciar Sesión</h2>
                
                {error && <div className="error-message">{error}</div>}
                
                <div className="input-group">
                    <input
                        type="text"
                        placeholder="Usuario"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        disabled={isLoading}
                        required
                    />
                </div>
                
                <div className="input-group">
                    <input
                        type="password"
                        placeholder="Contraseña"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        disabled={isLoading}
                        required
                    />
                </div>
                
                <button type="submit" disabled={isLoading}>
                    {isLoading ? "Ingresando..." : "Ingresar"}
                </button>
            </form>
        </div>
    );
};

export default Login;