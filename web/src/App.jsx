// import { useState } from 'react';
import './App.css';


import { message } from 'antd';
import { createContext } from 'react';
import { createBrowserRouter, createRoutesFromElements, Route, RouterProvider } from 'react-router-dom';
import MyLayout from "./components/Layout.jsx";
import About from "./pages/About.jsx";
import Login from './pages/Login.jsx';
import Menus from "./pages/Menus.jsx";
import NotFound from "./pages/NotFound.jsx";
import Resource from "./pages/Resource.jsx";
import Role from "./pages/Role.jsx";
import User from "./pages/User.jsx";

export const MessageContext = createContext(null);

const router = createBrowserRouter(
  createRoutesFromElements(
    <>
      <Route path="/" element={<Login />} />
      <Route path="login" element={<Login />} />
      <Route path="about" element={<About />} />
      <Route path="user" element={<MyLayout />} >
          <Route path="list" index element={<User />} />
          <Route path="role" element={<Role />} />
          <Route path="menu" element={<Menus />} />
          <Route path="resource" element={<Resource />} />
      </Route>
      <Route path="*" element={<NotFound />} />
    </>
  )
);

function App() {
  const [messageApi, contextHolder] = message.useMessage();
  return (
    <MessageContext.Provider value={messageApi}>
      {contextHolder}
      <RouterProvider router={router} />
    </MessageContext.Provider>
  );
}
export default App;
