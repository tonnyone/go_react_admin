import { ConfigProvider } from "antd";
import 'antd/dist/reset.css';
import zhCN from 'antd/locale/zh_CN';
import { createRoot } from 'react-dom/client';
import App from './App.jsx';


createRoot(document.getElementById('root')).render(
    <ConfigProvider locale={zhCN}>
      <App />
    </ConfigProvider>
)