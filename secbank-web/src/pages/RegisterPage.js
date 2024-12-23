import React, { useState, useRef, useCallback } from 'react';
import InputMask from 'react-input-mask';
import { Button, TextField, Box, Typography, Paper } from '@mui/material';
import { styled } from '@mui/system';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import Webcam from 'react-webcam';
import { useNavigate } from 'react-router-dom';
import api from '../api/axiosBase';

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
  const [activeStep, setActiveStep] = useState(0);
  const [fullName, setFullName] = useState('');
  const [phone, setPhone] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [birthday, setBirthday] = useState('');
  const [document, setDocument] = useState('');
  const [passwordError, setPasswordError] = useState(false);
  const [photo, setPhoto] = useState(null);

  const webcamRef = useRef(null);
  const navigate = useNavigate();

  const handleCapture = useCallback(() => {
    const imageSrc = webcamRef.current.getScreenshot();
    fetch(imageSrc)
      .then((res) => res.blob())
      .then((blob) => {
        const file = new File([blob], 'photo.jpg', { type: 'image/jpeg' });
        setPhoto(file);
      });
  }, [webcamRef]);

  const handleRetake = () => {
    setPhoto(null);
  };

  const handleNextStep = () => {
    if (
      !fullName.trim() ||
      !phone.trim() ||
      !email.trim() ||
      !password.trim() ||
      !confirmPassword.trim() ||
      !birthday.trim() ||
      !document.trim()
    ) {
      toast.error('Todos os campos são obrigatórios!');
      return;
    }

    if (password !== confirmPassword) {
      setPasswordError(true);
      toast.error('As senhas não coincidem!');
      return;
    }

    setActiveStep(1);
  };

  const handleBackStep = () => {
    setActiveStep(0);
  };

  const handleRegister = async () => {
    if (!photo) {
      toast.error('Por favor, capture uma foto!');
      return;
    }

    const formData = new FormData();
    formData.append('FullName', fullName);
    formData.append('Phone', phone.replace(/\D/g, '')); // Remove máscara
    formData.append('Email', email);
    formData.append('Birthday', new Date(birthday).toISOString());
    formData.append('Password', password);
    formData.append('Document', document.replace(/\D/g, '')); // Remove máscara
    formData.append('file', photo);

    try {
      const response = await api.post('/customer', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      if (response.ok) {
        toast.success('Usuário registrado com sucesso!');
        await handleLogin();
      } else {
        toast.error('Erro ao registrar o usuário.');
      }
    } catch (error) {
      console.log(error);
      const errorMessage =
        JSON.parse(error.message).messageError || 'Erro inesperado.';
      toast.error(errorMessage);
    }
  };

  const handleLogin = async () => {
    try {
      const response = await api.post(
        '/login',
        JSON.stringify({ email, password }),
        {
          headers: {
            'Content-Type': 'application/json',
          },
        }
      );

      const result = response.result.data;

      if (response.ok) {
        localStorage.clear();
        localStorage.setItem('token', result.token);
        toast.success('Login realizado com sucesso!');
        navigate('/home');
      } else {
        toast.error('Erro ao realizar login. Verifique suas credenciais.');
      }
    } catch (error) {
      const errorMessage =
        JSON.parse(error.message).messageError || 'Erro ao realizar login';
      toast.error(errorMessage);
    }
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
      <RegisterPaper elevation={6}>
        <Typography variant="h4" gutterBottom color="primary" fontWeight="bold">
          Registrar
        </Typography>
        {activeStep === 0 ? (
          <>
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
            <InputMask
              mask="(99) 99999-9999"
              value={phone}
              onChange={(e) => setPhone(e.target.value)}
            >
              {() => (
                <TextField
                  label="Telefone"
                  variant="outlined"
                  fullWidth
                  margin="normal"
                  InputProps={{
                    style: { borderRadius: '10px' },
                  }}
                />
              )}
            </InputMask>
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
            <InputMask
              mask="999.999.999-99"
              value={document}
              onChange={(e) => setDocument(e.target.value)}
            >
              {() => (
                <TextField
                  label="CPF"
                  variant="outlined"
                  fullWidth
                  margin="normal"
                  InputProps={{
                    style: { borderRadius: '10px' },
                  }}
                />
              )}
            </InputMask>
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
              onClick={handleNextStep}
              fullWidth
              error={passwordError}
            >
              Próximo
            </StyledButton>
          </>
        ) : (
          <>
            <Typography variant="h6" gutterBottom>
              Verificação de Identidade
            </Typography>
            {!photo ? (
              <Box>
                <Webcam
                  audio={false}
                  ref={webcamRef}
                  screenshotFormat="image/jpeg"
                  width="100%"
                />
                <Button
                  variant="contained"
                  color="primary"
                  onClick={handleCapture}
                  sx={{ mt: 2 }}
                >
                  Tirar Foto
                </Button>
                <Button
                  variant="outlined"
                  color="secondary"
                  onClick={handleBackStep}
                  sx={{ mt: 2, ml: 2 }}
                >
                  Voltar
                </Button>
              </Box>
            ) : (
              <Box mt={2} textAlign="center">
                <Typography>Foto Capturada:</Typography>
                <img
                  src={URL.createObjectURL(photo)}
                  alt="Foto capturada"
                  style={{ width: '100%', maxWidth: '300px' }}
                />
                <Box sx={{ display: 'flex', justifyContent: 'center', mt: 2 }}>
                  <Button variant="contained" color="secondary" onClick={handleRetake}>
                    Nova Tentativa
                  </Button>
                  <Button
                    variant="contained"
                    color="primary"
                    onClick={handleRegister}
                    sx={{ ml: 2 }}
                  >
                    Confirmar Registro
                  </Button>
                  <Button
                    variant="outlined"
                    color="secondary"
                    onClick={handleBackStep}
                    sx={{ ml: 2 }}
                  >
                    Voltar
                  </Button>
                </Box>
              </Box>
            )}
          </>
        )}
      </RegisterPaper>
    </Background>
  );
};

export default RegisterPage;