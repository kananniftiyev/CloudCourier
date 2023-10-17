import Link from "@mui/joy/Link";
import Typography from "@mui/joy/Typography";

function Footer() {
  return (
    <footer>
      <div className="md:container md:mx-auto pb-5">
        <div className="content flex justify-between">
          <div className="left-foot">
            <Typography>2023 CloudShareX. Warsaw. Poland</Typography>
          </div>
          <div className="right-foot flex items-center gap-4">
            <Link>Support</Link>
            <Link>Terms of Service</Link>
            <Link>Privacy Policy</Link>
          </div>
        </div>
      </div>
    </footer>
  );
}

export default Footer;
