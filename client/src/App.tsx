import React, { useState } from "react";
import "./App.css";
import { Board } from "./Board";
import { SubmitLevelForm } from "./SubmitLevelForm";
import { Direction, MovePlayerInput } from "./MovePlayerInput";

export const ADDRESS = "http://127.0.0.1:8080";
export const SUBMIT_LEVEL_ENDPOINT: string = "/level/submit";
export const PLAYER_MOVE_ENDPOINT: string = "/player/move";

function App() {
  let init: number[][] = [];
  const [board, setBoard] = useState(init);
  const [id, setCurrentId] = useState(0);

  const keyUpHandler = (event: React.KeyboardEvent<HTMLInputElement>) => {
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

  return (
    <>
      <div className="container" onKeyUp={keyUpHandler}>
        <Board id={id} cells={board} />
      </div>
      <div className="under-grid">
        <SubmitLevelForm setBoard={setBoard} setCurrentId={setCurrentId} />
        Level ID: {id}
        <br></br>
        <div className="example-level-container">
          Example Level:
          [[0,0,0,0,2],[0,0,4,0,2],[0,1,2,0,0],[0,1,1,3,0],[0,0,0,0,0]]
        </div>
      </div>
    </>
  );
}

export default App;
