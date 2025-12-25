import { GreenSignal, GreySignal, RedSignal } from "../consts/signal";
import { AIOutputEndType, AIOutputType, ChildProcessOutputType, ReconnectDelay, WebSocketURL } from '../consts/websocket';
import type { AsashishiAgentWsDependencies, ContextDependency, DisplayMsg, WebSocketMsg } from "../types/websocket_type";

class AsashishiAgentWs {
    private ws: WebSocket;
    private dependencies: AsashishiAgentWsDependencies;

    public constructor()  {
        this.dependencies = {};
        this.dependencies.wsSignal?.setValue({
            strength: 4,
            color: GreySignal,
        });
        this.ws = this.connect();
    }

    private connect(): WebSocket {
        this.ws = new WebSocket(WebSocketURL);
        this.ws.onopen = () => {
            this.dependencies.wsSignal?.setValue({
                strength: 4,
                color: GreenSignal,
            })
        }
        this.ws.onclose = (): void => {
            this.dependencies.wsSignal?.setValue({
                strength: 1,
                color: RedSignal,
            });
            setTimeout(() => this.connect(), ReconnectDelay);
        };
        this.ws.onerror = (error: unknown): void => {
            this.ws.close();
            console.warn((error as Error).message);
        }
        this.ws.onmessage = (event: MessageEvent<string>) => {
          // 改禁止发送按钮
          this.disPlayMsg(JSON.parse(event.data) as WebSocketMsg);
        };
        return this.ws;
    }

    public send(msg: WebSocketMsg): void {
        this.ws.send(JSON.stringify(msg));
    }

    public reconnect = (): void => {
        this.ws.close();
    }

    public injectDependencies(dependencies: ContextDependency<any>[]): void {
        for (const item of dependencies) {
            this.dependencies[item.key as keyof AsashishiAgentWsDependencies] = item;
        }
    }

    public disPlayMsg(msg: WebSocketMsg): void {
        switch (msg.type) {
            case AIOutputType: {
                if (this.dependencies.tab?.value !== "chat") {
                    this.dependencies.tab?.setValue("chat");
                }
                this.dependencies.aiOutput?.setValue((prev: string): string => prev + msg.content);
                break;
            }
            case AIOutputEndType: {
                this.dependencies.ioHistories?.setValue((prev: DisplayMsg[]): DisplayMsg[] => {
                    const next: DisplayMsg[] = prev;
                    next.push({
                        type: "output",
                        diplayPosition: "chat",
                        content: this.dependencies.aiOutput!.value!,
                    });
                    return next;
                });
                this.dependencies.aiOutput?.setValue("");
                break;
            }   
            case ChildProcessOutputType: {
                if (this.dependencies.tab?.value !== "shell") {
                    this.dependencies.tab?.setValue("shell");
                }
                this.dependencies.shellOutput?.setValue((prev: string): string => prev + msg.content);
                break;
            }
        }
    }
}

// single instance
const wsInstance: AsashishiAgentWs = new AsashishiAgentWs();

export default wsInstance;