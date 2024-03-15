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
import { jsx as _jsx } from "react/jsx-runtime";
import React from "react";
import { Square } from "./Square";
export function Board(_a) {
    var id = _a.id, cells = _a.cells;
    var board = cells.map(function (row, idx) {
        var rowHtml = row.map(function (num, jdx) {
            var key = "".concat(id, "-").concat(idx, "-").concat(jdx);
            return (_jsx(React.Fragment, { children: _jsx(Square, { value: num }) }, key));
        });
        return (_jsx("div", __assign({ className: "board-row" }, { children: rowHtml }), "".concat(id, "-row-").concat(idx)));
    });
    return _jsx("div", { children: board });
}
