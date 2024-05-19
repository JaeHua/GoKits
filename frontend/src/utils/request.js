import axios from 'axios';
import storageService from "@/service/storageService.js";

const service = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_URL,
    timeout: 1000 * 5,
    headers: { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` },
});

// 解决token读取到的是上一个的记录
service.interceptors.request.use((config) => {
    // Do something before request is sent
    Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` });
    return config;
}, (error) => {
    // Do something with request error
    return Promise.reject(error);
});
export default service