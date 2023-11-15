import React from "react";
import Navigation from "../components/Navigation";
import FileDrop from "../components/FileDrop";
import Footer from "../components/Footer";
function Home() {
  return (
    <div className="home min-h-screen">
      <Navigation />
      <div className="split md:container md:mx-auto mt-48">
        <FileDrop />
      </div>
      <Footer />
    </div>
  );
}

export default Home;
