import React from "react";
import { Typography } from "@mui/joy";
import Button from "@mui/joy/Button";
import Input from "@mui/joy/Input";
import Switch from "@mui/joy/Switch";
import Divider from "@mui/joy/Divider";
import { BsPlusLg } from "react-icons/bs";
import { BsFillTrashFill } from "react-icons/bs";
import { useEffect } from "react";

import { useState } from "react";

function FileDrop() {
  const [showElement, setShowElement] = useState(true);
  const [afterUploadElements, setAfterUploadElements] = useState(false);
  const [passowordElements, setPassowordElements] = useState(false);
  const [file, setFile] = useState(null);
  const [checked, setChecked] = useState(false);

  const formatFileSize = (size) => {
    if (size < 1024) {
      return `${size} B`;
    } else if (size < 1024 * 1024) {
      return `${(size / 1024).toFixed(0)} KB`;
    } else {
      return `${(size / (1024 * 1024)).toFixed(2)} MB`;
    }
  };

  const mapMimeTypeToDisplay = (mimeType) => {
    const mimeTypeMap = {
      "application/pdf": "pdf",
      "text/plain": "txt",
      "application/x-msdownload": "exe",
      // Add more mappings for other MIME types if needed
    };
    return mimeTypeMap[mimeType] || mimeType;
  };

  const handleFileChange = (event) => {
    const selectedFile = event.target.files[0];
    if (selectedFile) {
      setFile(selectedFile);
      setShowElement(false);
      setAfterUploadElements(true);
    }
  };

  const deleteFile = () => {
    setFile(null);
    setShowElement(true);
    setAfterUploadElements(false);
  };

  const handleSwitch = (event) => {
    const isChecked = event.target.checked;
    setChecked(isChecked);

    if (isChecked) {
      setPassowordElements(true);
    } else {
      setPassowordElements(false);
    }
  };

  useEffect(() => {
    console.log(file);
  }, [file]);

  return (
    <div className="filedrop flex flex-row gap-6">
      <div className="left  flex flex-col ">
        <Input placeholder="Title" className="mb-4" />
        <div className="sheet p-4 flex flex-row items-center mb-5 justify-around">
          <div className="left-sheet flex gap-0.5 flex-col">
            <Typography level="body-md">Password Protection</Typography>
            <Typography className="sss" maxWidth={200} level="body-xs">
              When enabled, it lets you create a password.
            </Typography>
          </div>
          <div className="right-sheet">
            <Switch
              className="switch"
              size="lg"
              onChange={handleSwitch}
              checked={checked}
            />
          </div>
        </div>
        {passowordElements ? (
          <>
            {" "}
            <Input placeholder="Password" type="password" className="mb-4" />
          </>
        ) : null}
        <Divider />
        <Typography className="text-center a" maxWidth={300} level="body-xs">
          By clicking 'Create Secure Link' you agree to the Terms of Service and
          Privacy & Cookie Notice.
        </Typography>
        <Button className="" disabled size="lg">
          Create Secure Link
        </Button>
      </div>
      <div className="right-upload p-8">
        {showElement ? (
          <>
            <label htmlFor="file">
              <div className="sheet pl-4 pr-12 add-files flex flex-row items-center justify-between gap-3">
                <BsPlusLg className="bsplus" />
                <div className="flex flex-col ">
                  <Typography level="body-lg" className="font-bold">
                    Add Files
                  </Typography>
                  <Typography level="body-xs">or Select Folder</Typography>
                </div>
              </div>
            </label>
            <input type="file" id="file" onChange={handleFileChange} />
          </>
        ) : null}

        {afterUploadElements ? (
          <>
            {file && file.name && (
              <div className="after sheet pl-4 pr-12">
                <Typography level="body-md">{file.name}</Typography>
                <Typography level="body-sm">
                  {formatFileSize(file.size)} |{" "}
                  {mapMimeTypeToDisplay(file.type)}
                </Typography>
              </div>
            )}
          </>
        ) : null}

        <div className="bottom-of flex flex-row justify-between items-center">
          <Typography>
            0 <span className="font-bold">files</span> (0 B)
          </Typography>
          <Button variant="plain" onClick={deleteFile}>
            <BsFillTrashFill />
          </Button>
        </div>
      </div>
    </div>
  );
}

export default FileDrop;
