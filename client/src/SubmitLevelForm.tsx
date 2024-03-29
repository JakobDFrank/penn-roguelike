import React, { useState } from "react";
import { ADDRESS, SUBMIT_LEVEL_ENDPOINT } from "./App";
import { HistoryLogger } from "./HistoryLogger";

interface SubmitLevelFormProps {
  setBoard: React.Dispatch<React.SetStateAction<number[][]>>;
  setCurrentId: React.Dispatch<React.SetStateAction<number>>;
  logger: HistoryLogger;
}

export function SubmitLevelForm({
  setBoard,
  setCurrentId,
  logger,
}: SubmitLevelFormProps) {
  const [text, setText] = useState("");

  function handleChange(event: any) {
    const text = event.target.value;
    setText(text);
  }

  function handleSubmit(event: any) {
    event.preventDefault();

    let level: number[][];

    try {
      level = JSON.parse(text);
    } catch {
      logger.error("invalid level: " + text);
      return;
    }

    const url = ADDRESS + SUBMIT_LEVEL_ENDPOINT;

    fetch(url, {
      method: "POST",
      body: text,
      headers: {
        "Content-type": "application/json; charset=UTF-8",
      },
    })
      .then((response) => response.json())
      .then((data) => {
        logger.log(data);

        if (data.id) {
          setBoard(level);
          setCurrentId(data.id);
        } else {
          logger.error("newBoard error");
        }
      })
      .catch((err) => {
        logger.log(err.message);
      });
  }

  return (
    <form onSubmit={handleSubmit}>
      <label>
        <input
          className="mr-1 p-2.5 border border-gray-300 rounded w-1/2 text-center"
          placeholder=" [[0,0,0,0,2],[0,0,4,0,2],[0,1,2,0,0],[0,1,1,3,0],[0,0,0,0,0]]"
          value={text}
          onChange={handleChange}
        />
      </label>
      <button
        type="submit"
        className="p-2.5 border border-transparent rounded bg-green-500 hover:bg-green-600 cursor-pointer"
      >
        Load
      </button>
    </form>
  );
}
