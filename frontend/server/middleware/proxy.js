// server/middleware/proxy.js
import { createProxyMiddleware } from 'http-proxy-middleware';
import { fromNodeMiddleware } from 'h3';

const apiProxy = createProxyMiddleware('/api', {
  target: 'http://localhost:3300', // Replace with your backend API URL
  changeOrigin: true,
  pathRewrite: { '^/api': '' }, // Remove /api prefix when forwarding
});

export default fromNodeMiddleware(apiProxy);