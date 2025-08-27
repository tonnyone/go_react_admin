// import { useState } from 'react';
import './App.css';


import { message } from 'antd';
import { createContext } from 'react';
import { RouterProvider } from 'react-router-dom';
import { router } from './router/router.jsx';

export const MessageContext = createContext(null);


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
