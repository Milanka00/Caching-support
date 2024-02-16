import React, { useState, useEffect } from 'react';

function App() {
  const [albums, setAlbums] = useState([]);

  useEffect(() => {
    async function fetchAlbums() {
      const response = await fetch('http://localhost:10000/albums');
      const data = await response.json();
      console.log('Cache-Control:', response.headers.get('Cache-Control'));
      setAlbums(data);
    }
    fetchAlbums();
    
  }, []);

  return (
    <div>
      <h1>Albums</h1>
      <ul>
        {albums.map(album => (
          <li key={album.id}>
            {album.title} by {album.artist} - ${album.price}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
