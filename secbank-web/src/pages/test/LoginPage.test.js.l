import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import api from '../../api/axiosBase';
import LoginPage from '../LoginPage';

jest.mock('../../api/axiosBase', () => ({
  post: jest.fn(),
}));

// Mock useNavigate globally
const mockNavigate = jest.fn();
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockNavigate,
}));

// Mock localStorage
beforeEach(() => {
  const localStorageMock = (() => {
    let store = {};
    return {
      getItem: jest.fn((key) => store[key] || null),
      setItem: jest.fn((key, value) => {
        store[key] = value;
      }),
      clear: jest.fn(() => {
        store = {};
      }),
    };
  })();

  Object.defineProperty(window, 'localStorage', {
    value: localStorageMock,
  });
});

const renderWithRouter = (ui) => {
  return render(<MemoryRouter>{ui}</MemoryRouter>);
};

describe('LoginPage Component', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  test('renders login form correctly', () => {
    renderWithRouter(
      <>
        <LoginPage />
        <ToastContainer />
      </>
    );

    expect(screen.getByLabelText(/Email/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/Senha/i)).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /Entrar/i })).toBeInTheDocument();
    expect(screen.getByText(/Esqueceu sua senha\?/i)).toBeInTheDocument();
    expect(screen.getByText(/Não tem uma conta\? Registre-se/i)).toBeInTheDocument();
  });

  test('displays loading spinner when login is in progress', async () => {
    api.post.mockImplementationOnce(() => new Promise(() => {})); // Simulate a pending request

    renderWithRouter(<LoginPage />);

    fireEvent.change(screen.getByLabelText(/Email/i), { target: { value: 'test@example.com' } });
    fireEvent.change(screen.getByLabelText(/Senha/i), { target: { value: 'password123' } });
    fireEvent.click(screen.getByRole('button', { name: /Entrar/i }));

    expect(await screen.findByRole('progressbar')).toBeInTheDocument();
  });

  test('calls API and handles successful login', async () => {
    api.post.mockResolvedValueOnce({
      ok: true,
      result: {
        data: { token: 'mock-token' },
      },
    });

    renderWithRouter(
      <>
        <LoginPage />
        <ToastContainer />
      </>
    );

    fireEvent.change(screen.getByLabelText(/Email/i), { target: { value: 'test@example.com' } });
    fireEvent.change(screen.getByLabelText(/Senha/i), { target: { value: 'password123' } });
    fireEvent.click(screen.getByRole('button', { name: /Entrar/i }));

    await waitFor(() => {
      expect(api.post).toHaveBeenCalledWith(
        '/login',
        JSON.stringify({ email: 'test@example.com', password: 'password123' })
      );
      expect(localStorage.setItem).toHaveBeenCalledWith('token', 'mock-token');
      expect(mockNavigate).toHaveBeenCalledWith('/home');
      expect(screen.getByText(/Login realizado com sucesso!/i)).toBeInTheDocument();
    });
  });

  test('handles login failure', async () => {
    api.post.mockResolvedValueOnce({
      ok: false,
    });

    renderWithRouter(
      <>
        <LoginPage />
        <ToastContainer />
      </>
    );

    fireEvent.change(screen.getByLabelText(/Email/i), { target: { value: 'test@example.com' } });
    fireEvent.change(screen.getByLabelText(/Senha/i), { target: { value: 'wrongpassword' } });
    fireEvent.click(screen.getByRole('button', { name: /Entrar/i }));

    await waitFor(() => {
      expect(api.post).toHaveBeenCalledWith(
        '/login',
        JSON.stringify({ email: 'test@example.com', password: 'wrongpassword' })
      );
      expect(screen.getByText(/Erro ao realizar login. Verifique suas credenciais./i)).toBeInTheDocument();
    });
  });

  test('handles API errors gracefully', async () => {
    api.post.mockRejectedValueOnce(new Error(JSON.stringify({ messageError: 'Server error' })));

    renderWithRouter(
      <>
        <LoginPage />
        <ToastContainer />
      </>
    );

    fireEvent.change(screen.getByLabelText(/Email/i), { target: { value: 'test@example.com' } });
    fireEvent.change(screen.getByLabelText(/Senha/i), { target: { value: 'password123' } });
    fireEvent.click(screen.getByRole('button', { name: /Entrar/i }));

    await waitFor(() => {
      expect(screen.getByText(/Server error/i)).toBeInTheDocument();
    });
  });

  test('navigates to forgot password page', () => {
    renderWithRouter(<LoginPage />);

    fireEvent.click(screen.getByText(/Esqueceu sua senha\?/i));
    expect(mockNavigate).toHaveBeenCalledWith('/passwordRecovery');
  });

  test('navigates to register page', () => {
    renderWithRouter(<LoginPage />);

    fireEvent.click(screen.getByText(/Não tem uma conta\? Registre-se/i));
    expect(mockNavigate).toHaveBeenCalledWith('/register');
  });
});
