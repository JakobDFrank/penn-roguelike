var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
import { jsx as _jsx, jsxs as _jsxs, Fragment as _Fragment } from "react/jsx-runtime";
import { useState } from "react";
import "./App.css";
import { Board } from "./Board";
import { SubmitLevelForm } from "./SubmitLevelForm";
import { Direction, MovePlayerInput } from "./MovePlayerInput";
export var ADDRESS = "http://127.0.0.1:8080";
export var SUBMIT_LEVEL_ENDPOINT = "/level/submit";
export var PLAYER_MOVE_ENDPOINT = "/player/move";
function App() {
    var init = [];
    var _a = useState(init), board = _a[0], setBoard = _a[1];
    var _b = useState(0), id = _b[0], setCurrentId = _b[1];
    var keyUpHandler = function (event) {
        event.preventDefault();
        switch (event.code) {
            case "KeyA":
            case "ArrowLeft":
                MovePlayerInput(setBoard, id, Direction.Left);
                break;
            case "KeyD":
            case "ArrowRight":
                MovePlayerInput(setBoard, id, Direction.Right);
                break;
            case "KeyW":
            case "ArrowUp":
                MovePlayerInput(setBoard, id, Direction.Up);
                break;
            case "KeyS":
            case "ArrowDown":
                MovePlayerInput(setBoard, id, Direction.Down);
                break;
        }
    };
    return (_jsxs(_Fragment, { children: [_jsx("div", __assign({ className: "container", onKeyUp: keyUpHandler }, { children: _jsx(Board, { id: id, cells: board }) })), _jsxs("div", __assign({ className: "under-grid" }, { children: [_jsx(SubmitLevelForm, { setBoard: setBoard, setCurrentId: setCurrentId }), "Level ID: ", id, _jsx("br", {}), _jsx("div", __assign({ className: "example-level-container" }, { children: "Example Level: [[0,0,0,0,2],[0,0,4,0,2],[0,1,2,0,0],[0,1,1,3,0],[0,0,0,0,0]]" }))] }))] }));
}
export default App;
