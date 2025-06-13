import { useState } from "react";
import '../styles/Login.css';
import { useNavigate } from "react-router-dom";

const getTokenPayload = (token) => {
    const parts = token.split('.');
    const decodedPaylod = atob(parts[1]);

    return JSON.parse(decodedPaylod);
}

const sha256 = async (text) => {
    const encoder = new TextEncoder();
    const data = encoder.encode(text);
    const hashBuffer = await crypto.subtle.digest('SHA-256', data);
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    return hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
}

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
                    password: await sha256(password)
                })
            });

            if (response.ok) {
                const data = await response.json();
                const payload = getTokenPayload(data.access_token)
                const admin = payload.is_admin;
                const idUsuario = payload.id_usuario

                localStorage.setItem("access_token", data.access_token);
                localStorage.setItem("idUsuario", parseInt(idUsuario));
                localStorage.setItem("isAdmin", admin.toString());
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

    const handleBack = () => {
        navigate('/');
    };

    return (
        <div className="login-container">
            <button onClick={handleBack} className="back-button">
                ← Volver
            </button>
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
