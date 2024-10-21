import React from 'react';
import { Link } from 'react-router-dom';

const Navbar: React.FC = () => {
  return (
    <nav style={styles.navbar}>
      <ul style={styles.navList}>
        <li style={styles.navItem}>
          <Link to="/" style={styles.navLink}>Home</Link>
        </li>
        <li style={styles.navItem}>
          <Link to="/about" style={styles.navLink}>About</Link>
        </li>
        <li style={styles.navItem}>
          <Link to="/contact" style={styles.navLink}>Contassssssssst</Link>
        </li>
      </ul>
    </nav>
  );
};

const styles = {
  navbar: {
    backgroundColor: '#333',
    padding: '1rem',
  },
  navList: {
    listStyle: 'none',
    display: 'flex',
    justifyContent: 'space-around',
    margin: 0,
    padding: 0,
  },
  navItem: {
    margin: '0 1rem',
  },
  navLink: {
    color: '#fff',
    textDecoration: 'none',
  },
};

export default Navbar;