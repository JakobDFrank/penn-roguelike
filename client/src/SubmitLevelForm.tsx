import React, { useState } from "react";
import { ADDRESS, SUBMIT_LEVEL_ENDPOINT } from "./App";

interface SubmitLevelFormProps {
  setBoard: React.Dispatch<React.SetStateAction<number[][]>>;
  setCurrentId: React.Dispatch<React.SetStateAction<number>>;
}

export function SubmitLevelForm({
  setBoard,
  setCurrentId,
}: SubmitLevelFormProps) {
  const [text, setText] = useState("");

  const handleChange = (event: any) => {
    const text = event.target.value;
    setText(text);
  };

  const handleSubmit = (event: any) => {
    event.preventDefault();

    let level: number[][];

    try {
      level = JSON.parse(text);
    } catch {
      console.error("invalid level: " + text);
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
        console.log(data);

        if (data.id) {
          setBoard(level);
          setCurrentId(data.id);
        } else {
          console.error("newBoard error");
        }
      })
      .catch((err) => {
        console.log(err.message);
      });
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        <input
          placeholder=" [[0,0,0,0,2],[0,0,4,0,2],[0,1,2,0,0],[0,1,1,3,0],[0,0,0,0,0]]"
          value={text}
          onChange={handleChange}
        />
      </label>
      <button type="submit">Load</button>
    </form>
  );
}
