import axios from 'axios';

export const api = axios.create({
  baseURL: 'http://localhost:8080',
});

export const fetchUsers = () => api.get('/users');
export const fetchUserById = (id: number) => api.get(`/users/${id}`);
export const createUser = (user: { name: string; age: number }) => api.post('/users', user);
export const updateUser = (id: number, user: { name: string; age: number }) => api.put(`/users/${id}`, user);
export const deleteUser = (id: number) => api.delete(`/users/${id}`);