import React, { useState } from 'react';
import { Button, TextField, Box, Typography, Paper, CircularProgress } from '@mui/material';
import { styled } from '@mui/system';
import { useNavigate } from 'react-router-dom';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import api from '../api/axiosBase';

const Background = styled(Box)({
  background: '#f5f5f5',
  minHeight: '100vh',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
});

const LoginPaper = styled(Paper)({
  padding: '40px',
  maxWidth: '400px',
  textAlign: 'center',
  borderRadius: '12px',
  boxShadow: '0px 10px 30px rgba(0,0,0,0.1)',
  backgroundColor: '#ffffff',
});

const StyledButton = styled(Button)({
  marginTop: '20px',
  background: '#3f51b5',
  padding: '10px 20px',
  borderRadius: '25px',
  color: '#fff',
  fontWeight: 'bold',
  textTransform: 'none',
  transition: '0.3s',
  '&:hover': {
    background: '#303f9f',
  },
});

const LoginPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async () => {
    setLoading(true);
    debugger
    try {
      const response = await api.post('/login', JSON.stringify({ email, password }));
      const result = response.result;

      if (response.ok) {
        localStorage.clear();
        localStorage.setItem('token', result.data.token);
        toast.success('Login realizado com sucesso!');
        navigate('/home');
      } else {
        toast.error('Erro ao realizar login. Verifique suas credenciais.');
      }
    } catch (error) {
      const errorMessage = JSON.parse(error.message).messageError || 'Erro ao realizar login';
      toast.error(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  const handleRegister = () => {
    navigate('/register');
  };

  return (
    <Background>
      <ToastContainer
        position="top-right"
        autoClose={3000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
      />
      <LoginPaper elevation={6}>
        <Typography variant="h4" gutterBottom color="primary" fontWeight="bold">
          Acesse sua Conta
        </Typography>
        <TextField
          label="Email"
          variant="outlined"
          fullWidth
          margin="normal"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          InputProps={{
            style: { borderRadius: '10px' },
          }}
        />
        <TextField
          label="Senha"
          variant="outlined"
          type="password"
          fullWidth
          margin="normal"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          InputProps={{
            style: { borderRadius: '10px' },
          }}
        />
        {loading ? (
          <CircularProgress sx={{ mt: 3 }} />
        ) : (
          <StyledButton variant="contained" onClick={handleLogin} fullWidth>
            Entrar
          </StyledButton>
        )}
        <Typography
          variant="body2"
          sx={{ mt: 2, color: '#888', cursor: 'pointer' }}
          onClick={handleRegister}
        >
          Não tem uma conta? Registre-se
        </Typography>
      </LoginPaper>
    </Background>
  );
};

export default LoginPage;
