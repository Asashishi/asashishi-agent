import type { ContextStorageItem } from "./context_storage";

export type WebSocketMsg = {
    type: string
    content: string
}

const WebSocketURL: string = "ws://localhost:3000/ws";

const Warn: string = "⚠️ Something went wrong: ";
const Connected: string = "✅ WebSocket service connected";
const Closed: string = "❌ WebSocket service connect closed";

class AsashishiAgentWs {

    private ws: WebSocket;
    private uInputState: ContextStorageItem<string> | undefined;
    private aiOutputState: ContextStorageItem<string> | undefined;
    private scpOutputState: ContextStorageItem<string> | undefined;

    public injectContextItems(
        uInputState: ContextStorageItem<string>,
        aiOutputState: ContextStorageItem<string>,
        scpOutputState: ContextStorageItem<string>,
    ): void {
        this.uInputState = uInputState;
        this.aiOutputState = aiOutputState;
        this.scpOutputState = scpOutputState;
    }

    public constructor()  {
        this.ws = new WebSocket(WebSocketURL);
        this.ws.onopen = () => console.log(Connected);
        this.ws.onclose = () => {
            console.log(Closed);
            this.ws = new WebSocket(WebSocketURL);
        };
        this.ws.onerror = (error: unknown) => {
            this.ws.close();
            console.warn(Warn + (error as Error).message);
        }
        this.ws.onmessage = (event: MessageEvent<string>) => {
          if (this.uInputState?.value != "") {
            this.uInputState?.setValue("");
          }
          this.disPlayMsg(JSON.parse(event.data) as WebSocketMsg);
        };
    }

    public send(msg: WebSocketMsg) {
        this.ws.send(JSON.stringify(msg));
    }

    public disPlayMsg(msg: WebSocketMsg): void {
        if (msg.type === "ai_msg") {
            this.aiOutputState?.setValue(this.aiOutputState.value + msg.content);
        } else {
            this.scpOutputState?.setValue(this.scpOutputState.value + msg.content + '\n');
        }
    }
}

// single instance
const wsInstance: AsashishiAgentWs = new AsashishiAgentWs();

export default wsInstance;