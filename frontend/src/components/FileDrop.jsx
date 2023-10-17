import React from "react";
import { Typography } from "@mui/joy";
import Button from "@mui/joy/Button";
import Input from "@mui/joy/Input";
import Switch from "@mui/joy/Switch";
import Divider from "@mui/joy/Divider";

import { useState } from "react";

function FileDrop() {
  const [showElement, setShowElement] = useState(true);
  const [afterUploadElements, setAfterUploadElements] = useState(false);
  const [file, setFile] = useState(null);
  const handleFileChange = (event) => {
    const selectedFile = event.target.files[0];
    if (selectedFile) {
      setFile(selectedFile);
      setShowElement(false);
      setAfterUploadElements(true);
      const fileType = selectedFile.type;
    }
  };

  return (
    <div className="left  flex flex-col ">
      <Typography className="ll" level="body-lg">
        From
      </Typography>
      <Input size="md" placeholder="Your Email" className="mb-4 abc" />
      <div className="sheet p-4 flex flex-row items-center mb-5 justify-around">
        <div className="left-sheet flex gap-0.5 flex-col">
          <Typography level="body-md">Password Protection</Typography>
          <Typography className="sss" maxWidth={200} level="body-xs">
            When enabled, it creates a password.
          </Typography>
        </div>
        <div className="right-sheet">
          <Switch className="switch" size="lg" />
        </div>
      </div>
      <Divider />
      <Typography className="text-center a" maxWidth={300} level="body-xs">
        By clicking 'Create Secure Link' you agree to the Terms of Service and
        Privacy & Cookie Notice.
      </Typography>
      <Button className="" disabled size="lg">
        Create Secure Link
      </Button>
    </div>
  );
}

export default FileDrop;
