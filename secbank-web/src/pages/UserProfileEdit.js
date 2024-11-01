import React, { useState, useEffect } from 'react';
import {
  Button, TextField, Typography, Paper, Box, InputAdornment
} from '@mui/material';
import { styled } from '@mui/system';
import { isValidPhoneNumber } from 'libphonenumber-js';
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
  backgroundColor: '#e0e0e0',
  color: '#757575',
  '& .MuiInputBase-input': {
    color: '#757575',
  },
});

const UserProfileEdit = () => {
  const [userData, setUserData] = useState({
    id: '',
    name: '',
    cpf: '',
    phone: '',
    birthDate: '',
    email: '',
    createdTimeStamp: '',
    updatedTimeStamp: '',
  });

  const [errors, setErrors] = useState({});
  const [loading, setLoading] = useState(true);

  const fetchUserData = async () => {
    const token = localStorage.getItem('token');
    if (!token) return alert("Token não encontrado");

    try {
      const response = await fetch('http://secbank-lb-1340144523.us-east-1.elb.amazonaws.com/customer/info', {
        method: 'GET',
        headers: { 'Authorization': `Bearer ${token}` },
      });
      const result = await response.json();

      if (response.ok && result.success) {
        const data = result.data;
        setUserData({
          id: data.ID || '',
          name: data.FullName || '',
          cpf: data.CPF || '',
          phone: data.Phone || '',
          birthDate: data.Birthday || '',
          email: data.Email || '',
          createdTimeStamp: data.CreatedTimeStamp || '',
          updatedTimeStamp: data.UpdatedTimeStamp || '',
        });
      } else {
        alert(result.messageError || 'Erro ao carregar dados do usuário');
      }
    } catch (error) {
      console.error('Erro ao buscar dados:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUserData();
  }, []);

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

  const handleSubmit = async (e) => {
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
      const token = localStorage.getItem('token');
      if (!token) return alert("Token não encontrado");

      const requestData = {
        ID: userData.id,
        FullName: userData.name,
        Phone: userData.phone,
        Email: userData.email,
        Birthday: new Date(userData.birthDate).toISOString(),
        CreatedTimeStamp: userData.createdTimeStamp,
        UpdatedTimeStamp: new Date().toISOString(),
      };

      try {
        const response = await fetch('http://secbank-lb-1340144523.us-east-1.elb.amazonaws.com/customer', {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
          body: JSON.stringify(requestData),
        });

        if (response.ok) {
          alert('Dados atualizados com sucesso!');
        } else {
          alert('Erro ao atualizar dados');
        }
      } catch (error) {
        console.error('Erro ao atualizar dados:', error);
      }
    }
  };

  if (loading) return <Typography>Carregando...</Typography>;

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
