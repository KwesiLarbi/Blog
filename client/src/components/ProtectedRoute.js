import { Navigate } from 'react-router-dom';
import { useAuth } from '../hooks/useAuth';

export default function ProtectedRoute({ children }) {
  const { user } = useAuth();

  if (!user) {
    // user is not authenticated
    return <Navigate to="/" />
  }

  return children;
}
