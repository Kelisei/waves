import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Navigation from "./components/Navigation"; // Import Navigation

// Example Components
import Home from "./pages/Home";
import About from "./pages/About";
import NotFound from "./pages/NotFound";

const App: React.FC = () => {
  return (
    <Router>
      <div
        style={{
          backgroundImage:
            "radial-gradient(circle farthest-corner at 12.3% 19.3%, rgba(85,88,218,1) 0%, rgba(95,209,249,1) 100.2%)",
        }}
      >
        <Navigation />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
