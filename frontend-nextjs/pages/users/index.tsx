import React, { useEffect, useState } from 'react'; 
import axios from 'axios';

interface User {
  id: number;
  name: string;
  age: number;
}

const UsersPage: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]); // State to store the list of users

  // useEffect hook runs when the component mounts (first load)
  useEffect(() => {
    // Make an HTTP GET request to fetch users from the backend API
    axios.get('http://localhost:8080/users') 
      .then(response => setUsers(response.data)) // Set users in state if request is successful
      .catch(error => console.error('Error fetching users:', error)); // Log error if the request fails
  }, []); // Empty dependency array ensures this effect runs only once, after the component mounts

  return (
    <div>
      <h1>User List</h1> {/* Page header */}
      <ul>
        {/* Loop through the users array and display each user's name and age */}
        {users.map(user => (
          <li key={user.id}>{user.name} - {user.age}</li> // Display each user's name and age in a list item
        ))}
      </ul>
    </div>
  );
};

export default UsersPage; // Export the UsersPage component for use in other parts of the app
