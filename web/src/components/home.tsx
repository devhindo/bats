import { useEffect, useState } from 'react'
import { Settings, LogOut, User, HomeIcon, UserCircle } from 'lucide-react'
import axios from 'axios'
import useAuth from './verifyAuth'

interface Comment {
  author: string;
  createdAt: string;
  content: string;
  likes: number;
}

interface Post {
  username: string;
  createdAt: string;
  content: string;
  likes: number;
  commentsNumber: number;
  comments: Comment[];
}

export default function Home() {
  const url = import.meta.env.VITE_BASE_URL;
  const [isSettingsOpen, setIsSettingsOpen] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const isLoggedIn = useAuth();
  return (
    <div>
        {isLoggedIn ? (
            <p>User is logged in</p>
        ) : (
            <p>User is not logged in</p>
        )}
    </div>
  );

  const toggleSettings = () => {
    setIsSettingsOpen(!isSettingsOpen)
  }

  const [posts, setPosts] = useState<Post[]>([])

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await axios.get(url + 'home', { withCredentials: true })
        console.log(response.status)
        if (response.status === 200) {
          setPosts(response.data)
          console.log("dtaaaaaa")
        } else {
          console.log("stttstxt" + response.statusText)
          setError('Error fetching posts: ' + response.statusText)
        }
      } catch (error) {
        console.error('Error fetching posts:', error)
        if (axios.isAxiosError(error)) {
          if (!error.response) {
            setError('Server is offline. Please try again later.')
          } else if (error.response.status === 401) {
            setError('You are not authenticated. Please log in.')
          } else {
            setError('Error fetching posts: ' + error.message)
          }
        } else {
          setError('An unexpected error occurred: ' + (error as Error).message)
        }
      }
    }

    fetchPosts()
  }, [url]) // Add url as a dependency

  if (error) {
    if (error === 'You are not authenticated. Please log in.') {
      return (
        <div className="min-h-screen flex items-center justify-center bg-black text-gray-200">
          <div className="text-center">
            <p>{error}</p>
            <button
              onClick={() => window.location.href = '/signin'}
              className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
            >
              Go to Sign In
            </button>
          </div>
        </div>
      )
    } else if (error === 'Server is offline. Please try again later.') {
      return (
        <div className="min-h-screen flex items-center justify-center bg-black text-gray-200">
          <div className="text-center">
            <p>{error}</p>
          </div>
        </div>
      )
    } else {
      return (
        <div className="min-h-screen flex items-center justify-center bg-black text-gray-200">
          <div className="text-center">
            <p>{error}</p>
          </div>
        </div>
      )
    }
  }

  console.log(posts)

  return (
    <div className={`min-h-screen dark`}>
      <div className="bg-black text-gray-200 text-sm sm:text-base">
        {/* Main Content */}
        <main className="pb-16"> {/* Add padding to bottom to account for mobile nav */}
            <div className="fixed top-4 left-4 z-50">
            <img src="/bats.svg" alt="Bats Logo" className="w-6 h-6 sm:w-10 sm:h-10" />
            </div>
            {/* Settings Button and Dropdown */}
          <div className="relative">
            <button
              onClick={toggleSettings}
              className="fixed top-4 right-4 z-50 p-1 sm:p-2 rounded-full bg-gray-800 hover:bg-gray-700 transition-colors"
            >
              <Settings className="h-4 w-4 sm:h-6 sm:w-6 text-gray-200" />
            </button>
            {isSettingsOpen && (
              <div className="absolute top-14 right-4 w-48 sm:w-56 bg-gray-900 rounded-md shadow-lg z-50">
                <button
                  onClick={() => console.log("Profile settings clicked")}
                  className="w-full text-left px-2 sm:px-4 py-1 sm:py-2 hover:bg-gray-700 flex items-center text-gray-200"
                >
                  <User className="mr-1 sm:mr-2 h-3 sm:h-4 w-3 sm:w-4" />
                  <span>Profile settings</span>
                </button>
                <button
                  onClick={
                    () => {console.log("Log out clicked")
                    localStorage.removeItem('token')
                    window.location.href = '/'
                    }}
                  className="w-full text-left px-2 sm:px-4 py-1 sm:py-2 hover:bg-gray-700 flex items-center text-gray-200"
                >
                  <LogOut className="mr-1 sm:mr-2 h-3 sm:h-4 w-3 sm:w-4" />
                  <span>Log out</span>
                </button>
              </div>
            )}
          </div>

          {/* Feed */}
          <div className="max-w-2xl mx-auto">
            {/* New Post Input */}
            <div className="p-2 sm:p-4 border-b border-gray-800">
              <div className="flex gap-2 sm:gap-4">
                <UserCircle className="w-6 h-6 sm:w-8 sm:h-8 text-gray-200" />
                <input
                  type="text"
                  placeholder="What's happening?"
                  className="flex-1 bg-transparent border-none text-sm sm:text-lg focus:outline-none text-gray-200"
                />
              </div>
              <div className="flex justify-end mt-2 sm:mt-4">
                <button className="px-2 sm:px-4 py-1 sm:py-2 border border-gray-500 text-gray-200 rounded-md hover:bg-gray-700 transition-colors">
                  Post
                </button>
              </div>
            </div>

            {/* Posts */}
            {Array.from({ length: 5 }).map((_, i) => (
              <article key={i} className="p-2 sm:p-4 border-b border-gray-800">
                <div className="flex gap-2 sm:gap-4">
                  <UserCircle className="w-6 h-6 sm:w-8 sm:h-8 text-gray-200" />
                  <div className="flex-1">
                    <div className="flex items-center gap-1 sm:gap-2">
                      <h3 className="font-semibold text-gray-200 text-xs sm:text-sm sm:text-base">User {i + 1}</h3>
                      <span className="text-gray-400 text-xs sm:text-sm">@user{i + 1}</span>
                      <span className="text-gray-400 text-xs sm:text-sm">· 1h</span>
                    </div>
                    <p className="mt-1 sm:mt-2 text-gray-200 text-xs sm:text-sm sm:text-base">This is a sample post content. It can be much longer and may include hashtags, mentions, and links.</p>
                    <div className="mt-2 sm:mt-4 flex gap-2 sm:gap-4 text-gray-400">
                      <button className="flex items-center gap-1 sm:gap-2 hover:text-gray-200 text-xs sm:text-sm sm:text-base">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="h-4 w-4 sm:h-5 sm:w-5"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>
                        <span>20</span>
                      </button>
                      <button className="flex items-center gap-1 sm:gap-2 hover:text-gray-200 text-xs sm:text-sm sm:text-base">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="h-4 w-4 sm:h-5 sm:w-5"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
                        <span>10</span>
                      </button>
                    </div>
                  </div>
                </div>
              </article>
            ))}
          </div>
        </main>

        {/* Mobile Navigation */}
        <nav className="fixed bottom-0 left-0 right-0 bg-black border-t border-gray-800">
          <div className="flex justify-around p-2 sm:p-4">
            <a href="#" className="text-gray-200 hover:text-gray-400">
              <HomeIcon className="h-4 w-4 sm:h-6 sm:w-6" />
            </a>
            <a href="#" className="text-gray-200 hover:text-gray-400">
              <User className="h-4 w-4 sm:h-6 sm:w-6" />
            </a>
          </div>
        </nav>
      </div>
    </div>
  )
}
