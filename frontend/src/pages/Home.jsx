import gymPortada from './../assets/maxresdefault.jpg'
import './../styles/index.css'

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