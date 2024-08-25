import React from "react";
import { Link } from "react-router-dom";

const RegisterPage = () => {
  return (
    <div>
      <h1>Register</h1>
      {/* Registration form goes here */}
      <button>
        <Link to="/feed">Go to Feed</Link>
      </button>
    </div>
  );
};

export default RegisterPage;
