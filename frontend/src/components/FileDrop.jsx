import { Typography } from "@mui/joy";
import Button from "@mui/joy/Button";
// TODO: Add Functionality
function FileDrop() {
  return (
    <div className="file-drop mt-16 w-2/4 mx-auto">
      <div className="md:container md:mx-auto back ">
        <label className="flex flex-col items-center justify-center gap-5 p-10 border-4 border-dotted  border-gray-200 back">
          <Typography level="h3">Drag a File Here</Typography>
          <span className="or">---Or---</span>
          <input className="input-file" type="file" id="file" />
          <label className="" htmlFor="file">
            Browse Disk
          </label>
        </label>
      </div>
    </div>
  );
}

export default FileDrop;
