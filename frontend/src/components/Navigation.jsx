import React from "react";
import Typography from "@mui/joy/Typography";
import Button from "@mui/joy/Button";
import Link from "@mui/joy/Link";
import Drawer from "@mui/joy/Drawer";
import Input from "@mui/joy/Input";
import { BsShieldLockFill } from "react-icons/bs";
import Checkbox from "@mui/joy/Checkbox";
import { useState } from "react";

// TODO: Seperate sign up and login from nav.
// TODO: give info when wrong form fill.
function Navigation() {
  const [open, setOpen] = React.useState(false);
  const [openSign, setOpenSign] = React.useState(false);
  const [emailError, setEmailError] = useState(false);
  const [passwordError, setPasswordError] = useState(false);

  const toggleDrawer = (inOpen) => (event) => {
    if (
      event.type === "keydown" &&
      (event.key === "Tab" || event.key === "Shift")
    ) {
      return;
    }

    setOpen(inOpen);
  };

  const toggleDrawerSignIn = (inOpenSign) => (event) => {
    if (
      event.type === "keydown" &&
      (event.key === "Tab" || event.key === "Shift")
    ) {
      return;
    }

    setOpenSign(inOpenSign);
  };

  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
    passwordAgain: "",
  });

  const isFormFilled = Object.values(formData).every(
    (value) => value.trim() !== ""
  );

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });

    if (name === "email") {
      // Check email format with a regular expression
      const emailPattern = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i;
      setEmailError(!emailPattern.test(value));
    }

    if (name === "passwordAgain") {
      if (value !== formData.password) {
        setPasswordError(true);
      } else {
        setPasswordError(false);
      }
    }
  };

  const handleSign = () => {
    if (emailError) {
      return;
    }
    // You can access the form data in the `formData` state here
    console.log("Form Data:", formData);

    // You can perform any further actions here, such as sending the data to a server.
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
              onClick={toggleDrawerSignIn(true)}
            >
              Sign in
            </Link>
            <Drawer
              className="drawer-sign flex flex-col"
              open={openSign}
              onClose={toggleDrawerSignIn(false)}
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

              <Input className="mb-4" placeholder="Email" />
              <Input type="password" placeholder="Password" />
              <div className="flex flex-row justify-between mt-8 mb-6">
                <Checkbox label="Remember me" />
                <Link>Forgot your password?</Link>
              </div>
              <Button size="lg" disabled>
                Login
              </Button>
            </Drawer>
            <Button
              className="button-signup"
              size="lg"
              onClick={toggleDrawer(true)}
            >
              Create a Account
            </Button>
            <Drawer
              className="drawer-create"
              open={open}
              onClose={toggleDrawer(false)}
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
                name="passwordAgain"
                value={formData.passwordAgain}
                error={passwordError}
                onChange={handleChange}
              />
              <Button
                size="lg"
                disabled={!isFormFilled || emailError || passwordError}
                onClick={handleSign}
              >
                Sign up
              </Button>
            </Drawer>
          </div>
        </div>
      </div>
    </header>
  );
}

export default Navigation;
