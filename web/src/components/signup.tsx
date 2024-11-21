import React, { useState, useRef } from 'react';
import axios from 'axios';
import { User, Mail, Lock, ArrowRight } from 'lucide-react';
import { Link } from 'react-router-dom';

const url = import.meta.env.VITE_BASE_URL;

const Signup: React.FC = () => {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [otp, setOtp] = useState(Array(6).fill(''));
    const [isOtpSent, setIsOtpSent] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');
    const [otpErrorMessage, setOtpErrorMessage] = useState('');
    const [passwordValidation, setPasswordValidation] = useState({
        length: false,
        uppercase: false,
        number: false,
        specialChar: false,
    });
    const [isPasswordTouched, setIsPasswordTouched] = useState(false);
    const [showOtpPopup, setShowOtpPopup] = useState(false);

    const otpRefs = useRef<(HTMLInputElement | null)[]>([]);

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
            return;
        }
        try {
            const response = await axios.post(url + 'signup', {
                username,
                email,
                password,
            });
            if (response.status === 201 && response.data.message === 'otp sent to email') {
                setIsOtpSent(true);
                setErrorMessage('');
                setShowOtpPopup(true);
            }
        } catch (error) {
            if (axios.isAxiosError(error) && error.response?.status === 409) {
                setErrorMessage('Username already exists');
            } else {
                console.error('There was an error registering!', error);
            }
        }
    };

    const handleOtpChange = (index: number) => (e: React.ChangeEvent<HTMLInputElement>) => {
        const newOtp = [...otp];
        const value = e.target.value;

        if (value === '' && (e.nativeEvent as InputEvent).inputType === 'deleteContentBackward') {
            newOtp[index] = '';
            setOtp(newOtp);
            if (index > 0) {
                otpRefs.current[index - 1]?.focus();
            }
        } else if (value.length === 1) {
            newOtp[index] = value;
            setOtp(newOtp);
            if (index < otpRefs.current.length - 1) {
                otpRefs.current[index + 1]?.focus();
            }
        }
    };

    const handleVerifyOtp = async (event: React.FormEvent) => {
        event.preventDefault();
        try {
            const response = await axios.post(url + 'signup/otp', {
                email,
                otp: otp.join(''),
            });
            window.location.href = '/home'; // Redirect to /home on successful registration
        } catch (error) {
            if (axios.isAxiosError(error) && error.response?.status === 401) {
                setOtpErrorMessage('Incorrect OTP. Please try again.');
            } else {
                console.error('There was an error verifying OTP!', error);
            }
        }
    };

    return (
        <div className="min-h-screen bg-gradient-to-b from-purple-900 to-indigo-900 flex flex-col">
            <nav className="bg-transparent p-2">
                <div className="container mx-auto flex justify-between items-center">
                    <Link to="/" className="flex items-center">
                        <img src="/bats.svg" alt="Bats Logo" className="w-6 h-6 sm:w-8 sm:h-8 mr-2" />
                        <span className="text-white text-sm sm:text-lg font-bold">Bats</span>
                    </Link>
                </div>
            </nav>
            <div className="flex-grow flex items-center justify-center p-4">
                <div className="w-full max-w-sm space-y-6 bg-purple-800 bg-opacity-50 p-6 rounded-xl shadow-2xl">
                    <div className="text-center">
                        <img src="/bats.svg" alt="Bats Logo" className="w-16 h-16 sm:w-20 sm:h-20 mx-auto mb-4" />
                        <h2 className="mt-4 text-2xl sm:text-3xl font-extrabold text-white">Join Bats</h2>
                        <p className="mt-2 text-sm text-purple-200">Create your account</p>
                    </div>
                    {!isOtpSent ? (
                        <form className="mt-6 space-y-4 sm:space-y-6" onSubmit={handleRegister}>
                            {errorMessage && (
                                <div className="mb-4 text-yellow-400">
                                    {errorMessage}
                                </div>
                            )}
                            <div className="rounded-md shadow-sm -space-y-px">
                                <div>
                                    <label htmlFor="username" className="sr-only">Username</label>
                                    <div className="relative">
                                        <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                            <User className="h-5 w-5 text-purple-400" />
                                        </div>
                                        <input
                                            id="username"
                                            name="username"
                                            type="text"
                                            required
                                            className="appearance-none rounded-none relative block w-full px-3 py-2 pl-10 border border-purple-500 placeholder-purple-400 text-white bg-purple-700 bg-opacity-50 rounded-t-md focus:outline-none focus:ring-2 focus:ring-purple-400 focus:border-purple-400 focus:z-10 text-sm sm:text-base"
                                            placeholder="Username"
                                            value={username}
                                            onChange={(e) => setUsername(e.target.value)}
                                        />
                                    </div>
                                </div>
                                <div>
                                    <label htmlFor="email" className="sr-only">Email address</label>
                                    <div className="relative">
                                        <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                            <Mail className="h-5 w-5 text-purple-400" />
                                        </div>
                                        <input
                                            id="email"
                                            name="email"
                                            type="email"
                                            required
                                            className="appearance-none rounded-none relative block w-full px-3 py-2 pl-10 border border-purple-500 placeholder-purple-400 text-white bg-purple-700 bg-opacity-50 focus:outline-none focus:ring-2 focus:ring-purple-400 focus:border-purple-400 focus:z-10 text-sm sm:text-base focus:bg-purple-700"
                                            placeholder="Email address"
                                            value={email}
                                            onChange={(e) => setEmail(e.target.value)}
                                        />
                                    </div>
                                </div>
                                <div>
                                    <label htmlFor="password" className="sr-only">Password</label>
                                    <div className="relative">
                                        <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                            <Lock className="h-5 w-5 text-purple-400" />
                                        </div>
                                        <input
                                            id="password"
                                            name="password"
                                            type="password"
                                            required
                                            className="appearance-none rounded-none relative block w-full px-3 py-2 pl-10 border border-purple-500 placeholder-purple-400 text-white bg-purple-700 bg-opacity-50 rounded-b-md focus:outline-none focus:ring-2 focus:ring-purple-400 focus:border-purple-400 focus:z-10 text-sm sm:text-base"
                                            placeholder="Password"
                                            value={password}
                                            onChange={handlePasswordChange}
                                        />
                                    </div>
                                    {isPasswordTouched && (
                                        <div className="mt-2">
                                            <div className={`text-sm ${passwordValidation.length ? 'text-green-400' : 'text-yellow-400'}`}>
                                                {passwordValidation.length ? '✔' : '✘'} At least 8 characters
                                            </div>
                                            <div className={`text-sm ${passwordValidation.uppercase ? 'text-green-400' : 'text-yellow-400'}`}>
                                                {passwordValidation.uppercase ? '✔' : '✘'} At least one uppercase letter
                                            </div>
                                            <div className={`text-sm ${passwordValidation.number ? 'text-green-400' : 'text-yellow-400'}`}>
                                                {passwordValidation.number ? '✔' : '✘'} At least one number
                                            </div>
                                            <div className={`text-sm ${passwordValidation.specialChar ? 'text-green-400' : 'text-yellow-400'}`}>
                                                {passwordValidation.specialChar ? '✔' : '✘'} At least one special character (@$!%*?&)
                                            </div>
                                        </div>
                                    )}
                                </div>
                            </div>
                            <div>
                                <button
                                    type="submit"
                                    className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm sm:text-base font-medium rounded-md text-white bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
                                    disabled={!passwordValidation.length || !passwordValidation.uppercase || !passwordValidation.number || !passwordValidation.specialChar}
                                >
                                    <span className="absolute left-0 inset-y-0 flex items-center pl-3">
                                        <ArrowRight className="h-5 w-5 text-purple-500 group-hover:text-purple-400" aria-hidden="true" />
                                    </span>
                                    Sign up
                                </button>
                            </div>
                        </form>
                    ) : (
                        <form className="mt-6 space-y-4 sm:space-y-6" onSubmit={handleVerifyOtp}>
                            {otpErrorMessage && (
                                <div className="mb-4 text-yellow-400">
                                    {otpErrorMessage}
                                </div>
                            )}
                            <div className="mb-4 text-center">
                                <label className="block text-white">OTP:</label>
                                <div className="flex justify-center space-x-2 mt-2">
                                    {otp.map((value, index) => (
                                        <input
                                            key={index}
                                            type="text"
                                            value={value}
                                            onChange={handleOtpChange(index)}
                                            maxLength={1}
                                            ref={(el) => (otpRefs.current[index] = el)}
                                            className="w-10 px-3 py-2 bg-purple-700 border border-purple-500 rounded-md shadow-sm focus:outline-none focus:ring-purple-400 focus:border-purple-400 text-white text-center"
                                            onKeyDown={(e) => {
                                                if (e.key === 'Backspace' && otp[index] === '') {
                                                    if (index > 0) {
                                                        otpRefs.current[index - 1]?.focus();
                                                    }
                                                }
                                            }}
                                        />
                                    ))}
                                </div>
                            </div>
                            <div>
                                <button
                                    type="submit"
                                    className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm sm:text-base font-medium rounded-md text-white bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
                                >
                                    <span className="absolute left-0 inset-y-0 flex items-center pl-3">
                                        <ArrowRight className="h-5 w-5 text-purple-500 group-hover:text-purple-400" aria-hidden="true" />
                                    </span>
                                    Confirm OTP
                                </button>
                            </div>
                        </form>
                    )}
                    <div className="text-center">
                        <p className="mt-2 text-sm text-purple-200">
                            Already have an account?{' '}
                            <Link to="/signin" className="font-medium text-purple-300 hover:text-purple-200">
                                Sign in
                            </Link>
                        </p>
                    </div>
                </div>
                {showOtpPopup && (
                    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 p-4">
                        <div className="bg-purple-800 p-6 rounded-lg shadow-lg w-full max-w-sm">
                            <p className="text-white">Mail has been sent with the OTP. Please check your mail.</p>
                            <button
                                onClick={() => setShowOtpPopup(false)}
                                className="mt-4 py-2 px-4 bg-purple-600 hover:bg-purple-700 text-white font-semibold rounded-md shadow-md focus:outline-none focus:ring-2 focus:ring-purple-500"
                            >
                                Close
                            </button>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
};

export default Signup;
