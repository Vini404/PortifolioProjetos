import React, { useState, useRef, useCallback } from 'react';
import {
  Box, Button, TextField, Typography, Paper, Stepper, Step, StepLabel,
} from '@mui/material';
import { useNavigate } from 'react-router-dom';
import Sidebar from '../components/Sidebar';
import Navbar from '../components/Navbar';
import api from '../api/axiosBase';
import Webcam from 'react-webcam';

const TransferPage = () => {
  const [activeStep, setActiveStep] = useState(0);
  const [transferData, setTransferData] = useState({
    amount: '',
    numberCreditAccount: '',
    digitCreditAccount: '',
  });
  const [photo, setPhoto] = useState(null); // Guarda a foto capturada

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
    setTransferData({ ...transferData, [e.target.name]: e.target.value });
  };

  const handleCapture = useCallback(() => {
    const imageSrc = webcamRef.current.getScreenshot();
    setPhoto(imageSrc); // Salva a imagem capturada no estado
  }, [webcamRef]);

  const handleRetake = () => {
    setPhoto(null); // Limpa a foto capturada para permitir uma nova tentativa
  };

  const handleConfirm = async () => {
    const token = localStorage.getItem('token');
    if (!token) {
      alert('Token não encontrado!');
      return;
    }

    // Preparar os dados para a transferência, incluindo a foto
    const requestData = {
      digit_credit_account: transferData.digitCreditAccount,
      number_credit_account: transferData.numberCreditAccount,
      amount: parseFloat(transferData.amount),
      photo, // Adiciona a foto capturada na requisição
    };

    try {
      const response = await api.post('/transaction', JSON.stringify(requestData));

      if (response.ok) {
        alert('Transferência realizada com sucesso!');
        navigate('/extract');
      }
    } catch (error) {
      const errorMessage = JSON.parse(error.message).messageError;
      alert(errorMessage);
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
            value={transferData.amount}
            onChange={handleChange}
            type="number"
            required
          />
        );
      case 1:
        return (
          <>
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
            <TextField
              label="Dígito da Conta de Destino"
              variant="outlined"
              fullWidth
              margin="normal"
              name="digitCreditAccount"
              value={transferData.digitCreditAccount}
              onChange={handleChange}
              inputProps={{ maxLength: 1 }}
              required
            />
          </>
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
              <strong>Valor:</strong> R$ {parseFloat(transferData.amount).toFixed(2)}
            </Typography>
          </Box>
        );
      case 3:
        return (
          <Box>
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
                <Box sx={{ display: 'flex', justifyContent: 'center', mt: 2 }}>
                  <Button variant="contained" color="primary" onClick={handleCapture}>
                    Tirar Foto
                  </Button>
                </Box>
              </Box>
            ) : (
              <Box mt={2} textAlign="center">
                <Typography>Foto Capturada:</Typography>
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
