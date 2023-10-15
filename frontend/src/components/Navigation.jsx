import Typography from "@mui/joy/Typography";
import Button from "@mui/joy/Button";
import Link from "@mui/joy/Link";

function Navigation() {
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
            <Link underline="none" className="button-login">
              Sign in
            </Link>
            <Button className="button-signup" size="lg">
              Create a Account
            </Button>
          </div>
        </div>
      </div>
    </header>
  );
}

export default Navigation;
