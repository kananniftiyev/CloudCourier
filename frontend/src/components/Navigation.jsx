import React from "react";
import Typography from "@mui/joy/Typography";
import Button from "@mui/joy/Button";
import Link from "@mui/joy/Link";
import Drawer from "@mui/joy/Drawer";

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
              className="drawer-sign"
              open={openSign}
              onClose={toggleDrawerSignIn(false)}
              anchor="right"
            >
              <Typography>llll</Typography>
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
              <Typography>ASDD</Typography>
            </Drawer>
          </div>
        </div>
      </div>
    </header>
  );
}

export default Navigation;
