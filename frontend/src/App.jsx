import "./App.css";
import Navigation from "./components/Navigation";
import Footer from "./components/Footer";
import FileDrop from "./components/FileDrop";
import Auth from "./pages/Auth";
import { BrowserRouter as Router, Route, Link, Routes } from "react-router-dom";
import "@fontsource/uncut-sans/300.css";
import "@fontsource/uncut-sans/400.css";
import "@fontsource/uncut-sans/500.css";
import "@fontsource/uncut-sans/600.css";
import "@fontsource/uncut-sans/700.css";

import Home from "./pages/Home";
function App() {
  return (
    <Router>
      <Routes>
        <Route path="/auth" element={<Auth />} />
        <Route path="/home" element={<Home />} />
      </Routes>
    </Router>
  );
}

export default App;
