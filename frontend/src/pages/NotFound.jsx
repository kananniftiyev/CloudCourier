import React from "react";
import Typography from "@mui/joy/Typography";
import { useNavigate } from "react-router-dom";

const NotFound = () => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate("/home");
  };
  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <span role="img" aria-label="Lost" style={{ fontSize: "4em" }}>
        🌍
      </span>
      <Typography level="h1" className="text-5xl grey font-bold mb-4">
        Oops! 404 Not Found
      </Typography>
      <Typography level="body-lg" className="grey text-2xl">
        Looks like you've wandered into uncharted territory. Let's get you back
        on track!
      </Typography>
      <a href="" onClick={handleClick}>
        <span
          role="img"
          aria-label="Explore"
          style={{ fontSize: "3em", marginTop: "1em" }}
        >
          🚀
        </span>
      </a>
    </div>
  );
};

export default NotFound;
