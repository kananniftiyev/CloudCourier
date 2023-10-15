import { Text } from "@chakra-ui/react";
import { Button, ButtonGroup } from "@chakra-ui/react";
import { Link } from "@chakra-ui/react";

function Navigation() {
  return (
    <header>
      <div className="md:container md:mx-auto mt-4">
        <div className="flex justify-between">
          <div className="logo">
            <Text className="logo-main" as="b" fontSize="2xl">
              CloudShare<span className="x-logo">X</span>
            </Text>
          </div>
          <div className="auth flex justify-between gap-6 items-center">
            <Link as="b" className="button-login">
              Login
            </Link>
            <Button className="button-signup" size="md">
              Sign Up
            </Button>
          </div>
        </div>
      </div>
    </header>
  );
}

export default Navigation;
