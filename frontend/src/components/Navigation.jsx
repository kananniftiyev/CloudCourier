import React from "react";
import Typography from "@mui/joy/Typography";
import Button from "@mui/joy/Button";
import Link from "@mui/joy/Link";
import Drawer from "@mui/joy/Drawer";
import Input from "@mui/joy/Input";
import { BsShieldLockFill } from "react-icons/bs";
import Checkbox from "@mui/joy/Checkbox";

function Navigation() {
  const [open, setOpen] = React.useState(false);
  const [openSign, setOpenSign] = React.useState(false);

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
              <Input className="mb-4" placeholder="Username" />
              <Input className="mb-4" placeholder="Email" />
              <Input className="mb-4" type="password" placeholder="Password" />
              <Input
                className="mb-8"
                type="password"
                placeholder="Password Again"
              />
              <Button size="lg" disabled>
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
