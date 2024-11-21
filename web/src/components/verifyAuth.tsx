import axios from 'axios';
import { useEffect, useState } from 'react';

const url = import.meta.env.VITE_BASE_URL + "usr/verify";

const useAuth = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    useEffect(() => {
        const checkAuth = async () => {
            try {
                const response = await axios.get(url, {
                    withCredentials: true // Include credentials (cookies)
                });
                if (response.status === 200 && response.data.message === 'authenticated') {
                    setIsLoggedIn(true);
                } else {
                    setIsLoggedIn(false);
                }
            } catch (error) {
                setIsLoggedIn(false);
            }
        };

        checkAuth();
    }, []);

    return isLoggedIn;
};

export default useAuth;