import React, { useEffect } from "react";
import { useLocation } from "react-router-dom";
import Typography from "@mui/joy/Typography";
import Input from "@mui/joy/Input";
import Button from "@mui/joy/Button";

function Login() {
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
              <Input placeholder="Email" className="field mb-4" />
              <Input placeholder="Password" className="field" type="password" />
            </div>
          </div>
          <div className="login-field-button max-w-xs w-full">
            <button className=" text-white font-bold py-3 px-4 rounded-lg w-full button-in-login">
              Sign In
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Login;
