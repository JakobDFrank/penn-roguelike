import React from "react";
import { ADDRESS, PLAYER_MOVE_ENDPOINT } from "./App";
import { HistoryLogger } from "./HistoryLogger";

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
  logger: HistoryLogger,
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

  logger.log(req);

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
      logger.log(data);

      if (data.level) {
        setNums(data.level);
      } else {
        logger.error("newBoard error");
      }
    })
    .catch((err) => {
      logger.log(err.message);
    });
}
