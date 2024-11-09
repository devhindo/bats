import React, { useState, useEffect } from 'react'
import { Heart, MessageCircle, Send } from 'lucide-react'

// Mock data to simulate fetching from /home endpoint
const mockPosts = [
  {
    id: 1,
    username: 'night_owl_1',
    profilePic: '/placeholder.svg?height=40&width=40',
    content: 'Just spotted a beautiful owl outside my window. Night life at its best! 🦉✨',
    likes: 15,
    comments: [
      { username: 'moon_gazer', content: 'Amazing! What species?' },
      { username: 'insomnia_queen', content: 'Lucky you! I wish I could see owls where I live.' }
    ]
  },
  {
    id: 2,
    username: 'midnight_coder',
    profilePic: '/placeholder.svg?height=40&width=40',
    content: 'Late night coding session. Who else is burning the midnight oil? 💻☕',
    likes: 32,
    comments: [
      { username: 'caffeine_addict', content: 'Right here with you! What are you working on?' },
      { username: 'sleepy_dev', content: 'Wish I had your energy. I\'m calling it a night.' }
    ]
  }
]

function HomePage() {
  const [posts, setPosts] = useState([])

  useEffect(() => {
    // Simulating API call to /home endpoint
    setTimeout(() => {
      setPosts(mockPosts)
    }, 1000)
  }, [])

  return (
    <div className="min-h-screen bg-gradient-to-b from-purple-900 to-indigo-900 text-white p-4">
      <header className="text-center mb-8">
        <h1 className="text-3xl font-bold">Bats Feed</h1>
      </header>
      <div className="max-w-2xl mx-auto space-y-6">
        {posts.map((post) => (
          <PostCard key={post.id} post={post} />
        ))}
      </div>
    </div>
  )
}

function PostCard({ post }) {
  const [likes, setLikes] = useState(post.likes)
  const [comments, setComments] = useState(post.comments)
  const [newComment, setNewComment] = useState('')

  const handleLike = () => {
    setLikes(likes + 1)
  }

  const handleComment = (e) => {
    e.preventDefault()
    if (newComment.trim()) {
      setComments([...comments, { username: 'current_user', content: newComment }])
      setNewComment('')
    }
  }

  return (
    <div className="bg-purple-800 bg-opacity-50 rounded-lg p-4 shadow-lg">
      <div className="flex items-center mb-4">
        <img src={post.profilePic} alt={post.username} className="w-10 h-10 rounded-full mr-3" />
        <span className="font-semibold">{post.username}</span>
      </div>
      <p className="mb-4">{post.content}</p>
      <div className="flex items-center space-x-4 mb-4">
        <button onClick={handleLike} className="flex items-center space-x-1 text-pink-400 hover:text-pink-300">
          <Heart size={20} />
          <span>{likes}</span>
        </button>
        <div className="flex items-center space-x-1 text-blue-400">
          <MessageCircle size={20} />
          <span>{comments.length}</span>
        </div>
      </div>
      <div className="space-y-2">
        {comments.map((comment, index) => (
          <div key={index} className="bg-purple-700 bg-opacity-50 rounded p-2">
            <span className="font-semibold mr-2">{comment.username}:</span>
            {comment.content}
          </div>
        ))}
      </div>
      <form onSubmit={handleComment} className="mt-4 flex">
        <input
          type="text"
          value={newComment}
          onChange={(e) => setNewComment(e.target.value)}
          placeholder="Add a comment..."
          className="flex-grow bg-purple-700 bg-opacity-50 rounded-l p-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
        />
        <button type="submit" className="bg-blue-500 hover:bg-blue-600 rounded-r p-2">
          <Send size={20} />
        </button>
      </form>
    </div>
  )
}

export default HomePage