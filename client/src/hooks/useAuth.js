import { createContext, useContext, useMemo } from 'react';
import { useNavigate } from 'react-router-dom';
import { useLocalStorage } from './useLocalStorage';

const AuthContext = createContext();

export default function AuthProvider({ children }) {
  const [user, setUser] = useLocalStorage('user', null);
  const navigate = useNavigate();

  // call this function when you want to authenticate the user
  async function login(data) {
    setUser(data);
    navigate('/home')
  }

  // call this function to sign out logged in user
  function logout() {
    setUser(null);
    navigate('/', { replace: true });
  }

  const value = useMemo(
    () => ({
      user,
      login, 
      logout
    }),
    [user]
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

export function useAuth() {
  return useContext(AuthContext);
}
