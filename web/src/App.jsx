// import { useState } from 'react';
import { Route, Routes } from 'react-router-dom';
import './App.css';
import MyLayout from './components/Layout';
import About from './pages/About';
import Menus from './pages/Menus';
import Resource from './pages/Resource';
import Role from './pages/Role';
import User from './pages/User';

function App() {
  return (
    <MyLayout> 
      <Routes>
        <Route path="list" element={<User />} />
        <Route path="role" element={<Role />} />
        <Route path="menu" element={<Menus />} />
        <Route path="resource" element={<Resource />} />
        <Route path="about" element={<About />} />
      </Routes>
    </MyLayout>
  )
}
export default App
