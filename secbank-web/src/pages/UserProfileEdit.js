import React, { useState, useEffect } from 'react';
import InputMask from 'react-input-mask';
import {
  Button, TextField, Typography, Paper, Box,
} from '@mui/material';
import { styled } from '@mui/system';
import dayjs from 'dayjs';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import api from '../api/axiosBase';

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
    if (!token) {
      toast.error('Token não encontrado.');
      return;
    }

    try {
      const response = await api.get('/customer/info');
      if (response.ok) {
        const data = response.result.data;
        setUserData({
          id: data.ID || '',
          name: data.FullName || '',
          cpf: data.Document || '',
          phone: data.Phone || '',
          birthDate: dayjs(data.Birthday).format('YYYY-MM-DD') || '',
          email: data.Email || '',
          createdTimeStamp: data.CreatedTimeStamp || '',
          updatedTimeStamp: data.UpdatedTimeStamp || '',
        });
      } else {
        toast.error(response.messageError || 'Erro ao carregar dados do usuário.');
      }
    } catch (error) {
      const errorMessage = JSON.parse(error.message)?.messageError || 'Erro inesperado.';
      toast.error(errorMessage);
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

  const handleSubmit = async (e) => {
    e.preventDefault();
    const newErrors = {};

    if (!userData.phone.replace(/\D/g, '').match(/^\d{10,11}$/)) {
      newErrors.phone = 'Número de telefone inválido!';
    }

    if (dayjs().diff(dayjs(userData.birthDate), 'year') < 18) {
      newErrors.birthDate = 'Usuário deve ser maior de idade!';
    }

    setErrors(newErrors);

    if (Object.keys(newErrors).length === 0) {
      const token = localStorage.getItem('token');
      if (!token) {
        toast.error('Token não encontrado.');
        return;
      }

      const requestData = {
        ID: userData.id,
        FullName: userData.name,
        Phone: userData.phone.replace(/\D/g, ''), // Remove máscara
        Email: userData.email,
        Birthday: new Date(userData.birthDate).toISOString(),
        CreatedTimeStamp: userData.createdTimeStamp,
        UpdatedTimeStamp: new Date().toISOString(),
      };

      try {
        const response = await api.put('/customer', JSON.stringify(requestData));

        if (response.ok) {
          toast.success('Dados atualizados com sucesso!');
        } else {
          toast.error('Erro ao atualizar dados.');
        }
      } catch (error) {
        const errorMessage = JSON.parse(error.message)?.messageError || 'Erro inesperado.';
        toast.error(errorMessage);
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
          <InputMask
            mask="999.999.999-99"
            value={userData.cpf}
            disabled
          >
            {() => (
              <ReadOnlyTextField
                label="CPF"
                variant="outlined"
                fullWidth
                margin="normal"
              />
            )}
          </InputMask>
          <InputMask
            mask="(99) 99999-9999"
            value={userData.phone}
            onChange={(e) => setUserData({ ...userData, phone: e.target.value })}
          >
            {() => (
              <TextField
                label="Telefone"
                name="phone"
                variant="outlined"
                fullWidth
                margin="normal"
                error={!!errors.phone}
                helperText={errors.phone}
              />
            )}
          </InputMask>
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
