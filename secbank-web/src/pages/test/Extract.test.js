import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import Extract from '../Extract';
import api from '../../api/axiosBase';

// Mock da API
jest.mock('../../api/axiosBase', () => ({
  get: jest.fn(),
}));

// Função auxiliar para renderizar com roteamento
const renderWithRouter = (ui) => {
  return render(
    <MemoryRouter>
      {ui}
    </MemoryRouter>
  );
};

// Variáveis para capturar avisos
let consoleWarnSpy;
let consoleErrorSpy;

describe('Extract Component', () => {
  beforeEach(() => {
    // Ignorar warnings e erros
    consoleWarnSpy = jest.spyOn(console, 'warn').mockImplementation(() => {});
    consoleErrorSpy = jest.spyOn(console, 'error').mockImplementation(() => {});
    jest.clearAllMocks();
    localStorage.setItem('token', 'mock-token');
  });

  afterEach(() => {
    // Restaurar métodos console
    consoleWarnSpy.mockRestore();
    consoleErrorSpy.mockRestore();
    jest.restoreAllMocks();
  });

  test('renders loading state initially', async () => {
    renderWithRouter(
      <>
        <Extract />
        <ToastContainer />
      </>
    );

    expect(screen.getByText(/Carregando.../i)).toBeInTheDocument();
  });

  test('handles token not found error', async () => {
    localStorage.removeItem('token'); // Simula token ausente

    renderWithRouter(
      <>
        <Extract />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Token não encontrado!/i)).toBeInTheDocument();
    });
  });

  test('handles user info API error', async () => {
    api.get.mockResolvedValueOnce({
      ok: false, // Simula falha no `/customer/info`
    });

    renderWithRouter(
      <>
        <Extract />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Erro ao obter informações do usuário/i)).toBeInTheDocument();
    });
  });

  test('handles extract API error', async () => {
    api.get
      .mockResolvedValueOnce({
        ok: true,
        result: { data: { ID: 123 } }, // Simula sucesso no `/customer/info`
      })
      .mockResolvedValueOnce({
        ok: false, // Simula falha no `/balance/extract/:id`
      });

    renderWithRouter(
      <>
        <Extract />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Erro ao obter extrato/i)).toBeInTheDocument();
    });
  });

  test('handles unexpected error gracefully', async () => {
    api.get.mockRejectedValueOnce(new Error(JSON.stringify({ messageError: 'Erro inesperado' })));

    renderWithRouter(
      <>
        <Extract />
        <ToastContainer />
      </>
    );

    await waitFor(() => {
      expect(screen.getByText(/Erro inesperado/i)).toBeInTheDocument();
    });
  });
});
