import type { ContextStorageItem, WebSocketMsg } from '../types/index';

const WebSocketURL: string = "ws://localhost:3000/ws";

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
        this.ws.onopen = () => console.log("✅ 已连接到 WebSocket 服务");
        this.ws.onclose = () => {
            console.log("\n ❌ 连接已关闭");
            this.ws = new WebSocket(WebSocketURL);
        };
        this.ws.onerror = (error: unknown) => {
            this.ws.close();
            console.error("\n ⚠️ 出错: " + (error as Error).message);
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