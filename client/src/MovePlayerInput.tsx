import React from "react";
import { ADDRESS, PLAYER_MOVE_ENDPOINT } from "./App";

export enum Direction {
  Left,
  Right,
  Up,
  Down,
}

interface MovePlayerRequest {
  id: number;
  direction: number;
}

export function MovePlayerInput(
  setNums: React.Dispatch<React.SetStateAction<number[][]>>,
  id: number,
  dir: Direction,
) {
  let num;

  switch (dir) {
    case Direction.Left:
      num = 0;
      break;
    case Direction.Right:
      num = 1;
      break;
    case Direction.Up:
      num = 2;
      break;
    case Direction.Down:
      num = 3;
      break;
    default:
      return null;
  }

  const req: MovePlayerRequest = {
    id: id,
    direction: num,
  };

  console.log(req);

  const url = ADDRESS + PLAYER_MOVE_ENDPOINT;

  fetch(url, {
    method: "POST",
    body: JSON.stringify(req),
    headers: {
      "Content-type": "application/json; charset=UTF-8",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      console.log(data);

      if (data.level) {
        setNums(data.level);
      } else {
        console.error("newBoard error");
      }
    })
    .catch((err) => {
      console.log(err.message);
    });
}
