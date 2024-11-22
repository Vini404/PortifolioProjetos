import React, { useState, useRef, useCallback } from 'react';
import {
  Box, Button, TextField, Typography, Paper, Stepper, Step, StepLabel, Grid,
} from '@mui/material';
import { useNavigate } from 'react-router-dom';
import Sidebar from '../components/Sidebar';
import Navbar from '../components/Navbar';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import api from '../api/axiosBase';
import Webcam from 'react-webcam';

const TransferPage = () => {
  const [activeStep, setActiveStep] = useState(0);
  const [transferData, setTransferData] = useState({
    amount: '',
    numberCreditAccount: '',
    digitCreditAccount: '',
  });
  const [photo, setPhoto] = useState(null);

  const webcamRef = useRef(null);
  const navigate = useNavigate();

  const steps = ['Valor e Conta de Origem', 'Conta de Destino', 'Resumo da Transferência', 'Verificação de Identidade'];

  const handleNext = () => {
    setActiveStep((prevActiveStep) => prevActiveStep + 1);
  };

  const handleBack = () => {
    setActiveStep((prevActiveStep) => prevActiveStep - 1);
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    if (name === 'amount') {
      // Formatando o valor com separador de milhar
      const formattedValue = value.replace(/\D/g, '').replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,');
      setTransferData({ ...transferData, [name]: formattedValue });
    } else {
      setTransferData({ ...transferData, [name]: value });
    }
  };

  const handleCapture = useCallback(() => {
    const imageSrc = webcamRef.current.getScreenshot();
    setPhoto(imageSrc);
  }, [webcamRef]);

  const handleRetake = () => {
    setPhoto(null);
  };

  const handleConfirm = async () => {
    const token = localStorage.getItem('token');
    if (!token) {
      toast.error('Token não encontrado!');
      return;
    }
  
    if (!transferData.amount) {
      toast.error('Por favor, insira um valor para o montante.');
      return;
    }
  
    const formData = new FormData();
    formData.append('DigitCreditAccount', transferData.digitCreditAccount);
    formData.append('NumberCreditAccount', transferData.numberCreditAccount);
    // Remover vírgulas antes de enviar o valor para o backend
    const formattedAmount = transferData.amount.replace(/\D/g, '');
    formData.append('Amount', formattedAmount);
  
    try {
      const responseBlob = await fetch(photo);
      const blob = await responseBlob.blob();
      formData.append('file', blob, 'photo.jpg');
  
      const response = await api.post('/transaction', formData, {
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'multipart/form-data',
        },
      });
  
      if (response.ok) {
        toast.success('Transferência realizada com sucesso!');
        navigate('/extract');
      } else {
        toast.error('Erro ao realizar a transferência.');
      }
    } catch (error) {
      const errorMessage = JSON.parse(error.message)?.messageError || 'Erro inesperado.';
      toast.error(errorMessage);
    }
  };
  

  const renderStepContent = (step) => {
    switch (step) {
      case 0:
        return (
          <TextField
            label="Valor da Transferência"
            variant="outlined"
            fullWidth
            margin="normal"
            name="amount"
            value={`R$ ${transferData.amount}`} // Adicionando o símbolo de real
            onChange={handleChange}
            type="text"
            required
          />
        );
      case 1:
        return (
          <Grid container spacing={2}>
            <Grid item xs={9}>
              <TextField
                label="Número da Conta de Destino"
                variant="outlined"
                fullWidth
                margin="normal"
                name="numberCreditAccount"
                value={transferData.numberCreditAccount}
                onChange={handleChange}
                inputProps={{ maxLength: 7 }}
                required
              />
            </Grid>
            <Grid item xs={3}>
              <TextField
                label="Dígito"
                variant="outlined"
                fullWidth
                margin="normal"
                name="digitCreditAccount"
                value={transferData.digitCreditAccount}
                onChange={handleChange}
                inputProps={{ maxLength: 1 }}
                required
              />
            </Grid>
          </Grid>
        );
      case 2:
        return (
          <Box>
            <Typography variant="h6" gutterBottom>
              Resumo da Transferência
            </Typography>
            <Typography variant="body1">
              <strong>Para:</strong> Conta {transferData.numberCreditAccount}-{transferData.digitCreditAccount}
            </Typography>
            <Typography variant="body1">
              <strong>Valor:</strong> R$ {transferData.amount}
            </Typography>
          </Box>
        );
      case 3:
        return (
          <Box textAlign="center">
            <Typography
              variant="h6"
              gutterBottom
              sx={{
                color: '#303f9f', // Cor azul predominante
                fontWeight: 'bold',
                mb: 2,
              }}
            >
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
                <Box sx={{ display: 'flex', justifyContent: 'center', mt: 2 }}>
                  <Button variant="contained" color="primary" onClick={handleCapture}>
                    Tirar Foto
                  </Button>
                </Box>
              </Box>
            ) : (
              <Box mt={2} textAlign="center">
                <Typography
                  variant="h5"
                  sx={{
                    color: '#303f9f', // Cor azul predominante
                    fontWeight: 'bold',
                    mb: 2,
                  }}
                >
                  Foto Capturada
                </Typography>
                <img src={photo} alt="Foto capturada" style={{ width: '100%', maxWidth: '300px' }} />
                <Box sx={{ display: 'flex', justifyContent: 'center', mt: 2 }}>
                  <Button variant="contained" color="secondary" onClick={handleRetake}>
                    Nova Tentativa
                  </Button>
                </Box>
              </Box>
            )}
          </Box>
        );
      default:
        return 'Etapa desconhecida';
    }
  };

  return (
    <Box
      sx={{
        width: '100%',
        height: '100vh',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
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
      <Paper sx={{ p: 4, maxWidth: 600 }}>
        <Stepper activeStep={activeStep} alternativeLabel>
          {steps.map((label, index) => (
            <Step key={index}>
              <StepLabel>{label}</StepLabel>
            </Step>
          ))}
        </Stepper>

        <Box sx={{ mt: 3 }}>
          {renderStepContent(activeStep)}

          <Box sx={{ mt: 3 }}>
            {activeStep > 0 && (
              <Button onClick={handleBack} sx={{ mr: 1 }}>
                Voltar
              </Button>
            )}
            {activeStep < steps.length - 1 ? (
              <Button variant="contained" color="primary" onClick={handleNext}>
                Próximo
              </Button>
            ) : (
              <Button variant="contained" color="primary" onClick={handleConfirm}>
                Confirmar Transferência
              </Button>
            )}
          </Box>
        </Box>
      </Paper>
    </Box>
  );
};

export default TransferPage;
