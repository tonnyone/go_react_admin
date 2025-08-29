import { ConfigProvider } from "antd";
import 'antd/dist/reset.css';
import zhCN from 'antd/locale/zh_CN';
import { createRoot } from 'react-dom/client';
import App from './App.jsx';

// dev 使用mock数据, npm run dev 的时候会自动注入
if (process.env.NODE_ENV === 'development') {
  // import('./mock');
}

createRoot(document.getElementById('root')).render(
    <ConfigProvider locale={zhCN}>
      <App />
    </ConfigProvider>
)

