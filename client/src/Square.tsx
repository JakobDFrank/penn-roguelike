import React from "react";

interface SquareProps {
  value: number;
}

export function Square({ value }: SquareProps) {
  return <button className="square">{value !== 0 ? value : ""}</button>;
}
