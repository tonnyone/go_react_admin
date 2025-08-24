import { ConfigProvider } from "antd";
import 'antd/dist/reset.css';
import zhCN from 'antd/locale/zh_CN';
import { createRoot } from 'react-dom/client';
import App from './App.jsx';
import Login from './pages/Login.jsx';

import { Route, HashRouter as Router, Routes } from 'react-router-dom';

createRoot(document.getElementById('root')).render(
  <Router>
     <ConfigProvider locale={zhCN}>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/user/*" element={<App />} />
        <Route path="/*" element={<App />} />
      </Routes>
    </ConfigProvider>
  </Router>,
)