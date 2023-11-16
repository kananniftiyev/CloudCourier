import { useParams } from "react-router-dom";
import React, { useState, useEffect } from "react";
import Input from "@mui/joy/Input";
import Typography from "@mui/joy/Typography";
import Button from "@mui/joy/Button";
import Alert from "@mui/joy/Alert";
import Box from "@mui/joy/Box";
import { motion, AnimatePresence } from "framer-motion";
import PlaylistAddCheckCircleRoundedIcon from "@mui/icons-material/PlaylistAddCheckCircleRounded";
import DownloadingIcon from "@mui/icons-material/Downloading";
import VisibilityIcon from "@mui/icons-material/Visibility";
import DeleteIcon from "@mui/icons-material/Delete";
function File() {
  const { uuid } = useParams();
  const [currentLink, setCurrentLink] = useState("");
  const [showAlert, setShowAlert] = useState(false);

  useEffect(() => {
    // Update the currentLink state with the current URL when the component mounts
    setCurrentLink(window.location.href);
  }, []);

  const copyLink = () => {
    navigator.clipboard
      .writeText(currentLink)
      .then(() => {
        console.log("Link copied to clipboard");
        // Show the alert with animation
        setShowAlert(true);
        // Hide the alert after a certain duration (adjust as needed)
        setTimeout(() => {
          setShowAlert(false);
        }, 3000);
      })
      .catch((err) => {
        console.error("Unable to copy link to clipboard", err);
      });
  };

  const AlertMessage = () => {
    return (
      <AnimatePresence>
        {showAlert && (
          <Box
            component={motion.div}
            initial={{ opacity: 0, y: -20 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: -20 }}
            transition={{ duration: 0.5 }}
            sx={{
              display: "flex",
              flexDirection: "column",
              gap: 2,
              width: "50%",
              position: "absolute",
              top: "1.5%",
              width: "50%",
              zIndex: 1000, // Adjust the zIndex as needed
            }}
          >
            <Alert
              startDecorator={<PlaylistAddCheckCircleRoundedIcon />}
              variant="soft"
              color="success"
            >
              Link copied successfully.
            </Alert>
          </Box>
        )}
      </AnimatePresence>
    );
  };

  return (
    <>
      <div className="logo-file px-36 pt-4 absolute">
        <div className="logo">
          <Typography className="logo-main" level="h3">
            CloudShare<span className="x-home">X</span>
          </Typography>
        </div>
      </div>
      <div className="file-container flex flex-col items-center justify-center">
        <AlertMessage />
      </div>
      <div className="file-layout flex flex-col justify-center py-36 px-36">
        <div className="file-layout-top">
          <div className="top-texts flex flex-col gap-1">
            <Typography level="h1" className="white">
              File Name
            </Typography>
            <Typography level="body-s grey">
              1 file · 12 MB · Sent 36 seconds ago
            </Typography>
            <div className="divider mt-6"></div>
          </div>
        </div>
        <div className="file-layout-mid flex flex-row justify-between items-center mt-4 mb-4">
          <Input
            endDecorator={<Button onClick={copyLink}>Copy link</Button>}
            value={currentLink}
            sx={{
              "--Input-radius": "14px",
              "--Input-minHeight": "30px",
            }}
          />

          <div className="mid-right flex flex-row gap-8">
            <div className="download blue">
              <a href="" className="flex-col flex items-center gap-1">
                <DownloadingIcon fontSize="large" className="icon" />
                <Typography className="text" level="body-sm">
                  Download
                </Typography>
              </a>
            </div>
            <div className="preview disabled ">
              <a href="#" className="flex-col gap-1 flex items-center">
                <VisibilityIcon fontSize="large" className="" />
                <Typography className="" level="body-sm">
                  Preview
                </Typography>
              </a>
            </div>
            <div className="download">
              <a href="" className="flex-col gap-1 flex items-center">
                <DeleteIcon fontSize="large" className="icon" />
                <Typography className="text" level="body-sm">
                  Delete
                </Typography>
              </a>
            </div>
          </div>
        </div>
        <div className="file-layout-bottom">
          <div className="divider"></div>
          <div className="bottom-layout mt-5">
            <div className="bottom-left flex flex-col gap-5">
              <div className="flex flex-col gap-0.5">
                <Typography className="white" level="h4">
                  Expiration date
                </Typography>
                <Typography className="grey" level="body-sm">
                  November 23, 2023
                </Typography>
              </div>
              <div className="flex flex-col gap-0.5">
                <Typography className="white mb-1" level="h4">
                  Password
                </Typography>
                <Typography className="grey" level="body-sm">
                  xxxxxx
                </Typography>
              </div>
              <div>
                <Typography className="white mb-1" level="h4">
                  Total Downloads
                </Typography>
                <Typography className="grey" level="h1">
                  0
                </Typography>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default File;
