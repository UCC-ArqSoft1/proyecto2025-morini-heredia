import './index.css'
import gymPortada from './assets/maxresdefault.jpg'

const Home = () => {

    return (
        <div>
            <img 
            className="img-gym"
            src={ gymPortada }></img>
            <a type='buttom'>Inscribase ya  </a>
        </div>
    )

}

export default Home;