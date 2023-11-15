import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import Typography from "@mui/joy/Typography";
import Input from "@mui/joy/Input";
import Button from "@mui/joy/Button";
import Alert from "@mui/joy/Alert";
import { motion } from "framer-motion";

function Login() {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const [emailError, setEmailError] = useState(false);

  const isFormFilled = Object.values(formData).every(
    (value) => value.trim() !== ""
  );

  const isEmailValid = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i;

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });

    if (name === "email") {
      const isValidEmail = isEmailValid.test(value);
      setEmailError(!isValidEmail);
    }
  };

  const handleSignIn = () => {
    if (emailError) {
      return;
    }

    console.log("Sign In Form Data:", formData);

    // Perform sign-in actions here
    // For example, send data to a server for authentication
  };

  return (
    <div className="wrapper">
      <nav className="flex items-center justify-center pt-4">
        <a href="">
          <Typography className="logo-main" level="h3">
            CloudShare<span className="x-home">X</span>
          </Typography>
        </a>
      </nav>
      <div className="flex items-center justify-center h-screen pb-6">
        <div className="login flex flex-col gap-6 items-center">
          <div className="login-text-field text-center flex flex-col gap-1">
            <Typography className="logo-main" level="h1">
              Sign in to your account
            </Typography>
            <Typography className="grey" level="body-s">
              Hello, Lets login to your account
            </Typography>
          </div>
          <div className="max-w-xs w-full ">
            <div className="login-form-field  gap-3 items-center">
              {emailError && (
                <motion.div
                  initial={{ opacity: 0 }}
                  animate={{ opacity: 1 }}
                  exit={{ opacity: 0 }}
                  transition={{ duration: 0.5 }}
                  className="mb-4"
                >
                  <Alert color="danger">Please enter valid email.</Alert>
                </motion.div>
              )}
              <Input
                name="email"
                value={formData.email}
                error={emailError}
                placeholder="Email"
                className="field mb-4"
                onChange={handleChange}
              />
              <Input
                type="password"
                placeholder="Password"
                name="password"
                onChange={handleChange}
                value={formData.password}
              />
            </div>
          </div>
          <div className="login-field-button max-w-xs w-full">
            <button
              className={`text-white font-bold py-3 px-4 rounded-lg w-full button-in-login ${
                !isFormFilled || emailError ? "disabled-button" : ""
              }`}
              disabled={!isFormFilled || emailError}
              onClick={handleSignIn}
            >
              Sign In
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Login;
