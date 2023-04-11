export class WorkerPool {
    private workers: Worker[] = [];
    private queue: ((event: MessageEvent) => void)[] = [];
    private maxWorkers: number;
  
    constructor(maxWorkers: number, workerUrl: string) {
      this.maxWorkers = maxWorkers;
  
      for (let i = 0; i < maxWorkers; i++) {
        const worker = new Worker(workerUrl, { type: "module" });
        worker.onmessage = (event) => {
          if (this.queue.length > 0) {
            const callback = this.queue.shift();
            if (callback) {
              callback(event);
            }
          } else {
            this.workers.push(worker);
          }
        };
        this.workers.push(worker);
      }
    }
  
    async run(data: any): Promise<any> {
      const worker = this.workers.pop();
  
      if (worker) {
        worker.postMessage(data);
        return new Promise((resolve) => {
          worker.onmessage = (event) => {
            resolve(event.data);
            this.workers.push(worker);
          };
        });
      } else {
        return new Promise((resolve) => {
          this.queue.push((event) => {
            resolve(event.data);
          });
        });
      }
    }
  }
  