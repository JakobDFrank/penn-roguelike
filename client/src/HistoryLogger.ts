import { CircularQueue } from "./CircularQueue";
import EventEmitter from "events";

export enum ErrorKind {
  Log,
  Error,
}

export class HistoryLogger extends EventEmitter {
  private history: CircularQueue<LogMessage>;

  constructor(size: number) {
    super();
    this.history = new CircularQueue(size);
  }

  log(...data: any[]) {
    console.log(...data);

    this.history.enqueue({
      type: ErrorKind.Log,
      data: data,
    });

    this.emit("stateChanged", null);
  }

  error(...data: any[]) {
    console.error(...data);

    this.history.enqueue({
      type: ErrorKind.Error,
      data: data,
    });

    this.emit("stateChanged", null);
  }

  toArray(): LogMessage[] {
    return this.history.toArray();
  }
}

export interface HistoryLogger {
  log(...data: any[]): void;
  error(...data: any[]): void;
}

export interface LogMessage {
  type: ErrorKind;
  data: any[];
}
