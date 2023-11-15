import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import Typography from "@mui/joy/Typography";
import Input from "@mui/joy/Input";
import Button from "@mui/joy/Button";
import { motion } from "framer-motion";
import Alert from "@mui/joy/Alert";

function Register() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  });

  const [emailError, setEmailError] = useState(false);
  const [passwordError, setPasswordError] = useState(false);

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

    if (name === "confirmPassword" && value !== formData.password) {
      setPasswordError(true);
    } else {
      setPasswordError(false);
    }
  };

  const handleSignUp = () => {
    if (emailError || passwordError) {
      return;
    }

    console.log("Sign Up Form Data:", formData);

    // Perform sign-up actions here
    // For example, send data to a server to create a new account
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
              Create Your Account
            </Typography>
            <Typography className="grey" level="body-s">
              Hello there! Let's create your account.
            </Typography>
          </div>
          <div className="max-w-xs w-full ">
            <div className="login-form-field gap-3 items-center">
              {passwordError && (
                <motion.div
                  initial={{ opacity: 0 }}
                  animate={{ opacity: 1 }}
                  exit={{ opacity: 0 }}
                  transition={{ duration: 0.5 }}
                  className="mb-4"
                >
                  <Alert color="danger">Password is not matching!</Alert>
                </motion.div>
              )}
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
                placeholder="Username"
                name="username"
                value={formData.username}
                onChange={handleChange}
                className="field mb-4"
              />
              <Input
                placeholder="Email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                type="email"
                error={emailError}
                className="field mb-4"
              />
              <Input
                type="password"
                placeholder="Password"
                name="password"
                value={formData.password}
                onChange={handleChange}
                className="field mb-4"
              />
              <Input
                placeholder="Password Again"
                className="field"
                type="password"
                name="confirmPassword"
                value={formData.confirmPassword}
                error={passwordError}
                onChange={handleChange}
              />
            </div>
          </div>
          <div className="login-field-button max-w-xs w-full">
            <button
              className={`text-white font-bold py-3 px-4 rounded-lg w-full button-in-login ${
                !isFormFilled || passwordError || emailError
                  ? "disabled-button"
                  : ""
              }`}
              disabled={!isFormFilled || passwordError || emailError}
            >
              Create my account
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Register;
