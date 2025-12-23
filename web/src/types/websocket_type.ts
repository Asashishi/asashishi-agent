import type { SignalStaus } from "./signal_status_type";

export type WebSocketMsgType = string;

export type WebSocketMsg = {
    content: string,
    type: WebSocketMsgType,
}

export type ContextDependency<T> = {
    key: string
    value: T,
    setValue: React.Dispatch<React.SetStateAction<T>>,
}

export type AsashishiAgentWsDependencies = {
    aiOutput?: ContextDependency<string>,
    shellOutput?: ContextDependency<string>,
    wsSignal?: ContextDependency<SignalStaus>,
    aiOutputHistories?: ContextDependency<string>,
}