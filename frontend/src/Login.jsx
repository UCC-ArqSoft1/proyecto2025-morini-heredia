import { Fragment, useState } from "react";
import './Login.css';
import { useNavigate } from "react-router-dom";

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handlerLogin = async (e) => {
        e.preventDefault(); // como es submit, se utiliza esto para evitar que se recargue la pagina

        if (username=="admin" && password=="admin"){
            console.log("Login OK");
            localStorage.setItem("isLoggedIn", "true");

            // Aca va la logica de los admin
            navigate("/actividades");
        } else {
            console.log("Login incorrecto");
        }
    }

    return (
        <div className="login-container"> 
            <form className="login-form" onSubmit={ handlerLogin }>
                <h2> Iniciar Sesion</h2>
                <input
                    type="text"
                    placeholder="Usuario"
                    value={ username }
                    onChange={(e) => setUsername(e.target.value)}
                    required
                />
                <input
                    type="password"
                    placeholder="Contrasena"
                    value = { password }
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />

                <button type="submit"> Ingresar </button>
            </form>
        </div>
    )
}

export default Login;