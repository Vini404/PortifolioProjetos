import React, { useState } from 'react';
import { Button, TextField, Typography, Paper, Box } from '@mui/material';
import { styled } from '@mui/system';
import { useNavigate } from 'react-router-dom'; // Importação para navegação

const FormContainer = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(4),
  margin: 'auto',
  maxWidth: 400,
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  backgroundColor: theme.palette.background.default,
}));

const PasswordRecoveryPage = () => {
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');
  const navigate = useNavigate(); // Hook de navegação

  const handleEmailChange = (e) => {
    setEmail(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Simular envio de solicitação de recuperação de senha
    setMessage(`Um e-mail de recuperação de senha foi enviado para ${email}`);
    setTimeout(() => {
      navigate('/'); // Redireciona para a página de login após 3 segundos
    }, 3000);
  };

  return (
    <Box
      sx={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
        backgroundColor: '#f5f5f5',
      }}
    >
      <FormContainer elevation={3}>
        <Typography variant="h6" gutterBottom color="primary" fontWeight="bold">
          Recuperar Senha
        </Typography>
        <Typography variant="body2" gutterBottom>
          Insira seu e-mail para receber um link de recuperação de senha.
        </Typography>
        <form onSubmit={handleSubmit}>
          <TextField
            label="E-mail"
            variant="outlined"
            fullWidth
            margin="normal"
            type="email"
            value={email}
            onChange={handleEmailChange}
            required
          />
          <Button
            variant="contained"
            color="primary"
            fullWidth
            type="submit"
            sx={{ mt: 2 }}
          >
            Enviar
          </Button>
        </form>
        {message && (
          <Typography variant="body2" color="success" sx={{ mt: 2 }}>
            {message}
          </Typography>
        )}
      </FormContainer>
    </Box>
  );
};

export default PasswordRecoveryPage;
