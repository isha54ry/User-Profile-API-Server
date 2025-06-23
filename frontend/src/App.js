import React, { useState, useEffect } from 'react';

function App() {
  const [users, setUsers] = useState([]);
  const [name, setName] = useState('');
  const [image, setImage] = useState(null);
  const [dark, setDark] = useState(false);

  const fetchUsers = () => {
    fetch('http://localhost:8081/users')
      .then(res => res.json())
      .then(data => setUsers(Object.values(data)));
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('name', name);
    formData.append('image', image);
    await fetch('http://localhost:8081/users', {
      method: 'POST',
      body: formData
    });
    setName('');
    setImage(null);
    fetchUsers();
  };

  const handleDelete = async (id) => {
    await fetch(`http://localhost:8081/users/${id}`, { method: 'DELETE' });
    fetchUsers();
  };

  const handleDownload = (id) => {
    window.open(`http://localhost:8081/profile/${id}`, '_blank');
  };

  return (
    <div className={`${dark ? 'bg-gray-900' : 'bg-gradient-to-br from-purple-100 to-blue-200'} min-h-screen p-8 transition-colors duration-300 relative overflow-hidden`}>
      {/* Background doodles - enhanced */}
      <svg
        className="absolute top-0 left-0 opacity-20 pointer-events-none"
        width="500"
        height="500"
        viewBox="0 0 500 500"
        fill="none"
        style={{ zIndex: 0 }}
      >
        {/* Large swirl */}
        <path
          d="M250,250 m-200,0 a200,200 0 1,0 400,0 a200,200 0 1,0 -400,0"
          stroke={dark ? "#a78bfa" : "#c4b5fd"}
          strokeWidth="10"
          fill="none"
        />
        {/* Zigzag */}
        <polyline
          points="60,400 100,350 140,400 180,350 220,400"
          stroke={dark ? "#818cf8" : "#6366f1"}
          strokeWidth="7"
          fill="none"
        />
        {/* Dotted circle */}
        <circle
          cx="400"
          cy="100"
          r="60"
          stroke={dark ? "#f472b6" : "#f9a8d4"}
          strokeWidth="7"
          fill="none"
          strokeDasharray="10,10"
        />
        {/* Star */}
        <polygon
          points="120,80 130,110 160,110 135,125 145,155 120,135 95,155 105,125 80,110 110,110"
          stroke={dark ? "#34d399" : "#6ee7b7"}
          strokeWidth="5"
          fill="none"
        />
      </svg>
      <svg
        className="absolute bottom-0 right-0 opacity-20 pointer-events-none"
        width="400"
        height="400"
        viewBox="0 0 400 400"
        fill="none"
        style={{ zIndex: 0 }}
      >
        {/* Concentric circles */}
        <circle cx="320" cy="320" r="60" stroke={dark ? "#f87171" : "#fca5a5"} strokeWidth="7" fill="none" />
        <circle cx="320" cy="320" r="30" stroke={dark ? "#fbbf24" : "#fde68a"} strokeWidth="4" fill="none" />
        {/* Squiggle */}
        <path
          d="M60 340 Q100 300 140 340 T220 340"
          stroke={dark ? "#818cf8" : "#6366f1"}
          strokeWidth="6"
          fill="none"
        />
        {/* Dotted ellipse */}
        <ellipse
          cx="80"
          cy="80"
          rx="40"
          ry="20"
          stroke={dark ? "#34d399" : "#6ee7b7"}
          strokeWidth="5"
          fill="none"
          strokeDasharray="8,8"
        />
        {/* Triangle */}
        <polygon
          points="300,60 340,120 260,120"
          stroke={dark ? "#f472b6" : "#f9a8d4"}
          strokeWidth="5"
          fill="none"
        />
      </svg>
      <div className="max-w-3xl mx-auto relative z-10">
        <div className="flex justify-between items-center mb-8">
          <div className="flex items-center gap-3">
            {/* User icon */}
            <svg width="40" height="40" viewBox="0 0 24 24" fill="none" className={`${dark ? 'text-purple-300' : 'text-purple-700'}`}>
              <circle cx="12" cy="8" r="4" stroke="currentColor" strokeWidth="2" />
              <path d="M4 20c0-3.3137 3.134-6 7-6s7 2.6863 7 6" stroke="currentColor" strokeWidth="2" />
            </svg>
            <h1 className={`${dark ? 'text-purple-300' : 'text-purple-700'} text-4xl font-bold text-center flex-1`}>User Manager Dashboard</h1>
          </div>
          <button
            onClick={() => setDark(!dark)}
            className={`ml-4 px-4 py-2 rounded transition-colors duration-300 ${dark ? 'bg-purple-700 text-white' : 'bg-gray-200 text-gray-800 hover:bg-gray-300'}`}
          >
            {dark ? (
              // Sun icon
              <span className="flex items-center gap-2">
                <svg width="20" height="20" fill="none" viewBox="0 0 24 24" className="inline-block"><circle cx="12" cy="12" r="5" stroke="currentColor" strokeWidth="2"/><path stroke="currentColor" strokeWidth="2" d="M12 1v2M12 21v2M4.22 4.22l1.42 1.42M18.36 18.36l1.42 1.42M1 12h2M21 12h2M4.22 19.78l1.42-1.42M18.36 5.64l1.42-1.42"/></svg>
                Light Mode
              </span>
            ) : (
              // Moon icon
              <span className="flex items-center gap-2">
                <svg width="20" height="20" fill="none" viewBox="0 0 24 24" className="inline-block"><path d="M21 12.79A9 9 0 1111.21 3a7 7 0 109.79 9.79z" stroke="currentColor" strokeWidth="2" /></svg>
                Dark Mode
              </span>
            )}
          </button>
        </div>

        <form onSubmit={handleSubmit} className={`${dark ? 'bg-gray-800' : 'bg-white'} rounded-xl shadow-lg p-6 mb-10 space-y-4`}>
          <div className="flex items-center gap-3">
            {/* Add user graphic */}
            <svg width="32" height="32" fill="none" viewBox="0 0 24 24" className={`${dark ? 'text-purple-400' : 'text-purple-600'}`}>
              <circle cx="12" cy="8" r="4" stroke="currentColor" strokeWidth="2"/>
              <path d="M4 20c0-3.3137 3.134-6 7-6s7 2.6863 7 6" stroke="currentColor" strokeWidth="2"/>
              <path d="M19 7v4M21 9h-4" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/>
            </svg>
            <span className={`${dark ? 'text-purple-200' : 'text-purple-700'} font-semibold text-lg`}>Add New User</span>
          </div>
          <input
            type="text"
            placeholder="Enter name"
            className={`w-full p-2 border rounded ${dark ? 'bg-gray-700 text-white border-gray-600' : ''}`}
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <input
            type="file"
            onChange={(e) => setImage(e.target.files[0])}
            required
            className={`${dark ? 'text-white' : ''}`}
          />
          <button type="submit" className="px-4 py-2 bg-purple-600 text-white rounded hover:bg-purple-700 flex items-center gap-2">
            {/* Plus icon */}
            <svg width="18" height="18" fill="none" viewBox="0 0 24 24"><path d="M12 5v14M5 12h14" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/></svg>
            Add User
          </button>
        </form>

        <div className="grid gap-6">
          {users.map(user => (
            <div key={user.ID} className={`${dark ? 'bg-gray-800' : 'bg-white'} p-5 rounded-xl shadow-md flex items-center justify-between`}>
              <div className="flex items-center gap-4">
                {/* Avatar graphic */}
                <svg width="40" height="40" fill="none" viewBox="0 0 24 24" className={`${dark ? 'text-purple-400' : 'text-purple-600'}`}>
                  <circle cx="12" cy="8" r="4" stroke="currentColor" strokeWidth="2"/>
                  <path d="M4 20c0-3.3137 3.134-6 7-6s7 2.6863 7 6" stroke="currentColor" strokeWidth="2"/>
                </svg>
                <div>
                  <h2 className={`text-xl font-semibold ${dark ? 'text-gray-100' : 'text-gray-800'}`}>{user.Name}</h2>
                  <p className={`text-sm ${dark ? 'text-gray-400' : 'text-gray-500'}`}>{user.ID}</p>
                </div>
              </div>
              <div className="flex gap-2">
                <button onClick={() => handleDownload(user.ID)} className="px-3 py-1 bg-green-500 text-white rounded hover:bg-green-600 flex items-center gap-1">
                  {/* PDF icon */}
                  <svg width="16" height="16" fill="none" viewBox="0 0 24 24"><path d="M6 2h7l5 5v13a2 2 0 01-2 2H6a2 2 0 01-2-2V4a2 2 0 012-2z" stroke="currentColor" strokeWidth="2"/><path d="M13 2v6h6" stroke="currentColor" strokeWidth="2"/></svg>
                  PDF
                </button>
                <button onClick={() => handleDelete(user.ID)} className="px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600 flex items-center gap-1">
                  {/* Trash icon */}
                  <svg width="16" height="16" fill="none" viewBox="0 0 24 24"><path d="M3 6h18M8 6v12a2 2 0 002 2h4a2 2 0 002-2V6M10 11v6M14 11v6" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/></svg>
                  Delete
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default App;