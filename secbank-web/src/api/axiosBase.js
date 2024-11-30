import axios from 'axios';

const baseURL = process.env.REACT_APP_API_BASE_URL

console.log(process.env.NODE_ENV)

const api = axios.create({
  baseURL: baseURL, 
  headers: {
    'Content-Type': 'application/json', // default headers (optional)
  },
});

// Optional: Set up interceptors for request or response if needed
api.interceptors.request.use(
  config => {
    // Add Authorization token if it exists
    const token = localStorage.getItem('token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  error => Promise.reject(error)
);

api.interceptors.response.use(
  response => {
    return { ok: true, result: response.data };
  },
  error => {
    if (error.response?.status === 400) {
      return Promise.reject(
        new Error(
          JSON.stringify({ ok: false, messageError: error.response.data.messageError })
        )
      );
    }
    return Promise.reject(error);
  }
);

export default api;
