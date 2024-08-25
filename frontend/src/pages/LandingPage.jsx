import { Link } from "react-router-dom";
import React, { useEffect, useState } from "react";

const LandingPage = () => {
  const [message, setMessage] = useState("");
  useEffect(() => {
    fetch("http://localhost:8081/api/hello")
      .then((response) => response.text())
      .then((data) => setMessage(data))
      .catch((error) => console.error("Error fetching data:", error));
  }, []);

  return (
    <div>
      <div className="App">
        <h1>{message}</h1>
      </div>
      <div>
        <h1>Welcome to Our App</h1>
        <button>
          <Link to="/register">Register</Link>
        </button>
      </div>
    </div>
  );
};

export default LandingPage;
