// pages/index.tsx
import { useRouter } from 'next/router';

const HomePage: React.FC = () => {
  const router = useRouter();

  const navigateToUsers = () => {
    router.push('/users');
  };

  return (
    <div>
      <h1>Welcome to the Home Page</h1>
      <button onClick={navigateToUsers}>Go to Users Page</button>
    </div>
  );
};

export default HomePage;
