const ReconnectDelay: number = 100; 
const WebSocketURL: string = "ws://localhost:3000/ws";

const Warn: string = "⚠️ Something went wrong: ";
const Connected: string = "✅ WebSocket service connected";
const Closed: string = "❌ WebSocket service connect closed";

export type WebSocketMsg = {
    type: string
    content: string
}

export type ContextDependency<T> = {
    key: string
    value: T,
    setValue: React.Dispatch<React.SetStateAction<T>>,
}

export type AsashishiAgentWsDependencies = {
    uInput?: ContextDependency<string>,
    aiOutput?: ContextDependency<string>,
    shellOutput?: ContextDependency<string>,
}

class AsashishiAgentWs {

    private ws: WebSocket;
    private dependencies: AsashishiAgentWsDependencies;

    public injectDependencies(dependencies: ContextDependency<any>[]): void {
        for (const item of dependencies) {
            this.dependencies[item.key as keyof AsashishiAgentWsDependencies] = item;
        }
    }

    public constructor()  {
        this.dependencies = {};
        this.ws = new WebSocket(WebSocketURL);
        this.ws.onopen = () => console.log(Connected);
        this.ws.onclose = () => {
            console.log(Closed);
            setTimeout(() => this.ws = new WebSocket(WebSocketURL), ReconnectDelay);
        };
        this.ws.onerror = (error: unknown) => {
            this.ws.close();
            console.warn(Warn + (error as Error).message);
        }
        this.ws.onmessage = (event: MessageEvent<string>) => {
          if (this.dependencies.uInput?.value != "") {
            this.dependencies.uInput?.setValue("");
          }
          this.disPlayMsg(JSON.parse(event.data) as WebSocketMsg);
        };
    }

    public send(msg: WebSocketMsg) {
        this.ws.send(JSON.stringify(msg));
    }

    public disPlayMsg(msg: WebSocketMsg): void {
        if (msg.type === "ai_msg") {
            this.dependencies.aiOutput?.setValue((prev: string): string => prev + msg.content);
        } else {
            this.dependencies.shellOutput?.setValue((prev: string): string => prev + msg.content + '\n');
        }
    }
}

// single instance
const wsInstance: AsashishiAgentWs = new AsashishiAgentWs();

export default wsInstance;