import React from "react";
import { ErrorKind, LogMessage } from "./HistoryLogger";

interface ConsoleMessagesProp {
  messages: LogMessage[];
}

export function ConsoleMessages({ messages }: ConsoleMessagesProp) {
  const logMessage = "p-1 mx-5 border-b border-gray-200";

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
          className = `${logMessage} text-red-600`;
        } else {
          className = `${logMessage} text-blue-600`;
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
