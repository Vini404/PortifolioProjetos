import React, { useState } from 'react';
import {
  Button, TextField, Typography, Paper, Box, InputAdornment
} from '@mui/material';
import { styled } from '@mui/system';
import { isValidPhoneNumber } from 'libphonenumber-js'; // Biblioteca para validar telefone (instalar via npm)
import dayjs from 'dayjs';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';

const FormContainer = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(4),
  margin: 'auto',
  maxWidth: 500,
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  backgroundColor: theme.palette.background.default,
}));

const ReadOnlyTextField = styled(TextField)({
  backgroundColor: '#e0e0e0', // Fundo cinza claro
  color: '#757575', // Texto com tom mais claro
  '& .MuiInputBase-input': {
    color: '#757575', // Cor do texto dentro do campo
  },
});

const UserProfileEdit = () => {
  const [userData, setUserData] = useState({
    name: 'João da Silva',
    cpf: '123.456.789-10',
    phone: '',
    birthDate: '',
    email: '',
  });

  const [errors, setErrors] = useState({});

  const handleChange = (e) => {
    setUserData({ ...userData, [e.target.name]: e.target.value });
  };

  const validatePhone = (phone) => {
    return isValidPhoneNumber(phone, 'BR');
  };

  const isUnderage = (date) => {
    const today = dayjs();
    const birthDate = dayjs(date);
    return today.diff(birthDate, 'year') < 18;
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const newErrors = {};

    if (!validatePhone(userData.phone)) {
      newErrors.phone = 'Número de telefone inválido!';
    }

    if (isUnderage(userData.birthDate)) {
      newErrors.birthDate = 'Usuário deve ser maior de idade!';
    }

    setErrors(newErrors);

    if (Object.keys(newErrors).length === 0) {
      alert('Dados atualizados com sucesso!');
      console.log(userData);
    }
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
    <Navbar />
    <Sidebar />
      <FormContainer elevation={3}>
        <Typography variant="h4" gutterBottom color="primary" fontWeight="bold">
          Editar Registro
        </Typography>
        <form onSubmit={handleSubmit} style={{ width: '100%' }}>
          <ReadOnlyTextField
            label="Nome"
            variant="outlined"
            fullWidth
            margin="normal"
            value={userData.name}
            InputProps={{
              readOnly: true,
            }}
          />
          <ReadOnlyTextField
            label="CPF"
            variant="outlined"
            fullWidth
            margin="normal"
            value={userData.cpf}
            InputProps={{
              readOnly: true,
            }}
          />
          <TextField
            label="Telefone"
            name="phone"
            variant="outlined"
            fullWidth
            margin="normal"
            value={userData.phone}
            onChange={handleChange}
            InputProps={{
              startAdornment: <InputAdornment position="start">+55</InputAdornment>,
              placeholder: '(XX) XXXXX-XXXX',
            }}
            error={!!errors.phone}
            helperText={errors.phone}
          />
          <TextField
            label="Data de Nascimento"
            name="birthDate"
            variant="outlined"
            type="date"
            fullWidth
            margin="normal"
            value={userData.birthDate}
            onChange={handleChange}
            InputLabelProps={{
              shrink: true,
            }}
            error={!!errors.birthDate}
            helperText={errors.birthDate}
          />
          <TextField
            label="Email"
            name="email"
            variant="outlined"
            fullWidth
            margin="normal"
            value={userData.email}
            onChange={handleChange}
          />
          <Button variant="contained" color="primary" type="submit" fullWidth sx={{ mt: 2 }}>
            Salvar
          </Button>
        </form>
      </FormContainer>
    </Box>
  );
};

export default UserProfileEdit;
