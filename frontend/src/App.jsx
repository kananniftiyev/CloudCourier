import "./App.css";
import Navigation from "./components/Navigation";
import Footer from "./components/Footer";
import Hero from "./components/Hero";
import FileDrop from "./components/FileDrop";
import "@fontsource/inter";

function App() {
  return (
    <>
      <Navigation />
      <Hero />
      <FileDrop />
      <Footer />
    </>
  );
}

export default App;
