import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

function App() {
  const navigate = useNavigate();

  useEffect(() => {
    navigate('/signup');
  }, [navigate]);

  return null;
}

export default App;
