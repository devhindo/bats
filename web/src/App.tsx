import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

function App() {
  const navigate = useNavigate();

  useEffect(() => {
    const accessToken = localStorage.getItem('accessToken');
    if (!accessToken) {
      navigate('/signup');
    }
  }, [navigate]);

  return (
    <div>
      <h1>App</h1>
    </div>
  );
}

export default App;
