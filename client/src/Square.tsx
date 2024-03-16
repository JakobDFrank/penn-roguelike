import React from "react";

interface SquareProps {
  value: number;
}

export function Square({ value }: SquareProps) {
  return (
    <button className="bg-white border border-gray-300 text-lg font-bold leading-9 h-9 p-0 text-center w-9 hover:bg-gray-200 active:bg-gray-400 active:cursor-crosshair focus:outline-none">
      {value !== 0 ? value : ""}
    </button>
  );
}
