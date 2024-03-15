import React from "react";
import { Square } from "./Square";

interface BoardProps {
  id: number;
  cells: number[][];
}

export function Board({ id, cells }: BoardProps) {
  const board = cells.map((row: number[], idx: number) => {
    const rowHtml = row.map((num: number, jdx: number) => {
      let key: string = `${id}-${idx}-${jdx}`;

      return (
        <React.Fragment key={key}>
          <Square value={num} />
        </React.Fragment>
      );
    });

    return (
      <div key={`${id}-row-${idx}`} className="board-row">
        {rowHtml}
      </div>
    );
  });

  return <div>{board}</div>;
}
