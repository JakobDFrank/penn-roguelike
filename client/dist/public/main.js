"use strict";
var _a = require('electron'), app = _a.app, BrowserWindow = _a.BrowserWindow;
function createWindow() {
    // Create the browser window.
    var win = new BrowserWindow({
        width: 800,
        height: 600,
        webPreferences: {
            nodeIntegration: true
        }
    });
    //load the index.html from React's local server
    win.loadURL('http://localhost:3000');
    win.webContents.openDevTools();
}
app.whenReady().then(createWindow);
app.on('window-all-closed', function () {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});
app.on('activate', function () {
    if (BrowserWindow.getAllWindows().length === 0) {
        createWindow();
    }
});
