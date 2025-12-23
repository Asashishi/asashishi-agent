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
    tab?: ContextDependency<Tab>;
    aiOutput?: ContextDependency<string>;
    shellOutput?: ContextDependency<string>;
    wsSignal?: ContextDependency<SignalStaus>;
    ioHistories?: ContextDependency<DisplayMsg[]>;
}

export type Tab = "chat" | "shell";

export type DisplayMsg = {
    content: string;
    type: "input" | "output";
    diplayPosition: Tab;
}