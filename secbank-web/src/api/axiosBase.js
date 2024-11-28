import axios from 'axios';

// Create an Axios instance with a centralized base URL
const api = axios.create({
  baseURL: 'https://secbank.api.vinilab.dev', // replace with your API base URL
  headers: {
    'Content-Type': 'application/json', // default headers (optional)
    // Add any other headers you need, e.g., Authorization
  },
});

// Optional: Set up interceptors for request or response if needed
api.interceptors.request.use(
  config => {
    config.headers['Authorization'] = `Bearer ${localStorage.getItem('token')}`
    return config;
  },
  error => Promise.reject(error)
);

api.interceptors.response.use(
  response => {
    return {ok:true, result:response.data }
  },
  error => {
    if(error.status === 400)
        return Promise.reject(new Error( JSON.stringify({ok:false,messageError:error.response.data.messageError})));

    else return Promise.reject(error)
  }
);

export default api;
