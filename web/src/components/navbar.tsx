// FILE: Navbar.tsx
import React from 'react';
import { Link } from 'react-router-dom';

const Navbar: React.FC = () => {
    return (
        <nav className="bg-transparent p-2">
            <div className="container mx-auto flex justify-between items-center">
                <Link to="/" className="flex items-center">
                    <img src="/bats.svg" alt="Bats Logo" className="w-6 h-6 sm:w-8 sm:h-8 mr-2" />
                    <span className="text-white text-sm sm:text-lg font-bold">Bats</span>
                </Link>
            </div>
        </nav>
    );
};

export default Navbar;