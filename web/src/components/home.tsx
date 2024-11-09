import React from 'react';

const Home: React.FC = () => {
    return (
        <div className="home-page">
            <header className="home-header">
                <h1>Welcome to Social Media App</h1>
            </header>
            <main className="home-content">
                <section className="posts-section">
                    <h2>Recent Posts</h2>
                    {/* Posts will be rendered here */}
                </section>
                <section className="friends-section">
                    <h2>Friends</h2>
                    {/* Friends list will be rendered here */}
                </section>
            </main>
        </div>
    );
};

export default Home;