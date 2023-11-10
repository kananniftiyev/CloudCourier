import "./App.css";
import Navigation from "./components/Navigation";
import Footer from "./components/Footer";
import FileDrop from "./components/FileDrop";
// Supports weights 300-800
import "@fontsource/source-sans-pro";
function App() {
  return (
    <>
      <Navigation />
      <div className="split md:container md:mx-auto mt-48">
        <FileDrop />
      </div>
      <Footer />
    </>
  );
}

export default App;
