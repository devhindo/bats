import React, { useState } from 'react';
import axios from 'axios';
import logo from '../assets/bats.png'; // Adjust the path if necessary

const url = import.meta.env.VITE_BASE_URL;

const Signup: React.FC = () => {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [otp, setOtp] = useState('');
    const [isOtpSent, setIsOtpSent] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');
    const [passwordError, setPasswordError] = useState('');
    const [passwordValidation, setPasswordValidation] = useState({
        length: false,
        uppercase: false,
        number: false,
        specialChar: false,
    });
    const [isPasswordTouched, setIsPasswordTouched] = useState(false);

    const validatePassword = (password: string) => {
        const length = password.length >= 8;
        const uppercase = /[A-Z]/.test(password);
        const number = /\d/.test(password);
        const specialChar = /[@$!%*?&]/.test(password);

        setPasswordValidation({
            length,
            uppercase,
            number,
            specialChar,
        });

        return length && uppercase && number && specialChar;
    };

    const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const newPassword = e.target.value;
        setPassword(newPassword);
        setIsPasswordTouched(true);
        validatePassword(newPassword);
    };

    const handleRegister = async (event: React.FormEvent) => {
        event.preventDefault();
        if (!validatePassword(password)) {
            setPasswordError('Password must be at least 8 characters long, include an uppercase letter, a number, and a special character.');
            return;
        }
        try {
            await axios.post(url + 'signup', {
                username,
                email,
                password,
            });
            setIsOtpSent(true);
            setErrorMessage('');
            setPasswordError('');
        } catch (error) {
            if (axios.isAxiosError(error) && error.response?.status === 409) {
                setErrorMessage('Username already exists');
            } else {
                console.error('There was an error registering!', error);
            }
        }
    };

    const handleVerifyOtp = async (event: React.FormEvent) => {
        event.preventDefault();
        try {
            const response = await axios.post(url + 'signup/otp', {
                email,
                otp,
            });
            const token = response.data.token;
            localStorage.setItem('token', token);
            alert('Registration successful');
        } catch (error) {
            console.error('There was an error verifying OTP!', error);
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-900 px-4 sm:px-6 lg:px-8">
            <div className="bg-white dark:bg-gray-800 p-8 rounded-lg shadow-md w-full max-w-md">
                <div className="flex items-center justify-center mb-6">
                    <img src={logo} alt="Bats Logo" className="h-12 w-12 mr-2" />
                    <span className="text-2xl font-bold text-gray-700 dark:text-gray-300">Bats</span>
                </div>
                {!isOtpSent ? (
                    <form onSubmit={handleRegister}>
                        {errorMessage && (
                            <div className="mb-4 text-red-500">
                                {errorMessage}
                            </div>
                        )}
                        <div className="mb-4">
                            <label className="block text-gray-700 dark:text-gray-300">Username:</label>
                            <input
                                type="text"
                                value={username}
                                onChange={(e) => setUsername(e.target.value)}
                                className="mt-1 block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 dark:focus:ring-indigo-400 dark:focus:border-indigo-400 text-gray-700 dark:text-gray-300"
                            />
                        </div>
                        <div className="mb-4">
                            <label className="block text-gray-700 dark:text-gray-300">Email:</label>
                            <input
                                type="email"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                className="mt-1 block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 dark:focus:ring-indigo-400 dark:focus:border-indigo-400 text-gray-700 dark:text-gray-300"
                            />
                        </div>
                        <div className="mb-4">
                            <label className="block text-gray-700 dark:text-gray-300">Password:</label>
                            <input
                                type="password"
                                value={password}
                                onChange={handlePasswordChange}
                                className="mt-1 block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 dark:focus:ring-indigo-400 dark:focus:border-indigo-400 text-gray-700 dark:text-gray-300"
                            />
                            {passwordError && (
                                <div className="mt-2 text-red-500">
                                    {passwordError}
                                </div>
                            )}
                            {isPasswordTouched && (
                                <div className="mt-2">
                                    <div className={`text-sm ${passwordValidation.length ? 'text-green-500' : 'text-red-500'}`}>
                                        {passwordValidation.length ? '✔' : '✘'} At least 8 characters
                                    </div>
                                    <div className={`text-sm ${passwordValidation.uppercase ? 'text-green-500' : 'text-red-500'}`}>
                                        {passwordValidation.uppercase ? '✔' : '✘'} At least one uppercase letter
                                    </div>
                                    <div className={`text-sm ${passwordValidation.number ? 'text-green-500' : 'text-red-500'}`}>
                                        {passwordValidation.number ? '✔' : '✘'} At least one number
                                    </div>
                                    <div className={`text-sm ${passwordValidation.specialChar ? 'text-green-500' : 'text-red-500'}`}>
                                        {passwordValidation.specialChar ? '✔' : '✘'} At least one special character (@$!%*?&)
                                    </div>
                                </div>
                            )}
                        </div>
                        <button
                            type="submit"
                            className="w-full py-2 px-4 bg-indigo-600 hover:bg-indigo-700 text-white font-semibold rounded-md shadow-md focus:outline-none focus:ring-2 focus:ring-indigo-500 dark:focus:ring-indigo-400"
                        >
                            Register
                        </button>
                    </form>
                ) : (
                    <form onSubmit={handleVerifyOtp}>
                        <div className="mb-4">
                            <label className="block text-gray-700 dark:text-gray-300">OTP:</label>
                            <input
                                type="text"
                                value={otp}
                                onChange={(e) => setOtp(e.target.value)}
                                className="mt-1 block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 dark:focus:ring-indigo-400 dark:focus:border-indigo-400 text-gray-700 dark:text-gray-300"
                            />
                        </div>
                        <button
                            type="submit"
                            className="w-full py-2 px-4 bg-indigo-600 hover:bg-indigo-700 text-white font-semibold rounded-md shadow-md focus:outline-none focus:ring-2 focus:ring-indigo-500 dark:focus:ring-indigo-400"
                        >
                            Verify OTP
                        </button>
                    </form>
                )}
            </div>
        </div>
    );
};

export default Signup;
