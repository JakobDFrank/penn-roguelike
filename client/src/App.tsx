import React, { useEffect, useRef, useState } from "react";
import "./App.css";
import { Board } from "./Board";
import { SubmitLevelForm } from "./SubmitLevelForm";
import { Direction, MovePlayerInput } from "./MovePlayerInput";
import { ConsoleMessages } from "./ConsoleMessages";
import { HistoryLogger, LogMessage } from "./HistoryLogger";

export const ADDRESS = "http://127.0.0.1:8080";
export const SUBMIT_LEVEL_ENDPOINT: string = "/level/submit";
export const PLAYER_MOVE_ENDPOINT: string = "/player/move";

function App() {
  const loggerRef = useRef(new HistoryLogger(10));

  if (!loggerRef.current) {
    loggerRef.current = new HistoryLogger(10);
  }

  const logger = loggerRef.current;

  const init: number[][] = [];
  const [board, setBoard] = useState(init);
  const [id, setCurrentId] = useState(0);

  const initMessages: LogMessage[] = [];
  const [messages, setMessages] = useState(initMessages);

  useEffect(() => {
    const l = loggerRef.current;
    const updateMessages = () => {
      const arr = l.toArray();
      setMessages(arr);
    };

    l.on("stateChanged", updateMessages);

    // unsubscribe when the component unmounts
    return () => {
      l.off("stateChanged", updateMessages);
    };
  }, []); // empty array ensures this effect runs only once

  const keyUpHandler = (event: React.KeyboardEvent<HTMLInputElement>) => {
    event.preventDefault();

    switch (event.code) {
      case "KeyA":
      case "ArrowLeft":
        MovePlayerInput(setBoard, id, Direction.Left, logger);
        break;
      case "KeyD":
      case "ArrowRight":
        MovePlayerInput(setBoard, id, Direction.Right, logger);
        break;
      case "KeyW":
      case "ArrowUp":
        MovePlayerInput(setBoard, id, Direction.Up, logger);
        break;
      case "KeyS":
      case "ArrowDown":
        MovePlayerInput(setBoard, id, Direction.Down, logger);
        break;
    }
  };

  return (
    <>
      <div className="container" onKeyUp={keyUpHandler}>
        <Board id={id} cells={board} />
      </div>
      <div className="under-grid">
        <SubmitLevelForm
          setBoard={setBoard}
          setCurrentId={setCurrentId}
          logger={logger}
        />
        Level ID: {id}
        <br></br>
        <div className="log-container">
          <ConsoleMessages messages={messages} />
        </div>
        <div className="example-level-container">
          Example Level:
          [[0,0,0,0,2],[0,0,4,0,2],[0,1,2,0,0],[0,1,1,3,0],[0,0,0,0,0]]
        </div>
      </div>
    </>
  );
}

export default App;
