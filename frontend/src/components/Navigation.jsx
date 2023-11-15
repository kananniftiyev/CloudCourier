import React, { useState } from "react";
import Typography from "@mui/joy/Typography";
import Button from "@mui/joy/Button";
import { useNavigate } from "react-router-dom";

function Navigation() {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate("/auth");
  };
  return (
    <header>
      <div className="md:container md:mx-auto pt-4">
        <div className="flex justify-between">
          <div className="logo">
            <Typography className="logo-main" level="h3">
              CloudShare<span className="x-home">X</span>
            </Typography>
          </div>
          <div className="auth flex justify-between gap-6 items-center">
            <Button className="button-signup" size="lg" onClick={handleClick}>
              Get Started
            </Button>
          </div>
        </div>
      </div>
    </header>
  );
}

export default Navigation;
