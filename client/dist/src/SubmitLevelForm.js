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
import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { useState } from "react";
import { ADDRESS, SUBMIT_LEVEL_ENDPOINT } from "./App";
export function SubmitLevelForm(_a) {
    var setBoard = _a.setBoard, setCurrentId = _a.setCurrentId;
    var _b = useState(""), text = _b[0], setText = _b[1];
    var handleChange = function (event) {
        var text = event.target.value;
        setText(text);
    };
    var handleSubmit = function (event) {
        event.preventDefault();
        var level;
        try {
            level = JSON.parse(text);
        }
        catch (_a) {
            console.error("invalid level: " + text);
            return;
        }
        var url = ADDRESS + SUBMIT_LEVEL_ENDPOINT;
        fetch(url, {
            method: "POST",
            body: text,
            headers: {
                "Content-type": "application/json; charset=UTF-8",
            },
        })
            .then(function (response) { return response.json(); })
            .then(function (data) {
            console.log(data);
            if (data.id) {
                setBoard(level);
                setCurrentId(data.id);
            }
            else {
                console.error("newBoard error");
            }
        })
            .catch(function (err) {
            console.log(err.message);
        });
    };
    return (_jsxs("form", __assign({ onSubmit: handleSubmit }, { children: [_jsx("label", { children: _jsx("input", { placeholder: " [[0,0,0,0,2],[0,0,4,0,2],[0,1,2,0,0],[0,1,1,3,0],[0,0,0,0,0]]", value: text, onChange: handleChange }) }), _jsx("button", __assign({ type: "submit" }, { children: "Load" }))] })));
}
