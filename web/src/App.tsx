import { useEffect } from 'react';
import React from 'react'
import { Link, useNavigate } from 'react-router-dom';
import { MoonIcon, SunIcon, UserIcon, UserPlusIcon, MenuIcon } from 'lucide-react'


function App() {
  const navigate = useNavigate();
  const [isMenuOpen, setIsMenuOpen] = React.useState(false)



  useEffect(() => {
    const accessToken = localStorage.getItem('token');
    if (accessToken) {
      navigate('/home');
    } else {
      navigate('/');
    }
  }, [navigate]);

  return (
    <div className="min-h-screen bg-gradient-to-b from-purple-900 to-indigo-900 text-white">
      {/* Navbar */}
      <nav className="fixed top-0 left-0 right-0 z-50 bg-transparent">
        <div className="container mx-auto px-4 py-4">
          <div className="flex justify-between items-center">
            <div className="flex items-center space-x-2">
              <img src="/bats.svg" alt="Bats Logo" className="w-8 h-8 sm:w-10 sm:h-10" />
              <span className="text-xl sm:text-2xl font-bold">Bats</span>
            </div>
            <div className="hidden sm:flex space-x-4">
                <button
                className="text-white hover:bg-white hover:bg-opacity-20 px-3 py-2 rounded-md text-sm font-medium flex items-center"
                onClick={() => navigate('/signin')}
                >
                <UserIcon className="w-4 h-4 mr-2" />
                Login
                </button>
                <button
                className="text-white hover:bg-white hover:bg-opacity-20 px-3 py-2 rounded-md text-sm font-medium flex items-center"
                onClick={() => navigate('/signup')}
                >
                <UserPlusIcon className="w-4 h-4 mr-2" />
                Sign Up
                </button>
            </div>
            <button className="sm:hidden" onClick={() => setIsMenuOpen(!isMenuOpen)}>
              <MenuIcon className="w-6 h-6" />
            </button>
          </div>
          {isMenuOpen && (
            <div className="mt-4 sm:hidden">
              <button
                className="text-white hover:bg-white hover:bg-opacity-20 px-3 py-2 rounded-md text-sm font-medium flex items-center w-full mb-2"
                onClick={() => {
                  setIsMenuOpen(false);
                  navigate('/signin');
                }}
              >
                <UserIcon className="w-4 h-4 mr-2" />
                Login
              </button>
              <button
                className="text-white hover:bg-white hover:bg-opacity-20 px-3 py-2 rounded-md text-sm font-medium flex items-center w-full"
                onClick={() => {
                  setIsMenuOpen(false);
                  navigate('/signup');
                }}
              >
                <UserPlusIcon className="w-4 h-4 mr-2" />
                Sign Up
              </button>
            </div>
          )}
        </div>
      </nav>

      {/* Hero Section */}
      <section className="pt-32 pb-20 px-4">
        <div className="container mx-auto text-center">
          <img src="/bats.svg" alt="Bats Logo" className="w-24 h-24 sm:w-32 sm:h-32 mx-auto mb-8" />
            <h1 className="text-4xl sm:text-5xl font-bold mb-6">
            Welcome to <span className="text-transparent bg-clip-text bg-gradient-to-r from-purple-400 to-indigo-400">Bats</span>
            </h1>
          <p className="text-lg sm:text-xl mb-8">The social media app that comes alive when the sun goes down</p>
          <div className="flex justify-center items-center space-x-4 mb-12">
            <SunIcon className="w-6 h-6 sm:w-8 sm:h-8 text-yellow-400" />
            <div className="text-xl sm:text-2xl font-semibold">6 PM - 6 AM</div>
            <MoonIcon className="w-6 h-6 sm:w-8 sm:h-8 text-blue-200" />
          </div>
            <Link to="/signup" className="bg-gradient-to-r from-purple-500 via-pink-500 to-red-500 text-white hover:from-purple-600 hover:via-pink-600 hover:to-red-600 px-8 py-4 rounded-full text-lg font-medium shadow-xl transform transition-transform duration-300 hover:scale-110 hover:rotate-1">
            Join the Night
            </Link>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-20 px-4 bg-purple-800 bg-opacity-50">
        <div className="container mx-auto">
          <h2 className="text-2xl sm:text-3xl font-bold text-center mb-12">Why Bats?</h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
            <FeatureCard
              icon={<MoonIcon className="w-10 h-10 sm:w-12 sm:h-12 mb-4 text-yellow-400" />}
              title="Night-time Exclusivity"
              description="Connect with fellow night owls when the rest of the world sleeps."
            />
            <FeatureCard
              icon={<UserIcon className="w-10 h-10 sm:w-12 sm:h-12 mb-4 text-blue-300" />}
              title="Unique Interactions"
              description="Engage in conversations and activities tailored for the night."
            />
            <FeatureCard
              icon={<SunIcon className="w-10 h-10 sm:w-12 sm:h-12 mb-4 text-orange-400" />}
              title="Daily Reset"
              description="Start fresh every evening with new content and challenges."
            />
          </div>
        </div>
      </section>

      {/* Call to Action */}
      <section className="py-20 px-4">
        <div className="container mx-auto text-center">
          <h2 className="text-2xl sm:text-3xl font-bold mb-6">Ready to embrace the night?</h2>
          <p className="text-lg sm:text-xl mb-8">Join Bats and discover a world that comes alive after dark.</p>
            <Link to="/signup" className="bg-gradient-to-r from-purple-500 via-pink-500 to-red-500 text-white hover:from-purple-600 hover:via-pink-600 hover:to-red-600 px-8 py-4 rounded-full text-lg font-medium shadow-xl transform transition-transform duration-300 hover:scale-110 hover:rotate-1">
            Get Started
            </Link>
        </div>
      </section>
    </div>
  )
}

interface FeatureCardProps {
  icon: React.ReactNode;
  title: string;
  description: string;
}

function FeatureCard({ icon, title, description }: FeatureCardProps) {
  return (
    <div className="bg-purple-700 bg-opacity-50 p-6 rounded-lg text-center">
      <div className="flex justify-center">{icon}</div>
      <h3 className="text-lg sm:text-xl font-semibold mb-2">{title}</h3>
      <p className="text-sm sm:text-base">{description}</p>
    </div>
  )
}

export default App