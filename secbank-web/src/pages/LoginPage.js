import React, { useState } from 'react';
import { Button, TextField, Box, Typography, Paper, CircularProgress } from '@mui/material';
import { styled } from '@mui/system';
import { useNavigate } from 'react-router-dom';

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
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleLogin = async () => {
    setLoading(true);
    setError('');

    try {
      const response = await fetch('http://secbank-lb-1340144523.us-east-1.elb.amazonaws.com/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });

      const result = await response.json();

      if (result.success) {
        // Salva o token no localStorage
        localStorage.setItem('token', result.data.token);
        navigate('/home');
      } else {
        setError(result.messageError || 'Erro ao autenticar');
      }
    } catch (err) {
      setError('Erro de conexão com o servidor');
    } finally {
      setLoading(false);
    }
  };

  const handleForgotPassword = () => {
    navigate('/passwordRecovery');
  };

  const handleRegister = () => {
    navigate('/register');
  };

  return (
    <Background>
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
        {error && (
          <Typography color="error" sx={{ mt: 2 }}>
            {error}
          </Typography>
        )}
        <Typography
          variant="body2"
          sx={{ mt: 2, color: '#888', cursor: 'pointer' }}
          onClick={handleForgotPassword}
        >
          Esqueceu sua senha?
        </Typography>
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
