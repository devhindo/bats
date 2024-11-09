import React, { useState } from 'react'
import { Moon, User, Lock, ArrowRight } from 'lucide-react'
import { Link } from 'react-router-dom'

export default function Signin() {
  const [identifier, setIdentifier] = useState('')
  const [password, setPassword] = useState('')

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // Here you would typically handle the login logic
    console.log('Login attempt with:', { identifier, password })
  }

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
            <h2 className="mt-4 text-2xl sm:text-3xl font-extrabold text-white">Welcome to Bats</h2>
            <p className="mt-2 text-sm text-purple-200">Sign in to your account</p>
          </div>
          <form className="mt-6 space-y-4 sm:space-y-6" onSubmit={handleSubmit}>
            <div className="rounded-md shadow-sm -space-y-px">
              <div>
                <label htmlFor="identifier" className="sr-only">Email or Username</label>
                <div className="relative">
                  <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <User className="h-5 w-5 text-purple-400" />
                  </div>
                  <input
                    id="identifier"
                    name="identifier"
                    type="text"
                    required
                    className="appearance-none rounded-none relative block w-full px-3 py-2 pl-10 border border-purple-500 placeholder-purple-400 text-white bg-purple-700 bg-opacity-50 rounded-t-md focus:outline-none focus:ring-2 focus:ring-purple-400 focus:border-purple-400 focus:z-10 text-sm sm:text-base"
                    placeholder="Email or Username"
                    value={identifier}
                    onChange={(e) => setIdentifier(e.target.value)}
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
                    onChange={(e) => setPassword(e.target.value)}
                  />
                </div>
              </div>
            </div>

            <div className="flex items-center justify-end">
              <div className="text-sm">
                <a href="#" className="font-medium text-purple-300 hover:text-purple-200">
                  Forgot password?
                </a>
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
                Sign in
              </button>
            </div>
          </form>
          <div className="text-center">
            <p className="mt-2 text-sm text-purple-200">
              Don't have an account?{' '}
              <Link to="/signup" className="font-medium text-purple-300 hover:text-purple-200">
                Sign up
              </Link>
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}