import React, { useState } from 'react';

interface UserFormProps {
  onSubmit: (name: string, age: number) => void;
  initialName?: string;
  initialAge?: number;
}

const UserForm: React.FC<UserFormProps> = ({ onSubmit, initialName = '', initialAge = 0 }) => {
  const [name, setName] = useState(initialName);
  const [age, setAge] = useState(initialAge);

  const handleSubmit = () => {
    onSubmit(name, age);
  };

  return (
    <div>
      <h2>User Form</h2>
      <input
        type="text"
        placeholder="Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <input
        type="number"
        placeholder="Age"
        value={age}
        onChange={(e) => setAge(Number(e.target.value))}
      />
      <button onClick={handleSubmit}>Submit</button>
    </div>
  );
};

export default UserForm;