import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Login from './Login'
import Layout from './Layout'
import Home from './Home'
import Actividades from './Actividades'
import AdminPanel from './AdminPanel'

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
