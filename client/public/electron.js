const { app, BrowserWindow } = require("electron");
const path = require("path");

function createWindow() {
  // Create the browser window.
  const win = new BrowserWindow({
    width: 1024,
    height: 768,
    webPreferences: {
      nodeIntegration: true,
    },
  });

  let pth;

  if (!app.isPackaged) {
    pth = "http://localhost:3000";
  } else {
    pth = `file://${path.join(__dirname, "../build/index.html")}`;
  }

  win
    .loadURL(pth)
    .then()
    .catch((e) => {
      console.error("loadUrl error: " + e);
    });

  //win.webContents.openDevTools();
}

app.whenReady().then(createWindow);

app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    app.quit();
  }
});

app.on("activate", () => {
  if (BrowserWindow.getAllWindows().length === 0) {
    createWindow();
  }
});
