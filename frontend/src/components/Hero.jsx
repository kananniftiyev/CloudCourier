import Typography from "@mui/joy/Typography";

function Hero() {
  return (
    <div className="hero mt-36">
      <div className="container-sm flex flex-col gap-4">
        <Typography level="h1" className="text-center hero-text">
          Send Files Securely
        </Typography>
        <Typography
          level="title-lg"
          variant="plain"
          className="text-center hero-desc"
        >
          No one can ever access your files that are sent through TransferChain
          besides the intended recipients.
        </Typography>
      </div>
    </div>
  );
}

export default Hero;
