import React from 'react';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginPage from './pages/LoginPage';

import RegisterPage from './pages/RegisterPage';

import TransferPage from './pages/TransferPage';
import UserProfileEdit from './pages/UserProfileEdit';
import Extract from './pages/Extract';
import Home from './pages/Home';



// Definindo o tema do Material-UI
const theme = createTheme({
  palette: {
    primary: {
      main: '#3f51b5', // Azul
    },
    secondary: {
      main: '#f50057', // Rosa
    },
  },
});

const App = () => {
  return (
    
    <ThemeProvider theme={theme}>
    <Router>
    <Routes>
      <Route path="/" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/transfer" element={<TransferPage />} />
      <Route path="/user/edit" element={<UserProfileEdit/>} />
      <Route path="/extract" element={<Extract/>} />
      <Route path="/home" element={<Home/>} />

    </Routes>
  </Router>
  </ThemeProvider>
  );
};

export default App;
