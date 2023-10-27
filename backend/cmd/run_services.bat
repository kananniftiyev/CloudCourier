@echo off

:: Run Auth Service in a new Command Prompt window
echo Running Auth Service
start cmd /k "cd /d C:\Users\kenan\Documents\GitHub\CloudShareX\backend\cmd\auth && go run main.go"

:: Run File Service in a new Command Prompt window
echo Running File Service
start cmd /k "cd /d C:\Users\kenan\Documents\GitHub\CloudShareX\backend\cmd\file && go run main.go"
