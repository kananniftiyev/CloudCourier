import React, { useState } from "react";
import Typography from "@mui/joy/Typography";
import Button from "@mui/joy/Button";
import Link from "@mui/joy/Link";
import Drawer from "@mui/joy/Drawer";
import Input from "@mui/joy/Input";
import { BsShieldLockFill } from "react-icons/bs";
import Checkbox from "@mui/joy/Checkbox";

function SignInForm({ open, onClose }) {
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
    <Drawer
      className="drawer-sign flex flex-col"
      open={open}
      onClose={onClose}
      anchor="right"
    >
      <div className="flex flex-row justify-between mb-10">
        <Typography level="h3">Login</Typography>
        <div className="flex flex-row items-center gap-1">
          <BsShieldLockFill className="lock" />
          <Typography level="body-sm">
            Accounts are secure and encrypted
          </Typography>
        </div>
      </div>

      <Input
        className="mb-4"
        placeholder="Email"
        name="email"
        onChange={handleChange}
        value={formData.email}
        error={emailError}
      />
      <Input
        type="password"
        placeholder="Password"
        name="password"
        onChange={handleChange}
        value={formData.password}
      />
      <div className="flex flex-row justify-between mt-8 mb-6">
        <Checkbox label="Remember me" />
        <Link>Forgot your password?</Link>
      </div>
      <Button
        size="lg"
        disabled={!isFormFilled || emailError}
        onClick={handleSignIn}
      >
        Login
      </Button>
    </Drawer>
  );
}

function SignUpForm({ open, onClose }) {
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
    <Drawer
      className="drawer-create"
      open={open}
      onClose={onClose}
      anchor="right"
    >
      <div className="flex flex-row justify-between mb-10">
        <Typography level="h3">Sign up</Typography>
        <div className="flex flex-row items-center gap-1">
          <BsShieldLockFill className="lock" />
          <Typography level="body-sm">
            Accounts are secure and encrypted
          </Typography>
        </div>
      </div>
      <Input
        className="mb-4"
        placeholder="Username"
        name="username"
        value={formData.username}
        onChange={handleChange}
      />
      <Input
        className="mb-4"
        placeholder="Email"
        name="email"
        value={formData.email}
        onChange={handleChange}
        type="email"
        error={emailError}
      />
      <Input
        className="mb-4"
        type="password"
        placeholder="Password"
        name="password"
        value={formData.password}
        onChange={handleChange}
      />
      <Input
        className="mb-8"
        type="password"
        placeholder="Password Again"
        name="confirmPassword"
        value={formData.confirmPassword}
        error={passwordError}
        onChange={handleChange}
      />
      <Button
        size="lg"
        disabled={!isFormFilled || emailError || passwordError}
        onClick={handleSignUp}
      >
        Sign up
      </Button>
    </Drawer>
  );
}

function Navigation() {
  const [isSignInDrawerOpen, setSignInDrawerOpen] = useState(false);
  const [isSignUpDrawerOpen, setSignUpDrawerOpen] = useState(false);

  const toggleSignInDrawer = (open) => () => {
    setSignInDrawerOpen(open);
  };

  const toggleSignUpDrawer = (open) => () => {
    setSignUpDrawerOpen(open);
  };

  return (
    <header>
      <div className="md:container md:mx-auto mt-4">
        <div className="flex justify-between">
          <div className="logo">
            <Typography className="logo-main" level="h3">
              CloudShare<span className="x-logo">X</span>
            </Typography>
          </div>
          <div className="auth flex justify-between gap-6 items-center">
            <Link
              underline="none"
              className="button-login"
              onClick={toggleSignInDrawer(true)}
            >
              Sign in
            </Link>
            <Button
              className="button-signup"
              size="lg"
              onClick={toggleSignUpDrawer(true)}
            >
              Create an Account
            </Button>
          </div>
        </div>
      </div>
      <SignInForm
        open={isSignInDrawerOpen}
        onClose={toggleSignInDrawer(false)}
      />
      <SignUpForm
        open={isSignUpDrawerOpen}
        onClose={toggleSignUpDrawer(false)}
      />
    </header>
  );
}

export default Navigation;
