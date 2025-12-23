import { GreenSignal, GreySignal, RedSignal } from "../consts/signal";
import { AIOutputEndType, AIOutputType, ChildProcessOutputType, SysErrorType, SystMsgType, SystWarnType } from '../consts/websocket';
import type { AsashishiAgentWsDependencies, ContextDependency, DisplayMsg, WebSocketMsg } from "../types/websocket_type";

const ReconnectDelay: number = 500; 
const WebSocketURL: string = "ws://localhost:3000/ws";

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
        this.dependencies.outputHistories?.setValue((prev: DisplayMsg[]): DisplayMsg[] => {
            const next: DisplayMsg[] = prev;
            next.push({
                type: "input",
                content: msg.content,
            });
            return next;
        });
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
                this.dependencies.aiOutput?.setValue((prev: string): string => prev + msg.content);
                break;
            }
            case AIOutputEndType: {
                this.dependencies.outputHistories?.setValue((prev: DisplayMsg[]): DisplayMsg[] => {
                    const next: DisplayMsg[] = prev;
                    next.push({
                        type: "output",
                        content: this.dependencies.aiOutput!.value!,
                    });
                    return next;
                });
                this.dependencies.aiOutput?.setValue("");
                break;
            }   
            case SystMsgType || SystWarnType || SysErrorType || ChildProcessOutputType: {
                this.dependencies.shellOutput?.setValue((prev: string): string => prev + msg.content + '\n');
                break;
            }
        }
    }
}

// single instance
const wsInstance: AsashishiAgentWs = new AsashishiAgentWs();

export default wsInstance;