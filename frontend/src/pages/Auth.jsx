import Typography from "@mui/joy/Typography";
import Button from "@mui/joy/Button";

function Auth() {
  return (
    <div className="layout grid grid-cols-5 h-screen">
      <div className="left-side col-span-3 px-8 py-5">
        <nav className="left-0 top-8 flex w-full px-6 sm:absolute md:top-[22px] md:px-6 lg:px-8">
          <a href="">
            <Typography className="logo-main" level="h2">
              CloudShare<span className="x-logo">X</span>
            </Typography>
          </a>
        </nav>
      </div>
      <div className="right-side col-span-2 px-5 py-8 flex flex-col items-center justify-center">
        <div className="flex grow flex-col items-center w-full max-w-[440px]">
          <div className="relative flex gap-6 w-full grow flex-col items-center justify-center">
            <Typography className="logo-main" level="h2">
              Welcome
            </Typography>
            <div className="grid gap-x-3 gap-y-2 sm:grid-cols-2 sm:gap-y-0 w-full">
              <Button className="login-button" size="lg">
                Log in
              </Button>
              <Button className="login-button" size="lg">
                Sign up
              </Button>
            </div>
          </div>
          <div style={{ marginTop: "auto" }} className="flex flex-row gap-1">
            <a href="">
              <Typography className="Terms" level="body-sm">
                Terms of use
              </Typography>
            </a>
            <span className="line">|</span>
            <a href="">
              <Typography className="Terms" level="body-sm">
                Privacy Policy
              </Typography>
            </a>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Auth;
