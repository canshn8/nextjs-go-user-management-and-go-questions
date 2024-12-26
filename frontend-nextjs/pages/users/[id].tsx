import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';

interface User {
  id: number;
  name: string;
  age: number;
}

const UserDetailPage: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const router = useRouter();
  const { id } = router.query;

  useEffect(() => {
    if (id) {
      axios.get(`http://localhost:8080/users/${id}`)
        .then(response => setUser(response.data))
        .catch(error => console.error('Error fetching user:', error));
    }
  }, [id]);

  if (!user) return <div>Loading...</div>;

  return (
    <div>
      <h1>User Detail</h1>
      <p>ID: {user.id}</p>
      <p>Name: {user.name}</p>
      <p>Age: {user.age}</p>
    </div>
  );
};

export default UserDetailPage;