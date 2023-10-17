import "./App.css";
import Navigation from "./components/Navigation";
import Footer from "./components/Footer";
import Hero from "./components/Hero";
import FileDrop from "./components/FileDrop";
// Supports weights 100-700
import "@fontsource-variable/roboto-mono";
function App() {
  return (
    <>
      <Navigation />
      <div className="split md:container md:mx-auto mt-48">
        <FileDrop />
        <Hero />
      </div>

      <Footer />
    </>
  );
}

export default App;
