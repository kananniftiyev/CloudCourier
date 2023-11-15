import React, { useEffect } from "react";
import { useLocation } from "react-router-dom";

function Login() {
  const location = useLocation();

  useEffect(() => {
    // This function will be called whenever the route changes
    // You can modify the body styles or perform any other actions here
    document.body.style.backgroundColor = "white"; // Change to your desired style

    return () => {
      document.body.style.backgroundColor = ""; // Reset to default or remove the style
    };
  }, [location.pathname]);
}

export default Login;
