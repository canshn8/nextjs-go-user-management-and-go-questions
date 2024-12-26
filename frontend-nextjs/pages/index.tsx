import { useRouter } from 'next/router'; // Import the useRouter hook from Next.js to enable navigation

const HomePage: React.FC = () => {
  const router = useRouter(); 

  // Function to navigate to the Users page
  const navigateToUsers = () => {
    router.push('/users'); // `router.push()` is used to navigate programmatically to another page (Users page in this case)
  };

  return (
    <div>
      <h1>Welcome to the Home Page</h1> {/* Display the title of the page */}
      {/* Button that triggers navigation to the Users page when clicked */}
      <button onClick={navigateToUsers}>Go to Users Page</button>
    </div>
  );
};

export default HomePage; // Export the HomePage component so it can be used in other parts of the app
