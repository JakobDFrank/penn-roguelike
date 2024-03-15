import React from "react";
import { ErrorKind, LogMessage } from "./HistoryLogger";

interface ConsoleMessagesProp {
  messages: LogMessage[];
}

export function ConsoleMessages({ messages }: ConsoleMessagesProp) {
  return (
    <div>
      {messages.map((msg, index) => {
        let text;

        try {
          text = JSON.stringify(msg.data);
        } catch {
          text = "*COULD_NOT_STRINGIFY*";
        }

        let className: string;

        if (msg.type === ErrorKind.Error) {
          className = "log-message log-error";
        } else {
          className = "log-message log-info";
        }

        return (
          <div className={className} key={index}>
            {text}
          </div>
        );
      })}
    </div>
  );
}
