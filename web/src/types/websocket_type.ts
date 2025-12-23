import type { SignalStaus } from "./signal_status_type";

export type WebSocketMsg = {
    type: string;
    content: string;
}

export type ContextDependency<T> = {
    key: string;
    value: T;
    setValue: React.Dispatch<React.SetStateAction<T>>;
}

export type AsashishiAgentWsDependencies = {
    aiOutput?: ContextDependency<string>;
    shellOutput?: ContextDependency<string>;
    wsSignal?: ContextDependency<SignalStaus>;
    ioHistories?: ContextDependency<DisplayMsg[]>;
}

export type DisplayMsg = {
    content: string;
    type: "input" | "output";
}