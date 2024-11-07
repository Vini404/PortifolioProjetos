import React, { useState } from 'react';
import { Button, TextField, Box, Typography, Paper } from '@mui/material';
import { styled } from '@mui/system';
import api from '../api/axiosBase'

const Background = styled(Box)({
  background: '#f5f5f5',
  minHeight: '100vh',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
});

const RegisterPaper = styled(Paper)({
  padding: '40px',
  maxWidth: '500px',
  textAlign: 'center',
  borderRadius: '12px',
  boxShadow: '0px 10px 30px rgba(0,0,0,0.1)',
  backgroundColor: '#ffffff',
});

const StyledButton = styled(Button)(({ error }) => ({
  marginTop: '20px',
  background: '#3f51b5',
  padding: '10px 20px',
  borderRadius: '25px',
  color: '#fff',
  fontWeight: 'bold',
  textTransform: 'none',
  transition: '0.3s',
  border: error ? '2px solid red' : 'none',
  '&:hover': {
    background: '#303f9f',
  },
}));

const RegisterPage = () => {
  const [fullName, setFullName] = useState('');
  const [phone, setPhone] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [birthday, setBirthday] = useState('');
  const [document, setDocument] = useState('');
  const [passwordError, setPasswordError] = useState(false);
  const [registrationError, setRegistrationError] = useState('');
  const [registrationSuccess, setRegistrationSuccess] = useState(false);

  const handleRegister = async () => {
    if (password !== confirmPassword) {
      setPasswordError(true);
      return;
    }

    // Formatação dos dados para o contrato
    const requestData = {
      FullName: fullName,
      Phone: phone,
      Email: email,
      Birthday: new Date(birthday).toISOString(),
      Password: password,
      Document: document,
    };

    try {
      const response = await api.post('/customer', JSON.stringify(requestData));

      if (response.ok) {
        setRegistrationSuccess(true);
        setRegistrationError('');
        alert('Usuário registrado com sucesso!');
      }
    } catch (error) {
      const errorMessage = JSON.parse(error.message).messageError
      alert(errorMessage);
    }
  };

  return (
    <Background>
      <RegisterPaper elevation={6}>
        <Typography variant="h4" gutterBottom color="primary" fontWeight="bold">
          Registrar
        </Typography>
        <TextField
          label="Nome Completo"
          variant="outlined"
          fullWidth
          margin="normal"
          value={fullName}
          onChange={(e) => setFullName(e.target.value)}
          InputProps={{
            style: { borderRadius: '10px' },
          }}
        />
        <TextField
          label="Telefone"
          variant="outlined"
          fullWidth
          margin="normal"
          value={phone}
          onChange={(e) => setPhone(e.target.value)}
          InputProps={{
            style: { borderRadius: '10px' },
          }}
        />
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
          label="Documento"
          variant="outlined"
          fullWidth
          margin="normal"
          value={document}
          onChange={(e) => setDocument(e.target.value)}
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
        <TextField
          label="Confirmar Senha"
          variant="outlined"
          type="password"
          fullWidth
          margin="normal"
          value={confirmPassword}
          onChange={(e) => {
            setConfirmPassword(e.target.value);
            setPasswordError(false);
          }}
          InputProps={{
            style: { borderRadius: '10px' },
          }}
          error={passwordError}
          helperText={passwordError && 'As senhas não coincidem'}
        />
        <TextField
          label="Data de Nascimento"
          variant="outlined"
          type="date"
          fullWidth
          margin="normal"
          value={birthday}
          onChange={(e) => setBirthday(e.target.value)}
          InputLabelProps={{
            shrink: true,
          }}
          InputProps={{
            style: { borderRadius: '10px' },
          }}
        />
        <StyledButton
          variant="contained"
          onClick={handleRegister}
          fullWidth
          error={passwordError}
        >
          Registrar
        </StyledButton>
        {registrationError && (
          <Typography color="error" sx={{ mt: 2 }}>
            {registrationError}
          </Typography>
        )}
        {registrationSuccess && (
          <Typography color="primary" sx={{ mt: 2 }}>
            Registro bem-sucedido!
          </Typography>
        )}
      </RegisterPaper>
    </Background>
  );
};

export default RegisterPage;
