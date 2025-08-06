const Navbar = () => {
    return (
        <nav className="navbar">
            <h1>Jijo's blog</h1>
            <div className="links">
                <a href="/">Home</a>
                <a href="/create" style={{ 
          color: 'white', 
          backgroundColor: '#f1356d',
          borderRadius: '8px' 
        }}>Create Blog</a>
            </div>
        </nav>
      );
}
 
export default Navbar;