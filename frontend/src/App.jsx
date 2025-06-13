import { useState } from 'react'
import './styles/App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Login from './pages/Login'
import Home from './Home'
import Actividades from './pages/Actividades'
import AdminPanel from './pages/AdminPanel'
import Layout from './Layout'

function App() {
  const [count, setCount] = useState(0)

  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/" element={<Layout />}>
                  <Route index element={<Home />} />
                  <Route path="/actividades" element={<Actividades />} />
                  <Route path="/admin" element={<AdminPanel />} />
        </Route>
      </Routes>
    </Router>
  )
}

export default App
