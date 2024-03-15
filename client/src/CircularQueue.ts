export class CircularQueue<T> {
  private readonly queue: T[];
  public readonly maxSize: number;

  constructor(maxSize: number) {
    if (maxSize < 1) {
      throw new Error("size must greater than 0");
    }

    this.maxSize = maxSize;
    this.queue = [];
  }

  enqueue(item: T) {
    if (this.queue.length > this.maxSize) {
      this.dequeue();
    }

    this.queue.push(item);
  }

  dequeue(): T {
    const item = this.queue.shift();

    if (item === undefined) {
      throw Error("queue is empty");
    }

    return item;
  }

  peek(): T {
    if (this.queue.length === 0) {
      throw Error("queue is empty");
    }

    return this.queue[0];
  }

  toArray(): T[] {
    // todo - check if primitive or not for perf

    // const array: T[] = Object.assign([], this.queue);
    // return array;

    const array: T[] = [];
    this.queue.forEach((val) => array.push(Object.assign({}, val)));
    return array;
  }
}
