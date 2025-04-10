@echo off
echo Cleaning up node_modules and package-lock.json
rmdir /s /q node_modules
del package-lock.json

echo Installing dependencies...
npm install

echo All dependencies installed. Starting the app...
npm start
